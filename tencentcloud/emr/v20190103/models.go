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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddUsersForUserManagerRequestParams struct {
	// Cluster string ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// User information list
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`
}

type AddUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// Cluster string ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// User information list
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`
}

func (r *AddUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserManagerUserList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersForUserManagerResponseParams struct {
	// The user list that is successfully added
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SuccessUserList []*string `json:"SuccessUserList,omitnil" name:"SuccessUserList"`

	// The user list that is not successfully added
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedUserList []*string `json:"FailedUserList,omitnil" name:"FailedUserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *AddUsersForUserManagerResponseParams `json:"Response"`
}

func (r *AddUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllNodeResourceSpec struct {
	// The description of master nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterResourceSpec *NodeResourceSpec `json:"MasterResourceSpec,omitnil" name:"MasterResourceSpec"`

	// The description of core nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoreResourceSpec *NodeResourceSpec `json:"CoreResourceSpec,omitnil" name:"CoreResourceSpec"`

	// The description of task nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskResourceSpec *NodeResourceSpec `json:"TaskResourceSpec,omitnil" name:"TaskResourceSpec"`

	// The description of common nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommonResourceSpec *NodeResourceSpec `json:"CommonResourceSpec,omitnil" name:"CommonResourceSpec"`

	// The number of master nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterCount *int64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// The number of core nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoreCount *int64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// The number of task nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskCount *int64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// The number of common nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CommonCount *int64 `json:"CommonCount,omitnil" name:"CommonCount"`
}

type ApplicationStatics struct {
	// Queue name
	Queue *string `json:"Queue,omitnil" name:"Queue"`

	// Username
	User *string `json:"User,omitnil" name:"User"`

	// Application type
	ApplicationType *string `json:"ApplicationType,omitnil" name:"ApplicationType"`

	// `SumMemorySeconds` meaning
	SumMemorySeconds *int64 `json:"SumMemorySeconds,omitnil" name:"SumMemorySeconds"`


	SumVCoreSeconds *int64 `json:"SumVCoreSeconds,omitnil" name:"SumVCoreSeconds"`

	// SumHDFSBytesWritten (with unit)
	SumHDFSBytesWritten *string `json:"SumHDFSBytesWritten,omitnil" name:"SumHDFSBytesWritten"`

	// SumHDFSBytesRead (with unit)
	SumHDFSBytesRead *string `json:"SumHDFSBytesRead,omitnil" name:"SumHDFSBytesRead"`

	// Application count
	CountApps *int64 `json:"CountApps,omitnil" name:"CountApps"`
}

type COSSettings struct {
	// COS `SecretId`
	CosSecretId *string `json:"CosSecretId,omitnil" name:"CosSecretId"`

	// COS `SecrectKey`
	CosSecretKey *string `json:"CosSecretKey,omitnil" name:"CosSecretKey"`

	// COS path to log
	LogOnCosPath *string `json:"LogOnCosPath,omitnil" name:"LogOnCosPath"`
}

type CdbInfo struct {
	// Database instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Database IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Database port
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// Database memory specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// Database disk specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// Service flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Service *string `json:"Service,omitnil" name:"Service"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplyTime *string `json:"ApplyTime,omitnil" name:"ApplyTime"`

	// Payment type
	// Note: this field may return null, indicating that no valid values can be obtained.
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// Expiration flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireFlag *bool `json:"ExpireFlag,omitnil" name:"ExpireFlag"`

	// Database status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Renewal flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`

	// Database string
	// Note: this field may return null, indicating that no valid values can be obtained.
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// ZoneId
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// RegionId
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`
}

type ClusterExternalServiceInfo struct {
	// Dependency. `0`: Other clusters depend on the current cluster. `1`: The current cluster depends on other clusters.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DependType *int64 `json:"DependType,omitnil" name:"DependType"`

	// Shared component
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Service *string `json:"Service,omitnil" name:"Service"`

	// Sharing cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Sharing cluster status
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterStatus *int64 `json:"ClusterStatus,omitnil" name:"ClusterStatus"`
}

type ClusterInstancesInfo struct {
	// ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// Cluster ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Title
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ftitle *string `json:"Ftitle,omitnil" name:"Ftitle"`

	// Cluster name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// Region ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// User APPID
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Cluster `VPCID`
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcId *int64 `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubnetId *int64 `json:"SubnetId,omitnil" name:"SubnetId"`

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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// Execution duration
	// Note: this field may return null, indicating that no valid values can be obtained.
	RunTime *string `json:"RunTime,omitnil" name:"RunTime"`

	// Cluster product configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Config *EmrProductConfigOutter `json:"Config,omitnil" name:"Config"`

	// Public IP of master node
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterIp *string `json:"MasterIp,omitnil" name:"MasterIp"`

	// EMR version
	// Note: this field may return null, indicating that no valid values can be obtained.
	EmrVersion *string `json:"EmrVersion,omitnil" name:"EmrVersion"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// Transaction version
	// Note: this field may return null, indicating that no valid values can be obtained.
	TradeVersion *int64 `json:"TradeVersion,omitnil" name:"TradeVersion"`

	// Resource order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceOrderId *int64 `json:"ResourceOrderId,omitnil" name:"ResourceOrderId"`

	// Whether this is a paid cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsTradeCluster *int64 `json:"IsTradeCluster,omitnil" name:"IsTradeCluster"`

	// Alarm information for cluster error
	// Note: this field may return null, indicating that no valid values can be obtained.
	AlarmInfo *string `json:"AlarmInfo,omitnil" name:"AlarmInfo"`

	// Whether the new architecture is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWoodpeckerCluster *int64 `json:"IsWoodpeckerCluster,omitnil" name:"IsWoodpeckerCluster"`

	// Metadatabase information
	// Note: this field may return null, indicating that no valid values can be obtained.
	MetaDb *string `json:"MetaDb,omitnil" name:"MetaDb"`

	// Tag information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Hive metadata
	// Note: this field may return null, indicating that no valid values can be obtained.
	HiveMetaDb *string `json:"HiveMetaDb,omitnil" name:"HiveMetaDb"`

	// Cluster type: EMR, CLICKHOUSE, DRUID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceClass *string `json:"ServiceClass,omitnil" name:"ServiceClass"`

	// Alias serialization of all nodes in cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	AliasInfo *string `json:"AliasInfo,omitnil" name:"AliasInfo"`

	// Cluster version ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// Availability zone
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Scenario name
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// Scenario-based cluster type
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneServiceClass *string `json:"SceneServiceClass,omitnil" name:"SceneServiceClass"`

	// Scenario-based EMR version
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneEmrVersion *string `json:"SceneEmrVersion,omitnil" name:"SceneEmrVersion"`

	// Scenario-based cluster type
	// Note: This field may return `null`, indicating that no valid value was found.
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// Subnet name
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// Cluster dependency
	// Note: This field may return `null`, indicating that no valid value was found.
	ClusterExternalServiceInfo []*ClusterExternalServiceInfo `json:"ClusterExternalServiceInfo,omitnil" name:"ClusterExternalServiceInfo"`

	// The VPC ID string type of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// The subnet ID string type of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TopologyInfoList []*TopologyInfo `json:"TopologyInfoList,omitnil" name:"TopologyInfoList"`

	// Multi-AZ cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil" name:"IsMultiZoneCluster"`

	// Whether the feature of automatic abnormal node replacement is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCvmReplace *bool `json:"IsCvmReplace,omitnil" name:"IsCvmReplace"`
}

type ComponentBasicRestartInfo struct {
	// The process name (required), such as NameNode.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// The target IP list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IpList []*string `json:"IpList,omitnil" name:"IpList"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// The EMR version, such as `EMR-V2.3.0` that indicates the version 2.3.0 of EMR. You can query the EMR version [here](https://intl.cloud.tencent.com/document/product/589/66338?from_cn_redirect=1).
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// Whether to enable high availability for nodes. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil" name:"EnableSupportHAFlag"`

	// The instance name.
	// <li>Length limit: 6–36 characters.</li>
	// <li>Can contain only Chinese characters, letters, digits, hyphens (-), and underscores (_).</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The instance billing mode. Valid values:
	// <li>`POSTPAID_BY_HOUR`: The postpaid mode by hour.</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// The instance login setting. This parameter allows you to set a login password or key for your purchased node.
	// <li>If a key is set, the password will be used for login to the native component WebUI only.</li>
	// <li>If no key is set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// The configuration of cluster application scenario and supported components.
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil" name:"SceneSoftwareConfig"`

	// The details of the monthly subscription, including the instance period and auto-renewal. It is required if `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// The ID of the security group to which the instance belongs, in the format of `sg-xxxxxxxx`. You can call the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API and obtain this parameter from the `SecurityGroupId` field in the response.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// The [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings.
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// A unique random token, which is valid for 5 minutes and needs to be specified by the caller to prevent the client from repeatedly creating resources. An example value is `a9a90aa6-751a-41b6-aad6-fae360632808`.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Whether to enable public IP access for master nodes. Valid values:
	// <li>`NEED_MASTER_WAN`: Enable public IP for master nodes.</li>
	// <li>`NOT_NEED_MASTER_WAN`: Disable.</li>The public IP is enabled for master nodes by default.
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// Whether to enable remote login over the public network. It is invalid if `SecurityGroupId` is passed in. It is disabled by default. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil" name:"EnableRemoteLoginFlag"`

	// Whether to enable Kerberos authentication. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false` (default): Disable</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil" name:"EnableKerberosFlag"`

	// [Custom software configuration](https://intl.cloud.tencent.com/document/product/589/35655?from_cn_redirect=1?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil" name:"CustomConf"`

	// The tag description list. This parameter is used to bind a tag to a resource instance.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The list of spread placement group IDs. Only one can be specified.
	// You can call the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/product/213/17810?from_cn_redirect=1) API and obtain this parameter from the `DisasterRecoverGroupId` field in the response.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// Whether to enable the cluster-level CBS encryption. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false` (default): Disable</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil" name:"EnableCbsEncryptFlag"`

	// The metadatabase information. If `MetaType` is `EMR_NEW_META`, `MetaDataJdbcUrl`, `MetaDataUser`, `MetaDataPass`, and `UnifyMetaInstanceId` are not required.
	// If `MetaType` is `EMR_EXIT_META`, `UnifyMetaInstanceId` is required.
	// If `MetaType` is `USER_CUSTOM_META`, `MetaDataJdbcUrl`, `MetaDataUser`, and `MetaDataPass` are required.
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// The shared component information.
	DependService []*DependService `json:"DependService,omitnil" name:"DependService"`

	// The node resource specs. A spec is specified for each AZ, with the first spec for the primary AZ, the second for the backup AZ, and the third for the arbitrator AZ. If the multi-AZ mode is not enabled, only one spec is required.
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil" name:"ZoneResourceConfiguration"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// The EMR version, such as `EMR-V2.3.0` that indicates the version 2.3.0 of EMR. You can query the EMR version [here](https://intl.cloud.tencent.com/document/product/589/66338?from_cn_redirect=1).
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// Whether to enable high availability for nodes. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil" name:"EnableSupportHAFlag"`

	// The instance name.
	// <li>Length limit: 6–36 characters.</li>
	// <li>Can contain only Chinese characters, letters, digits, hyphens (-), and underscores (_).</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The instance billing mode. Valid values:
	// <li>`POSTPAID_BY_HOUR`: The postpaid mode by hour.</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// The instance login setting. This parameter allows you to set a login password or key for your purchased node.
	// <li>If a key is set, the password will be used for login to the native component WebUI only.</li>
	// <li>If no key is set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// The configuration of cluster application scenario and supported components.
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil" name:"SceneSoftwareConfig"`

	// The details of the monthly subscription, including the instance period and auto-renewal. It is required if `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// The ID of the security group to which the instance belongs, in the format of `sg-xxxxxxxx`. You can call the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API and obtain this parameter from the `SecurityGroupId` field in the response.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// The [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings.
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// A unique random token, which is valid for 5 minutes and needs to be specified by the caller to prevent the client from repeatedly creating resources. An example value is `a9a90aa6-751a-41b6-aad6-fae360632808`.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Whether to enable public IP access for master nodes. Valid values:
	// <li>`NEED_MASTER_WAN`: Enable public IP for master nodes.</li>
	// <li>`NOT_NEED_MASTER_WAN`: Disable.</li>The public IP is enabled for master nodes by default.
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// Whether to enable remote login over the public network. It is invalid if `SecurityGroupId` is passed in. It is disabled by default. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false`: Disable</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil" name:"EnableRemoteLoginFlag"`

	// Whether to enable Kerberos authentication. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false` (default): Disable</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil" name:"EnableKerberosFlag"`

	// [Custom software configuration](https://intl.cloud.tencent.com/document/product/589/35655?from_cn_redirect=1?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil" name:"CustomConf"`

	// The tag description list. This parameter is used to bind a tag to a resource instance.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The list of spread placement group IDs. Only one can be specified.
	// You can call the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/product/213/17810?from_cn_redirect=1) API and obtain this parameter from the `DisasterRecoverGroupId` field in the response.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// Whether to enable the cluster-level CBS encryption. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false` (default): Disable</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil" name:"EnableCbsEncryptFlag"`

	// The metadatabase information. If `MetaType` is `EMR_NEW_META`, `MetaDataJdbcUrl`, `MetaDataUser`, `MetaDataPass`, and `UnifyMetaInstanceId` are not required.
	// If `MetaType` is `EMR_EXIT_META`, `UnifyMetaInstanceId` is required.
	// If `MetaType` is `USER_CUSTOM_META`, `MetaDataJdbcUrl`, `MetaDataUser`, and `MetaDataPass` are required.
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// The shared component information.
	DependService []*DependService `json:"DependService,omitnil" name:"DependService"`

	// The node resource specs. A spec is specified for each AZ, with the first spec for the primary AZ, the second for the backup AZ, and the third for the arbitrator AZ. If the multi-AZ mode is not enabled, only one spec is required.
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil" name:"ZoneResourceConfiguration"`
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
	delete(f, "ProductVersion")
	delete(f, "EnableSupportHAFlag")
	delete(f, "InstanceName")
	delete(f, "InstanceChargeType")
	delete(f, "LoginSettings")
	delete(f, "SceneSoftwareConfig")
	delete(f, "InstanceChargePrepaid")
	delete(f, "SecurityGroupIds")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "EnableRemoteLoginFlag")
	delete(f, "EnableKerberosFlag")
	delete(f, "CustomConf")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "EnableCbsEncryptFlag")
	delete(f, "MetaDBInfo")
	delete(f, "DependService")
	delete(f, "ZoneResourceConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// The instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateInstanceRequestParams struct {
	// The product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>16: EMR v2.3.0</li>
	// <li>20: EMR v2.5.0</li>
	// <li>25: EMR v3.1.0</li>
	// <li>27: Kafka v1.0.0</li>
	// <li>30: EMR v2.6.0</li>
	// <li>33: EMR v3.2.1</li>
	// <li>34: EMR v3.3.0</li>
	// <li>36: StarRocks v1.0.0</li>
	// <li>37: EMR v3.4.0</li>
	// <li>38: EMR v2.7.0</li>
	// <li>39: StarRocks v1.1.0</li>
	// <li>41: Druid v1.1.0</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// List of deployed components. The list of component options varies by EMR product ID (i.e., `ProductId`; for specific meanings, please see the `ProductId` input parameter). For more information, please see [Component Version](https://intl.cloud.tencent.com/document/product/589/20279?from_cn_redirect=1).
	// Enter an instance value: `hive` or `flink`.
	Software []*string `json:"Software,omitnil" name:"Software"`

	// Whether to enable high node availability. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// Instance name.
	// <li>Length limit: 6-36 characters.</li>
	// <li>Only letters, numbers, dashes (-), and underscores (_) are supported.</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Instance login settings. This parameter allows you to set the login password or key for your purchased node.
	// <li>If the key is set, the password will be only used for login to the native component WebUI.</li>
	// <li>If the key is not set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// Node resource specification.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// Parameter required for enabling COS access.
	COSSettings *COSSettings `json:"COSSettings,omitnil" name:"COSSettings"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Security group to which an instance belongs in the format of `sg-xxxxxxxx`. This parameter can be obtained from the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) API.
	SgId *string `json:"SgId,omitnil" name:"SgId"`

	// [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// Whether auto-renewal is enabled. Valid values:
	// <li>0: auto-renewal not enabled.</li>
	// <li>1: auto-renewal enabled.</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Whether to enable public IP access for master node. Valid values:
	// <li>NEED_MASTER_WAN: enables public IP for master node.</li>
	// <li>NOT_NEED_MASTER_WAN: does not enable.</li>Public IP is enabled for master node by default.
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// Whether to enable remote public network login, i.e., port 22. When `SgId` is not empty, this parameter does not take effect.
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil" name:"RemoteLoginAtCreate"`

	// Whether to enable secure cluster. 0: no; other values: yes.
	CheckSecurity *int64 `json:"CheckSecurity,omitnil" name:"CheckSecurity"`

	// Accesses to external file system.
	ExtendFsField *string `json:"ExtendFsField,omitnil" name:"ExtendFsField"`

	// Tag description list. This parameter is used to bind a tag to a resource instance.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// List of spread placement group IDs. Only one can be specified currently.
	// This parameter can be obtained in the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1) API.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// CBS disk encryption at the cluster level. 0: not encrypted, 1: encrypted
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// Custom application role.
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`


	VersionID *int64 `json:"VersionID,omitnil" name:"VersionID"`

	// `true` indicates that the multi-AZ deployment mode is enabled. This parameter is available only in cluster creation and cannot be changed after setting.
	MultiZone *bool `json:"MultiZone,omitnil" name:"MultiZone"`

	// Node resource specs. The actual number of AZs is set, with the first AZ as the primary AZ, the second as the backup AZ, and the third as the arbitrator AZ. If the multi-AZ mode is not enabled, set the value to `1`.
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>16: EMR v2.3.0</li>
	// <li>20: EMR v2.5.0</li>
	// <li>25: EMR v3.1.0</li>
	// <li>27: Kafka v1.0.0</li>
	// <li>30: EMR v2.6.0</li>
	// <li>33: EMR v3.2.1</li>
	// <li>34: EMR v3.3.0</li>
	// <li>36: StarRocks v1.0.0</li>
	// <li>37: EMR v3.4.0</li>
	// <li>38: EMR v2.7.0</li>
	// <li>39: StarRocks v1.1.0</li>
	// <li>41: Druid v1.1.0</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// List of deployed components. The list of component options varies by EMR product ID (i.e., `ProductId`; for specific meanings, please see the `ProductId` input parameter). For more information, please see [Component Version](https://intl.cloud.tencent.com/document/product/589/20279?from_cn_redirect=1).
	// Enter an instance value: `hive` or `flink`.
	Software []*string `json:"Software,omitnil" name:"Software"`

	// Whether to enable high node availability. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// Instance name.
	// <li>Length limit: 6-36 characters.</li>
	// <li>Only letters, numbers, dashes (-), and underscores (_) are supported.</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Instance login settings. This parameter allows you to set the login password or key for your purchased node.
	// <li>If the key is set, the password will be only used for login to the native component WebUI.</li>
	// <li>If the key is not set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// Node resource specification.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// Parameter required for enabling COS access.
	COSSettings *COSSettings `json:"COSSettings,omitnil" name:"COSSettings"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Security group to which an instance belongs in the format of `sg-xxxxxxxx`. This parameter can be obtained from the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) API.
	SgId *string `json:"SgId,omitnil" name:"SgId"`

	// [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// Whether auto-renewal is enabled. Valid values:
	// <li>0: auto-renewal not enabled.</li>
	// <li>1: auto-renewal enabled.</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Whether to enable public IP access for master node. Valid values:
	// <li>NEED_MASTER_WAN: enables public IP for master node.</li>
	// <li>NOT_NEED_MASTER_WAN: does not enable.</li>Public IP is enabled for master node by default.
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// Whether to enable remote public network login, i.e., port 22. When `SgId` is not empty, this parameter does not take effect.
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil" name:"RemoteLoginAtCreate"`

	// Whether to enable secure cluster. 0: no; other values: yes.
	CheckSecurity *int64 `json:"CheckSecurity,omitnil" name:"CheckSecurity"`

	// Accesses to external file system.
	ExtendFsField *string `json:"ExtendFsField,omitnil" name:"ExtendFsField"`

	// Tag description list. This parameter is used to bind a tag to a resource instance.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// List of spread placement group IDs. Only one can be specified currently.
	// This parameter can be obtained in the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1) API.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// CBS disk encryption at the cluster level. 0: not encrypted, 1: encrypted
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// Custom application role.
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	VersionID *int64 `json:"VersionID,omitnil" name:"VersionID"`

	// `true` indicates that the multi-AZ deployment mode is enabled. This parameter is available only in cluster creation and cannot be changed after setting.
	MultiZone *bool `json:"MultiZone,omitnil" name:"MultiZone"`

	// Node resource specs. The actual number of AZs is set, with the first AZ as the primary AZ, the second as the backup AZ, and the third as the arbitrator AZ. If the multi-AZ mode is not enabled, set the value to `1`.
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Software")
	delete(f, "SupportHA")
	delete(f, "InstanceName")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "LoginSettings")
	delete(f, "VPCSettings")
	delete(f, "ResourceSpec")
	delete(f, "COSSettings")
	delete(f, "Placement")
	delete(f, "SgId")
	delete(f, "PreExecutedFileSettings")
	delete(f, "AutoRenew")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "RemoteLoginAtCreate")
	delete(f, "CheckSecurity")
	delete(f, "ExtendFsField")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "CbsEncrypt")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ApplicationRole")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZone")
	delete(f, "MultiZoneSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// Instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetaDBInfo struct {
	// The JDBC URL of the custom metadatabase instance. Example: jdbc:mysql://10.10.10.10:3306/dbname
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil" name:"MetaDataJdbcUrl"`

	// The custom metadatabase instance username.
	MetaDataUser *string `json:"MetaDataUser,omitnil" name:"MetaDataUser"`

	// The custom metadatabase instance password.
	MetaDataPass *string `json:"MetaDataPass,omitnil" name:"MetaDataPass"`

	// The Hive-shared metadatabase type. Valid values:
	// <li>`EMR_DEFAULT_META`: The cluster creates one by default.</li>
	// <li>`EMR_EXIST_META`: The cluster uses the specified EMR metadatabase instance.</li>
	// <li>`USER_CUSTOM_META`: The cluster uses a custom metadatabase instance.</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// The EMR-MetaDB instance.
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`
}

type CustomMetaInfo struct {
	// JDBC connection to custom MetaDB instance beginning with `jdbc:mysql://`
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil" name:"MetaDataJdbcUrl"`

	// Custom MetaDB instance username
	MetaDataUser *string `json:"MetaDataUser,omitnil" name:"MetaDataUser"`

	// Custom MetaDB instance password
	MetaDataPass *string `json:"MetaDataPass,omitnil" name:"MetaDataPass"`
}

