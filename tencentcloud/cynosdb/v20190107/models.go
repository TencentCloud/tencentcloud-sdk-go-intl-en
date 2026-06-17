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

package v20190107

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AIOptimizerStatus struct {
	// <p>Status. closing-Closing, closed-Closed, opening-Opening, training-Training, trained-Trained, train_failed-Training failed.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Start time.</p>
	OpenTime *string `json:"OpenTime,omitnil,omitempty" name:"OpenTime"`

	// <p>Training progress</p>
	TrainingProgress *int64 `json:"TrainingProgress,omitnil,omitempty" name:"TrainingProgress"`
}

type AIOptimizerTaskData struct {
	// <p>Template ID</p>
	TemplateID *string `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

type Ability struct {
	// Whether secondary AZ is supported
	IsSupportSlaveZone *string `json:"IsSupportSlaveZone,omitnil,omitempty" name:"IsSupportSlaveZone"`

	// The causes for no support from an availability zone.
	NonsupportSlaveZoneReason *string `json:"NonsupportSlaveZoneReason,omitnil,omitempty" name:"NonsupportSlaveZoneReason"`

	// Whether read-only instance is supported
	IsSupportRo *string `json:"IsSupportRo,omitnil,omitempty" name:"IsSupportRo"`

	// Reasons why RO instances are not supported.
	NonsupportRoReason *string `json:"NonsupportRoReason,omitnil,omitempty" name:"NonsupportRoReason"`

	// Whether manual snapshot backup initiation is supported.
	IsSupportManualSnapshot *string `json:"IsSupportManualSnapshot,omitnil,omitempty" name:"IsSupportManualSnapshot"`

	// Whether transparent data encryption is supported.
	IsSupportTransparentDataEncryption *string `json:"IsSupportTransparentDataEncryption,omitnil,omitempty" name:"IsSupportTransparentDataEncryption"`

	// Reasons for no support of transparent data encryption.
	NoSupportTransparentDataEncryptionReason *string `json:"NoSupportTransparentDataEncryptionReason,omitnil,omitempty" name:"NoSupportTransparentDataEncryptionReason"`

	// Whether manual initiation of logical backup is supported.
	IsSupportManualLogic *string `json:"IsSupportManualLogic,omitnil,omitempty" name:"IsSupportManualLogic"`

	// Enable global encryption.
	IsSupportGlobalEncryption *string `json:"IsSupportGlobalEncryption,omitnil,omitempty" name:"IsSupportGlobalEncryption"`

	// The causes for unsupported global encryption.
	NoSupportGlobalEncryptionReason *string `json:"NoSupportGlobalEncryptionReason,omitnil,omitempty" name:"NoSupportGlobalEncryptionReason"`

	// Status code for unsupported tde reason.
	NoSupportTransparentDataEncryptionReasonCode *string `json:"NoSupportTransparentDataEncryptionReasonCode,omitnil,omitempty" name:"NoSupportTransparentDataEncryptionReasonCode"`

	// Status code for unsupported global encryption.
	NoSupportGlobalEncryptionReasonCode *string `json:"NoSupportGlobalEncryptionReasonCode,omitnil,omitempty" name:"NoSupportGlobalEncryptionReasonCode"`
}

type Account struct {
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The max connections
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// Whether password rotation is enabled (0: turn off; 1: turn on)
	PasswordRotation *int64 `json:"PasswordRotation,omitnil,omitempty" name:"PasswordRotation"`
}

// Predefined struct for user
type ActivateInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type ActivateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of instance IDs in the format of `cynosdbmysql-ins-n7ocdslw` as displayed in the TDSQL-C for MySQL console. You can use the instance list querying API to query the ID, i.e., the `InstanceId` value in the output parameters.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

func (r *ActivateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateInstanceResponseParams struct {
	// Task flow ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateInstanceResponseParams `json:"Response"`
}

func (r *ActivateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateLibraDBClusterRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ActivateLibraDBClusterRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *ActivateLibraDBClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateLibraDBClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateLibraDBClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateLibraDBClusterResponseParams struct {
	// flow id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateLibraDBClusterResponse struct {
	*tchttp.BaseResponse
	Response *ActivateLibraDBClusterResponseParams `json:"Response"`
}

func (r *ActivateLibraDBClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateLibraDBClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateLibraDBInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type ActivateLibraDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

func (r *ActivateLibraDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateLibraDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateLibraDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateLibraDBInstanceResponseParams struct {
	// task flow id
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateLibraDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateLibraDBInstanceResponseParams `json:"Response"`
}

func (r *ActivateLibraDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateLibraDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Binlog sync mode. Default value: async. Optional values: sync, semisync, async.
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// Semi-sync timeout period in milliseconds. To ensure business stability, semi-sync replication has a degradation logic. When the primary availability zone cluster waits for the secondary availability zone cluster to confirm a transaction and exceeds the timeout period, the replication method will degrade to asynchronous replication. The minimum value is set to 1000 ms, support up to 4294967295 ms, and defaults to 10000 ms.
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type AddClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Binlog sync mode. Default value: async. Optional values: sync, semisync, async.
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// Semi-sync timeout period in milliseconds. To ensure business stability, semi-sync replication has a degradation logic. When the primary availability zone cluster waits for the secondary availability zone cluster to confirm a transaction and exceeds the timeout period, the replication method will degrade to asynchronous replication. The minimum value is set to 1000 ms, support up to 4294967295 ms, and defaults to 10000 ms.
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

func (r *AddClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	delete(f, "BinlogSyncWay")
	delete(f, "SemiSyncTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterSlaveZoneResponseParams struct {
	// Async FlowId.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *AddClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cpu cores</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// <p>Instance Machine Type. Supported values are as follows:</p><ul><li>common: indicates universal type</li><li>exclusive: indicates exclusive</li></ul>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Instance group ID, used when adding new instances to an existing RO group. If not passed, a new RO group will be created. The current version does not recommend transmitting this value.</p>
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// <p>ID of the associated VPC network.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID. If VpcId is set up, SubnetId is required.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Database type, value ranges from...to...: </p><li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Order source, string length range [0,64)</p>
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// <p>Transaction mode. 0: place order and pay; 1: place order</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Parameter template ID</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Parameter list. InstanceParams is valid only when ParamTemplateId is passed in.</p>
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// <p>Security group ID. You can specify security groups when creating a read-only instance.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>proxy sync upgrade</p>
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cpu cores</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of added read-only instances. Value range: (0,16].
	ReadOnlyCount *int64 `json:"ReadOnlyCount,omitnil,omitempty" name:"ReadOnlyCount"`

	// <p>Instance Machine Type. Supported values are as follows:</p><ul><li>common: indicates universal type</li><li>exclusive: indicates exclusive</li></ul>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Instance group ID, used when adding new instances to an existing RO group. If not passed, a new RO group will be created. The current version does not recommend transmitting this value.</p>
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// <p>ID of the associated VPC network.</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID. If VpcId is set up, SubnetId is required.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The port used when adding an RO group. Value range: [0,65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance name. String length range: [0,64).
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Database type, value ranges from...to...: </p><li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Order source, string length range [0,64)</p>
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// <p>Transaction mode. 0: place order and pay; 1: place order</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Parameter template ID</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Parameter list. InstanceParams is valid only when ParamTemplateId is passed in.</p>
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// <p>Security group ID. You can specify security groups when creating a read-only instance.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>proxy sync upgrade</p>
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "ReadOnlyCount")
	delete(f, "DeviceType")
	delete(f, "InstanceGrpId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "InstanceName")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "OrderSource")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "InstanceParams")
	delete(f, "SecurityGroupIds")
	delete(f, "UpgradeProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesResponseParams struct {
	// <p>Freeze transaction. One frozen flow is activated at a time.</p>
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// <p>Order ID for payment. Description: Recommend you use the <a href="https://www.tencentcloud.com/document/product/1098/40735">DescribeResourcesByDealName</a> api to query order associated instance.</p>
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// <p>Delivery resource id list.</p>
	//
	// Deprecated: ResourceIds is deprecated.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>Large order number</p>
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddInstancesResponseParams `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLibraDBInstancesRequestParams struct {
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Disk size.
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Sync object list
	Objects *Objects `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Port used when adding new RO groups, value range [0,65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Number of newly-added read-only instances, value range (0,15]
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance name, string length range [0,64), value range uppercase and lowercase letters, digits 0-9, '_', '-', '.'
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of replicas
	ReplicasNum *int64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// Value is 'Exclusive' when ReplicasNum>1 or ReplicasNum=1 and Cpu>=32 cores, and 'Common' in other scenarios.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Disk type.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Order source, string length range [0,64)
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Transaction mode 0-Place order and pay 1-Place order
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// ID of the associated VPC network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If VpcId is set up, SubnetId is required.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Security group ID. You can specify security groups when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Analysis engine version
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// Purchase period, combined with TimeUnit for the changes to take effect
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, takes effect when combined with TimeSpan. Options: day:d, month:m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Source instance id
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`
}

type AddLibraDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory in GB
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Disk size.
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Sync object list
	Objects *Objects `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Port used when adding new RO groups, value range [0,65535)
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Number of newly-added read-only instances, value range (0,15]
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance name, string length range [0,64), value range uppercase and lowercase letters, digits 0-9, '_', '-', '.'
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Number of replicas
	ReplicasNum *int64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// Value is 'Exclusive' when ReplicasNum>1 or ReplicasNum=1 and Cpu>=32 cores, and 'Common' in other scenarios.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Disk type.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Order source, string length range [0,64)
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Transaction mode 0-Place order and pay 1-Place order
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// ID of the associated VPC network.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If VpcId is set up, SubnetId is required.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Security group ID. You can specify security groups when creating a read-only instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Analysis engine version
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// Purchase period, combined with TimeUnit for the changes to take effect
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, takes effect when combined with TimeSpan. Options: day:d, month:m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Source instance id
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`
}

func (r *AddLibraDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLibraDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "StorageSize")
	delete(f, "PayMode")
	delete(f, "Objects")
	delete(f, "Port")
	delete(f, "GoodsNum")
	delete(f, "InstanceName")
	delete(f, "ReplicasNum")
	delete(f, "InstanceType")
	delete(f, "StorageType")
	delete(f, "AutoVoucher")
	delete(f, "OrderSource")
	delete(f, "DealMode")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SecurityGroupIds")
	delete(f, "LibraDBVersion")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "SrcInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLibraDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLibraDBInstancesResponseParams struct {
	// Large order number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Freeze transaction. One frozen flow is activated at a time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Post-paid order number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Delivery resource id list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddLibraDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddLibraDBInstancesResponseParams `json:"Response"`
}

func (r *AddLibraDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLibraDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddServerlessRoInstancesRequestParams struct {
	// <p>Cluster Id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Minimum specification of the ro instance</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Maximum specification of ro instance</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>ro instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID. If VpcId is set up, SubnetId is required.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Port used when adding new RO groups, value range [0,65535)</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Security group ID. You can specify security groups when creating a read-only instance.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Whether to automatically pause</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>Auto-pause time</p><p>Unit: s</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>Instance parameter</p>
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// <p>Parameter template</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Number of newly-added read-only instances</p>
	RoCount *int64 `json:"RoCount,omitnil,omitempty" name:"RoCount"`
}

type AddServerlessRoInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster Id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Minimum specification of the ro instance</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Maximum specification of ro instance</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>ro instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID. If VpcId is set up, SubnetId is required.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Port used when adding new RO groups, value range [0,65535)</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Security group ID. You can specify security groups when creating a read-only instance.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Whether to automatically pause</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>Auto-pause time</p><p>Unit: s</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>Instance parameter</p>
	InstanceParams []*ModifyParamItem `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// <p>Parameter template</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Number of newly-added read-only instances</p>
	RoCount *int64 `json:"RoCount,omitnil,omitempty" name:"RoCount"`
}

func (r *AddServerlessRoInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddServerlessRoInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "InstanceParams")
	delete(f, "ParamTemplateId")
	delete(f, "RoCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddServerlessRoInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddServerlessRoInstancesResponseParams struct {
	// <p>Freeze transaction. One frozen flow is activated at a time.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// <p>Post-paid order number</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// <p>Delivery resource id list.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>Large order number</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddServerlessRoInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddServerlessRoInstancesResponseParams `json:"Response"`
}

func (r *AddServerlessRoInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddServerlessRoInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Addr struct {
	// IP address
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Array of instance group IDs, starting with the prefix cynosdbmysql-grp- or a cluster ID.
	// Description: To get the instance group ID of a cluster, perform [query cluster instance group](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Security group ID list to be modified. An array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance group IDs, starting with the prefix cynosdbmysql-grp- or a cluster ID.
	// Description: To get the instance group ID of a cluster, perform [query cluster instance group](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Security group ID list to be modified. An array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInstanceFilters struct {
	// Filter condition values. Supported values include: InstanceId - Instance ID, InstanceName - Instance name, ProjectId - Project ID, TagKey - Tag key, Tag - Tag (separated by vertical bar, e.g., TagKey|TagValue).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// True indicates exact match, false indicates fuzzy match.
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// Filter value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuditInstanceInfo struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Tag information.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type AuditRuleFilters struct {
	// A single audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// Rule template ID
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Filter of the rule template
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time of a rule template
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Rule template modification time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// Alarm level. valid values: 1 (low risk), 2 (medium risk), 3 (high risk).
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. 0 - no alert, 1 - alert.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Template status. 0 - no task, 1 - modifying.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Template application is used in which instances.
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`
}

type AutoCopyConfig struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe region
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`

	// Delivery Type: binlog, redolog, snapshot, logic
	CopyType *string `json:"CopyType,omitnil,omitempty" name:"CopyType"`
}

type AutoMapRule struct {
	// Source instance Id
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Source database regular
	SrcDatabaseRegex *string `json:"SrcDatabaseRegex,omitnil,omitempty" name:"SrcDatabaseRegex"`

	// Source table regular
	SrcTableRegex *string `json:"SrcTableRegex,omitnil,omitempty" name:"SrcTableRegex"`

	// Target database regular
	DstDatabaseRegex *string `json:"DstDatabaseRegex,omitnil,omitempty" name:"DstDatabaseRegex"`

	// Target table regular
	DstTableRegex *string `json:"DstTableRegex,omitnil,omitempty" name:"DstTableRegex"`
}

type BackupConfigInfo struct {
	// System automation time.
	BackupCustomAutoTime *bool `json:"BackupCustomAutoTime,omitnil,omitempty" name:"BackupCustomAutoTime"`

	// Indicates the full backup start time. value range: [0-24*3600]. for example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// Indicates the full backup end time. value range: [0-24*3600]. for example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// Currently this parameter cannot be modified. no need to specify. backup frequency is an array of length 7, corresponding to the backup method from sunday to saturday, full for full backup and increment for incremental backup.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupWeekDays []*string `json:"BackupWeekDays,omitnil,omitempty" name:"BackupWeekDays"`

	// Interval.
	BackupIntervalTime *int64 `json:"BackupIntervalTime,omitnil,omitempty" name:"BackupIntervalTime"`

	// Indicates the backup retention period in seconds. data will be cleaned up longer than this time. 7 days means 3600247=604800. the maximum is 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// Enable cross-region backup.
	// Enable.
	// 0: disabled.
	CrossRegionsEnable *string `json:"CrossRegionsEnable,omitnil,omitempty" name:"CrossRegionsEnable"`

	// Cross-Regional backup region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`

	// Automatic data backup trigger policy, periodically: automatic periodic backup, frequent: high-frequency backup
	BackupTriggerStrategy *string `json:"BackupTriggerStrategy,omitnil,omitempty" name:"BackupTriggerStrategy"`

	// Backup delivery relationship
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

type BackupFileInfo struct {
	// <p>Snapshot file ID, abandoned, please use BackupId</p>
	SnapshotId *uint64 `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// <p>Backup file name</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Backup file size</p>
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>Backup start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Backup completion time.</p>
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// <p>Backup type: snapshot, snapshot backup; logic, logical backup</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Backup method: auto, automatic backup; manual, manual backup</p>
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// <p>Backup file status: success: backup successful; fail: backup failed; creating: backup file being created; deleting: backup file deleting</p>
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// <p>Backup file time</p>
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// <p>Backup ID</p>
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Snapshot type. Value range: full, full; increment, incremental</p>
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// <p>Backup file remark</p>
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// <p>Delivery status</p>
	CopyStatus *string `json:"CopyStatus,omitnil,omitempty" name:"CopyStatus"`

	// <p>Key id</p>
	EncryptKeyId *string `json:"EncryptKeyId,omitnil,omitempty" name:"EncryptKeyId"`

	// <p>Key region</p>
	EncryptRegion *string `json:"EncryptRegion,omitnil,omitempty" name:"EncryptRegion"`

	// <p>Safe info</p>
	VaultInfos []*VaultInfo `json:"VaultInfos,omitnil,omitempty" name:"VaultInfos"`

	// <p>Backup cycle policy</p>
	BackupPeriodStrategy *string `json:"BackupPeriodStrategy,omitnil,omitempty" name:"BackupPeriodStrategy"`
}

type BackupLimitClusterRestriction struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Download limit configuration.
	BackupLimitRestriction *BackupLimitRestriction `json:"BackupLimitRestriction,omitnil,omitempty" name:"BackupLimitRestriction"`
}

type BackupLimitRestriction struct {
	// Restriction type.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// This parameter only supports In, which indicates that the vpc specified by LimitVpc can be downloaded. the default is In.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Specified ips can download; specified ips are not allowed to download.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// Specifies the vpc setting for download restrictions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// Specifies the ip settings for limiting downloads.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type BackupLimitVpcItem struct {
	// Specifies the region for limiting download sources. currently only supports the current region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Limit the vpc list for downloads.
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BackupRegionAndIds struct {
	// Backup region.
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// Backup ID.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type BillingResourceInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Order ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type BindClusterResourcePackagesRequestParams struct {
	// The unique ID of a resource pack
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *BindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageIds")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindClusterResourcePackagesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *BindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *BindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindInstanceInfo struct {
	// ID of the bound cluster.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The region where the bound instance is located.
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// Type of the bound instance.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// The instance ID under the bound cluster.
	ExtendIds []*string `json:"ExtendIds,omitnil,omitempty" name:"ExtendIds"`
}

type BinlogConfigInfo struct {
	// Specifies the retention time of binlogs.
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// Whether binlog cross-region backup is enabled.
	BinlogCrossRegionsEnable *string `json:"BinlogCrossRegionsEnable,omitnil,omitempty" name:"BinlogCrossRegionsEnable"`

	// binlog in a different region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BinlogCrossRegions []*string `json:"BinlogCrossRegions,omitnil,omitempty" name:"BinlogCrossRegions"`

	// Safe info
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

type BinlogItem struct {
	// Binlog filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// binlog cross-region coverage
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`

	// Backup delivery status
	CopyStatus *string `json:"CopyStatus,omitnil,omitempty" name:"CopyStatus"`

	// Safe info
	VaultInfos []*VaultInfo `json:"VaultInfos,omitnil,omitempty" name:"VaultInfos"`

	// Encryption key
	EncryptKeyId *string `json:"EncryptKeyId,omitnil,omitempty" name:"EncryptKeyId"`

	// Encrypt key region
	EncryptRegion *string `json:"EncryptRegion,omitnil,omitempty" name:"EncryptRegion"`
}

type BizTaskInfo struct {
	// <p>Task ID.</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>User appid</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Task creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Delayed execution time</p>
	DelayTime *string `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// <p>Task failure information</p>
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// <p>Asynchronous task flow id</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Task input information</p>
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// <p>Instance group id.</p>
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// <p>Instance group id.</p>
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// <p>Instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Task operation object id</p>
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// <p>Task operation object type.</p>
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// <p>Operator uin</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>Task output information</p>
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>Task status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Task type</p>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>Parent task ID that triggers this task</p>
	TriggerTaskId *int64 `json:"TriggerTaskId,omitnil,omitempty" name:"TriggerTaskId"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Task start time</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Task end time</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Task progress</p>
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// <p>Modify parameter task information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ModifyParamsData is deprecated.
	ModifyParamsData []*ModifyParamsData `json:"ModifyParamsData,omitnil,omitempty" name:"ModifyParamsData"`

	// <p>Create cluster task information</p>
	CreateClustersData *CreateClustersData `json:"CreateClustersData,omitnil,omitempty" name:"CreateClustersData"`

	// <p>Cluster rollback task information</p>
	RollbackData *RollbackData `json:"RollbackData,omitnil,omitempty" name:"RollbackData"`

	// <p>Instance configuration change task information</p>
	ModifyInstanceData *ModifyInstanceData `json:"ModifyInstanceData,omitnil,omitempty" name:"ModifyInstanceData"`

	// <p>Manual backup task information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManualBackupData *ManualBackupData `json:"ManualBackupData,omitnil,omitempty" name:"ManualBackupData"`

	// <p>Modify kernel version task information</p>
	ModifyDbVersionData *ModifyDbVersionData `json:"ModifyDbVersionData,omitnil,omitempty" name:"ModifyDbVersionData"`

	// <p>Cluster Availability Zone Information</p>
	ClusterSlaveData *ClusterSlaveData `json:"ClusterSlaveData,omitnil,omitempty" name:"ClusterSlaveData"`

	// <p>Convert cluster logs</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SwitchClusterLogBin *SwitchClusterLogBin `json:"SwitchClusterLogBin,omitnil,omitempty" name:"SwitchClusterLogBin"`

	// <p>Modify instance parameter data</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModifyInstanceParamsData *BizTaskModifyParamsData `json:"ModifyInstanceParamsData,omitnil,omitempty" name:"ModifyInstanceParamsData"`

	// <p>Maintenance time</p>
	TaskMaintainInfo *TaskMaintainInfo `json:"TaskMaintainInfo,omitnil,omitempty" name:"TaskMaintainInfo"`

	// <p>Instance Log Delivery Information</p>
	InstanceCLSDeliveryInfos []*InstanceCLSDeliveryInfo `json:"InstanceCLSDeliveryInfos,omitnil,omitempty" name:"InstanceCLSDeliveryInfos"`

	// <p>Task progress information</p>
	TaskProgressInfo *TaskProgressInfo `json:"TaskProgressInfo,omitnil,omitempty" name:"TaskProgressInfo"`

	// <p>Global database network task</p>
	GdnTaskInfo *GdnTaskInfo `json:"GdnTaskInfo,omitnil,omitempty" name:"GdnTaskInfo"`

	// <p>Safe id</p>
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// <p>Safe name</p>
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// <p>AI optimizer task information</p>
	AIOptimizerTaskData *AIOptimizerTaskData `json:"AIOptimizerTaskData,omitnil,omitempty" name:"AIOptimizerTaskData"`
}

type BizTaskModifyInstanceParam struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Details of instance parameter modification task.
	ModifyInstanceParamList []*ModifyParamItem `json:"ModifyInstanceParamList,omitnil,omitempty" name:"ModifyInstanceParamList"`
}

type BizTaskModifyParamsData struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster parameter modification data.
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// Instance parameter modification data.
	ModifyInstanceParams []*BizTaskModifyInstanceParam `json:"ModifyInstanceParams,omitnil,omitempty" name:"ModifyInstanceParams"`
}

