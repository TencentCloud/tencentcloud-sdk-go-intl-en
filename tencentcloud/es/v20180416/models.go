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

package v20180416

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type BackingIndexMetaField struct {
	// Backing index name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Backing index status
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// Backing index size
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// Current lifecycle phase of backing index
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexPhrase *string `json:"IndexPhrase,omitnil,omitempty" name:"IndexPhrase"`

	// Backing index creation time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`
}

type ClusterView struct {
	// Cluster health status
	Health *float64 `json:"Health,omitnil,omitempty" name:"Health"`

	// Whether the cluster is visible
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// Whether the cluster encounters circuit breaking
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// Average disk usage
	AvgDiskUsage *float64 `json:"AvgDiskUsage,omitnil,omitempty" name:"AvgDiskUsage"`

	// Average memory usage
	AvgMemUsage *float64 `json:"AvgMemUsage,omitnil,omitempty" name:"AvgMemUsage"`

	// Average CPU usage
	AvgCpuUsage *float64 `json:"AvgCpuUsage,omitnil,omitempty" name:"AvgCpuUsage"`

	// Total disk size of the cluster
	TotalDiskSize *uint64 `json:"TotalDiskSize,omitnil,omitempty" name:"TotalDiskSize"`

	// Types of nodes to receive client requests
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// Number of online nodes
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Total number of nodes
	TotalNodeNum *int64 `json:"TotalNodeNum,omitnil,omitempty" name:"TotalNodeNum"`

	// Number of data nodes
	DataNodeNum *int64 `json:"DataNodeNum,omitnil,omitempty" name:"DataNodeNum"`

	// Number of indices
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// Number of documents
	DocNum *int64 `json:"DocNum,omitnil,omitempty" name:"DocNum"`

	// Used disk size (in bytes)
	DiskUsedInBytes *int64 `json:"DiskUsedInBytes,omitnil,omitempty" name:"DiskUsedInBytes"`

	// Number of shards
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// Number of primary shards
	PrimaryShardNum *int64 `json:"PrimaryShardNum,omitnil,omitempty" name:"PrimaryShardNum"`

	// Number of relocating shards
	RelocatingShardNum *int64 `json:"RelocatingShardNum,omitnil,omitempty" name:"RelocatingShardNum"`

	// Number of initializing shards
	InitializingShardNum *int64 `json:"InitializingShardNum,omitnil,omitempty" name:"InitializingShardNum"`

	// Number of unassigned shards
	UnassignedShardNum *int64 `json:"UnassignedShardNum,omitnil,omitempty" name:"UnassignedShardNum"`

	// Total COS storage of an enterprise cluster, in GB
	TotalCosStorage *int64 `json:"TotalCosStorage,omitnil,omitempty" name:"TotalCosStorage"`

	// Name of the COS bucket that stores searchable snapshots of an enterprise cluster
	// Note: This field may return `null`, indicating that no valid value was found.
	SearchableSnapshotCosBucket *string `json:"SearchableSnapshotCosBucket,omitnil,omitempty" name:"SearchableSnapshotCosBucket"`

	// COS app ID of the searchable snapshots of an enterprise cluster
	// Note: This field may return `null`, indicating that no valid value was found.
	SearchableSnapshotCosAppId *string `json:"SearchableSnapshotCosAppId,omitnil,omitempty" name:"SearchableSnapshotCosAppId"`
}

