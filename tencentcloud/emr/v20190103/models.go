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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddUsersForUserManagerRequestParams struct {
	// Cluster string ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// User information list
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitempty" name:"UserManagerUserList"`
}

type AddUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// Cluster string ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// User information list
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitempty" name:"UserManagerUserList"`
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
	SuccessUserList []*string `json:"SuccessUserList,omitempty" name:"SuccessUserList"`

	// The user list that is not successfully added
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FailedUserList []*string `json:"FailedUserList,omitempty" name:"FailedUserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type ApplicationStatics struct {
	// Queue name
	Queue *string `json:"Queue,omitempty" name:"Queue"`

	// Username
	User *string `json:"User,omitempty" name:"User"`

	// Application type
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// `SumMemorySeconds` meaning
	SumMemorySeconds *int64 `json:"SumMemorySeconds,omitempty" name:"SumMemorySeconds"`


	SumVCoreSeconds *int64 `json:"SumVCoreSeconds,omitempty" name:"SumVCoreSeconds"`

	// SumHDFSBytesWritten (with unit)
	SumHDFSBytesWritten *string `json:"SumHDFSBytesWritten,omitempty" name:"SumHDFSBytesWritten"`

	// SumHDFSBytesRead (with unit)
	SumHDFSBytesRead *string `json:"SumHDFSBytesRead,omitempty" name:"SumHDFSBytesRead"`

	// Application count
	CountApps *int64 `json:"CountApps,omitempty" name:"CountApps"`
}

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

type ClusterExternalServiceInfo struct {
	// Dependency. `0`: Other clusters depend on the current cluster. `1`: The current cluster depends on other clusters.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DependType *int64 `json:"DependType,omitempty" name:"DependType"`

	// Shared component
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Service *string `json:"Service,omitempty" name:"Service"`

	// Sharing cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Sharing cluster status
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Hive metadata
	// Note: this field may return null, indicating that no valid values can be obtained.
	HiveMetaDb *string `json:"HiveMetaDb,omitempty" name:"HiveMetaDb"`

	// Cluster type: EMR, CLICKHOUSE, DRUID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceClass *string `json:"ServiceClass,omitempty" name:"ServiceClass"`

	// Alias serialization of all nodes in cluster
	// Note: this field may return null, indicating that no valid values can be obtained.
	AliasInfo *string `json:"AliasInfo,omitempty" name:"AliasInfo"`

	// Cluster version ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// Availability zone
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Scenario name
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// Scenario-based cluster type
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneServiceClass *string `json:"SceneServiceClass,omitempty" name:"SceneServiceClass"`

	// Scenario-based EMR version
	// Note: This field may return `null`, indicating that no valid value was found.
	SceneEmrVersion *string `json:"SceneEmrVersion,omitempty" name:"SceneEmrVersion"`

	// Scenario-based cluster type
	// Note: This field may return `null`, indicating that no valid value was found.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value was found.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Subnet name
	// Note: This field may return `null`, indicating that no valid value was found.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Cluster dependency
	// Note: This field may return `null`, indicating that no valid value was found.
	ClusterExternalServiceInfo []*ClusterExternalServiceInfo `json:"ClusterExternalServiceInfo,omitempty" name:"ClusterExternalServiceInfo"`

	// The VPC ID string type of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// The subnet ID string type of the cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TopologyInfoList []*TopologyInfo `json:"TopologyInfoList,omitempty" name:"TopologyInfoList"`

	// Multi-AZ cluster
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitempty" name:"IsMultiZoneCluster"`
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1</li>
	// <li>2: EMR v2.0.1</li>
	// <li>4: EMR v2.1.0</li>
	// <li>7: EMR v3.0.0</li>
	// <li>9: EMR v2.2.0</li>
	// <li>11: ClickHouse v1.0.0</li>
	// <li>13: Druid v1.0.0</li>
	// <li>15: EMR v2.2.1</li>
	// <li>16: EMR v2.3.0</li>
	// <li>17: ClickHouse v1.1.0</li>
	// <li>19: EMR v2.4.0</li>
	// <li>20: EMR v2.5.0</li>
	// <li>22: ClickHouse v1.2.0</li>
	// <li>24: EMR TianQiong v1.0.0</li>
	// <li>25: EMR v3.1.0</li>
	// <li>26: Doris v1.0.0</li>
	// <li>27: Kafka v1.0.0</li>
	// <li>28: EMR v3.2.0</li>
	// <li>29: EMR v2.5.1</li>
	// <li>30: EMR v2.6.0</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// List of deployed components. The list of component options varies by EMR product ID (i.e., `ProductId`; for specific meanings, please see the `ProductId` input parameter). For more information, please see [Component Version](https://intl.cloud.tencent.com/document/product/589/20279?from_cn_redirect=1).
	// Enter an instance value: `hive` or `flink`.
	Software []*string `json:"Software,omitempty" name:"Software"`

	// Whether to enable high node availability. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// Instance name.
	// <li>Length limit: 6-36 characters.</li>
	// <li>Only letters, numbers, dashes (-), and underscores (_) are supported.</li>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Instance login settings. This parameter allows you to set the login password or key for your purchased node.
	// <li>If the key is set, the password will be only used for login to the native component WebUI.</li>
	// <li>If the key is not set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// Node resource specification.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// Parameter required for enabling COS access.
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Security group to which an instance belongs in the format of `sg-xxxxxxxx`. This parameter can be obtained from the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) API.
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// List of spread placement group IDs. Only one can be specified currently.
	// This parameter can be obtained in the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1) API.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

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

	// Custom application role.
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitempty" name:"ExternalService"`


	VersionID *int64 `json:"VersionID,omitempty" name:"VersionID"`

	// `true` indicates that the multi-AZ deployment mode is enabled. This parameter is available only in cluster creation and cannot be changed after setting.
	MultiZone *bool `json:"MultiZone,omitempty" name:"MultiZone"`

	// Node resource specs. The actual number of AZs is set, with the first AZ as the primary AZ, the second as the backup AZ, and the third as the arbitrator AZ. If the multi-AZ mode is not enabled, set the value to `1`.
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitempty" name:"MultiZoneSettings"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Product ID. Different product IDs represent different EMR product versions. Valid values:
	// <li>1: EMR v1.3.1</li>
	// <li>2: EMR v2.0.1</li>
	// <li>4: EMR v2.1.0</li>
	// <li>7: EMR v3.0.0</li>
	// <li>9: EMR v2.2.0</li>
	// <li>11: ClickHouse v1.0.0</li>
	// <li>13: Druid v1.0.0</li>
	// <li>15: EMR v2.2.1</li>
	// <li>16: EMR v2.3.0</li>
	// <li>17: ClickHouse v1.1.0</li>
	// <li>19: EMR v2.4.0</li>
	// <li>20: EMR v2.5.0</li>
	// <li>22: ClickHouse v1.2.0</li>
	// <li>24: EMR TianQiong v1.0.0</li>
	// <li>25: EMR v3.1.0</li>
	// <li>26: Doris v1.0.0</li>
	// <li>27: Kafka v1.0.0</li>
	// <li>28: EMR v3.2.0</li>
	// <li>29: EMR v2.5.1</li>
	// <li>30: EMR v2.6.0</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// List of deployed components. The list of component options varies by EMR product ID (i.e., `ProductId`; for specific meanings, please see the `ProductId` input parameter). For more information, please see [Component Version](https://intl.cloud.tencent.com/document/product/589/20279?from_cn_redirect=1).
	// Enter an instance value: `hive` or `flink`.
	Software []*string `json:"Software,omitempty" name:"Software"`

	// Whether to enable high node availability. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// Instance name.
	// <li>Length limit: 6-36 characters.</li>
	// <li>Only letters, numbers, dashes (-), and underscores (_) are supported.</li>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Instance login settings. This parameter allows you to set the login password or key for your purchased node.
	// <li>If the key is set, the password will be only used for login to the native component WebUI.</li>
	// <li>If the key is not set, the password will be used for login to all purchased nodes and the native component WebUI.</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Configuration information of VPC. This parameter is used to specify the VPC ID, subnet ID, etc.
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// Node resource specification.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// Parameter required for enabling COS access.
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// Instance location. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Security group to which an instance belongs in the format of `sg-xxxxxxxx`. This parameter can be obtained from the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) API.
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// [Bootstrap action](https://intl.cloud.tencent.com/document/product/589/35656?from_cn_redirect=1) script settings
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// List of spread placement group IDs. Only one can be specified currently.
	// This parameter can be obtained in the `SecurityGroupId` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1) API.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

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

	// Custom application role.
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitempty" name:"ExternalService"`

	VersionID *int64 `json:"VersionID,omitempty" name:"VersionID"`

	// `true` indicates that the multi-AZ deployment mode is enabled. This parameter is available only in cluster creation and cannot be changed after setting.
	MultiZone *bool `json:"MultiZone,omitempty" name:"MultiZone"`

	// Node resource specs. The actual number of AZs is set, with the first AZ as the primary AZ, the second as the backup AZ, and the third as the arbitrator AZ. If the multi-AZ mode is not enabled, set the value to `1`.
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitempty" name:"MultiZoneSettings"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type CustomMetaInfo struct {
	// JDBC connection to custom MetaDB instance beginning with `jdbc:mysql://`
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitempty" name:"MetaDataJdbcUrl"`

	// Custom MetaDB instance username
	MetaDataUser *string `json:"MetaDataUser,omitempty" name:"MetaDataUser"`

	// Custom MetaDB instance password
	MetaDataPass *string `json:"MetaDataPass,omitempty" name:"MetaDataPass"`
}

