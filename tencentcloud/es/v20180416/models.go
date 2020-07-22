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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CosBackup struct {

	// Whether to enable auto-backup to COS
	IsAutoBackup *bool `json:"IsAutoBackup,omitempty" name:"IsAutoBackup"`

	// Auto-backup time (accurate down to the hour), such as "22:00"
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// Availability Zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance version ("5.6.4", "6.4.3", "6.8.2", or "7.5.1")
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Access password, which must contain 8 to 16 characters, and include at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitempty" name:"Password"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2–50)
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Billing mode <li>POSTPAID_BY_HOUR: Pay-as-you-go hourly </li>Default value: POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// This parameter is not used on the global website
	ChargePeriod *uint64 `json:"ChargePeriod,omitempty" name:"ChargePeriod"`

	// This parameter is not used on the global website
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node storage type <li>CLOUD_SSD: SSD cloud storage </li><li>CLOUD_PREMIUM: premium cloud storage </li>Default value: CLOUD_SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// This parameter is not used on the global website
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// This parameter has been disused. Please use `NodeInfoList`
	// Whether to create a dedicated main node <li>true: yes </li><li>false: no </li>Default value: false
	EnableDedicatedMain *bool `json:"EnableDedicatedMain,omitempty" name:"EnableDedicatedMain"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated main nodes (only 3 and 5 are supported. This value must be passed in if `EnableDedicatedMain` is `true`)
	MainNodeNum *uint64 `json:"MainNodeNum,omitempty" name:"MainNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated main node type, which must be passed in if `EnableDedicatedMain` is `true` <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MainNodeType *string `json:"MainNodeType,omitempty" name:"MainNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated main node disk size in GB, which is optional. If passed in, it can only be 50 and cannot be customized currently
	MainNodeDiskSize *uint64 `json:"MainNodeDiskSize,omitempty" name:"MainNodeDiskSize"`

	// ClusterName in the cluster configuration file, which is the instance ID by default and currently cannot be customized
	ClusterNameInConf *string `json:"ClusterNameInConf,omitempty" name:"ClusterNameInConf"`

	// Cluster deployment mode <li>0: single-AZ deployment </li><li>1: multi-AZ deployment </li>Default value: 0
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// Details of AZs in multi-AZ deployment mode (which is required when DeployMode is 1)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitempty" name:"MultiZoneInfo" list`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`
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

		// Instance ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

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

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest

	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Log type. Default value: 1
	// <li>1: main log</li>
	// <li>2: search slow log</li>
	// <li>3: index slow log</li>
	// <li>4: GC log</li>
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// Search keyword, which supports LUCENE syntax, such as `level:WARN`, `ip:1.1.1.1`, and `message:test-index`
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Log start time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Log end time in the format of YYYY-MM-DD HH:MM:SS, such as 2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 100. Maximum value: 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Time sorting order. Default value: 0
	// <li>0: descending</li>
	// <li>1: ascending</li>
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned logs
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Log details list
		InstanceLogList []*InstanceLog `json:"InstanceLogList,omitempty" name:"InstanceLogList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest

	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time, such as "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time, such as "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination start value
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of operation records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Operation history
		Operations []*Operation `json:"Operations,omitempty" name:"Operations" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// AZ of the cluster instance. If this is not passed in, all AZs are used by default
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// List of cluster instance IDs
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// List of cluster instance names
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sort by field <li>1: instance ID </li><li>2: instance name </li><li>3: AZ </li><li>4: creation time </li>If `orderKey` is not passed in, sort by creation time in descending order
	OrderByKey *uint64 `json:"OrderByKey,omitempty" name:"OrderByKey"`

	// Sorting order <li>0: ascending </li><li>1: descending </li>If orderByKey is passed in but orderByType is not, ascending order is used by default
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`

	// Node tag information list
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// VPC VIP list
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
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

		// Number of returned instances
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance details
		InstanceList []*InstanceInfo `json:"InstanceList,omitempty" name:"InstanceList" list`

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

type DictInfo struct {

	// Dictionary key value
	Key *string `json:"Key,omitempty" name:"Key"`

	// Dictionary name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Dictionary size in B
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type EsAcl struct {

	// Kibana access blacklist
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

	// Kibana access whitelist
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`
}