type CosBackup struct {
	// Whether to enable auto-backup to COS
	IsAutoBackup *bool `json:"IsAutoBackup,omitnil,omitempty" name:"IsAutoBackup"`

	// Auto-backup time (accurate down to the hour), such as "22:00"
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`
}

// Predefined struct for user
type CreateIndexRequestParams struct {
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to create. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to create
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// JSON-formatted index metadata to create, such as `mappings` and `settings`
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to create. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to create
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// JSON-formatted index metadata to create, such as `mappings` and `settings`
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "IndexMetaJson")
	delete(f, "Username")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateIndexResponseParams `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// Availability Zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance version. Valid values: `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1`, `7.10.1`
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Access password, which must contain 8 to 16 characters, and include at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2-50)
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Billing mode <li>POSTPAID_BY_HOUR: Pay-as-you-go hourly </li>Default value: POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// This parameter is not used on the global website
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// This parameter is not used on the global website
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node storage type <li>CLOUD_SSD: SSD cloud storage </li><li>CLOUD_PREMIUM: premium cloud storage </li>Default value: CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// This parameter is not used on the global website
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Whether to create a dedicated primary node <li>true: yes </li><li>false: no </li>Default value: false
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated primary nodes (only 3 and 5 are supported. This value must be passed in if `EnableDedicatedMaster` is `true`)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node type, which must be passed in if `EnableDedicatedMaster` is `true` <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node disk size in GB, which is optional. If passed in, it can only be 50 and cannot be customized currently
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// ClusterName in the cluster configuration file, which is the instance ID by default and currently cannot be customized
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// Cluster deployment mode <li>0: single-AZ deployment </li><li>1: multi-AZ deployment </li>Default value: 0
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Details of AZs in multi-AZ deployment mode (which is required when DeployMode is 1)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Scenario template type. 0: not enabled; 1: general; 2: log; 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Visual node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Valid values: `https`, `http` (default)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The maintenance time slot
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether to enable the storage-computing separation feature.
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// Whether to enable enhanced SSD
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// Whether to enable smart inspection.
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Availability Zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance version. Valid values: `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1`, `7.10.1`
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Access password, which must contain 8 to 16 characters, and include at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2-50)
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Billing mode <li>POSTPAID_BY_HOUR: Pay-as-you-go hourly </li>Default value: POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// This parameter is not used on the global website
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// This parameter is not used on the global website
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node storage type <li>CLOUD_SSD: SSD cloud storage </li><li>CLOUD_PREMIUM: premium cloud storage </li>Default value: CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// This parameter is not used on the global website
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Whether to create a dedicated primary node <li>true: yes </li><li>false: no </li>Default value: false
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated primary nodes (only 3 and 5 are supported. This value must be passed in if `EnableDedicatedMaster` is `true`)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node type, which must be passed in if `EnableDedicatedMaster` is `true` <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node disk size in GB, which is optional. If passed in, it can only be 50 and cannot be customized currently
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// ClusterName in the cluster configuration file, which is the instance ID by default and currently cannot be customized
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// Cluster deployment mode <li>0: single-AZ deployment </li><li>1: multi-AZ deployment </li>Default value: 0
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Details of AZs in multi-AZ deployment mode (which is required when DeployMode is 1)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Scenario template type. 0: not enabled; 1: general; 2: log; 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Visual node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Valid values: `https`, `http` (default)
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The maintenance time slot
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether to enable the storage-computing separation feature.
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// Whether to enable enhanced SSD
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// Whether to enable smart inspection.
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`
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
	delete(f, "Zone")
	delete(f, "EsVersion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "InstanceName")
	delete(f, "NodeNum")
	delete(f, "ChargeType")
	delete(f, "ChargePeriod")
	delete(f, "RenewFlag")
	delete(f, "NodeType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "TimeUnit")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "EnableDedicatedMaster")
	delete(f, "MasterNodeNum")
	delete(f, "MasterNodeType")
	delete(f, "MasterNodeDiskSize")
	delete(f, "ClusterNameInConf")
	delete(f, "DeployMode")
	delete(f, "MultiZoneInfo")
	delete(f, "LicenseType")
	delete(f, "NodeInfoList")
	delete(f, "TagList")
	delete(f, "BasicSecurityType")
	delete(f, "SceneType")
	delete(f, "WebNodeTypeInfo")
	delete(f, "Protocol")
	delete(f, "OperationDuration")
	delete(f, "EnableHybridStorage")
	delete(f, "DiskEnhance")
	delete(f, "EnableDiagnose")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Order ID
	// Note: This field may return `null`, indicating that no valid value was found.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DeleteIndexRequestParams struct {
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to delete. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to delete
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Backing index name
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to delete. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to delete
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Backing index name
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "BackingIndexName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIndexResponseParams `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListRequestParams struct {
	// Index type. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index name. `null` indicates that all indexes are requested.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The starting position of paging
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting condition field, which can be `IndexName`, `IndexStorage`, or `IndexCreateTime`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filtering by index status
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// Sorting mode, which can be `asc` and `desc`.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
	
	// Index type. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index name. `null` indicates that all indexes are requested.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The starting position of paging
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting condition field, which can be `IndexName`, `IndexStorage`, or `IndexCreateTime`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filtering by index status
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// Sorting mode, which can be `asc` and `desc`.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeIndexListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IndexType")
	delete(f, "InstanceId")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "IndexStatusList")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListResponseParams struct {
	// Index metadata field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexMetaFields []*IndexMetaField `json:"IndexMetaFields,omitnil,omitempty" name:"IndexMetaFields"`

	// Total number of results
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexListResponseParams `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexMetaRequestParams struct {
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index type. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name. `null` indicates that all indexes are requested.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DescribeIndexMetaRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index type. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name. `null` indicates that all indexes are requested.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *DescribeIndexMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexMetaResponseParams struct {
	// Index metadata field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexMetaField *IndexMetaField `json:"IndexMetaField,omitnil,omitempty" name:"IndexMetaField"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexMetaResponseParams `json:"Response"`
}