type CustomServiceDefine struct {
	// Custom parameter key
	Name *string `json:"Name,omitempty" name:"Name"`

	// Custom parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type DescribeClusterNodesRequestParams struct {
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

	// Resource type. Valid values: all, host, pod. Default value: all
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// Searchable field
	SearchFields []*SearchItem `json:"SearchFields,omitempty" name:"SearchFields"`
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

	// Resource type. Valid values: all, host, pod. Default value: all
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// Searchable field
	SearchFields []*SearchItem `json:"SearchFields,omitempty" name:"SearchFields"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodesResponseParams struct {
	// Total number of queried nodes
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// List of node details
	// Note: this field may return null, indicating that no valid values can be obtained.
	NodeList []*NodeHardwareInfo `json:"NodeList,omitempty" name:"NodeList"`

	// List of tag keys owned by user
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Resource type list
	// Note: this field may return null, indicating that no valid values can be obtained.
	HardwareResourceTypeList []*string `json:"HardwareResourceTypeList,omitempty" name:"HardwareResourceTypeList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Queue name used for filtering
	Queues []*string `json:"Queues,omitempty" name:"Queues"`

	// Username used for filtering
	Users []*string `json:"Users,omitempty" name:"Users"`

	// Application type used for filtering
	ApplicationTypes []*string `json:"ApplicationTypes,omitempty" name:"ApplicationTypes"`

	// Group field. Valid values: `queue`, `user`, and `applicationType`.
	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`

	// Sorting field. Valid values: `sumMemorySeconds`, `sumVCoreSeconds`, `sumHDFSBytesWritten`, and `sumHDFSBytesRead`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Order type. Valid values: `0` (descending) and `1`(ascending).
	IsAsc *int64 `json:"IsAsc,omitempty" name:"IsAsc"`

	// Page number
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Page limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeEmrApplicationStaticsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time in the format of timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// Queue name used for filtering
	Queues []*string `json:"Queues,omitempty" name:"Queues"`

	// Username used for filtering
	Users []*string `json:"Users,omitempty" name:"Users"`

	// Application type used for filtering
	ApplicationTypes []*string `json:"ApplicationTypes,omitempty" name:"ApplicationTypes"`

	// Group field. Valid values: `queue`, `user`, and `applicationType`.
	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`

	// Sorting field. Valid values: `sumMemorySeconds`, `sumVCoreSeconds`, `sumHDFSBytesWritten`, and `sumHDFSBytesRead`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Order type. Valid values: `0` (descending) and `1`(ascending).
	IsAsc *int64 `json:"IsAsc,omitempty" name:"IsAsc"`

	// Page number
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Page limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	Statics []*ApplicationStatics `json:"Statics,omitempty" name:"Statics"`

	// Total count
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Available queue name
	Queues []*string `json:"Queues,omitempty" name:"Queues"`

	// Available usernames
	Users []*string `json:"Users,omitempty" name:"Users"`

	// Available application type
	ApplicationTypes []*string `json:"ApplicationTypes,omitempty" name:"ApplicationTypes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// Page number. Default value: `0`, indicating the first page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: `10`; maximum value: `100`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting field. Valid values: <li>clusterId: Sorting by instance ID. </li><li>addTime: Sorting by instance creation time.</li><li>status: Sorting by instance status code.</li>
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Sort ascending or descending based on `OrderField`. Valid values:<li>0: Descending.</li><li>1: Ascending.</li>Default value: `0`.
	Asc *int64 `json:"Asc,omitempty" name:"Asc"`

	// Custom query
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeInstancesListRequest struct {
	*tchttp.BaseRequest
	
	// Cluster filtering policy. Valid values: <li>clusterList: Queries the list of clusters excluding terminated ones.</li><li>monitorManage: Queries the list of clusters excluding those terminated, under creation and not successfully created.</li><li>cloudHardwareManage/componentManage: Two reserved values, which have the same implications as those of `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// Page number. Default value: `0`, indicating the first page.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results per page. Default value: `10`; maximum value: `100`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting field. Valid values: <li>clusterId: Sorting by instance ID. </li><li>addTime: Sorting by instance creation time.</li><li>status: Sorting by instance status code.</li>
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Sort ascending or descending based on `OrderField`. Valid values:<li>0: Descending.</li><li>1: Ascending.</li>Default value: `0`.
	Asc *int64 `json:"Asc,omitempty" name:"Asc"`

	// Custom query
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
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
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// Cluster instance list.
	InstancesList []*EmrListInstance `json:"InstancesList,omitempty" name:"InstancesList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// Queries by one or more instance IDs in the format of `emr-xxxxxxxx`. For the format of this parameter, please see the `id.N` section in [API Overview](https://intl.cloud.tencent.com/document/api/213/15688). If no instance ID is entered, the list of all instances under this `APPID` will be returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster filtering policy. Valid values:
	// <li>clusterList: queries the list of clusters except terminated ones.</li>
	// <li>monitorManage: queries the list of clusters except those that have been terminated, are being created, or failed to be created.</li>
	// <li>cloudHardwareManage/componentManage: reserved fields with the same meaning as `monitorManage`.</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// Queries by one or more instance IDs in the format of `emr-xxxxxxxx`. For the format of this parameter, please see the `id.N` section in [API Overview](https://intl.cloud.tencent.com/document/api/213/15688). If no instance ID is entered, the list of all instances under this `APPID` will be returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// List of EMR instance details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitempty" name:"ClusterList"`

	// List of tag keys associated to an instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeResourceScheduleRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	OpenSwitch *bool `json:"OpenSwitch,omitempty" name:"OpenSwitch"`

	// The resource scheduler in service
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// Fair Scheduler information
	FSInfo *string `json:"FSInfo,omitempty" name:"FSInfo"`

	// Capacity Scheduler information
	CSInfo *string `json:"CSInfo,omitempty" name:"CSInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Page number
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Page size
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// User list query filter
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitempty" name:"UserManagerFilter"`

	// Whether the Keytab file information is required. This field is only valid for clusters with Kerberos enabled and defaults to `false`.
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitempty" name:"NeedKeytabInfo"`
}

type DescribeUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Page number
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// Page size
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// User list query filter
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitempty" name:"UserManagerFilter"`

	// Whether the Keytab file information is required. This field is only valid for clusters with Kerberos enabled and defaults to `false`.
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitempty" name:"NeedKeytabInfo"`
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
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// User information list
	// Note: This field may return null, indicating that no valid value can be obtained.
	UserManagerUserList []*UserManagerUserBriefInfo `json:"UserManagerUserList,omitempty" name:"UserManagerUserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type DynamicPodSpec struct {
	// Minimum number of CPUs
	RequestCpu *float64 `json:"RequestCpu,omitempty" name:"RequestCpu"`

	// Maximum number of CPUs
	LimitCpu *float64 `json:"LimitCpu,omitempty" name:"LimitCpu"`

	// Minimum memory in MB
	RequestMemory *float64 `json:"RequestMemory,omitempty" name:"RequestMemory"`

	// Maximum memory in MB
	LimitMemory *float64 `json:"LimitMemory,omitempty" name:"LimitMemory"`
}

type EmrListInstance struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Status description
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster region
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// User APPID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Creation time
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Running time
	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`

	// Cluster IP
	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`

	// Cluster version
	EmrVersion *string `json:"EmrVersion,omitempty" name:"EmrVersion"`

	// Cluster billing mode
	ChargeType *uint64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// EMR ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Product ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// Project ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Subnet ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetId *uint64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VpcId *uint64 `json:"VpcId,omitempty" name:"VpcId"`

	// Region
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Status code
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Instance tag
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Alarm information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AlarmInfo *string `json:"AlarmInfo,omitempty" name:"AlarmInfo"`

	// Whether it is a Woodpecker cluster
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsWoodpeckerCluster *uint64 `json:"IsWoodpeckerCluster,omitempty" name:"IsWoodpeckerCluster"`

	// VPC name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Subnet name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// VPC ID string
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Subnet ID string
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// Cluster type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClusterClass *string `json:"ClusterClass,omitempty" name:"ClusterClass"`

	// Whether it is a multi-AZ cluster
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitempty" name:"IsMultiZoneCluster"`
}