type CustomServiceDefine struct {
	// Custom parameter key
	Name *string `json:"Name,omitnil" name:"Name"`

	// Custom parameter value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DependService struct {
	// The shared component name.
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// The cluster to which the shared component belongs.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type DescribeClusterNodesRequestParams struct {
	// Cluster instance ID in the format of emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Node flag. Valid values:
	// <li>all: gets the information of nodes in all types except TencentDB information.</li>
	// <li>master: gets master node information.</li>
	// <li>core: gets core node information.</li>
	// <li>task: gets task node information.</li>
	// <li>common: gets common node information.</li>
	// <li>router: gets router node information.</li>
	// <li>db: gets TencentDB information in normal status.</li>
	// Note: only the above values are supported for the time being. Entering other values will cause errors.
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// Page number. Default value: 0, indicating the first page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: 100. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Resource type. Valid values: all, host, pod. Default value: all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// Searchable field
	SearchFields []*SearchItem `json:"SearchFields,omitnil" name:"SearchFields"`

	// None
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// None
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID in the format of emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Node flag. Valid values:
	// <li>all: gets the information of nodes in all types except TencentDB information.</li>
	// <li>master: gets master node information.</li>
	// <li>core: gets core node information.</li>
	// <li>task: gets task node information.</li>
	// <li>common: gets common node information.</li>
	// <li>router: gets router node information.</li>
	// <li>db: gets TencentDB information in normal status.</li>
	// Note: only the above values are supported for the time being. Entering other values will cause errors.
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// Page number. Default value: 0, indicating the first page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: 100. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Resource type. Valid values: all, host, pod. Default value: all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// Searchable field
	SearchFields []*SearchItem `json:"SearchFields,omitnil" name:"SearchFields"`

	// None
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// None
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeFlag")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "HardwareResourceType")
	delete(f, "SearchFields")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodesResponseParams struct {
	// Total number of queried nodes
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// List of node details
	// Note: this field may return null, indicating that no valid values can be obtained.
	NodeList []*NodeHardwareInfo `json:"NodeList,omitnil" name:"NodeList"`

	// List of tag keys owned by user
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Resource type list
	// Note: this field may return null, indicating that no valid values can be obtained.
	HardwareResourceTypeList []*string `json:"HardwareResourceTypeList,omitnil" name:"HardwareResourceTypeList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodesResponseParams `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time in the format of timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Queue name used for filtering
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// Username used for filtering
	Users []*string `json:"Users,omitnil" name:"Users"`

	// Application type used for filtering
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// Group field. Valid values: `queue`, `user`, and `applicationType`.
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// Sorting field. Valid values: `sumMemorySeconds`, `sumVCoreSeconds`, `sumHDFSBytesWritten`, and `sumHDFSBytesRead`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Order type. Valid values: `0` (descending) and `1`(ascending).
	IsAsc *int64 `json:"IsAsc,omitnil" name:"IsAsc"`

	// Page number
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Page limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeEmrApplicationStaticsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time in the format of timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Queue name used for filtering
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// Username used for filtering
	Users []*string `json:"Users,omitnil" name:"Users"`

	// Application type used for filtering
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// Group field. Valid values: `queue`, `user`, and `applicationType`.
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// Sorting field. Valid values: `sumMemorySeconds`, `sumVCoreSeconds`, `sumHDFSBytesWritten`, and `sumHDFSBytesRead`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Order type. Valid values: `0` (descending) and `1`(ascending).
	IsAsc *int64 `json:"IsAsc,omitnil" name:"IsAsc"`

	// Page number
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Page limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeEmrApplicationStaticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Queues")
	delete(f, "Users")
	delete(f, "ApplicationTypes")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "IsAsc")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmrApplicationStaticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsResponseParams struct {
	// Application statistics
	Statics []*ApplicationStatics `json:"Statics,omitnil" name:"Statics"`

	// Total count
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Available queue name
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// Available usernames
	Users []*string `json:"Users,omitnil" name:"Users"`

	// Available application type
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEmrApplicationStaticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmrApplicationStaticsResponseParams `json:"Response"`
}

func (r *DescribeEmrApplicationStaticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListRequestParams struct {
	// Cluster filtering policy. Valid values: <li>clusterList: Queries the list of clusters excluding terminated ones.</li><li>monitorManage: Queries the list of clusters excluding those terminated, under creation and not successfully created.</li><li>cloudHardwareManage/componentManage: Two reserved values, which have the same implications as those of `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// Page number. Default value: `0`, indicating the first page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: `10`; maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Sorting field. Valid values: <li>clusterId: Sorting by instance ID. </li><li>addTime: Sorting by instance creation time.</li><li>status: Sorting by instance status code.</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sort ascending or descending based on `OrderField`. Valid values:<li>0: Descending.</li><li>1: Ascending.</li>Default value: `0`.
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`

	// Custom query
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeInstancesListRequest struct {
	*tchttp.BaseRequest
	
	// Cluster filtering policy. Valid values: <li>clusterList: Queries the list of clusters excluding terminated ones.</li><li>monitorManage: Queries the list of clusters excluding those terminated, under creation and not successfully created.</li><li>cloudHardwareManage/componentManage: Two reserved values, which have the same implications as those of `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// Page number. Default value: `0`, indicating the first page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: `10`; maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Sorting field. Valid values: <li>clusterId: Sorting by instance ID. </li><li>addTime: Sorting by instance creation time.</li><li>status: Sorting by instance status code.</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sort ascending or descending based on `OrderField`. Valid values:<li>0: Descending.</li><li>1: Ascending.</li>Default value: `0`.
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`

	// Custom query
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeInstancesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Asc")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListResponseParams struct {
	// Number of eligible instances.
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// Cluster instance list.
	InstancesList []*EmrListInstance `json:"InstancesList,omitnil" name:"InstancesList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesListResponseParams `json:"Response"`
}

func (r *DescribeInstancesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Cluster filtering policy. Valid values:
	// <li>clusterList: queries the list of clusters except terminated ones.</li>
	// <li>monitorManage: queries the list of clusters except those that have been terminated, are being created, or failed to be created.</li>
	// <li>cloudHardwareManage/componentManage: reserved fields with the same meaning as `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// Queries by one or more instance IDs in the format of `emr-xxxxxxxx`. For the format of this parameter, please see the `id.N` section in [API Overview](https://intl.cloud.tencent.com/document/api/213/15688). If no instance ID is entered, the list of all instances under this `APPID` will be returned.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Page number. Default value: 0, indicating the first page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: 10. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` field in the return value of the `DescribeProject` API. If this value is -1, the list of all instances will be returned.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Sorting field. Valid values:
	// <li>clusterId: sorts by cluster ID.</li>
	// <li>addTime: sorts by instance creation time.</li>
	// <li>status: sorts by instance status code.</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sorts according to `OrderField` in ascending or descending order. Valid values:
	// <li>0: descending order.</li>
	// <li>1: ascending order.</li>Default value: 0.
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster filtering policy. Valid values:
	// <li>clusterList: queries the list of clusters except terminated ones.</li>
	// <li>monitorManage: queries the list of clusters except those that have been terminated, are being created, or failed to be created.</li>
	// <li>cloudHardwareManage/componentManage: reserved fields with the same meaning as `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// Queries by one or more instance IDs in the format of `emr-xxxxxxxx`. For the format of this parameter, please see the `id.N` section in [API Overview](https://intl.cloud.tencent.com/document/api/213/15688). If no instance ID is entered, the list of all instances under this `APPID` will be returned.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Page number. Default value: 0, indicating the first page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results per page. Default value: 10. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` field in the return value of the `DescribeProject` API. If this value is -1, the list of all instances will be returned.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Sorting field. Valid values:
	// <li>clusterId: sorts by cluster ID.</li>
	// <li>addTime: sorts by instance creation time.</li>
	// <li>status: sorts by instance status code.</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sorts according to `OrderField` in ascending or descending order. Valid values:
	// <li>0: descending order.</li>
	// <li>1: ascending order.</li>Default value: 0.
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of eligible instances.
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// List of EMR instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitnil" name:"ClusterList"`

	// List of tag keys associated to an instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleRequestParams struct {
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeResourceScheduleRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeResourceScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleResponseParams struct {
	// Whether to enable the resource scheduling feature
	OpenSwitch *bool `json:"OpenSwitch,omitnil" name:"OpenSwitch"`

	// The resource scheduler in service
	Scheduler *string `json:"Scheduler,omitnil" name:"Scheduler"`

	// Fair Scheduler information
	FSInfo *string `json:"FSInfo,omitnil" name:"FSInfo"`

	// Capacity Scheduler information
	CSInfo *string `json:"CSInfo,omitnil" name:"CSInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResourceScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceScheduleResponseParams `json:"Response"`
}

func (r *DescribeResourceScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Page number
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// Page size
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// User list query filter
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil" name:"UserManagerFilter"`

	// Whether the Keytab file information is required. This field is only valid for clusters with Kerberos enabled and defaults to `false`.
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil" name:"NeedKeytabInfo"`
}

type DescribeUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Page number
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// Page size
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// User list query filter
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil" name:"UserManagerFilter"`

	// Whether the Keytab file information is required. This field is only valid for clusters with Kerberos enabled and defaults to `false`.
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil" name:"NeedKeytabInfo"`
}

func (r *DescribeUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "UserManagerFilter")
	delete(f, "NeedKeytabInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerResponseParams struct {
	// Total number
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// User information list
	// Note: This field may return null, indicating that no valid value can be obtained.
	UserManagerUserList []*UserManagerUserBriefInfo `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersForUserManagerResponseParams `json:"Response"`
}

func (r *DescribeUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpecInfo struct {
	// The number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// The system disk type. Valid values:
	// <li>`CLOUD_SSD`: Cloud SSD</li>
	// <li>`CLOUD_PREMIUM`: Premium cloud disk</li>
	// <li>`CLOUD_BASIC`: Cloud HDD</li>
	// <li>`LOCAL_BASIC`: Local disk</li>
	// <li>`LOCAL_SSD`: Local SSD</li>
	// 
	// The data disk type. Valid values:
	// <li>`CLOUD_SSD`: Cloud SSD</li>
	// <li>`CLOUD_PREMIUM`: Premium cloud disk</li>
	// <li>`CLOUD_BASIC`: Cloud HDD</li>
	// <li>`LOCAL_BASIC`: Local disk</li>
	// <li>`LOCAL_SSD`: Local SSD</li>
	// <li>`CLOUD_HSSD`: Enhanced SSD</li>
	// <li>`CLOUD_THROUGHPUT`: Throughput HDD</li>
	// <li>CLOUD_TSSD: ulTra SSD</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// The disk capacity in GB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type DynamicPodSpec struct {
	// Minimum number of CPUs
	RequestCpu *float64 `json:"RequestCpu,omitnil" name:"RequestCpu"`

	// Maximum number of CPUs
	LimitCpu *float64 `json:"LimitCpu,omitnil" name:"LimitCpu"`

	// Minimum memory in MB
	RequestMemory *float64 `json:"RequestMemory,omitnil" name:"RequestMemory"`

	// Maximum memory in MB
	LimitMemory *float64 `json:"LimitMemory,omitnil" name:"LimitMemory"`
}

type EmrListInstance struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Status description
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Cluster region
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// User APPID
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// Creation time
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// Running time
	RunTime *string `json:"RunTime,omitnil" name:"RunTime"`

	// Cluster IP
	MasterIp *string `json:"MasterIp,omitnil" name:"MasterIp"`

	// Cluster version
	EmrVersion *string `json:"EmrVersion,omitnil" name:"EmrVersion"`

	// Cluster billing mode
	ChargeType *uint64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// EMR ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Product ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// Project ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetId *uint64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VpcId *uint64 `json:"VpcId,omitnil" name:"VpcId"`

	// Region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Status code
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Instance tag
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Alarm information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlarmInfo *string `json:"AlarmInfo,omitnil" name:"AlarmInfo"`

	// Whether it is a Woodpecker cluster
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsWoodpeckerCluster *uint64 `json:"IsWoodpeckerCluster,omitnil" name:"IsWoodpeckerCluster"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// Subnet name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// VPC ID string
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Subnet ID string
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Cluster type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClusterClass *string `json:"ClusterClass,omitnil" name:"ClusterClass"`

	// Whether it is a multi-AZ cluster
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil" name:"IsMultiZoneCluster"`

	// Whether it is a manually deployed cluster
	// Note: This field may return null, indicating that no valid value can be obtained. 
	IsHandsCluster *bool `json:"IsHandsCluster,omitnil" name:"IsHandsCluster"`

	// Client component information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OutSideSoftInfo []*SoftDependInfo `json:"OutSideSoftInfo,omitnil" name:"OutSideSoftInfo"`

	// Whether the current cluster supports external clients.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSupportOutsideCluster *bool `json:"IsSupportOutsideCluster,omitnil" name:"IsSupportOutsideCluster"`
}

type EmrPrice struct {
	// The published price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// The discounted price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountCost *string `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// The unit of the billable item.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// The queried spec.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceSpec *PriceResource `json:"PriceSpec,omitnil" name:"PriceSpec"`

	// Whether spot instances are supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportSpotPaid *bool `json:"SupportSpotPaid,omitnil" name:"SupportSpotPaid"`
}

type EmrProductConfigOutter struct {
	// Software information
	// Note: this field may return null, indicating that no valid values can be obtained.
	SoftInfo []*string `json:"SoftInfo,omitnil" name:"SoftInfo"`

	// Number of master nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterNodeSize *int64 `json:"MasterNodeSize,omitnil" name:"MasterNodeSize"`

	// Number of core nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoreNodeSize *int64 `json:"CoreNodeSize,omitnil" name:"CoreNodeSize"`

	// Number of task nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskNodeSize *int64 `json:"TaskNodeSize,omitnil" name:"TaskNodeSize"`

	// Number of common nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComNodeSize *int64 `json:"ComNodeSize,omitnil" name:"ComNodeSize"`

	// Master node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterResource *OutterResource `json:"MasterResource,omitnil" name:"MasterResource"`

	// Core node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	CoreResource *OutterResource `json:"CoreResource,omitnil" name:"CoreResource"`

	// Task node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskResource *OutterResource `json:"TaskResource,omitnil" name:"TaskResource"`

	// Common node resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	ComResource *OutterResource `json:"ComResource,omitnil" name:"ComResource"`

	// Whether COS is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	OnCos *bool `json:"OnCos,omitnil" name:"OnCos"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// Number of router nodes
	// Note: this field may return null, indicating that no valid values can be obtained.
	RouterNodeSize *int64 `json:"RouterNodeSize,omitnil" name:"RouterNodeSize"`

	// Whether HA is supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	SupportHA *bool `json:"SupportHA,omitnil" name:"SupportHA"`

	// Whether secure mode is supported
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecurityOn *bool `json:"SecurityOn,omitnil" name:"SecurityOn"`

	// Security group name
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecurityGroup *string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// Whether to enable CBS encryption
	// Note: this field may return null, indicating that no valid values can be obtained.
	CbsEncrypt *int64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// Custom application role
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// Security groups
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	SecurityGroups []*string `json:"SecurityGroups,omitnil" name:"SecurityGroups"`

	// SSH key ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PublicKeyId *string `json:"PublicKeyId,omitnil" name:"PublicKeyId"`
}

type ExternalService struct {
	// Shared component type, which can be EMR or CUSTOM
	ShareType *string `json:"ShareType,omitnil" name:"ShareType"`

	// Custom parameters
	CustomServiceDefineList []*CustomServiceDefine `json:"CustomServiceDefineList,omitnil" name:"CustomServiceDefineList"`

	// Shared component name
	Service *string `json:"Service,omitnil" name:"Service"`

	// Shared component cluster
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type Filters struct {
	// Field name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type HostVolumeContext struct {
	// The directory for mounting the host in the pod, which is the mount point of the host in the resource. A specified mount point corresponds to the host path and is used as the data storage directory in the pod.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumePath *string `json:"VolumePath,omitnil" name:"VolumePath"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Whether to enable high availability of node. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// List of deployed components. Different required components need to be selected for different EMR product IDs (i.e., `ProductId`; for specific meanings, please see the `ProductId` field in the input parameter):
	// <li>When `ProductId` is 1, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 2, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 4, the required components include hadoop-2.8.4, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 7, the required components include hadoop-3.1.2, knox-1.2.0, and zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitnil" name:"Software"`

	// Node specification queried for price.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1.</li>
	// <li>2: EMR v2.0.1.</li>
	// <li>4: EMR v2.1.0.</li>
	// <li>7: EMR v3.0.0.</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`


	VersionID *uint64 `json:"VersionID,omitnil" name:"VersionID"`

	// AZ specs
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Whether to enable high availability of node. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// List of deployed components. Different required components need to be selected for different EMR product IDs (i.e., `ProductId`; for specific meanings, please see the `ProductId` field in the input parameter):
	// <li>When `ProductId` is 1, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 2, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 4, the required components include hadoop-2.8.4, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 7, the required components include hadoop-3.1.2, knox-1.2.0, and zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitnil" name:"Software"`

	// Node specification queried for price.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// Hive-shared metadatabase type. Valid values:
	// <li>EMR_DEFAULT_META: the cluster creates one by default.</li>
	// <li>EMR_EXIST_META: the cluster uses the specified EMR-MetaDB instance.</li>
	// <li>USER_CUSTOM_META: the cluster uses a custom MetaDB instance.</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB instance
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// Custom MetaDB instance information
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1.</li>
	// <li>2: EMR v2.0.1.</li>
	// <li>4: EMR v2.1.0.</li>
	// <li>7: EMR v3.0.0.</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	VersionID *uint64 `json:"VersionID,omitnil" name:"VersionID"`

	// AZ specs
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "Currency")
	delete(f, "PayMode")
	delete(f, "SupportHA")
	delete(f, "Software")
	delete(f, "ResourceSpec")
	delete(f, "Placement")
	delete(f, "VPCSettings")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ProductId")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZoneSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Purchase duration of instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// The price list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceList []*ZoneDetailPriceResult `json:"PriceList,omitnil" name:"PriceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceRequestParams struct {
	// How long the instance will be renewed for, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// List of resource IDs of the node to be renewed. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// Location of the instance. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Instance billing mode.
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// Unit of time for instance renewal.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Whether to change from pay-as-you-go billing to monthly subscription billing. `0`: no; `1`: yes
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// How long the instance will be renewed for, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// List of resource IDs of the node to be renewed. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// Location of the instance. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Instance billing mode.
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// Unit of time for instance renewal.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Whether to change from pay-as-you-go billing to monthly subscription billing. `0`: no; `1`: yes
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "ResourceIds")
	delete(f, "Placement")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	delete(f, "ModifyPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceResponseParams struct {
	// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// Unit of time for instance renewal.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// How long the instance will be renewed for.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceRequestParams struct {
	// Time unit of scale-out. Valid value:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// ID of the AZ where an instance resides.
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// Number of master nodes to be added.
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of scale-out. Valid value:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// ID of the AZ where an instance resides.
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// Number of master nodes to be added.
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`
}

func (r *InquiryPriceScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "ZoneId")
	delete(f, "PayMode")
	delete(f, "InstanceId")
	delete(f, "CoreCount")
	delete(f, "TaskCount")
	delete(f, "Currency")
	delete(f, "RouterCount")
	delete(f, "MasterCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceResponseParams struct {
	// Original price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Discounted price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountCost *string `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// Time unit of scale-out. Valid value:
	// <li>s: Second.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Node spec queried for price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceSpec *PriceResource `json:"PriceSpec,omitnil" name:"PriceSpec"`

	// The inquiry results corresponding to the specs specified by the input parameter `MultipleResources`, with the result of the first spec returned by other output parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MultipleEmrPrice []*EmrPrice `json:"MultipleEmrPrice,omitnil" name:"MultipleEmrPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceScaleOutInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceRequestParams struct {
	// Time unit of scaling. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Duration of scaling, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Target node specification.
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil" name:"UpdateSpec"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// The resource ID list for batch configuration change.
	ResourceIdList []*string `json:"ResourceIdList,omitnil" name:"ResourceIdList"`
}

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of scaling. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Duration of scaling, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Target node specification.
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil" name:"UpdateSpec"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// Currency.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// The resource ID list for batch configuration change.
	ResourceIdList []*string `json:"ResourceIdList,omitnil" name:"ResourceIdList"`
}

func (r *InquiryPriceUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "UpdateSpec")
	delete(f, "PayMode")
	delete(f, "Placement")
	delete(f, "Currency")
	delete(f, "ResourceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceResponseParams struct {
	// Original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Discounted price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// Time unit of scaling. Valid values:
	// <li>s: seconds.</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Duration of scaling.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Pricing details
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceDetail []*PriceDetail `json:"PriceDetail,omitnil" name:"PriceDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpdateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// The period of monthly subscription, which defaults to 1 and is expressed in month.
	// Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Whether to enable auto-renewal. Valid values:
	// <li>`true`: Enable</li>
	// <li>`false` (default): Disable</li>
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type LoginSettings struct {
	// The login password of the instance, which contains 8 to 16 uppercase letters, lowercase letters, digits, and special characters (only !@%^*) and cannot start with a special character.
	Password *string `json:"Password,omitnil" name:"Password"`

	// The key ID. After an instance is associated with a key, you can access it with the private key in the key pair. You can call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) to obtain `PublicKeyId`.
	PublicKeyId *string `json:"PublicKeyId,omitnil" name:"PublicKeyId"`
}

// Predefined struct for user
type ModifyResourceScheduleConfigRequestParams struct {
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Business identifier. `fair`: Edit fair configuration items; `fairPlan`: Edit the execution plan; `capacity`: Edit capacity configuration items.
	Key *string `json:"Key,omitnil" name:"Key"`

	// Modified module information
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ModifyResourceScheduleConfigRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Business identifier. `fair`: Edit fair configuration items; `fairPlan`: Edit the execution plan; `capacity`: Edit capacity configuration items.
	Key *string `json:"Key,omitnil" name:"Key"`

	// Modified module information
	Value *string `json:"Value,omitnil" name:"Value"`
}

func (r *ModifyResourceScheduleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceScheduleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceScheduleConfigResponseParams struct {
	// `true`: Draft, indicating the resource pool is not refreshed.
	IsDraft *bool `json:"IsDraft,omitnil" name:"IsDraft"`

	// Verification error information. If it is not null, the verification fails and thus the configuration fails.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// The response data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourceScheduleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceScheduleConfigResponseParams `json:"Response"`
}

func (r *ModifyResourceScheduleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerRequestParams struct {
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The original scheduler: `fair`
	OldValue *string `json:"OldValue,omitnil" name:"OldValue"`

	// The new scheduler: `capacity`
	NewValue *string `json:"NewValue,omitnil" name:"NewValue"`
}

type ModifyResourceSchedulerRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The original scheduler: `fair`
	OldValue *string `json:"OldValue,omitnil" name:"OldValue"`

	// The new scheduler: `capacity`
	NewValue *string `json:"NewValue,omitnil" name:"NewValue"`
}

func (r *ModifyResourceSchedulerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldValue")
	delete(f, "NewValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceSchedulerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourceSchedulerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceSchedulerResponseParams `json:"Response"`
}

func (r *ModifyResourceSchedulerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiDisk struct {
	// Cloud disk type
	// <li>`CLOUD_SSD`: SSD</li>
	// <li>`CLOUD_PREMIUM`: Premium Cloud Storage</li>
	// <li>`CLOUD_HSSD`: Enhanced SSD</li>
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Cloud disk size
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// Number of cloud disks of this type
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type MultiDiskMC struct {
	// Number of cloud disks in this type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Cloud disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

type MultiZoneSetting struct {
	// "master", "standby", "third-party"
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneTag *string `json:"ZoneTag,omitnil" name:"ZoneTag"`

	// None
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// None
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// None
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`
}

type NewResourceSpec struct {
	// Describes master node resource
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitnil" name:"MasterResourceSpec"`

	// Describes core node resource
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitnil" name:"CoreResourceSpec"`

	// Describes task node resource
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitnil" name:"TaskResourceSpec"`

	// Number of master nodes
	MasterCount *int64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// Number of core nodes
	CoreCount *int64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Number of task nodes
	TaskCount *int64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Describes common node resource
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitnil" name:"CommonResourceSpec"`

	// Number of common nodes
	CommonCount *int64 `json:"CommonCount,omitnil" name:"CommonCount"`
}

type NodeDetailPriceResult struct {
	// The node type. Valid values: `master`, `core`, `task`, `common`, `router`, `mysql`
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// Price details by node part
	PartDetailPrice []*PartDetailPriceItem `json:"PartDetailPrice,omitnil" name:"PartDetailPrice"`
}

type NodeHardwareInfo struct {
	// User `APPID`
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Serial number
	// Note: this field may return null, indicating that no valid values can be obtained.
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// Machine instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderNo *string `json:"OrderNo,omitnil" name:"OrderNo"`

	// Public IP bound to master node
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// Node type. 0: common node; 1: master node;
	// 2: core node; 3: task node
	// Note: this field may return null, indicating that no valid values can be obtained.
	Flag *int64 `json:"Flag,omitnil" name:"Flag"`

	// Node specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// Number of node cores
	// Note: this field may return null, indicating that no valid values can be obtained.
	CpuNum *int64 `json:"CpuNum,omitnil" name:"CpuNum"`

	// Node memory size
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// Node memory description
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemDesc *string `json:"MemDesc,omitnil" name:"MemDesc"`

	// Node region
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// Node AZ
	// Note: this field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplyTime *string `json:"ApplyTime,omitnil" name:"ApplyTime"`

	// Release time
	// Note: this field may return null, indicating that no valid values can be obtained.
	FreeTime *string `json:"FreeTime,omitnil" name:"FreeTime"`

	// Disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *string `json:"DiskSize,omitnil" name:"DiskSize"`

	// Node description
	// Note: this field may return null, indicating that no valid values can be obtained.
	NameTag *string `json:"NameTag,omitnil" name:"NameTag"`

	// Services deployed on node
	// Note: this field may return null, indicating that no valid values can be obtained.
	Services *string `json:"Services,omitnil" name:"Services"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// System disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// Payment type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// Database IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbIp *string `json:"CdbIp,omitnil" name:"CdbIp"`

	// Database port
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbPort *int64 `json:"CdbPort,omitnil" name:"CdbPort"`

	// Disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwDiskSize *int64 `json:"HwDiskSize,omitnil" name:"HwDiskSize"`

	// Disk capacity description
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwDiskSizeDesc *string `json:"HwDiskSizeDesc,omitnil" name:"HwDiskSizeDesc"`

	// Memory capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwMemSize *int64 `json:"HwMemSize,omitnil" name:"HwMemSize"`

	// Memory capacity description
	// Note: this field may return null, indicating that no valid values can be obtained.
	HwMemSizeDesc *string `json:"HwMemSizeDesc,omitnil" name:"HwMemSizeDesc"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Node resource ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	EmrResourceId *string `json:"EmrResourceId,omitnil" name:"EmrResourceId"`

	// Renewal flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`

	// Device flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeviceClass *string `json:"DeviceClass,omitnil" name:"DeviceClass"`

	// Support for configuration adjustment
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mutable *int64 `json:"Mutable,omitnil" name:"Mutable"`

	// Multi-cloud disk
	// Note: this field may return null, indicating that no valid values can be obtained.
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitnil" name:"MCMultiDisk"`

	// Database information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CdbNodeInfo *CdbInfo `json:"CdbNodeInfo,omitnil" name:"CdbNodeInfo"`

	// Private IP
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Whether this node can be terminated. 1: yes, 0: no
	// Note: this field may return null, indicating that no valid values can be obtained.
	Destroyable *int64 `json:"Destroyable,omitnil" name:"Destroyable"`

	// Tags bound to node
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Wether the node is auto-scaling. 0 means common node. 1 means auto-scaling node.
	AutoFlag *int64 `json:"AutoFlag,omitnil" name:"AutoFlag"`

	// Resource type. Valid values: host, pod
	// Note: this field may return null, indicating that no valid values can be obtained.
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// Whether floating specification is used. `1`: yes; `0`: no
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	IsDynamicSpec *int64 `json:"IsDynamicSpec,omitnil" name:"IsDynamicSpec"`

	// Floating specification in JSON string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DynamicPodSpec *string `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// Whether to support billing mode change. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SupportModifyPayMode *int64 `json:"SupportModifyPayMode,omitnil" name:"SupportModifyPayMode"`

	// System disk type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RootStorageType *int64 `json:"RootStorageType,omitnil" name:"RootStorageType"`

	// AZ information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Subnet
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil" name:"SubnetInfo"`

	// Client
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Clients *string `json:"Clients,omitnil" name:"Clients"`

	// The current system time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentTime *string `json:"CurrentTime,omitnil" name:"CurrentTime"`

	// Whether it is used in a federation. Valid values: `0` (no), `1` (yes).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFederation *int64 `json:"IsFederation,omitnil" name:"IsFederation"`

	// Device name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Service
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceClient *string `json:"ServiceClient,omitnil" name:"ServiceClient"`

	// Enabling instance protection for this instance. Valid values: `true` (enable) and `false` (disable).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil" name:"DisableApiTermination"`

	// The billing version. Valid values: `0` (original billing) and `1` (new billing)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TradeVersion *int64 `json:"TradeVersion,omitnil" name:"TradeVersion"`
}

type NodeResourceSpec struct {
	// The spec type, such as `S2.MEDIUM8`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// The system disk, which can be up to 1 PCS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SystemDisk []*DiskSpecInfo `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// The list of tags to be bound.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The cloud data disk, which can be up to 15 PCS.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDisk []*DiskSpecInfo `json:"DataDisk,omitnil" name:"DataDisk"`

	// The local data disk.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalDataDisk []*DiskSpecInfo `json:"LocalDataDisk,omitnil" name:"LocalDataDisk"`
}

type OpScope struct {
	// The information of the services to operate on.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceInfoList []*ServiceBasicRestartInfo `json:"ServiceInfoList,omitnil" name:"ServiceInfoList"`
}

type OutterResource struct {
	// Specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// Specification name
	// Note: this field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// Disk type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// System disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// Memory size
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of CPUs
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Disk size
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Specification
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type PartDetailPriceItem struct {
	// The type. Valid values: `node` (node); `rootDisk` (system disk); `dataDisk` and `metaDB` (cloud data disk)
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// Rate (original)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// Rate (discounted)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealCost *float64 `json:"RealCost,omitnil" name:"RealCost"`

	// Total price (discounted)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *float64 `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// Discount
	// Note: This field may return null, indicating that no valid values can be obtained.
	Policy *float64 `json:"Policy,omitnil" name:"Policy"`

	// Quantity
	// Note: This field may return null, indicating that no valid values can be obtained.
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`
}

type PersistentVolumeContext struct {
	// Disk size in GB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Disk type. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskNum *int64 `json:"DiskNum,omitnil" name:"DiskNum"`
}

type Placement struct {
	// The ID of the availability zone where the instance resides, such as `ap-guangzhou-1`. You can call the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API and obtain this ID from the `Zone` field in the response.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Project ID of the instance. If no ID is passed in, the default project ID is used.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type PodNewParameter struct {
	// The TKE or EKS cluster ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Custom permissions
	// Examples:
	// {
	//   "apiVersion": "v1",
	//   "clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil" name:"Config"`

	// Custom parameters
	// Examples:
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil" name:"Parameter"`
}

type PodNewSpec struct {
	// The identifier of an external resource provider, such as "cls-a1cd23fa".
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil" name:"ResourceProviderIdentifier"`

	// The type of the external resource provider, such as "tke". Currently, only "tke" is supported.
	ResourceProviderType *string `json:"ResourceProviderType,omitnil" name:"ResourceProviderType"`

	// The purpose of the resource, which means the node type and can only be "TASK".
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// The number of CPUs.
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// The memory size in GB.
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// The EKS cluster - CPU type. Valid values: `intel` and `amd`.
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// The data directory mounting information of the pod node.
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil" name:"PodVolumes"`

	// Whether the dynamic spec is used. Valid values:
	// <li>`true`: Yes</li>
	// <li>`false` (default): No</li>
	EnableDynamicSpecFlag *bool `json:"EnableDynamicSpecFlag,omitnil" name:"EnableDynamicSpecFlag"`

	// The dynamic spec.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// The unique VPC ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The unique VPC subnet ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The pod name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type PodParameter struct {
	// ID of TKE or EKS cluster
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Custom permissions
	// Example:
	// {
	//   "apiVersion": "v1",
	//   "Clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil" name:"Config"`

	// Custom parameters
	// Example:
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil" name:"Parameter"`
}

type PodSpec struct {
	// Identifier of external resource provider, such as "cls-a1cd23fa".
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil" name:"ResourceProviderIdentifier"`

	// Type of external resource provider, such as "tke". Currently, only "tke" is supported.
	ResourceProviderType *string `json:"ResourceProviderType,omitnil" name:"ResourceProviderType"`

	// Purpose of the resource, which means the node type and can only be "TASK".
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// Number of CPUs
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// Memory size in GB.
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// Mount point of resources for the host. A specified mount point corresponds to the host path and is used as the data storage directory in the pod. (This parameter has been disused)
	DataVolumes []*string `json:"DataVolumes,omitnil" name:"DataVolumes"`

	// EKS cluster - CPU type. Valid values: `intel` and `amd`.
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// Data directory mounting information of the pod node.
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil" name:"PodVolumes"`

	// Whether floating specification is used. `1`: Yes; `0`: No.
	IsDynamicSpec *uint64 `json:"IsDynamicSpec,omitnil" name:"IsDynamicSpec"`

	// Floating specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// Unique VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Unique VPC subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// pod name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type PodSpecInfo struct {
	// The specified information such as pod spec and source for scale-out with pod resources.
	PodSpec *PodNewSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// The custom pod permission and parameter.
	PodParameter *PodNewParameter `json:"PodParameter,omitnil" name:"PodParameter"`
}

type PodVolume struct {
	// Storage type. Valid values: `pvc` and `hostpath`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeType *string `json:"VolumeType,omitnil" name:"VolumeType"`

	// This field will take effect if `VolumeType` is `pvc`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PVCVolume *PersistentVolumeContext `json:"PVCVolume,omitnil" name:"PVCVolume"`

	// This field will take effect if `VolumeType` is `hostpath`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostVolume *HostVolumeContext `json:"HostVolume,omitnil" name:"HostVolume"`
}

type PreExecuteFileSettings struct {
	// COS path to script, which has been disused
	Path *string `json:"Path,omitnil" name:"Path"`

	// Execution script parameter
	Args []*string `json:"Args,omitnil" name:"Args"`

	// COS bucket name, which has been disused
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// COS region name, which has been disused
	Region *string `json:"Region,omitnil" name:"Region"`

	// COS domain data, which has been disused
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Execution sequence
	RunOrder *int64 `json:"RunOrder,omitnil" name:"RunOrder"`

	// `resourceAfter` or `clusterAfter`
	WhenRun *string `json:"WhenRun,omitnil" name:"WhenRun"`

	// Script name, which has been disused
	CosFileName *string `json:"CosFileName,omitnil" name:"CosFileName"`

	// COS address of script
	CosFileURI *string `json:"CosFileURI,omitnil" name:"CosFileURI"`

	// COS `SecretId`
	CosSecretId *string `json:"CosSecretId,omitnil" name:"CosSecretId"`

	// COS `SecretKey`
	CosSecretKey *string `json:"CosSecretKey,omitnil" name:"CosSecretKey"`

	// COS `appid`, which has been disused
	AppId *string `json:"AppId,omitnil" name:"AppId"`
}

type PriceDetail struct {
	// The node ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// The price formula
	Formula *string `json:"Formula,omitnil" name:"Formula"`

	// The original price
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// The discount price
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`
}

type PriceResource struct {
	// Target specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// Disk type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *uint64 `json:"StorageType,omitnil" name:"StorageType"`

	// Disk type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// System disk size
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// Memory size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of CPUs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Disk size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// List of cloud disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil" name:"MultiDisks"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCnt *int64 `json:"DiskCnt,omitnil" name:"DiskCnt"`

	// Specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// Tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskNum *int64 `json:"DiskNum,omitnil" name:"DiskNum"`

	// Number of local disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalDiskNum *int64 `json:"LocalDiskNum,omitnil" name:"LocalDiskNum"`
}

type Resource struct {
	// Node specification description, such as CVM.SA2
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// Storage type
	// Valid values:
	// <li>4: SSD</li>
	// <li>5: Premium Cloud Storage</li>
	// <li>6: Enhanced SSD</li>
	// <li>11: High-Throughput cloud disk</li>
	// <li>12: Tremendous SSD</li>
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// Disk type
	// Valid values:
	// <li>`CLOUD_SSD`: SSD</li>
	// <li>`CLOUD_PREMIUM`: Premium Cloud Storage</li>
	// <li>`CLOUD_BASIC`: HDD</li>
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Memory capacity in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of CPU cores
	// Note: this field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Data disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// System disk capacity
	// Note: this field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// List of cloud disks. When the data disk is a cloud disk, `DiskType` and `DiskSize` are used directly; `MultiDisks` will be used for the excessive part
	// Note: this field may return null, indicating that no valid values can be obtained.
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil" name:"MultiDisks"`

	// List of tags to be bound
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Specification type, such as S2.MEDIUM8
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// Number of local disks. This field has been disused.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitnil" name:"LocalDiskNum"`

	// Number of local disks, such as 2
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskNum *uint64 `json:"DiskNum,omitnil" name:"DiskNum"`
}

// Predefined struct for user
type ScaleOutClusterRequestParams struct {
	// The node billing mode. Valid values:
	// <li>`POSTPAID_BY_HOUR`: The postpaid mode by hour.</li>
	// <li>`SPOTPAID`: The spot instance mode (for task nodes only).</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// The cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The type and number of nodes to be added.
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil" name:"ScaleOutNodeConfig"`

	// A unique random token, which is valid for 5 minutes and needs to be specified by the caller to prevent the client from repeatedly creating resources. An example value is `a9a90aa6-751a-41b6-aad6-fae36063280`.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// The details of the monthly subscription, including the instance period and auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// The [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings.
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// The services to be deployed for new nodes. By default, new nodes will inherit all services deployed for the current node type. Deployed services include default optional services. This parameter only supports optional services. For example, if `HDFS`, `YARN`, and `Impala` have been deployed for existing task nodes, only `HDFS` and `YARN` are passed in with this parameter if `Impala` is not deployed during the task node scale-out with API.
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// The processes to be deployed. All processes for services to be added are deployed by default. Deployed processes can be changed. For example, `HDFS`, `YARN`, and `Impala` have been deployed for current task nodes, and default services are `DataNode`, `NodeManager`, and `ImpalaServer`; if you want to change deployed processes, you can set this parameter to `DataNode,NodeManager,ImpalaServerCoordinator` or `DataNode,NodeManager,ImpalaServerExecutor`.
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// The list of spread placement group IDs. Only one can be specified.
	// You can call the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/product/213/17810?from_cn_redirect=1) API and obtain this parameter from the `DisasterRecoverGroupId` field in the response.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// The list of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The type of resources to add. Valid values: `host` (general CVM resources) and `pod` (resources provided by a TKE or EKS cluster).
	HardwareSourceType *string `json:"HardwareSourceType,omitnil" name:"HardwareSourceType"`

	// The pod resource information.
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil" name:"PodSpecInfo"`

	// The server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// The server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// The YARN node label specified for scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// Whether to start services after scale-out.
	// <li>`true`: Yes</li>
	// <li>`false` (default): No</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil" name:"EnableStartServiceFlag"`

	// The spec settings.
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// The ID of the AZ where the instance resides, such as `ap-guangzhou-1`. You can call the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API and obtain this ID from the `Zone` field in the response.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// The subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type ScaleOutClusterRequest struct {
	*tchttp.BaseRequest
	
	// The node billing mode. Valid values:
	// <li>`POSTPAID_BY_HOUR`: The postpaid mode by hour.</li>
	// <li>`SPOTPAID`: The spot instance mode (for task nodes only).</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// The cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The type and number of nodes to be added.
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil" name:"ScaleOutNodeConfig"`

	// A unique random token, which is valid for 5 minutes and needs to be specified by the caller to prevent the client from repeatedly creating resources. An example value is `a9a90aa6-751a-41b6-aad6-fae36063280`.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// The details of the monthly subscription, including the instance period and auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// The [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings.
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// The services to be deployed for new nodes. By default, new nodes will inherit all services deployed for the current node type. Deployed services include default optional services. This parameter only supports optional services. For example, if `HDFS`, `YARN`, and `Impala` have been deployed for existing task nodes, only `HDFS` and `YARN` are passed in with this parameter if `Impala` is not deployed during the task node scale-out with API.
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// The processes to be deployed. All processes for services to be added are deployed by default. Deployed processes can be changed. For example, `HDFS`, `YARN`, and `Impala` have been deployed for current task nodes, and default services are `DataNode`, `NodeManager`, and `ImpalaServer`; if you want to change deployed processes, you can set this parameter to `DataNode,NodeManager,ImpalaServerCoordinator` or `DataNode,NodeManager,ImpalaServerExecutor`.
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// The list of spread placement group IDs. Only one can be specified.
	// You can call the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/product/213/17810?from_cn_redirect=1) API and obtain this parameter from the `DisasterRecoverGroupId` field in the response.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// The list of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// The type of resources to add. Valid values: `host` (general CVM resources) and `pod` (resources provided by a TKE or EKS cluster).
	HardwareSourceType *string `json:"HardwareSourceType,omitnil" name:"HardwareSourceType"`

	// The pod resource information.
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil" name:"PodSpecInfo"`

	// The server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// The server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// The YARN node label specified for scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// Whether to start services after scale-out.
	// <li>`true`: Yes</li>
	// <li>`false` (default): No</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil" name:"EnableStartServiceFlag"`

	// The spec settings.
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// The ID of the AZ where the instance resides, such as `ap-guangzhou-1`. You can call the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1) API and obtain this ID from the `Zone` field in the response.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// The subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

func (r *ScaleOutClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceChargeType")
	delete(f, "InstanceId")
	delete(f, "ScaleOutNodeConfig")
	delete(f, "ClientToken")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareSourceType")
	delete(f, "PodSpecInfo")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "EnableStartServiceFlag")
	delete(f, "ResourceSpec")
	delete(f, "Zone")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutClusterResponseParams struct {
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The client token.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// The scale-out workflow ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScaleOutClusterResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutClusterResponseParams `json:"Response"`
}

func (r *ScaleOutClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// Time unit of scale-out. Valid values:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: Month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Processes unnecessary for scale-out.
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil" name:"UnNecessaryNodeList"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// Deployed service.
	// <li>`SoftDeployInfo` and `ServiceNodeInfo` are in the same group and mutually exclusive with `UnNecessaryNodeList`.</li>
	// <li>The combination of `SoftDeployInfo` and `ServiceNodeInfo` is recommended.</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// Started process.
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// List of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Resource type selected for scaling. Valid values: `host` (general CVM resource) and `pod` (resource provided by TKE or EKS cluster).
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// Specified information such as pod specification and source for scale-out with pod resources.
	PodSpec *PodSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// Server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// Server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// Yarn node label specified for rule-based scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// Custom pod permission and parameter
	PodParameter *PodParameter `json:"PodParameter,omitnil" name:"PodParameter"`

	// Number of master nodes to be added.
	// When a ClickHouse cluster is scaled, this parameter does not take effect.
	// When a Kafka cluster is scaled, this parameter does not take effect.
	// When `HardwareResourceType` is `pod`, this parameter does not take effect.
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// Whether to start the service after scale-out. `true`: Yes; `false`: No.
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil" name:"StartServiceAfterScaleOut"`

	// AZ, which defaults to the primary AZ of the cluster.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Pre-defined configuration set
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil" name:"ScaleOutServiceConfAssign"`

	// Whether to enable auto-renewal. Valid values: `0` (no), `1` (yes).
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of scale-out. Valid values:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: Month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Processes unnecessary for scale-out.
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil" name:"UnNecessaryNodeList"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// Deployed service.
	// <li>`SoftDeployInfo` and `ServiceNodeInfo` are in the same group and mutually exclusive with `UnNecessaryNodeList`.</li>
	// <li>The combination of `SoftDeployInfo` and `ServiceNodeInfo` is recommended.</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// Started process.
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// List of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Resource type selected for scaling. Valid values: `host` (general CVM resource) and `pod` (resource provided by TKE or EKS cluster).
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// Specified information such as pod specification and source for scale-out with pod resources.
	PodSpec *PodSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// Server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// Server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// Yarn node label specified for rule-based scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// Custom pod permission and parameter
	PodParameter *PodParameter `json:"PodParameter,omitnil" name:"PodParameter"`

	// Number of master nodes to be added.
	// When a ClickHouse cluster is scaled, this parameter does not take effect.
	// When a Kafka cluster is scaled, this parameter does not take effect.
	// When `HardwareResourceType` is `pod`, this parameter does not take effect.
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// Whether to start the service after scale-out. `true`: Yes; `false`: No.
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil" name:"StartServiceAfterScaleOut"`

	// AZ, which defaults to the primary AZ of the cluster.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Pre-defined configuration set
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil" name:"ScaleOutServiceConfAssign"`

	// Whether to enable auto-renewal. Valid values: `0` (no), `1` (yes).
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "InstanceId")
	delete(f, "PayMode")
	delete(f, "ClientToken")
	delete(f, "PreExecutedFileSettings")
	delete(f, "TaskCount")
	delete(f, "CoreCount")
	delete(f, "UnNecessaryNodeList")
	delete(f, "RouterCount")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareResourceType")
	delete(f, "PodSpec")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "PodParameter")
	delete(f, "MasterCount")
	delete(f, "StartServiceAfterScaleOut")
	delete(f, "ZoneId")
	delete(f, "SubnetId")
	delete(f, "ScaleOutServiceConfAssign")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Order number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// Client token.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Scale-out workflow ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Big order ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutNodeConfig struct {
	// Valid values of node type:
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// The number of nodes.
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`
}

type SceneSoftwareConfig struct {
	// The list of deployed components. The list of component options varies by `ProductVersion` (EMR version). For more information, see [Component Version](https://intl.cloud.tencent.com/document/product/589/20279?from_cn_redirect=1).
	// The instance type, `hive` or `flink`.
	Software []*string `json:"Software,omitnil" name:"Software"`

	// The scenario name, which defaults to `Hadoop-Default`. For more details, see [here](https://intl.cloud.tencent.com/document/product/589/14624?from_cn_redirect=1). Valid values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	// Hadoop-Default
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`
}

type ScriptBootstrapActionConfig struct {
	// The COS URL of the script, in the format of `https://beijing-111111.cos.ap-beijing.myqcloud.com/data/test.sh`. For the COS bucket list, see [Bucket List](https://console.cloud.tencent.com/cos/bucket).
	CosFileURI *string `json:"CosFileURI,omitnil" name:"CosFileURI"`

	// The execution time of the bootstrap action script. Valid values:
	// <li>`resourceAfter`: After node initialization</li>
	// <li>`clusterAfter`: After cluster start</li>
	// <li>`clusterBefore`: Before cluster start</li>
	ExecutionMoment *string `json:"ExecutionMoment,omitnil" name:"ExecutionMoment"`

	// The execution script parameter. The parameter format must comply with standard shell specifications.
	Args []*string `json:"Args,omitnil" name:"Args"`

	// The script file name.
	CosFileName *string `json:"CosFileName,omitnil" name:"CosFileName"`
}

type SearchItem struct {
	// Searchable type
	SearchType *string `json:"SearchType,omitnil" name:"SearchType"`

	// Searchable value
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type ServiceBasicRestartInfo struct {
	// The service name (required), such as HDFS.
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// If it is left empty, all processes will be operated on.
	ComponentInfoList []*ComponentBasicRestartInfo `json:"ComponentInfoList,omitnil" name:"ComponentInfoList"`
}

type ShortNodeInfo struct {
	// Node type: Master/Core/Task/Router/Common
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// Number of nodes
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeSize *uint64 `json:"NodeSize,omitnil" name:"NodeSize"`
}

type SoftDependInfo struct {
	// The component name.
	SoftName *string `json:"SoftName,omitnil" name:"SoftName"`

	// Whether the component is required.
	Required *bool `json:"Required,omitnil" name:"Required"`
}

// Predefined struct for user
type StartStopServiceOrMonitorRequestParams struct {
	// The cluster ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The operation type. Valid values:
	// <li>StartService: Start service</li>
	// <li>StopService: Stop service</li>
	// <li>StartMonitor: Start maintenance</li>
	// <li>StopMonitor: Stop maintenance</li>
	// <li>RestartService: Restart service. If this type is selected, "StrategyConfig" is required.</li>
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// The operation scope.
	OpScope *OpScope `json:"OpScope,omitnil" name:"OpScope"`

	// The operation policy.
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil" name:"StrategyConfig"`
}

type StartStopServiceOrMonitorRequest struct {
	*tchttp.BaseRequest
	
	// The cluster ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The operation type. Valid values:
	// <li>StartService: Start service</li>
	// <li>StopService: Stop service</li>
	// <li>StartMonitor: Start maintenance</li>
	// <li>StopMonitor: Stop maintenance</li>
	// <li>RestartService: Restart service. If this type is selected, "StrategyConfig" is required.</li>
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// The operation scope.
	OpScope *OpScope `json:"OpScope,omitnil" name:"OpScope"`

	// The operation policy.
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil" name:"StrategyConfig"`
}

func (r *StartStopServiceOrMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpType")
	delete(f, "OpScope")
	delete(f, "StrategyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStopServiceOrMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStopServiceOrMonitorResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartStopServiceOrMonitorResponse struct {
	*tchttp.BaseResponse
	Response *StartStopServiceOrMonitorResponseParams `json:"Response"`
}

func (r *StartStopServiceOrMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyConfig struct {
	// `0`: Disable rolling restart
	// `1`: Enable rolling restart
	// Note: This field may return null, indicating that no valid values can be obtained.
	RollingRestartSwitch *int64 `json:"RollingRestartSwitch,omitnil" name:"RollingRestartSwitch"`

	// The number of nodes to be restarted per batch in rolling restart, with a maximum value of 99,999.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BatchSize *int64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// The wait time (in seconds) per batch in rolling restart, with a maximum value of 5 minutes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeWait *int64 `json:"TimeWait,omitnil" name:"TimeWait"`

	// The failure handling policy. Valid values: `0` (blocks the process) and `1` (skips).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealOnFail *int64 `json:"DealOnFail,omitnil" name:"DealOnFail"`
}

type SubnetInfo struct {
	// Subnet information (name)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// Subnet information (ID)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type TerminateClusterNodesRequestParams struct {
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The list of resources to be terminated.
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil" name:"CvmInstanceIds"`

	// Valid values of node type:
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// The graceful scale-in feature. Valid values:
	//   <li>`true`: Enabled.</li>
	//   <li>`false`: Disabled.</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil" name:"GraceDownFlag"`

	// The graceful scale-in wait time in seconds. Value range: 60–1800.
	GraceDownTime *int64 `json:"GraceDownTime,omitnil" name:"GraceDownTime"`
}

type TerminateClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The list of resources to be terminated.
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil" name:"CvmInstanceIds"`

	// Valid values of node type:
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// The graceful scale-in feature. Valid values:
	//   <li>`true`: Enabled.</li>
	//   <li>`false`: Disabled.</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil" name:"GraceDownFlag"`

	// The graceful scale-in wait time in seconds. Value range: 60–1800.
	GraceDownTime *int64 `json:"GraceDownTime,omitnil" name:"GraceDownTime"`
}

func (r *TerminateClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CvmInstanceIds")
	delete(f, "NodeFlag")
	delete(f, "GraceDownFlag")
	delete(f, "GraceDownTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateClusterNodesResponseParams struct {
	// The scale-in process ID.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateClusterNodesResponseParams `json:"Response"`
}

func (r *TerminateClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// ID of terminated node. This parameter is reserved and does not need to be configured.
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// ID of terminated node. This parameter is reserved and does not need to be configured.
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *TerminateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstanceResponseParams `json:"Response"`
}

func (r *TerminateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of resource IDs of the node to be terminated. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of resource IDs of the node to be terminated. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *TerminateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTasksResponseParams `json:"Response"`
}

func (r *TerminateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopologyInfo struct {
	// AZ ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// AZ information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Subnet information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetInfoList []*SubnetInfo `json:"SubnetInfoList,omitnil" name:"SubnetInfoList"`

	// Node information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeInfoList []*ShortNodeInfo `json:"NodeInfoList,omitnil" name:"NodeInfoList"`
}

type UpdateInstanceSettings struct {
	// Memory capacity in GB
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// Number of CPU cores
	CPUCores *uint64 `json:"CPUCores,omitnil" name:"CPUCores"`

	// Machine resource ID (EMR resource identifier)
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// Target machine specification
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type UserInfoForUserManager struct {
	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// The group to which the user belongs
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`


	PassWord *string `json:"PassWord,omitnil" name:"PassWord"`


	ReMark *string `json:"ReMark,omitnil" name:"ReMark"`
}

type UserManagerFilter struct {
	// Username
	// Note: This field may return null, indicating that no valid value can be obtained.
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type UserManagerUserBriefInfo struct {
	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// The group to which the user belongs
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// `Manager` represents an admin, and `NormalUser` represents a general user.
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// Account creation time
	// Note: This field may return null, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whether the corresponding Keytab file of the user is available for download. This parameter applies only to a Kerberos-enabled cluster.
	SupportDownLoadKeyTab *bool `json:"SupportDownLoadKeyTab,omitnil" name:"SupportDownLoadKeyTab"`

	// Download link of the Keytab file
	// Note: This field may return null, indicating that no valid value can be obtained.
	DownLoadKeyTabUrl *string `json:"DownLoadKeyTabUrl,omitnil" name:"DownLoadKeyTabUrl"`
}

type VPCSettings struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type VirtualPrivateCloud struct {
	// The VPC ID.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The subnet ID.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type ZoneDetailPriceResult struct {
	// AZ ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// Price details by node
	NodeDetailPrice []*NodeDetailPriceResult `json:"NodeDetailPrice,omitnil" name:"NodeDetailPrice"`
}

type ZoneResourceConfiguration struct {
	// The VPC configuration information. This parameter is used to specify the VPC ID, subnet ID and other information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// The instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// The specs of all nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AllNodeResourceSpec *AllNodeResourceSpec `json:"AllNodeResourceSpec,omitnil" name:"AllNodeResourceSpec"`

	// For a single AZ, `ZoneTag` can be left out. For a double-AZ mode, `ZoneTag` is set to `master` and `standby` for the first and second AZs, respectively. If there are three AZs, `ZoneTag` is set to `master`, `standby`, and `third-party` for the first, second, and third AZs, respectively. Valid values:
	//   <li>master</li>
	//   <li>standby</li>
	//   <li>third-party</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneTag *string `json:"ZoneTag,omitnil" name:"ZoneTag"`
}