type EsDictionaryInfo struct {

	// List of non-stop words
	MainDict []*DictInfo `json:"MainDict,omitempty" name:"MainDict" list`

	// List of stop words
	Stopwords []*DictInfo `json:"Stopwords,omitempty" name:"Stopwords" list`
}

type EsPublicAcl struct {

	// Access blacklist
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

	// Access whitelist
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`
}

type InstanceInfo struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Availability Zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// User ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// UID of the VPC where the instance resides
	VpcUid *string `json:"VpcUid,omitempty" name:"VpcUid"`

	// UID of the subnet where the instance resides
	SubnetUid *string `json:"SubnetUid,omitempty" name:"SubnetUid"`

	// Instance status. 0: processing; 1: normal; -1: stopped; -2: terminating; -3: terminated
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Instance billing method. Valid values: POSTPAID_BY_HOUR (pay-as-you-go hourly); CDHPAID (billed based on CDH, i.e., only CDH is billed but not the instances on CDH)
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// This parameter is not used on the global website
	ChargePeriod *uint64 `json:"ChargePeriod,omitempty" name:"ChargePeriod"`

	// This parameter is not used on the global website
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// Number of nodes
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Number of CPU cores of the node
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// Node memory size in GB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Node disk type
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ES domain name
	EsDomain *string `json:"EsDomain,omitempty" name:"EsDomain"`

	// ES VIP
	EsVip *string `json:"EsVip,omitempty" name:"EsVip"`

	// ES port
	EsPort *uint64 `json:"EsPort,omitempty" name:"EsPort"`

	// Kibana access URL
	KibanaUrl *string `json:"KibanaUrl,omitempty" name:"KibanaUrl"`

	// ES version number
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// ES configuration item
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// Kibana access control configuration
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last modified time of the instance
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// This parameter is not used on the global website
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// Instance type (instance type identifier, which can be only 1 or 2 currently)
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// IK analyzer configuration
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitempty" name:"IkConfig"`

	// Dedicated main node configuration
	MainNodeInfo *MainNodeInfo `json:"MainNodeInfo,omitempty" name:"MainNodeInfo"`

	// Auto-backup to COS configuration
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// Whether to allow auto-backup to COS
	AllowCosBackup *bool `json:"AllowCosBackup,omitempty" name:"AllowCosBackup"`

	// List of tags owned by the instance
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// Whether it is a hot/warm cluster <li>true: yes </li><li>false: no</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitempty" name:"EnableHotWarmMode"`

	// Warm node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmNodeType *string `json:"WarmNodeType,omitempty" name:"WarmNodeType"`

	// Number of warm nodes
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitempty" name:"WarmNodeNum"`

	// Number of warm node CPU cores
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitempty" name:"WarmCpuNum"`

	// Warm node memory size in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmMemSize *uint64 `json:"WarmMemSize,omitempty" name:"WarmMemSize"`

	// Warm node disk type
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmDiskType *string `json:"WarmDiskType,omitempty" name:"WarmDiskType"`

	// Warm node disk size in GB
	// Note: This field may return null, indicating that no valid values can be obtained.
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitempty" name:"WarmDiskSize"`

	// Cluster node information list
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// ES public IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	EsPublicUrl *string `json:"EsPublicUrl,omitempty" name:"EsPublicUrl"`

	// Multi-AZ network information
	// Note: This field may return null, indicating that no valid values can be obtained.
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitempty" name:"MultiZoneInfo" list`

	// Deployment mode <li>0: single-AZ </li><li>1: multi-AZ</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// ES public access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicAccess *string `json:"PublicAccess,omitempty" name:"PublicAccess"`

	// ES public access control configuration
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitempty" name:"EsPublicAcl"`

	// Kibana private IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitempty" name:"KibanaPrivateUrl"`

	// Kibana public access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Kibana private access status <li>OPEN: enabled </li><li>CLOSE: disabled
	// Note: This field may return null, indicating that no valid values can be obtained.
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityType *uint64 `json:"SecurityType,omitempty" name:"SecurityType"`
}