type EmrProductConfigOutter struct {
	// Software information
	// Note: this field may return null, indicating that no valid values can be obtained.
	SoftInfo []*string `json:"SoftInfo,omitempty" name:"SoftInfo"`

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

	// Custom application role
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`

	// Security groups
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// SSH key ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PublicKeyId *string `json:"PublicKeyId,omitempty" name:"PublicKeyId"`
}

type ExternalService struct {
	// Shared component type, which can be EMR or CUSTOM
	ShareType *string `json:"ShareType,omitempty" name:"ShareType"`

	// Custom parameters
	CustomServiceDefineList []*CustomServiceDefine `json:"CustomServiceDefineList,omitempty" name:"CustomServiceDefineList"`

	// Shared component name
	Service *string `json:"Service,omitempty" name:"Service"`

	// Shared component cluster
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type Filters struct {
	// Field name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filters by the field value
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type HostVolumeContext struct {
	// The directory for mounting the host in the pod, which is the mount point of the host in the resource. A specified mount point corresponds to the host path and is used as the data storage directory in the pod.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumePath *string `json:"VolumePath,omitempty" name:"VolumePath"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable high availability of node. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// List of deployed components. Different required components need to be selected for different EMR product IDs (i.e., `ProductId`; for specific meanings, please see the `ProductId` field in the input parameter):
	// <li>When `ProductId` is 1, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 2, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 4, the required components include hadoop-2.8.4, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 7, the required components include hadoop-3.1.2, knox-1.2.0, and zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitempty" name:"Software"`

	// Node specification queried for price.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

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

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitempty" name:"ExternalService"`


	VersionID *uint64 `json:"VersionID,omitempty" name:"VersionID"`

	// AZ specs
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitempty" name:"MultiZoneSettings"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of instance purchase duration. Valid values:
	// <li>s: seconds. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Purchase duration of instance, which needs to be used together with `TimeUnit`.
	// <li>When `TimeUnit` is `s`, this parameter can only be filled with 3600, indicating a pay-as-you-go instance.</li>
	// <li>When `TimeUnit` is `m`, the number entered in this parameter indicates the purchase duration of the monthly-subscription instance; for example, 1 means one month</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Instance billing mode. Valid values:
	// <li>0: pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable high availability of node. Valid values:
	// <li>0: does not enable high availability of node.</li>
	// <li>1: enables high availability of node.</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// List of deployed components. Different required components need to be selected for different EMR product IDs (i.e., `ProductId`; for specific meanings, please see the `ProductId` field in the input parameter):
	// <li>When `ProductId` is 1, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 2, the required components include hadoop-2.7.3, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 4, the required components include hadoop-2.8.4, knox-1.2.0, and zookeeper-3.4.9</li>
	// <li>When `ProductId` is 7, the required components include hadoop-3.1.2, knox-1.2.0, and zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitempty" name:"Software"`

	// Node specification queried for price.
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

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

	// Scenario-based values:
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// Shared component information
	ExternalService []*ExternalService `json:"ExternalService,omitempty" name:"ExternalService"`

	VersionID *uint64 `json:"VersionID,omitempty" name:"VersionID"`

	// AZ specs
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitempty" name:"MultiZoneSettings"`
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
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// List of resource IDs of the node to be renewed. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Location of the instance. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance billing mode.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Unit of time for instance renewal.
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Whether to change from pay-as-you-go billing to monthly subscription billing. `0`: no; `1`: yes
	ModifyPayMode *int64 `json:"ModifyPayMode,omitempty" name:"ModifyPayMode"`
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// How long the instance will be renewed for, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// List of resource IDs of the node to be renewed. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Location of the instance. This parameter is used to specify the AZ, project, and other attributes of the instance.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance billing mode.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Unit of time for instance renewal.
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Whether to change from pay-as-you-go billing to monthly subscription billing. `0`: no; `1`: yes
	ModifyPayMode *int64 `json:"ModifyPayMode,omitempty" name:"ModifyPayMode"`
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
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// ID of the AZ where an instance resides.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// Number of master nodes to be added.
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of scale-out. Valid value:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	// <li>When `PayMode` is 0, `TimeSpan` can only be 3,600.</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// ID of the AZ where an instance resides.
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Currency.
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// Number of master nodes to be added.
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`
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
	OriginalCost *string `json:"OriginalCost,omitempty" name:"OriginalCost"`

	// Discounted price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`

	// Time unit of scale-out. Valid value:
	// <li>s: Second.</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// Node spec queried for price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceSpec *PriceResource `json:"PriceSpec,omitempty" name:"PriceSpec"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceResponseParams struct {
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

type LoginSettings struct {
	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public Key
	PublicKeyId *string `json:"PublicKeyId,omitempty" name:"PublicKeyId"`
}

// Predefined struct for user
type ModifyResourceScheduleConfigRequestParams struct {
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Business identifier. `fair`: Edit fair configuration items; `fairPlan`: Edit the execution plan; `capacity`: Edit capacity configuration items.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Modified module information
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyResourceScheduleConfigRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Business identifier. `fair`: Edit fair configuration items; `fairPlan`: Edit the execution plan; `capacity`: Edit capacity configuration items.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Modified module information
	Value *string `json:"Value,omitempty" name:"Value"`
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
	IsDraft *bool `json:"IsDraft,omitempty" name:"IsDraft"`

	// Verification error information. If it is not null, the verification fails and thus the configuration fails.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The original scheduler: `fair`
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// The new scheduler: `capacity`
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
}

type ModifyResourceSchedulerRequest struct {
	*tchttp.BaseRequest
	
	// EMR cluster ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The original scheduler: `fair`
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// The new scheduler: `capacity`
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type MultiZoneSetting struct {
	// "master", "standby", "third-party"
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ZoneTag *string `json:"ZoneTag,omitempty" name:"ZoneTag"`

	// None
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// None
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// None
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`
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

	// Node type. 0: common node; 1: master node;
	// 2: core node; 3: task node
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
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitempty" name:"MCMultiDisk"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Wether the node is auto-scaling. 0 means common node. 1 means auto-scaling node.
	AutoFlag *int64 `json:"AutoFlag,omitempty" name:"AutoFlag"`

	// Resource type. Valid values: host, pod
	// Note: this field may return null, indicating that no valid values can be obtained.
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// Whether floating specification is used. `1`: yes; `0`: no
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	IsDynamicSpec *int64 `json:"IsDynamicSpec,omitempty" name:"IsDynamicSpec"`

	// Floating specification in JSON string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DynamicPodSpec *string `json:"DynamicPodSpec,omitempty" name:"DynamicPodSpec"`

	// Whether to support billing mode change. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SupportModifyPayMode *int64 `json:"SupportModifyPayMode,omitempty" name:"SupportModifyPayMode"`

	// System disk type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RootStorageType *int64 `json:"RootStorageType,omitempty" name:"RootStorageType"`

	// AZ information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Subnet
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitempty" name:"SubnetInfo"`

	// Client
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Clients *string `json:"Clients,omitempty" name:"Clients"`
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

type PersistentVolumeContext struct {
	// Disk size in GB.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Disk type. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskNum *int64 `json:"DiskNum,omitempty" name:"DiskNum"`
}

type Placement struct {
	// ID of the project to which the instance belongs. This parameter can be obtained from the `projectId` field in the return value of the `DescribeProject` API. If 0 is entered, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// AZ where the instance resides, such as ap-guangzhou-1. You can call the `DescribeZones` API and see the `Zone` field to get the value of this parameter.
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type PodParameter struct {
	// ID of TKE or EKS cluster
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

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
	Config *string `json:"Config,omitempty" name:"Config"`

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
	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`
}

type PodSpec struct {
	// Identifier of external resource provider, such as "cls-a1cd23fa".
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitempty" name:"ResourceProviderIdentifier"`

	// Type of external resource provider, such as "tke". Currently, only "tke" is supported.
	ResourceProviderType *string `json:"ResourceProviderType,omitempty" name:"ResourceProviderType"`

	// Purpose of the resource, which means the node type and can only be "TASK".
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// Number of CPUs
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in GB.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Mount point of resources for the host. A specified mount point corresponds to the host path and is used as the data storage directory in the pod. (This parameter has been disused)
	DataVolumes []*string `json:"DataVolumes,omitempty" name:"DataVolumes"`

	// EKS cluster - CPU type. Valid values: `intel` and `amd`.
	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`

	// Data directory mounting information of the pod node.
	PodVolumes []*PodVolume `json:"PodVolumes,omitempty" name:"PodVolumes"`

	// Whether floating specification is used. `1`: Yes; `0`: No.
	IsDynamicSpec *uint64 `json:"IsDynamicSpec,omitempty" name:"IsDynamicSpec"`

	// Floating specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitempty" name:"DynamicPodSpec"`

	// Unique VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Unique VPC subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// pod name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type PodVolume struct {
	// Storage type. Valid values: `pvc` and `hostpath`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`

	// This field will take effect if `VolumeType` is `pvc`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PVCVolume *PersistentVolumeContext `json:"PVCVolume,omitempty" name:"PVCVolume"`

	// This field will take effect if `VolumeType` is `hostpath`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostVolume *HostVolumeContext `json:"HostVolume,omitempty" name:"HostVolume"`
}

type PreExecuteFileSettings struct {
	// COS path to script, which has been disused
	Path *string `json:"Path,omitempty" name:"Path"`

	// Execution script parameter
	Args []*string `json:"Args,omitempty" name:"Args"`

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
	// Note: This field may return null, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Disk type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *uint64 `json:"StorageType,omitempty" name:"StorageType"`

	// Disk type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk size
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// Memory size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of CPUs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Disk size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// List of cloud disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCnt *int64 `json:"DiskCnt,omitempty" name:"DiskCnt"`

	// Specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Tag
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskNum *int64 `json:"DiskNum,omitempty" name:"DiskNum"`

	// Number of local disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LocalDiskNum *int64 `json:"LocalDiskNum,omitempty" name:"LocalDiskNum"`
}

type Resource struct {
	// Node specification description, such as CVM.SA2
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// Storage type
	// Valid values:
	// <li>4: SSD</li>
	// <li>5: Premium Cloud Storage</li>
	// <li>6: Enhanced SSD</li>
	// <li>11: High-Throughput cloud disk</li>
	// <li>12: Tremendous SSD</li>
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// Disk type
	// Valid values:
	// <li>`CLOUD_SSD`: SSD</li>
	// <li>`CLOUD_PREMIUM`: Premium Cloud Storage</li>
	// <li>`CLOUD_BASIC`: HDD</li>
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks"`

	// List of tags to be bound
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Specification type, such as S2.MEDIUM8
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of local disks. This field has been disused.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitempty" name:"LocalDiskNum"`

	// Number of local disks, such as 2
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// Time unit of scale-out. Valid values:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: Month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Processes unnecessary for scale-out.
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitempty" name:"UnNecessaryNodeList"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// Deployed service.
	// <li>`SoftDeployInfo` and `ServiceNodeInfo` are in the same group and mutually exclusive with `UnNecessaryNodeList`.</li>
	// <li>The combination of `SoftDeployInfo` and `ServiceNodeInfo` is recommended.</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitempty" name:"SoftDeployInfo"`

	// Started process.
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitempty" name:"ServiceNodeInfo"`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// List of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Resource type selected for scaling. Valid values: `host` (general CVM resource) and `pod` (resource provided by TKE or EKS cluster).
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// Specified information such as pod specification and source for scale-out with pod resources.
	PodSpec *PodSpec `json:"PodSpec,omitempty" name:"PodSpec"`

	// Server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitempty" name:"ClickHouseClusterName"`

	// Server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitempty" name:"ClickHouseClusterType"`

	// Yarn node label specified for rule-based scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitempty" name:"YarnNodeLabel"`

	// Custom pod permission and parameter
	PodParameter *PodParameter `json:"PodParameter,omitempty" name:"PodParameter"`

	// Number of master nodes to be added.
	// When a ClickHouse cluster is scaled, this parameter does not take effect.
	// When a Kafka cluster is scaled, this parameter does not take effect.
	// When `HardwareResourceType` is `pod`, this parameter does not take effect.
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Whether to start the service after scale-out. `true`: Yes; `false`: No.
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitempty" name:"StartServiceAfterScaleOut"`

	// AZ, which defaults to the primary AZ of the cluster.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Pre-defined configuration set
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitempty" name:"ScaleOutServiceConfAssign"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Time unit of scale-out. Valid values:
	// <li>s: Second. When `PayMode` is 0, `TimeUnit` can only be `s`.</li>
	// <li>m: Month. When `PayMode` is 1, `TimeUnit` can only be `m`.</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Time span of scale-out, which needs to be used together with `TimeUnit`.
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance billing mode. Valid value:
	// <li>0: Pay-as-you-go.</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Client token.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Bootstrap script settings.
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// Number of task nodes to be added.
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Number of core nodes to be added.
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Processes unnecessary for scale-out.
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitempty" name:"UnNecessaryNodeList"`

	// Number of router nodes to be added.
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// Deployed service.
	// <li>`SoftDeployInfo` and `ServiceNodeInfo` are in the same group and mutually exclusive with `UnNecessaryNodeList`.</li>
	// <li>The combination of `SoftDeployInfo` and `ServiceNodeInfo` is recommended.</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitempty" name:"SoftDeployInfo"`

	// Started process.
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitempty" name:"ServiceNodeInfo"`

	// List of spread placement group IDs. Only one can be specified currently.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// List of tags bound to added nodes.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Resource type selected for scaling. Valid values: `host` (general CVM resource) and `pod` (resource provided by TKE or EKS cluster).
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// Specified information such as pod specification and source for scale-out with pod resources.
	PodSpec *PodSpec `json:"PodSpec,omitempty" name:"PodSpec"`

	// Server group name selected for ClickHouse cluster scale-out.
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitempty" name:"ClickHouseClusterName"`

	// Server group type selected for ClickHouse cluster scale-out. Valid values: `new` (create a group) and `old` (select an existing group).
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitempty" name:"ClickHouseClusterType"`

	// Yarn node label specified for rule-based scale-out.
	YarnNodeLabel *string `json:"YarnNodeLabel,omitempty" name:"YarnNodeLabel"`

	// Custom pod permission and parameter
	PodParameter *PodParameter `json:"PodParameter,omitempty" name:"PodParameter"`

	// Number of master nodes to be added.
	// When a ClickHouse cluster is scaled, this parameter does not take effect.
	// When a Kafka cluster is scaled, this parameter does not take effect.
	// When `HardwareResourceType` is `pod`, this parameter does not take effect.
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Whether to start the service after scale-out. `true`: Yes; `false`: No.
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitempty" name:"StartServiceAfterScaleOut"`

	// AZ, which defaults to the primary AZ of the cluster.
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subnet, which defaults to the subnet used when the cluster is created.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Pre-defined configuration set
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitempty" name:"ScaleOutServiceConfAssign"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Order number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// Client token.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Scale-out workflow ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Big order ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SearchItem struct {
	// Searchable type
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// Searchable value
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type ShortNodeInfo struct {
	// Node type: Master/Core/Task/Router/Common
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// Number of nodes
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeSize *uint64 `json:"NodeSize,omitempty" name:"NodeSize"`
}

type SubnetInfo struct {
	// Subnet information (name)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// Subnet information (ID)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateInstanceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of terminated node. This parameter is reserved and does not need to be configured.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of terminated node. This parameter is reserved and does not need to be configured.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of resource IDs of the node to be terminated. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of resource IDs of the node to be terminated. The resource ID is in the format of `emr-vm-xxxxxxxx`. A valid resource ID can be queried in the [console](https://console.cloud.tencent.com/emr/static/hardware).
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Subnet information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SubnetInfoList []*SubnetInfo `json:"SubnetInfoList,omitempty" name:"SubnetInfoList"`

	// Node information
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NodeInfoList []*ShortNodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList"`
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

type UserInfoForUserManager struct {
	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// The group to which the user belongs
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`


	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`


	ReMark *string `json:"ReMark,omitempty" name:"ReMark"`
}

type UserManagerFilter struct {
	// Username
	// Note: This field may return null, indicating that no valid value can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type UserManagerUserBriefInfo struct {
	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// The group to which the user belongs
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// `Manager` represents an admin, and `NormalUser` represents a general user.
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// Account creation time
	// Note: This field may return null, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the corresponding Keytab file of the user is available for download. This parameter applies only to a Kerberos-enabled cluster.
	SupportDownLoadKeyTab *bool `json:"SupportDownLoadKeyTab,omitempty" name:"SupportDownLoadKeyTab"`

	// Download link of the Keytab file
	// Note: This field may return null, indicating that no valid value can be obtained.
	DownLoadKeyTabUrl *string `json:"DownLoadKeyTabUrl,omitempty" name:"DownLoadKeyTabUrl"`
}

type VPCSettings struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}