type CLSInfo struct {
	// Log topic operation: Options are create, reuse. 
	// create: Create a new log topic, using TopicName to create the log topic.
	// reuse: Use an existing log topic, using TopicId to specify the log topic.
	// The combination of using an existing log topic while creating a new logset is not allowed.
	TopicOperation *string `json:"TopicOperation,omitnil,omitempty" name:"TopicOperation"`

	// Logset operation: Options are create, reuse.
	// create: Create a new logset, using GroupName to create the logset.
	// reuse: Use an existing logset, using GroupId to specify the logset.
	// The combination of using an existing log topic while creating a new logset is not allowed.
	GroupOperation *string `json:"GroupOperation,omitnil,omitempty" name:"GroupOperation"`

	// Log delivery region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Log topic ID.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Logset ID.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Logset name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

// Predefined struct for user
type CalculateBackupSaveSecExpiresRequestParams struct {
	// Backup safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup retention period (seconds). Must be greater than 0.
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Number of items per page, range (0,100], default 10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: endTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type CalculateBackupSaveSecExpiresRequest struct {
	*tchttp.BaseRequest
	
	// Backup safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup retention period (seconds). Must be greater than 0.
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Number of items per page, range (0,100], default 10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: endTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *CalculateBackupSaveSecExpiresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CalculateBackupSaveSecExpiresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "BackupSaveSeconds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CalculateBackupSaveSecExpiresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CalculateBackupSaveSecExpiresResponseParams struct {
	// Total number of backup files that will be deleted
	WillDeleteBackupFileCount *int64 `json:"WillDeleteBackupFileCount,omitnil,omitempty" name:"WillDeleteBackupFileCount"`

	// Backup file list to be deleted
	WillDeleteBackupFiles []*WillDeleteItem `json:"WillDeleteBackupFiles,omitnil,omitempty" name:"WillDeleteBackupFiles"`

	// Total files of binlogs that will be deleted
	WillDeleteBinlogFileCount *int64 `json:"WillDeleteBinlogFileCount,omitnil,omitempty" name:"WillDeleteBinlogFileCount"`

	// Binlog file list to be deleted
	WillDeleteBinlogFiles []*WillDeleteItem `json:"WillDeleteBinlogFiles,omitnil,omitempty" name:"WillDeleteBinlogFiles"`

	// Total Redolog files to be deleted
	WillDeleteRedoLogFileCount *int64 `json:"WillDeleteRedoLogFileCount,omitnil,omitempty" name:"WillDeleteRedoLogFileCount"`

	// List of Redolog files that will be deleted
	WillDeleteRedoLogFiles []*WillDeleteItem `json:"WillDeleteRedoLogFiles,omitnil,omitempty" name:"WillDeleteRedoLogFiles"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CalculateBackupSaveSecExpiresResponse struct {
	*tchttp.BaseResponse
	Response *CalculateBackupSaveSecExpiresResponseParams `json:"Response"`
}

func (r *CalculateBackupSaveSecExpiresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CalculateBackupSaveSecExpiresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCreateLibraDBInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CheckCreateLibraDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CheckCreateLibraDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCreateLibraDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCreateLibraDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCreateLibraDBInstanceResponseParams struct {
	// Overall verification status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Check item
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckItem []*CheckItem `json:"CheckItem,omitnil,omitempty" name:"CheckItem"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckCreateLibraDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CheckCreateLibraDBInstanceResponseParams `json:"Response"`
}

func (r *CheckCreateLibraDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCreateLibraDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckItem struct {
	// Check item name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Item *string `json:"Item,omitnil,omitempty" name:"Item"`

	// Verification result of this item
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Detailed explanation of validation failure and modification suggestions
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Details of validation failure and modification suggestions
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpectedValue *string `json:"ExpectedValue,omitnil,omitempty" name:"ExpectedValue"`
}

// Predefined struct for user
type CheckTransferClusterZoneRequestParams struct {
	// Source Cluster Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Target availability zone
	DstZone *string `json:"DstZone,omitnil,omitempty" name:"DstZone"`

	// Target zone info for proxy migration
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type CheckTransferClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// Source Cluster Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Target availability zone
	DstZone *string `json:"DstZone,omitnil,omitempty" name:"DstZone"`

	// Target zone info for proxy migration
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

func (r *CheckTransferClusterZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransferClusterZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DstZone")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTransferClusterZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTransferClusterZoneResponseParams struct {
	// Whether the check is successful
	CheckStatus *bool `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// Reason for check failure
	CheckMsg *string `json:"CheckMsg,omitnil,omitempty" name:"CheckMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckTransferClusterZoneResponse struct {
	*tchttp.BaseResponse
	Response *CheckTransferClusterZoneResponseParams `json:"Response"`
}

func (r *CheckTransferClusterZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransferClusterZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseClusterPasswordComplexityRequestParams struct {
	// Cluster IDs in array
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type CloseClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster IDs in array
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *CloseClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CloseClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CloseClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyEndPointRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type CloseProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *CloseProxyEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyEndPointResponseParams struct {
	// Async process ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Asynchronous Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseProxyEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CloseProxyEndPointResponseParams `json:"Response"`
}

func (r *CloseProxyEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

func (r *CloseProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "OnlyCloseRW")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseProxyResponse struct {
	*tchttp.BaseResponse
	Response *CloseProxyResponseParams `json:"Response"`
}

func (r *CloseProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// Process ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanRequestParams struct {
	// Instance group ID
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

type CloseWanRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

func (r *CloseWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseWanResponse struct {
	*tchttp.BaseResponse
	Response *CloseWanResponseParams `json:"Response"`
}

func (r *CloseWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstanceDetail struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Engine type</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Instance status</p>
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>Instance status description</p>
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// <p>cpu cores</p>
	InstanceCpu *int64 `json:"InstanceCpu,omitnil,omitempty" name:"InstanceCpu"`

	// <p>Memory</p>
	InstanceMemory *int64 `json:"InstanceMemory,omitnil,omitempty" name:"InstanceMemory"`

	// <p>hard disk</p>
	InstanceStorage *int64 `json:"InstanceStorage,omitnil,omitempty" name:"InstanceStorage"`

	// <p>Instance role</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Execution start time (seconds from 00:00).</p>
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// <p>Duration (unit: s)</p>
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// <p>The time when it can be executed, enumeration value: ["Mon","Tue","Wed","Thu","Fri", "Sat", "Sun"]</p>
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// <p>serverless instance substatus</p>
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// <p>Instance task information</p>
	InstanceTasks []*ObjectTask `json:"InstanceTasks,omitnil,omitempty" name:"InstanceTasks"`

	// <p>Instance Machine Type</p><ol><li>common, universal type.</li><li>exclusive, dedicated.</li></ol>
	InstanceDeviceType *string `json:"InstanceDeviceType,omitnil,omitempty" name:"InstanceDeviceType"`

	// <p>Instance storage type<br>Description: This parameter returns a value only when the resource belonging to the query is LibraDB.</p>
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`

	// <p>Database type</p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Node list<br>Description: This parameter returns a value only when querying resources belonging to LibraDB.</p>
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// <p>AI optimizer status</p>
	AIOptimizerStatus *AIOptimizerStatus `json:"AIOptimizerStatus,omitnil,omitempty" name:"AIOptimizerStatus"`
}

type ClusterReadOnlyValue struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Specifies the value of the read-only switch.
	ReadOnlyValue *string `json:"ReadOnlyValue,omitnil,omitempty" name:"ReadOnlyValue"`
}

type ClusterSlaveData struct {
	// Old primary availability zone.
	OldMasterZone *string `json:"OldMasterZone,omitnil,omitempty" name:"OldMasterZone"`

	// Old slave availability zone.
	OldSlaveZone []*string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// New primary availability zone.
	NewMasterZone *string `json:"NewMasterZone,omitnil,omitempty" name:"NewMasterZone"`

	// New slave availability zone.
	NewSlaveZone []*string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// New from availability zone attribute.
	NewSlaveZoneAttr []*SlaveZoneAttrItem `json:"NewSlaveZoneAttr,omitnil,omitempty" name:"NewSlaveZoneAttr"`

	// Old availability zone attributes.
	OldSlaveZoneAttr []*SlaveZoneAttrItem `json:"OldSlaveZoneAttr,omitnil,omitempty" name:"OldSlaveZoneAttr"`
}

type ClusterTaskId struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CopyBackupToVaultRequestParams struct {
	// <p>Target safe ID. The backup file will be copied to this safe.</p>
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// <p>List of backup file IDs. Supports batch operations to copy multiple backup files.</p>
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type CopyBackupToVaultRequest struct {
	*tchttp.BaseRequest
	
	// <p>Target safe ID. The backup file will be copied to this safe.</p>
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// <p>List of backup file IDs. Supports batch operations to copy multiple backup files.</p>
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

func (r *CopyBackupToVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyBackupToVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyBackupToVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyBackupToVaultResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyBackupToVaultResponse struct {
	*tchttp.BaseResponse
	Response *CopyBackupToVaultResponseParams `json:"Response"`
}

func (r *CopyBackupToVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyBackupToVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyClusterPasswordComplexityRequestParams struct {
	// A parameter used to replicate the array of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster ID
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
}

type CopyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// A parameter used to replicate the array of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster ID
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`
}

func (r *CopyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "SourceClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *CopyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *CopyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of new accounts
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountsResponseParams `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateRequestParams struct {
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateAuditRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateResponseParams struct {
	// The generated rule template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditRuleTemplateResponseParams `json:"Response"`
}

func (r *CreateAuditRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 	Message sent to safe
	Vaults []*CreateBackupVaultItem `json:"Vaults,omitnil,omitempty" name:"Vaults"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup type. Valid values: `logic` (logic backup), `snapshot` (physical backup)
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup database, which is valid when `BackupType` is `logic`.
	BackupDatabases []*string `json:"BackupDatabases,omitnil,omitempty" name:"BackupDatabases"`

	// Backup table, which is valid when `BackupType` is `logic`.
	BackupTables []*DatabaseTables `json:"BackupTables,omitnil,omitempty" name:"BackupTables"`

	// Backup name
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 	Message sent to safe
	Vaults []*CreateBackupVaultItem `json:"Vaults,omitnil,omitempty" name:"Vaults"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupType")
	delete(f, "BackupDatabases")
	delete(f, "BackupTables")
	delete(f, "BackupName")
	delete(f, "Vaults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupVaultItem struct {
	// Safe id
	// Note: This field may return null, indicating that no valid values can be obtained.
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe region
	// Note: This field may return null, indicating that no valid values can be obtained.
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`
}

// Predefined struct for user
type CreateCLSDeliveryRequestParams struct {
	// Intance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log shipping configuration.
	CLSInfoList []*CLSInfo `json:"CLSInfoList,omitnil,omitempty" name:"CLSInfoList"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type CreateCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Intance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log shipping configuration.
	CLSInfoList []*CLSInfo `json:"CLSInfoList,omitnil,omitempty" name:"CLSInfoList"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *CreateCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSInfoList")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSDeliveryResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSDeliveryResponseParams `json:"Response"`
}

func (r *CreateCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterDatabaseRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Host permissions of the authorized user
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Host permissions of the authorized user
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "CharacterSet")
	delete(f, "CollateRule")
	delete(f, "UserHostPrivileges")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterDatabaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterDatabaseResponseParams `json:"Response"`
}

func (r *CreateClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClustersData struct {
	// Instance CPU.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Cluster storage upper limit.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`
}

// Predefined struct for user
type CreateClustersRequestParams struct {
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Database type</p><p>Enumeration value:</p><ul><li>MYSQL: MYSQL</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Database version</p><p>Enumeration value:</p><ul><li>5.7: MySQL 5.7 edition</li><li>8.0: MySQL 8.0 edition</li></ul>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>project-ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Required when DbMode is NORMAL or left blank<br>Cpu cores of the regular instance</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Required when DbMode is NORMAL or left blank<br>Ordinary instance memory in GB</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance count</p><p>Value ranges from [1, 16]</p><p>Default value: 2</p><ul><li>A value of 2 means one rw instance + one ro instance.</li><li>The transmitted n refers to one rw instance + n-1 ro instances (identical specifications).</li><li>For a more precise cluster composition collocation, please use InstanceInitInfos.</li><li>The value set by this parameter is suitable for provisioned resource cluster. If needed to set the specifications and quantity of Serverless cluster, please use the InstanceInitInfo structure in InstanceInitInfos.n.</li></ul>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>This parameter is meaningless and abandoned.<br>Storage size, in GB.</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Cluster name, length less than 64 characters, each character value ranges from uppercase/lowercase letters, digits, special symbols ('-', '_', '.')</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Account (8-64 characters, a combination of uppercase and lowercase letters, digits and symbols ~!@#$%^&amp;*_-+=|(){}[]:;&#39;&lt;&gt;,.?/ with any three types required)</p>
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// <p>Port, default 3306, in the range of [0, 65535)</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Billing mode</p><p>Enumeration value:</p><ul><li>0: Pay-as-you-go billing</li><li>1: Monthly Subscription</li></ul><p>Default value: 0</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Number of clusters to purchase. Optional value range: [1,50]. Default is 1.</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>Rollback type</p><p>Enumeration value:</p><ul><li>noneRollback: No rollback</li><li>snapRollback: Snapshot rollback</li><li>timeRollback: Point-in-time rollback</li></ul>
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// <p>Snapshot rollback means snapshotId; point-in-time rollback means queryId. A value of 0 indicates requirement to judge whether the time point is valid.</p>
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// <p>During rollback, input the source cluster ID to search for the source poolId</p>
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// <p>Point-in-time rollback, specified time; snapshot rollback, snapshot time</p>
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// <p>This parameter is meaningless and abandoned.<br>Point-in-time rollback, allowed range of specified time.</p>
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// <p>Ordinary instance storage cap, in GB<br>When DbType is MYSQL and the storage billing mode is prepaid, this parameter should not exceed the maximum storage specification corresponding to cpu and memory.</p>
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// <p>Annual and monthly subscription duration</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>Annual and monthly subscription duration unit, ['s', 'd', 'm', 'y']</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Whether Annual/Monthly Subscription is auto-renewed</p><p>Enumeration value:</p><ul><li>0: Default renewal method</li><li>1: Auto-renewal</li><li>2: No auto-renewal</li></ul><p>Default value: 0</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Whether to automatically select voucher 1 Yes 0 No Default value: 0</p><p>Enumeration value:</p><ul><li>1: Yes</li><li>0: No</li></ul><p>Default value: 0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Instance count (this parameter is deprecated and only for existing accommodative)</p>
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// <p>Order source</p>
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// <p>tag Array info that should be bound for cluster creation</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Db type</p><p>Enumeration value:</p><ul><li>NORMAL: NORMAL instance</li><li>SERVERLESS: SERVERLESS instance</li></ul><p>Default value: NORMAL</p><p>Selectable when DbType is MYSQL (default NORMAL)</p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Required when DbMode is SERVERLESS<br>Minimum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Required when DbMode is SERVERLESS:<br>Maximum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>No auto pause</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul><p>Default value: yes</p><p>Take effect when DbMode is SERVERLESS</p>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>When DbMode is SERVERLESS, specify the delay for Cluster Auto-Pause in seconds. Optional range: [600,691200]<br>Default value: 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>The storage billing mode of the cluster. Pay-As-You-Go: 0, Monthly Subscription: 1. Default is Pay-As-You-Go.<br>When DbType is MYSQL and the compute billing mode of the cluster is postpaid (including DbMode as SERVERLESS), the storage billing mode can only be Pay-As-You-Go.<br>Rollback and clone do not support Monthly Subscription storage.</p>
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// <p>Security group id array</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Alarm policy Id array</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// <p>Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names, 1-case-insensitive, 0-case-sensitive</p>
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// <p>Transaction mode</p><p>Enumeration value:</p><ul><li>0: Place order and pay</li><li>1: Place order</li></ul><p>Default value: 0</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Parameter template ID. The parameter template ID can be obtained through querying parameter template information DescribeParamTemplates.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Multi-AZ address</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Instance initialization configuration information is mainly used to select different specification instances during cluster purchase.</p>
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// <p>Global database unique ID</p>
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// <p>Database proxy configuration</p>
	ProxyConfig *ProxyConfig `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// <p>Auto archive or not</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul><p>Default value: no</p><p>This parameter takes effect only when the primary instance of the current cluster is SERVERLESS</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Archiving processing time after pausing</p><p>Measurement unit: hour</p><p>Default value: 12</p><p>This parameter takes effect only when the primary instance of the current cluster is SERVERLESS.</p>
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// <p>Cluster level, optional. Example: P0, P1. (This field can be ignored)</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// <p>Kernel minor version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`
}

type CreateClustersRequest struct {
	*tchttp.BaseRequest
	
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Database type</p><p>Enumeration value:</p><ul><li>MYSQL: MYSQL</li></ul>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Database version</p><p>Enumeration value:</p><ul><li>5.7: MySQL 5.7 edition</li><li>8.0: MySQL 8.0 edition</li></ul>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>project-ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Required when DbMode is NORMAL or left blank<br>Cpu cores of the regular instance</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Required when DbMode is NORMAL or left blank<br>Ordinary instance memory in GB</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance count</p><p>Value ranges from [1, 16]</p><p>Default value: 2</p><ul><li>A value of 2 means one rw instance + one ro instance.</li><li>The transmitted n refers to one rw instance + n-1 ro instances (identical specifications).</li><li>For a more precise cluster composition collocation, please use InstanceInitInfos.</li><li>The value set by this parameter is suitable for provisioned resource cluster. If needed to set the specifications and quantity of Serverless cluster, please use the InstanceInitInfo structure in InstanceInitInfos.n.</li></ul>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>This parameter is meaningless and abandoned.<br>Storage size, in GB.</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Cluster name, length less than 64 characters, each character value ranges from uppercase/lowercase letters, digits, special symbols ('-', '_', '.')</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Account (8-64 characters, a combination of uppercase and lowercase letters, digits and symbols ~!@#$%^&amp;*_-+=|(){}[]:;&#39;&lt;&gt;,.?/ with any three types required)</p>
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// <p>Port, default 3306, in the range of [0, 65535)</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Billing mode</p><p>Enumeration value:</p><ul><li>0: Pay-as-you-go billing</li><li>1: Monthly Subscription</li></ul><p>Default value: 0</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Number of clusters to purchase. Optional value range: [1,50]. Default is 1.</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>Rollback type</p><p>Enumeration value:</p><ul><li>noneRollback: No rollback</li><li>snapRollback: Snapshot rollback</li><li>timeRollback: Point-in-time rollback</li></ul>
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// <p>Snapshot rollback means snapshotId; point-in-time rollback means queryId. A value of 0 indicates requirement to judge whether the time point is valid.</p>
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// <p>During rollback, input the source cluster ID to search for the source poolId</p>
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// <p>Point-in-time rollback, specified time; snapshot rollback, snapshot time</p>
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// <p>This parameter is meaningless and abandoned.<br>Point-in-time rollback, allowed range of specified time.</p>
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// <p>Ordinary instance storage cap, in GB<br>When DbType is MYSQL and the storage billing mode is prepaid, this parameter should not exceed the maximum storage specification corresponding to cpu and memory.</p>
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// <p>Annual and monthly subscription duration</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>Annual and monthly subscription duration unit, ['s', 'd', 'm', 'y']</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Whether Annual/Monthly Subscription is auto-renewed</p><p>Enumeration value:</p><ul><li>0: Default renewal method</li><li>1: Auto-renewal</li><li>2: No auto-renewal</li></ul><p>Default value: 0</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Whether to automatically select voucher 1 Yes 0 No Default value: 0</p><p>Enumeration value:</p><ul><li>1: Yes</li><li>0: No</li></ul><p>Default value: 0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>Instance count (this parameter is deprecated and only for existing accommodative)</p>
	HaCount *int64 `json:"HaCount,omitnil,omitempty" name:"HaCount"`

	// <p>Order source</p>
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// <p>tag Array info that should be bound for cluster creation</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Db type</p><p>Enumeration value:</p><ul><li>NORMAL: NORMAL instance</li><li>SERVERLESS: SERVERLESS instance</li></ul><p>Default value: NORMAL</p><p>Selectable when DbType is MYSQL (default NORMAL)</p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Required when DbMode is SERVERLESS<br>Minimum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Required when DbMode is SERVERLESS:<br>Maximum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>No auto pause</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul><p>Default value: yes</p><p>Take effect when DbMode is SERVERLESS</p>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>When DbMode is SERVERLESS, specify the delay for Cluster Auto-Pause in seconds. Optional range: [600,691200]<br>Default value: 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>The storage billing mode of the cluster. Pay-As-You-Go: 0, Monthly Subscription: 1. Default is Pay-As-You-Go.<br>When DbType is MYSQL and the compute billing mode of the cluster is postpaid (including DbMode as SERVERLESS), the storage billing mode can only be Pay-As-You-Go.<br>Rollback and clone do not support Monthly Subscription storage.</p>
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// <p>Security group id array</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Alarm policy Id array</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// <p>Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names, 1-case-insensitive, 0-case-sensitive</p>
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// <p>Transaction mode</p><p>Enumeration value:</p><ul><li>0: Place order and pay</li><li>1: Place order</li></ul><p>Default value: 0</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Parameter template ID. The parameter template ID can be obtained through querying parameter template information DescribeParamTemplates.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Multi-AZ address</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Instance initialization configuration information is mainly used to select different specification instances during cluster purchase.</p>
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// <p>Global database unique ID</p>
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// <p>Database proxy configuration</p>
	ProxyConfig *ProxyConfig `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// <p>Auto archive or not</p><p>Enumeration value:</p><ul><li>yes: Yes</li><li>no: No</li></ul><p>Default value: no</p><p>This parameter takes effect only when the primary instance of the current cluster is SERVERLESS</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Archiving processing time after pausing</p><p>Measurement unit: hour</p><p>Default value: 12</p><p>This parameter takes effect only when the primary instance of the current cluster is SERVERLESS.</p>
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// <p>Cluster level, optional. Example: P0, P1. (This field can be ignored)</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// <p>Kernel minor version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`
}

func (r *CreateClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbType")
	delete(f, "DbVersion")
	delete(f, "ProjectId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "InstanceCount")
	delete(f, "Storage")
	delete(f, "ClusterName")
	delete(f, "AdminPassword")
	delete(f, "Port")
	delete(f, "PayMode")
	delete(f, "Count")
	delete(f, "RollbackStrategy")
	delete(f, "RollbackId")
	delete(f, "OriginalClusterId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "HaCount")
	delete(f, "OrderSource")
	delete(f, "ResourceTags")
	delete(f, "DbMode")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "StoragePayMode")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "SlaveZone")
	delete(f, "InstanceInitInfos")
	delete(f, "GdnId")
	delete(f, "ProxyConfig")
	delete(f, "AutoArchive")
	delete(f, "AutoArchiveDelayHours")
	delete(f, "ClusterLevel")
	delete(f, "CynosVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClustersResponseParams struct {
	// <p>Frozen transaction ID</p>
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// <p>Order ID</p>
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// <p>Resource ID list (This field is no longer maintained. Please use the dealNames field to query the API DescribeResourcesByDealName to obtain resource IDs)</p>
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>Cluster ID list (This field is no longer maintained. Please use the dealNames field and the query API DescribeResourcesByDealName to get the cluster ID)</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>Large Order ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClustersResponse struct {
	*tchttp.BaseResponse
	Response *CreateClustersResponseParams `json:"Response"`
}

func (r *CreateClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrateClusterRequestParams struct {
	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specifies the ID of the VPC network it belongs to.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database version. valid values:. 
	// <li>Specifies the available values for MYSQL: 5.7, 8.0.</li>.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Cluster name, length less than 64 characters. each character value ranges from uppercase/lowercase letters, digits, special symbols ('-','_','.').
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Account password (8-64 characters, a combination of uppercase and lowercase letters, digits and symbols ~!@#$%^&*_-+=|\(){}[]:;'<>,.?/ with any three types).
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Port, default 3306, in the range of [0, 65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Billing mode. 0: pay-as-you-go; 1: monthly subscription. default is pay-as-you-go.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Number of clusters to purchase. value range: [1,3]. default is 1.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Maximum storage limit of a regular instance, in GB.
	// When DbType is MYSQL and the storage billing mode is prepaid, this parameter should not exceed the maximum storage specification corresponding to cpu and memory.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Specifies the annual and monthly subscription duration.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Specifies the measurement unit for annual and monthly subscription duration. valid values: 's', 'd', 'm', 'y'.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether annual/monthly subscription is auto-renewed. default value is 0.
	// 0 identifies the default renewal method, 1 means auto-renew, 2 indicates no auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. 1: yes; 0: no. default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Specifies the tag array information that needs to be bound during cluster creation.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Specifies the cluster storage billing mode. 0: pay-as-you-go; 1: monthly subscription. default is pay-as-you-go.
	// When DbType is MYSQL and the cluster billing mode for computing is postpaid (including DbMode as SERVERLESS), the storage billing mode can only be pay-as-you-go.
	// Rollback and clone do not support monthly subscription storage.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Security group id array.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Specifies the Alarm policy Id array.
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names. valid values: 1 (case-insensitive), 0 (case-sensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// Transaction mode. valid values: 0 (place order and pay), 1 (place order).
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID. can be obtained through querying parameter template information DescribeParamTemplates.
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Initializes configuration information, mainly used to purchase clusters with different specification instances.
	InstanceInitInfos []*IntegrateInstanceInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// Global database unique identifier.
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// Database proxy configuration.
	ProxyConfig *ProxyConfigInfo `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// Specifies whether to automatically archive.
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// Processing time after pausing.
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// Encryption method (consists of encryption algorithm and key pair version).
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// Describes the cluster configuration information.
	IntegrateCreateClusterConfig *IntegrateCreateClusterConfig `json:"IntegrateCreateClusterConfig,omitnil,omitempty" name:"IntegrateCreateClusterConfig"`

	// Storage architecture type. valid values: 1.0/2.0. default value: 1.0.
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

type CreateIntegrateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specifies the ID of the VPC network it belongs to.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database version. valid values:. 
	// <li>Specifies the available values for MYSQL: 5.7, 8.0.</li>.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Cluster name, length less than 64 characters. each character value ranges from uppercase/lowercase letters, digits, special symbols ('-','_','.').
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Account password (8-64 characters, a combination of uppercase and lowercase letters, digits and symbols ~!@#$%^&*_-+=|\(){}[]:;'<>,.?/ with any three types).
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Port, default 3306, in the range of [0, 65535).
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Billing mode. 0: pay-as-you-go; 1: monthly subscription. default is pay-as-you-go.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Number of clusters to purchase. value range: [1,3]. default is 1.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Maximum storage limit of a regular instance, in GB.
	// When DbType is MYSQL and the storage billing mode is prepaid, this parameter should not exceed the maximum storage specification corresponding to cpu and memory.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Specifies the annual and monthly subscription duration.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Specifies the measurement unit for annual and monthly subscription duration. valid values: 's', 'd', 'm', 'y'.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Whether annual/monthly subscription is auto-renewed. default value is 0.
	// 0 identifies the default renewal method, 1 means auto-renew, 2 indicates no auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher. 1: yes; 0: no. default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Specifies the tag array information that needs to be bound during cluster creation.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Specifies the cluster storage billing mode. 0: pay-as-you-go; 1: monthly subscription. default is pay-as-you-go.
	// When DbType is MYSQL and the cluster billing mode for computing is postpaid (including DbMode as SERVERLESS), the storage billing mode can only be pay-as-you-go.
	// Rollback and clone do not support monthly subscription storage.
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Security group id array.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Specifies the Alarm policy Id array.
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names. valid values: 1 (case-insensitive), 0 (case-sensitive).
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// Transaction mode. valid values: 0 (place order and pay), 1 (place order).
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Parameter template ID. can be obtained through querying parameter template information DescribeParamTemplates.
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Multi-AZ address.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Initializes configuration information, mainly used to purchase clusters with different specification instances.
	InstanceInitInfos []*IntegrateInstanceInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// Global database unique identifier.
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// Database proxy configuration.
	ProxyConfig *ProxyConfigInfo `json:"ProxyConfig,omitnil,omitempty" name:"ProxyConfig"`

	// Specifies whether to automatically archive.
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// Processing time after pausing.
	AutoArchiveDelayHours *int64 `json:"AutoArchiveDelayHours,omitnil,omitempty" name:"AutoArchiveDelayHours"`

	// Encryption method (consists of encryption algorithm and key pair version).
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// Describes the cluster configuration information.
	IntegrateCreateClusterConfig *IntegrateCreateClusterConfig `json:"IntegrateCreateClusterConfig,omitnil,omitempty" name:"IntegrateCreateClusterConfig"`

	// Storage architecture type. valid values: 1.0/2.0. default value: 1.0.
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

func (r *CreateIntegrateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersion")
	delete(f, "ProjectId")
	delete(f, "ClusterName")
	delete(f, "AdminPassword")
	delete(f, "Port")
	delete(f, "PayMode")
	delete(f, "Count")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTags")
	delete(f, "StoragePayMode")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "DealMode")
	delete(f, "ParamTemplateId")
	delete(f, "SlaveZone")
	delete(f, "InstanceInitInfos")
	delete(f, "GdnId")
	delete(f, "ProxyConfig")
	delete(f, "AutoArchive")
	delete(f, "AutoArchiveDelayHours")
	delete(f, "EncryptMethod")
	delete(f, "IntegrateCreateClusterConfig")
	delete(f, "StorageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrateClusterResponseParams struct {
	// Frozen transaction ID.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Order ID.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Resource ID list (this field is no longer maintained. please use the dealNames field and the query API DescribeResourcesByDealName to obtain resource ids).
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Cluster ID list (this field is no longer maintained. please use the dealNames field to get cluster ids via the DescribeResourcesByDealName api.).
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Large order number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrateClusterResponseParams `json:"Response"`
}

func (r *CreateIntegrateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraDBClusterAccountsRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

type CreateLibraDBClusterAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Accounts []*NewAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

func (r *CreateLibraDBClusterAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraDBClusterAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	delete(f, "EncryptMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLibraDBClusterAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraDBClusterAccountsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLibraDBClusterAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CreateLibraDBClusterAccountsResponseParams `json:"Response"`
}

func (r *CreateLibraDBClusterAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraDBClusterAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraDBClustersRequestParams struct {
	// Quantity.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance initialization information
	InstanceInitInfos []*LibraDBInstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// user password
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Whether to perform auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Order placement mode
	DealMode *string `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// LibraDBVersion version. Defaults to the latest version.
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// Order Source
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Duration
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Instance creation binds Tag array information
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// VPC ID of the cluster location
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID of the cluster location
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type CreateLibraDBClustersRequest struct {
	*tchttp.BaseRequest
	
	// Quantity.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance initialization information
	InstanceInitInfos []*LibraDBInstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// user password
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Whether to perform auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to automatically select a voucher.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Order placement mode
	DealMode *string `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// LibraDBVersion version. Defaults to the latest version.
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// Order Source
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Security group
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Duration
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Instance creation binds Tag array information
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// VPC ID of the cluster location
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID of the cluster location
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

func (r *CreateLibraDBClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraDBClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Count")
	delete(f, "Zone")
	delete(f, "InstanceInitInfos")
	delete(f, "AdminPassword")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "ClusterName")
	delete(f, "DealMode")
	delete(f, "EncryptMethod")
	delete(f, "LibraDBVersion")
	delete(f, "OrderSource")
	delete(f, "PayMode")
	delete(f, "ProjectId")
	delete(f, "SecurityGroupIds")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ResourceTags")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLibraDBClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraDBClustersResponseParams struct {
	// Prepaid Total Order Number
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Each resource corresponds to a dealName. The business needs to ensure the delivery API idempotency based on the dealName.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Freeze transaction
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Instance ID.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLibraDBClustersResponse struct {
	*tchttp.BaseResponse
	Response *CreateLibraDBClustersResponseParams `json:"Response"`
}

func (r *CreateLibraDBClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraDBClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// MySQL version number
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// ID of the template to be copied
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Database type. Valid values:  `NORMAL` (default), `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// List of the parameters
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// MySQL version number
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// ID of the template to be copied
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Database type. Valid values:  `NORMAL` (default), `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// List of the parameters
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "EngineVersion")
	delete(f, "TemplateDescription")
	delete(f, "TemplateId")
	delete(f, "DbMode")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyEndPointRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>VPC ID.</p>
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// <p>VPC subnet ID.</p>
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// <p>Connection pool Type: SessionConnectionPool (session-level connection pool).</p>
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// <p>Whether the connection pool is enabled.<br>yes: enabled.<br>no: disabled.</p>
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// <p>Connection pool threshold: measurement unit (seconds), optional range: 0 - 300 seconds.</p>
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// <p>Bound security group ID array.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>vip information that should be bound must correspond to UniqueVpcId.</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Weight mode:<br>system: system-assigned.<br>custom: custom.</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Automatically add read-only instance.<br>yes: means automatically add read-only instance.<br>no: means not automatic add read-only instance.</p>
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether fault migration is enabled.<br>yes: Enable. Once enabled, when a database proxy failure occurs, the connection address will be routed to the primary instance.<br>no: Disabled.<br>Description:<br>This item can be set to only when the RwType parameter value is READWRITE.</p>
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Consistency type:<br>eventual: Final consistency.<br>global: Global consistency.<br>session: Session consistency.<br>Description:<br>This item can be set to only when the RwType parameter value is READWRITE.</p>
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// <p>Read-write attribute:<br>READWRITE: means read-write separation. When the parameter value is READWRITE, FailOver and ConsistencyType parameters can be set.<br>READONLY: means read-only.</p>
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// <p>Consistency timeout. Value ranges from 0 to 1000000 (microseconds). When set to 0, it means the request will wait if the read-only instance fails to satisfy the consistency policy due to latency.</p>
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// <p>Whether to enable transaction split. Once enabled, read and write operations in a transaction are split and executed on different instances.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Access mode:<br>nearby: proximity access.<br>balance: balanced allocation.</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Instance weight.</p>
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// <p>Load balancing mode</p><p>Enumeration value:</p><ul><li>static: Static load</li><li>dynamic: Dynamic load</li></ul>
	LoadBalanceMode *string `json:"LoadBalanceMode,omitnil,omitempty" name:"LoadBalanceMode"`
}

type CreateProxyEndPointRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>VPC ID.</p>
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// <p>VPC subnet ID.</p>
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// <p>Connection pool Type: SessionConnectionPool (session-level connection pool).</p>
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// <p>Whether the connection pool is enabled.<br>yes: enabled.<br>no: disabled.</p>
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// <p>Connection pool threshold: measurement unit (seconds), optional range: 0 - 300 seconds.</p>
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// <p>Bound security group ID array.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>vip information that should be bound must correspond to UniqueVpcId.</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Weight mode:<br>system: system-assigned.<br>custom: custom.</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Automatically add read-only instance.<br>yes: means automatically add read-only instance.<br>no: means not automatic add read-only instance.</p>
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether fault migration is enabled.<br>yes: Enable. Once enabled, when a database proxy failure occurs, the connection address will be routed to the primary instance.<br>no: Disabled.<br>Description:<br>This item can be set to only when the RwType parameter value is READWRITE.</p>
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Consistency type:<br>eventual: Final consistency.<br>global: Global consistency.<br>session: Session consistency.<br>Description:<br>This item can be set to only when the RwType parameter value is READWRITE.</p>
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// <p>Read-write attribute:<br>READWRITE: means read-write separation. When the parameter value is READWRITE, FailOver and ConsistencyType parameters can be set.<br>READONLY: means read-only.</p>
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// <p>Consistency timeout. Value ranges from 0 to 1000000 (microseconds). When set to 0, it means the request will wait if the read-only instance fails to satisfy the consistency policy due to latency.</p>
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// <p>Whether to enable transaction split. Once enabled, read and write operations in a transaction are split and executed on different instances.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Access mode:<br>nearby: proximity access.<br>balance: balanced allocation.</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Instance weight.</p>
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// <p>Load balancing mode</p><p>Enumeration value:</p><ul><li>static: Static load</li><li>dynamic: Dynamic load</li></ul>
	LoadBalanceMode *string `json:"LoadBalanceMode,omitnil,omitempty" name:"LoadBalanceMode"`
}

func (r *CreateProxyEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "Vip")
	delete(f, "WeightMode")
	delete(f, "AutoAddRo")
	delete(f, "FailOver")
	delete(f, "ConsistencyType")
	delete(f, "RwType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "InstanceWeights")
	delete(f, "LoadBalanceMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyEndPointResponseParams struct {
	// <p>Async process ID.</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Asynchronous Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Database Proxy Group ID.</p>
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxyEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyEndPointResponseParams `json:"Response"`
}

func (r *CreateProxyEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Database node information
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type CreateProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// VPC ID, which is the same as that of the cluster by default.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID, which is the same as that of the cluster by default.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Connection pool type. Valid value: `SessionConnectionPool` (session-level connection pool)
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether to enable connection pool. Valid value: `yes` (enable), `no` (disable).
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Connection pool threshold in seconds
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Array of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Database node information
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

func (r *CreateProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "UniqueVpcId")
	delete(f, "UniqueSubnetId")
	delete(f, "ProxyCount")
	delete(f, "ConnectionPoolType")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "SecurityGroupIds")
	delete(f, "Description")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxyResponseParams `json:"Response"`
}

func (r *CreateProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageRequestParams struct {
	// Instance type. currently fixed to cynosdb-serverless.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Resource package region of use: chineseMainland - common in the chinese mainland. overseas - universally applicable in hong kong (china), macao (china), taiwan (china), and overseas.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack edition. Valid values: `base` (basic edition), `common` (general edition), `enterprise` (enterprise edition).
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Resource package size. Compute resource unit: unit; storage resource: GB.
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// Validity period of a resource pack in days
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// Number of the resource packs to be purchased
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// Resource pack name
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type CreateResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// Instance type. currently fixed to cynosdb-serverless.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Resource package region of use: chineseMainland - common in the chinese mainland. overseas - universally applicable in hong kong (china), macao (china), taiwan (china), and overseas.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack edition. Valid values: `base` (basic edition), `common` (general edition), `enterprise` (enterprise edition).
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Resource package size. Compute resource unit: unit; storage resource: GB.
	PackageSpec *float64 `json:"PackageSpec,omitnil,omitempty" name:"PackageSpec"`

	// Validity period of a resource pack in days
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`

	// Number of the resource packs to be purchased
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// Resource pack name
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *CreateResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "PackageVersion")
	delete(f, "PackageSpec")
	delete(f, "ExpireDay")
	delete(f, "PackageCount")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourcePackageResponseParams struct {
	// Big order number
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Each item has only one `dealName`, through which you need to ensure the idempotency of the delivery API.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourcePackageResponseParams `json:"Response"`
}

func (r *CreateResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVaultRequestParams struct {
	// Safe name, length must be greater than 0
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Backup retention period (seconds). Must be greater than 0.
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Safe description
	VaultDescribe *string `json:"VaultDescribe,omitnil,omitempty" name:"VaultDescribe"`

	// KMS key ID, length 0-36 characters
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key type, available values: cloud (KMS managed keys), custom (custom key)
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key region, length 0-32 characters
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Lock time, format: YYYY-MM-DD HH:mm:ss
	LockedTime *string `json:"LockedTime,omitnil,omitempty" name:"LockedTime"`
}

type CreateVaultRequest struct {
	*tchttp.BaseRequest
	
	// Safe name, length must be greater than 0
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Backup retention period (seconds). Must be greater than 0.
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Safe description
	VaultDescribe *string `json:"VaultDescribe,omitnil,omitempty" name:"VaultDescribe"`

	// KMS key ID, length 0-36 characters
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key type, available values: cloud (KMS managed keys), custom (custom key)
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key region, length 0-32 characters
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Lock time, format: YYYY-MM-DD HH:mm:ss
	LockedTime *string `json:"LockedTime,omitnil,omitempty" name:"LockedTime"`
}

func (r *CreateVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultName")
	delete(f, "BackupSaveSeconds")
	delete(f, "VaultDescribe")
	delete(f, "KeyId")
	delete(f, "KeyType")
	delete(f, "KeyRegion")
	delete(f, "LockedTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVaultResponseParams struct {
	// Task ID, used for query task execution status
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVaultResponse struct {
	*tchttp.BaseResponse
	Response *CreateVaultResponseParams `json:"Response"`
}

func (r *CreateVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossRegionBackupItem struct {
	// Target region for backup.
	CrossRegion *string `json:"CrossRegion,omitnil,omitempty" name:"CrossRegion"`

	// Target region's backup task ID.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup status in the target region.
	BackupStatus *string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`
}

type CynosdbCluster struct {
	// <p>Cluster status. Valid values include:<br>creating: Creating<br>running: Running<br>isolating: Isolation<br>isolated: Isolated<br>activating: De-isolating<br>offlining: Offline<br>offlined: Offline<br>deleting: Deleting<br>deleted: Deleted</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Database version</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Number of instances</p>
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// <p>User uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>Engine type</p>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>User appid</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cluster status description</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Cluster creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Payment mode. 0: pay-as-you-go; 1: monthly subscription</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>End time</p>
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Cluster read/write vip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Cluster read/write vport</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Project ID.</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>cynos kernel version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// <p>cynos version tag</p>
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// <p>Storage capacity</p>
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// <p>Renewal flag</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Currently processing task</p>
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// <p>Task array of the cluster</p>
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>tag Array of cluster binding</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Db type (NORMAL, SERVERLESS)</p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>When the Db type is SERVERLESS, the SERVERLESS cluster status. Available values:<br>resume<br>pause</p>
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// <p>Cluster Prepaid storage size</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Storage ID of cluster storage in prepayment, used for prepaid storage configuration change</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>Cluster storage payment mode. 0: pay-as-you-go; 1: monthly subscription</p>
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// <p>Minimum storage of the cluster compute specification</p>
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// <p>Maximum storage value of the cluster compute specification</p>
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// <p>Cluster network information</p>
	NetAddrs []*NetAddr `json:"NetAddrs,omitnil,omitempty" name:"NetAddrs"`

	// <p>Physical AZ</p>
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// <p>Primary AZ</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>Whether there is a secondary AZ</p>
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// <p>Secondary AZ</p>
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// <p>Business type</p>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// <p>Freeze or not</p>
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// <p>Order source</p>
	OrderSource *string `json:"OrderSource,omitnil,omitempty" name:"OrderSource"`

	// <p>Capacity</p>
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// <p>Instance bind resource package info (here only return storage resource package, for example packageType=DISK)</p>
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// <p>Global database unique ID</p>
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// <p>Cluster role. Primary cluster - primary, slave cluster - standby. If GdnId is empty, the field is invalid.</p>
	GdnRole *string `json:"GdnRole,omitnil,omitempty" name:"GdnRole"`
}

type CynosdbClusterDetail struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Physical AZ</p>
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// <p>Status. Supported values are as follows:</p><ul><li>creating: Creating</li><li>running: Running</li><li>isolating: Isolation</li><li>isolated: Isolated</li><li>activating: Restore from recycle bin</li><li>offlining: Offline</li><li>offlined: Offline</li><li>deleting: Deleting</li><li>deleted: Deleted</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>When the Db type is SERVERLESS, the SERVERLESS cluster status. Available values:<br>resume<br>resuming<br>pause<br>pausing</p>
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// <p>Storage Id</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>Storage size in GB</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Maximum storage specification, in GB</p>
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// <p>Minimum storage specification, in GB</p>
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// <p>Storage billing type. Valid values: 1 (yearly/monthly subscription); 0 (pay-as-you-go).</p>
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// <p>VPC name</p>
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// <p>vpc Unique id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet name.</p>
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Character set.</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Database type</p>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Db type: <li>NORMAL</li><li>SERVERLESS</li></p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Database version</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>Storage space limit</p>
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// <p>Used capacity</p>
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// <p>vip address</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>vport port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vip address and vport of the cluster read-only instance</p>
	RoAddr []*Addr `json:"RoAddr,omitnil,omitempty" name:"RoAddr"`

	// <p>Functions supported by the cluster</p>
	Ability *Ability `json:"Ability,omitnil,omitempty" name:"Ability"`

	// <p>cynos version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// <p>Business type</p>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// <p>Whether there is a secondary AZ</p>
	HasSlaveZone *string `json:"HasSlaveZone,omitnil,omitempty" name:"HasSlaveZone"`

	// <p>Freeze or not</p>
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// <p>Task List</p>
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Primary AZ</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>From the AZ list</p>
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// <p>Instance information</p>
	InstanceSet []*ClusterInstanceDetail `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// <p>Payment mode</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Expiry time.</p>
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Project ID.</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>tag Array information for instance binding</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Proxy status</p>
	ProxyStatus *string `json:"ProxyStatus,omitnil,omitempty" name:"ProxyStatus"`

	// <p>binlog switch, available values: ON, OFF</p>
	LogBin *string `json:"LogBin,omitnil,omitempty" name:"LogBin"`

	// <p>Skip transaction or not</p>
	IsSkipTrade *string `json:"IsSkipTrade,omitnil,omitempty" name:"IsSkipTrade"`

	// <p>PITR type, available values: normal, redo_pitr</p>
	PitrType *string `json:"PitrType,omitnil,omitempty" name:"PitrType"`

	// <p>Whether to toggle on password complexity</p>
	IsOpenPasswordComplexity *string `json:"IsOpenPasswordComplexity,omitnil,omitempty" name:"IsOpenPasswordComplexity"`

	// <p>Network type</p>
	NetworkStatus *string `json:"NetworkStatus,omitnil,omitempty" name:"NetworkStatus"`

	// <p>Package info of the bound resource for the cluster</p>
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// <p>Auto-renewal flag. 1 means auto-renewal, 0 means non-renewal upon expiration.</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Node network type</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>Secondary availability zone property</p>
	SlaveZoneAttr []*SlaveZoneAttrItem `json:"SlaveZoneAttr,omitnil,omitempty" name:"SlaveZoneAttr"`

	// <p>Version tag</p>
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// <p>Unique ID of the global database network</p>
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// <p>Role of the cluster in the global data network.<br>Primary cluster - primary<br>Slave cluster - standby<br>If empty, the field is invalid.</p>
	GdnRole *string `json:"GdnRole,omitnil,omitempty" name:"GdnRole"`

	// <p>Second-level storage usage in GB</p>
	UsedArchiveStorage *int64 `json:"UsedArchiveStorage,omitnil,omitempty" name:"UsedArchiveStorage"`

	// <p>Archiving status. Enumeration value <li>normal: Normal</li><li>archiving: Archiving</li><li>resuming: Recovering</li><li>archived: Archived</li></p>
	ArchiveStatus *string `json:"ArchiveStatus,omitnil,omitempty" name:"ArchiveStatus"`

	// <p>Archive progress, percentage.</p>
	ArchiveProgress *int64 `json:"ArchiveProgress,omitnil,omitempty" name:"ArchiveProgress"`

	// <p>Cluster level. For example P0, P1</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// <p>Whether to enable transparent data encryption</p>
	IsOpenTDE *bool `json:"IsOpenTDE,omitnil,omitempty" name:"IsOpenTDE"`
}

type CynosdbErrorLogItem struct {
	// Log timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Log level.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Log content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CynosdbInstance struct {
	// <p>User Uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>User AppId</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Project ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance status description in Chinese</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Instance form, whether it is a serverless instance</p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Database type</p>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Database version.</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>Cpu, unit: core</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Memory, unit: GB</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Stored amount, unit: GB</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Instance type</p><p>Enumeration value:</p><ul><li>rw: Read-write instance</li><li>ro: Read-only instance</li></ul>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Current instance role</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Instance private IP address</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Instance Intranet Port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Payment mode</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Instance expiration time</p>
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Termination phase</p>
	DestroyDeadlineText *string `json:"DestroyDeadlineText,omitnil,omitempty" name:"DestroyDeadlineText"`

	// <p>Isolation time</p>
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// <p>Network type</p>
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>Public network domain name</p>
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// <p>Public network IP</p>
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// <p>Public network port</p>
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// <p>Public network status</p>
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// <p>Instance destruction time</p>
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// <p>Cynos kernel version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// <p>Currently processing task</p>
	ProcessingTask *string `json:"ProcessingTask,omitnil,omitempty" name:"ProcessingTask"`

	// <p>Renewal flag</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>serverless instance cpu minimum</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>cpu cap of the serverless instance</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>serverless instance status, available values:<br>resume<br>pause</p>
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// <p>Prepayment storage Id</p>
	StorageId *string `json:"StorageId,omitnil,omitempty" name:"StorageId"`

	// <p>Storage billing type</p>
	StoragePayMode *int64 `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// <p>Physical zone</p>
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// <p>Business type</p>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// <p>Task</p>
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Whether to freeze</p>
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// <p>Resource tag</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Primary AZ</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>Standby availability zone</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// <p>Instance network info</p>
	InstanceNetInfo []*InstanceNetInfo `json:"InstanceNetInfo,omitnil,omitempty" name:"InstanceNetInfo"`

	// <p>Instance bind resource package info (only return compute resource package, packageType=CCU)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourcePackages []*ResourcePackage `json:"ResourcePackages,omitnil,omitempty" name:"ResourcePackages"`

	// <p>Instance index form, available values [mixedRowColumn (hybrid row-column storage), onlyRowIndex (row storage only)]</p>
	InstanceIndexMode *string `json:"InstanceIndexMode,omitnil,omitempty" name:"InstanceIndexMode"`

	// <p>Supported capabilities of the current instance</p>
	InstanceAbility *InstanceAbility `json:"InstanceAbility,omitnil,omitempty" name:"InstanceAbility"`

	// <p>Instance Machine Type</p><ol><li>common, universal type.</li><li>exclusive, dedicated.</li></ol>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Instance storage type</p>
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`

	// <p>Unknown field</p>
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// <p>libradb node information</p>
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// <p>Unique ID of the global database</p>
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`
}

type CynosdbInstanceDetail struct {
	// <p>User Uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>User AppId</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Project ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance status<br>creating: Creating<br>running: Running<br>isolating: Isolation<br>isolated: Isolated<br>activating: Recovering<br>offlining: Offline<br>offlined: Offline</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description in Chinese of the instance</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Status of the serverless instance. Possible value:<br>resume<br>pause</p>
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`

	// <p>Database type.</p>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Database version</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>Cpu, unit: cores</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Memory, unit: GB</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Stored amount, unit: GB</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Instance type</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Current instance role</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Payment mode</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Instance expiration time</p>
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Network type.</p>
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>VPC network ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Instance private IP address</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Instance Intranet Port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Instance public network domain name</p>
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// <p>Character set</p>
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// <p>Cynos kernel version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// <p>Renewal flag</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>serverless instance cpu minimum</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>cpu cap of the serverless instance</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>Db type:<li>NORMAL</li><li>SERVERLESS</li></p>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Cluster read/write instance Availability Zone</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`
}

type CynosdbInstanceGrp struct {
	// User `appId`
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Deletion time
	DeletedTime *string `json:"DeletedTime,omitnil,omitempty" name:"DeletedTime"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance group type. ha: HA group; ro: RO group
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Update time
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// Private IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public IP
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public port
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// Information of instances contained in instance group
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Information of the old IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldAddrInfo *OldAddrInfo `json:"OldAddrInfo,omitnil,omitempty" name:"OldAddrInfo"`

	// Task in progress
	ProcessingTasks []*string `json:"ProcessingTasks,omitnil,omitempty" name:"ProcessingTasks"`

	// Task list
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// biz_net_service table ID
	NetServiceId *int64 `json:"NetServiceId,omitnil,omitempty" name:"NetServiceId"`
}

type DataSourceItem struct {
	// Source instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Source cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Source database type
	DBType *string `json:"DBType,omitnil,omitempty" name:"DBType"`

	// Source database IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Source database port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Source instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Source instance availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Source main account uin
	SrcUin *string `json:"SrcUin,omitnil,omitempty" name:"SrcUin"`

	// Account type
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// sync task status
	ReplicationJobStatus *string `json:"ReplicationJobStatus,omitnil,omitempty" name:"ReplicationJobStatus"`
}

type DatabasePrivileges struct {
	// Database.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DatabaseTables struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name list.
	Tables []*string `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type DbInfo struct {
	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`

	// Database status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Collation
	CollateRule *string `json:"CollateRule,omitnil,omitempty" name:"CollateRule"`

	// Specifies the remark of the database.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// User permission
	UserHostPrivileges []*UserHostPrivilege `json:"UserHostPrivileges,omitnil,omitempty" name:"UserHostPrivileges"`

	// Database ID
	DbId *int64 `json:"DbId,omitnil,omitempty" name:"DbId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// User appid.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DbTable struct {
	// Database name.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Database table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Accounts in array, which contains `account` and `host`.
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Accounts in array, which contains `account` and `host`.
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountsResponseParams `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

func (r *DeleteAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DeleteAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID. This field is used by legacy versions and thus not recommended.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup file ID. This field is recommended.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SnapshotIdList")
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupResponseParams `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupVaultRequestParams struct {
	// Backup safe ID. The length must be greater than 0.
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup file ID list to be deleted, cannot be empty
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

type DeleteBackupVaultRequest struct {
	*tchttp.BaseRequest
	
	// Backup safe ID. The length must be greater than 0.
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup file ID list to be deleted, cannot be empty
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`
}

func (r *DeleteBackupVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "BackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupVaultResponseParams struct {
	// task ID, for querying task execution status
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupVaultResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupVaultResponseParams `json:"Response"`
}

func (r *DeleteBackupVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCLSDeliveryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type DeleteCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *DeleteCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCLSDeliveryResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCLSDeliveryResponseParams `json:"Response"`
}

func (r *DeleteCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterDatabaseRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name.
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
}

type DeleteClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name.
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`
}

func (r *DeleteClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterDatabaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterDatabaseResponseParams `json:"Response"`
}

func (r *DeleteClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterSaveBackupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Reserved backup file ID. recommended.
	SaveBackupId *int64 `json:"SaveBackupId,omitnil,omitempty" name:"SaveBackupId"`
}

type DeleteClusterSaveBackupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Reserved backup file ID. recommended.
	SaveBackupId *int64 `json:"SaveBackupId,omitnil,omitempty" name:"SaveBackupId"`
}

func (r *DeleteClusterSaveBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterSaveBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SaveBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterSaveBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterSaveBackupResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterSaveBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterSaveBackupResponseParams `json:"Response"`
}

func (r *DeleteClusterSaveBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterSaveBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibraDBClusterAccountsRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteLibraDBClusterAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Accounts []*InputAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *DeleteLibraDBClusterAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraDBClusterAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLibraDBClusterAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibraDBClusterAccountsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLibraDBClusterAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLibraDBClusterAccountsResponseParams `json:"Response"`
}

func (r *DeleteLibraDBClusterAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraDBClusterAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibraDBClusterRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteLibraDBClusterRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteLibraDBClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraDBClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLibraDBClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibraDBClusterResponseParams struct {
	// flow id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLibraDBClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLibraDBClusterResponseParams `json:"Response"`
}

func (r *DeleteLibraDBClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraDBClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVaultTask struct {
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type DeleteVaultsRequestParams struct {
	// List of backup safe IDs to be deleted, cannot be empty. The safe must have all files cleared.
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

type DeleteVaultsRequest struct {
	*tchttp.BaseRequest
	
	// List of backup safe IDs to be deleted, cannot be empty. The safe must have all files cleared.
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`
}

func (r *DeleteVaultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVaultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVaultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVaultsResponseParams struct {
	// Delete task list. Each safe corresponds to a task.
	VaultTask []*DeleteVaultTask `json:"VaultTask,omitnil,omitempty" name:"VaultTask"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVaultsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVaultsResponseParams `json:"Response"`
}

func (r *DeleteVaultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVaultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliverSummary struct {
	// Select the delivery type, storage class, message channel.
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// Specifies the delivery subtype: cls, ckafka.
	DeliverSubType *string `json:"DeliverSubType,omitnil,omitempty" name:"DeliverSubType"`

	// Sender.
	DeliverConsumer *string `json:"DeliverConsumer,omitnil,omitempty" name:"DeliverConsumer"`

	// Specifies the name of the sender.
	DeliverConsumerName *string `json:"DeliverConsumerName,omitnil,omitempty" name:"DeliverConsumerName"`

	// Exception error in delivery
	DeliverError *string `json:"DeliverError,omitnil,omitempty" name:"DeliverError"`
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// When the database name is “*”, the value specified in `Type` and `TableName` will be ignored, indicating that the user's global permissions are being modified.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Object type in a specified database. Valid values: `table`, `*`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The database name can be specified when `Type` is 'table'.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// When the database name is “*”, the value specified in `Type` and `TableName` will be ignored, indicating that the user's global permissions are being modified.
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Object type in a specified database. Valid values: `table`, `*`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// The database name can be specified when `Type` is 'table'.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "Host")
	delete(f, "Db")
	delete(f, "Type")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// The list of permissions, such as  ["select","update","delete","create","drop","references","index","alter","show_db","create_tmp_table","lock_tables","execute","create_view","show_view","create_routine","alter_routine","event","trigger"]
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Keywords for fuzzy search (match `AccountName` and `AccountHost` at the same time), which supports regex. The union results will be returned.
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of accounts to be filtered
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// List of accounts to be filtered
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Keywords for fuzzy search (match `AccountName` and `AccountHost` at the same time), which supports regex. The union results will be returned.
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountNames")
	delete(f, "DbType")
	delete(f, "Hosts")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountRegular")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Database account list.
	AccountSet []*Account `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// Total number of accounts
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListRequestParams struct {
	// Enabling status of instance audit. 1 - enabled; 0 - not enabled.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// Filtering conditions for querying the instance list.
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Instance audit rule mode. 1 - rule-based audit; 0 - full audit.
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// Number of entries returned per request. The default value is 30, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// Enabling status of instance audit. 1 - enabled; 0 - not enabled.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// Filtering conditions for querying the instance list.
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Instance audit rule mode. 1 - rule-based audit; 0 - full audit.
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// Number of entries returned per request. The default value is 30, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditSwitch")
	delete(f, "Filters")
	delete(f, "AuditMode")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// Total number of instances meeting the query conditions.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of detailed information about the audit instance.
	Items []*InstanceAuditStatus `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesRequestParams struct {
	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of results returned per request. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleTemplateNames")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesResponseParams struct {
	// Number of eligible instances
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Rule template detail list.
	Items []*AuditRuleTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsRequestParams struct {
	// Instance ID. Currently, only one single instance can be queried.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeAuditRuleWithInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Currently, only one single instance can be queried.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeAuditRuleWithInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleWithInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleWithInstanceIdsResponseParams struct {
	// None
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit rule information of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*InstanceAuditRule `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRuleWithInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleWithInstanceIdsResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleWithInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleWithInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigResponseParams struct {
	// <p>Indicates the full backup start time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// <p>Indicates the full backup start time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// <p>Indicates the backup retention period in seconds. Backups longer than this time will be cleaned up. 7 days is represented as 3600<em>24</em>7=604800.</p>
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// <p>Backup frequency, an array of length 7, corresponding to Monday to Sunday backup method, full-full backup, increment-incremental backup</p>
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// <p>Backup method, logic-logical backup, snapshot-snapshot backup</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Cross-regional logical backup configuration modification time</p>
	LogicCrossRegionsConfigUpdateTime *string `json:"LogicCrossRegionsConfigUpdateTime,omitnil,omitempty" name:"LogicCrossRegionsConfigUpdateTime"`

	// <p>Automatic logical backup configuration</p>
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// <p>Second-level snapshot backup configuration information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotSecondaryBackupConfig *BackupConfigInfo `json:"SnapshotSecondaryBackupConfig,omitnil,omitempty" name:"SnapshotSecondaryBackupConfig"`

	// <p>Sparse backup configuration</p>
	SparseBackupConfig *SparseBackupConfigRsp `json:"SparseBackupConfig,omitnil,omitempty" name:"SparseBackupConfig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {
	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// Cluster backup download limit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupLimitClusterRestrictions []*BackupLimitClusterRestriction `json:"BackupLimitClusterRestrictions,omitnil,omitempty" name:"BackupLimitClusterRestrictions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUrlRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup download source restriction condition.
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

type DescribeBackupDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup download source restriction condition.
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

func (r *DescribeBackupDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	delete(f, "DownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUrlResponseParams struct {
	// Backup download address
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUserRestrictionRequestParams struct {
	// Pagination size
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Specifies whether to query only user-level download limits. true - yes, false - no.
	OnlyUserRestriction *bool `json:"OnlyUserRestriction,omitnil,omitempty" name:"OnlyUserRestriction"`
}

type DescribeBackupDownloadUserRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Pagination size
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Specifies whether to query only user-level download limits. true - yes, false - no.
	OnlyUserRestriction *bool `json:"OnlyUserRestriction,omitnil,omitempty" name:"OnlyUserRestriction"`
}

func (r *DescribeBackupDownloadUserRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUserRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OnlyUserRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadUserRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadUserRestrictionResponseParams struct {
	// Cluster backup download limit information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupLimitClusterRestrictions []*BackupLimitClusterRestriction `json:"BackupLimitClusterRestrictions,omitnil,omitempty" name:"BackupLimitClusterRestrictions"`

	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadUserRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadUserRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadUserRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadUserRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListByVaultItem struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Backup information</p>
	BackupFileInfo *BackupFileInfo `json:"BackupFileInfo,omitnil,omitempty" name:"BackupFileInfo"`
}

// Predefined struct for user
type DescribeBackupListByVaultRequestParams struct {
	// Safe ID, length must be greater than 0
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup file ID list for filtering specific backups
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Cluster ID for filtering backups of a specific cluster
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup name list for exact matching filter criteria
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list for exact matching filter
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of pages, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Range: [0, INF). Default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeBackupListByVaultRequest struct {
	*tchttp.BaseRequest
	
	// Safe ID, length must be greater than 0
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Backup file ID list for filtering specific backups
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Cluster ID for filtering backups of a specific cluster
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup name list for exact matching filter criteria
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list for exact matching filter
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of pages, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Range: [0, INF). Default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeBackupListByVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListByVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "BackupIds")
	delete(f, "ClusterId")
	delete(f, "BackupNames")
	delete(f, "FileNames")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListByVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListByVaultResponseParams struct {
	// Total files of backup that meets the conditions
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Backup file list
	BackupList []*DescribeBackupListByVaultItem `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupListByVaultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupListByVaultResponseParams `json:"Response"`
}

func (r *DescribeBackupListByVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListByVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100].
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Backup ID.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Snapshot type. Optional values: full - Full snapshot; increment - Incremental snapshot.
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// Backup start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Backup file name, fuzzy search.
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup region.
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// Whether cross-region backup.
	IsCrossRegionsBackup *string `json:"IsCrossRegionsBackup,omitnil,omitempty" name:"IsCrossRegionsBackup"`

	// Status you want to query.
	BackupStatus []*string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`
}

type DescribeBackupListRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The number of results to be returned. Value range: (0,100].
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Value range: [0,INF).
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Backup ID.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup type. Valid values: `snapshot` (snapshot backup), `logic` (logic backup).
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Back mode. Valid values: `auto` (automatic backup), `manual` (manual backup)
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Snapshot type. Optional values: full - Full snapshot; increment - Incremental snapshot.
	SnapShotType *string `json:"SnapShotType,omitnil,omitempty" name:"SnapShotType"`

	// Backup start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Backup file name, fuzzy search.
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Backup alias, which supports fuzzy query.
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// ID list of the snapshot backup.
	SnapshotIdList []*int64 `json:"SnapshotIdList,omitnil,omitempty" name:"SnapshotIdList"`

	// Backup region.
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// Whether cross-region backup.
	IsCrossRegionsBackup *string `json:"IsCrossRegionsBackup,omitnil,omitempty" name:"IsCrossRegionsBackup"`

	// Status you want to query.
	BackupStatus []*string `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`
}

func (r *DescribeBackupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DbType")
	delete(f, "BackupIds")
	delete(f, "BackupType")
	delete(f, "BackupMethod")
	delete(f, "SnapShotType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FileNames")
	delete(f, "BackupNames")
	delete(f, "SnapshotIdList")
	delete(f, "BackupRegion")
	delete(f, "IsCrossRegionsBackup")
	delete(f, "BackupStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupListResponseParams struct {
	// Total number of backup files.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Backup file list.
	BackupList []*BackupFileInfo `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupListResponseParams `json:"Response"`
}

func (r *DescribeBackupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogConfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogConfigResponseParams struct {
	// Configuration update time for cross-regional Binlog.
	BinlogCrossRegionsConfigUpdateTime *string `json:"BinlogCrossRegionsConfigUpdateTime,omitnil,omitempty" name:"BinlogCrossRegionsConfigUpdateTime"`

	// Specifies the Binlog configuration message.
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogConfigResponseParams `json:"Response"`
}

func (r *DescribeBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// Backup download source restriction condition.
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

type DescribeBinlogDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog file ID
	BinlogId *int64 `json:"BinlogId,omitnil,omitempty" name:"BinlogId"`

	// Backup download source restriction condition.
	DownloadRestriction *BackupLimitRestriction `json:"DownloadRestriction,omitnil,omitempty" name:"DownloadRestriction"`
}

func (r *DescribeBinlogDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogId")
	delete(f, "DownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogDownloadUrlResponseParams struct {
	// Download address
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBinlogDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogListByVaultItem struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Binlog file information</p>
	BinlogFileInfo *BinlogItem `json:"BinlogFileInfo,omitnil,omitempty" name:"BinlogFileInfo"`
}

// Predefined struct for user
type DescribeBinlogListByVaultRequestParams struct {
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the backup IDs.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup name list
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of returned results, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0, INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeBinlogListByVaultRequest struct {
	*tchttp.BaseRequest
	
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the backup IDs.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Backup name list
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of returned results, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0, INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeBinlogListByVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogListByVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "ClusterId")
	delete(f, "BackupIds")
	delete(f, "BackupNames")
	delete(f, "FileNames")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogListByVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogListByVaultResponseParams struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Binlog file list
	BinlogList []*DescribeBinlogListByVaultItem `json:"BinlogList,omitnil,omitempty" name:"BinlogList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogListByVaultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogListByVaultResponseParams `json:"Response"`
}

func (r *DescribeBinlogListByVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogListByVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogSaveDaysResponseParams struct {
	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *DescribeBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on the number of records. The default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on the number of records. The default value is 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// Total number of records
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Binlog list.
	Binlogs []*BinlogItem `json:"Binlogs,omitnil,omitempty" name:"Binlogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogsResponseParams `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangedParamsAfterUpgradeRequestParams struct {
	// Cluster ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CPU after specification adjustment.
	DstCpu *int64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// Memory after specification adjustment, in GB.
	DstMem *int64 `json:"DstMem,omitnil,omitempty" name:"DstMem"`
}

type DescribeChangedParamsAfterUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CPU after specification adjustment.
	DstCpu *int64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// Memory after specification adjustment, in GB.
	DstMem *int64 `json:"DstMem,omitnil,omitempty" name:"DstMem"`
}

func (r *DescribeChangedParamsAfterUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangedParamsAfterUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstCpu")
	delete(f, "DstMem")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChangedParamsAfterUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangedParamsAfterUpgradeResponseParams struct {
	// Number of parameters.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance parameter list.
	Items []*ParamItemInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChangedParamsAfterUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChangedParamsAfterUpgradeResponseParams `json:"Response"`
}

func (r *DescribeChangedParamsAfterUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangedParamsAfterUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabaseTablesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the table type.
	// Specifies that "view" only returns the view, "base_table" only returns the basic table, and "all" returns both the view and the table. the default value is all.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type DescribeClusterDatabaseTablesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Specifies the table type.
	// Specifies that "view" only returns the view, "base_table" only returns the basic table, and "all" returns both the view and the table. the default value is all.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

func (r *DescribeClusterDatabaseTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabaseTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Db")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDatabaseTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabaseTablesResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Pagination Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database table list.
	Tables []*string `json:"Tables,omitnil,omitempty" name:"Tables"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDatabaseTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDatabaseTablesResponseParams `json:"Response"`
}

func (r *DescribeClusterDatabaseTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabaseTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabasesRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination number limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination number limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeClusterDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDatabasesResponseParams struct {
	// Total number of entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database list.
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Pagination number limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDatabasesResponseParams `json:"Response"`
}

func (r *DescribeClusterDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailDatabasesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DescribeClusterDetailDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of returned results. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

func (r *DescribeClusterDetailDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailDatabasesResponseParams struct {
	// Database information
	DbInfos []*DbInfo `json:"DbInfos,omitnil,omitempty" name:"DbInfos"`

	// The total count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDetailDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailDatabasesResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// Cluster details
	Detail *CynosdbClusterDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGrpsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterInstanceGrpsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInstanceGrpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstanceGrpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstanceGrpsResponseParams struct {
	// Number of instance groups
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance group list
	//
	// Deprecated: InstanceGrpInfoList is deprecated.
	InstanceGrpInfoList []*CynosdbInstanceGrp `json:"InstanceGrpInfoList,omitnil,omitempty" name:"InstanceGrpInfoList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInstanceGrpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstanceGrpsResponseParams `json:"Response"`
}

func (r *DescribeClusterInstanceGrpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstanceGrpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Parameter name.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether it is a global parameter.
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
}

type DescribeClusterParamsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Parameter name.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether it is a global parameter.
	IsGlobal *string `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`
}

func (r *DescribeClusterParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamName")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterParamsResponseParams struct {
	// Number of parameters.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance parameter list.
	Items []*ParamInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterParamsResponseParams `json:"Response"`
}

func (r *DescribeClusterParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPasswordComplexityResponseParams struct {
	// Data dictionary parameter.
	ValidatePasswordDictionary *ParamInfo `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`

	// Specifies the password length.
	ValidatePasswordLength *ParamInfo `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Case-Sensitive character count.
	ValidatePasswordMixedCaseCount *ParamInfo `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of digits.
	ValidatePasswordNumberCount *ParamInfo `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password level.
	ValidatePasswordPolicy *ParamInfo `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Number of special characters.
	ValidatePasswordSpecialCharCount *ParamInfo `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *DescribeClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReadOnlyRequestParams struct {
	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeClusterReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReadOnlyResponseParams struct {
	// List of cluster read-only switches.
	ClusterReadOnlyValues []*ClusterReadOnlyValue `json:"ClusterReadOnlyValues,omitnil,omitempty" name:"ClusterReadOnlyValues"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterReadOnlyResponseParams `json:"Response"`
}

func (r *DescribeClusterReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTransparentEncryptInfoRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterTransparentEncryptInfoRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterTransparentEncryptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTransparentEncryptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterTransparentEncryptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTransparentEncryptInfoResponseParams struct {
	// Encryption key id.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Encryption key region.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Key type
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// To be enabled global encryption
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterTransparentEncryptInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterTransparentEncryptInfoResponseParams `json:"Response"`
}

func (r *DescribeClusterTransparentEncryptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTransparentEncryptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "DbType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Number of clusters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Cluster list
	ClusterSet []*CynosdbCluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDBSecurityGroupsRequestParams struct {
	// Specifies that the instance ID must be provided by selecting either InstanceId or InstanceGroupId.
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance group ID, which can be obtained through the [DescribeClusterInstanceGroups](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1) API.
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Specifies that the instance ID must be provided by selecting either InstanceId or InstanceGroupId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance group ID, which can be obtained through the [DescribeClusterInstanceGroups](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1) API.
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group information
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowRequestParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// Task flow status. Valid values: `0` (succeeded), `1` (failed), `2` (Processing).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCLSLogDeliveryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type DescribeInstanceCLSLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

func (r *DescribeInstanceCLSLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCLSLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCLSLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCLSLogDeliveryResponseParams struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance shipping information.
	InstanceCLSDeliveryInfos []*InstanceCLSDeliveryInfo `json:"InstanceCLSDeliveryInfos,omitnil,omitempty" name:"InstanceCLSDeliveryInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceCLSLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCLSLogDeliveryResponseParams `json:"Response"`
}

func (r *DescribeInstanceCLSLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCLSLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailResponseParams struct {
	// Instance details
	Detail *CynosdbInstanceDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceErrorLogsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit on the number of logs.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the log number.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting field. Valid value: 'Timestamp'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC`, `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level, which supports combo search by multiple levels. Valid values: `error`, `warning`, `note`.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// Keywords, supports fuzzy search.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

type DescribeInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit on the number of logs.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the log number.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting field. Valid value: 'Timestamp'.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC`, `DESC`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Log level, which supports combo search by multiple levels. Valid values: `error`, `warning`, `note`.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// Keywords, supports fuzzy search.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

func (r *DescribeInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceErrorLogsResponseParams struct {
	// Number of logs.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Error log list.
	ErrorLogs []*CynosdbErrorLogItem `json:"ErrorLogs,omitnil,omitempty" name:"ErrorLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID, which supports batch query.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Search condition for a parameter name, which supports fuzzy search.
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID, which supports batch query.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Search condition for a parameter name, which supports fuzzy search.
	ParamKeyword *string `json:"ParamKeyword,omitnil,omitempty" name:"ParamKeyword"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ParamKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// List of instance parameters
	Items []*InstanceParamItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesRequestParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Earliest transaction start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Latest transaction start time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Limit.</p><p>It is recommended to control the limit size. Large limits may cause truncation due to the platform's size limit for return results.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Client host.</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>Database name.</p>
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// <p>Sorting field.</p><p>Enumeration values:</p><ul><li>QueryTime: sorts by the total execution time of the SQL statements.</li><li>LockTime: sorts by the time consumed by the SQL statements waiting for locks (such as table locks or row locks).</li><li>RowsExamined: sorts by the number of rows scanned during SQL statement execution.</li><li>RowsSent: sorts by the number of result rows returned to the client for the SQL statements.</li><li>Timestamp: sorts by the timestamp when the slow query statement occurs.</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type. Valid values: asc and desc.</p>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>SQL statement.</p>
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`
}

type DescribeInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Earliest transaction start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Latest transaction start time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Limit.</p><p>It is recommended to control the limit size. Large limits may cause truncation due to the platform's size limit for return results.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Offset.</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Username.</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Client host.</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>Database name.</p>
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// <p>Sorting field.</p><p>Enumeration values:</p><ul><li>QueryTime: sorts by the total execution time of the SQL statements.</li><li>LockTime: sorts by the time consumed by the SQL statements waiting for locks (such as table locks or row locks).</li><li>RowsExamined: sorts by the number of rows scanned during SQL statement execution.</li><li>RowsSent: sorts by the number of result rows returned to the client for the SQL statements.</li><li>Timestamp: sorts by the timestamp when the slow query statement occurs.</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type. Valid values: asc and desc.</p>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>SQL statement.</p>
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`
}

func (r *DescribeInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "SqlText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSlowQueriesResponseParams struct {
	// <p>Total quantity.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Slow query records.</p>
	SlowQueries []*SlowQueriesItem `json:"SlowQueries,omitnil,omitempty" name:"SlowQueries"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsByOperationTypeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Operation type.
	// Add new RO instance
	// modifyInstance: Modify configuration
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Instance ID, required when querying configuration modification specifications
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance Machine Type
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type DescribeInstanceSpecsByOperationTypeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Operation type.
	// Add new RO instance
	// modifyInstance: Modify configuration
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Instance ID, required when querying configuration modification specifications
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance Machine Type
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

func (r *DescribeInstanceSpecsByOperationTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsByOperationTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OperationType")
	delete(f, "InstanceId")
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsByOperationTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsByOperationTypeResponseParams struct {
	// instance specifications
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitnil,omitempty" name:"InstanceSpecSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecsByOperationTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecsByOperationTypeResponseParams `json:"Response"`
}

func (r *DescribeInstanceSpecsByOperationTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsByOperationTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsRequestParams struct {
	// <p>Database type, value ranges from...to... </p><li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Whether required to return AZ information</p>
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`

	// <p>Instance machine type.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Cluster level, optional. For example P0, P1</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

type DescribeInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Database type, value ranges from...to... </p><li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// <p>Whether required to return AZ information</p>
	IncludeZoneStocks *bool `json:"IncludeZoneStocks,omitnil,omitempty" name:"IncludeZoneStocks"`

	// <p>Instance machine type.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Cluster level, optional. For example P0, P1</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

func (r *DescribeInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbType")
	delete(f, "IncludeZoneStocks")
	delete(f, "DeviceType")
	delete(f, "ClusterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecsResponseParams struct {
	// <p>Specification information.</p>
	InstanceSpecSet []*InstanceSpec `json:"InstanceSpecSet,omitnil,omitempty" name:"InstanceSpecSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecsResponseParams `json:"Response"`
}

func (r *DescribeInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Cluster type, value ranges from <li>CYNOSDB: transactional cluster</li><li>LIBRADB: analysis cluster</li><li>ALL: all</li>, default to ALL.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort by field. Valid values:
	// <li> CREATETIME: creation time</li>
	// <li> PERIODENDTIME: expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> ASC: ascending</li>
	// <li> DESC: descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter. If more than one filter exists, the logical relationship between these filters is `AND`.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type. Currently, `MYSQL` is supported.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance status. Valid values:
	// creating
	// running
	// isolating
	// isolated
	// activating: Removing the instance from isolation
	// offlining: Eliminating the instance
	// offlined: Instance eliminated
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Cluster type, value ranges from <li>CYNOSDB: transactional cluster</li><li>LIBRADB: analysis cluster</li><li>ALL: all</li>, default to ALL.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	delete(f, "DbType")
	delete(f, "Status")
	delete(f, "InstanceIds")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

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
type DescribeInstancesWithinSameClusterRequestParams struct {
	// vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DescribeInstancesWithinSameClusterRequest struct {
	*tchttp.BaseRequest
	
	// vpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DescribeInstancesWithinSameClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesWithinSameClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesWithinSameClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesWithinSameClusterResponseParams struct {
	// Number of instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesWithinSameClusterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesWithinSameClusterResponseParams `json:"Response"`
}

func (r *DescribeInstancesWithinSameClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesWithinSameClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrateTaskRequestParams struct {
	// Large order id. must choose between large order id and sub-order id.
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// Order list
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeIntegrateTaskRequest struct {
	*tchttp.BaseRequest
	
	// Large order id. must choose between large order id and sub-order id.
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// Order list
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeIntegrateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BigDealId")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrateTaskResponseParams struct {
	// Current step.
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// Current progress.
	CurrentProgress *string `json:"CurrentProgress,omitnil,omitempty" name:"CurrentProgress"`

	// Indicates the task status.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntegrateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrateTaskResponseParams `json:"Response"`
}

func (r *DescribeIntegrateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIsolatedInstancesRequestParams struct {
	// Number of returned results. the default value is 20. the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. valid values:.
	// <Li>CREATETIME: creation time</li>.
	// <li> PERIODENDTIME: expiration time</li>.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. value range:.
	// <Li>ASC: specifies ascending sort.</li>.
	// <li> DESC: sorts in descending order. </li>.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Search criteria. when multiple filters exist, the relationship between filters is logical AND.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type: currently supports "MYSQL", "POSTGRESQL".
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

type DescribeIsolatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. the default value is 20. the maximum value is 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. valid values:.
	// <Li>CREATETIME: creation time</li>.
	// <li> PERIODENDTIME: expiration time</li>.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. value range:.
	// <Li>ASC: specifies ascending sort.</li>.
	// <li> DESC: sorts in descending order. </li>.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Search criteria. when multiple filters exist, the relationship between filters is logical AND.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Engine type: currently supports "MYSQL", "POSTGRESQL".
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`
}

func (r *DescribeIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIsolatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	delete(f, "DbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIsolatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIsolatedInstancesResponseParams struct {
	// Number of instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list
	InstanceSet []*CynosdbInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIsolatedInstancesResponseParams `json:"Response"`
}

func (r *DescribeIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIsolatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountAllPrivilegesRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type DescribeLibraDBClusterAccountAllPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

func (r *DescribeLibraDBClusterAccountAllPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountAllPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterAccountAllPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountAllPrivilegesResponseParams struct {
	// Permission statement
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivilegeStatements []*string `json:"PrivilegeStatements,omitnil,omitempty" name:"PrivilegeStatements"`

	// Global permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterAccountAllPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterAccountAllPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterAccountAllPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountAllPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountPrivilegesRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// host name
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DescribeLibraDBClusterAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// host name
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

func (r *DescribeLibraDBClusterAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "Host")
	delete(f, "Db")
	delete(f, "Type")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountPrivilegesResponseParams struct {
	// Permission list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountsRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Fuzzy matching keyword
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`

	// host name
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLibraDBClusterAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountNames []*string `json:"AccountNames,omitnil,omitempty" name:"AccountNames"`

	// Fuzzy matching keyword
	AccountRegular *string `json:"AccountRegular,omitnil,omitempty" name:"AccountRegular"`

	// host name
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLibraDBClusterAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountNames")
	delete(f, "AccountRegular")
	delete(f, "Hosts")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAccountsResponseParams struct {
	// Account information
	AccountSet []*Account `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterAccountsResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAutoMapRuleRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Analyze instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLibraDBClusterAutoMapRuleRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Analyze instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLibraDBClusterAutoMapRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAutoMapRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterAutoMapRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterAutoMapRuleResponseParams struct {
	// Advanced mapping rule
	AutoMapRules []*AutoMapRule `json:"AutoMapRules,omitnil,omitempty" name:"AutoMapRules"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterAutoMapRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterAutoMapRuleResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterAutoMapRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterAutoMapRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterDetailRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to get more server information, available values yes no
	GetServerInfo *string `json:"GetServerInfo,omitnil,omitempty" name:"GetServerInfo"`
}

type DescribeLibraDBClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to get more server information, available values yes no
	GetServerInfo *string `json:"GetServerInfo,omitnil,omitempty" name:"GetServerInfo"`
}

func (r *DescribeLibraDBClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GetServerInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterDetailResponseParams struct {
	// Cluster information.
	Detail *LibraDBClusterDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterTableMappingRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page history limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Source schema list
	SrcSchemas []*string `json:"SrcSchemas,omitnil,omitempty" name:"SrcSchemas"`

	// Source table list
	SrcTableName []*string `json:"SrcTableName,omitnil,omitempty" name:"SrcTableName"`

	// Status list
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// Map database name
	MapSchemas []*string `json:"MapSchemas,omitnil,omitempty" name:"MapSchemas"`

	// Mapping table name
	MapTableName []*string `json:"MapTableName,omitnil,omitempty" name:"MapTableName"`

	// Query records where the database name is not null
	MapSchemaNotEmpty *bool `json:"MapSchemaNotEmpty,omitnil,omitempty" name:"MapSchemaNotEmpty"`

	// Query records where the mapping table name is not null
	MapTableNameNotEmpty *bool `json:"MapTableNameNotEmpty,omitnil,omitempty" name:"MapTableNameNotEmpty"`
}

type DescribeLibraDBClusterTableMappingRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page history limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Source schema list
	SrcSchemas []*string `json:"SrcSchemas,omitnil,omitempty" name:"SrcSchemas"`

	// Source table list
	SrcTableName []*string `json:"SrcTableName,omitnil,omitempty" name:"SrcTableName"`

	// Status list
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// Map database name
	MapSchemas []*string `json:"MapSchemas,omitnil,omitempty" name:"MapSchemas"`

	// Mapping table name
	MapTableName []*string `json:"MapTableName,omitnil,omitempty" name:"MapTableName"`

	// Query records where the database name is not null
	MapSchemaNotEmpty *bool `json:"MapSchemaNotEmpty,omitnil,omitempty" name:"MapSchemaNotEmpty"`

	// Query records where the mapping table name is not null
	MapTableNameNotEmpty *bool `json:"MapTableNameNotEmpty,omitnil,omitempty" name:"MapTableNameNotEmpty"`
}

func (r *DescribeLibraDBClusterTableMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterTableMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "NodeId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SrcSchemas")
	delete(f, "SrcTableName")
	delete(f, "StatusList")
	delete(f, "MapSchemas")
	delete(f, "MapTableName")
	delete(f, "MapSchemaNotEmpty")
	delete(f, "MapTableNameNotEmpty")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClusterTableMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClusterTableMappingResponseParams struct {
	// Total number of records
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// Database mapping info
	TableMappings []*TableMappingObject `json:"TableMappings,omitnil,omitempty" name:"TableMappings"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClusterTableMappingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClusterTableMappingResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClusterTableMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClusterTableMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClustersRequestParams struct {
	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sort method
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter criteria.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLibraDBClustersRequest struct {
	*tchttp.BaseRequest
	
	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sort method
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter criteria.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLibraDBClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBClustersResponseParams struct {
	// Cluster information.
	ClusterSet []*LibraDBClusterSet `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// Number of clusters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBClustersResponseParams `json:"Response"`
}

func (r *DescribeLibraDBClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBDataSourceRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLibraDBDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLibraDBDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBDataSourceResponseParams struct {
	// Source information list
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataSourceList []*DataSourceItem `json:"DataSourceList,omitnil,omitempty" name:"DataSourceList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBDataSourceResponseParams `json:"Response"`
}

func (r *DescribeLibraDBDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBForwardConfigRequestParams struct {
	// Read-only analysis engine instance id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLibraDBForwardConfigRequest struct {
	*tchttp.BaseRequest
	
	// Read-only analysis engine instance id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLibraDBForwardConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBForwardConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBForwardConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBForwardConfigResponseParams struct {
	// Whether to enable forwarding
	ForwardMode *string `json:"ForwardMode,omitnil,omitempty" name:"ForwardMode"`

	// Forwarding list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ForwardList []*ForwardInstanceInfo `json:"ForwardList,omitnil,omitempty" name:"ForwardList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBForwardConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBForwardConfigResponseParams `json:"Response"`
}

func (r *DescribeLibraDBForwardConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBForwardConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBInstanceDetailRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Read-only analysis engine instance ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLibraDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Read-only analysis engine instance ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLibraDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBInstanceDetailResponseParams struct {
	// <p>root account</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>Account unique ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Cluster ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Instance ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance name.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Project ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Region</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>AZ.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance status.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Libra analysis engine version</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// <p>cpu cores</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Memory size</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Storage size</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>Storage type</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>Instance type</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Instance role</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Update time.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Creation time.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Selling mode</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Start time</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodStartTime *string `json:"PeriodStartTime,omitnil,omitempty" name:"PeriodStartTime"`

	// <p>End time of sale</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Renewal flag</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Network type</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>VPC ID</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet ID.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Virtual IP</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Port</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Instance network information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNetInfo *InstanceNetInfo `json:"InstanceNetInfo,omitnil,omitempty" name:"InstanceNetInfo"`

	// <p>Instance tag information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Instance node information</p>
	NodeInfo []*LibraDBNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// <p>Number of instance nodes</p>
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// <p>Analyze the information after instance upgrade version</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisUpgradeVersionInfo *UpgradeAnalysisInstanceVersionInfo `json:"AnalysisUpgradeVersionInfo,omitnil,omitempty" name:"AnalysisUpgradeVersionInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeLibraDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBInstanceSpecsRequestParams struct {

}

type DescribeLibraDBInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLibraDBInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBInstanceSpecsResponseParams struct {
	// Instance specifications list in this region
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceSpecSet []*RegionInstanceSpecInfo `json:"InstanceSpecSet,omitnil,omitempty" name:"InstanceSpecSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBInstanceSpecsResponseParams `json:"Response"`
}

func (r *DescribeLibraDBInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBSlowLogsRequestParams struct {
	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, 1753171200.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, 1753171200.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Log items per page limit: 0-200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log paginate, more than 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log sorting method, DESC - descending order, ASC - ascending order.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Log sorting condition.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter criteria
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type DescribeLibraDBSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time, 1753171200.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time, 1753171200.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Log items per page limit: 0-200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log paginate, more than 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log sorting method, DESC - descending order, ASC - ascending order.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Log sorting condition.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter criteria
	LogFilter []*LogFilter `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

func (r *DescribeLibraDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "LogFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBSlowLogsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeLibraDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBVersionRequestParams struct {

}

type DescribeLibraDBVersionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLibraDBVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibraDBVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibraDBVersionResponseParams struct {
	// Version list
	VersionList []*LibraDBVersion `json:"VersionList,omitnil,omitempty" name:"VersionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibraDBVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibraDBVersionResponseParams `json:"Response"`
}

func (r *DescribeLibraDBVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibraDBVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainPeriodRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintainPeriodRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintainPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintainPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintainPeriodResponseParams struct {
	// Maintenance days of the week
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// Maintenance start time in seconds
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaintainPeriodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintainPeriodResponseParams `json:"Response"`
}

func (r *DescribeMaintainPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintainPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateDetailRequestParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateDetailRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateDetailResponseParams struct {
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Engine version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Total number of parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of parameters
	Items []*ParamDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// Database type. Valid values:  `NORMAL`, `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateDetailResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Database engine version number
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Template name
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbModes []*string `json:"DbModes,omitnil,omitempty" name:"DbModes"`

	// Offset for query
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on queries
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Product type of the queried template
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// Template type
	TemplateTypes []*string `json:"TemplateTypes,omitnil,omitempty" name:"TemplateTypes"`

	// Version type
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// The sorting order of the returned results
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc `.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersions")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	delete(f, "DbModes")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Products")
	delete(f, "TemplateTypes")
	delete(f, "EngineTypes")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// Number of parameter templates
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter template information
	Items []*ParamTemplateListInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Maximum entries returned per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Search by keyword
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The total number of groups
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxiesRequestParams struct {
	// <p>Cluster ID (this parameter is required, such as cynosdbmysql-2u2mh111).</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Number of returned results, defaults to 20 with a maximum value of 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Record offset. Default value is 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field. Value ranges from...to...:</p><li>CREATETIME: Creation time</li><li>PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type. Value ranges from:</p><li> ASC: ascending order </li><li> DESC: sort in descending order </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>Search criteria. If there are multiple Filters, the relationship between Filters is logical AND. <br>Description: This parameter currently only supports two filter conditions: Status and ProxyGroupId.</p>
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxiesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID (this parameter is required, such as cynosdbmysql-2u2mh111).</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Number of returned results, defaults to 20 with a maximum value of 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Record offset. Default value is 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Sorting field. Value ranges from...to...:</p><li>CREATETIME: Creation time</li><li>PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting type. Value ranges from:</p><li> ASC: ascending order </li><li> DESC: sort in descending order </li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// <p>Search criteria. If there are multiple Filters, the relationship between Filters is logical AND. <br>Description: This parameter currently only supports two filter conditions: Status and ProxyGroupId.</p>
	Filters []*QueryParamFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxiesResponseParams struct {
	// <p>Number of Database Proxy Groups</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Database Proxy Group list</p>
	ProxyGroupInfos []*ProxyGroupInfo `json:"ProxyGroupInfos,omitnil,omitempty" name:"ProxyGroupInfos"`

	// <p>database proxy node</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// <p>sql automatic forwarding</p>
	ColumnStoreProxyForward *string `json:"ColumnStoreProxyForward,omitnil,omitempty" name:"ColumnStoreProxyForward"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxiesResponseParams `json:"Response"`
}

func (r *DescribeProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesRequestParams struct {
	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Search criteria. When multiple filters exist, the relationship between filters is logical AND. Currently supported search fields: Status, ProxyNodeId, ClusterId, OssProxyNodeName.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProxyNodesRequest struct {
	*tchttp.BaseRequest
	
	// Number of returned results. Default value: `20`. Maximum value: `100`,
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values:
	// <li> CREATETIME: Creation time</li>
	// <li> PERIODENDTIME: Expiration time</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values:
	// <li> `ASC`: Ascending.</li>
	// <li> `DESC`: Descending</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Search criteria. When multiple filters exist, the relationship between filters is logical AND. Currently supported search fields: Status, ProxyNodeId, ClusterId, OssProxyNodeName.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProxyNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyNodesResponseParams struct {
	// Number of the database proxy nodes
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of the database proxy nodes
	ProxyNodeInfos []*ProxyNodeInfo `json:"ProxyNodeInfos,omitnil,omitempty" name:"ProxyNodeInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxyNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyNodesResponseParams `json:"Response"`
}

func (r *DescribeProxyNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeProxySpecsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeProxySpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySpecsResponseParams struct {
	// List of database proxyspecifications
	ProxySpecs []*ProxySpec `json:"ProxySpecs,omitnil,omitempty" name:"ProxySpecs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySpecsResponseParams `json:"Response"`
}

func (r *DescribeProxySpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRedoLogListByVaultItem struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>RedoLog file information</p>
	RedoFileInfo *RedoLogItem `json:"RedoFileInfo,omitnil,omitempty" name:"RedoFileInfo"`
}

// Predefined struct for user
type DescribeRedoLogListByVaultRequestParams struct {
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// List of the backup IDs.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup name list
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of items per page, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeRedoLogListByVaultRequest struct {
	*tchttp.BaseRequest
	
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// List of the backup IDs.
	BackupIds []*int64 `json:"BackupIds,omitnil,omitempty" name:"BackupIds"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup name list
	BackupNames []*string `json:"BackupNames,omitnil,omitempty" name:"BackupNames"`

	// Filename list
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// Number of items per page, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,INF). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime. Default: createTime.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC, default desc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeRedoLogListByVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedoLogListByVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "BackupIds")
	delete(f, "ClusterId")
	delete(f, "BackupNames")
	delete(f, "FileNames")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedoLogListByVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedoLogListByVaultResponseParams struct {
	// Total eligible redo log files
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Redo log file list
	RedoLogList []*DescribeRedoLogListByVaultItem `json:"RedoLogList,omitnil,omitempty" name:"RedoLogList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedoLogListByVaultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedoLogListByVaultResponseParams `json:"Response"`
}

func (r *DescribeRedoLogListByVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedoLogListByVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedoLogsRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Entries Per Page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Log file name.
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`
}

type DescribeRedoLogsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Entries Per Page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Log file name.
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`
}

func (r *DescribeRedoLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedoLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FileNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedoLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedoLogsResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// redo log information.
	RedoLogs []*RedoLogItem `json:"RedoLogs,omitnil,omitempty" name:"RedoLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedoLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedoLogsResponseParams `json:"Response"`
}

func (r *DescribeRedoLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedoLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeResourcePackageDetailRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance D
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeResourcePackageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "ClusterIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageDetailResponseParams struct {
	// Total number of deducted resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Resource pack details
	Detail []*PackageDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageDetailResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListRequestParams struct {
	// The unique ID of a resource pack
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack name
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack status. Valid values: `using`, `expired`, `normal_finish` (used up), `apply_refund` (requesting a refund), `refund` (refunded).
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting conditions. Valid values: `startTime` (effective time), `expireTime` (expiration date), `packageUsedSpec` (used capacity), `packageTotalSpec` (total storage capacity). 
	// Sorting by array order
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending).
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageListRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId []*string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource pack name
	PackageName []*string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Resource pack type. Valid values: `CCU` (compute resource pack), `DISK` (storage resource pack).
	PackageType []*string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion []*string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack status. Valid values: `using`, `expired`, `normal_finish` (used up), `apply_refund` (requesting a refund), `refund` (refunded).
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting conditions. Valid values: `startTime` (effective time), `expireTime` (expiration date), `packageUsedSpec` (used capacity), `packageTotalSpec` (total storage capacity). 
	// Sorting by array order
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending).
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	delete(f, "PackageType")
	delete(f, "PackageRegion")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageListResponseParams struct {
	// Total number of resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Resource pack details.
	Detail []*Package `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageListResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecRequestParams struct {
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack, `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeResourcePackageSaleSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Region of the resource pack. Valid values: `China` (Chinese mainland), `overseas` (outside Chinese mainland).
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Resource pack type. Valid values: `CCU` (compute resource pack, `DISK` (storage resource pack).
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeResourcePackageSaleSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "PackageRegion")
	delete(f, "PackageType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePackageSaleSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePackageSaleSpecResponseParams struct {
	// Total number of available resource packs
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Specifies the details of the resource pack.
	Detail []*SalePackageSpec `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePackageSaleSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePackageSaleSpecResponseParams `json:"Response"`
}

func (r *DescribeResourcePackageSaleSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePackageSaleSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameRequestParams struct {
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeResourcesByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// Order ID. (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Order ID, which can be used to query the resource information of multiple orders ID (If the cluster is not delivered yet, the `DescribeResourcesByDealName` API may return the `InvalidParameterValue.DealNameNotFound` error. Call the API again until it succeeds.)
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeResourcesByDealNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealName")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByDealNameResponseParams struct {
	// Billable resource ID information array
	BillingResourceInfos []*BillingResourceInfo `json:"BillingResourceInfos,omitnil,omitempty" name:"BillingResourceInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcesByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByDealNameResponseParams `json:"Response"`
}

func (r *DescribeResourcesByDealNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByDealNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRangeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Safe id
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe region
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`
}

type DescribeRollbackTimeRangeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Safe id
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe region
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`
}

func (r *DescribeRollbackTimeRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VaultId")
	delete(f, "VaultRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRangeResponseParams struct {
	// Valid regression time range start time point (abandoned).
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// Valid regression time range end time point (abandoned).
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// Time range available for rollback
	RollbackTimeRanges []*RollbackTimeRange `json:"RollbackTimeRanges,omitnil,omitempty" name:"RollbackTimeRanges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeRangeResponseParams `json:"Response"`
}

func (r *DescribeRollbackTimeRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSQLExecutionPlanRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SQL Template ID</p>
	TemplateID *string `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// <p>Plan details serial number</p>
	PlanDetailId *int64 `json:"PlanDetailId,omitnil,omitempty" name:"PlanDetailId"`
}

type DescribeSQLExecutionPlanRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>SQL Template ID</p>
	TemplateID *string `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// <p>Plan details serial number</p>
	PlanDetailId *int64 `json:"PlanDetailId,omitnil,omitempty" name:"PlanDetailId"`
}

func (r *DescribeSQLExecutionPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSQLExecutionPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "TemplateID")
	delete(f, "PlanDetailId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSQLExecutionPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSQLExecutionPlanResponseParams struct {
	// <p>Execution plan details</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	PlanDetail *ExecutionPlanDetail `json:"PlanDetail,omitnil,omitempty" name:"PlanDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSQLExecutionPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSQLExecutionPlanResponseParams `json:"Response"`
}

func (r *DescribeSQLExecutionPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSQLExecutionPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaveBackupClustersRequestParams struct {
	// Entries Per Page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Retrieval criteria.
	Filters []*QuerySimpleFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSaveBackupClustersRequest struct {
	*tchttp.BaseRequest
	
	// Entries Per Page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Retrieval criteria.
	Filters []*QuerySimpleFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSaveBackupClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaveBackupClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSaveBackupClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSaveBackupClustersResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Legacy backup information.
	SaveBackupClusterInfos []*SaveBackupClusterInfo `json:"SaveBackupClusterInfos,omitnil,omitempty" name:"SaveBackupClusterInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSaveBackupClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSaveBackupClustersResponseParams `json:"Response"`
}

func (r *DescribeSaveBackupClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSaveBackupClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstanceSpecsRequestParams struct {
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Cluster level</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

type DescribeServerlessInstanceSpecsRequest struct {
	*tchttp.BaseRequest
	
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Cluster level</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

func (r *DescribeServerlessInstanceSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstanceSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ClusterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessInstanceSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstanceSpecsResponseParams struct {
	// <p>Serverless instance available specifications</p>
	Specs []*ServerlessSpec `json:"Specs,omitnil,omitempty" name:"Specs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessInstanceSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessInstanceSpecsResponseParams `json:"Response"`
}

func (r *DescribeServerlessInstanceSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstanceSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessStrategyRequestParams struct {
	// Specifies the serverless cluster id.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeServerlessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the serverless cluster id.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeServerlessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessStrategyResponseParams struct {
	// Specifies how long (in seconds) the auto-pause is triggered when the cpu load is 0.
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// Specifies how long (in seconds) the auto-scaling is triggered when the cpu load exceeds the number of cores in the current specifications.
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// Specifies how long (in seconds) the system will wait for the cpu load to be lower than the number of cores in the lower specification before triggering automatic scaling down.
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// Whether to automatically pause. valid values:.
	// yes
	// no
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// Specifies whether the cluster allows upward scaling. valid values: <li>yes</li><li>no</li>.
	AutoScaleUp *string `json:"AutoScaleUp,omitnil,omitempty" name:"AutoScaleUp"`

	// Whether the cluster is allowed to scale down. valid values: <li>yes</li><li>no</li>.
	AutoScaleDown *string `json:"AutoScaleDown,omitnil,omitempty" name:"AutoScaleDown"`

	// Whether archiving is enabled. Optional range</p><li>yes</li><li>no</li>Default value: yes</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessStrategyResponseParams `json:"Response"`
}

func (r *DescribeServerlessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlaveZonesRequestParams struct {
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cloud frame cluster ID.
	OssClusterId *int64 `json:"OssClusterId,omitnil,omitempty" name:"OssClusterId"`

	// Storage architecture type. Enumeration value: 1.0/2.0. Default value: 1.0.
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

type DescribeSlaveZonesRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cloud frame cluster ID.
	OssClusterId *int64 `json:"OssClusterId,omitnil,omitempty" name:"OssClusterId"`

	// Storage architecture type. Enumeration value: 1.0/2.0. Default value: 1.0.
	StorageVersion *string `json:"StorageVersion,omitnil,omitempty" name:"StorageVersion"`
}

func (r *DescribeSlaveZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlaveZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "OssClusterId")
	delete(f, "StorageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlaveZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlaveZonesResponseParams struct {
	// Secondary AZ.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlaveZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlaveZonesResponseParams `json:"Response"`
}

func (r *DescribeSlaveZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlaveZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeSupportProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeSupportProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportProxyVersionResponseParams struct {
	// Supported database proxy version collections.
	SupportProxyVersions []*string `json:"SupportProxyVersions,omitnil,omitempty" name:"SupportProxyVersions"`

	// The current proxy version number.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// Specifies the proxy version details.
	SupportProxyVersionDetail []*ProxyVersionInfo `json:"SupportProxyVersionDetail,omitnil,omitempty" name:"SupportProxyVersionDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportProxyVersionResponseParams `json:"Response"`
}

func (r *DescribeSupportProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// Starting value of the task start time.
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// End value of the task start time.
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// Filtering Conditions. Supported search fields: "ClusterId", "ClusterName", "InstanceId", "InstanceName", "Status", "TaskId", "TaskType".
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Length of the list to be queried.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the list to be queried.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Starting value of the task start time.
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// End value of the task start time.
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// Filtering Conditions. Supported search fields: "ClusterId", "ClusterName", "InstanceId", "InstanceName", "Status", "TaskId", "TaskType".
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Length of the list to be queried.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset of the list to be queried.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeBegin")
	delete(f, "StartTimeEnd")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// Total number of entries in the task list.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Task list.
	TaskList []*BizTaskInfo `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVaultBackupClusterInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Cluster status
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// cluster region
	ClusterRegion *string `json:"ClusterRegion,omitnil,omitempty" name:"ClusterRegion"`

	// AZ of the cluster
	ClusterZone *string `json:"ClusterZone,omitnil,omitempty" name:"ClusterZone"`
}

// Predefined struct for user
type DescribeVaultBackupClusterInfoRequestParams struct {
	// Backup safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`
}

type DescribeVaultBackupClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// Backup safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`
}

func (r *DescribeVaultBackupClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVaultBackupClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVaultBackupClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVaultBackupClusterInfoResponseParams struct {
	// Safe info
	ClusterInfos []*DescribeVaultBackupClusterInfo `json:"ClusterInfos,omitnil,omitempty" name:"ClusterInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVaultBackupClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVaultBackupClusterInfoResponseParams `json:"Response"`
}

func (r *DescribeVaultBackupClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVaultBackupClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVaultsItem struct {
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe name
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Safe description
	VaultDescribe *string `json:"VaultDescribe,omitnil,omitempty" name:"VaultDescribe"`

	// Encryption key ID.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key region
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Key type
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Number of backup files
	BackupFileCount *int64 `json:"BackupFileCount,omitnil,omitempty" name:"BackupFileCount"`

	// Total size of backup files (byte)
	BackupFileSize *int64 `json:"BackupFileSize,omitnil,omitempty" name:"BackupFileSize"`

	// Binlog file count
	BinlogFileCount *int64 `json:"BinlogFileCount,omitnil,omitempty" name:"BinlogFileCount"`

	// Total size of the Binlog file (byte)
	BinlogFileSize *int64 `json:"BinlogFileSize,omitnil,omitempty" name:"BinlogFileSize"`

	// Number of RedoLog files
	RedoLogFileCount *int64 `json:"RedoLogFileCount,omitnil,omitempty" name:"RedoLogFileCount"`

	// Total size of RedoLog files (byte)
	RedoLogFileSize *int64 `json:"RedoLogFileSize,omitnil,omitempty" name:"RedoLogFileSize"`

	// Safe status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Backup retention period (seconds)
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// lock time
	LockedTime *string `json:"LockedTime,omitnil,omitempty" name:"LockedTime"`

	// Associated task list
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Safe region
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`

	// Automatically deliver relationship
	AutoCopyConfigs []*AutoCopyConfig `json:"AutoCopyConfigs,omitnil,omitempty" name:"AutoCopyConfigs"`
}

// Predefined struct for user
type DescribeVaultsRequestParams struct {
	// Safe ID list for precise filtering
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`

	// Safe name, used for fuzzy filtering
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Status list of safe for filtering
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of items per page, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,+∞). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeVaultsRequest struct {
	*tchttp.BaseRequest
	
	// Safe ID list for precise filtering
	VaultIds []*string `json:"VaultIds,omitnil,omitempty" name:"VaultIds"`

	// Safe name, used for fuzzy filtering
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Status list of safe for filtering
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of items per page, range (0,100], default 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Range: [0,+∞). Default is 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Available values: VaultId, VaultName, BackupSaveSeconds, LockedTime, CreateTime, UpdateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, available values: desc, asc, DESC, ASC
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeVaultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVaultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultIds")
	delete(f, "VaultName")
	delete(f, "Status")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVaultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVaultsResponseParams struct {
	// Safe list
	Vaults []*DescribeVaultsItem `json:"Vaults,omitnil,omitempty" name:"Vaults"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVaultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVaultsResponseParams `json:"Response"`
}

func (r *DescribeVaultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVaultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// Whether the virtual zone is included.
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user's permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Whether the virtual zone is included.
	IncludeVirtualZones *bool `json:"IncludeVirtualZones,omitnil,omitempty" name:"IncludeVirtualZones"`

	// Whether to display all AZs in a region and the user's permissions in each AZ.
	ShowPermission *bool `json:"ShowPermission,omitnil,omitempty" name:"ShowPermission"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IncludeVirtualZones")
	delete(f, "ShowPermission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Region information
	RegionSet []*SaleRegion `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Array of instance group IDs, starting with the prefix cynosdbmysql-grp- or cluster ID.
	// Description: To get the instance group ID of a cluster, perform [query cluster instance group](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Security group ID list to modify. An array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone.
	// Description: Please enter the primary AZ of the cluster location correctly. If you enter a non-primary AZ of the cluster location, the call may display success but the actual execution will fail.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance group IDs, starting with the prefix cynosdbmysql-grp- or cluster ID.
	// Description: To get the instance group ID of a cluster, perform [query cluster instance group](https://www.tencentcloud.com/document/product/1003/103934?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Security group ID list to modify. An array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone.
	// Description: Please enter the primary AZ of the cluster location correctly. If you enter a non-primary AZ of the cluster location, the call may display success but the actual execution will fail.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadLibraDBClusterListRequestParams struct {
	// Limit Quantity
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, desc, asc, DESC, ASC
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter criteria.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DownloadLibraDBClusterListRequest struct {
	*tchttp.BaseRequest
	
	// Limit Quantity
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method, desc, asc, DESC, ASC
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Filter criteria.
	Filters []*QueryFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DownloadLibraDBClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadLibraDBClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadLibraDBClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadLibraDBClusterListResponseParams struct {
	// Analyze cluster information
	ClusterSet []*LibraClusterSet `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadLibraDBClusterListResponse struct {
	*tchttp.BaseResponse
	Response *DownloadLibraDBClusterListResponseParams `json:"Response"`
}

func (r *DownloadLibraDBClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadLibraDBClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorLogItemExport struct {
	// Time
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Specifies the log level. valid values are note, warning, and error.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Log content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ExchangeInstanceInfo struct {
	// Source instance information.
	SrcInstanceInfo *RollbackInstanceInfo `json:"SrcInstanceInfo,omitnil,omitempty" name:"SrcInstanceInfo"`

	// Target instance information.
	DstInstanceInfo *RollbackInstanceInfo `json:"DstInstanceInfo,omitnil,omitempty" name:"DstInstanceInfo"`
}

type ExchangeRoGroupInfo struct {
	// Source RO group information.
	SrcRoGroupInfo *RollbackRoGroupInfo `json:"SrcRoGroupInfo,omitnil,omitempty" name:"SrcRoGroupInfo"`

	// Target RO group information.
	DstRoGroupInfo *RollbackRoGroupInfo `json:"DstRoGroupInfo,omitnil,omitempty" name:"DstRoGroupInfo"`
}

type ExecutionPlanDetail struct {
	// <p>Template ID</p>
	TemplateID *string `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// <p>Database name.</p>
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// <p>Original SQL sample</p>
	SQLSample *string `json:"SQLSample,omitnil,omitempty" name:"SQLSample"`

	// <p>Sample SQL after rewrite</p>
	SQLSampleRewritten *string `json:"SQLSampleRewritten,omitnil,omitempty" name:"SQLSampleRewritten"`

	// <p>Execution plan before optimization - List</p>
	TablePlanBefore []*ExplainRow `json:"TablePlanBefore,omitnil,omitempty" name:"TablePlanBefore"`

	// <p>Execution plan after optimization - List</p>
	TablePlanAfter []*ExplainRow `json:"TablePlanAfter,omitnil,omitempty" name:"TablePlanAfter"`

	// <p>Tree execution plan before optimization</p>
	TreePlanBefore *string `json:"TreePlanBefore,omitnil,omitempty" name:"TreePlanBefore"`

	// <p>Tree execution plan after optimization</p>
	TreePlanAfter *string `json:"TreePlanAfter,omitnil,omitempty" name:"TreePlanAfter"`

	// <p>Query time before optimization</p>
	QueryTimeBefore *float64 `json:"QueryTimeBefore,omitnil,omitempty" name:"QueryTimeBefore"`

	// <p>Query time after optimization</p>
	QueryTimeAfter *float64 `json:"QueryTimeAfter,omitnil,omitempty" name:"QueryTimeAfter"`

	// <p>Number of scanned rows before optimization</p>
	SQLScanRowsBefore *int64 `json:"SQLScanRowsBefore,omitnil,omitempty" name:"SQLScanRowsBefore"`

	// <p>Number of scanned rows after optimization</p>
	SQLScanRowsAfter *int64 `json:"SQLScanRowsAfter,omitnil,omitempty" name:"SQLScanRowsAfter"`
}

type ExplainRow struct {
	// <p>Serial number of the query</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Query type. Common values: SIMPLE (simple query, excluding subqueries or UNION), PRIMARY (outermost query), SUBQUERY (first SELECT in a subquery), DERIVED (derivative table/subquery in FROM clause), UNION (second and subsequent SELECTs in UNION), UNION RESULT (result set of UNION).</p>
	SelectType *string `json:"SelectType,omitnil,omitempty" name:"SelectType"`

	// <p>Data table name</p>
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// <p>Query matching partition</p>
	Partitions *string `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// <p>Access type (very important, a key metric for measuring query efficiency), ranked from best to worst: system &gt; const &gt; eq_ref &gt; ref &gt; fulltext &gt; ref_or_null &gt; index_merge &gt; unique_subquery &gt; index_subquery &gt; range &gt; index &gt; ALL. Common value descriptions: • system: The table has only one row (system table). • const: Matches one row through primary key or unique index, commonly seen in WHERE pk = 1. • eq_ref: Uses primary key or unique index when connecting, each index value matches only one row. • ref: Uses non-unique index to search, may match multiple rows. • range: Index range scan, such as BETWEEN, &gt;, &lt;, IN. • index: Full index scan (traversal of the entire index tree). • ALL: Full-table scan (worst, requires optimization).</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Indexes that may be used in querying. NULL means no available index.</p>
	PossibleKeys *string `json:"PossibleKeys,omitnil,omitempty" name:"PossibleKeys"`

	// <p>Indexes actually used. NULL means no index is used.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>The length of the actual used index (number of bytes). It can be used to determine which columns are actually used in the composite index. A shorter value indicates fewer index columns are used.</p>
	KeyLen *string `json:"KeyLen,omitnil,omitempty" name:"KeyLen"`

	// <p>Display which columns or constants are compared with indexes in the key column. Common values: const (constant), certain column name, func (function result).</p>
	Ref *string `json:"Ref,omitnil,omitempty" name:"Ref"`

	// <p>Estimate the number of rows to scan</p>
	Rows *int64 `json:"Rows,omitnil,omitempty" name:"Rows"`

	// <p>Indicates the estimation of the percentage of rows remaining after table conditional filtering. 100% means no additional filtering. The higher the value, the better.</p>
	Filtered *float64 `json:"Filtered,omitnil,omitempty" name:"Filtered"`

	// <p>Additional information (very important), common values: • Using index: covering index, no need to return to the table (good) • Using where: filter with WHERE after the storage engine returns rows • Using temporary: used a temporary table (commonly used in GROUP BY/ORDER BY, needs optimization) • Using filesort: used file sort instead of index sorting (needs optimization) • Using index condition: used pushdown of index (ICP)</p>
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

// Predefined struct for user
type ExportInstanceErrorLogsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The max number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log level.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// KeyWords.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The template type. Valid values: `csv`, `original`.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Valid value: `Timestamp`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. valid values: ASC or DESC.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceErrorLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Log end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The max number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Log level.
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// KeyWords.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The template type. Valid values: `csv`, `original`.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Valid value: `Timestamp`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. valid values: ASC or DESC.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *ExportInstanceErrorLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "LogLevels")
	delete(f, "KeyWords")
	delete(f, "FileType")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceErrorLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceErrorLogsResponseParams struct {
	// Exported error log content.
	ErrorLogItems []*ErrorLogItemExport `json:"ErrorLogItems,omitnil,omitempty" name:"ErrorLogItems"`

	// Error log string.
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportInstanceErrorLogsResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceErrorLogsResponseParams `json:"Response"`
}

func (r *ExportInstanceErrorLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceErrorLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Sorting field. Valid values: `QueryTime`, `LockTime`, `RowsExamined`, and `RowsSent`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc`, `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type ExportInstanceSlowQueriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Transaction start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Transaction end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Maximum number
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Username
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Client host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// File type. Valid values: csv, original.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Sorting field. Valid values: `QueryTime`, `LockTime`, `RowsExamined`, and `RowsSent`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc`, `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *ExportInstanceSlowQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Username")
	delete(f, "Host")
	delete(f, "Database")
	delete(f, "FileType")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportInstanceSlowQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportInstanceSlowQueriesResponseParams struct {
	// Slow query export content
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportInstanceSlowQueriesResponse struct {
	*tchttp.BaseResponse
	Response *ExportInstanceSlowQueriesResponseParams `json:"Response"`
}

func (r *ExportInstanceSlowQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportInstanceSlowQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportResourcePackageDeductDetailsRequestParams struct {
	// Resource package ID to be exported.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Specifies the cluster ID of the cynos cluster that uses the resource package capacity.
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Sorting field. currently supports: createTime (resource package deduction time), successDeductSpec (resource package deduction amount).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. supports ASC, DESC, ASC, DESC.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// A maximum of 2000 rows of data can be exported at a time. currently, a maximum of 2000 rows are supported.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset and page number.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Specifies the format for exporting data. currently only supports csv format, reserved for future expansion.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type ExportResourcePackageDeductDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Resource package ID to be exported.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Specifies the cluster ID of the cynos cluster that uses the resource package capacity.
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Sorting field. currently supports: createTime (resource package deduction time), successDeductSpec (resource package deduction amount).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting type. supports ASC, DESC, ASC, DESC.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// A maximum of 2000 rows of data can be exported at a time. currently, a maximum of 2000 rows are supported.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset and page number.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Specifies the format for exporting data. currently only supports csv format, reserved for future expansion.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *ExportResourcePackageDeductDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportResourcePackageDeductDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "ClusterIds")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportResourcePackageDeductDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportResourcePackageDeductDetailsResponseParams struct {
	// File detail.
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportResourcePackageDeductDetailsResponse struct {
	*tchttp.BaseResponse
	Response *ExportResourcePackageDeductDetailsResponseParams `json:"Response"`
}

func (r *ExportResourcePackageDeductDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportResourcePackageDeductDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardInstanceInfo struct {
	// Forward the instance id
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Forward instance region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type GdnTaskInfo struct {
	// Unique ID of the global database
	GdnId *string `json:"GdnId,omitnil,omitempty" name:"GdnId"`

	// unique alias of the global database
	GdnName *string `json:"GdnName,omitnil,omitempty" name:"GdnName"`

	// Primary Cluster ID
	PrimaryClusterId *string `json:"PrimaryClusterId,omitnil,omitempty" name:"PrimaryClusterId"`

	// Master cluster region
	PrimaryClusterRegion *string `json:"PrimaryClusterRegion,omitnil,omitempty" name:"PrimaryClusterRegion"`

	// from the cluster region
	StandbyClusterRegion *string `json:"StandbyClusterRegion,omitnil,omitempty" name:"StandbyClusterRegion"`

	// Slave cluster ID
	StandbyClusterId *string `json:"StandbyClusterId,omitnil,omitempty" name:"StandbyClusterId"`

	// From the cluster name
	StandbyClusterName *string `json:"StandbyClusterName,omitnil,omitempty" name:"StandbyClusterName"`

	// Whether forced switchover is performed
	ForceSwitchGdn *string `json:"ForceSwitchGdn,omitnil,omitempty" name:"ForceSwitchGdn"`

	// Return code
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// prompt message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Whether forced switchover is supported
	IsSupportForce *string `json:"IsSupportForce,omitnil,omitempty" name:"IsSupportForce"`
}

type GoodsPrice struct {
	// Specifies the instance price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Specifies the storage price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// Specifies the product specification.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GoodsSpec *GoodsSpec `json:"GoodsSpec,omitnil,omitempty" name:"GoodsSpec"`
}

type GoodsSpec struct {
	// Number of products
	// Note: This field may return null, indicating that no valid values can be obtained.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Number of CPU cores. required for PREPAID and POSTPAID instance types.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in gb. required for PREPAID and POSTPAID instance types.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Ccu size. required for serverless type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// Storage size. required for PREPAID storage type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Purchase duration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Instance machine type
	// 1. common, universal type.
	// 2. exclusive, dedicated.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of table permission codes
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// Database table information
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of table permission codes
	DbTablePrivileges []*string `json:"DbTablePrivileges,omitnil,omitempty" name:"DbTablePrivileges"`

	// Database table information
	DbTables []*DbTable `json:"DbTables,omitnil,omitempty" name:"DbTables"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "DbTablePrivileges")
	delete(f, "DbTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *GrantAccountPrivilegesResponseParams `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputAccount struct {
	// Account
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type InquirePriceCreateRequestParams struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Instance device type. Supported values are as follows:
	// - common: indicates the general type
	// - exclusive: indicates the exclusive type.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceCreateRequest struct {
	*tchttp.BaseRequest
	
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of compute node to purchase
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance type for purchase. Valid values: `PREPAID`, `POSTPAID`, `SERVERLESS`.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage type for purchase. Valid values: `PREPAID`, `POSTPAID`.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Instance device type. Supported values are as follows:
	// - common: indicates the general type
	// - exclusive: indicates the exclusive type.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Number of CPU cores, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in GB, which is required when `InstancePayMode` is `PREPAID` or `POSTPAID`.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CCU size, which is required when `InstancePayMode` is `SERVERLESS`.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`

	// Storage size, which is required when `StoragePayMode` is `PREPAID`.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Validity period, which is required when `InstancePayMode` is `PREPAID`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit, which is required when `InstancePayMode` is `PREPAID`. Valid values: `m` (month), `d` (day).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceCreateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "GoodsNum")
	delete(f, "InstancePayMode")
	delete(f, "StoragePayMode")
	delete(f, "DeviceType")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Ccu")
	delete(f, "StorageLimit")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateResponseParams struct {
	// Instance price
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Storage price
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateResponseParams `json:"Response"`
}

func (r *InquirePriceCreateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Specifies the number of CPU cores.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory Size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage size, storage resource adjustment: ClusterId, StorageLimit.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Instance ID. computational resource adjustment requires passing: ClusterId, instance ID, Cpu, Memory.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance device type.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Specifies the ccu size of the serverless instance.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`
}

type InquirePriceModifyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Specifies the number of CPU cores.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory Size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage size, storage resource adjustment: ClusterId, StorageLimit.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Instance ID. computational resource adjustment requires passing: ClusterId, instance ID, Cpu, Memory.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance device type.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Specifies the ccu size of the serverless instance.
	Ccu *float64 `json:"Ccu,omitnil,omitempty" name:"Ccu"`
}

func (r *InquirePriceModifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "StorageLimit")
	delete(f, "InstanceId")
	delete(f, "DeviceType")
	delete(f, "Ccu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyResponseParams struct {
	// Instance price.
	InstancePrice *TradePrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// Specifies the storage price.
	StoragePrice *TradePrice `json:"StoragePrice,omitnil,omitempty" name:"StoragePrice"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceModifyResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyResponseParams `json:"Response"`
}

func (r *InquirePriceModifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceMultiSpecRequestParams struct {
	// Availability zone. specifies the best practice for region provision.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance purchase type. available values are: PREPAID, POSTPAID, SERVERLESS.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage purchase type. available values are: PREPAID, POSTPAID.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Specifies the product specification.
	GoodsSpecs []*GoodsSpec `json:"GoodsSpecs,omitnil,omitempty" name:"GoodsSpecs"`
}

type InquirePriceMultiSpecRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone. specifies the best practice for region provision.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance purchase type. available values are: PREPAID, POSTPAID, SERVERLESS.
	InstancePayMode *string `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Storage purchase type. available values are: PREPAID, POSTPAID.
	StoragePayMode *string `json:"StoragePayMode,omitnil,omitempty" name:"StoragePayMode"`

	// Specifies the product specification.
	GoodsSpecs []*GoodsSpec `json:"GoodsSpecs,omitnil,omitempty" name:"GoodsSpecs"`
}

func (r *InquirePriceMultiSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceMultiSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstancePayMode")
	delete(f, "StoragePayMode")
	delete(f, "GoodsSpecs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceMultiSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceMultiSpecResponseParams struct {
	// Specifies the product price.
	GoodsPrice []*GoodsPrice `json:"GoodsPrice,omitnil,omitempty" name:"GoodsPrice"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceMultiSpecResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceMultiSpecResponseParams `json:"Response"`
}

func (r *InquirePriceMultiSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceMultiSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type InquirePriceRenewRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Validity period, which needs to be used together with `TimeUnit`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Unit of validity period, which needs to be used together with `TimeSpan`. Valid values: `d` (day), `m` (month).
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

func (r *InquirePriceRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Price of instance specification in array
	Prices []*TradePrice `json:"Prices,omitnil,omitempty" name:"Prices"`

	// Total renewal price of compute node
	InstanceRealTotalPrice *int64 `json:"InstanceRealTotalPrice,omitnil,omitempty" name:"InstanceRealTotalPrice"`

	// Total renewal price of storage node
	StorageRealTotalPrice *int64 `json:"StorageRealTotalPrice,omitnil,omitempty" name:"StorageRealTotalPrice"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRenewResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewResponseParams `json:"Response"`
}

func (r *InquirePriceRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceAbility struct {
	// Specifies whether the instance supports forced reboot. valid values: yes (supported), no (unsupported).
	IsSupportForceRestart *string `json:"IsSupportForceRestart,omitnil,omitempty" name:"IsSupportForceRestart"`

	// Specifies the causes for not supporting forced reboot.
	NonsupportForceRestartReason *string `json:"NonsupportForceRestartReason,omitnil,omitempty" name:"NonsupportForceRestartReason"`
}

type InstanceAuditRule struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies whether it is a rule-based audit. true - rule-based audit; false - comprehensive audit.
	AuditRule *bool `json:"AuditRule,omitnil,omitempty" name:"AuditRule"`

	// Specifies the audit rule details. valid when AuditRule=true.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Whether it is an audit policy.
	OldRule *bool `json:"OldRule,omitnil,omitempty" name:"OldRule"`

	// The rule template details of the instance application.
	RuleTemplates []*RuleTemplateInfo `json:"RuleTemplates,omitnil,omitempty" name:"RuleTemplates"`
}

type InstanceAuditStatus struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit status. ON - Indicates audit is enabled, OFF - Indicates audit is disabled.
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// Specifies the log retention period.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High-Frequency storage duration.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Specifies the duration of low-frequency storage.
	LowLogExpireDay *uint64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// Specifies the log storage volume.
	BillingAmount *float64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// Specifies the high-frequency storage volume.
	HighRealStorage *float64 `json:"HighRealStorage,omitnil,omitempty" name:"HighRealStorage"`

	// Specifies the infrequent access storage size.
	LowRealStorage *float64 `json:"LowRealStorage,omitnil,omitempty" name:"LowRealStorage"`

	// Specifies whether it is a full audit. true - indicates a full audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Specifies the audit activation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Instance-Related information.
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// Specifies the total storage capacity.
	RealStorage *float64 `json:"RealStorage,omitnil,omitempty" name:"RealStorage"`

	// The rule template applied to the instance.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Specifies whether to enable log delivery: ON, OFF.
	Deliver *string `json:"Deliver,omitnil,omitempty" name:"Deliver"`

	// Log shipping type.
	DeliverSummary []*DeliverSummary `json:"DeliverSummary,omitnil,omitempty" name:"DeliverSummary"`
}

type InstanceCLSDeliveryInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Log topic id.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Log topic name.
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Log set id.
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// Log set name.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Log delivery region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Delivery status: creating, running, offlining, offlined.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type InstanceInitInfo struct {
	// <p>Instance cpu</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Instance memory</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance type rw/ro</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Number of instances, range [1,15]</p>
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>Minimum count of Serverless instance, range [1,15]</p>
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// <p>Maximum count of Serverless instances, range [1,15]</p>
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// <p>Minimum specification of Serverless instance</p>
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// <p>Maximum specification of Serverless instance</p>
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// <p>Instance Machine Type</p><ol><li>common, universal type.</li><li>exclusive, dedicated.</li></ol>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type InstanceNameWeight struct {
	// Instance name. specifies the name defined by InstanceInitInfo.InstanceName in cluster creation.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Weight
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type InstanceNetInfo struct {
	// Network type.
	InstanceGroupType *string `json:"InstanceGroupType,omitnil,omitempty" name:"InstanceGroupType"`

	// Instance group ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Specifies the virtual private cloud ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Specifies the subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Network type. valid values: 0 (basic network), 1 (vpc network), 2 (blackstone network).
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// VPC IP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Specifies the public network domain name.
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public IP address
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public network port.
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network enabled status.
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`
}

type InstanceParamItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of instance parameters
	ParamsItems []*ParamItemDetail `json:"ParamsItems,omitnil,omitempty" name:"ParamsItems"`
}

type InstanceSet struct {
	// Database schema
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Number of CPU cores
	InstanceCpu *int64 `json:"InstanceCpu,omitnil,omitempty" name:"InstanceCpu"`

	// Instance type
	InstanceDeviceType *string `json:"InstanceDeviceType,omitnil,omitempty" name:"InstanceDeviceType"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// memory
	InstanceMemory *int64 `json:"InstanceMemory,omitnil,omitempty" name:"InstanceMemory"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance role
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Status description.
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// hard disk
	InstanceStorage *int64 `json:"InstanceStorage,omitnil,omitempty" name:"InstanceStorage"`

	// Hard disk type.
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`

	// Engine type.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Duration.
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Execution start time (seconds from 00:00)
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// The time when it can be executed. Enumeration value: ["Mon","Tue","Wed","Thu","Fri", "Sat", "Sun"].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`

	// Node list
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// Instance task
	InstanceTasks []*ObjectTask `json:"InstanceTasks,omitnil,omitempty" name:"InstanceTasks"`
}

type InstanceSpec struct {
	// Number of instance CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory in GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Maximum instance storage capacity GB
	MaxStorageSize *uint64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// Minimum instance storage capacity GB
	MinStorageSize *uint64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Machine type
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// Maximum IOPS
	MaxIops *int64 `json:"MaxIops,omitnil,omitempty" name:"MaxIops"`

	// Maximum bandwidth
	MaxIoBandWidth *int64 `json:"MaxIoBandWidth,omitnil,omitempty" name:"MaxIoBandWidth"`

	// Regional inventory information.
	ZoneStockInfos []*ZoneStockInfo `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`

	// Inventory quantity.
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// Maximum cpu
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// Minimum cpu
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`
}

type IntegrateCreateClusterConfig struct {
	// Retention days of binlog. value range: 7-1830.
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`

	// Specifies the backup retention days. value range: 7-1830.
	BackupSaveDays *int64 `json:"BackupSaveDays,omitnil,omitempty" name:"BackupSaveDays"`

	// Specifies the semi-sync timeout period. value range: [1000,4294967295].
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`

	// proxy connection address configuration message.
	ProxyEndPointConfigs []*ProxyEndPointConfigInfo `json:"ProxyEndPointConfigs,omitnil,omitempty" name:"ProxyEndPointConfigs"`
}

type IntegrateInstanceInfo struct {
	// Specifies the cpu of the instance.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Specifies the instance memory.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance type (rw/ro).
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Number of instances. value range: [1,15].
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Instance machine type. valid values: universal type (common), exclusive type.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

// Predefined struct for user
type IsolateClusterRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance return reason type.
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Instance return reason supplement.
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`

	// Retain backup, true - Retained (incur fees).
	SaveBackup *bool `json:"SaveBackup,omitnil,omitempty" name:"SaveBackup"`
}

type IsolateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance return reason type.
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Instance return reason supplement.
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`

	// Retain backup, true - Retained (incur fees).
	SaveBackup *bool `json:"SaveBackup,omitnil,omitempty" name:"SaveBackup"`
}

func (r *IsolateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbType")
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	delete(f, "SaveBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateClusterResponseParams struct {
	// Task flow ID (returned for pay-as-you-go or serverless resources. if necessary to sync task status, please use the DescribeFlow api).
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Refund order number (returned for prepaid resources. if necessary, synchronize the order status by using the billing product's DescribeDealsByCond to synchronize the order status).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateClusterResponse struct {
	*tchttp.BaseResponse
	Response *IsolateClusterResponseParams `json:"Response"`
}

func (r *IsolateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance Refund Reason Types
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Instance Refund Reason Supplement
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`

	// Backup retention
	SaveBackup *bool `json:"SaveBackup,omitnil,omitempty" name:"SaveBackup"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Instance Refund Reason Types
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Instance Refund Reason Supplement
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`

	// Backup retention
	SaveBackup *bool `json:"SaveBackup,omitnil,omitempty" name:"SaveBackup"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "DbType")
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	delete(f, "SaveBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order ID for isolated instance (prepaid instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstanceResponseParams `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateLibraDBClusterRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Isolation reason type
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Cause of isolation
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

type IsolateLibraDBClusterRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Isolation reason type
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Cause of isolation
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

func (r *IsolateLibraDBClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateLibraDBClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateLibraDBClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateLibraDBClusterResponseParams struct {
	// flow id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Return order ID
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateLibraDBClusterResponse struct {
	*tchttp.BaseResponse
	Response *IsolateLibraDBClusterResponseParams `json:"Response"`
}

func (r *IsolateLibraDBClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateLibraDBClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateLibraDBInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Whether to force isolation
	ForceIsolate *bool `json:"ForceIsolate,omitnil,omitempty" name:"ForceIsolate"`

	// Isolation reason type
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Cause of isolation
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

type IsolateLibraDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Whether to force isolation
	ForceIsolate *bool `json:"ForceIsolate,omitnil,omitempty" name:"ForceIsolate"`

	// Isolation reason type
	IsolateReasonTypes []*int64 `json:"IsolateReasonTypes,omitnil,omitempty" name:"IsolateReasonTypes"`

	// Cause of isolation
	IsolateReason *string `json:"IsolateReason,omitnil,omitempty" name:"IsolateReason"`
}

func (r *IsolateLibraDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateLibraDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "ForceIsolate")
	delete(f, "IsolateReasonTypes")
	delete(f, "IsolateReason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateLibraDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateLibraDBInstanceResponseParams struct {
	// task flow id
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Order ID list
	// Note: This field may return null, indicating that no valid values can be obtained.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateLibraDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateLibraDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateLibraDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateLibraDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LibraClusterSet struct {
	// User ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// instance information
	InstanceSet []*LibraInstanceSet `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Expiration time.
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Auto-renewal flag. 1: Automatic renewal. 0: Non-renewal upon expiration.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Storage size
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Used Capacity
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// VIP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vport port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type LibraDBClusterDetail struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Cluster name.</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Storage size</p>
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>VPC name</p>
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// <p>vpc Unique id</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Subnet name.</p>
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// <p>Subnet ID.</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Database version</p>
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// <p>Used capacity</p>
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// <p>vip address</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>vport</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>vip address and vport of the cluster read-only instance</p>
	RoAddr []*RoAddr `json:"RoAddr,omitnil,omitempty" name:"RoAddr"`

	// <p>cynos version</p>
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// <p>Freeze or not</p>
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// <p>Task List</p>
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Primary AZ</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>Instance collection</p>
	InstanceSet []*InstanceSet `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// <p>Payment mode</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Expiry time</p>
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// <p>Project ID.</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>Auto-renewal flag</p>
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>Version tag</p>
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// <p>Additions are not supported when ro is yes. Additions are supported when ro is no/null/&quot;&quot;.</p>
	NoSupportAddRo *string `json:"NoSupportAddRo,omitnil,omitempty" name:"NoSupportAddRo"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Physical AZ</p>
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// <p>Grayscale information for version upgrade</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisUpgradeVersionInfo *UpgradeAnalysisInstanceVersionInfo `json:"AnalysisUpgradeVersionInfo,omitnil,omitempty" name:"AnalysisUpgradeVersionInfo"`
}

type LibraDBClusterSet struct {
	// User ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// cynos version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Version tag
	CynosVersionTag *string `json:"CynosVersionTag,omitnil,omitempty" name:"CynosVersionTag"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Instance count
	InstanceNum *int64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// Whether to freeze
	IsFreeze *string `json:"IsFreeze,omitnil,omitempty" name:"IsFreeze"`

	// Network address.
	NetAddrs []*NetAddr `json:"NetAddrs,omitnil,omitempty" name:"NetAddrs"`

	// Payment mode
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Expiration time.
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// Project ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Auto-renewal flag. 1: Automatic renewal. 0: Non-renewal upon expiration.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Storage size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Task list
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// Account ID.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// VIP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Unique ID of VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vport port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Primary AZ.
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// physical AZ
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type LibraDBClusterSrcInfo struct {
	// Source endpoint type
	SrcInstanceType *string `json:"SrcInstanceType,omitnil,omitempty" name:"SrcInstanceType"`

	// Network type.
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// Source instance ID
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Source cluster ID
	SrcClusterId *string `json:"SrcClusterId,omitnil,omitempty" name:"SrcClusterId"`

	// Address.
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Port.
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Username.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Password.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Source sql mode
	SqlMode *string `json:"SqlMode,omitnil,omitempty" name:"SqlMode"`

	// Source application id
	SrcAppId *int64 `json:"SrcAppId,omitnil,omitempty" name:"SrcAppId"`

	// source account
	SrcUin *string `json:"SrcUin,omitnil,omitempty" name:"SrcUin"`

	// sub-account
	SrcSubAccountUin *string `json:"SrcSubAccountUin,omitnil,omitempty" name:"SrcSubAccountUin"`

	// Account
	AccountMode *string `json:"AccountMode,omitnil,omitempty" name:"AccountMode"`

	// Source instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Operation of the source instance
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type LibraDBInstanceInitInfo struct {
	// cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// memory
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// hard disk
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// Storage type
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance version
	LibraDBVersion *string `json:"LibraDBVersion,omitnil,omitempty" name:"LibraDBVersion"`

	// Instance count
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// subnet id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Number of instance replicas to purchase
	ReplicasNum *int64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`
}

type LibraDBNodeInfo struct {
	// LibraDB Node ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// syncing data
	DataStatus *string `json:"DataStatus,omitnil,omitempty" name:"DataStatus"`

	// Number of CPU cores.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size. Unit: GB.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size, in GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Error message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type LibraDBVersion struct {
	// Version
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Version tag
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Can be used this version
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`
}

type LibraInstanceSet struct {
	// Database schema
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Number of CPU cores
	InstanceCpu *int64 `json:"InstanceCpu,omitnil,omitempty" name:"InstanceCpu"`

	// Instance type
	InstanceDeviceType *string `json:"InstanceDeviceType,omitnil,omitempty" name:"InstanceDeviceType"`

	// Group ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// memory
	InstanceMemory *int64 `json:"InstanceMemory,omitnil,omitempty" name:"InstanceMemory"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Payment mode
	InstancePayMode *int64 `json:"InstancePayMode,omitnil,omitempty" name:"InstancePayMode"`

	// Payment end time
	InstancePeriodEndTime *string `json:"InstancePeriodEndTime,omitnil,omitempty" name:"InstancePeriodEndTime"`

	// Instance role
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Instance status
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// Instance status description
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// Network type.
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// Subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Virtual IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Virtual port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Public network region
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public IP address
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Public network status
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// hard disk
	InstanceStorage *int64 `json:"InstanceStorage,omitnil,omitempty" name:"InstanceStorage"`

	// Hard disk type.
	InstanceStorageType *string `json:"InstanceStorageType,omitnil,omitempty" name:"InstanceStorageType"`
}

type LogFilter struct {
	// Filter items.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter criteria. Support the following conditions: WINC-include (segment dimension), WEXC-exclude (segment dimension), INC-include, EXC-exclude, EQS-equal, NEQ-not equal to, RA-range.
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filtered value. For reverse query, multiple values have an AND relationship. For forward query, multiple values have an OR relationship.
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogicBackupConfigInfo struct {
	// Whether automatic logical backup is enabled.
	LogicBackupEnable *string `json:"LogicBackupEnable,omitnil,omitempty" name:"LogicBackupEnable"`

	// Specifies the automatic logic backup start time.
	LogicBackupTimeBeg *uint64 `json:"LogicBackupTimeBeg,omitnil,omitempty" name:"LogicBackupTimeBeg"`

	// Specifies the termination time of automatic logical backup.
	LogicBackupTimeEnd *uint64 `json:"LogicBackupTimeEnd,omitnil,omitempty" name:"LogicBackupTimeEnd"`

	// Automatic logical backup retention time.
	// Unit: seconds.
	LogicReserveDuration *uint64 `json:"LogicReserveDuration,omitnil,omitempty" name:"LogicReserveDuration"`

	// Is cross-regional logical backup enabled?.
	// Valid values: ON/OFF.
	LogicCrossRegionsEnable *string `json:"LogicCrossRegionsEnable,omitnil,omitempty" name:"LogicCrossRegionsEnable"`

	// Regions covered by logical backup.
	LogicCrossRegions []*string `json:"LogicCrossRegions,omitnil,omitempty" name:"LogicCrossRegions"`

	// Backup delivery relationship
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

type ManualBackupData struct {
	// Backup type. snapshot - Snapshot backup.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Backup method. auto - Automatic backup, manual - Manual backup.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Backup time.
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// Detailed information of cross-region backup items.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CrossRegionBackupInfos []*CrossRegionBackupItem `json:"CrossRegionBackupInfos,omitnil,omitempty" name:"CrossRegionBackupInfos"`
}

type MigrateDBItem struct {
	// Database name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Data table migration mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrateTableMode *string `json:"MigrateTableMode,omitnil,omitempty" name:"MigrateTableMode"`

	// Data table information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tables []*MigrateTableItem `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type MigrateObject struct {
	// Database migration mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrateDBMode *string `json:"MigrateDBMode,omitnil,omitempty" name:"MigrateDBMode"`

	// Database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Databases []*MigrateDBItem `json:"Databases,omitnil,omitempty" name:"Databases"`
}

type MigrateOpt struct {
	// Database table information
	DatabaseTables *MigrateObject `json:"DatabaseTables,omitnil,omitempty" name:"DatabaseTables"`
}

type MigrateTableItem struct {
	// Data table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type ModifiableInfo struct {
	// Whether the parameter can be modified, 1: Yes 0: No.
	IsModifiable *int64 `json:"IsModifiable,omitnil,omitempty" name:"IsModifiable"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Database account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountName")
	delete(f, "Description")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New host
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type ModifyAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// New host
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

func (r *ModifyAccountHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NewHost")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountHostResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountHostResponseParams `json:"Response"`
}

func (r *ModifyAccountHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of global permissions
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Array of database permissions
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Array of table permissions
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account infomation
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Array of global permissions
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Array of database permissions
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Array of table permissions
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alert Levels. 1 - Low Risk, 2 - Medium Risk, 3 - High Risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alert policy. 0 - No alert, 1 - Alert.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit rule after modification
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// New name of the rule template
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// New description of the rule template
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alert Levels. 1 - Low Risk, 2 - Medium Risk, 3 - High Risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alert policy. 0 - No alert, 1 - Alert.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

func (r *ModifyAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *ModifyAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// The parameter used to change the audit rule of the instance to full audit
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Rule audit
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditAll")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Indicates the full backup start time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// <p>Indicates the full backup end time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// <p>Indicates the backup retention period in seconds. Backups will be cleaned up longer than this time. Seven days is represented as 3600<em>24</em>7=604800. The maximum value is 158112000.</p>
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// <p>This parameter currently does not support modification and is not required. Backup frequency is an array of length 7, corresponding to Monday to Sunday backup method, full-full backup, increment-incremental backup.</p>
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// <p>This parameter currently does not support modification. No need to specify.</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Logical backup configuration</p>
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// <p>Whether to delete automatic logical backup</p>
	DeleteAutoLogicBackup *bool `json:"DeleteAutoLogicBackup,omitnil,omitempty" name:"DeleteAutoLogicBackup"`

	// <p>Second-level snapshot backup parameter</p>
	SnapshotSecondaryBackupConfig *SnapshotBackupConfig `json:"SnapshotSecondaryBackupConfig,omitnil,omitempty" name:"SnapshotSecondaryBackupConfig"`

	// <p>Sparse backup configuration</p>
	SparseBackupConfig *SparseBackupConfig `json:"SparseBackupConfig,omitnil,omitempty" name:"SparseBackupConfig"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Indicates the full backup start time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// <p>Indicates the full backup end time, [0-24*3600]. For example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.</p>
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// <p>Indicates the backup retention period in seconds. Backups will be cleaned up longer than this time. Seven days is represented as 3600<em>24</em>7=604800. The maximum value is 158112000.</p>
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// <p>This parameter currently does not support modification and is not required. Backup frequency is an array of length 7, corresponding to Monday to Sunday backup method, full-full backup, increment-incremental backup.</p>
	BackupFreq []*string `json:"BackupFreq,omitnil,omitempty" name:"BackupFreq"`

	// <p>This parameter currently does not support modification. No need to specify.</p>
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// <p>Logical backup configuration</p>
	LogicBackupConfig *LogicBackupConfigInfo `json:"LogicBackupConfig,omitnil,omitempty" name:"LogicBackupConfig"`

	// <p>Whether to delete automatic logical backup</p>
	DeleteAutoLogicBackup *bool `json:"DeleteAutoLogicBackup,omitnil,omitempty" name:"DeleteAutoLogicBackup"`

	// <p>Second-level snapshot backup parameter</p>
	SnapshotSecondaryBackupConfig *SnapshotBackupConfig `json:"SnapshotSecondaryBackupConfig,omitnil,omitempty" name:"SnapshotSecondaryBackupConfig"`

	// <p>Sparse backup configuration</p>
	SparseBackupConfig *SparseBackupConfig `json:"SparseBackupConfig,omitnil,omitempty" name:"SparseBackupConfig"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupTimeBeg")
	delete(f, "BackupTimeEnd")
	delete(f, "ReserveDuration")
	delete(f, "BackupFreq")
	delete(f, "BackupType")
	delete(f, "LogicBackupConfig")
	delete(f, "DeleteAutoLogicBackup")
	delete(f, "SnapshotSecondaryBackupConfig")
	delete(f, "SparseBackupConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Download limit type. valid values: NoLimit (unlimited), LimitOnlyIntranet (limited to private network), Customize (custom).
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// This parameter only supports In, which indicates that the vpc specified by LimitVpc can be downloaded. the default is In.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Specified ips can download; specified ips are not allowed to download.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// Limit the vpc settings for downloads.
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// Specifies the ip settings for limiting downloads.
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Download limit type. valid values: NoLimit (unlimited), LimitOnlyIntranet (limited to private network), Customize (custom).
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// This parameter only supports In, which indicates that the vpc specified by LimitVpc can be downloaded. the default is In.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Specified ips can download; specified ips are not allowed to download.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// Limit the vpc settings for downloads.
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// Specifies the ip settings for limiting downloads.
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpcs")
	delete(f, "LimitIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadUserRestrictionRequestParams struct {
	// Download limit type. valid values: NoLimit (unlimited), LimitOnlyIntranet (limited to private network), Customize (custom).
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// This parameter only supports In, which indicates that the vpc specified by LimitVpc can be downloaded. the default is In.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Specified ips can download; specified ips are not allowed to download.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// Limit the vpc settings for downloads.
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// Specifies the ip settings for limiting downloads.
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

type ModifyBackupDownloadUserRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Download limit type. valid values: NoLimit (unlimited), LimitOnlyIntranet (limited to private network), Customize (custom).
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// This parameter only supports In, which indicates that the vpc specified by LimitVpc can be downloaded. the default is In.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Specified ips can download; specified ips are not allowed to download.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// Limit the vpc settings for downloads.
	LimitVpcs []*BackupLimitVpcItem `json:"LimitVpcs,omitnil,omitempty" name:"LimitVpcs"`

	// Specifies the ip settings for limiting downloads.
	LimitIps []*string `json:"LimitIps,omitnil,omitempty" name:"LimitIps"`
}

func (r *ModifyBackupDownloadUserRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadUserRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpcs")
	delete(f, "LimitIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadUserRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadUserRestrictionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadUserRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadUserRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadUserRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadUserRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Backup file ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup name, which can contain up to 60 characters.
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

func (r *ModifyBackupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BackupId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupNameResponseParams `json:"Response"`
}

func (r *ModifyBackupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogConfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Specifies the Binlog configuration message.
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`
}

type ModifyBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Specifies the Binlog configuration message.
	BinlogConfig *BinlogConfigInfo `json:"BinlogConfig,omitnil,omitempty" name:"BinlogConfig"`
}

func (r *ModifyBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBinlogConfigResponseParams `json:"Response"`
}

func (r *ModifyBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
}

type ModifyBinlogSaveDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Binlog retention period in days
	BinlogSaveDays *int64 `json:"BinlogSaveDays,omitnil,omitempty" name:"BinlogSaveDays"`
}

func (r *ModifyBinlogSaveDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "BinlogSaveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBinlogSaveDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBinlogSaveDaysResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBinlogSaveDaysResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBinlogSaveDaysResponseParams `json:"Response"`
}

func (r *ModifyBinlogSaveDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBinlogSaveDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterBinlogRedoLogAutoCopyVaultRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Automatically copy the safe configuration list
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

type ModifyClusterBinlogRedoLogAutoCopyVaultRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Automatically copy the safe configuration list
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

func (r *ModifyClusterBinlogRedoLogAutoCopyVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterBinlogRedoLogAutoCopyVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoCopyVaults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterBinlogRedoLogAutoCopyVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterBinlogRedoLogAutoCopyVaultResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterBinlogRedoLogAutoCopyVaultResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterBinlogRedoLogAutoCopyVaultResponseParams `json:"Response"`
}

func (r *ModifyClusterBinlogRedoLogAutoCopyVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterBinlogRedoLogAutoCopyVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Host permissions of the new authorized user
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Host permissions of previously authorized users
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
}

type ModifyClusterDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Host permissions of the new authorized user
	NewUserHostPrivileges []*UserHostPrivilege `json:"NewUserHostPrivileges,omitnil,omitempty" name:"NewUserHostPrivileges"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Host permissions of previously authorized users
	OldUserHostPrivileges []*UserHostPrivilege `json:"OldUserHostPrivileges,omitnil,omitempty" name:"OldUserHostPrivileges"`
}

func (r *ModifyClusterDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DbName")
	delete(f, "NewUserHostPrivileges")
	delete(f, "Description")
	delete(f, "OldUserHostPrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterDatabaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterDatabaseResponseParams `json:"Response"`
}

func (r *ModifyClusterDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterGlobalEncryptionRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Turn on or off global encryption
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

type ModifyClusterGlobalEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Turn on or off global encryption
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

func (r *ModifyClusterGlobalEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterGlobalEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsOpenGlobalEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterGlobalEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterGlobalEncryptionResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterGlobalEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterGlobalEncryptionResponseParams `json:"Response"`
}

func (r *ModifyClusterGlobalEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterGlobalEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNameResponseParams `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterParamRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// List of the parameters to be modified. Each element in the list is a combination of `ParamName`, `CurrentValue`, and `OldValue`. `ParamName` is the parameter name; `CurrentValue` is the current value; `OldValue` is the old value that doesn’t need to be verified.
	ParamList []*ParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Valid values: `yes` (execute during maintenance time), `no` (execute now)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyClusterParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterParamResponseParams struct {
	// Async request ID used to query the result
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterParamResponseParams `json:"Response"`
}

func (r *ModifyClusterParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type ModifyClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *ModifyClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *ModifyClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterReadOnlyRequestParams struct {
	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster read-only switch, valid values: ON, OFF.
	ReadOnlyOperation *string `json:"ReadOnlyOperation,omitnil,omitempty" name:"ReadOnlyOperation"`

	// Valid values: `yes` (modify in maintenance window), `no` (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyClusterReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// List of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// Cluster read-only switch, valid values: ON, OFF.
	ReadOnlyOperation *string `json:"ReadOnlyOperation,omitnil,omitempty" name:"ReadOnlyOperation"`

	// Valid values: `yes` (modify in maintenance window), `no` (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyClusterReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "ReadOnlyOperation")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterReadOnlyResponseParams struct {
	// List of cluster task ids.
	ClusterTaskIds []*ClusterTaskId `json:"ClusterTaskIds,omitnil,omitempty" name:"ClusterTaskIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterReadOnlyResponseParams `json:"Response"`
}

func (r *ModifyClusterReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// Specifies the binlog synchronization mode. the default value is async. valid values are sync, semisync, and async.
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// Semi-sync timeout in ms. To ensure business stability, semi-synchronous replication has a degradation logic. When the primary availability zone cluster waits for the secondary availability zone cluster to confirm a transaction, if the timeout period is exceeded, the replication method will degrade to asynchronous replication. The minimum is set to 1000 ms, supporting up to 4294967295 ms, with a default of 10000 ms.
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type ModifyClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Old replica AZ
	OldSlaveZone *string `json:"OldSlaveZone,omitnil,omitempty" name:"OldSlaveZone"`

	// New replica AZ
	NewSlaveZone *string `json:"NewSlaveZone,omitnil,omitempty" name:"NewSlaveZone"`

	// Specifies the binlog synchronization mode. the default value is async. valid values are sync, semisync, and async.
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// Semi-sync timeout in ms. To ensure business stability, semi-synchronous replication has a degradation logic. When the primary availability zone cluster waits for the secondary availability zone cluster to confirm a transaction, if the timeout period is exceeded, the replication method will degrade to asynchronous replication. The minimum is set to 1000 ms, supporting up to 4294967295 ms, with a default of 10000 ms.
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

func (r *ModifyClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldSlaveZone")
	delete(f, "NewSlaveZone")
	delete(f, "BinlogSyncWay")
	delete(f, "SemiSyncTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterSlaveZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *ModifyClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// <p>Network Group id (cynosdbmysql-grp- prefix at the beginning) or Cluster id (such as cynosdbmysql-xxxxxxxx prefix). When configuring security group with instance IP address triplet (UniqVpcId, Vip, Vport), this field must be set to Cluster id (such as cynosdbmysql-xxxxxxxx prefix).</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>List of security group IDs to modify, an array of one or more security group IDs.<br>Note: This input will perform a full replacement of the existing collection, not an incremental update. The modification requires the expected full collection.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>ID of the VPC network for instance ownership. (UniqVpcId, Vip, and Vport must be specified simultaneously to uniquely identify the network instance)</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Instance IP address, instance IP address triplet (UniqVpcId, Vip, and Vport). These three parameters must be specified simultaneously for unique identification of network instances.</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Instance port, instance IP address triplet (UniqVpcId, Vip, and Vport). These three parameters must be specified simultaneously for unique identification of network instances.</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Network Group id (cynosdbmysql-grp- prefix at the beginning) or Cluster id (such as cynosdbmysql-xxxxxxxx prefix). When configuring security group with instance IP address triplet (UniqVpcId, Vip, Vport), this field must be set to Cluster id (such as cynosdbmysql-xxxxxxxx prefix).</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>List of security group IDs to modify, an array of one or more security group IDs.<br>Note: This input will perform a full replacement of the existing collection, not an incremental update. The modification requires the expected full collection.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>ID of the VPC network for instance ownership. (UniqVpcId, Vip, and Vport must be specified simultaneously to uniquely identify the network instance)</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Instance IP address, instance IP address triplet (UniqVpcId, Vip, and Vport). These three parameters must be specified simultaneously for unique identification of network instances.</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Instance port, instance IP address triplet (UniqVpcId, Vip, and Vport). These three parameters must be specified simultaneously for unique identification of network instances.</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	delete(f, "Zone")
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDbVersionData struct {
	// Version before modification.
	OldVersion *string `json:"OldVersion,omitnil,omitempty" name:"OldVersion"`

	// Version after modification.
	NewVersion *string `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// Upgrade method.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type ModifyInstanceData struct {
	// CPU after resizing.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory after resizing.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Storage upper limit after resizing.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// CPU before resizing.
	OldCpu *int64 `json:"OldCpu,omitnil,omitempty" name:"OldCpu"`

	// Memory before resizing.
	OldMemory *int64 `json:"OldMemory,omitnil,omitempty" name:"OldMemory"`

	// Storage upper limit before resizing.
	OldStorageLimit *int64 `json:"OldStorageLimit,omitnil,omitempty" name:"OldStorageLimit"`

	// Instance Machine Type Before Scaling
	// 1. common, universal type.
	// 2. exclusive, dedicated.
	OldDeviceType *string `json:"OldDeviceType,omitnil,omitempty" name:"OldDeviceType"`

	// Instance Machine Type After Scaling
	// 1. common, universal type.
	// 2. exclusive, dedicated.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Upgrade method. Switch after upgrade completes or switch within maintenance window.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// Specifies the quantity of libra nodes.
	LibraNodeCount *int64 `json:"LibraNodeCount,omitnil,omitempty" name:"LibraNodeCount"`

	// Specifies the original quantity of libra nodes.
	OldLibraNodeCount *int64 `json:"OldLibraNodeCount,omitnil,omitempty" name:"OldLibraNodeCount"`
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster parameters
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// List of the instance parameters
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// Valid values: `yes` (modify in maintenance window),  `no`  (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of cluster parameters
	ClusterParamList []*ModifyParamItem `json:"ClusterParamList,omitnil,omitempty" name:"ClusterParamList"`

	// List of the instance parameters
	InstanceParamList []*ModifyParamItem `json:"InstanceParamList,omitnil,omitempty" name:"InstanceParamList"`

	// Valid values: `yes` (modify in maintenance window),  `no`  (execute now by default).
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "ClusterParamList")
	delete(f, "InstanceParamList")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceUpgradeLimitDaysRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Upgrade time limit.
	UpgradeLimitDays *int64 `json:"UpgradeLimitDays,omitnil,omitempty" name:"UpgradeLimitDays"`
}

type ModifyInstanceUpgradeLimitDaysRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Upgrade time limit.
	UpgradeLimitDays *int64 `json:"UpgradeLimitDays,omitnil,omitempty" name:"UpgradeLimitDays"`
}

func (r *ModifyInstanceUpgradeLimitDaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceUpgradeLimitDaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "UpgradeLimitDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceUpgradeLimitDaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceUpgradeLimitDaysResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceUpgradeLimitDaysResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceUpgradeLimitDaysResponseParams `json:"Response"`
}

func (r *ModifyInstanceUpgradeLimitDaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceUpgradeLimitDaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountDescriptionRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// host name
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ModifyLibraDBClusterAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account name.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// host name
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ModifyLibraDBClusterAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountName")
	delete(f, "Description")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountDescriptionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountHostRequestParams struct {
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// host name
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`
}

type ModifyLibraDBClusterAccountHostRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account information
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// host name
	NewHost *string `json:"NewHost,omitnil,omitempty" name:"NewHost"`
}

func (r *ModifyLibraDBClusterAccountHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "NewHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterAccountHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountHostResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterAccountHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterAccountHostResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterAccountHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountPrivilegeRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Global permission
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database permission
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table permission
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type ModifyLibraDBClusterAccountPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Account
	Account *InputAccount `json:"Account,omitnil,omitempty" name:"Account"`

	// Global permission
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database permission
	DatabasePrivileges []*DatabasePrivileges `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table permission
	TablePrivileges []*TablePrivileges `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

func (r *ModifyLibraDBClusterAccountPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Account")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterAccountPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterAccountPrivilegeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterAccountPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterAccountPrivilegeResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterAccountPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterAccountPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterDataSourceRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Source information
	SrcInfo []*LibraDBClusterSrcInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`
}

type ModifyLibraDBClusterDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Source information
	SrcInfo []*LibraDBClusterSrcInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`
}

func (r *ModifyLibraDBClusterDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "SrcInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterDataSourceResponseParams struct {
	// Asynchronous Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterDataSourceResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterNameRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ModifyLibraDBClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

func (r *ModifyLibraDBClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterNameResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterProjectRequestParams struct {
	// Analysis Cluster ID List
	ClusterIdSet []*string `json:"ClusterIdSet,omitnil,omitempty" name:"ClusterIdSet"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyLibraDBClusterProjectRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID List
	ClusterIdSet []*string `json:"ClusterIdSet,omitnil,omitempty" name:"ClusterIdSet"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyLibraDBClusterProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterProjectResponseParams struct {
	// This example shows you how to obtain the cluster list.
	AffectedClusterIdSet []*string `json:"AffectedClusterIdSet,omitnil,omitempty" name:"AffectedClusterIdSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterProjectResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterReplicationObjectRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Map mode
	ForceDefaultMapRule *string `json:"ForceDefaultMapRule,omitnil,omitempty" name:"ForceDefaultMapRule"`

	// sync object
	Objects []*ReplicationObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Automated mapping rule
	AutoMapRules []*AutoMapRule `json:"AutoMapRules,omitnil,omitempty" name:"AutoMapRules"`

	// Whether to refresh existing mapping relationships according to the latest mapping rule
	RefreshMapping *bool `json:"RefreshMapping,omitnil,omitempty" name:"RefreshMapping"`
}

type ModifyLibraDBClusterReplicationObjectRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Map mode
	ForceDefaultMapRule *string `json:"ForceDefaultMapRule,omitnil,omitempty" name:"ForceDefaultMapRule"`

	// sync object
	Objects []*ReplicationObject `json:"Objects,omitnil,omitempty" name:"Objects"`

	// Automated mapping rule
	AutoMapRules []*AutoMapRule `json:"AutoMapRules,omitnil,omitempty" name:"AutoMapRules"`

	// Whether to refresh existing mapping relationships according to the latest mapping rule
	RefreshMapping *bool `json:"RefreshMapping,omitnil,omitempty" name:"RefreshMapping"`
}

func (r *ModifyLibraDBClusterReplicationObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterReplicationObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "ForceDefaultMapRule")
	delete(f, "Objects")
	delete(f, "AutoMapRules")
	delete(f, "RefreshMapping")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBClusterReplicationObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBClusterReplicationObjectResponseParams struct {
	// Asynchronous Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBClusterReplicationObjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBClusterReplicationObjectResponseParams `json:"Response"`
}

func (r *ModifyLibraDBClusterReplicationObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBClusterReplicationObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBForwardConfigRequestParams struct {
	// Read-only analysis engine instance Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Forwarding Mode
	ForwardMode *string `json:"ForwardMode,omitnil,omitempty" name:"ForwardMode"`

	// Forward instance list
	ForwardList []*ForwardInstanceInfo `json:"ForwardList,omitnil,omitempty" name:"ForwardList"`
}

type ModifyLibraDBForwardConfigRequest struct {
	*tchttp.BaseRequest
	
	// Read-only analysis engine instance Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Forwarding Mode
	ForwardMode *string `json:"ForwardMode,omitnil,omitempty" name:"ForwardMode"`

	// Forward instance list
	ForwardList []*ForwardInstanceInfo `json:"ForwardList,omitnil,omitempty" name:"ForwardList"`
}

func (r *ModifyLibraDBForwardConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBForwardConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ForwardMode")
	delete(f, "ForwardList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraDBForwardConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraDBForwardConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraDBForwardConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraDBForwardConfigResponseParams `json:"Response"`
}

func (r *ModifyLibraDBForwardConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraDBForwardConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type ModifyMaintainPeriodConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance start time in seconds. For example, 03:00 AM is represented by 10800
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Maintenance duration in seconds. For example, one hour is represented by 3600
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Maintenance days of the week. Valid values: [Mon, Tue, Wed, Thu, Fri, Sat, Sun].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

func (r *ModifyMaintainPeriodConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintainStartTime")
	delete(f, "MaintainDuration")
	delete(f, "MaintainWeekDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintainPeriodConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintainPeriodConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMaintainPeriodConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintainPeriodConfigResponseParams `json:"Response"`
}

func (r *ModifyMaintainPeriodConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintainPeriodConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamItem struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Current parameter value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Old parameter value (useful only in output parameters).
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// libra component type.
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Parameter list
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Parameter list
	ParamList []*ModifyParamItem `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyParamsData struct {
	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value before modification.
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// Parameter value after modification.
	CurValue *string `json:"CurValue,omitnil,omitempty" name:"CurValue"`
}

// Predefined struct for user
type ModifyProxyDescRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Database proxy description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyProxyDescRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Database proxy description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyProxyDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyDescResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProxyDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyDescResponseParams `json:"Response"`
}

func (r *ModifyProxyDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitRequestParams struct {
	// <p>Cluster ID, for example, cynosdbmysql-asd123</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Database Proxy Group ID, such as cynosdbmysql-proxy-qwe123</p>
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// <p>Consistency Type "eventual"-final consistency, "session"-session consistency, "global"-global consistency</p>
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// <p>Consistency timeout.<br>Value ranges from 0 to 1000000 (microseconds). When set to 0, it means the request will wait if the read-only instance fails to satisfy the consistency policy due to delay.</p>
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// <p>Read-write weight allocation mode; system Auto-Assignment: "system", custom: "custom"</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Instance read-only weight.</p>
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// <p>Whether fault migration is enabled. After a failure occurs, the connection address will be routed to primary instance. Value: "yes", "no"</p>
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Automatically add read-only instance, value: "yes", "no"</p>
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether to enable read-write separation.<br>This parameter is deprecated. Set the read-write attribute through RwType.</p>
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// <p>Read-write type:<br>READWRITE, READONLY</p>
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// <p>Transaction split.<br>Read and write operations in a transaction are split and executed on different instances.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Connection mode:<br>nearby, balance</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Whether to enable the connection pool:<br>yes, no</p>
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// <p>Connection pool Type:<br>SessionConnectionPool</p>
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// <p>Connection pool time.<br>Optional range: 0-300 (seconds).</p>
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// <p>Whether to treat the libra node as an ordinary RO node</p>
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// <p>Whether to forward to other nodes when a libra node fault occurs</p>
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`

	// <p>Load balancing mode</p><p>Enumeration value:</p><ul><li>static: Static load</li><li>dynamic: Dynamic load</li></ul>
	LoadBalanceMode *string `json:"LoadBalanceMode,omitnil,omitempty" name:"LoadBalanceMode"`
}

type ModifyProxyRwSplitRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID, for example, cynosdbmysql-asd123</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Database Proxy Group ID, such as cynosdbmysql-proxy-qwe123</p>
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// <p>Consistency Type "eventual"-final consistency, "session"-session consistency, "global"-global consistency</p>
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// <p>Consistency timeout.<br>Value ranges from 0 to 1000000 (microseconds). When set to 0, it means the request will wait if the read-only instance fails to satisfy the consistency policy due to delay.</p>
	ConsistencyTimeOut *string `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// <p>Read-write weight allocation mode; system Auto-Assignment: "system", custom: "custom"</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Instance read-only weight.</p>
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// <p>Whether fault migration is enabled. After a failure occurs, the connection address will be routed to primary instance. Value: "yes", "no"</p>
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Automatically add read-only instance, value: "yes", "no"</p>
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether to enable read-write separation.<br>This parameter is deprecated. Set the read-write attribute through RwType.</p>
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// <p>Read-write type:<br>READWRITE, READONLY</p>
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// <p>Transaction split.<br>Read and write operations in a transaction are split and executed on different instances.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Connection mode:<br>nearby, balance</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Whether to enable the connection pool:<br>yes, no</p>
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// <p>Connection pool Type:<br>SessionConnectionPool</p>
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// <p>Connection pool time.<br>Optional range: 0-300 (seconds).</p>
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// <p>Whether to treat the libra node as an ordinary RO node</p>
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// <p>Whether to forward to other nodes when a libra node fault occurs</p>
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`

	// <p>Load balancing mode</p><p>Enumeration value:</p><ul><li>static: Static load</li><li>dynamic: Dynamic load</li></ul>
	LoadBalanceMode *string `json:"LoadBalanceMode,omitnil,omitempty" name:"LoadBalanceMode"`
}

func (r *ModifyProxyRwSplitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	delete(f, "ConsistencyType")
	delete(f, "ConsistencyTimeOut")
	delete(f, "WeightMode")
	delete(f, "InstanceWeights")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "OpenRw")
	delete(f, "RwType")
	delete(f, "TransSplit")
	delete(f, "AccessMode")
	delete(f, "OpenConnectionPool")
	delete(f, "ConnectionPoolType")
	delete(f, "ConnectionPoolTimeOut")
	delete(f, "ApNodeAsRoNode")
	delete(f, "ApQueryToOtherNode")
	delete(f, "LoadBalanceMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProxyRwSplitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProxyRwSplitResponseParams struct {
	// <p>Async FlowId</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Asynchronous task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProxyRwSplitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProxyRwSplitResponseParams `json:"Response"`
}

func (r *ModifyProxyRwSplitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProxyRwSplitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// ID of the cluster to be bound
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// ID of the cluster to be unbound
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
}

type ModifyResourcePackageClustersRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// ID of the cluster to be bound
	BindClusterIds []*string `json:"BindClusterIds,omitnil,omitempty" name:"BindClusterIds"`

	// ID of the cluster to be unbound
	UnbindClusterIds []*string `json:"UnbindClusterIds,omitnil,omitempty" name:"UnbindClusterIds"`
}

func (r *ModifyResourcePackageClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "BindClusterIds")
	delete(f, "UnbindClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageClustersResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackageClustersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageClustersResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Custom resource pack name, which can contains up to 120 characters.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ModifyResourcePackageNameRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Custom resource pack name, which can contains up to 120 characters.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *ModifyResourcePackageNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackageNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackageNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackageNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackageNameResponseParams `json:"Response"`
}

func (r *ModifyResourcePackageNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackageNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackagesDeductionPriorityRequestParams struct {
	// Specifies the resource package type whose priority needs to be modified. CCU: compute resource package. DISK: storage resource package.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// The modified deduction priority takes effect for which cynos resource.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Resource package deduction priority.
	DeductionPriorities []*PackagePriority `json:"DeductionPriorities,omitnil,omitempty" name:"DeductionPriorities"`
}

type ModifyResourcePackagesDeductionPriorityRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the resource package type whose priority needs to be modified. CCU: compute resource package. DISK: storage resource package.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// The modified deduction priority takes effect for which cynos resource.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Resource package deduction priority.
	DeductionPriorities []*PackagePriority `json:"DeductionPriorities,omitnil,omitempty" name:"DeductionPriorities"`
}

func (r *ModifyResourcePackagesDeductionPriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackagesDeductionPriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageType")
	delete(f, "ClusterId")
	delete(f, "DeductionPriorities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePackagesDeductionPriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePackagesDeductionPriorityResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePackagesDeductionPriorityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePackagesDeductionPriorityResponseParams `json:"Response"`
}

func (r *ModifyResourcePackagesDeductionPriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePackagesDeductionPriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServerlessStrategyRequestParams struct {
	// <p>serverless cluster id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Whether the cluster auto-pause feature is enabled. Optional range</p><li>yes</li><li>no</li>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>Delay of Cluster Auto-Pause in seconds, optional range [600,691200], default 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>This parameter is temporarily unavailable</p>
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// <p>This parameter is temporarily invalid</p>
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// <p>cpu minimum value. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Maximum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>Minimum value of read-only instance cpu. For the optional range, refer to the API response of DescribeServerlessInstanceSpecs.</p>
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// <p>Read-only cpu maximum value of the optional range. For reference, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// <p>Minimum count of read-only nodes</p>
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// <p>Maximum number of read-only nodes</p>
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// <p>Whether archiving is enabled. Optional range</p><li>yes</li><li>no</li>Default value: yes</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Upgrade type. Default value: upgradeImmediate. Available values: upgradeImmediate - immediately complete the modification; upgradeInMaintain - complete the modification during maintenance window.</p>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// <p>List of security groups to which newly-added read-only instances need to be bound. This only applies to binding security groups to read-only instances generated during the process of this adjustment of policies. Existing instances are not bound.</p>
	SecurityGroupIdsForNewRo []*string `json:"SecurityGroupIdsForNewRo,omitnil,omitempty" name:"SecurityGroupIdsForNewRo"`
}

type ModifyServerlessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <p>serverless cluster id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Whether the cluster auto-pause feature is enabled. Optional range</p><li>yes</li><li>no</li>
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>Delay of Cluster Auto-Pause in seconds, optional range [600,691200], default 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>This parameter is temporarily unavailable</p>
	AutoScaleUpDelay *int64 `json:"AutoScaleUpDelay,omitnil,omitempty" name:"AutoScaleUpDelay"`

	// <p>This parameter is temporarily invalid</p>
	AutoScaleDownDelay *int64 `json:"AutoScaleDownDelay,omitnil,omitempty" name:"AutoScaleDownDelay"`

	// <p>cpu minimum value. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Maximum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>Minimum value of read-only instance cpu. For the optional range, refer to the API response of DescribeServerlessInstanceSpecs.</p>
	MinRoCpu *float64 `json:"MinRoCpu,omitnil,omitempty" name:"MinRoCpu"`

	// <p>Read-only cpu maximum value of the optional range. For reference, see API response of DescribeServerlessInstanceSpecs.</p>
	MaxRoCpu *float64 `json:"MaxRoCpu,omitnil,omitempty" name:"MaxRoCpu"`

	// <p>Minimum count of read-only nodes</p>
	MinRoCount *int64 `json:"MinRoCount,omitnil,omitempty" name:"MinRoCount"`

	// <p>Maximum number of read-only nodes</p>
	MaxRoCount *int64 `json:"MaxRoCount,omitnil,omitempty" name:"MaxRoCount"`

	// <p>Whether archiving is enabled. Optional range</p><li>yes</li><li>no</li>Default value: yes</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Upgrade type. Default value: upgradeImmediate. Available values: upgradeImmediate - immediately complete the modification; upgradeInMaintain - complete the modification during maintenance window.</p>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// <p>List of security groups to which newly-added read-only instances need to be bound. This only applies to binding security groups to read-only instances generated during the process of this adjustment of policies. Existing instances are not bound.</p>
	SecurityGroupIdsForNewRo []*string `json:"SecurityGroupIdsForNewRo,omitnil,omitempty" name:"SecurityGroupIdsForNewRo"`
}

func (r *ModifyServerlessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServerlessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "AutoScaleUpDelay")
	delete(f, "AutoScaleDownDelay")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "MinRoCpu")
	delete(f, "MaxRoCpu")
	delete(f, "MinRoCount")
	delete(f, "MaxRoCount")
	delete(f, "AutoArchive")
	delete(f, "UpgradeType")
	delete(f, "SecurityGroupIdsForNewRo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServerlessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServerlessStrategyResponseParams struct {
	// <p>Async process id</p>
	//
	// Deprecated: FlowId is deprecated.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServerlessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServerlessStrategyResponseParams `json:"Response"`
}

func (r *ModifyServerlessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServerlessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapBackupCrossRegionConfigRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether cross-region snapshot backup is enabled.
	CrossRegionsEnable *string `json:"CrossRegionsEnable,omitnil,omitempty" name:"CrossRegionsEnable"`

	// Cross-Regional snapshot backup.
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`
}

type ModifySnapBackupCrossRegionConfigRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether cross-region snapshot backup is enabled.
	CrossRegionsEnable *string `json:"CrossRegionsEnable,omitnil,omitempty" name:"CrossRegionsEnable"`

	// Cross-Regional snapshot backup.
	CrossRegions []*string `json:"CrossRegions,omitnil,omitempty" name:"CrossRegions"`
}

func (r *ModifySnapBackupCrossRegionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapBackupCrossRegionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CrossRegionsEnable")
	delete(f, "CrossRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapBackupCrossRegionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapBackupCrossRegionConfigResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySnapBackupCrossRegionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapBackupCrossRegionConfigResponseParams `json:"Response"`
}

func (r *ModifySnapBackupCrossRegionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapBackupCrossRegionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVaultRequestParams struct {
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe name, maximum 255 characters
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Safe description, maximum 1024 characters
	VaultDescribe *string `json:"VaultDescribe,omitnil,omitempty" name:"VaultDescribe"`

	// Backup retention period (seconds). Range: (0, 632448000].
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Encryption key ID, maximum 36 characters
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key type, available values: cloud, custom
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key region, maximum 32 characters
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Is the safe locked
	IsLock *bool `json:"IsLock,omitnil,omitempty" name:"IsLock"`

	// Lock expiry time, format: 2006-01-02 15:04:05, lock time current maximum 15 days
	LockedTime *string `json:"LockedTime,omitnil,omitempty" name:"LockedTime"`

	// Whether it is encrypted
	IsEncryption *bool `json:"IsEncryption,omitnil,omitempty" name:"IsEncryption"`
}

type ModifyVaultRequest struct {
	*tchttp.BaseRequest
	
	// Safe ID
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe name, maximum 255 characters
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Safe description, maximum 1024 characters
	VaultDescribe *string `json:"VaultDescribe,omitnil,omitempty" name:"VaultDescribe"`

	// Backup retention period (seconds). Range: (0, 632448000].
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`

	// Encryption key ID, maximum 36 characters
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key type, available values: cloud, custom
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key region, maximum 32 characters
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Is the safe locked
	IsLock *bool `json:"IsLock,omitnil,omitempty" name:"IsLock"`

	// Lock expiry time, format: 2006-01-02 15:04:05, lock time current maximum 15 days
	LockedTime *string `json:"LockedTime,omitnil,omitempty" name:"LockedTime"`

	// Whether it is encrypted
	IsEncryption *bool `json:"IsEncryption,omitnil,omitempty" name:"IsEncryption"`
}

func (r *ModifyVaultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVaultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VaultId")
	delete(f, "VaultName")
	delete(f, "VaultDescribe")
	delete(f, "BackupSaveSeconds")
	delete(f, "KeyId")
	delete(f, "KeyType")
	delete(f, "KeyRegion")
	delete(f, "IsLock")
	delete(f, "LockedTime")
	delete(f, "IsEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVaultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVaultResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVaultResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVaultResponseParams `json:"Response"`
}

func (r *ModifyVaultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVaultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance group ID
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type ModifyVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`

	// Target IP to be modified
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Target port to be modified
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Database type. Valid values: 
	// <li> MYSQL </li>
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Valid hours of old IPs. If it is set to `0` hours, the IPs will be released immediately.
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

func (r *ModifyVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceGrpId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "DbType")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVipVportResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVipVportResponseParams `json:"Response"`
}

func (r *ModifyVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Module struct {
	// Whether it is supported. Valid values: `yes`, `no`.
	IsDisable *string `json:"IsDisable,omitnil,omitempty" name:"IsDisable"`

	// Module name
	ModuleName *string `json:"ModuleName,omitnil,omitempty" name:"ModuleName"`
}

type MonthDay struct {
	// Month info
	Month *int64 `json:"Month,omitnil,omitempty" name:"Month"`

	// date info
	Day *int64 `json:"Day,omitnil,omitempty" name:"Day"`
}

type NetAddr struct {
	// Private IP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Specifies the private network port number.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Specifies the public network domain name.
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Specifies the public network port number.
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Network type (ro - read-only, rw/ha - read-write).
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// Specifies the subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Specifies the virtual private cloud ID.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Description information
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Public IP address
	WanIP *string `json:"WanIP,omitnil,omitempty" name:"WanIP"`

	// Specifies the public network status.
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// Instance group ID
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`
}

type NewAccount struct {
	// Account name, which can contain 1-16 letters, digits, and underscores. It must begin with a letter and end with a letter or digit.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Password, which can contain 8-64 characters.
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Whether password rotation is enabled (0: turn off; 1: turn on)
	PasswordRotation *int64 `json:"PasswordRotation,omitnil,omitempty" name:"PasswordRotation"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Maximum number of user connections, which cannot be above 10,240.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ObjectTask struct {
	// Task auto-increment ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Status of tasks.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task ID (cluster ID | instance group ID | instance ID).
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// Task type
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`
}

type Objects struct {
	// Database table information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseTables *MigrateObject `json:"DatabaseTables,omitnil,omitempty" name:"DatabaseTables"`
}

// Predefined struct for user
type OfflineClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type OfflineClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *OfflineClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineClusterResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineClusterResponse struct {
	*tchttp.BaseResponse
	Response *OfflineClusterResponseParams `json:"Response"`
}

func (r *OfflineClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type OfflineInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Instance ID array
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

func (r *OfflineInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineInstanceResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineInstanceResponseParams `json:"Response"`
}

func (r *OfflineInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineLibraDBClusterRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type OfflineLibraDBClusterRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *OfflineLibraDBClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineLibraDBClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineLibraDBClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineLibraDBClusterResponseParams struct {
	// flow id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineLibraDBClusterResponse struct {
	*tchttp.BaseResponse
	Response *OfflineLibraDBClusterResponseParams `json:"Response"`
}

func (r *OfflineLibraDBClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineLibraDBClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineLibraDBInstanceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

type OfflineLibraDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Read-only analysis engine instance ID list
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`
}

func (r *OfflineLibraDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineLibraDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineLibraDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineLibraDBInstanceResponseParams struct {
	// task flow id
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineLibraDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineLibraDBInstanceResponseParams `json:"Response"`
}

func (r *OfflineLibraDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineLibraDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OldAddrInfo struct {
	// IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Expect recycle time.
	ReturnTime *string `json:"ReturnTime,omitnil,omitempty" name:"ReturnTime"`
}

// Predefined struct for user
type OpenAIOptimizerRequestParams struct {
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenAIOptimizerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster ID.</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OpenAIOptimizerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAIOptimizerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAIOptimizerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAIOptimizerResponseParams struct {
	// <p>Task flow id.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAIOptimizerResponse struct {
	*tchttp.BaseResponse
	Response *OpenAIOptimizerResponseParams `json:"Response"`
}

func (r *OpenAIOptimizerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAIOptimizerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule (deprecated).
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. true - Full audit; default false - Rule-based audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Frequent log retention period.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule (deprecated).
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and `AuditRuleFilters` are left empty, full audit will be applied.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. true - Full audit; default false - Rule-based audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	delete(f, "AuditAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterPasswordComplexityRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

type OpenClusterPasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password length
	ValidatePasswordLength *int64 `json:"ValidatePasswordLength,omitnil,omitempty" name:"ValidatePasswordLength"`

	// Number of letters
	ValidatePasswordMixedCaseCount *int64 `json:"ValidatePasswordMixedCaseCount,omitnil,omitempty" name:"ValidatePasswordMixedCaseCount"`

	// Number of symbols
	ValidatePasswordSpecialCharCount *int64 `json:"ValidatePasswordSpecialCharCount,omitnil,omitempty" name:"ValidatePasswordSpecialCharCount"`

	// Number of digits
	ValidatePasswordNumberCount *int64 `json:"ValidatePasswordNumberCount,omitnil,omitempty" name:"ValidatePasswordNumberCount"`

	// Password strength. Valid values: `MEDIUM`, `STRONG`.
	ValidatePasswordPolicy *string `json:"ValidatePasswordPolicy,omitnil,omitempty" name:"ValidatePasswordPolicy"`

	// Data dictionary
	ValidatePasswordDictionary []*string `json:"ValidatePasswordDictionary,omitnil,omitempty" name:"ValidatePasswordDictionary"`
}

func (r *OpenClusterPasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ValidatePasswordLength")
	delete(f, "ValidatePasswordMixedCaseCount")
	delete(f, "ValidatePasswordSpecialCharCount")
	delete(f, "ValidatePasswordNumberCount")
	delete(f, "ValidatePasswordPolicy")
	delete(f, "ValidatePasswordDictionary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterPasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterPasswordComplexityResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterPasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterPasswordComplexityResponseParams `json:"Response"`
}

func (r *OpenClusterPasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterPasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterReadOnlyInstanceGroupAccessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group ID.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type OpenClusterReadOnlyInstanceGroupAccessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group ID.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterReadOnlyInstanceGroupAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterReadOnlyInstanceGroupAccessResponseParams struct {
	// Initiate process ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterReadOnlyInstanceGroupAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterReadOnlyInstanceGroupAccessResponseParams `json:"Response"`
}

func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterReadOnlyInstanceGroupAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterTransparentEncryptRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Key type (cloud, custom).
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key Id.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key region.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Whether global encryption is enabled
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

type OpenClusterTransparentEncryptRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Key type (cloud, custom).
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// Key Id.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key region.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// Whether global encryption is enabled
	IsOpenGlobalEncryption *bool `json:"IsOpenGlobalEncryption,omitnil,omitempty" name:"IsOpenGlobalEncryption"`
}

func (r *OpenClusterTransparentEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterTransparentEncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "KeyType")
	delete(f, "KeyId")
	delete(f, "KeyRegion")
	delete(f, "IsOpenGlobalEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenClusterTransparentEncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenClusterTransparentEncryptResponseParams struct {
	// Asynchronous task ID.
	// 
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenClusterTransparentEncryptResponse struct {
	*tchttp.BaseResponse
	Response *OpenClusterTransparentEncryptResponseParams `json:"Response"`
}

func (r *OpenClusterTransparentEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenClusterTransparentEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessRequestParams struct {
	// Please use the [cluster information description](https://www.tencentcloud.com/document/product/1098/40752) to obtain the clusterId.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Please use the [cluster information description](https://www.tencentcloud.com/document/product/1098/40752) to obtain the instanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the designated vpc ID. please use the [query vpc list](https://www.tencentcloud.com/document/product/215/15778) to obtain the vpc ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Specified subnet ID. if vpc ID is set, SubnetId is required. please use [query subnet list](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) to get SubnetId.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// User-Defined port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group ID. use [view security group](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) to obtain the SecurityGroupId.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type OpenReadOnlyInstanceExclusiveAccessRequest struct {
	*tchttp.BaseRequest
	
	// Please use the [cluster information description](https://www.tencentcloud.com/document/product/1098/40752) to obtain the clusterId.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Please use the [cluster information description](https://www.tencentcloud.com/document/product/1098/40752) to obtain the instanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the designated vpc ID. please use the [query vpc list](https://www.tencentcloud.com/document/product/215/15778) to obtain the vpc ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Specified subnet ID. if vpc ID is set, SubnetId is required. please use [query subnet list](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) to get SubnetId.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// User-Defined port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Security group ID. use [view security group](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) to obtain the SecurityGroupId.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Port")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenReadOnlyInstanceExclusiveAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenReadOnlyInstanceExclusiveAccessResponseParams struct {
	// Activation process ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenReadOnlyInstanceExclusiveAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenReadOnlyInstanceExclusiveAccessResponseParams `json:"Response"`
}

func (r *OpenReadOnlyInstanceExclusiveAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenReadOnlyInstanceExclusiveAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanRequestParams struct {
	// Instance group ID
	//
	// Deprecated: InstanceGrpId is deprecated.
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

type OpenWanRequest struct {
	*tchttp.BaseRequest
	
	// Instance group ID
	InstanceGrpId *string `json:"InstanceGrpId,omitnil,omitempty" name:"InstanceGrpId"`
}

func (r *OpenWanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceGrpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanResponseParams struct {
	// Task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenWanResponse struct {
	*tchttp.BaseResponse
	Response *OpenWanResponseParams `json:"Response"`
}

func (r *OpenWanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Package struct {
	// AppID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique ID of the resource package.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource package name.
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Specifies the resource package type.
	// CCU: compute resource package. DISK: storage resource package.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource package region of use.
	// China - common in the chinese mainland. overseas - universally applicable in hong kong (china), macao (china), taiwan (china), and overseas.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Specifies the status of the resource package.
	// creating - indicates that it is in the process of being created.
	// {using} specifies that it is in use.
	// expired-expired;.
	// normal_finish - specifies that it is used up.
	// `Apply_refund`: apply for a refund.
	// Specifies that the fee has been refunded.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Total resource package quantity.
	PackageTotalSpec *float64 `json:"PackageTotalSpec,omitnil,omitempty" name:"PackageTotalSpec"`

	// Used amount of resource package.
	PackageUsedSpec *float64 `json:"PackageUsedSpec,omitnil,omitempty" name:"PackageUsedSpec"`

	// Whether there is inventory surplus.
	HasQuota *bool `json:"HasQuota,omitnil,omitempty" name:"HasQuota"`

	// Specifies the bound instance information.
	BindInstanceInfos []*BindInstanceInfo `json:"BindInstanceInfos,omitnil,omitempty" name:"BindInstanceInfos"`

	// Specifies the effective time: 2022-07-01 00:00:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the expiration time: 2022-08-01 00:00:00.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Information of the instance historically bound (now unbound) to the resource pack.
	HistoryBindResourceInfos []*BindInstanceInfo `json:"HistoryBindResourceInfos,omitnil,omitempty" name:"HistoryBindResourceInfos"`
}

type PackageDetail struct {
	// AppId account ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique ID of the resource package.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Successfully deduct capacity.
	SuccessDeductSpec *float64 `json:"SuccessDeductSpec,omitnil,omitempty" name:"SuccessDeductSpec"`

	// The used capacity of the resource package up to the present.
	PackageTotalUsedSpec *float64 `json:"PackageTotalUsedSpec,omitnil,omitempty" name:"PackageTotalUsedSpec"`

	// Deduction start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Deduction end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Extension Information
	ExtendInfo *string `json:"ExtendInfo,omitnil,omitempty" name:"ExtendInfo"`
}

type PackagePriority struct {
	// The resource pack whose deduction priority needs to be customized.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Custom deduction priority.
	DeductionPriority *int64 `json:"DeductionPriority,omitnil,omitempty" name:"DeductionPriority"`
}

type ParamDetail struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter type. Valid values:  `integer`, `enum`, `float`, `string`, `func`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Whether `func` is supported. Valid values: `true` (supported), `false` (not supported).
	SupportFunc *bool `json:"SupportFunc,omitnil,omitempty" name:"SupportFunc"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Whether to restart the database for the modified parameters to take effect. Valid values:  `0` (no), `1` (yes).
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Optional enumerated values of the parameter. if it is a non-enumerated value, it is empty.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Valid values: `1` (global parameter),  `0`  (non-global parameter).
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// The match type. Valid value: `multiVal`.
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// Match values, which will be separated by comma when `MatchType` is `multiVal`.
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// true - indicates a formula. false - indicates it is not a formula.
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// Specifies that when the parameter is set as a formula, Func returns the set formula content.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// Whether the parameter is modifiable.
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// The default formula style of parameters that support formulas.
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// If the parameter is of type enum/string/bool, the available options list.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Maximum value when parameter type is `float` or `integer`.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value when parameter type is `float` or `integer`.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether to restart the instance for the modified parameters to take effect.
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter type: `integer`, `float`, `string`, `enum`, `bool`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Match type. Regex can be used when parameter type is `string`. Valid value: `multiVal`.
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// Match values, which will be separated by semicolon when match type is `multiVal`.
	MatchValue *string `json:"MatchValue,omitnil,omitempty" name:"MatchValue"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether it is a global parameter.
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// Whether the parameter is modifiable.
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// Whether it is a function.
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// Function.
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// The default formula style of parameters that support formulas.
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamItem struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// New value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Original value
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`
}

type ParamItemDetail struct {
	// <p>Current value</p>
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// <p>Default value</p>
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// <p>Optional enumerated value of the parameter. If it is a non-enumerated value, it is empty.</p>
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// <p>1: Global parameter, 0: Non-global parameter</p>
	IsGlobal *int64 `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// <p>Maximum value</p>
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// <p>Minimum value</p>
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// <p>After modifying parameters, whether database restart is required to take effect. 0-no restart required, 1-restart required.</p>
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// <p>Parameter name</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <p>Parameter type: integer, enum, float, string, func</p>
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// <p>Whether the parameter is modifiable</p>
	ModifiableInfo *ModifiableInfo `json:"ModifiableInfo,omitnil,omitempty" name:"ModifiableInfo"`

	// <p>Parameter description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Whether the type is formula</p>
	IsFunc *bool `json:"IsFunc,omitnil,omitempty" name:"IsFunc"`

	// <p>Parameter configuration formula</p>
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`

	// <p>Default formula style for parameters that support formulas</p>
	FuncPattern *string `json:"FuncPattern,omitnil,omitempty" name:"FuncPattern"`
}

type ParamItemInfo struct {
	// Parameter name.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// New parameter value.
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// Old parameter value.
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// Parameter formula.
	ValueFunction *string `json:"ValueFunction,omitnil,omitempty" name:"ValueFunction"`
}

type ParamTemplateListInfo struct {
	// Parameter template ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// Engine version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Database Type. Valid values: `NORMAL`, `SERVERLESS`.
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Parameter template details
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamInfoSet []*TemplateParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`
}

// Predefined struct for user
type PauseServerlessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
}

type PauseServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Whether to pause forcibly and ignore the current user connections. Valid values: `0` (no), `1` (yes). Default value: `1`
	ForcePause *int64 `json:"ForcePause,omitnil,omitempty" name:"ForcePause"`
}

func (r *PauseServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ForcePause")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseServerlessResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseServerlessResponse struct {
	*tchttp.BaseResponse
	Response *PauseServerlessResponseParams `json:"Response"`
}

func (r *PauseServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyRule struct {
	// Policy, which can be `ACCEPT` or `DROP`
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Source IP or source IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol, such as UDP and TCP
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// Protocol port ID or protocol port group ID.
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// IP address ID or IP address group ID.
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ProxyConfig struct {
	// <p>Number of database proxy group nodes. This parameter is no longer recommended. Recommend using ProxyZones.</p>
	//
	// Deprecated: ProxyCount is deprecated.
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// <p>cpu cores</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Memory</p>
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// <p>Connection pool type: SessionConnectionPool (session-level connection pool)</p>
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// <p>Whether the connection pool is enabled, yes-enable, no-disable</p>
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// <p>Connection pool threshold: Measurement unit (seconds)</p>
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// <p>description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Database node information (this parameter is used in combination with ProxyCount and either one must be input)</p>
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type ProxyConfigInfo struct {
	// Number of database proxy group nodes. this parameter is no longer recommended. recommend using ProxyZones.
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Specifies the memory.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Database node information (this parameter is used in combination with ProxyCount, either one must be manually input).
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type ProxyConnectionPoolInfo struct {
	// Specifies the persistence threshold of the connection pool. the unit is seconds.
	ConnectionPoolTimeOut *int64 `json:"ConnectionPoolTimeOut,omitnil,omitempty" name:"ConnectionPoolTimeOut"`

	// Whether the connection pool is enabled.
	OpenConnectionPool *string `json:"OpenConnectionPool,omitnil,omitempty" name:"OpenConnectionPool"`

	// Specifies the connection pool type. valid values: SessionConnectionPool (session-level connection pool).
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`
}

type ProxyEndPointConfigInfo struct {
	// Specifies the ID of the VPC network it belongs to.
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// Subnet ID.
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// Security group id array.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Weight mode: system-system allocation, custom-custom.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Specifies whether to automatically add a read-only instance. valid values: yes, no.
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Read-Write attribute. valid values: READWRITE, READONLY.
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// Weight information.
	InstanceNameWeights []*InstanceNameWeight `json:"InstanceNameWeights,omitnil,omitempty" name:"InstanceNameWeights"`
}

type ProxyGroup struct {
	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Number of nodes in the proxy group
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`

	// Database proxy group status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// User AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Specifies that a read-write node activates the database proxy.
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`
}

type ProxyGroupInfo struct {
	// Database proxy group.
	ProxyGroup *ProxyGroup `json:"ProxyGroup,omitnil,omitempty" name:"ProxyGroup"`

	// Database proxy group read-write separation information.
	ProxyGroupRwInfo *ProxyGroupRwInfo `json:"ProxyGroupRwInfo,omitnil,omitempty" name:"ProxyGroupRwInfo"`

	// Node information of the database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProxyNodes []*ProxyNodeInfo `json:"ProxyNodes,omitnil,omitempty" name:"ProxyNodes"`

	// Database proxy connection pool information.
	ConnectionPool *ProxyConnectionPoolInfo `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Network information for database proxy
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetAddrInfos []*NetAddr `json:"NetAddrInfos,omitnil,omitempty" name:"NetAddrInfos"`

	// Database proxy task set.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type ProxyGroupRwInfo struct {
	// <p>Consistency Type eventual-final consistency,global-global consistency,session-session consistency</p>
	ConsistencyType *string `json:"ConsistencyType,omitnil,omitempty" name:"ConsistencyType"`

	// <p>Consistency timeout</p>
	ConsistencyTimeOut *int64 `json:"ConsistencyTimeOut,omitnil,omitempty" name:"ConsistencyTimeOut"`

	// <p>Weight mode system-system-assigned, custom-custom</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Whether fault migration is enabled</p>
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Automatically add read-only instance, yes-yes, no-no</p>
	AutoAddRo *string `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Instance weight array</p>
	InstanceWeights []*ProxyInstanceWeight `json:"InstanceWeights,omitnil,omitempty" name:"InstanceWeights"`

	// <p>Whether to enable the read-write node, yes - enable, no - disable</p>
	OpenRw *string `json:"OpenRw,omitnil,omitempty" name:"OpenRw"`

	// <p>Read-write attribute, value range: READWRITE, READONLY</p>
	RwType *string `json:"RwType,omitnil,omitempty" name:"RwType"`

	// <p>Transaction split</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Connection mode, available values: balance, nearby</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Whether to treat the libra node as an ordinary RO node</p>
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// <p>Whether to forward to other nodes when a libra node fault occurs</p>
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`

	// <p>Auto load</p><p>Enumeration value:</p><ul><li>static: Static load</li><li>dynamic: Dynamic load</li></ul>
	LoadBalanceMode *string `json:"LoadBalanceMode,omitnil,omitempty" name:"LoadBalanceMode"`
}

type ProxyInstanceWeight struct {
	// InstanID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance weight
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ProxyNodeInfo struct {
	// Database proxy node ID.
	ProxyNodeId *string `json:"ProxyNodeId,omitnil,omitempty" name:"ProxyNodeId"`

	// Current node connections, which is not returned by the `DescribeProxyNodes` API.
	ProxyNodeConnections *int64 `json:"ProxyNodeConnections,omitnil,omitempty" name:"ProxyNodeConnections"`

	// CPU of the database proxy node.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory of the database proxy node.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Status of the database proxy node.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Database proxy group ID.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// User AppID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Database proxy node name.
	OssProxyNodeName *string `json:"OssProxyNodeName,omitnil,omitempty" name:"OssProxyNodeName"`
}

type ProxySpec struct {
	// Number of database proxy CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database proxy memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`
}

type ProxyVersionInfo struct {
	// proxy version number.
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// Version description: GA: stable version. BETA: BETA version. DEPRECATED: outdated.
	ProxyVersionType *string `json:"ProxyVersionType,omitnil,omitempty" name:"ProxyVersionType"`
}

type ProxyZone struct {
	// AZ of the proxy node
	ProxyNodeZone *string `json:"ProxyNodeZone,omitnil,omitempty" name:"ProxyNodeZone"`

	// The number of proxy nodes
	ProxyNodeCount *int64 `json:"ProxyNodeCount,omitnil,omitempty" name:"ProxyNodeCount"`
}

type QueryFilter struct {
	// Search field. Valid values: "InstanceId", "ProjectId", "InstanceName", "Vip"
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Search string
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to use exact match
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// Search field
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Operator
	//
	// Deprecated: Operator is deprecated.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type QueryParamFilter struct {
	// Search field. Currently supports: ProxyGroupId
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Search string
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Whether to use exact match
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type QuerySimpleFilter struct {
	// Field name
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// Field value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Fuzzy matching. valid values: true (yes), false (no).
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type RedoLogItem struct {
	// Filename.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Backup time
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// redoLogId
	RedoLogId *int64 `json:"RedoLogId,omitnil,omitempty" name:"RedoLogId"`

	// Cross-Region information.
	RedoCrossRegions []*BackupRegionAndIds `json:"RedoCrossRegions,omitnil,omitempty" name:"RedoCrossRegions"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Completion time.
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Safe info
	VaultInfos []*VaultInfo `json:"VaultInfos,omitnil,omitempty" name:"VaultInfos"`

	// Backup delivery status
	CopyStatus *string `json:"CopyStatus,omitnil,omitempty" name:"CopyStatus"`

	// Encryption key
	EncryptKeyId *string `json:"EncryptKeyId,omitnil,omitempty" name:"EncryptKeyId"`

	// Key region for encryption
	EncryptRegion *string `json:"EncryptRegion,omitnil,omitempty" name:"EncryptRegion"`
}

// Predefined struct for user
type RefundResourcePackageRequestParams struct {
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
}

type RefundResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a resource pack
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`
}

func (r *RefundResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundResourcePackageResponseParams struct {
	// Each item has only one `dealName`, through which you can ensure the idempotency of the delivery API.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefundResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *RefundResourcePackageResponseParams `json:"Response"`
}

func (r *RefundResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInstanceSpecInfo struct {
	// Number of CPU cores
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Minimum storage size
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinStorageSize *int64 `json:"MinStorageSize,omitnil,omitempty" name:"MinStorageSize"`

	// Maximum storage size
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// Whether there is inventory
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Instance type
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Storage type
	// Note: This field may return null, indicating that no valid values can be obtained.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// minimum replica count
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinReplicaNum *uint64 `json:"MinReplicaNum,omitnil,omitempty" name:"MinReplicaNum"`

	// Maximum number of replicas
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxReplicaNum *uint64 `json:"MaxReplicaNum,omitnil,omitempty" name:"MaxReplicaNum"`

	// Availability zone inventory information list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneStockInfos []*ZoneStockInfo4Libra `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`
}

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *ReloadBalanceProxyNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadBalanceProxyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReloadBalanceProxyNodeResponse struct {
	*tchttp.BaseResponse
	Response *ReloadBalanceProxyNodeResponseParams `json:"Response"`
}

func (r *ReloadBalanceProxyNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveClusterSlaveZoneRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

type RemoveClusterSlaveZoneRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Replica AZ
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

func (r *RemoveClusterSlaveZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SlaveZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveClusterSlaveZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveClusterSlaveZoneResponseParams struct {
	// Async FlowId
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveClusterSlaveZoneResponse struct {
	*tchttp.BaseResponse
	Response *RemoveClusterSlaveZoneResponseParams `json:"Response"`
}

func (r *RemoveClusterSlaveZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveClusterSlaveZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewClustersRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Renewal period.
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit, such as y, m, d, h, i, and s.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// "Transaction mode. 	0 - place an order and pay; 1 - place an order."
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

type RenewClustersRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Renewal period.
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit, such as y, m, d, h, i, and s.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// "Transaction mode. 	0 - place an order and pay; 1 - place an order."
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

func (r *RenewClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewClustersResponseParams struct {
	// Prepaid total order number.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Refund order number.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Frozen flow. One frozen flow is activated at a time.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Delivery resource id list corresponding to each order number.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// List of Cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewClustersResponse struct {
	*tchttp.BaseResponse
	Response *RenewClustersResponseParams `json:"Response"`
}

func (r *RenewClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLibraDBClustersRequestParams struct {
	// Time interval
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Order mode
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

type RenewLibraDBClustersRequest struct {
	*tchttp.BaseRequest
	
	// Time interval
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Order mode
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`
}

func (r *RenewLibraDBClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLibraDBClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ClusterId")
	delete(f, "DealMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewLibraDBClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLibraDBClustersResponseParams struct {
	// Prepaid Total Order Number
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Freeze transaction
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Order name
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Resource ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Cluster ID.
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewLibraDBClustersResponse struct {
	*tchttp.BaseResponse
	Response *RenewLibraDBClustersResponseParams `json:"Response"`
}

func (r *RenewLibraDBClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLibraDBClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplayInstanceAuditLogRequestParams struct {
	// Source cluster id.
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`

	// Source instance id.
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// Target cluster id.
	// Specifies the target cluster must be a cluster cloned from the original cluster within three days.
	TargetClusterId *string `json:"TargetClusterId,omitnil,omitempty" name:"TargetClusterId"`

	// Target instance id.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Username. host must be % username.
	TargetUserName *string `json:"TargetUserName,omitnil,omitempty" name:"TargetUserName"`

	// Password.
	TargetPassword *string `json:"TargetPassword,omitnil,omitempty" name:"TargetPassword"`

	// Start time. time format: yyyy-DD-mm hh:mm:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the end time in the time format yyyy-DD-mm hh:mm:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ReplayInstanceAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// Source cluster id.
	SourceClusterId *string `json:"SourceClusterId,omitnil,omitempty" name:"SourceClusterId"`

	// Source instance id.
	SourceInstanceId *string `json:"SourceInstanceId,omitnil,omitempty" name:"SourceInstanceId"`

	// Target cluster id.
	// Specifies the target cluster must be a cluster cloned from the original cluster within three days.
	TargetClusterId *string `json:"TargetClusterId,omitnil,omitempty" name:"TargetClusterId"`

	// Target instance id.
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// Username. host must be % username.
	TargetUserName *string `json:"TargetUserName,omitnil,omitempty" name:"TargetUserName"`

	// Password.
	TargetPassword *string `json:"TargetPassword,omitnil,omitempty" name:"TargetPassword"`

	// Start time. time format: yyyy-DD-mm hh:mm:ss.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Specifies the end time in the time format yyyy-DD-mm hh:mm:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ReplayInstanceAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplayInstanceAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceClusterId")
	delete(f, "SourceInstanceId")
	delete(f, "TargetClusterId")
	delete(f, "TargetInstanceId")
	delete(f, "TargetUserName")
	delete(f, "TargetPassword")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplayInstanceAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplayInstanceAuditLogResponseParams struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplayInstanceAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *ReplayInstanceAuditLogResponseParams `json:"Response"`
}

func (r *ReplayInstanceAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplayInstanceAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationObject struct {
	// Source instance type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInstanceType *string `json:"SrcInstanceType,omitnil,omitempty" name:"SrcInstanceType"`

	// Source Cluster Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcClusterId *string `json:"SrcClusterId,omitnil,omitempty" name:"SrcClusterId"`

	// Source instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Copy task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicationJobId *string `json:"ReplicationJobId,omitnil,omitempty" name:"ReplicationJobId"`

	// sync object details
	// Note: This field may return null, indicating that no valid values can be obtained.
	MigrateObjects *MigrateOpt `json:"MigrateObjects,omitnil,omitempty" name:"MigrateObjects"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Database account name
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// New password of the database account
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Host. Default value: `%`
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ClusterId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetLibraDBClusterAccountPasswordRequestParams struct {
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password.
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Account
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetLibraDBClusterAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Password.
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// Account
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Encryption method
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ResetLibraDBClusterAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetLibraDBClusterAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AccountPassword")
	delete(f, "AccountName")
	delete(f, "EncryptMethod")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetLibraDBClusterAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetLibraDBClusterAccountPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetLibraDBClusterAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetLibraDBClusterAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetLibraDBClusterAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetLibraDBClusterAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePackage struct {
	// The unique ID of the resource package.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Resource package type: CCU: compute resource package.
	// DISK: storage resource package.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Deduction priority of the current resource package bound to the current instance.
	DeductionPriority *int64 `json:"DeductionPriority,omitnil,omitempty" name:"DeductionPriority"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

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
type RestartLibraDBInstanceRequestParams struct {
	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartLibraDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Read-only analysis engine instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *RestartLibraDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLibraDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartLibraDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartLibraDBInstanceResponseParams struct {
	// Asynchronous task ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartLibraDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartLibraDBInstanceResponseParams `json:"Response"`
}

func (r *RestartLibraDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLibraDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ResumeServerlessRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *ResumeServerlessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeServerlessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeServerlessResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeServerlessResponse struct {
	*tchttp.BaseResponse
	Response *ResumeServerlessResponseParams `json:"Response"`
}

func (r *ResumeServerlessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeServerlessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoAddr struct {
	// IP address
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type RollBackClusterRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Rollback policy. timeRollback - roll back by time point; snapRollback - roll back by backup file.
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// Backup file ID. This parameter is required when the rollback policy is rolling back by backup file.
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// Expected rollback time. This parameter is required when the rollback policy is timeRollback (roll back by time point).
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// Expected threshold (deprecated).
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// List of rollback databases.
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// List of rollback databases and tables.
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// Mode of rolling back by time point. full: normal; db: fast; table: ultra-fast (the default value is normal).
	RollbackMode *string `json:"RollbackMode,omitnil,omitempty" name:"RollbackMode"`

	// Safe id
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`
}

type RollBackClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Rollback policy. timeRollback - roll back by time point; snapRollback - roll back by backup file.
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// Backup file ID. This parameter is required when the rollback policy is rolling back by backup file.
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// Expected rollback time. This parameter is required when the rollback policy is timeRollback (roll back by time point).
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// Expected threshold (deprecated).
	ExpectTimeThresh *uint64 `json:"ExpectTimeThresh,omitnil,omitempty" name:"ExpectTimeThresh"`

	// List of rollback databases.
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// List of rollback databases and tables.
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// Mode of rolling back by time point. full: normal; db: fast; table: ultra-fast (the default value is normal).
	RollbackMode *string `json:"RollbackMode,omitnil,omitempty" name:"RollbackMode"`

	// Safe id
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`
}

func (r *RollBackClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollBackClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RollbackStrategy")
	delete(f, "RollbackId")
	delete(f, "ExpectTime")
	delete(f, "ExpectTimeThresh")
	delete(f, "RollbackDatabases")
	delete(f, "RollbackTables")
	delete(f, "RollbackMode")
	delete(f, "VaultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollBackClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollBackClusterResponseParams struct {
	// Task flow ID.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollBackClusterResponse struct {
	*tchttp.BaseResponse
	Response *RollBackClusterResponseParams `json:"Response"`
}

func (r *RollBackClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollBackClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollbackData struct {
	// Instance CPU.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance memory.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Cluster storage upper limit.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Original cluster ID.
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// Original cluster name.
	OriginalClusterName *string `json:"OriginalClusterName,omitnil,omitempty" name:"OriginalClusterName"`

	// Rollback method.
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// Snapshot time.
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`

	// Roll back to the Serverless cluster with minimum CPU.
	MinCpu *int64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// Roll back to the Serverless cluster with maximum CPU.
	MaxCpu *int64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// Snapshot ID.
	SnapShotId *uint64 `json:"SnapShotId,omitnil,omitempty" name:"SnapShotId"`

	// Rollback database.
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// Rollback data table.
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// Specifies the backup file name.
	BackupFileName *string `json:"BackupFileName,omitnil,omitempty" name:"BackupFileName"`

	// Rollback process.
	RollbackProcess *RollbackProcessInfo `json:"RollbackProcess,omitnil,omitempty" name:"RollbackProcess"`
}

type RollbackDatabase struct {
	// Old database name.
	OldDatabase *string `json:"OldDatabase,omitnil,omitempty" name:"OldDatabase"`

	// New database name.
	NewDatabase *string `json:"NewDatabase,omitnil,omitempty" name:"NewDatabase"`
}

type RollbackInstanceInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// VPC information
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet information.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vip information.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Specifies the vport information.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CPU Size.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory Size.
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Storage size.
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`
}

type RollbackProcessInfo struct {
	// Is it possible to switch vip.
	IsVipSwitchable *bool `json:"IsVipSwitchable,omitnil,omitempty" name:"IsVipSwitchable"`

	// The exchangeable time of vip.
	VipSwitchableTime *string `json:"VipSwitchableTime,omitnil,omitempty" name:"VipSwitchableTime"`

	// Swap instance list.
	ExchangeInstanceInfoList []*ExchangeInstanceInfo `json:"ExchangeInstanceInfoList,omitnil,omitempty" name:"ExchangeInstanceInfoList"`

	// Swap RO group list.
	ExchangeRoGroupInfoList []*ExchangeRoGroupInfo `json:"ExchangeRoGroupInfoList,omitnil,omitempty" name:"ExchangeRoGroupInfoList"`

	// Current step.
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// Current step progress.
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil,omitempty" name:"CurrentStepProgress"`

	// Remaining time of the current step.
	CurrentStepRemainingTime *string `json:"CurrentStepRemainingTime,omitnil,omitempty" name:"CurrentStepRemainingTime"`
}

type RollbackRoGroupInfo struct {
	// Instance group ID.
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// VPC information.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet information.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vip information.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Specifies the vport information.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type RollbackTable struct {
	// Database name.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Database table.
	Tables []*RollbackTableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type RollbackTableInfo struct {
	// Old table name.
	OldTable *string `json:"OldTable,omitnil,omitempty" name:"OldTable"`

	// New table name.
	NewTable *string `json:"NewTable,omitnil,omitempty" name:"NewTable"`
}

type RollbackTimeRange struct {
	// Start time
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// End time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`
}

// Predefined struct for user
type RollbackToNewClusterRequestParams struct {
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>During rollback, input the source cluster ID to search for the source poolId</p>
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// <p>VPC network ID</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Cluster name, length less than 64 characters. Each character value ranges from uppercase/lowercase letters, digits, to special symbols ('-', '_', '.').</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Snapshot rollback means snapshotId; point-in-time rollback means queryId. A value of 0 indicates requirement to judge whether the time point is valid.</p>
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// <p>Point-in-time rollback, specified time; snapshot rollback, snapshot time</p>
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// <p>Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>tag Array information that should be bound for cluster creation</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>DB type<br>Selectable when DbType is MYSQL (default NORMAL):</p><li>NORMAL</li><li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Required when DbMode is SEVERLESS<br>Minimum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Required when DbMode is SEVERLESS:<br>Maximum value of cpu. For the optional range, see the API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>When DbMode is SEVERLESS, whether to automatically pause within specified clusters. Optional range</p><li>yes</li><li>no</li>Default value: yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>When DbMode is SEVERLESS, specify the delay for Cluster Auto-Pause in seconds, optional range [600,691200]<br>Default value: 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>Security group id array</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Alarm policy Id array</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// <p>Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names, 1-case-insensitive, 0-case-sensitive.</p>
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// <p>Parameter template ID. The parameter template ID can be obtained through querying parameter template information DescribeParamTemplates.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Instance initialization configuration information is mainly used to select different specification instances during cluster purchase.</p>
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// <p>0-Place order and pay 1-Placing order</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Pay-per-compute-node model: 0-Pay-As-You-Go, 1-Prepayment</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Time</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>Unit</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Rollback database info</p>
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// <p>Roll back table information</p>
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// <p>Original ro instance information</p>
	OriginalROInstanceList []*string `json:"OriginalROInstanceList,omitnil,omitempty" name:"OriginalROInstanceList"`

	// <p>Project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Whether to enable archive. Optional range <li>yes</li><li>no</li> Default value: yes</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Whether to restore from the saved backup</p>
	FromSaveBackup *bool `json:"FromSaveBackup,omitnil,omitempty" name:"FromSaveBackup"`
}

type RollbackToNewClusterRequest struct {
	*tchttp.BaseRequest
	
	// <p>AZ.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>During rollback, input the source cluster ID to search for the source poolId</p>
	OriginalClusterId *string `json:"OriginalClusterId,omitnil,omitempty" name:"OriginalClusterId"`

	// <p>VPC network ID</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Cluster name, length less than 64 characters. Each character value ranges from uppercase/lowercase letters, digits, to special symbols ('-', '_', '.').</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Snapshot rollback means snapshotId; point-in-time rollback means queryId. A value of 0 indicates requirement to judge whether the time point is valid.</p>
	RollbackId *uint64 `json:"RollbackId,omitnil,omitempty" name:"RollbackId"`

	// <p>Point-in-time rollback, specified time; snapshot rollback, snapshot time</p>
	ExpectTime *string `json:"ExpectTime,omitnil,omitempty" name:"ExpectTime"`

	// <p>Whether to automatically select a voucher. 1: Yes; 0: No. Default is 0.</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>tag Array information that should be bound for cluster creation</p>
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>DB type<br>Selectable when DbType is MYSQL (default NORMAL):</p><li>NORMAL</li><li>SERVERLESS</li>
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// <p>Required when DbMode is SEVERLESS<br>Minimum value of cpu. For optional range, see API response of DescribeServerlessInstanceSpecs</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Required when DbMode is SEVERLESS:<br>Maximum value of cpu. For the optional range, see the API response of DescribeServerlessInstanceSpecs.</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>When DbMode is SEVERLESS, whether to automatically pause within specified clusters. Optional range</p><li>yes</li><li>no</li>Default value: yes
	AutoPause *string `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// <p>When DbMode is SEVERLESS, specify the delay for Cluster Auto-Pause in seconds, optional range [600,691200]<br>Default value: 600</p>
	AutoPauseDelay *int64 `json:"AutoPauseDelay,omitnil,omitempty" name:"AutoPauseDelay"`

	// <p>Security group id array</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Alarm policy Id array</p>
	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitnil,omitempty" name:"AlarmPolicyIds"`

	// <p>Parameter array, temporarily supports character_set_server (utf8|latin1|gbk|utf8mb4), lower_case_table_names, 1-case-insensitive, 0-case-sensitive.</p>
	ClusterParams []*ParamItem `json:"ClusterParams,omitnil,omitempty" name:"ClusterParams"`

	// <p>Parameter template ID. The parameter template ID can be obtained through querying parameter template information DescribeParamTemplates.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Instance initialization configuration information is mainly used to select different specification instances during cluster purchase.</p>
	InstanceInitInfos []*InstanceInitInfo `json:"InstanceInitInfos,omitnil,omitempty" name:"InstanceInitInfos"`

	// <p>0-Place order and pay 1-Placing order</p>
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// <p>Pay-per-compute-node model: 0-Pay-As-You-Go, 1-Prepayment</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Time</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>Unit</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>Rollback database info</p>
	RollbackDatabases []*RollbackDatabase `json:"RollbackDatabases,omitnil,omitempty" name:"RollbackDatabases"`

	// <p>Roll back table information</p>
	RollbackTables []*RollbackTable `json:"RollbackTables,omitnil,omitempty" name:"RollbackTables"`

	// <p>Original ro instance information</p>
	OriginalROInstanceList []*string `json:"OriginalROInstanceList,omitnil,omitempty" name:"OriginalROInstanceList"`

	// <p>Project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Whether to enable archive. Optional range <li>yes</li><li>no</li> Default value: yes</p>
	AutoArchive *string `json:"AutoArchive,omitnil,omitempty" name:"AutoArchive"`

	// <p>Whether to restore from the saved backup</p>
	FromSaveBackup *bool `json:"FromSaveBackup,omitnil,omitempty" name:"FromSaveBackup"`
}

func (r *RollbackToNewClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackToNewClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "OriginalClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ClusterName")
	delete(f, "RollbackId")
	delete(f, "ExpectTime")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTags")
	delete(f, "DbMode")
	delete(f, "MinCpu")
	delete(f, "MaxCpu")
	delete(f, "AutoPause")
	delete(f, "AutoPauseDelay")
	delete(f, "SecurityGroupIds")
	delete(f, "AlarmPolicyIds")
	delete(f, "ClusterParams")
	delete(f, "ParamTemplateId")
	delete(f, "InstanceInitInfos")
	delete(f, "DealMode")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "RollbackDatabases")
	delete(f, "RollbackTables")
	delete(f, "OriginalROInstanceList")
	delete(f, "ProjectId")
	delete(f, "AutoArchive")
	delete(f, "FromSaveBackup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackToNewClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackToNewClusterResponseParams struct {
	// <p>Frozen transaction ID</p>
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// <p>Order ID</p>
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// <p>Resource ID list (This field is no longer maintained. Please use the dealNames field to query the API DescribeResourcesByDealName to obtain resource IDs.)</p>
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>Cluster ID list (This field is no longer maintained. Please use the dealNames field to query the API DescribeResourcesByDealName to obtain cluster IDs.)</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>Large order number</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackToNewClusterResponse struct {
	*tchttp.BaseResponse
	Response *RollbackToNewClusterResponseParams `json:"Response"`
}

func (r *RollbackToNewClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackToNewClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleFilters struct {
	// Filter parameter name of the audit rule. Valid values: `host` (client IP), `user` (database account), `dbName` (database name), `sqlType` (SQL type), `sql` (SQL statement).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter match type of the audit rule. Valid values: `INC` (including), `EXC` (excluding), `EQS` (equal to), `NEQ` (not equal to).
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filter match value of the audit rule
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RuleTemplateInfo struct {
	// Rule template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// The rule content.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Alarm level. valid values: 1 (low risk), 2 (medium risk), 3 (high risk).
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. 0 - no alert, 1 - alert.
	AlarmPolicy *int64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Rule description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type SalePackageSpec struct {
	// Resource package region of use.
	PackageRegion *string `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// Specifies the resource package type.
	// CCU - compute resource package.
	// Storage resource package.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Resource pack version.
	// base - basic version; common - general version; enterprise - enterprise edition.
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// Minimum number of resources in the current version of the resource package. compute resource unit: pieces; storage resource: GB.
	MinPackageSpec *float64 `json:"MinPackageSpec,omitnil,omitempty" name:"MinPackageSpec"`

	// Specifies the maximum number of resources in the current version of the resource package. valid values: compute resource unit: pieces; storage resource: GB.
	MaxPackageSpec *float64 `json:"MaxPackageSpec,omitnil,omitempty" name:"MaxPackageSpec"`

	// Specifies the resource pack validity period. the measurement unit is day.
	ExpireDay *int64 `json:"ExpireDay,omitnil,omitempty" name:"ExpireDay"`
}

type SaleRegion struct {
	// Region name
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Numeric ID of a region
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// List of purchasable AZs
	ZoneSet []*SaleZone `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// Engine type
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Supported modules in a region
	Modules []*Module `json:"Modules,omitnil,omitempty" name:"Modules"`
}

type SaleZone struct {
	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Numeric ID of an AZ
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// AZ name
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// Whether serverless cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportServerless *int64 `json:"IsSupportServerless,omitnil,omitempty" name:"IsSupportServerless"`

	// Whether standard cluster is supported. Valid values: <br>
	// `0`: No<br>
	// `1`: Yes
	IsSupportNormal *int64 `json:"IsSupportNormal,omitnil,omitempty" name:"IsSupportNormal"`

	// Physical zone
	PhysicalZone *string `json:"PhysicalZone,omitnil,omitempty" name:"PhysicalZone"`

	// Whether the user has AZ permission.
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`

	// Whether it is a full-linkage RDMA AZ.
	IsWholeRdmaZone *string `json:"IsWholeRdmaZone,omitnil,omitempty" name:"IsWholeRdmaZone"`

	// Specifies whether a newly purchased cluster is allowed in the current availability zone. valid values: 1 (allowed), 0 (not allowed).
	IsSupportCreateCluster *int64 `json:"IsSupportCreateCluster,omitnil,omitempty" name:"IsSupportCreateCluster"`
}

type SaveBackupClusterInfo struct {
	// Backup id of the deceased.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Backup time
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Db type (NORMAL, SERVERLESS).
	DbMode *string `json:"DbMode,omitnil,omitempty" name:"DbMode"`

	// Cluster status.
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// Task list.
	Tasks []*ObjectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

// Predefined struct for user
type SearchClusterDatabasesRequestParams struct {
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
}

type SearchClusterDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// The cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Whether to search exactly
	// Valid values: `0` (fuzzy search), `1` (exact search). 
	// Default value: `0`.
	MatchType *int64 `json:"MatchType,omitnil,omitempty" name:"MatchType"`
}

func (r *SearchClusterDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "MatchType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterDatabasesResponseParams struct {
	// Database List
	// Note: This field may return null, indicating that no valid values can be obtained.
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchClusterDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterDatabasesResponseParams `json:"Response"`
}

func (r *SearchClusterDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type SearchClusterTablesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Data table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Data table type. Valid values:
	// `view`: Only return to view,
	// `base_table`: Only return to basic table,
	// `all`: Return to view and table.
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`
}

func (r *SearchClusterTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClusterTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClusterTablesResponseParams struct {
	// Data table list.
	Tables []*DatabaseTables `json:"Tables,omitnil,omitempty" name:"Tables"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchClusterTablesResponse struct {
	*tchttp.BaseResponse
	Response *SearchClusterTablesResponseParams `json:"Response"`
}

func (r *SearchClusterTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClusterTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*PolicyRule `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*PolicyRule `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

type ServerlessSpec struct {
	// <p>cpu minimum value</p>
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// <p>Maximum value of cpu</p>
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// <p>Maximum storage space</p>
	MaxStorageSize *int64 `json:"MaxStorageSize,omitnil,omitempty" name:"MaxStorageSize"`

	// <p>Is the default specification</p>
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// <p>Whether there is inventory</p>
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// <p>Inventory quantity</p>
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// <p>Availability zone inventory information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneStockInfos []*ServerlessZoneStockInfo `json:"ZoneStockInfos,omitnil,omitempty" name:"ZoneStockInfos"`
}

type ServerlessZoneStockInfo struct {
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specifies the stored amount.
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// Whether it contains inventory.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Inventory information from availability zone.
	SlaveZoneStockInfos []*SlaveZoneStockInfo `json:"SlaveZoneStockInfos,omitnil,omitempty" name:"SlaveZoneStockInfos"`
}

// Predefined struct for user
type SetLibraDBClusterRenewFlagRequestParams struct {
	// Analysis Cluster ID List
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag 0: Normal renewal 1: Automatic renewal 2: Non-renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetLibraDBClusterRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Analysis Cluster ID List
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag 0: Normal renewal 1: Automatic renewal 2: Non-renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *SetLibraDBClusterRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLibraDBClusterRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLibraDBClusterRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLibraDBClusterRenewFlagResponseParams struct {
	// Quantity.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLibraDBClusterRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetLibraDBClusterRenewFlagResponseParams `json:"Response"`
}

func (r *SetLibraDBClusterRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLibraDBClusterRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRenewFlagRequestParams struct {
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be manipulated
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Auto-renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *SetRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRenewFlagResponseParams struct {
	// Number of successfully manipulated instances
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetRenewFlagResponseParams `json:"Response"`
}

func (r *SetRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlaveZoneAttrItem struct {
	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// binlog synchronization mode.
	BinlogSyncWay *string `json:"BinlogSyncWay,omitnil,omitempty" name:"BinlogSyncWay"`

	// Semi-Sync timeout period in milliseconds.
	SemiSyncTimeout *int64 `json:"SemiSyncTimeout,omitnil,omitempty" name:"SemiSyncTimeout"`
}

type SlaveZoneStockInfo struct {
	// Replica AZ.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Inventory quantity in spare availability zone.	
	StockCount *uint64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// Whether there is inventory in the spare availability zone.	
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`
}

type SlowQueriesItem struct {
	// Execution timestamp
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Execution duration in seconds
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// SQL statement
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Client host
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Lock duration in seconds
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Number of returned rows
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// MD5 value of the SQL statement
	SqlMd5 *string `json:"SqlMd5,omitnil,omitempty" name:"SqlMd5"`

	// Remote reading count.
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncReadCountRemote *int64 `json:"SyncReadCountRemote,omitnil,omitempty" name:"SyncReadCountRemote"`

	// Number of remote read bytes.
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncReadBytesRemote *int64 `json:"SyncReadBytesRemote,omitnil,omitempty" name:"SyncReadBytesRemote"`

	// Time spent on remote reads (µs).
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncReadTimeRemote *int64 `json:"SyncReadTimeRemote,omitnil,omitempty" name:"SyncReadTimeRemote"`

	// Remote write count.
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncWriteCountRemote *int64 `json:"SyncWriteCountRemote,omitnil,omitempty" name:"SyncWriteCountRemote"`

	// Specifies the number of remote written bytes.
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncWriteBytesRemote *int64 `json:"SyncWriteBytesRemote,omitnil,omitempty" name:"SyncWriteBytesRemote"`

	// Time spent on remote writing (µs).
	// Specifies that the database kernel version is larger than 3.1.12.
	SyncWriteTimeRemote *int64 `json:"SyncWriteTimeRemote,omitnil,omitempty" name:"SyncWriteTimeRemote"`

	// Transaction submission delay (µs).
	// Specifies that the database kernel version is larger than 3.1.12.
	TrxCommitDelay *int64 `json:"TrxCommitDelay,omitnil,omitempty" name:"TrxCommitDelay"`
}

type SnapshotBackupConfig struct {
	// System automation time.
	BackupCustomAutoTime *bool `json:"BackupCustomAutoTime,omitnil,omitempty" name:"BackupCustomAutoTime"`

	// Indicates the full backup start time. value range: [0-24*3600]. for example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.
	BackupTimeBeg *uint64 `json:"BackupTimeBeg,omitnil,omitempty" name:"BackupTimeBeg"`

	// Indicates the full backup end time. value range: [0-24*3600]. for example, 0:00, 1:00, and 2:00 are 0, 3600, and 7200 respectively.
	BackupTimeEnd *uint64 `json:"BackupTimeEnd,omitnil,omitempty" name:"BackupTimeEnd"`

	// This parameter currently does not support modification and no need to specify. backup frequency is an array of length 7, corresponding to the backup method from sunday to saturday, full for full backup and increment for incremental backup.
	BackupWeekDays []*string `json:"BackupWeekDays,omitnil,omitempty" name:"BackupWeekDays"`

	// Interval.
	BackupIntervalTime *int64 `json:"BackupIntervalTime,omitnil,omitempty" name:"BackupIntervalTime"`

	// Indicates the backup retention period in seconds. data will be cleaned up longer than this time. 7 days means 3600247=604800. the maximum is 158112000.
	ReserveDuration *uint64 `json:"ReserveDuration,omitnil,omitempty" name:"ReserveDuration"`

	// Automatic data backup trigger policy, periodically: automatic periodic backup, frequent: high-frequency backup
	BackupTriggerStrategy *string `json:"BackupTriggerStrategy,omitnil,omitempty" name:"BackupTriggerStrategy"`

	// Safe info
	AutoCopyVaults []*CreateBackupVaultItem `json:"AutoCopyVaults,omitnil,omitempty" name:"AutoCopyVaults"`
}

type SparseBackupConfig struct {
	// Sparse backup switch: ON/OFF
	SparseBackupSwitch *string `json:"SparseBackupSwitch,omitnil,omitempty" name:"SparseBackupSwitch"`

	// Sparse backup policy list (1-3)
	SparseBackupConfigInfos []*SparseBackupConfigInfo `json:"SparseBackupConfigInfos,omitnil,omitempty" name:"SparseBackupConfigInfos"`
}

type SparseBackupConfigInfo struct {
	// <p>Operation type: add, modify, remove</p>
	OpType *string `json:"OpType,omitnil,omitempty" name:"OpType"`

	// <p>Configuration ID</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Policy type: weekly/monthly/yearly</p>
	SparsePeriodConfig *string `json:"SparsePeriodConfig,omitnil,omitempty" name:"SparsePeriodConfig"`

	// <p>Cycle time configuration</p>
	SparsePeriodTime *SparsePeriodTime `json:"SparsePeriodTime,omitnil,omitempty" name:"SparsePeriodTime"`

	// <p>Retention days (7-7320 days, longest 20 years)</p>
	SparseBackupSaveDays *int64 `json:"SparseBackupSaveDays,omitnil,omitempty" name:"SparseBackupSaveDays"`
}

type SparseBackupConfigRsp struct {
	// <p>Sparse backup switch: ON/OFF</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparseBackupSwitch *string `json:"SparseBackupSwitch,omitnil,omitempty" name:"SparseBackupSwitch"`

	// <p>Sparse backup policy list (1-3)</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SparseBackupConfigInfos []*SparseBackupConfigInfo `json:"SparseBackupConfigInfos,omitnil,omitempty" name:"SparseBackupConfigInfos"`
}

type SparsePeriodTime struct {
	// Weekly: List of days of the week, 1-7, 1=Monday, 7=Sunday (Only used for weekly cycle, up to 7).
	WeekDays []*int64 `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// By month: Date list, 1-31 (for monthly cycle only, up to 7)
	Days []*int64 `json:"Days,omitnil,omitempty" name:"Days"`

	// Yearly: Month-day composite column (Only for yearly cycle, up to 7)
	MonthDays []*MonthDay `json:"MonthDays,omitnil,omitempty" name:"MonthDays"`
}

// Predefined struct for user
type StartCLSDeliveryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Enabled log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type StartCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Enabled log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *StartCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCLSDeliveryResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *StartCLSDeliveryResponseParams `json:"Response"`
}

func (r *StartCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCLSDeliveryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type StopCLSDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log topic ID.
	CLSTopicIds []*string `json:"CLSTopicIds,omitnil,omitempty" name:"CLSTopicIds"`

	// Log type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the maintenance time is in operation.
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *StopCLSDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCLSDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CLSTopicIds")
	delete(f, "LogType")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCLSDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCLSDeliveryResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCLSDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *StopCLSDeliveryResponseParams `json:"Response"`
}

func (r *StopCLSDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCLSDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchClusterLogBin struct {
	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type SwitchClusterVpcRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

type SwitchClusterVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`
}

func (r *SwitchClusterVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterVpcResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchClusterVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterVpcResponseParams `json:"Response"`
}

func (r *SwitchClusterVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneRequestParams struct {
	// <p>Cluster Id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Current availability zone</p>
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// <p>Availability zone to switch to</p>
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// <p>Execute during maintenance period - yes, immediately execute - no</p>
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type SwitchClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// <p>Cluster Id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Current availability zone</p>
	OldZone *string `json:"OldZone,omitnil,omitempty" name:"OldZone"`

	// <p>Availability zone to switch to</p>
	NewZone *string `json:"NewZone,omitnil,omitempty" name:"NewZone"`

	// <p>Execute during maintenance period - yes, immediately execute - no</p>
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *SwitchClusterZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldZone")
	delete(f, "NewZone")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterZoneResponseParams struct {
	// <p>Task ID.</p>
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>Async FlowId</p>
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchClusterZoneResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterZoneResponseParams `json:"Response"`
}

func (r *SwitchClusterZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchProxyVpcRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// VPC ID in string
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID in string
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Valid hours of old IP
	OldIpReserveHours *int64 `json:"OldIpReserveHours,omitnil,omitempty" name:"OldIpReserveHours"`

	// Database proxy group ID (required), which can be obtained through the `DescribeProxies` API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *SwitchProxyVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "OldIpReserveHours")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyVpcResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchProxyVpcResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyVpcResponseParams `json:"Response"`
}

func (r *SwitchProxyVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableMappingObject struct {
	// Source instance Id
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// Database name.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Map database name
	MapDatabaseName *string `json:"MapDatabaseName,omitnil,omitempty" name:"MapDatabaseName"`

	// Mapping table name
	MapTableName *string `json:"MapTableName,omitnil,omitempty" name:"MapTableName"`

	// Synchronization status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Sync progress
	Process *float64 `json:"Process,omitnil,omitempty" name:"Process"`

	// Delay
	Lag *int64 `json:"Lag,omitnil,omitempty" name:"Lag"`

	// Message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Whether it is the main table
	IsPrimary *bool `json:"IsPrimary,omitnil,omitempty" name:"IsPrimary"`

	// Virtual column padding value
	VirtualColValue *string `json:"VirtualColValue,omitnil,omitempty" name:"VirtualColValue"`
}

type TablePrivileges struct {
	// Database name
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskMaintainInfo struct {
	// Execution start time (seconds from 0 o'clock).
	MaintainStartTime *int64 `json:"MaintainStartTime,omitnil,omitempty" name:"MaintainStartTime"`

	// Specifies the continuous time. the unit is second.
	MaintainDuration *int64 `json:"MaintainDuration,omitnil,omitempty" name:"MaintainDuration"`

	// Specifies the time when it can be executed. valid values: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"].
	MaintainWeekDays []*string `json:"MaintainWeekDays,omitnil,omitempty" name:"MaintainWeekDays"`
}

type TaskProgressInfo struct {
	// Current step.
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// Current progress.
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil,omitempty" name:"CurrentStepProgress"`

	// Estimated Time
	CurrentStepRemainingTime *string `json:"CurrentStepRemainingTime,omitnil,omitempty" name:"CurrentStepRemainingTime"`
}

type TemplateParamInfo struct {
	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Default value
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// The collection of optional value types when the parameter type is `enum`.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// The maximum value when the parameter type is float/integer.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value when the parameter type is float/integer.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Whether to restart the instance for the parameter to take effect
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Parameter type. Valid value: `integer`, `float`, `string`, `enum`.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`
}

type TradePrice struct {
	// Resource total price under prepaid mode, excluding discounts. unit: microCent. 1 US dollar equals 1e8 microCents.
	TotalPrice *int64 `json:"TotalPrice,omitnil,omitempty" name:"TotalPrice"`

	// Total discount. `100` means no discount.
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Discounted total price under prepaid mode, unit: cent. 1 US dollar equals 1e8 microCents. for example, the user enjoys a Discount = TotalPrice * Discount.
	TotalPriceDiscount *int64 `json:"TotalPriceDiscount,omitnil,omitempty" name:"TotalPriceDiscount"`

	// Unit resource price in postpaid mode, excluding discounts. unit: cent. 1 US dollar equals 1e2 cents
	UnitPrice *int64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// Unit resource price in postpaid mode after Discount, unit: cent. 1 US dollar equals 1e2 cents. for example, the user enjoys a Discount = unitprice * Discount.
	UnitPriceDiscount *int64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// Price unit
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// Excludes discounted rates under high precision.
	UnitPriceHighPrecision *string `json:"UnitPriceHighPrecision,omitnil,omitempty" name:"UnitPriceHighPrecision"`

	// Discounted price under high precision.
	UnitPriceDiscountHighPrecision *string `json:"UnitPriceDiscountHighPrecision,omitnil,omitempty" name:"UnitPriceDiscountHighPrecision"`

	// Currency unit.
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

// Predefined struct for user
type TransferClusterZoneRequestParams struct {
	// Source Cluster Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Target availability zone
	DstZone *string `json:"DstZone,omitnil,omitempty" name:"DstZone"`

	// Cross-AZ migration master-slave data delay verification threshold, milliseconds (ms)
	MaxLag *int64 `json:"MaxLag,omitnil,omitempty" name:"MaxLag"`

	// Immediate: Immediate execution, InMaintain: Time window execution
	TransferType *string `json:"TransferType,omitnil,omitempty" name:"TransferType"`

	// Multi-availability zone backup zone
	DstSlaveZone *string `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// Target zone info for proxy migration
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

type TransferClusterZoneRequest struct {
	*tchttp.BaseRequest
	
	// Source Cluster Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Target availability zone
	DstZone *string `json:"DstZone,omitnil,omitempty" name:"DstZone"`

	// Cross-AZ migration master-slave data delay verification threshold, milliseconds (ms)
	MaxLag *int64 `json:"MaxLag,omitnil,omitempty" name:"MaxLag"`

	// Immediate: Immediate execution, InMaintain: Time window execution
	TransferType *string `json:"TransferType,omitnil,omitempty" name:"TransferType"`

	// Multi-availability zone backup zone
	DstSlaveZone *string `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// Target zone info for proxy migration
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`
}

func (r *TransferClusterZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferClusterZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DstZone")
	delete(f, "MaxLag")
	delete(f, "TransferType")
	delete(f, "DstSlaveZone")
	delete(f, "ProxyZones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferClusterZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferClusterZoneResponseParams struct {
	// Asynchronous task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferClusterZoneResponse struct {
	*tchttp.BaseResponse
	Response *TransferClusterZoneResponseParams `json:"Response"`
}

func (r *TransferClusterZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferClusterZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindClusterResourcePackagesRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The unique ID of a resource pack. If this parameter is left empty, all resource packs of the instance will be unbound.
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
}

type UnbindClusterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The unique ID of a resource pack. If this parameter is left empty, all resource packs of the instance will be unbound.
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`
}

func (r *UnbindClusterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindClusterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindClusterResourcePackagesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindClusterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindClusterResourcePackagesResponseParams `json:"Response"`
}

func (r *UnbindClusterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindClusterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAnalysisInstanceVersionInfo struct {
	// <p>ip</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Port</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Version</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Expiry time</p>
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

// Predefined struct for user
type UpgradeClusterVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Upgrade time type. Valid values: `upgradeImmediate`, `upgradeInMaintain`.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

type UpgradeClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Kernel version
	CynosVersion *string `json:"CynosVersion,omitnil,omitempty" name:"CynosVersion"`

	// Upgrade time type. Valid values: `upgradeImmediate`, `upgradeInMaintain`.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`
}

func (r *UpgradeClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "CynosVersion")
	delete(f, "UpgradeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterVersionResponseParams `json:"Response"`
}

func (r *UpgradeClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database CPU.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database memory in GB.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// Instance Machine Type
	// 1. common: general.
	// 2. exclusive, dedicated.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order).
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Proxy synchronous upgrade.
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database CPU.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Database memory in GB.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Upgrade type. Valid values: upgradeImmediate, upgradeInMaintain.
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// Instance Machine Type
	// 1. common: general.
	// 2. exclusive, dedicated.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// This parameter has been disused.
	StorageLimit *uint64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// Whether to automatically select a voucher. 1: yes; 0: no. Default value: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// This parameter has been disused.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Transaction mode. Valid values: `0` (place and pay for an order), `1` (place an order).
	DealMode *int64 `json:"DealMode,omitnil,omitempty" name:"DealMode"`

	// Valid values: `NormalUpgrade` (Normal mode), `FastUpgrade` (QuickChange). If the system detects that the configuration modification process will cause a momentary disconnection, the process will be terminated.
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// Proxy synchronous upgrade.
	UpgradeProxy *UpgradeProxy `json:"UpgradeProxy,omitnil,omitempty" name:"UpgradeProxy"`
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
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "UpgradeType")
	delete(f, "DeviceType")
	delete(f, "StorageLimit")
	delete(f, "AutoVoucher")
	delete(f, "DbType")
	delete(f, "DealMode")
	delete(f, "UpgradeMode")
	delete(f, "UpgradeProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// Frozen transaction ID.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Large order number.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// Order ID.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

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

type UpgradeProxy struct {
	// cpu
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Proxy node information
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// Rebalance
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`
}

// Predefined struct for user
type UpgradeProxyRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// ID of the database proxy group (disused)
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Load rebalance mode. Valid values: `auto`, `manual`
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `timeWindow` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// Node information of the atabase proxy
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// whether rolling upgrade
	IsRollUpgrade *string `json:"IsRollUpgrade,omitnil,omitempty" name:"IsRollUpgrade"`

	// Rolling upgrade waiting time, unit: second
	RollUpgradeWaitingTime *int64 `json:"RollUpgradeWaitingTime,omitnil,omitempty" name:"RollUpgradeWaitingTime"`
}

type UpgradeProxyRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Number of nodes in the proxy group
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// ID of the database proxy group (disused)
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Load rebalance mode. Valid values: `auto`, `manual`
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `timeWindow` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`

	// Node information of the atabase proxy
	ProxyZones []*ProxyZone `json:"ProxyZones,omitnil,omitempty" name:"ProxyZones"`

	// whether rolling upgrade
	IsRollUpgrade *string `json:"IsRollUpgrade,omitnil,omitempty" name:"IsRollUpgrade"`

	// Rolling upgrade waiting time, unit: second
	RollUpgradeWaitingTime *int64 `json:"RollUpgradeWaitingTime,omitnil,omitempty" name:"RollUpgradeWaitingTime"`
}

func (r *UpgradeProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "ProxyCount")
	delete(f, "ProxyGroupId")
	delete(f, "ReloadBalance")
	delete(f, "IsInMaintainPeriod")
	delete(f, "ProxyZones")
	delete(f, "IsRollUpgrade")
	delete(f, "RollUpgradeWaitingTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeProxyResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyResponseParams `json:"Response"`
}

func (r *UpgradeProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `yes` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Database proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Upgrade time. Valid values: `no` (upon upgrade completion), `yes` (upgrade during instance maintenance time)
	IsInMaintainPeriod *string `json:"IsInMaintainPeriod,omitnil,omitempty" name:"IsInMaintainPeriod"`
}

func (r *UpgradeProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SrcProxyVersion")
	delete(f, "DstProxyVersion")
	delete(f, "ProxyGroupId")
	delete(f, "IsInMaintainPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserHostPrivilege struct {
	// Authorized user
	DbUserName *string `json:"DbUserName,omitnil,omitempty" name:"DbUserName"`

	// Client IP
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// User permission
	DbPrivilege *string `json:"DbPrivilege,omitnil,omitempty" name:"DbPrivilege"`
}

type VaultInfo struct {
	// Safe id
	VaultId *string `json:"VaultId,omitnil,omitempty" name:"VaultId"`

	// Safe name
	VaultName *string `json:"VaultName,omitnil,omitempty" name:"VaultName"`

	// Safe region
	VaultRegion *string `json:"VaultRegion,omitnil,omitempty" name:"VaultRegion"`

	// Safe status
	VaultStatus *string `json:"VaultStatus,omitnil,omitempty" name:"VaultStatus"`

	// Backup retention period
	BackupSaveSeconds *int64 `json:"BackupSaveSeconds,omitnil,omitempty" name:"BackupSaveSeconds"`
}

type WillDeleteItem struct {
	// Backup file ID
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup file name
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`
}

type ZoneStockInfo struct {
	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Whether there is an inventory.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`

	// Quantity in stock
	StockCount *int64 `json:"StockCount,omitnil,omitempty" name:"StockCount"`

	// Available zone inventory information.
	SlaveZoneStockInfos []*SlaveZoneStockInfo `json:"SlaveZoneStockInfos,omitnil,omitempty" name:"SlaveZoneStockInfos"`
}

type ZoneStockInfo4Libra struct {
	// Availability zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// whether there is inventory
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasStock *bool `json:"HasStock,omitnil,omitempty" name:"HasStock"`
}