type InstanceLog struct {

	// Log time
	Time *string `json:"Time,omitempty" name:"Time"`

	// Log level
	Level *string `json:"Level,omitempty" name:"Level"`

	// Cluster node IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Log content
	Message *string `json:"Message,omitempty" name:"Message"`
}

type KeyValue struct {

	// Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LocalDiskInfo struct {

	// Local disk type <li>LOCAL_SATA: big data </li><li>NVME_SSD: high IO</li>
	LocalDiskType *string `json:"LocalDiskType,omitempty" name:"LocalDiskType"`

	// Size of a single local disk
	LocalDiskSize *uint64 `json:"LocalDiskSize,omitempty" name:"LocalDiskSize"`

	// Number of local disks
	LocalDiskCount *uint64 `json:"LocalDiskCount,omitempty" name:"LocalDiskCount"`
}

type MainNodeInfo struct {

	// Whether to enable the dedicated main node
	EnableDedicatedMain *bool `json:"EnableDedicatedMain,omitempty" name:"EnableDedicatedMain"`

	// Dedicated main node specification <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MainNodeType *string `json:"MainNodeType,omitempty" name:"MainNodeType"`

	// Number of dedicated main nodes
	MainNodeNum *uint64 `json:"MainNodeNum,omitempty" name:"MainNodeNum"`

	// Number of CPU cores of the dedicated main node
	MainNodeCpuNum *uint64 `json:"MainNodeCpuNum,omitempty" name:"MainNodeCpuNum"`

	// Memory size of the dedicated main node in GB
	MainNodeMemSize *uint64 `json:"MainNodeMemSize,omitempty" name:"MainNodeMemSize"`

	// Disk size of the dedicated main node in GB
	MainNodeDiskSize *uint64 `json:"MainNodeDiskSize,omitempty" name:"MainNodeDiskSize"`

	// Disk type of the dedicated main node
	MainNodeDiskType *string `json:"MainNodeDiskType,omitempty" name:"MainNodeDiskType"`
}

