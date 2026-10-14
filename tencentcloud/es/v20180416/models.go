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

type AutoScaleDiskInfo struct {
	// Node type. Valid values: `hotData`, `warmData`.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 0: Percentage-based scale-out; 1: Absolute value-based scale-out
	ScaleType *uint64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// Trigger threshold. Unit: %. For example, 80%.
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// Trigger duration in minutes, for example, 60
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Expansion ratio per time, unit: %, for example, 20%
	PercentSize *uint64 `json:"PercentSize,omitnil,omitempty" name:"PercentSize"`

	// Absolute value-based scale-out, in GB, for example, 100 GB
	FixSize *uint64 `json:"FixSize,omitnil,omitempty" name:"FixSize"`

	// Capacity limit for scale-out, in GB, for example, 500 GB
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type BackingIndexMetaField struct {
	// <p>Backing index name</p>
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// <p>Backup index status</p>
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// <p>Backup index storage size</p>
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// <p>Backup index current lifecycle</p>
	IndexPhrase *string `json:"IndexPhrase,omitnil,omitempty" name:"IndexPhrase"`

	// <p>Backup index creation time</p>
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`

	// <p>Index Uuid</p>
	IndexUuid *string `json:"IndexUuid,omitnil,omitempty" name:"IndexUuid"`
}

type ClusterView struct {
	// Cluster health status
	Health *float64 `json:"Health,omitnil,omitempty" name:"Health"`

	// Whether the cluster is visible
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// Whether to fuse the cluster
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// Average Disk Usage
	AvgDiskUsage *float64 `json:"AvgDiskUsage,omitnil,omitempty" name:"AvgDiskUsage"`

	// Average memory utilization
	AvgMemUsage *float64 `json:"AvgMemUsage,omitnil,omitempty" name:"AvgMemUsage"`

	// Average CPU Usage
	AvgCpuUsage *float64 `json:"AvgCpuUsage,omitnil,omitempty" name:"AvgCpuUsage"`

	// Total storage size of a cluster.
	TotalDiskSize *uint64 `json:"TotalDiskSize,omitnil,omitempty" name:"TotalDiskSize"`

	// Client request node
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// Number of online nodes
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Total number of nodes
	TotalNodeNum *int64 `json:"TotalNodeNum,omitnil,omitempty" name:"TotalNodeNum"`

	// Number of Data Nodes
	DataNodeNum *int64 `json:"DataNodeNum,omitnil,omitempty" name:"DataNodeNum"`

	// Index count
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// Number of documents
	DocNum *int64 `json:"DocNum,omitnil,omitempty" name:"DocNum"`

	// Used bytes of the disk
	DiskUsedInBytes *int64 `json:"DiskUsedInBytes,omitnil,omitempty" name:"DiskUsedInBytes"`

	// Number of shards
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// Number of primary shards
	PrimaryShardNum *int64 `json:"PrimaryShardNum,omitnil,omitempty" name:"PrimaryShardNum"`

	// Number of migrating shards
	RelocatingShardNum *int64 `json:"RelocatingShardNum,omitnil,omitempty" name:"RelocatingShardNum"`

	// Number of initializing shards.
	InitializingShardNum *int64 `json:"InitializingShardNum,omitnil,omitempty" name:"InitializingShardNum"`

	// Number of unassigned shards
	UnassignedShardNum *int64 `json:"UnassignedShardNum,omitnil,omitempty" name:"UnassignedShardNum"`

	// Enterprise Edition COS storage capacity size in GB
	TotalCosStorage *int64 `json:"TotalCosStorage,omitnil,omitempty" name:"TotalCosStorage"`

	// bucket name for storing cos searchable snapshots in the Enterprise Edition cluster
	SearchableSnapshotCosBucket *string `json:"SearchableSnapshotCosBucket,omitnil,omitempty" name:"SearchableSnapshotCosBucket"`

	// App ID of the cos bucket for searchable snapshots in the Enterprise Edition cluster
	SearchableSnapshotCosAppId *string `json:"SearchableSnapshotCosAppId,omitnil,omitempty" name:"SearchableSnapshotCosAppId"`
}

type CosBackup struct {
	// <p>Whether to enable cos automatic backup</p>
	IsAutoBackup *bool `json:"IsAutoBackup,omitnil,omitempty" name:"IsAutoBackup"`

	// <p>Automatic backup execution time (accurate to the hour), e.g. "22:00"</p>
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// <p>Backup snapshot prefix</p>
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// <p>0 (Tencent Cloud repository); 1 (customer repository)</p>
	EsRepositoryType *uint64 `json:"EsRepositoryType,omitnil,omitempty" name:"EsRepositoryType"`

	// <p>Managed snapshot repository name.</p>
	PaasEsRepository *string `json:"PaasEsRepository,omitnil,omitempty" name:"PaasEsRepository"`

	// <p>Customer snapshot repository name</p>
	UserEsRepository *string `json:"UserEsRepository,omitnil,omitempty" name:"UserEsRepository"`

	// <p>cos storage folder directory</p>
	CosBasePath *string `json:"CosBasePath,omitnil,omitempty" name:"CosBasePath"`

	// <p>Snapshot storage cycle in days</p>
	StorageDuration *uint64 `json:"StorageDuration,omitnil,omitempty" name:"StorageDuration"`

	// <p>Auto backup frequency in hours.</p>
	AutoBackupInterval *uint64 `json:"AutoBackupInterval,omitnil,omitempty" name:"AutoBackupInterval"`

	// <p>Backup lock. Valid values: 0 (not locked) and 1 (locked).</p>
	CosRetention *uint64 `json:"CosRetention,omitnil,omitempty" name:"CosRetention"`

	// <p>Lock expiration date, such as 2022-12-10T08:34:48.000Z.</p>
	RetainUntilDate *string `json:"RetainUntilDate,omitnil,omitempty" name:"RetainUntilDate"`

	// <p>Lock grace period</p>
	RetentionGraceTime *uint64 `json:"RetentionGraceTime,omitnil,omitempty" name:"RetentionGraceTime"`

	// <p>Cross-region backup. Valid values: 0 (disabled) and 1 (enabled).</p>
	RemoteCos *uint64 `json:"RemoteCos,omitnil,omitempty" name:"RemoteCos"`

	// <p>Region name for cross-regional backup, such as ap-guangzhou.</p>
	RemoteCosRegion *string `json:"RemoteCosRegion,omitnil,omitempty" name:"RemoteCosRegion"`

	// <p>Policy name.</p>
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// <p>Backup index list. If left empty, it means to back up all indexes.</p>
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`

	// <p>cos multi-AZ backup. 0: single AZ; 1: multi-AZ</p>
	MultiAz *uint64 `json:"MultiAz,omitnil,omitempty" name:"MultiAz"`

	// <p>Maximum speed of write to the repository per node max_snapshot_bytes_per_sec, default 40m</p>
	MaxSnapshotPerSec *string `json:"MaxSnapshotPerSec,omitnil,omitempty" name:"MaxSnapshotPerSec"`

	// <p>Maximum speed of reading the repository per node max_restore_bytes_per_sec. Default is 40m before 710, and default unlimited from 710 onwards</p>
	MaxRestorePerSec *string `json:"MaxRestorePerSec,omitnil,omitempty" name:"MaxRestorePerSec"`

	// <p>Policy creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type CreateAutoBackUpStrategyRequestParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Policy information.
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`
}

type CreateAutoBackUpStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Policy information.
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`
}

func (r *CreateAutoBackUpStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoBackUpStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CosBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoBackUpStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoBackUpStrategyResponseParams struct {
	// Valid values: true (success) and false (failure).
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoBackUpStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoBackUpStrategyResponseParams `json:"Response"`
}

func (r *CreateAutoBackUpStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoBackUpStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexRequestParams struct {
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index type of the created index. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name of the created index
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// [Required] JSON-formatted index metadata created, such as mappings and settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index type of the created index. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name of the created index
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// [Required] JSON-formatted index metadata created, such as mappings and settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
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
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance version (supported: "5.6.4", "6.4.3", "6.8.2", "7.5.1", "7.10.1")</p>
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Access password (8-16 characters, including at least two of the following: [a-z,A-Z], [0-9], and special characters in [-!@#$%&amp;^*+=_:;,.?])</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Instance name (1-50 English letters, Chinese characters, digits, hyphens -, or underscores _)</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Abandoned. Please use NodeInfoList<br>Node quantity (2-50)</p>
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// <p>Billing type<li>PREPAID: prepaid, i.e., monthly subscription</li><li>POSTPAID_BY_HOUR: pay-as-you-go by hour</li>Default value: POSTPAID_BY_HOUR</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>Annual and monthly subscription duration (unit determined by the TimeUnit parameter)</p>
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// <p>Auto-renewal flag<li>RENEW_FLAG_AUTO: auto renewal</li><li>RENEW_FLAG_MANUAL: no auto-renewal, user renews manually</li>This needs to be set when ChargeType is PREPAID. If this parameter is not passed, no automatic renewal by default for general users, and auto renewal for SVIP users</p>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Abandoned. Please use NodeInfoList<br>Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>Abandoned. Please use NodeInfoList<br>Node disk type<li>CLOUD_SSD: ssd cloud disk</li><li>CLOUD_PREMIUM: high-performance cloud block storage</li><li>CLOUD_HSSD: enhanced SSD cloud disk</li><li>CLOUD_BSSD: universal type SSD cloud disk</li>Default value: CLOUD_SSD</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Abandoned. Please use NodeInfoList<br>Node disk capacity (unit: GB)</p>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>Billing duration unit (needs to be set when ChargeType is PREPAID. Default value: "m", which means month. Currently only support "m")</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Indicates whether to use vouchers automatically<li>0: Do not use automatically.</li><li>1: Use automatically.</li>Default value: 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Voucher ID list. Only one voucher can be specified.</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>Abandoned. Please use NodeInfoList<br>Whether to create a dedicated master node<li>true: enable a dedicated master node</li><li>false: disable a dedicated master node</li>Default value: false</p>
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// <p>Abandoned. Please use NodeInfoList<br>Number of dedicated primary nodes (supports only 3 or 5. Required if EnableDedicatedMaster is true)</p>
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// <p>Abandoned. Please use NodeInfoList<br>Dedicated primary node type (required if EnableDedicatedMaster is true)<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// <p>Abandoned. Please use NodeInfoList<br>Dedicated primary node disk size (in GB, optional; if transmitted, it must be 50. Customization is not supported temporarily)</p>
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// <p>ClusterName in the cluster configuration file (system default configuration is the instance ID, customization is not supported temporarily)</p>
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// <p>Cluster deployment method<li>0: single-AZ deployment</li><li>1: multi-AZ deployment, Beijing, Shanghai, Shanghai Finance, Guangzhou, Nanjing, Hong Kong, Singapore, Frankfurt (allowlist control)</li>Default value: 0</p>
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>Detailed information of availability zones in multi-AZ deployment (required when DeployMode is 1)</p>
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// <p>License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li>Default value: platinum</p>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <p>Node information list, for description of the specification information of various types of nodes in the cluster, such as node type, number of nodes, node specification, disk type, and disk size</p>
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// <p>Node tag information list</p>
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li></p>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// <p>Scenario template type 0: disabled 1: common 2: log 3: search</p>
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>Visualize node configuration</p>
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// <p>Create an https cluster. http is selected by default.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Maintenance time slot</p>
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// <p>Whether storage-compute separation is enabled</p>
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// <p>Additional hard disk performance</p>
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// <p>Whether intelligent check is enabled</p>
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`

	// <p>Whether to enable automatic backup</p>
	EnableCosBackup *bool `json:"EnableCosBackup,omitnil,omitempty" name:"EnableCosBackup"`

	// <p>cdcId, pass it when using a CDC subnet</p>
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// <p>Placement group affinity. Value range: [0,10]. 0 indicates disabled.</p>
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// <p>Enumeration values of sub-product ID: open source edition: "sp_es_io2", basic version: "sp_es_basic", platinum version: "sp_es_platinum", enterprise edition: "sp_es_enterprise", CDC platinum version: "sp_es_CDC_platinum", Log Enhancement Edition: "sp_es_enlogging", tsearch: "sp_tsearch_io2", logstash: "sp_es_logstash". It can be empty. When it is empty, the backend maps this field based on LicenseType.</p>
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// <p>Read-write separation mode: 0: not enabled; 1: local read/write separation; 2: remote read/write separation.</p>
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// <p>Whether the placement group has asynchronous tasks enabled</p>
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// <p>Maintenance time slot for enabling async tasks in a placement group</p>
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// <p>Auto capacity expansion parameter list</p>
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// <p>Whether to enable public network access to kibana. Enabled by default if not passed</p>
	EnableKibanaPublicAccess *string `json:"EnableKibanaPublicAccess,omitnil,omitempty" name:"EnableKibanaPublicAccess"`

	// <p>Existing cloud monitoring and alarm policy ID</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance version (supported: "5.6.4", "6.4.3", "6.8.2", "7.5.1", "7.10.1")</p>
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Access password (8-16 characters, including at least two of the following: [a-z,A-Z], [0-9], and special characters in [-!@#$%&amp;^*+=_:;,.?])</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Instance name (1-50 English letters, Chinese characters, digits, hyphens -, or underscores _)</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Abandoned. Please use NodeInfoList<br>Node quantity (2-50)</p>
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// <p>Billing type<li>PREPAID: prepaid, i.e., monthly subscription</li><li>POSTPAID_BY_HOUR: pay-as-you-go by hour</li>Default value: POSTPAID_BY_HOUR</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>Annual and monthly subscription duration (unit determined by the TimeUnit parameter)</p>
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// <p>Auto-renewal flag<li>RENEW_FLAG_AUTO: auto renewal</li><li>RENEW_FLAG_MANUAL: no auto-renewal, user renews manually</li>This needs to be set when ChargeType is PREPAID. If this parameter is not passed, no automatic renewal by default for general users, and auto renewal for SVIP users</p>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Abandoned. Please use NodeInfoList<br>Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>Abandoned. Please use NodeInfoList<br>Node disk type<li>CLOUD_SSD: ssd cloud disk</li><li>CLOUD_PREMIUM: high-performance cloud block storage</li><li>CLOUD_HSSD: enhanced SSD cloud disk</li><li>CLOUD_BSSD: universal type SSD cloud disk</li>Default value: CLOUD_SSD</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Abandoned. Please use NodeInfoList<br>Node disk capacity (unit: GB)</p>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>Billing duration unit (needs to be set when ChargeType is PREPAID. Default value: "m", which means month. Currently only support "m")</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Indicates whether to use vouchers automatically<li>0: Do not use automatically.</li><li>1: Use automatically.</li>Default value: 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Voucher ID list. Only one voucher can be specified.</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>Abandoned. Please use NodeInfoList<br>Whether to create a dedicated master node<li>true: enable a dedicated master node</li><li>false: disable a dedicated master node</li>Default value: false</p>
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// <p>Abandoned. Please use NodeInfoList<br>Number of dedicated primary nodes (supports only 3 or 5. Required if EnableDedicatedMaster is true)</p>
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// <p>Abandoned. Please use NodeInfoList<br>Dedicated primary node type (required if EnableDedicatedMaster is true)<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// <p>Abandoned. Please use NodeInfoList<br>Dedicated primary node disk size (in GB, optional; if transmitted, it must be 50. Customization is not supported temporarily)</p>
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// <p>ClusterName in the cluster configuration file (system default configuration is the instance ID, customization is not supported temporarily)</p>
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// <p>Cluster deployment method<li>0: single-AZ deployment</li><li>1: multi-AZ deployment, Beijing, Shanghai, Shanghai Finance, Guangzhou, Nanjing, Hong Kong, Singapore, Frankfurt (allowlist control)</li>Default value: 0</p>
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>Detailed information of availability zones in multi-AZ deployment (required when DeployMode is 1)</p>
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// <p>License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li>Default value: platinum</p>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <p>Node information list, for description of the specification information of various types of nodes in the cluster, such as node type, number of nodes, node specification, disk type, and disk size</p>
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// <p>Node tag information list</p>
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li></p>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// <p>Scenario template type 0: disabled 1: common 2: log 3: search</p>
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>Visualize node configuration</p>
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// <p>Create an https cluster. http is selected by default.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Maintenance time slot</p>
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// <p>Whether storage-compute separation is enabled</p>
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// <p>Additional hard disk performance</p>
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// <p>Whether intelligent check is enabled</p>
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`

	// <p>Whether to enable automatic backup</p>
	EnableCosBackup *bool `json:"EnableCosBackup,omitnil,omitempty" name:"EnableCosBackup"`

	// <p>cdcId, pass it when using a CDC subnet</p>
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// <p>Placement group affinity. Value range: [0,10]. 0 indicates disabled.</p>
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// <p>Enumeration values of sub-product ID: open source edition: "sp_es_io2", basic version: "sp_es_basic", platinum version: "sp_es_platinum", enterprise edition: "sp_es_enterprise", CDC platinum version: "sp_es_CDC_platinum", Log Enhancement Edition: "sp_es_enlogging", tsearch: "sp_tsearch_io2", logstash: "sp_es_logstash". It can be empty. When it is empty, the backend maps this field based on LicenseType.</p>
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// <p>Read-write separation mode: 0: not enabled; 1: local read/write separation; 2: remote read/write separation.</p>
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// <p>Whether the placement group has asynchronous tasks enabled</p>
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// <p>Maintenance time slot for enabling async tasks in a placement group</p>
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// <p>Auto capacity expansion parameter list</p>
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// <p>Whether to enable public network access to kibana. Enabled by default if not passed</p>
	EnableKibanaPublicAccess *string `json:"EnableKibanaPublicAccess,omitnil,omitempty" name:"EnableKibanaPublicAccess"`

	// <p>Existing cloud monitoring and alarm policy ID</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`
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
	delete(f, "EnableCosBackup")
	delete(f, "CdcId")
	delete(f, "DisasterRecoverGroupAffinity")
	delete(f, "SubProductCode")
	delete(f, "ReadWriteMode")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "AutoScaleDiskInfoList")
	delete(f, "EnableKibanaPublicAccess")
	delete(f, "AlarmPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Order ID.</p>
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DeleteAutoBackUpStrategyRequestParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Policy name.
	StrategyName []*string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`
}

type DeleteAutoBackUpStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Policy name.
	StrategyName []*string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`
}

func (r *DeleteAutoBackUpStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoBackUpStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoBackUpStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoBackUpStrategyResponseParams struct {
	// Valid values: true (success) and false (failure).
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoBackUpStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoBackUpStrategyResponseParams `json:"Response"`
}

func (r *DeleteAutoBackUpStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoBackUpStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexRequestParams struct {
	// <p>ES cluster ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Deleted index type. auto: autonomous index; normal: general index</p>
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// <p>Name of the deleted index</p>
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// <p>Cluster access username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Cluster access password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Backing index name</p>
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`

	// <p>Index life cycle</p>
	IndexPhrase *string `json:"IndexPhrase,omitnil,omitempty" name:"IndexPhrase"`
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest
	
	// <p>ES cluster ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Deleted index type. auto: autonomous index; normal: general index</p>
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// <p>Name of the deleted index</p>
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// <p>Cluster access username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Cluster access password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Backing index name</p>
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`

	// <p>Index life cycle</p>
	IndexPhrase *string `json:"IndexPhrase,omitnil,omitempty" name:"IndexPhrase"`
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
	delete(f, "IndexPhrase")
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
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Recycle bin lock settings. If true, it cannot be manually terminated and will be automatically destroyed after the preset time.</p><p>Default value: false</p>
	LockEnabled *bool `json:"LockEnabled,omitnil,omitempty" name:"LockEnabled"`

	// <p>Lock time after being put into the recycle bin</p><p>Value range: [1, 7]</p><p>Unit: day</p><p>Default value: 1</p>
	LockDuration *int64 `json:"LockDuration,omitnil,omitempty" name:"LockDuration"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Recycle bin lock settings. If true, it cannot be manually terminated and will be automatically destroyed after the preset time.</p><p>Default value: false</p>
	LockEnabled *bool `json:"LockEnabled,omitnil,omitempty" name:"LockEnabled"`

	// <p>Lock time after being put into the recycle bin</p><p>Value range: [1, 7]</p><p>Unit: day</p><p>Default value: 1</p>
	LockDuration *int64 `json:"LockDuration,omitnil,omitempty" name:"LockDuration"`
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
	delete(f, "LockEnabled")
	delete(f, "LockDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
type DescribeAutoBackUpStrategyRequestParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAutoBackUpStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoBackUpStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackUpStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoBackUpStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackUpStrategyResponseParams struct {
	// Policy information.
	CosBackupList []*CosBackup `json:"CosBackupList,omitnil,omitempty" name:"CosBackupList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoBackUpStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoBackUpStrategyResponseParams `json:"Response"`
}

func (r *DescribeAutoBackUpStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackUpStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListRequestParams struct {
	// Index type. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index name. If this parameter is left empty, all indexes are obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Paging start position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items displayed per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting fields, which support the following index names: IndexName, IndexStorage, and IndexCreateTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter by index status
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// Sort order. Supports asc and desc. Defaults to desc. Data format: "asc", "desc"
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
	
	// Index type. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index name. If this parameter is left empty, all indexes are obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Paging start position
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items displayed per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting fields, which support the following index names: IndexName, IndexStorage, and IndexCreateTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter by index status
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// Sort order. Supports asc and desc. Defaults to desc. Data format: "asc", "desc"
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
	IndexMetaFields []*IndexMetaField `json:"IndexMetaFields,omitnil,omitempty" name:"IndexMetaFields"`

	// Total query count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Index type. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name. If this parameter is left empty, all indexes are obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DescribeIndexMetaRequest struct {
	*tchttp.BaseRequest
	
	// ES cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Index type. auto: autonomous index; normal: general index
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name. If this parameter is left empty, all indexes are obtained.
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Cluster access username.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Cluster access password
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
	IndexMetaField *IndexMetaField `json:"IndexMetaField,omitnil,omitempty" name:"IndexMetaField"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Default value: 1.
	// <li>1, master log</li>
	// <li>2. Search slow logs</li>
	// <li>3. Index slow log</li>
	// <li>4. GC log</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Search term. LUCENE grammar is supported, for example, level:WARN, ip:1.1.1.1, message:test-index
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Log start time in the format of YYYY-MM-DD HH:MM:SS, for example, 2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time in the format of YYYY-MM-DD HH:MM:SS, for example, 2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination starting value. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination size. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Time sort method. Default value: 0.
	// <li>0: descending order</li>
	// <li>1, ascending</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// Node ID.
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// Slow log index name
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Slow log index shards
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// Slow log query time
	QueryCost *uint64 `json:"QueryCost,omitnil,omitempty" name:"QueryCost"`
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Default value: 1.
	// <li>1, master log</li>
	// <li>2. Search slow logs</li>
	// <li>3. Index slow log</li>
	// <li>4. GC log</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Search term. LUCENE grammar is supported, for example, level:WARN, ip:1.1.1.1, message:test-index
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Log start time in the format of YYYY-MM-DD HH:MM:SS, for example, 2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time in the format of YYYY-MM-DD HH:MM:SS, for example, 2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination starting value. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination size. Default value: 100. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Time sort method. Default value: 0.
	// <li>0: descending order</li>
	// <li>1, ascending</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// Node ID.
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// Slow log index name
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Slow log index shards
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// Slow log query time
	QueryCost *uint64 `json:"QueryCost,omitnil,omitempty" name:"QueryCost"`
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
	delete(f, "LogLevels")
	delete(f, "NodeIds")
	delete(f, "IndexName")
	delete(f, "ShardId")
	delete(f, "QueryCost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsResponseParams struct {
	// Number of returned log entries
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of log details
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
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, e.g., "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, e.g., "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging Starting Value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination size.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, e.g., "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, e.g., "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging Starting Value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination size.
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

	// Operation record
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// <p>AZ to which the cluster instance belongs. All AZs are selected by default if this parameter is not specified.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Cluster instance ID list</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Cluster instance name list.</p>
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// <p>Pagination starting value. Default value: 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Page size. Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sorting field<li>1: Instance ID</li><li>2: Instance name</li><li>3: Availability zone</li><li>4: Creation time</li>If orderByKey is not transmitted, sort by creation time in descending order</p>
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// <p>Sorting method<li>0: ascending</li><li>1: descending</li>If orderByKey is passed but orderByType is not, ascending by default</p>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>Node tag information list</p>
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>Private network vip list</p>
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// <p>IP list of cluster nodes</p>
	NodeIps []*string `json:"NodeIps,omitnil,omitempty" name:"NodeIps"`

	// <p>Availability zone list</p>
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <p>Health status filter list: 0 represents green, 1 represents yellow, 2 represents red, -1 represents unknown</p>
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// <p>Vpc list filter item</p>
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// <p>cdc cluster id</p>
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>AZ to which the cluster instance belongs. All AZs are selected by default if this parameter is not specified.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Cluster instance ID list</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Cluster instance name list.</p>
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// <p>Pagination starting value. Default value: 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Page size. Default value: 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sorting field<li>1: Instance ID</li><li>2: Instance name</li><li>3: Availability zone</li><li>4: Creation time</li>If orderByKey is not transmitted, sort by creation time in descending order</p>
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// <p>Sorting method<li>0: ascending</li><li>1: descending</li>If orderByKey is passed but orderByType is not, ascending by default</p>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>Node tag information list</p>
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>Private network vip list</p>
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// <p>IP list of cluster nodes</p>
	NodeIps []*string `json:"NodeIps,omitnil,omitempty" name:"NodeIps"`

	// <p>Availability zone list</p>
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// <p>Health status filter list: 0 represents green, 1 represents yellow, 2 represents red, -1 represents unknown</p>
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// <p>Vpc list filter item</p>
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// <p>cdc cluster id</p>
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
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
	delete(f, "NodeIps")
	delete(f, "ZoneList")
	delete(f, "HealthStatus")
	delete(f, "VpcIds")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// <p>Number of returned instances</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Instance detail list</p>
	InstanceList []*InstanceInfo `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID.
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
	// Cluster dimension view
	ClusterView *ClusterView `json:"ClusterView,omitnil,omitempty" name:"ClusterView"`

	// Node dimension view
	NodesView []*NodeView `json:"NodesView,omitnil,omitempty" name:"NodesView"`

	// Kibana dimensional view
	KibanasView []*KibanaView `json:"KibanasView,omitnil,omitempty" name:"KibanasView"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Dictionary key-value
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Dictionary name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Dictionary size. Unit: B.
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type EnableScheduleOperationDuration struct {
	// Date when enabling an asynchronous task is supported
	Periods []*string `json:"Periods,omitnil,omitempty" name:"Periods"`

	// Enable async start time
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// Supports enabling asynchronous termination time
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Enable async time zone
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type EsAcl struct {
	// kibana access blocklist
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// kibana access allowlist
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

type EsConfigSetInfo struct {
	// Configuration group type, for example, ldap and ad
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// "{\"order\":0,\"url\":\"ldap://10.0.1.72:389\",\"bind_dn\":\"cn=admin,dc=tencent,dc=com\",\"user_search.base_dn\":\"dc=tencent,dc=com\",\"user_search.filter\":\"(cn={0})\",\"group_search.base_dn\":\"dc=tencent,dc=com\"}"
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`
}

type EsDictionaryInfo struct {
	// List of enabled word dictionaries
	MainDict []*DictInfo `json:"MainDict,omitnil,omitempty" name:"MainDict"`

	// List of stopword dictionaries
	Stopwords []*DictInfo `json:"Stopwords,omitnil,omitempty" name:"Stopwords"`

	// List of QQ word segmentation dictionaries
	QQDict []*DictInfo `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// List of synonym dictionaries
	Synonym []*DictInfo `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// Update the dictionary type
	UpdateType *string `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// List of ansj enabled word dictionaries
	AnsjMain []*DictInfo `json:"AnsjMain,omitnil,omitempty" name:"AnsjMain"`

	// List of ansj stopword dictionaries
	AnsjStop []*DictInfo `json:"AnsjStop,omitnil,omitempty" name:"AnsjStop"`

	// List of ansj ambiguity lexicons
	AnsjAmbiguity []*DictInfo `json:"AnsjAmbiguity,omitnil,omitempty" name:"AnsjAmbiguity"`

	// List of ansj synonym dictionaries
	AnsjSynonyms []*DictInfo `json:"AnsjSynonyms,omitnil,omitempty" name:"AnsjSynonyms"`
}

type EsPublicAcl struct {
	// Access the blocklist
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// Access allowlist
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

// Predefined struct for user
type GetRequestTargetNodeTypesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GetRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
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
	// List of target node types that receive requests
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

type GpuInfo struct {
	// Number of Gpu blocks
	GpuCount *int64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// GPU type.
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`
}

type IndexMetaField struct {
	// Index type
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// Index name
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// JSON-formatted index metadata
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// Index status
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// Index storage size in bytes
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// Index creation time
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`

	// Backup index
	BackingIndices []*BackingIndexMetaField `json:"BackingIndices,omitnil,omitempty" name:"BackingIndices"`

	// Cluster ID of the index
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name of the index
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Version number of the cluster to which the index belongs
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// Index lifecycle field
	IndexPolicyField *IndexPolicyField `json:"IndexPolicyField,omitnil,omitempty" name:"IndexPolicyField"`

	// Index autonomous field
	IndexOptionsField *IndexOptionsField `json:"IndexOptionsField,omitnil,omitempty" name:"IndexOptionsField"`

	// Index configuration field
	IndexSettingsField *IndexSettingsField `json:"IndexSettingsField,omitnil,omitempty" name:"IndexSettingsField"`

	// Index alias field
	IndexAliasesField []*string `json:"IndexAliasesField,omitnil,omitempty" name:"IndexAliasesField"`

	// APP ID of the cluster to which the index belongs
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Number of indexed documents
	IndexDocs *uint64 `json:"IndexDocs,omitnil,omitempty" name:"IndexDocs"`
}

type IndexOptionsField struct {
	// <p>Expiration time.</p>
	ExpireMaxAge *string `json:"ExpireMaxAge,omitnil,omitempty" name:"ExpireMaxAge"`

	// <p>Expiration size</p>
	ExpireMaxSize *string `json:"ExpireMaxSize,omitnil,omitempty" name:"ExpireMaxSize"`

	// <p>Rollover period</p>
	RolloverMaxAge *string `json:"RolloverMaxAge,omitnil,omitempty" name:"RolloverMaxAge"`

	// <p>Whether to enable dynamic scrolling</p>
	RolloverDynamic *string `json:"RolloverDynamic,omitnil,omitempty" name:"RolloverDynamic"`

	// <p>Whether to enable dynamic sharding</p>
	ShardNumDynamic *string `json:"ShardNumDynamic,omitnil,omitempty" name:"ShardNumDynamic"`

	// <p>Time partition field</p>
	TimestampField *string `json:"TimestampField,omitnil,omitempty" name:"TimestampField"`

	// <p>Write mode</p>
	WriteMode *string `json:"WriteMode,omitnil,omitempty" name:"WriteMode"`

	// <p>Whether to enable complete uninstallation</p><p>Enumeration values:</p><ul><li>true: enable complete uninstallation</li><li>false: disable complete uninstallation</li></ul>
	FullOffloadedEnable *string `json:"FullOffloadedEnable,omitnil,omitempty" name:"FullOffloadedEnable"`

	// <p>Full uninstall lifecycle</p>
	FullOffloadedMaxAge *string `json:"FullOffloadedMaxAge,omitnil,omitempty" name:"FullOffloadedMaxAge"`

	// <p>Completely uninstall the lifecycle after backup index retrieval</p>
	FullOffloadedRetrieveMaxAge *string `json:"FullOffloadedRetrieveMaxAge,omitnil,omitempty" name:"FullOffloadedRetrieveMaxAge"`
}

type IndexPolicyField struct {
	// Whether the warm phase is enabled
	WarmEnable *string `json:"WarmEnable,omitnil,omitempty" name:"WarmEnable"`

	// Transfer time to the warm phase
	WarmMinAge *string `json:"WarmMinAge,omitnil,omitempty" name:"WarmMinAge"`

	// Whether the cold phase is enabled
	ColdEnable *string `json:"ColdEnable,omitnil,omitempty" name:"ColdEnable"`

	// Transfer time to the cold phase
	ColdMinAge *string `json:"ColdMinAge,omitnil,omitempty" name:"ColdMinAge"`

	// Whether the frozen phase is enabled
	FrozenEnable *string `json:"FrozenEnable,omitnil,omitempty" name:"FrozenEnable"`

	// Time of transfer to the frozen stage
	FrozenMinAge *string `json:"FrozenMinAge,omitnil,omitempty" name:"FrozenMinAge"`

	// /
	ColdAction *string `json:"ColdAction,omitnil,omitempty" name:"ColdAction"`
}

type IndexSettingsField struct {
	// Number of index primary shards
	NumberOfShards *string `json:"NumberOfShards,omitnil,omitempty" name:"NumberOfShards"`

	// Replica shard count of an index
	NumberOfReplicas *string `json:"NumberOfReplicas,omitnil,omitempty" name:"NumberOfReplicas"`

	// Index refresh frequency
	RefreshInterval *string `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`
}

type InstanceInfo struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Region.</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>uid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>User UIN.</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>UID of the VPC to which the instance belongs</p>
	VpcUid *string `json:"VpcUid,omitnil,omitempty" name:"VpcUid"`

	// <p>UID of the subnet to which the instance belongs</p>
	SubnetUid *string `json:"SubnetUid,omitnil,omitempty" name:"SubnetUid"`

	// <p>Instance status. 0: processing, 1: normal, -1: stopped, -2: being destroyed, -3: terminated, -4: isolated, 2: initializing when creating a cluster</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Auto-renewal flag. Value range: <br>RENEW_FLAG_AUTO: auto renewal<br>RENEW_FLAG_MANUAL: no auto-renewal<br>Default value: <br>RENEW_FLAG_DEFAULT: no auto-renewal<br>If this parameter is set to RENEW_FLAG_AUTO, the instance will automatically renew on a monthly basis after expiration if the account balance is sufficient.</p>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Instance billing mode. Value ranges: PREPAID: prepaid mode, i.e., monthly subscription. POSTPAID_BY_HOUR: postpaid, i.e., pay-as-you-go. CDHPAID: CDH-paid mode, i.e., only charges for CDH, without charging for instances on CDH.</p>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>Annual and monthly subscription duration, unit: month</p>
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// <p>Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>Number of nodes</p>
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// <p>Number of node CPU cores.</p>
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// <p>Node memory size (GB).</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Node disk type</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Node disk size (GB).</p>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>ES domain name</p>
	EsDomain *string `json:"EsDomain,omitnil,omitempty" name:"EsDomain"`

	// <p>ES VIP</p>
	EsVip *string `json:"EsVip,omitnil,omitempty" name:"EsVip"`

	// <p>ES port</p>
	EsPort *uint64 `json:"EsPort,omitnil,omitempty" name:"EsPort"`

	// <p>Kibana access url</p>
	KibanaUrl *string `json:"KibanaUrl,omitnil,omitempty" name:"KibanaUrl"`

	// <p>ES version number</p>
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// <p>ES configuration item</p>
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// <p>Kibana Access Control Configuration</p>
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// <p>Instance creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Last modification time of the instance</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Instance expiration time</p>
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// <p>Instance type (instance type flag, currently only 1 and 2)</p>
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Ik analyzer configuration</p>
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitnil,omitempty" name:"IkConfig"`

	// <p>Dedicated master node configuration</p>
	MasterNodeInfo *MasterNodeInfo `json:"MasterNodeInfo,omitnil,omitempty" name:"MasterNodeInfo"`

	// <p>cos automatic backup configuration</p>
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// <p>Whether cos automatic backup is allowed</p>
	AllowCosBackup *bool `json:"AllowCosBackup,omitnil,omitempty" name:"AllowCosBackup"`

	// <p>Tag list of the instance</p>
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li>Default value: platinum</p>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <p>License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li><li>enterprise: enterprise edition</li>Default value: platinum</p>
	RealLicenseType *string `json:"RealLicenseType,omitnil,omitempty" name:"RealLicenseType"`

	// <p>Whether it is a hot-cold cluster<li>true: Hot-cold cluster</li><li>false: Non-hot-cold cluster</li></p>
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitnil,omitempty" name:"EnableHotWarmMode"`

	// <p>Warm node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	WarmNodeType *string `json:"WarmNodeType,omitnil,omitempty" name:"WarmNodeType"`

	// <p>Number of warm nodes</p>
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitnil,omitempty" name:"WarmNodeNum"`

	// <p>Warm node CPU cores</p>
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitnil,omitempty" name:"WarmCpuNum"`

	// <p>Warm node memory size (GB).</p>
	WarmMemSize *uint64 `json:"WarmMemSize,omitnil,omitempty" name:"WarmMemSize"`

	// <p>Warm node disk type</p>
	WarmDiskType *string `json:"WarmDiskType,omitnil,omitempty" name:"WarmDiskType"`

	// <p>Warm node disk size (GB).</p>
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitnil,omitempty" name:"WarmDiskSize"`

	// <p>Cluster node information list</p>
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// <p>Es public network address</p>
	EsPublicUrl *string `json:"EsPublicUrl,omitnil,omitempty" name:"EsPublicUrl"`

	// <p>Multi-AZ network info</p>
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// <p>Deployment mode<li>0: single availability zone</li><li>1: multi-availability zone, Beijing, Shanghai, Shanghai Finance, Guangzhou, Nanjing, Hong Kong, Singapore, Frankfurt (allowlist control)</li></p>
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>ES public access status<li>OPEN: enable</li><li>CLOSE: disable</li></p>
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// <p>ES public network access control configuration</p>
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// <p>Kibana private network address</p>
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitnil,omitempty" name:"KibanaPrivateUrl"`

	// <p>Kibana public access status<li>OPEN: enabled</li><li>CLOSE: disabled</li></p>
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// <p>Kibana internal access status<li>OPEN: enable</li><li>CLOSE: disable</li></p>
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// <p>Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li></p>
	SecurityType *uint64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// <p>Scenario template type: 0, not enabled; 1, common scenarios; 2, log scenario; 3, search scenario</p>
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>Kibana configuration item</p>
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// <p>Kibana node information</p>
	KibanaNodeInfo *KibanaNodeInfo `json:"KibanaNodeInfo,omitnil,omitempty" name:"KibanaNodeInfo"`

	// <p>Visualize node configuration</p>
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// <p>JDK type. Valid values: oracle, kona</p>
	Jdk *string `json:"Jdk,omitnil,omitempty" name:"Jdk"`

	// <p>Cluster network communication protocol</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>Security group id</p>
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// <p>Cold node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	ColdNodeType *string `json:"ColdNodeType,omitnil,omitempty" name:"ColdNodeType"`

	// <p>Number of cold nodes</p>
	ColdNodeNum *uint64 `json:"ColdNodeNum,omitnil,omitempty" name:"ColdNodeNum"`

	// <p>Cold node CPU cores</p>
	ColdCpuNum *uint64 `json:"ColdCpuNum,omitnil,omitempty" name:"ColdCpuNum"`

	// <p>Cold node memory size (GB).</p>
	ColdMemSize *uint64 `json:"ColdMemSize,omitnil,omitempty" name:"ColdMemSize"`

	// <p>Cold node disk type</p>
	ColdDiskType *string `json:"ColdDiskType,omitnil,omitempty" name:"ColdDiskType"`

	// <p>Cold node disk size (GB).</p>
	ColdDiskSize *uint64 `json:"ColdDiskSize,omitnil,omitempty" name:"ColdDiskSize"`

	// <p>Frozen node specification<li>ES.S1.SMALL2: 1-core 2GB</li><li>ES.S1.MEDIUM4: 2-core 4GB</li><li>ES.S1.MEDIUM8: 2-core 8GB</li><li>ES.S1.LARGE16: 4-core 16GB</li><li>ES.S1.2XLARGE32: 8-core 32GB</li><li>ES.S1.4XLARGE32: 16-core 32GB</li><li>ES.S1.4XLARGE64: 16-core 64GB</li></p>
	FrozenNodeType *string `json:"FrozenNodeType,omitnil,omitempty" name:"FrozenNodeType"`

	// <p>Number of frozen nodes</p>
	FrozenNodeNum *uint64 `json:"FrozenNodeNum,omitnil,omitempty" name:"FrozenNodeNum"`

	// <p>Frozen node CPU cores</p>
	FrozenCpuNum *uint64 `json:"FrozenCpuNum,omitnil,omitempty" name:"FrozenCpuNum"`

	// <p>Frozen node memory size (GB).</p>
	FrozenMemSize *uint64 `json:"FrozenMemSize,omitnil,omitempty" name:"FrozenMemSize"`

	// <p>Frozen node disk type</p>
	FrozenDiskType *string `json:"FrozenDiskType,omitnil,omitempty" name:"FrozenDiskType"`

	// <p>Frozen node disk size (GB).</p>
	FrozenDiskSize *uint64 `json:"FrozenDiskSize,omitnil,omitempty" name:"FrozenDiskSize"`

	// <p>Cluster health status: -1 unknown; 0 Green; 1 Yellow; 2 Red</p>
	HealthStatus *int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// <p>Intracluster network url for https.</p>
	EsPrivateUrl *string `json:"EsPrivateUrl,omitnil,omitempty" name:"EsPrivateUrl"`

	// <p>https cluster private network domain name</p>
	EsPrivateDomain *string `json:"EsPrivateDomain,omitnil,omitempty" name:"EsPrivateDomain"`

	// <p>Cluster configuration group information.</p>
	EsConfigSets []*EsConfigSetInfo `json:"EsConfigSets,omitnil,omitempty" name:"EsConfigSets"`

	// <p>Cluster maintenance time slot</p>
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// <p>web node list</p>
	OptionalWebServiceInfos []*OptionalWebServiceInfo `json:"OptionalWebServiceInfos,omitnil,omitempty" name:"OptionalWebServiceInfos"`

	// <p>Autonomous index switch</p>
	AutoIndexEnabled *bool `json:"AutoIndexEnabled,omitnil,omitempty" name:"AutoIndexEnabled"`

	// <p>Whether storage and computation separation is supported.</p>
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// <p>Process progress.</p>
	ProcessPercent *float64 `json:"ProcessPercent,omitnil,omitempty" name:"ProcessPercent"`

	// <p>Kibana alerting public network alert policy<li>OPEN: enable</li><li>CLOSE: disable</li></p>
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// <p>Whether there is a kernel update this month: false (no) and true (yes).</p>
	HasKernelUpgrade *bool `json:"HasKernelUpgrade,omitnil,omitempty" name:"HasKernelUpgrade"`

	// <p>cdcId, pass it when using a CDC subnet</p>
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// <p>kibana intranet vip</p>
	KibanaPrivateVip *string `json:"KibanaPrivateVip,omitnil,omitempty" name:"KibanaPrivateVip"`

	// <p>Customize the private network url of kibana</p>
	CustomKibanaPrivateUrl *string `json:"CustomKibanaPrivateUrl,omitnil,omitempty" name:"CustomKibanaPrivateUrl"`

	// <p>Node outbound access detailed information.</p>
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// <p>Network connection solution</p>
	NetConnectScheme *string `json:"NetConnectScheme,omitnil,omitempty" name:"NetConnectScheme"`

	// <p>Placement group related parameters</p>
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// <p>Sub-product ID enumeration values: open source edition: "sp_es_io2", basic version: "sp_es_basic", platinum version: "sp_es_platinum", AI Search Enhancement Edition: "sp_es_aisearch", enterprise edition: "sp_es_enterprise", CDC platinum version: "sp_es_CDC_platinum", Log Enhancement Edition: "sp_es_enlogging", tsearch: "sp_tsearch_io2", logstash: "sp_es_logstash". It can be empty. When it is empty, the backend maps this field based on LicenseType.</p>
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// <p>Storage-compute separation cos usage, Unit: M</p>
	CosBucketStorageSize *uint64 `json:"CosBucketStorageSize,omitnil,omitempty" name:"CosBucketStorageSize"`

	// <p>Read-write separation mode: 0: not enabled; 1: local read/write separation; 2: remote read/write separation.</p>
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// <p>Whether there is an asynchronous task scheduling for placement groups</p>
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// <p>The time when an asynchronous task is scheduled</p>
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// <p>Cluster protection: OPEN: enable, CLOSE: disable</p>
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// <p>kibana intranet access address</p>
	ShowKibanaIpPort *string `json:"ShowKibanaIpPort,omitnil,omitempty" name:"ShowKibanaIpPort"`

	// <p>Whether it is a CDZLite availability zone</p>
	IsCdzLite *bool `json:"IsCdzLite,omitnil,omitempty" name:"IsCdzLite"`

	// <p>Cluster private network tcp address</p>
	EsPrivateTcpUrl *string `json:"EsPrivateTcpUrl,omitnil,omitempty" name:"EsPrivateTcpUrl"`

	// <p>Whether it is in the recycle bin</p><p>Default value: false</p>
	IsInRecycleBin *bool `json:"IsInRecycleBin,omitnil,omitempty" name:"IsInRecycleBin"`

	// <p>Is locked in the recycle bin</p><p>Default value: false</p>
	RecycleLockEnabled *bool `json:"RecycleLockEnabled,omitnil,omitempty" name:"RecycleLockEnabled"`

	// <p>The time when termination is expected</p><p>Parameter format: YYYY-MM-DD hh:mm:ss</p>
	MayDestroyPoint *string `json:"MayDestroyPoint,omitnil,omitempty" name:"MayDestroyPoint"`

	// <p>The time when termination is delayed</p><p>Unit: day</p>
	DelayDestroyInterval *int64 `json:"DelayDestroyInterval,omitnil,omitempty" name:"DelayDestroyInterval"`

	// <p>Enable automatic replacement of unhealthy nodes. 0: disable, 1: enable</p>
	EnableAutoReplace *int64 `json:"EnableAutoReplace,omitnil,omitempty" name:"EnableAutoReplace"`

	// <p>Enable mtls. 0: disable, 1: enable</p>
	OpenMTLS *uint64 `json:"OpenMTLS,omitnil,omitempty" name:"OpenMTLS"`

	// <p>Certificate type</p>
	ServerCertSource *string `json:"ServerCertSource,omitnil,omitempty" name:"ServerCertSource"`

	// <p>VIP pending migration for the cluster</p>
	OldEsVip *string `json:"OldEsVip,omitnil,omitempty" name:"OldEsVip"`

	// <p>Cluster TCP address to be migrated</p>
	OldEsPrivateTcpUrl *string `json:"OldEsPrivateTcpUrl,omitnil,omitempty" name:"OldEsPrivateTcpUrl"`

	// <p>Whether to take old nodes offline from the VIP</p>
	NeedOfflineOldNodesNotify *bool `json:"NeedOfflineOldNodesNotify,omitnil,omitempty" name:"NeedOfflineOldNodesNotify"`
}

type InstanceLog struct {
	// Log time
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Log level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Cluster node IP address.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Log content
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Cluster node ID.
	NodeID *string `json:"NodeID,omitnil,omitempty" name:"NodeID"`

	// Slow log index name
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// Slow log index shards
	Shard *string `json:"Shard,omitnil,omitempty" name:"Shard"`

	// Slow log index query time
	QueryCost *string `json:"QueryCost,omitnil,omitempty" name:"QueryCost"`
}

type KeyValue struct {
	// Key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KibanaNodeInfo struct {
	// Kibana node specification
	KibanaNodeType *string `json:"KibanaNodeType,omitnil,omitempty" name:"KibanaNodeType"`

	// Number of Kibana nodes
	KibanaNodeNum *uint64 `json:"KibanaNodeNum,omitnil,omitempty" name:"KibanaNodeNum"`

	// Number of CPUs of the Kibana node
	KibanaNodeCpuNum *uint64 `json:"KibanaNodeCpuNum,omitnil,omitempty" name:"KibanaNodeCpuNum"`

	// Kibana node memory in GB
	KibanaNodeMemSize *uint64 `json:"KibanaNodeMemSize,omitnil,omitempty" name:"KibanaNodeMemSize"`

	// Kibana node disk type
	KibanaNodeDiskType *string `json:"KibanaNodeDiskType,omitnil,omitempty" name:"KibanaNodeDiskType"`

	// Kibana node disk size
	KibanaNodeDiskSize *uint64 `json:"KibanaNodeDiskSize,omitnil,omitempty" name:"KibanaNodeDiskSize"`
}

type KibanaView struct {
	// <p>Kibana node IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Total node disk size</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>Disk utilization rate</p>
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// <p>Node memory size</p>
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Memory usage rate</p>
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// <p>Number of CPUs on the node</p>
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// <p>cpu utilization</p>
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>ts-0noqayxu-az6-hot-03222010-0</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>Client-side Ip info</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type LocalDiskInfo struct {
	// Local disk type<li>LOCAL_SATA: big data type</li><li>NVME_SSD: high IO type</li>
	LocalDiskType *string `json:"LocalDiskType,omitnil,omitempty" name:"LocalDiskType"`

	// Single local disk size
	LocalDiskSize *uint64 `json:"LocalDiskSize,omitnil,omitempty" name:"LocalDiskSize"`

	// Number of local disk blocks
	LocalDiskCount *uint64 `json:"LocalDiskCount,omitnil,omitempty" name:"LocalDiskCount"`
}

type MasterNodeInfo struct {
	// Whether the dedicated master node is enabled or not
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// Dedicated primary node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// Number of dedicated primary nodes
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// Number of CPU cores of the dedicated master node
	MasterNodeCpuNum *uint64 `json:"MasterNodeCpuNum,omitnil,omitempty" name:"MasterNodeCpuNum"`

	// Dedicated primary node memory size, in GB.
	MasterNodeMemSize *uint64 `json:"MasterNodeMemSize,omitnil,omitempty" name:"MasterNodeMemSize"`

	// Dedicated primary node disk size in GB
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Dedicated primary node disk type
	MasterNodeDiskType *string `json:"MasterNodeDiskType,omitnil,omitempty" name:"MasterNodeDiskType"`
}

// Predefined struct for user
type ModifyAutoBackUpCommonInfoRequestParams struct {

}

type ModifyAutoBackUpCommonInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyAutoBackUpCommonInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackUpCommonInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackUpCommonInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackUpCommonInfoResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoBackUpCommonInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackUpCommonInfoResponseParams `json:"Response"`
}

func (r *ModifyAutoBackUpCommonInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackUpCommonInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackUpStrategyRequestParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current policy name.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Policy information.
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`
}

type ModifyAutoBackUpStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Current policy name.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Policy information.
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`
}

func (r *ModifyAutoBackUpStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackUpStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StrategyName")
	delete(f, "CosBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackUpStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackUpStrategyResponseParams struct {
	// Valid values: true (success) and false (failure).
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoBackUpStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackUpStrategyResponseParams `json:"Response"`
}

func (r *ModifyAutoBackUpStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackUpStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// <p>Number of nodes.</p>
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// <p>Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li></p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>Node type</p><p>Enumeration values:</p><ul><li>hotData: hot data node</li><li>warmData: cold data node</li><li>dedicatedMaster: dedicated master node</li><li>dedicatedCoordinating: dedicated coordinator node</li><li>dedicatedMl: dedicated Machine Learning node</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Node disk type<li>CLOUD_SSD: ssd cloud disk</li><li>CLOUD_PREMIUM: high-performance cloud disk</li>Default value: CLOUD_SSD</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Node disk capacity (GB)</p>
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>Node local disk information.</p>
	LocalDiskInfo *LocalDiskInfo `json:"LocalDiskInfo,omitnil,omitempty" name:"LocalDiskInfo"`

	// <p>Number of node disk blocks</p>
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// <p>Whether the node disk is encrypted. 0: no encryption; 1: encryption; not encrypted by default</p>
	DiskEncrypt *uint64 `json:"DiskEncrypt,omitnil,omitempty" name:"DiskEncrypt"`

	// <p>Custom encryption key id</p>
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// <p>Custom encryption key name</p>
	KmsKeyName *string `json:"KmsKeyName,omitnil,omitempty" name:"KmsKeyName"`

	// <p>Number of CPUs</p>
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// <p>Memory size, in GB.</p>
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>Additional hard disk performance</p>
	DiskEnhance *int64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// <p>Gpu information of the node</p>
	GpuInfo *GpuInfo `json:"GpuInfo,omitnil,omitempty" name:"GpuInfo"`
}

type NodeView struct {
	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node IP
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// Whether the node is visible
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// Whether to fuse
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// Total node disk size
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Disk utilization rate
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// Node memory size (unit: GB)
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// Memory usage rate
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// Number of CPUs on the node
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// CPU utilization
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// AZ.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Node HTTP IP
	NodeHttpIp *string `json:"NodeHttpIp,omitnil,omitempty" name:"NodeHttpIp"`

	// JVM memory usage rate
	JvmMemUsage *float64 `json:"JvmMemUsage,omitnil,omitempty" name:"JvmMemUsage"`

	// Number of shards per node
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// Disk ID list on node
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// Whether it is a hidden AZ
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// Whether to act as the coordination node role
	IsCoordinationNode *bool `json:"IsCoordinationNode,omitnil,omitempty" name:"IsCoordinationNode"`

	// CVM running state
	CVMStatus *string `json:"CVMStatus,omitnil,omitempty" name:"CVMStatus"`

	// id of the placement group bound to the cvm
	CVMDisasterRecoverGroupId *string `json:"CVMDisasterRecoverGroupId,omitnil,omitempty" name:"CVMDisasterRecoverGroupId"`

	// Status of the placement group bound to the cvm. 2: Bound; 1: Binding; 0: Not enabled
	CVMDisasterRecoverGroupStatus *int64 `json:"CVMDisasterRecoverGroupStatus,omitnil,omitempty" name:"CVMDisasterRecoverGroupStatus"`
}

type Operation struct {
	// <p>Unique id of the operation</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Operation start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Operation type.</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Operation details.</p>
	Detail *OperationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// <p>Operation result.</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>Process task info</p>
	Tasks []*TaskDetail `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Operation progress.</p>
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>Rollback tag. 0: not rolled back, 1: rolling back, 2: rolled back</p>
	RollbackTag *int64 `json:"RollbackTag,omitnil,omitempty" name:"RollbackTag"`

	// <p>Operator Uin</p>
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// <p>Auto scaling flag. Valid values: 0 (non-auto) and 1 (auto).</p>
	AutoScaleTag *uint64 `json:"AutoScaleTag,omitnil,omitempty" name:"AutoScaleTag"`

	// <p>Process exception cause</p>
	SuspendedReason *string `json:"SuspendedReason,omitnil,omitempty" name:"SuspendedReason"`
}

type OperationDetail struct {
	// Original configuration message of the instance
	OldInfo []*KeyValue `json:"OldInfo,omitnil,omitempty" name:"OldInfo"`

	// Configuration message after the instance update
	NewInfo []*KeyValue `json:"NewInfo,omitnil,omitempty" name:"NewInfo"`
}

type OperationDuration struct {
	// Maintenance period, which refers to Monday to Sunday. Valid values: [0, 6].
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// Maintenance start time
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// Maintenance end time.
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Time zone in UTC format
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type OperationDurationUpdated struct {
	// Maintenance period, which refers to Monday to Sunday. Valid values: [0, 6].
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// Maintenance start time
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// Maintenance end time.
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// Time zone in UTC format
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// ES cluster ID array
	MoreInstances []*string `json:"MoreInstances,omitnil,omitempty" name:"MoreInstances"`
}

type OptionalWebServiceInfo struct {
	// Type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Public url
	PublicUrl *string `json:"PublicUrl,omitnil,omitempty" name:"PublicUrl"`

	// Private network url
	PrivateUrl *string `json:"PrivateUrl,omitnil,omitempty" name:"PrivateUrl"`

	// Public access permissions
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Private network access permission
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// Version
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Intranet vip of the web service
	PrivateVip *string `json:"PrivateVip,omitnil,omitempty" name:"PrivateVip"`

	// Custom url of cerebro on the private network
	CustomPrivateUrl *string `json:"CustomPrivateUrl,omitnil,omitempty" name:"CustomPrivateUrl"`
}

type OtherConfig struct {
	// Additional yml configuration for es
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// jvm heap configuration of es
	JvmHeapConfig *string `json:"JvmHeapConfig,omitnil,omitempty" name:"JvmHeapConfig"`
}

type OutboundPublicAcl struct {
	// Node types with permission for node outbound access
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Allowlist for outbound access from nodes
	WhiteHostList []*string `json:"WhiteHostList,omitnil,omitempty" name:"WhiteHostList"`
}

type ProcessDetail struct {
	// <p>Number of completed</p>
	Completed *int64 `json:"Completed,omitnil,omitempty" name:"Completed"`

	// <p>Remaining quantity</p>
	Remain *int64 `json:"Remain,omitnil,omitempty" name:"Remain"`

	// <p>Total quantity.</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>Task type:<br>60: Restart task<br>70: Shard migration task<br>80: Node configuration modification task</p>
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>Estimated remaining time</p><p>Unit: seconds</p>
	EstimatedTimeRemaining *int64 `json:"EstimatedTimeRemaining,omitnil,omitempty" name:"EstimatedTimeRemaining"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether to force restart<li>true: force restart</li><li>false: do not force restart</li>Default value: false</p>
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// <p>Restart mode: 0 rolling restart; 1 full restart</p>
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// <p>Select whether to upgrade the kernel patch version upon restart</p>
	UpgradeKernel *bool `json:"UpgradeKernel,omitnil,omitempty" name:"UpgradeKernel"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Whether to force restart<li>true: force restart</li><li>false: do not force restart</li>Default value: false</p>
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// <p>Restart mode: 0 rolling restart; 1 full restart</p>
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// <p>Select whether to upgrade the kernel patch version upon restart</p>
	UpgradeKernel *bool `json:"UpgradeKernel,omitnil,omitempty" name:"UpgradeKernel"`
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
	delete(f, "UpgradeKernel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// ES instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartKibanaRequest struct {
	*tchttp.BaseRequest
	
	// ES instance ID.
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node name list
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// Whether to perform forced reboot.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Optional restart mode. `in-place`: restart; `blue-green`: blue-green restart. Default value: `in-place`.
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// Node status, for use in blue-green mode; offline node blue/green at risk
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`

	// cvm delay listing time
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Whether to enable the placement group asynchronous task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Time period of the async task for the placement group
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// Event ID list.
	EventTypeIds []*string `json:"EventTypeIds,omitnil,omitempty" name:"EventTypeIds"`
}

type RestartNodesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node name list
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// Whether to perform forced reboot.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Optional restart mode. `in-place`: restart; `blue-green`: blue-green restart. Default value: `in-place`.
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// Node status, for use in blue-green mode; offline node blue/green at risk
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`

	// cvm delay listing time
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Whether to enable the placement group asynchronous task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Time period of the async task for the placement group
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// Event ID list.
	EventTypeIds []*string `json:"EventTypeIds,omitnil,omitempty" name:"EventTypeIds"`
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
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "EventTypeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Subtask status. 0: processing, 1: success, -1: failure
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Index name that failed the upgrade check
	FailedIndices []*string `json:"FailedIndices,omitnil,omitempty" name:"FailedIndices"`

	// Subtask end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Subtask level. 1: warning; 2: failure
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type TagInfo struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskDetail struct {
	// Task Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Task progress
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task completion time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Subtask
	SubTasks []*SubTaskDetail `json:"SubTasks,omitnil,omitempty" name:"SubTasks"`

	// Time taken by the task
	ElapsedTime *int64 `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`

	// Task progress details
	ProcessInfo *ProcessDetail `json:"ProcessInfo,omitnil,omitempty" name:"ProcessInfo"`
}

// Predefined struct for user
type UpdateDictionariesRequestParams struct {
	// ES instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Address of the COS main dictionary for IK word segmentation entered during installation, and the dictionary name such as test.dic entered during deletion.
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// Address of the COS stopword dictionary for IK word segmentation during installation. When deleting, enter the dictionary name, for example, test.dic.
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// Address of the synonym dictionary COS during installation, and the dictionary name such as test.dic when deleting.
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// Address of the COS QQ tokenizer dictionary during installation, and the dictionary name such as test.dic during deletion.
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// 0: installation; 1: deletion. Default value: 0
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// Whether to force restart the cluster. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpdateDictionariesRequest struct {
	*tchttp.BaseRequest
	
	// ES instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Address of the COS main dictionary for IK word segmentation entered during installation, and the dictionary name such as test.dic entered during deletion.
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// Address of the COS stopword dictionary for IK word segmentation during installation. When deleting, enter the dictionary name, for example, test.dic.
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// Address of the synonym dictionary COS during installation, and the dictionary name such as test.dic when deleting.
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// Address of the COS QQ tokenizer dictionary during installation, and the dictionary name such as test.dic during deletion.
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// 0: installation; 1: deletion. Default value: 0
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// Whether to force restart the cluster. Default value: false
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// <p>ES cluster ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Updated index type. auto: autonomous index; normal: general index</p>
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// <p>Updated index name</p>
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// <p>Updated JSON-formatted index metadata, such as mappings and settings</p>
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// <p>Cluster access username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Cluster access password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Whether to scroll the backup index</p>
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`

	// <p>Whether to retrieve the full uninstall index</p>
	MountIndex *bool `json:"MountIndex,omitnil,omitempty" name:"MountIndex"`

	// <p>Index Uuid</p>
	IndexUuid *string `json:"IndexUuid,omitnil,omitempty" name:"IndexUuid"`

	// <p>Backing index name</p>
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
}

type UpdateIndexRequest struct {
	*tchttp.BaseRequest
	
	// <p>ES cluster ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Updated index type. auto: autonomous index; normal: general index</p>
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// <p>Updated index name</p>
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// <p>Updated JSON-formatted index metadata, such as mappings and settings</p>
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// <p>Cluster access username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Cluster access password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Whether to scroll the backup index</p>
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`

	// <p>Whether to retrieve the full uninstall index</p>
	MountIndex *bool `json:"MountIndex,omitnil,omitempty" name:"MountIndex"`

	// <p>Index Uuid</p>
	IndexUuid *string `json:"IndexUuid,omitnil,omitempty" name:"IndexUuid"`

	// <p>Backing index name</p>
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
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
	delete(f, "MountIndex")
	delete(f, "IndexUuid")
	delete(f, "BackingIndexName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIndexResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Instance name (1-50 English letters, Chinese characters, digits, hyphens -, or underscores _)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Abandoned. Please use NodeInfoList.
	// Number of nodes (2-50).
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES configuration item (JSON format string)
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Password of the default user elastic (8–16 characters, including at least two of the following: [a-z,A-Z], [0-9], and special characters [-!@#$%&^*+=_:;,.?])
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Public access policy for visualization components (Kibana and Cerebro)
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// Abandoned. Please use NodeInfoList.
	// Disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Abandoned. Please use NodeInfoList.
	// Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Abandoned. Please use NodeInfoList.
	// Number of dedicated primary nodes (supports only 3 or 5)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// Abandoned. Please use NodeInfoList.
	// Dedicated primary node specification<li>ES.S1.SMALL2: 1-core 2GB</li><li>ES.S1.MEDIUM4: 2-core 4GB</li><li>ES.S1.MEDIUM8: 2-core 8GB</li><li>ES.S1.LARGE16: 4-core 16GB</li><li>ES.S1.2XLARGE32: 8-core 32GB</li><li>ES.S1.4XLARGE32: 16-core 32GB</li><li>ES.S1.4XLARGE64: 16-core 64GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// Abandoned. Please use NodeInfoList.
	// Dedicated primary node disk size (in GB. The system default configuration is 50 GB. Customization is temporarily not supported.)
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Whether to force restart when updating the configuration <li>true: force restart</li><li>false: do not force restart</li>Currently, this needs to be set only when updating EsConfig. The default value is false.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// COS auto backup information
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Node information list. You can transmit only the node to be updated and its corresponding specification information. Supported operations include <li>modifying the count of a node type</li><li>modifying the node specification and disk size of a node type</li><li>adding a node type (you need to simultaneously specify the node type, count, specification, disk, and other information)</li>. You can perform only one of the above operations at a time, and the disk type cannot be modified.
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// Public access status of the ES cluster
	// OPEN: enabled.
	// -CLOSE: disabled.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Public network access control list
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana public access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana internal access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// Turn on or off user authentication for the basic version of ES 6.8 and above
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana private network port
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: Scale out via blue-green change without restarting the cluster (default) 1: Scale out by unmounting disks with a cluster rolling restart
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// Multi-AZ deployment
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// Scenario template type. -1: not enabled. 1: common. 2: log. 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana configuration item (JSON format string)
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Visualize node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Switch to the new network architecture
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// Enable Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// Cerebro public access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// Cerebro private network access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// Newly-added or modified configuration group information
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// Maintenance time slot.
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether Alerting public network alert output is enabled:
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// Custom domain name of kibana on the private network
	KibanaPrivateDomain *string `json:"KibanaPrivateDomain,omitnil,omitempty" name:"KibanaPrivateDomain"`

	// Custom domain name of cerebro on the private network
	CerebroPrivateDomain *string `json:"CerebroPrivateDomain,omitnil,omitempty" name:"CerebroPrivateDomain"`

	// Change to an https cluster. http is selected by default.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Outbound access details of the node
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// Outbound access operations of the node
	// OPEN
	// -CLOSE: disabled.
	OutboundPublicAccess *string `json:"OutboundPublicAccess,omitnil,omitempty" name:"OutboundPublicAccess"`

	// Parameter for delayed cvm launch
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Read-write separation mode: -1: not enabled; 1: local read/write separation; 2: remote read/write separation.
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// Whether to enable the placement group async task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Placement group async task maintenance time slot
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// Cluster protection: OPEN: enable, CLOSE: disable
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// Auto capacity expansion parameter
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// Auto-expansion delete parameter
	AutoScaleDiskDeleteNodeTypeList []*string `json:"AutoScaleDiskDeleteNodeTypeList,omitnil,omitempty" name:"AutoScaleDiskDeleteNodeTypeList"`

	// Other additional configuration, jvm or yml.
	OtherConfig *OtherConfig `json:"OtherConfig,omitnil,omitempty" name:"OtherConfig"`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name (1-50 English letters, Chinese characters, digits, hyphens -, or underscores _)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Abandoned. Please use NodeInfoList.
	// Number of nodes (2-50).
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES configuration item (JSON format string)
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Password of the default user elastic (8–16 characters, including at least two of the following: [a-z,A-Z], [0-9], and special characters [-!@#$%&^*+=_:;,.?])
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Public access policy for visualization components (Kibana and Cerebro)
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// Abandoned. Please use NodeInfoList.
	// Disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Abandoned. Please use NodeInfoList.
	// Node specification<li>ES.S1.SMALL2: 1-core 2G</li><li>ES.S1.MEDIUM4: 2-core 4G</li><li>ES.S1.MEDIUM8: 2-core 8G</li><li>ES.S1.LARGE16: 4-core 16G</li><li>ES.S1.2XLARGE32: 8-core 32G</li><li>ES.S1.4XLARGE32: 16-core 32G</li><li>ES.S1.4XLARGE64: 16-core 64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Abandoned. Please use NodeInfoList.
	// Number of dedicated primary nodes (supports only 3 or 5)
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// Abandoned. Please use NodeInfoList.
	// Dedicated primary node specification<li>ES.S1.SMALL2: 1-core 2GB</li><li>ES.S1.MEDIUM4: 2-core 4GB</li><li>ES.S1.MEDIUM8: 2-core 8GB</li><li>ES.S1.LARGE16: 4-core 16GB</li><li>ES.S1.2XLARGE32: 8-core 32GB</li><li>ES.S1.4XLARGE32: 16-core 32GB</li><li>ES.S1.4XLARGE64: 16-core 64GB</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// Abandoned. Please use NodeInfoList.
	// Dedicated primary node disk size (in GB. The system default configuration is 50 GB. Customization is temporarily not supported.)
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// Whether to force restart when updating the configuration <li>true: force restart</li><li>false: do not force restart</li>Currently, this needs to be set only when updating EsConfig. The default value is false.
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// COS auto backup information
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// Node information list. You can transmit only the node to be updated and its corresponding specification information. Supported operations include <li>modifying the count of a node type</li><li>modifying the node specification and disk size of a node type</li><li>adding a node type (you need to simultaneously specify the node type, count, specification, disk, and other information)</li>. You can perform only one of the above operations at a time, and the disk type cannot be modified.
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// Public access status of the ES cluster
	// OPEN: enabled.
	// -CLOSE: disabled.
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// Public network access control list
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana public access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana internal access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// Turn on or off user authentication for the basic version of ES 6.8 and above
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana private network port
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: Scale out via blue-green change without restarting the cluster (default) 1: Scale out by unmounting disks with a cluster rolling restart
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// Multi-AZ deployment
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// Scenario template type. -1: not enabled. 1: common. 2: log. 3: search
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana configuration item (JSON format string)
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Visualize node configuration
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// Switch to the new network architecture
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// Enable Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// Cerebro public access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// Cerebro private network access status
	// OPEN: enabled.
	// -CLOSE: disabled.
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// Newly-added or modified configuration group information
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// Maintenance time slot.
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// Whether Alerting public network alert output is enabled:
	// OPEN: enabled.
	// -CLOSE: disabled.
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// Custom domain name of kibana on the private network
	KibanaPrivateDomain *string `json:"KibanaPrivateDomain,omitnil,omitempty" name:"KibanaPrivateDomain"`

	// Custom domain name of cerebro on the private network
	CerebroPrivateDomain *string `json:"CerebroPrivateDomain,omitnil,omitempty" name:"CerebroPrivateDomain"`

	// Change to an https cluster. http is selected by default.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Outbound access details of the node
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// Outbound access operations of the node
	// OPEN
	// -CLOSE: disabled.
	OutboundPublicAccess *string `json:"OutboundPublicAccess,omitnil,omitempty" name:"OutboundPublicAccess"`

	// Parameter for delayed cvm launch
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Read-write separation mode: -1: not enabled; 1: local read/write separation; 2: remote read/write separation.
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// Whether to enable the placement group async task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Placement group async task maintenance time slot
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// Cluster protection: OPEN: enable, CLOSE: disable
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// Auto capacity expansion parameter
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// Auto-expansion delete parameter
	AutoScaleDiskDeleteNodeTypeList []*string `json:"AutoScaleDiskDeleteNodeTypeList,omitnil,omitempty" name:"AutoScaleDiskDeleteNodeTypeList"`

	// Other additional configuration, jvm or yml.
	OtherConfig *OtherConfig `json:"OtherConfig,omitnil,omitempty" name:"OtherConfig"`
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
	delete(f, "KibanaPrivateDomain")
	delete(f, "CerebroPrivateDomain")
	delete(f, "Protocol")
	delete(f, "OutboundPublicAcls")
	delete(f, "OutboundPublicAccess")
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "ReadWriteMode")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "EnableDestroyProtection")
	delete(f, "AutoScaleDiskInfoList")
	delete(f, "AutoScaleDiskDeleteNodeTypeList")
	delete(f, "OtherConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInstanceResponseParams struct {
	// Order ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Name list of plug-ins that needs to be installed
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// List of plug-in names to uninstall
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// Whether to force restart. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Whether to reinstall. Default value: false
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0: System plugin
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

type UpdatePluginsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name list of plug-ins that needs to be installed
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// List of plug-in names to uninstall
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// Whether to force restart. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// Whether to reinstall. Default value: false
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0: System plugin
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of target node types that receive requests
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`
}

type UpdateRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of target node types that receive requests
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
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// Target ES version. Valid values: "6.4.3", "6.8.2", "7.5.1", "7.10.1", and "7.14.2".
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// Whether to perform an upgrade check only. Default value: false.
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// Target business feature version: <li>oss open source edition</li><li>basic edition</li>Currently used only when upgrading from 5.6.4 to 6.x. Default value: basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Upgrade method: <li>scale blue-green change</li><li>restart rolling restart</li>Default value: scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Whether to back up the cluster before upgrading the version. By default, no backup is performed.
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// In rolling mode, whether to skip the check and perform a forced reboot. Default value: false
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`

	// Parameter for delayed cvm launch
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Whether to enable the placement group asynchronous task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Time period of the async task for the placement group
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target ES version. Valid values: "6.4.3", "6.8.2", "7.5.1", "7.10.1", and "7.14.2".
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// Whether to perform an upgrade check only. Default value: false.
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// Target business feature version: <li>oss open source edition</li><li>basic edition</li>Currently used only when upgrading from 5.6.4 to 6.x. Default value: basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Upgrade method: <li>scale blue-green change</li><li>restart rolling restart</li>Default value: scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Whether to back up the cluster before upgrading the version. By default, no backup is performed.
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// In rolling mode, whether to skip the check and perform a forced reboot. Default value: false
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`

	// Parameter for delayed cvm launch
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// Number of concurrent shard migrations
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// Shard migration concurrency speed
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// Whether to enable the placement group asynchronous task
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// Time period of the async task for the placement group
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`
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
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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

	// License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li>Default value: platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Indicates whether to use vouchers automatically<li>0: Do not use automatically.</li><li>1: Use automatically.</li>Default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Only one voucher can be specified.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li><li>If no parameter is passed, the original status is maintained by default. When passing a parameter, note that it can only change from not enabled to enabled, and cannot change from enabled to not enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Whether to force restart <li>true: force restart</li><li>false: do not force restart</li> Default value: false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// License type<li>oss: open source edition</li><li>basic: basic version</li><li>platinum: platinum version</li>Default value: platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// Indicates whether to use vouchers automatically<li>0: Do not use automatically.</li><li>1: Use automatically.</li>Default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Only one voucher can be specified.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Whether xpack security authentication is enabled for basic version 6.8 (and above versions)<li>1: not enabled</li><li>2: enabled</li><li>If no parameter is passed, the original status is maintained by default. When passing a parameter, note that it can only change from not enabled to enabled, and cannot change from enabled to not enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Whether to force restart <li>true: force restart</li><li>false: do not force restart</li> Default value: false
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
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
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
	// Number of nodes to visualize, fixed as 1
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// Visualize node specifications
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type ZoneDetail struct {
	// AZ.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether it is a hidden AZ
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`
}