func (r *DescribeIndexMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Default value: 1
	// <li>1: primary log</li>
	// <li>2: search slow log</li>
	// <li>3: index slow log</li>
	// <li>4: GC log</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Search keyword, which supports LUCENE syntax, such as `level:WARN`, `ip:1.1.1.1`, and `message:test-index`
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Log start time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 100. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Time sorting order. Default value: 0
	// <li>0: descending</li>
	// <li>1: ascending</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Default value: 1
	// <li>1: primary log</li>
	// <li>2: search slow log</li>
	// <li>3: index slow log</li>
	// <li>4: GC log</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Search keyword, which supports LUCENE syntax, such as `level:WARN`, `ip:1.1.1.1`, and `message:test-index`
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Log start time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 100. Maximum value: 100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Time sorting order. Default value: 0
	// <li>0: descending</li>
	// <li>1: ascending</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "SearchKey")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsResponseParams struct {
	// Number of returned logs
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Log details list
	InstanceLogList []*InstanceLog `json:"InstanceLogList,omitnil,omitempty" name:"InstanceLogList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination start value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, such as "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, such as "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination start value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsResponseParams struct {
	// Total number of operation records
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Operation history
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// AZ of the cluster instance. If this is not passed in, all AZs are used by default
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// List of cluster instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster instance names
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The sorting field. <li>1: Instance ID </li><li>2: Instance name </li><li>3: AZ </li><li>4: Creation time </li>If `OrderByKey` is not passed in, sorting is performed by creation time in descending order.
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// Sorting order <li>0: ascending </li><li>1: descending </li>If orderByKey is passed in but orderByType is not, ascending order is used by default
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// VPC VIP list
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// List of availability zones
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// The health status filter. Valid values: `0` (green), `1` (yellow), `2` (red), `-1` (unknown).
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// VPC IDs
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// AZ of the cluster instance. If this is not passed in, all AZs are used by default
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// List of cluster instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster instance names
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The sorting field. <li>1: Instance ID </li><li>2: Instance name </li><li>3: AZ </li><li>4: Creation time </li>If `OrderByKey` is not passed in, sorting is performed by creation time in descending order.
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// Sorting order <li>0: ascending </li><li>1: descending </li>If orderByKey is passed in but orderByType is not, ascending order is used by default
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// VPC VIP list
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// List of availability zones
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// The health status filter. Valid values: `0` (green), `1` (yellow), `2` (red), `-1` (unknown).
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// VPC IDs
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`
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
	delete(f, "Zone")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByKey")
	delete(f, "OrderByType")
	delete(f, "TagList")
	delete(f, "IpList")
	delete(f, "ZoneList")
	delete(f, "HealthStatus")
	delete(f, "VpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of returned instances
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instance details
	InstanceList []*InstanceInfo `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeViewsRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeViewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsResponseParams struct {
	// Cluster view
	// Note: This field may return `null`, indicating that no valid value was found.
	ClusterView *ClusterView `json:"ClusterView,omitnil,omitempty" name:"ClusterView"`

	// Node view
	// Note: This field may return `null`, indicating that no valid value was found.
	NodesView []*NodeView `json:"NodesView,omitnil,omitempty" name:"NodesView"`

	// Kibana view
	// Note: This field may return `null`, indicating that no valid value was found.
	KibanasView []*KibanaView `json:"KibanasView,omitnil,omitempty" name:"KibanasView"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeViewsResponseParams `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DictInfo struct {
	// Dictionary key value
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dictionary name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Dictionary size in B
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type EsAcl struct {
	// Kibana access blocklist
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// Kibana access allowlist
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

type EsConfigSetInfo struct {
	// Configuration set type, such as `LDAP` and `AD`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// "{\"order\":0,\"url\":\"ldap://10.0.1.72:389\",\"bind_dn\":\"cn=admin,dc=tencent,dc=com\",\"user_search.base_dn\":\"dc=tencent,dc=com\",\"user_search.filter\":\"(cn={0})\",\"group_search.base_dn\":\"dc=tencent,dc=com\"}"
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`
}

type EsDictionaryInfo struct {
	// List of non-stop words
	MainDict []*DictInfo `json:"MainDict,omitnil,omitempty" name:"MainDict"`

	// List of stop words
	Stopwords []*DictInfo `json:"Stopwords,omitnil,omitempty" name:"Stopwords"`

	// QQ dictionary list
	QQDict []*DictInfo `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// Synonym dictionary list
	Synonym []*DictInfo `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// Update dictionary type
	UpdateType *string `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`
}

type EsPublicAcl struct {
	// Access blocklist
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// Access allowlist
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

// Predefined struct for user
type GetRequestTargetNodeTypesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GetRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *GetRequestTargetNodeTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestTargetNodeTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRequestTargetNodeTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRequestTargetNodeTypesResponseParams struct {
	// A list of node types used to receive requests.
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRequestTargetNodeTypesResponse struct {
	*tchttp.BaseResponse
	Response *GetRequestTargetNodeTypesResponseParams `json:"Response"`
}

func (r *GetRequestTargetNodeTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestTargetNodeTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IndexMetaField struct {
	// Index type
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Index status
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// Index size (in byte)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// Index creation time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`

	// Backing index
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BackingIndices []*BackingIndexMetaField `json:"BackingIndices,omitnil,omitempty" name:"BackingIndices"`

	// Cluster ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Cluster version
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// Index lifecycle policy field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexPolicyField *IndexPolicyField `json:"IndexPolicyField,omitnil,omitempty" name:"IndexPolicyField"`

	// Index automation field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexOptionsField *IndexOptionsField `json:"IndexOptionsField,omitnil,omitempty" name:"IndexOptionsField"`

	// Index setting field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IndexSettingsField *IndexSettingsField `json:"IndexSettingsField,omitnil,omitempty" name:"IndexSettingsField"`

	// Cluster APP ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The number of index docs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexDocs *uint64 `json:"IndexDocs,omitnil,omitempty" name:"IndexDocs"`
}

type IndexOptionsField struct {
	// Max age for expiry purpose
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireMaxAge *string `json:"ExpireMaxAge,omitnil,omitempty" name:"ExpireMaxAge"`

	// Max size for expiry purpose
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ExpireMaxSize *string `json:"ExpireMaxSize,omitnil,omitempty" name:"ExpireMaxSize"`

	// Rollover cycle
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RolloverMaxAge *string `json:"RolloverMaxAge,omitnil,omitempty" name:"RolloverMaxAge"`

	// Whether to enable the dynamic rollover
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RolloverDynamic *string `json:"RolloverDynamic,omitnil,omitempty" name:"RolloverDynamic"`

	// Whether to enable dynamic sharding
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ShardNumDynamic *string `json:"ShardNumDynamic,omitnil,omitempty" name:"ShardNumDynamic"`

	// Timestamp field
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TimestampField *string `json:"TimestampField,omitnil,omitempty" name:"TimestampField"`

	// Write mode
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WriteMode *string `json:"WriteMode,omitnil,omitempty" name:"WriteMode"`
}

type IndexPolicyField struct {
	// Whether to enable the warm phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WarmEnable *string `json:"WarmEnable,omitnil,omitempty" name:"WarmEnable"`

	// Min age before data transitions to the warm phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	WarmMinAge *string `json:"WarmMinAge,omitnil,omitempty" name:"WarmMinAge"`

	// Whether to enable the cold phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ColdEnable *string `json:"ColdEnable,omitnil,omitempty" name:"ColdEnable"`

	// Min age before data transitions to the cold phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ColdMinAge *string `json:"ColdMinAge,omitnil,omitempty" name:"ColdMinAge"`

	// Whether to enable the frozen phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FrozenEnable *string `json:"FrozenEnable,omitnil,omitempty" name:"FrozenEnable"`

	// Min age before data transitions to the frozen phase
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	FrozenMinAge *string `json:"FrozenMinAge,omitnil,omitempty" name:"FrozenMinAge"`

	// /
	// Note: This field may return null, indicating that no valid value can be obtained.
	ColdAction *string `json:"ColdAction,omitnil,omitempty" name:"ColdAction"`
}

type IndexSettingsField struct {
	// Number of primary shards
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NumberOfShards *string `json:"NumberOfShards,omitnil,omitempty" name:"NumberOfShards"`

	// Number of replica shards
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NumberOfReplicas *string `json:"NumberOfReplicas,omitnil,omitempty" name:"NumberOfReplicas"`

	// Index refresh interval
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RefreshInterval *string `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`
}

type InstanceInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability Zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// User ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// UID of the VPC where the instance resides
	VpcUid *string `json:"VpcUid,omitnil,omitempty" name:"VpcUid"`

	// UID of the subnet where the instance resides
	SubnetUid *string `json:"SubnetUid,omitnil,omitempty" name:"SubnetUid"`

	// Instance status. `0`: Processing; `1`: Normal; `-1`: `Stopped`; `-2`: Being terminated; `-3`: Terminated; `2`: Initializing during the cluster creation.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// This parameter is not used on the global website
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Instance billing method. Valid values: POSTPAID_BY_HOUR (pay-as-you-go hourly); CDHPAID (billed based on CDH, i.e., only CDH is billed but not the instances on CDH)
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// This parameter is not used on the global website
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Number of nodes
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Number of CPU cores of the node
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// Node memory size in GB
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Node disk type
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// ES domain name
	EsDomain *string `json:"EsDomain,omitnil,omitempty" name:"EsDomain"`

	// ES VIP
	EsVip *string `json:"EsVip,omitnil,omitempty" name:"EsVip"`

	// ES port
	EsPort *uint64 `json:"EsPort,omitnil,omitempty" name:"EsPort"`

	// Kibana access URL
	KibanaUrl *string `json:"KibanaUrl,omitnil,omitempty" name:"KibanaUrl"`

	// ES version number
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// ES configuration item
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Kibana access control configuration
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of the instance
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// This parameter is not used on the global website
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// Instance type (instance type identifier, which can be only 1 or 2 currently)
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// IK analyzer configuration
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitnil,omitempty" name:"IkConfig"`

	// Dedicated primary node configuration
	MasterNodeInfo *MasterNodeInfo `json:"MasterNodeInfo,omitnil,omitempty" name:"MasterNodeInfo"`

	// Auto-backup to COS configuration
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Whether to allow auto-backup to COS
	AllowCosBackup *bool `json:"AllowCosBackup,omitnil,omitempty" name:"AllowCosBackup"`

	// List of tags owned by the instance
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether it is a hot/warm cluster <li>true: yes </li><li>false: no</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitnil,omitempty" name:"EnableHotWarmMode"`

	// Warm node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmNodeType *string `json:"WarmNodeType,omitnil,omitempty" name:"WarmNodeType"`

	// Number of warm nodes
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitnil,omitempty" name:"WarmNodeNum"`

	// Number of warm node CPU cores
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitnil,omitempty" name:"WarmCpuNum"`

	// Warm node memory size (in GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmMemSize *uint64 `json:"WarmMemSize,omitnil,omitempty" name:"WarmMemSize"`

	// Warm node disk type
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmDiskType *string `json:"WarmDiskType,omitnil,omitempty" name:"WarmDiskType"`

	// Warm node disk size (in GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitnil,omitempty" name:"WarmDiskSize"`

	// Cluster node information list
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// ES public IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	EsPublicUrl *string `json:"EsPublicUrl,omitnil,omitempty" name:"EsPublicUrl"`

	// Multi-AZ network information
	// Note: This field may return null, indicating that no valid values can be obtained.
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// Deployment mode <li>0: single-AZ </li><li>1: multi-AZ</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// ES public access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// ES public access control configuration
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana private IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitnil,omitempty" name:"KibanaPrivateUrl"`

	// Kibana public access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana private access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityType *uint64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// Scenario template type. 0: not enabled; 1: general scenario; 2: log scenario; 3: search scenario
	// Note: this field may return null, indicating that no valid values can be obtained.
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana configuration item.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Kibana node information
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	KibanaNodeInfo *KibanaNodeInfo `json:"KibanaNodeInfo,omitnil,omitempty" name:"KibanaNodeInfo"`

	// Visual node configuration
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// JDK type. Valid values: `oracle`, `kona`
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Jdk *string `json:"Jdk,omitnil,omitempty" name:"Jdk"`

	// Cluster network communication protocol
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Security group ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// Cold node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdNodeType *string `json:"ColdNodeType,omitnil,omitempty" name:"ColdNodeType"`

	// Number of cold nodes
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdNodeNum *uint64 `json:"ColdNodeNum,omitnil,omitempty" name:"ColdNodeNum"`

	// Number of cold node CPU cores
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdCpuNum *uint64 `json:"ColdCpuNum,omitnil,omitempty" name:"ColdCpuNum"`

	// Cold node memory size (in GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdMemSize *uint64 `json:"ColdMemSize,omitnil,omitempty" name:"ColdMemSize"`

	// Cold node disk type
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdDiskType *string `json:"ColdDiskType,omitnil,omitempty" name:"ColdDiskType"`

	// Cold node disk size (in GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	ColdDiskSize *uint64 `json:"ColdDiskSize,omitnil,omitempty" name:"ColdDiskSize"`

	// Frozen node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenNodeType *string `json:"FrozenNodeType,omitnil,omitempty" name:"FrozenNodeType"`

	// Number of frozen nodes
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenNodeNum *uint64 `json:"FrozenNodeNum,omitnil,omitempty" name:"FrozenNodeNum"`

	// Number of frozen node CPU cores
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenCpuNum *uint64 `json:"FrozenCpuNum,omitnil,omitempty" name:"FrozenCpuNum"`

	// Frozen node memory size (GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenMemSize *uint64 `json:"FrozenMemSize,omitnil,omitempty" name:"FrozenMemSize"`

	// Frozen node disk type
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenDiskType *string `json:"FrozenDiskType,omitnil,omitempty" name:"FrozenDiskType"`

	// Frozen node disk size (in GB)
	// Note: This field may return `null`, indicating that no valid value was found.
	FrozenDiskSize *uint64 `json:"FrozenDiskSize,omitnil,omitempty" name:"FrozenDiskSize"`

	// Cluster health status. `-1`: Unknown; `0`: Green; `1`: Yellow; `2`: Red
	// Note: This field may return `null`, indicating that no valid value was found.
	HealthStatus *int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Private URL of the HTTPS cluster
	// Note: This field may return `null`, indicating that no valid value was found.
	EsPrivateUrl *string `json:"EsPrivateUrl,omitnil,omitempty" name:"EsPrivateUrl"`

	// Private domain of the HTTPS cluster
	// Note: This field may return `null`, indicating that no valid value was found.
	EsPrivateDomain *string `json:"EsPrivateDomain,omitnil,omitempty" name:"EsPrivateDomain"`

	// Configuration set info of the cluster.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EsConfigSets []*EsConfigSetInfo `json:"EsConfigSets,omitnil,omitempty" name:"EsConfigSets"`

	// The maintenance time slot of the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Web node list
	// Note: This field may return null, indicating that no valid values can be obtained.
	OptionalWebServiceInfos []*OptionalWebServiceInfo `json:"OptionalWebServiceInfos,omitnil,omitempty" name:"OptionalWebServiceInfos"`

	// Autonomous index option
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoIndexEnabled *bool `json:"AutoIndexEnabled,omitnil,omitempty" name:"AutoIndexEnabled"`

	// Whether the storage-computing separation feature is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// The process progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessPercent *float64 `json:"ProcessPercent,omitnil,omitempty" name:"ProcessPercent"`

	// The alerting policy of Kibana over the public network. <li>`OPEN`: Enable the policy;</li><li>`CLOSE`: Disable the policy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`
}

type InstanceLog struct {
	// Log time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Log level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Cluster node IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Log content
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type KeyValue struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KibanaNodeInfo struct {
	// Kibana node specification
	KibanaNodeType *string `json:"KibanaNodeType,omitnil,omitempty" name:"KibanaNodeType"`

	// Number of Kibana nodes
	KibanaNodeNum *uint64 `json:"KibanaNodeNum,omitnil,omitempty" name:"KibanaNodeNum"`

	// Number of Kibana node's CPUs
	KibanaNodeCpuNum *uint64 `json:"KibanaNodeCpuNum,omitnil,omitempty" name:"KibanaNodeCpuNum"`

	// Kibana node's memory in GB
	KibanaNodeMemSize *uint64 `json:"KibanaNodeMemSize,omitnil,omitempty" name:"KibanaNodeMemSize"`

	// Kibana node's disk type
	KibanaNodeDiskType *string `json:"KibanaNodeDiskType,omitnil,omitempty" name:"KibanaNodeDiskType"`

	// Kibana node's disk size
	KibanaNodeDiskSize *uint64 `json:"KibanaNodeDiskSize,omitnil,omitempty" name:"KibanaNodeDiskSize"`
}

type KibanaView struct {
	// Kibana node IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Node disk size
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Disk usage
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// Node memory size
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Memory usage
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// Number of node CPUs
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// CPU usage
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type LocalDiskInfo struct {
	// Local disk type <li>LOCAL_SATA: big data </li><li>NVME_SSD: high IO</li>
	LocalDiskType *string `json:"LocalDiskType,omitnil,omitempty" name:"LocalDiskType"`

	// Size of a single local disk
	LocalDiskSize *uint64 `json:"LocalDiskSize,omitnil,omitempty" name:"LocalDiskSize"`

	// Number of local disks
	LocalDiskCount *uint64 `json:"LocalDiskCount,omitnil,omitempty" name:"LocalDiskCount"`
}

type MasterNodeInfo struct {
	// Whether to enable the dedicated primary node
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// Dedicated primary node specification <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// Number of dedicated primary nodes
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// Number of CPU cores of the dedicated primary node
	MasterNodeCpuNum *uint64 `json:"MasterNodeCpuNum,omitnil,omitempty" name:"MasterNodeCpuNum"`

	// Memory size of the dedicated primary node in GB
	MasterNodeMemSize *uint64 `json:"MasterNodeMemSize,omitnil,omitempty" name:"MasterNodeMemSize"`

	// Disk size of the dedicated primary node in GB
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Disk type of the dedicated primary node
	MasterNodeDiskType *string `json:"MasterNodeDiskType,omitnil,omitempty" name:"MasterNodeDiskType"`
}

type NodeInfo struct {
	// Number of nodes
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Node type<li>`hotData`: hot data node</li>
	// <li>`warmData`: warm data node</li>
	// <li>`dedicatedMaster`: dedicated master node</li>
	// Default value: `hotData`
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Node disk type <li>CLOUD_SSD: SSD cloud storage </li><li>CLOUD_PREMIUM: Premium cloud disk </li>Default value: CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Local disk information
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocalDiskInfo *LocalDiskInfo `json:"LocalDiskInfo,omitnil,omitempty" name:"LocalDiskInfo"`

	// Number of node disks
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Whether to encrypt node disk. 0: no (default); 1: yes.
	DiskEncrypt *uint64 `json:"DiskEncrypt,omitnil,omitempty" name:"DiskEncrypt"`

	// CPU number
	// Note: This field may return null, indicating that no valid values can be obtained.
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// Memory size in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// /
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskEnhance *int64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`
}

type NodeView struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node IP
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// Whether the node is visible
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// Whether the node encounters circuit breaking
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// Node disk size
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Disk usage
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// Node memory size (in GB)
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Memory usage
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// Number of node CPUs
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// CPU usage
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Node HTTP IP
	NodeHttpIp *string `json:"NodeHttpIp,omitnil,omitempty" name:"NodeHttpIp"`

	// JVM memory usage
	JvmMemUsage *float64 `json:"JvmMemUsage,omitnil,omitempty" name:"JvmMemUsage"`

	// Number of node shards
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// ID list of node disks
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// Whether a hidden availability zone
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`
}

type Operation struct {
	// Unique operation ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Operation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Operation type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Operation details
	Detail *OperationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Operation result
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Workflow task information
	Tasks []*TaskDetail `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Operation progress
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type OperationDetail struct {
	// Original instance configuration information
	OldInfo []*KeyValue `json:"OldInfo,omitnil,omitempty" name:"OldInfo"`

	// Updated instance configuration information
	NewInfo []*KeyValue `json:"NewInfo,omitnil,omitempty" name:"NewInfo"`
}

type OperationDuration struct {
	// Maintenance period, which can be one or more days from Monday to Sunday. Valid values: [0, 6].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// The maintenance start time
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// The maintenance end time
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// The time zone expressed in UTC.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type OperationDurationUpdated struct {
	// Maintenance period, which can be one or more days from Monday to Sunday. Valid values: [0, 6].
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// The maintenance start time
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// The maintenance end time
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// The time zone expressed in UTC.
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// The array of ES cluster IDs
	MoreInstances []*string `json:"MoreInstances,omitnil,omitempty" name:"MoreInstances"`
}

type OptionalWebServiceInfo struct {
	// Type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Public URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicUrl *string `json:"PublicUrl,omitnil,omitempty" name:"PublicUrl"`

	// Private URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateUrl *string `json:"PrivateUrl,omitnil,omitempty" name:"PrivateUrl"`

	// Public network access
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Private network access
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to force restart <li>true: Yes </li><li>false: No </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Restart mode. `0`: rolling restart; `1`: full restart
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to force restart <li>true: Yes </li><li>false: No </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Restart mode. `0`: rolling restart; `1`: full restart
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ForceRestart")
	delete(f, "RestartMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartInstanceResponseParams `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartKibanaRequestParams struct {
	// ES instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartKibanaRequest struct {
	*tchttp.BaseRequest
	
	// ES instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *RestartKibanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartKibanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartKibanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartKibanaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartKibanaResponse struct {
	*tchttp.BaseResponse
	Response *RestartKibanaResponseParams `json:"Response"`
}

func (r *RestartKibanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartKibanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node name list
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// Whether to force restart
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// The restart mode. Valid values: `in-place` (default), `blue-green`.
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// The node status, applicable in the blue/green mode. The blue/green restart is risky if the node is offline.
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`
}

type RestartNodesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node name list
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// Whether to force restart
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// The restart mode. Valid values: `in-place` (default), `blue-green`.
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// The node status, applicable in the blue/green mode. The blue/green restart is risky if the node is offline.
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`
}

func (r *RestartNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeNames")
	delete(f, "ForceRestart")
	delete(f, "RestartMode")
	delete(f, "IsOffline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartNodesResponse struct {
	*tchttp.BaseResponse
	Response *RestartNodesResponseParams `json:"Response"`
}

func (r *RestartNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubTaskDetail struct {
	// Subtask name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Subtask result
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// Subtask error message
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// Subtask type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Subtask status. 0: processing, 1: succeeded, -1: failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Name of the index for which the check for upgrade failed
	FailedIndices []*string `json:"FailedIndices,omitnil,omitempty" name:"FailedIndices"`

	// Subtask end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Subtask level. 1: warning, 2: failed
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskDetail struct {
	// Task name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Task progress
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task completion time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Subtask
	SubTasks []*SubTaskDetail `json:"SubTasks,omitnil,omitempty" name:"SubTasks"`

	// The task time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ElapsedTime *int64 `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`
}

// Predefined struct for user
type UpdateDictionariesRequestParams struct {
	// ES instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// COS address of the main dictionary for the IK analyzer
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// COS address of the stopword dictionary for the IK analyzer
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// COS address of the synonym dictionary
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// COS address of the QQ dictionary
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// `0` (default): Install, `1`: Delete
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// Whether to force restart the cluster. The default value is `false`.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpdateDictionariesRequest struct {
	*tchttp.BaseRequest
	
	// ES instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// COS address of the main dictionary for the IK analyzer
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// COS address of the stopword dictionary for the IK analyzer
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// COS address of the synonym dictionary
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// COS address of the QQ dictionary
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// `0` (default): Install, `1`: Delete
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// Whether to force restart the cluster. The default value is `false`.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

func (r *UpdateDictionariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictionariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IkMainDicts")
	delete(f, "IkStopwords")
	delete(f, "Synonym")
	delete(f, "QQDict")
	delete(f, "UpdateType")
	delete(f, "ForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDictionariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDictionariesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDictionariesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDictionariesResponseParams `json:"Response"`
}

func (r *UpdateDictionariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictionariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIndexRequestParams struct {
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to update. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to update
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// JSON-formatted index metadata to update, such as `mappings` and `settings`.
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to roll over the backup index
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`
}

type UpdateIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Type of the index to update. `auto`: Automated; `normal`: General.
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Name of the index to update
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// JSON-formatted index metadata to update, such as `mappings` and `settings`.
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// Username for cluster access
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Password for cluster access
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Whether to roll over the backup index
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`
}

func (r *UpdateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "UpdateMetaJson")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "RolloverBackingIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIndexResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateIndexResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIndexResponseParams `json:"Response"`
}

func (r *UpdateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2-50)
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES configuration item (JSON string)
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Password of the default user 'elastic', which must contain 8 to 16 characters, including at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The policy for visual component (Kibana and Cerebro) access over public network.
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated primary nodes (only 3 and 5 are supported)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node specification <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node disk size in GB. This is 50 GB by default and currently cannot be customized
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Whether to force restart during configuration update <li>true: Yes </li><li>false: No </li>This needs to be set only for EsConfig. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Auto-backup to COS
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Node information list. You can pass in only the nodes to be updated and their corresponding specification information. Supported operations include: <li>modifying the number of nodes in the same type </li><li>modifying the specification and disk size of nodes in the same type </li><li>adding a node type (you must also specify the node type, quantity, specification, disk, etc.) </li>The above operations can only be performed one at a time, and the disk type cannot be modified
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// The status of ES cluster access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Public network ACL
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// The status of Kibana access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// The status of Kibana access over private network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// Enables or disables user authentication for ES Basic Edition v6.8 and above
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana private port
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: scaling in blue/green deployment mode without cluster restart (default); 1: scaling by unmounting disk with rolling cluster restart
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// Multi-AZ deployment
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// Scenario template type. -1: not enabled; 1: general; 2: log; 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana configuration item (JSON string)
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Visual node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Whether to switch to the new network architecture
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// Whether to enable Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// The status of Cerebro access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// The status of Cerebro access over private network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// Added or modified configuration set information
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// The maintenance time slot
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether to enable the option for sending alerting messages over the public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2-50)
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES configuration item (JSON string)
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Password of the default user 'elastic', which must contain 8 to 16 characters, including at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The policy for visual component (Kibana and Cerebro) access over public network.
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated primary nodes (only 3 and 5 are supported)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node specification <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated primary node disk size in GB. This is 50 GB by default and currently cannot be customized
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Whether to force restart during configuration update <li>true: Yes </li><li>false: No </li>This needs to be set only for EsConfig. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Auto-backup to COS
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Node information list. You can pass in only the nodes to be updated and their corresponding specification information. Supported operations include: <li>modifying the number of nodes in the same type </li><li>modifying the specification and disk size of nodes in the same type </li><li>adding a node type (you must also specify the node type, quantity, specification, disk, etc.) </li>The above operations can only be performed one at a time, and the disk type cannot be modified
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// The status of ES cluster access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Public network ACL
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// The status of Kibana access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// The status of Kibana access over private network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// Enables or disables user authentication for ES Basic Edition v6.8 and above
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana private port
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: scaling in blue/green deployment mode without cluster restart (default); 1: scaling by unmounting disk with rolling cluster restart
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// Multi-AZ deployment
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// Scenario template type. -1: not enabled; 1: general; 2: log; 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana configuration item (JSON string)
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Visual node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Whether to switch to the new network architecture
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// Whether to enable Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// The status of Cerebro access over public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// The status of Cerebro access over private network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// Added or modified configuration set information
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// The maintenance time slot
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether to enable the option for sending alerting messages over the public network.
	// `OPEN`: Enabled.
	// `CLOSE`: Disabled.
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`
}

func (r *UpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "NodeNum")
	delete(f, "EsConfig")
	delete(f, "Password")
	delete(f, "EsAcl")
	delete(f, "DiskSize")
	delete(f, "NodeType")
	delete(f, "MasterNodeNum")
	delete(f, "MasterNodeType")
	delete(f, "MasterNodeDiskSize")
	delete(f, "ForceRestart")
	delete(f, "CosBackup")
	delete(f, "NodeInfoList")
	delete(f, "PublicAccess")
	delete(f, "EsPublicAcl")
	delete(f, "KibanaPublicAccess")
	delete(f, "KibanaPrivateAccess")
	delete(f, "BasicSecurityType")
	delete(f, "KibanaPrivatePort")
	delete(f, "ScaleType")
	delete(f, "MultiZoneInfo")
	delete(f, "SceneType")
	delete(f, "KibanaConfig")
	delete(f, "WebNodeTypeInfo")
	delete(f, "SwitchPrivateLink")
	delete(f, "EnableCerebro")
	delete(f, "CerebroPublicAccess")
	delete(f, "CerebroPrivateAccess")
	delete(f, "EsConfigSet")
	delete(f, "OperationDuration")
	delete(f, "KibanaAlteringPublicAccess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInstanceResponseParams struct {
	// Order ID
	// Note: This field may return `null`, indicating that no valid value was found.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateInstanceResponseParams `json:"Response"`
}

func (r *UpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePluginsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of names of the plugins to be installed
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// List of names of the plugins to be uninstalled
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// Whether to force restart the cluster. The default value is `false`.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Whether to reinstall the cluster. The default value is `false`.
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0: system plugin
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

type UpdatePluginsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of names of the plugins to be installed
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// List of names of the plugins to be uninstalled
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// Whether to force restart the cluster. The default value is `false`.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Whether to reinstall the cluster. The default value is `false`.
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0: system plugin
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

func (r *UpdatePluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstallPluginList")
	delete(f, "RemovePluginList")
	delete(f, "ForceRestart")
	delete(f, "ForceUpdate")
	delete(f, "PluginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePluginsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePluginsResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePluginsResponseParams `json:"Response"`
}

func (r *UpdatePluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRequestTargetNodeTypesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A list of node types used to receive requests.
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`
}

type UpdateRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// A list of node types used to receive requests.
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`
}

func (r *UpdateRequestTargetNodeTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRequestTargetNodeTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TargetNodeTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRequestTargetNodeTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRequestTargetNodeTypesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRequestTargetNodeTypesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRequestTargetNodeTypesResponseParams `json:"Response"`
}

func (r *UpdateRequestTargetNodeTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRequestTargetNodeTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target ES version. Valid values: 6.4.3, 6.8.2, 7.5.1
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// Whether to check for upgrade only. Default value: false
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// Target X-Pack edition: <li>OSS: Open-source Edition </li><li>basic: Basic Edition </li>Currently only used for v5.6.4 to v6.x upgrade. Default value: basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Upgrade mode. <li>scale: blue/green deployment</li><li>restart: rolling restart</li>Default value: scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Whether to back up the cluster before version upgrade (no backup by default)
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Whether to skip the check and perform a force restart in the rolling mode. Default value: `false`.
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target ES version. Valid values: 6.4.3, 6.8.2, 7.5.1
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// Whether to check for upgrade only. Default value: false
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// Target X-Pack edition: <li>OSS: Open-source Edition </li><li>basic: Basic Edition </li>Currently only used for v5.6.4 to v6.x upgrade. Default value: basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Upgrade mode. <li>scale: blue/green deployment</li><li>restart: rolling restart</li>Default value: scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Whether to back up the cluster before version upgrade (no backup by default)
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Whether to skip the check and perform a force restart in the rolling mode. Default value: `false`.
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EsVersion")
	delete(f, "CheckOnly")
	delete(f, "LicenseType")
	delete(f, "BasicSecurityType")
	delete(f, "UpgradeMode")
	delete(f, "CosBackup")
	delete(f, "SkipCheckForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLicenseRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Whether to force restart <li>true: yes </li><li>false: no </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Whether to force restart <li>true: yes </li><li>false: no </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

func (r *UpgradeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LicenseType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "BasicSecurityType")
	delete(f, "ForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLicenseResponseParams struct {
	// Order ID
	// Note: This field may return `null`, indicating that no valid value was found.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLicenseResponseParams `json:"Response"`
}

func (r *UpgradeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebNodeTypeInfo struct {
	// Number of visual nodes. The value is always `1`.
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Visual node specification
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type ZoneDetail struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}