type NodeInfo struct {

	// Number of nodes
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// Node type <li>hotData: hot data node</li>
	// <li>warmData: warm data node</li>
	// <li>dedicatedMain: dedicated main node</li>
	// Default value: hotData
	Type *string `json:"Type,omitempty" name:"Type"`

	// Node disk type <li>CLOUD_SSD: SSD cloud storage </li><li>CLOUD_PREMIUM: Premium cloud disk </li>Default value: CLOUD_SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Node disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Local disk information
	// Note: this field may return null, indicating that no valid values can be obtained.
	LocalDiskInfo *LocalDiskInfo `json:"LocalDiskInfo,omitempty" name:"LocalDiskInfo"`

	// Number of node disks
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`
}

type Operation struct {

	// Unique operation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Operation start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Operation type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Operation details
	Detail *OperationDetail `json:"Detail,omitempty" name:"Detail"`

	// Operation result
	Result *string `json:"Result,omitempty" name:"Result"`

	// Workflow task information
	Tasks []*TaskDetail `json:"Tasks,omitempty" name:"Tasks" list`

	// Operation progress
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
}

type OperationDetail struct {

	// Original instance configuration information
	OldInfo []*KeyValue `json:"OldInfo,omitempty" name:"OldInfo" list`

	// Updated instance configuration information
	NewInfo []*KeyValue `json:"NewInfo,omitempty" name:"NewInfo" list`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether to force restart <li>true: Yes </li><li>false: No </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubTaskDetail struct {

	// Subtask name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Subtask result
	Result *bool `json:"Result,omitempty" name:"Result"`

	// Subtask error message
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// Subtask type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Subtask status. 0: processing, 1: succeeded, -1: failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Name of the index for which the check for upgrade failed
	FailedIndices []*string `json:"FailedIndices,omitempty" name:"FailedIndices" list`

	// Subtask end time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Subtask level. 1: warning, 2: failed
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TaskDetail struct {

	// Task name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Task progress
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// Task completion time
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// Subtask
	SubTasks []*SubTaskDetail `json:"SubTasks,omitempty" name:"SubTasks" list`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name, which can contain 1 to 50 English letters, Chinese characters, digits, dashes (-), or underscores (_)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of nodes (2–50)
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// Configuration item (JSON string). Only the following items are supported currently: <li>action.destructive_requires_name</li><li>indices.fielddata.cache.size</li><li>indices.query.bool.max_clause_count</li>
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// Password of the default user “elastic“, which must contain 8 to 16 characters, including at least two of the following three types of characters: [a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?]
	Password *string `json:"Password,omitempty" name:"Password"`

	// Access control list
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Disk size in GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Node specification <li>ES.S1.SMALL2: 1-core 2 GB </li><li>ES.S1.MEDIUM4: 2-core 4 GB </li><li>ES.S1.MEDIUM8: 2-core 8 GB </li><li>ES.S1.LARGE16: 4-core 16 GB </li><li>ES.S1.2XLARGE32: 8-core 32 GB </li><li>ES.S1.4XLARGE32: 16-core 32 GB </li><li>ES.S1.4XLARGE64: 16-core 64 GB </li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Number of dedicated main nodes (only 3 and 5 are supported)
	MainNodeNum *uint64 `json:"MainNodeNum,omitempty" name:"MainNodeNum"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated main node specification <li>ES.S1.SMALL2: 1-core 2 GB</li><li>ES.S1.MEDIUM4: 2-core 4 GB</li><li>ES.S1.MEDIUM8: 2-core 8 GB</li><li>ES.S1.LARGE16: 4-core 16 GB</li><li>ES.S1.2XLARGE32: 8-core 32 GB</li><li>ES.S1.4XLARGE32: 16-core 32 GB</li><li>ES.S1.4XLARGE64: 16-core 64 GB</li>
	MainNodeType *string `json:"MainNodeType,omitempty" name:"MainNodeType"`

	// This parameter has been disused. Please use `NodeInfoList`
	// Dedicated main node disk size in GB. This is 50 GB by default and currently cannot be customized
	MainNodeDiskSize *uint64 `json:"MainNodeDiskSize,omitempty" name:"MainNodeDiskSize"`

	// Whether to force restart during configuration update <li>true: Yes </li><li>false: No </li>This needs to be set only for EsConfig. Default value: false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`

	// Auto-backup to COS
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// Node information list. You can pass in only the nodes to be updated and their corresponding specification information. Supported operations include: <li>modifying the number of nodes in the same type </li><li>modifying the specification and disk size of nodes in the same type </li><li>adding a node type (you must also specify the node type, quantity, specification, disk, etc.) </li>The above operations can only be performed one at a time, and the disk type cannot be modified
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// Public network access status
	PublicAccess *string `json:"PublicAccess,omitempty" name:"PublicAccess"`

	// Public network ACL
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitempty" name:"EsPublicAcl"`

	// Public network access status of Kibana
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Private network access status of Kibana
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`

	// Enables or disables user authentication for ES Basic Edition v6.8 and above
	BasicSecurityType *int64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`

	// Kibana private port
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitempty" name:"KibanaPrivatePort"`
}

func (r *UpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Target ES version. Valid values: 6.4.3, 6.8.2, 7.5.1
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// Whether to check for upgrade only. Default value: false
	CheckOnly *bool `json:"CheckOnly,omitempty" name:"CheckOnly"`

	// Target X-Pack edition: <li>OSS: Open-source Edition </li><li>basic: Basic Edition </li>Currently only used for v5.6.4 to v6.x upgrade. Default value: basic
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// License type <li>oss: Open Source Edition </li><li>basic: Basic Edition </li><li>platinum: Platinum Edition </li>Default value: Platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// Whether to automatically use vouchers <li>0: No </li><li>1: Yes </li>Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// List of voucher IDs (only one voucher can be specified at a time currently)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// Whether to enable X-Pack security authentication in Basic Edition 6.8 (and above) <li>1: disabled </li><li>2: enabled</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`

	// Whether to force restart <li>true: yes </li><li>false: no </li>Default value: false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *UpgradeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneDetail struct {

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
