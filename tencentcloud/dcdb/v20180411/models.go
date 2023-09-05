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

package v20180411

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// Account name
	User *string `json:"User,omitnil" name:"User"`

	// Host address
	Host *string `json:"Host,omitnil" name:"Host"`
}

// Predefined struct for user
type ActiveHourDCDBInstanceRequestParams struct {
	// List of instance IDs in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type ActiveHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *ActiveHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActiveHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActiveHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActiveHourDCDBInstanceResponseParams struct {
	// IDs of instances removed from isolation
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil" name:"SuccessInstanceIds"`

	// IDs of instances failed to be removed from isolation
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil" name:"FailedInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ActiveHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActiveHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *ActiveHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActiveHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddShardConfig struct {
	// The number of shards to be added
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// Shard memory capacity in GB
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage capacity in GB
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of tdsqlshard-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of tdsqlshard-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type BriefNodeInfo struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node role. Valid values: `master`, `slave`
	Role *string `json:"Role,omitnil" name:"Role"`

	// The ID of the shard where the node resides
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`
}

// Predefined struct for user
type CancelDcnJobRequestParams struct {
	// Disaster recovery instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest
	
	// Disaster recovery instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDcnJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDcnJobResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelDcnJobResponseParams `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source user account name
	SrcUser *string `json:"SrcUser,omitnil" name:"SrcUser"`

	// Source user host
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target user account name
	DstUser *string `json:"DstUser,omitnil" name:"DstUser"`

	// Target user host
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// Target account description
	DstDesc *string `json:"DstDesc,omitnil" name:"DstDesc"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source user account name
	SrcUser *string `json:"SrcUser,omitnil" name:"SrcUser"`

	// Source user host
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target user account name
	DstUser *string `json:"DstUser,omitnil" name:"DstUser"`

	// Target user host
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// Target account description
	DstDesc *string `json:"DstDesc,omitnil" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUser")
	delete(f, "SrcHost")
	delete(f, "DstUser")
	delete(f, "DstHost")
	delete(f, "DstDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *CloneAccountResponseParams `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessRequestParams struct {
	// ID of an instance for which to disable public network access. The ID is in the format of dcdbt-ow728lmc and can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// ID of an instance for which to disable public network access. The ID is in the format of dcdbt-ow728lmc and can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessResponseParams struct {
	// Async task ID. The task status can be queried through the `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`

	// Column name
	Column *string `json:"Column,omitnil" name:"Column"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

type ConfigValue struct {
	// Configuration name, which supports `max_user_connections`.
	Config *string `json:"Config,omitnil" name:"Config"`

	// Configuration value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ConstraintRange struct {
	// Minimum value when the constraint type is `section`
	Min *string `json:"Min,omitnil" name:"Min"`

	// Maximum value when the constraint type is `section`
	Max *string `json:"Max,omitnil" name:"Max"`
}

// Predefined struct for user
type CopyAccountPrivilegesRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitnil" name:"SrcUserName"`

	// Access host allowed for a source user
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitnil" name:"DstUserName"`

	// Access host allowed for a target user
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// `ReadOnly` attribute of a source account
	SrcReadOnly *string `json:"SrcReadOnly,omitnil" name:"SrcReadOnly"`

	// `ReadOnly` attribute of a target account
	DstReadOnly *string `json:"DstReadOnly,omitnil" name:"DstReadOnly"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitnil" name:"SrcUserName"`

	// Access host allowed for a source user
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitnil" name:"DstUserName"`

	// Access host allowed for a target user
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// `ReadOnly` attribute of a source account
	SrcReadOnly *string `json:"SrcReadOnly,omitnil" name:"SrcReadOnly"`

	// `ReadOnly` attribute of a target account
	DstReadOnly *string `json:"DstReadOnly,omitnil" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUserName")
	delete(f, "SrcHost")
	delete(f, "DstUserName")
	delete(f, "DstHost")
	delete(f, "SrcReadOnly")
	delete(f, "DstReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAccountPrivilegesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *CopyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Account password. It must contain 8-32 characters in all of the following four types: lowercase letters, uppercase letters, digits, and symbols (()~!@#$%^&*-+=_|{}[]:<>,.?/), and cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail; 3: only the secondary will be read from.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitnil" name:"Description"`

	// If the secondary delay exceeds the set value of this parameter, the secondary will be deemed to have failed.
	// It is recommended that this parameter be set to a value greater than 10. This parameter takes effect when `ReadOnly` is 1 or 2.
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. If left empty or `0` is passed in, the connections will be unlimited. This parameter configuration is not supported for kernel version 10.1.
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Account password. It must contain 8-32 characters in all of the following four types: lowercase letters, uppercase letters, digits, and symbols (()~!@#$%^&*-+=_|{}[]:<>,.?/), and cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail; 3: only the secondary will be read from.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitnil" name:"Description"`

	// If the secondary delay exceeds the set value of this parameter, the secondary will be deemed to have failed.
	// It is recommended that this parameter be set to a value greater than 10. This parameter takes effect when `ReadOnly` is 1 or 2.
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. If left empty or `0` is passed in, the connections will be unlimited. This parameter configuration is not supported for kernel version 10.1.
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "ReadOnly")
	delete(f, "Description")
	delete(f, "DelayThresh")
	delete(f, "SlaveConst")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Instance ID, which is passed through from the input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is passed through from the input parameters.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host allowed for access, which is passed through from the input parameters.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Passed through from the input parameters.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDCDBInstanceRequestParams struct {
	// AZs to deploy shard nodes. You can specify up to two AZs. When the shard specification is 1-source-2-replica, two of the nodes are deployed in the first AZ.
	// The current purchasable AZ needs be pulled through `DescribeDCDBSaleInfo` API.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Shard memory size in GB, which can be obtained 
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage size in GB, which can be obtained
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// Number of nodes in a single shard, which can be obtained
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// The number of shards in the instance. Value range: 2-8. You can increase up to 64 shards by upgrading your instance.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// The number of instances to be purchased
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// VPC ID. If this parameter is left empty or not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Whether to automatically use vouchers. This option is disabled by default.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Custom name of the instance
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; `0`: case-sensitive; `1`: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; `0`: async; `1`: strong sync; `2`: downgradable strong sync; default value: `2`).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// DCN source region
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// DCN source instance ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// Renewal mode. Valid values: `0` (manual renewal, which is the default mode), `1` (auto-renewal), `2` (manual renewal, which is specified by users).  If no renewal is required, set it to `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Security group IDs in array. This parameter is compatible with the old parameter `SecurityGroupId`.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type CreateDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// AZs to deploy shard nodes. You can specify up to two AZs. When the shard specification is 1-source-2-replica, two of the nodes are deployed in the first AZ.
	// The current purchasable AZ needs be pulled through `DescribeDCDBSaleInfo` API.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Shard memory size in GB, which can be obtained 
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage size in GB, which can be obtained
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// Number of nodes in a single shard, which can be obtained
	//  by querying the instance specification through `DescribeShardSpec` API.
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// The number of shards in the instance. Value range: 2-8. You can increase up to 64 shards by upgrading your instance.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// The number of instances to be purchased
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// VPC ID. If this parameter is left empty or not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Whether to automatically use vouchers. This option is disabled by default.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Custom name of the instance
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; `0`: case-sensitive; `1`: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; `0`: async; `1`: strong sync; `2`: downgradable strong sync; default value: `2`).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// DCN source region
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// DCN source instance ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// Renewal mode. Valid values: `0` (manual renewal, which is the default mode), `1` (auto-renewal), `2` (manual renewal, which is specified by users).  If no renewal is required, set it to `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Security group IDs in array. This parameter is compatible with the old parameter `SecurityGroupId`.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *CreateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "Period")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardNodeCount")
	delete(f, "ShardCount")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersionId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceName")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "InitParams")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "AutoRenewFlag")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDCDBInstanceResponseParams struct {
	// Long order ID, which is used to call the `DescribeOrders` API.
	//  The parameter can be used to either query order details or call the user account APIs to make another payment when this payment fails.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// IDs of the instances you have purchased in this order. If no instance IDs are returned, you can query them with the `DescribeOrders` API. You can also use the `DescribeDBInstances` API to check whether an instance has been created successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDCDBInstanceRequestParams struct {
	// Shard memory in GB, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard capacity in GB, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// The number of nodes per shard, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// The number of shards in the instance. Value range: 2-8. Upgrade your instance to have up to 64 shards if you require more.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// The number of instances to be purchased
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// VPC ID. If this parameter is left empty or not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The number of CPU cores per shard, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardCpu *int64 `json:"ShardCpu,omitnil" name:"ShardCpu"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// AZs to deploy shard nodes. You can specify up to two AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Custom name of the instance
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the region of the associated source instance so that the disaster recovery instance can sync data with the source instance over the Data Communication Network (DCN).
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the ID of the associated source instance so that the disaster recovery instance can sync data with the source instance over the Data Communication Network (DCN).
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: 2).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// ID of the instance to be rolled back
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil" name:"RollbackInstanceId"`

	// Rollback time, such as "2021-11-22 00:00:00".
	RollbackTime *string `json:"RollbackTime,omitnil" name:"RollbackTime"`

	// Array of security group IDs (this parameter is compatible with the old parameter `SecurityGroupId`)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type CreateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Shard memory in GB, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard capacity in GB, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// The number of nodes per shard, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// The number of shards in the instance. Value range: 2-8. Upgrade your instance to have up to 64 shards if you require more.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// The number of instances to be purchased
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// VPC ID. If this parameter is left empty or not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The number of CPU cores per shard, which can be obtained through the `DescribeShardSpec` API.
	//   
	ShardCpu *int64 `json:"ShardCpu,omitnil" name:"ShardCpu"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// AZs to deploy shard nodes. You can specify up to two AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Custom name of the instance
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the region of the associated source instance so that the disaster recovery instance can sync data with the source instance over the Data Communication Network (DCN).
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the ID of the associated source instance so that the disaster recovery instance can sync data with the source instance over the Data Communication Network (DCN).
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: 2).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// ID of the instance to be rolled back
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil" name:"RollbackInstanceId"`

	// Rollback time, such as "2021-11-22 00:00:00".
	RollbackTime *string `json:"RollbackTime,omitnil" name:"RollbackTime"`

	// Array of security group IDs (this parameter is compatible with the old parameter `SecurityGroupId`)
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *CreateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardNodeCount")
	delete(f, "ShardCount")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ShardCpu")
	delete(f, "DbVersionId")
	delete(f, "Zones")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceName")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "InitParams")
	delete(f, "RollbackInstanceId")
	delete(f, "RollbackTime")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDCDBInstanceResponseParams struct {
	// IDs of the instances you have purchased in this order. If no instance IDs are returned, you can query them with the `DescribeOrders` API. You can also use the `DescribeDBInstances` API to check whether an instance has been created successfully.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Task ID, which can be used to query the creation progress
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Order ID, which is used for calling the `DescribeOrders` API.
	//  The parameter can be used to either query order details or call the user account APIs to make another payment when this payment fails.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {
	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host from which a user can log in (corresponding to the `host` field for a MySQL user; a user is uniquely identified by username and host; this parameter is in IP format and ends with % for IP range; % can be entered; if this parameter is left empty, % will be used by default).
	Host *string `json:"Host,omitnil" name:"Host"`

	// User remarks
	Description *string `json:"Description,omitnil" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Read-only flag. 0: no; 1: for the account's SQL requests, the replica will be used first, and if it is unavailable, the source will be used; 2: the replica will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// If the replica delay exceeds the set value of this parameter, the replica will be considered to have failed.
	// Set this parameter to a value above 10. This parameter takes effect when `ReadOnly` is 1 or 2.
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. `0` indicates no limit.	
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

type DBParamValue struct {
	// Parameter name
	Param *string `json:"Param,omitnil" name:"Param"`

	// Parameter value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DCDBInstanceInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Application ID
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitnil" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Numeric ID of a VPC
	VpcId *int64 `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet Digital ID
	SubnetId *int64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Instance status. Valid values: `0` (creating), `1` (running task), `2` (running), `3` (uninitialized), `-1` (isolated), `4` (initializing), `5` (eliminating), `6` (restarting), `7` (migrating data)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Private IP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Auto-renewal flag
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Number of shards
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Isolation time
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil" name:"IsolatedTimestamp"`

	// Account ID
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// Shard details
	ShardDetail []*ShardInfo `json:"ShardDetail,omitnil" name:"ShardDetail"`

	// Number of nodes. 2: one master and one slave; 3: one master and two slaves
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Temporary instance flag. 0: non-temporary instance
	IsTmp *int64 `json:"IsTmp,omitnil" name:"IsTmp"`

	// Dedicated cluster ID. If this parameter is empty, the instance is a non-dedicated cluster instance
	ExclusterId *string `json:"ExclusterId,omitnil" name:"ExclusterId"`

	// VPC ID in string type
	UniqueVpcId *string `json:"UniqueVpcId,omitnil" name:"UniqueVpcId"`

	// VPC subnet ID in string type
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil" name:"UniqueSubnetId"`

	// Numeric ID of instance (this field is obsolete and should not be depended on)
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Domain name for public network access, which can be resolved by the public network
	WanDomain *string `json:"WanDomain,omitnil" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network
	WanVip *string `json:"WanVip,omitnil" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitnil" name:"WanPort"`

	// Product type ID (this field is obsolete and should not be depended on)
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Last updated time of an instance in the format of 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Database engine
	DbEngine *string `json:"DbEngine,omitnil" name:"DbEngine"`

	// Database engine version
	DbVersion *string `json:"DbVersion,omitnil" name:"DbVersion"`

	// Billing mode
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Async task flow ID when an async task is in progress on an instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Locker *int64 `json:"Locker,omitnil" name:"Locker"`

	// Public network access status. 0: not enabled; 1: enabled; 2: disabled; 3: enabling
	WanStatus *int64 `json:"WanStatus,omitnil" name:"WanStatus"`

	// Whether the instance supports audit. 1: yes; 0: no
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitnil" name:"IsAuditSupported"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// Indicates whether the instance uses IPv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Private network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// Public network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitnil" name:"WanVipv6"`

	// Public network IPv6 port
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitnil" name:"WanPortIpv6"`

	// Public network IPv6 status
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitnil" name:"WanStatusIpv6"`

	// DCN type. Valid values: 0 (null), 1 (primary instance), 2 (disaster recovery instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: 0 (null), 1 (creating), 2 (syncing), 3 (disconnected)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// The number of DCN disaster recovery instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitnil" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (standard primary instance), `3` (standard disaster recovery instance), `4` (dedicated disaster recovery instance)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Instance tag information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// Database engine version
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`
}

type DCDBShardInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard SQL passthrough ID, which is used to pass through SQL statements to the specified shard for execution.
	ShardSerialId *string `json:"ShardSerialId,omitnil" name:"ShardSerialId"`

	// Globally unique shard ID
	ShardInstanceId *string `json:"ShardInstanceId,omitnil" name:"ShardInstanceId"`

	// Status. 0: creating; 1: processing; 2: running; 3: shard not initialized.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// VPC ID in string format
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID in string format
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitnil" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Number of nodes. 2: one source and one replica; 3: one source and two replicas
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Storage utilization in %
	StorageUsage *float64 `json:"StorageUsage,omitnil" name:"StorageUsage"`

	// Memory utilization in %
	MemoryUsage *float64 `json:"MemoryUsage,omitnil" name:"MemoryUsage"`

	// Numeric shard ID (this field is obsolete and should not be depended on)
	ShardId *int64 `json:"ShardId,omitnil" name:"ShardId"`

	// Product ID
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Proxy version
	ProxyVersion *string `json:"ProxyVersion,omitnil" name:"ProxyVersion"`

	// Billing mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Source AZ of the shard
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShardMasterZone *string `json:"ShardMasterZone,omitnil" name:"ShardMasterZone"`

	// List of replica AZs of the shard
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitnil" name:"ShardSlaveZones"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// The value range of shardkey, which includes 64 hash values, such as 0-31 or 32-63.
	Range *string `json:"Range,omitnil" name:"Range"`
}

type Database struct {
	// Database name
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

type DatabaseFunction struct {
	// Function name
	Func *string `json:"Func,omitnil" name:"Func"`
}

type DatabasePrivilege struct {
	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`
}

type DatabaseProcedure struct {
	// Stored procedure name
	Proc *string `json:"Proc,omitnil" name:"Proc"`
}

type DatabaseTable struct {
	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`
}

type DatabaseView struct {
	// View name
	View *string `json:"View,omitnil" name:"View"`
}

type DcnDetailItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Region where the instance resides
	Region *string `json:"Region,omitnil" name:"Region"`

	// Availability zone where the instance resides
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Instance IP address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Instance IPv6 address
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// Instance port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Instance status
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// DCN flag. Valid values: `1` (primary), `2` (disaster recovery)
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: `0` (none), `1` (creating), `2` (syncing), `3` (disconnected)
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// Number of CPU cores of the instance
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Instance memory capacity in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// Creation time of the instance in the format of 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Expiration time of the instance in the format of 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (non-dedicated primary instance), `3` (non-dedicated disaster recovery instance), and `4` (dedicated disaster recovery instance).
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Whether KMS is enabled.
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`
}

type Deal struct {
	// Order ID.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// Account
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Number of items
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// The associated process ID, which can be used to query the process execution status.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The ID of the created instance, which is required only for the order that creates an instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Billing mode. Valid values: `0` (postpaid), `1` (prepaid).
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user
	Host *string `json:"Host,omitnil" name:"Host"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user
	Host *string `json:"Host,omitnil" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty.
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried.
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty.
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried.
	ColName *string `json:"ColName,omitnil" name:"ColName"`
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
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Database account username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Database account host
	Host *string `json:"Host,omitnil" name:"Host"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Instance ID, which is passed through from the input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance user list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Users []*DBAccount `json:"Users,omitnil" name:"Users"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeBackupFilesRequestParams struct {
	// Query by instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Query by shard ID
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Query by start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query by end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Sorting dimension. Valid values: `Time`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Query by shard ID
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Query by start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query by end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Sorting dimension. Valid values: `Time`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "BackupType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesResponseParams struct {
	// List of backup files
	Files []*InstanceBackupFileItem `json:"Files,omitnil" name:"Files"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupFilesResponseParams `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesRequestParams struct {
	// Instance ID in the format of  `tdsqlshard-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of  `tdsqlshard-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesResponseParams struct {
	// Whether encryption is enabled. Valid values: `1` (enabled), `2` (disabled).
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`

	// DEK
	CipherText *string `json:"CipherText,omitnil" name:"CipherText"`

	// DEK expiration date
	ExpireDate *string `json:"ExpireDate,omitnil" name:"ExpireDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *DescribeDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID in the format of shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Requested log type. Valid values: 1 (binlog); 2 (cold backup); 3 (errlog); 4 (slowlog).
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID in the format of shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Requested log type. Valid values: 1 (binlog); 2 (cold backup); 3 (errlog); 4 (slowlog).
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesResponseParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog); 2 (cold backup); 3 (errlog); 4 (slowlog).
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// Total number of requested logs
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of log files
	Files []*LogFileInfo `json:"Files,omitnil" name:"Files"`

	// For an instance in a VPC, this prefix plus URI can be used as the download address
	VpcPrefix *string `json:"VpcPrefix,omitnil" name:"VpcPrefix"`

	// For an instance in a common network, this prefix plus URI can be used as the download address
	NormalPrefix *string `json:"NormalPrefix,omitnil" name:"NormalPrefix"`

	// Shard ID in the format of shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBLogFilesResponseParams `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersResponseParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requests the current parameter values of the database
	Params []*ParamDesc `json:"Params,omitnil" name:"Params"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParametersResponseParams `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Instance VIP
	// Note: This field may return null, indicating that no valid values can be obtained.
	VIP *string `json:"VIP,omitnil" name:"VIP"`

	// Instance Port
	// Note: This field may return null, indicating that no valid value can be obtained.
	VPort *string `json:"VPort,omitnil" name:"VPort"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeDBSlowLogsRequestParams struct {
	// Instance ID in the format of dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data entry number starting from which to return results
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Query start time in the format of 2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Shard ID of the instance in the format of shard-53ima8ln
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Query end time in the format of 2016-08-22 14:55:20. If this parameter is left empty, the current time will be used as the query end time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Specific name of the database to be queried
	Db *string `json:"Db,omitnil" name:"Db"`

	// Sorting metric. Valid values: `query_time_sum`, `query_count`. Default value: `query_time_sum`
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `desc` (descending), `asc` (ascending). Default value: `desc`
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Query slow queries from either the source or the replica. Valid values: `0` (source), `1` (replica). Default value: `0`
	Slave *int64 `json:"Slave,omitnil" name:"Slave"`
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data entry number starting from which to return results
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Query start time in the format of 2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Shard ID of the instance in the format of shard-53ima8ln
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Query end time in the format of 2016-08-22 14:55:20. If this parameter is left empty, the current time will be used as the query end time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Specific name of the database to be queried
	Db *string `json:"Db,omitnil" name:"Db"`

	// Sorting metric. Valid values: `query_time_sum`, `query_count`. Default value: `query_time_sum`
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `desc` (descending), `asc` (ascending). Default value: `desc`
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Query slow queries from either the source or the replica. Valid values: `0` (source), `1` (replica). Default value: `0`
	Slave *int64 `json:"Slave,omitnil" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "ShardId")
	delete(f, "EndTime")
	delete(f, "Db")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Slave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowLogsResponseParams struct {
	// Sum of all statement lock durations
	LockTimeSum *float64 `json:"LockTimeSum,omitnil" name:"LockTimeSum"`

	// Total number of statement queries
	QueryCount *int64 `json:"QueryCount,omitnil" name:"QueryCount"`

	// Total number of results
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Sum of all statement query durations
	QueryTimeSum *float64 `json:"QueryTimeSum,omitnil" name:"QueryTimeSum"`

	// Slow query log data
	Data []*SlowLogData `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSyncModeRequestParams struct {
	// ID of an instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest
	
	// ID of an instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSyncModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSyncModeResponseParams struct {
	// Sync mode. 0: async; 1: strong sync; 2: downgradable strong sync
	SyncMode *int64 `json:"SyncMode,omitnil" name:"SyncMode"`

	// Whether a modification is in progress. 1: yes; 0: no.
	IsModifying *int64 `json:"IsModifying,omitnil" name:"IsModifying"`

	// Current sync mode. Valid values: `0` (async), `1` (sync).
	CurrentSyncMode *int64 `json:"CurrentSyncMode,omitnil" name:"CurrentSyncMode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSyncModeResponseParams `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceDetailRequestParams struct {
	// Instance ID, such as dcdbt-7oaxtcb7.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDCDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as dcdbt-7oaxtcb7.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDCDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceDetailResponseParams struct {
	// Instance ID, such as dcdbt-7oaxtcb7.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance status. Valid values: `0` (creating), `1` (running task), `2` (running), `3` (uninitialized), `-1` (isolated).
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Current status of the instance
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Instance private IP address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private port of instance
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Number of instance nodes. Valid values: `2` (1-source-1-replica), `3` (1-source-2-replica).
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Instance region, such as ap-guangzhou.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Instance VPC ID, such as vpc-r9jr0de3.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID of an instance, such as subnet-6rqs61o2.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Public network status. Valid values: `0` (not enabled), `1` (enabled), `2` (disabled), `3`: (enabling), `4` (disabling).
	WanStatus *int64 `json:"WanStatus,omitnil" name:"WanStatus"`

	// Domain name for public network access, which can be resolved by the public network.
	WanDomain *string `json:"WanDomain,omitnil" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network.
	WanVip *string `json:"WanVip,omitnil" name:"WanVip"`

	// Public network access port
	WanPort *int64 `json:"WanPort,omitnil" name:"WanPort"`

	// Project ID of the instance
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Automatic renewal flag for an instance. Valid values: `0` (normal renewal), `1` (auto-renewal), `3` (no renewal upon expiration).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Dedicated cluster ID
	ExclusterId *string `json:"ExclusterId,omitnil" name:"ExclusterId"`

	// Billing mode. Valid values: `prepaid` (monthly subscription), `postpaid` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Creation time of the instance in the format of 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Expiration time of the instance in the format of 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Database version information
	DbVersion *string `json:"DbVersion,omitnil" name:"DbVersion"`

	// Whether the instance supports audit. Valid values: `1` (yes), `0` (no).
	IsAuditSupported *int64 `json:"IsAuditSupported,omitnil" name:"IsAuditSupported"`

	// Whether data encryption is supported for an instance. Valid values: `1` (yes), `0` (no).
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitnil" name:"IsEncryptSupported"`

	// Instance machine model
	Machine *string `json:"Machine,omitnil" name:"Machine"`

	// Instance memory size in GB, which is the sum of the memory of all shards.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance disk storage size in GB, which is the sum of the disk size of all shards.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Instance storage space utilization. It is calculated by dividing the sum of the used disk size of all shards by the total disk size of all shards.
	StorageUsage *float64 `json:"StorageUsage,omitnil" name:"StorageUsage"`

	// Size of log storage space in GB
	LogStorage *int64 `json:"LogStorage,omitnil" name:"LogStorage"`

	// Product type ID
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Source AZ
	MasterZone *string `json:"MasterZone,omitnil" name:"MasterZone"`

	// Replica AZ
	SlaveZones []*string `json:"SlaveZones,omitnil" name:"SlaveZones"`

	// Shard information
	Shards []*ShardBriefInfo `json:"Shards,omitnil" name:"Shards"`

	// Private network IPv6 address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// Number of CPU cores of an instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Instance QPS
	// Note: This field may return null, indicating that no valid values can be obtained.
	Qps *int64 `json:"Qps,omitnil" name:"Qps"`

	// Database engine
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbEngine *string `json:"DbEngine,omitnil" name:"DbEngine"`

	// Whether IPv6 is supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Public IPv6 address, which can be accessed over the public network
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitnil" name:"WanVipv6"`

	// Public network status. Valid values: `0` (not enabled), `1` (enabled), `2` (disabled), `3`: (enabling), `4` (disabling).
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *int64 `json:"WanStatusIpv6,omitnil" name:"WanStatusIpv6"`

	// Public network IPv6 port
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *int64 `json:"WanPortIpv6,omitnil" name:"WanPortIpv6"`

	// Tag information
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// DCN type. Valid values: `0` (N/A), `1` (source instance), `2` (disaster recovery read-only instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: `0` (N/A), `1` (creating), `2` (syncing), `3` (disconnected).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// The number of DCN disaster recovery instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitnil" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (non-dedicated primary instance), `3` (non-dedicated disaster recovery read-only instance), `4` (dedicated disaster recovery read-only instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Whether the instance supports setting the connection limit, which is not supported for kernel version 10.1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsMaxUserConnectionsSupported *bool `json:"IsMaxUserConnectionsSupported,omitnil" name:"IsMaxUserConnectionsSupported"`

	// The displayed database version
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Encryption status. Valid values: `0` (disabled), `1` (enabled).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`

	// Type of dedicated cluster. Valid values: `0` (public cloud), `1` (finance cage), `2` (CDC cluster).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Nearby VPC access
	// Note: This field may return null, indicating that no valid values can be obtained.
	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitnil" name:"RsAccessStrategy"`

	// Unclaimed network resource
	ReservedNetResources []*ReservedNetResource `json:"ReservedNetResources,omitnil" name:"ReservedNetResources"`


	IsPhysicalReplicationSupported *bool `json:"IsPhysicalReplicationSupported,omitnil" name:"IsPhysicalReplicationSupported"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDCDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceNodeInfoRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The maximum number of results returned at a time. Value range: (0-100]. Default value: `100`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset of the returned results. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeDCDBInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The maximum number of results returned at a time. Value range: (0-100]. Default value: `100`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset of the returned results. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeDCDBInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceNodeInfoResponseParams struct {
	// Total number of nodes
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Node information
	NodesInfo []*BriefNodeInfo `json:"NodesInfo,omitnil" name:"NodesInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDCDBInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstanceNodeInfoResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstancesRequestParams struct {
	// Query by instance ID or IDs. Instance ID is in the format of dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Search field name. Valid values: instancename (search by instance name); vip (search by private IP); all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitnil" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Query by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Sort by field. Valid values: projectId; createtime; instancename
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sort by type. Valid values: desc; asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 1: non-dedicated cluster; 2: dedicated cluster; 0: all
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Identifies whether to use the `ExclusterType` field. false: no; true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil" name:"IsFilterExcluster"`

	// Dedicated cluster ID
	ExclusterIds []*string `json:"ExclusterIds,omitnil" name:"ExclusterIds"`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitnil" name:"FilterInstanceType"`

	// Use this filter to include instances in specific statuses
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Use this filter to exclude instances in specific statuses
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil" name:"ExcludeStatus"`
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID or IDs. Instance ID is in the format of dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Search field name. Valid values: instancename (search by instance name); vip (search by private IP); all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitnil" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Query by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Sort by field. Valid values: projectId; createtime; instancename
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sort by type. Valid values: desc; asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 1: non-dedicated cluster; 2: dedicated cluster; 0: all
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Identifies whether to use the `ExclusterType` field. false: no; true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil" name:"IsFilterExcluster"`

	// Dedicated cluster ID
	ExclusterIds []*string `json:"ExclusterIds,omitnil" name:"ExclusterIds"`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitnil" name:"FilterInstanceType"`

	// Use this filter to include instances in specific statuses
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Use this filter to exclude instances in specific statuses
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil" name:"ExcludeStatus"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SearchName")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "IsFilterVpc")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ExclusterType")
	delete(f, "IsFilterExcluster")
	delete(f, "ExclusterIds")
	delete(f, "TagKeys")
	delete(f, "FilterInstanceType")
	delete(f, "Status")
	delete(f, "ExcludeStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstancesResponseParams struct {
	// Number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of instance details
	Instances []*DCDBInstanceInfo `json:"Instances,omitnil" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBPriceRequestParams struct {
	// AZ ID of the purchased instance.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// The number of instances to be purchased. You can purchase 1-10 instances.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Number of nodes in a single shard, which can be obtained
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// Shard memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage size in GB, which can be obtained
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// The number of shards in the instance. Value range: 2-8. Upgrade your instance to have up to 64 shards if you require more.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Price unit. Valid values:   
	// `* pent` (cent), 
	// `* microPent` (microcent).
	AmountUnit *string `json:"AmountUnit,omitnil" name:"AmountUnit"`
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID of the purchased instance.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// The number of instances to be purchased. You can purchase 1-10 instances.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Number of nodes in a single shard, which can be obtained
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`

	// Shard memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage size in GB, which can be obtained
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// The number of shards in the instance. Value range: 2-8. Upgrade your instance to have up to 64 shards if you require more.
	ShardCount *int64 `json:"ShardCount,omitnil" name:"ShardCount"`

	// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Price unit. Valid values:   
	// `* pent` (cent), 
	// `* microPent` (microcent).
	AmountUnit *string `json:"AmountUnit,omitnil" name:"AmountUnit"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Count")
	delete(f, "Period")
	delete(f, "ShardNodeCount")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardCount")
	delete(f, "Paymode")
	delete(f, "AmountUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBPriceResponseParams struct {
	// Original price  
	// * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description.
	// * Currency: CNY (Chinese site), USD (international site)
	OriginalPrice *int64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// The actual price may be different from the original price due to discounts. 
	// * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description.
	// * Currency: CNY (Chinese site), USD (international site)
	Price *int64 `json:"Price,omitnil" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBPriceResponseParams `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBShardsRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID list.
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil" name:"ShardInstanceIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sort by field. Only `createtime` is supported currently.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID list.
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil" name:"ShardInstanceIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sort by field. Only `createtime` is supported currently.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardInstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBShardsResponseParams struct {
	// Number of eligible shards
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Shard information list
	Shards []*DCDBShardInfo `json:"Shards,omitnil" name:"Shards"`

	// Disaster recovery flag. Valid values: 0 (none), 1 (source instance), 2 (disaster recovery instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBShardsResponseParams `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// Passed through from input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table list.
	Tables []*DatabaseTable `json:"Tables,omitnil" name:"Tables"`

	// View list.
	Views []*DatabaseView `json:"Views,omitnil" name:"Views"`

	// Stored procedure list.
	Procs []*DatabaseProcedure `json:"Procs,omitnil" name:"Procs"`

	// Function list.
	Funcs []*DatabaseFunction `json:"Funcs,omitnil" name:"Funcs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseObjectsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name, which can be obtained through the `DescribeDatabaseObjects` API.
	Table *string `json:"Table,omitnil" name:"Table"`
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name, which can be obtained through the `DescribeDatabaseObjects` API.
	Table *string `json:"Table,omitnil" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableResponseParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name.
	Table *string `json:"Table,omitnil" name:"Table"`

	// Column information.
	Cols []*TableColumn `json:"Cols,omitnil" name:"Cols"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseTableResponseParams `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// The database list of this instance.
	Databases []*Database `json:"Databases,omitnil" name:"Databases"`

	// Passed through from input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDcnDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDcnDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDcnDetailResponseParams struct {
	// DCN synchronization details
	DcnDetails []*DcnDetailItem `json:"DcnDetails,omitnil" name:"DcnDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDcnDetailResponseParams `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Unsigned file path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

type DescribeFileDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard ID
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Unsigned file path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

func (r *DescribeFileDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlResponseParams struct {
	// Signed download URL
	PreSignedUrl *string `json:"PreSignedUrl,omitnil" name:"PreSignedUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeFileDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowRequestParams struct {
	// Task ID returned by an async request API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// Task ID returned by an async request API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`
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
	// Task status. Valid values: `0` (succeeded), `1` (failed), `2` (running)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeOrdersRequestParams struct {
	// List of long order IDs to be queried, which are returned by the APIs for creating, renewing, or scaling instances.
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// List of long order IDs to be queried, which are returned by the APIs for creating, renewing, or scaling instances.
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// Returned number of orders
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Order information list
	Deals []*Deal `json:"Deals,omitnil" name:"Deals"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
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
	delete(f, "Product")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Number of security groups.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DestroyDCDBInstanceRequestParams struct {
	// Instance ID in the format of tdsqlshard-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsqlshard-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDCDBInstanceResponseParams struct {
	// Instance ID, which is the same as the request parameter `InstanceId`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/557/56485?from_cn_redirect=1) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDCDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDCDBInstanceRequestParams struct {
	// Instance ID in the format of tdsqlshard-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsqlshard-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDCDBInstanceResponseParams struct {
	// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/557/56485?from_cn_redirect=1) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Instance ID, which is the same as the request parameter `InstanceId`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type ExpandShardConfig struct {
	// Shard IDs in array
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil" name:"ShardInstanceIds"`

	// Shard memory capacity in GB
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage capacity in GB
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`

	// Number of shard nodes
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil" name:"ShardNodeCount"`
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Global permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`, `SHOW DATABASES`, `REPLICATION CLIENT`, `REPLICATION SLAVE`.
	// Database permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`. 
	// Table permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.  
	// Field permission. Valid values: `INSERT`, `REFERENCES`, `SELECT`, `UPDATE`.
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Type. Valid values: `table`, `\*`. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is table, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` = table and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Global permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`, `SHOW DATABASES`, `REPLICATION CLIENT`, `REPLICATION SLAVE`.
	// Database permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`. 
	// Table permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.  
	// Field permission. Valid values: `INSERT`, `REFERENCES`, `SELECT`, `UPDATE`.
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Type. Valid values: `table`, `\*`. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is table, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` = table and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitnil" name:"ColName"`
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
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Privileges")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type InitDCDBInstancesRequestParams struct {
	// List of IDs of instances to be initialized. The ID is in the format of `dcdbt-ow728lmc` and can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Parameter list. Valid values: character_set_server (character set; required); lower_case_table_names (table name case sensitivity; required; 0: case-sensitive, 1: case-insensitive); innodb_page_size (InnoDB data page; default size: 16 KB); sync_mode (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: strong sync).
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of instances to be initialized. The ID is in the format of `dcdbt-ow728lmc` and can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Parameter list. Valid values: character_set_server (character set; required); lower_case_table_names (table name case sensitivity; required; 0: case-sensitive, 1: case-insensitive); innodb_page_size (InnoDB data page; default size: 16 KB); sync_mode (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: strong sync).
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDCDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDCDBInstancesResponseParams struct {
	// Async task ID. The task status can be queried through the `DescribeFlow` API.
	FlowIds []*uint64 `json:"FlowIds,omitnil" name:"FlowIds"`

	// Passed through from input parameters.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InitDCDBInstancesResponseParams `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceBackupFileItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance status
	InstanceStatus *int64 `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// Shard ID
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// File path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// File name
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// File size
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Manual backup. Valid values: `0` (no), `1` (yes).
	ManualBackup *int64 `json:"ManualBackup,omitnil" name:"ManualBackup"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

// Predefined struct for user
type IsolateDCDBInstanceRequestParams struct {
	// Instance ID in the format of `tdsqlshard-avw0207d`,  which is the same as the instance ID displayed on the TencentDB console and can be queried through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type IsolateDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsqlshard-avw0207d`,  which is the same as the instance ID displayed on the TencentDB console and can be queried through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *IsolateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDCDBInstanceResponseParams struct {
	// IDs of isolated instances
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil" name:"SuccessInstanceIds"`

	// IDs of instances failed to be isolated
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil" name:"FailedInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDCDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceRequestParams struct {
	// Instance ID in the format of `dcdbt-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type IsolateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `dcdbt-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *IsolateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateHourDCDBInstanceRequestParams struct {
	// ID list of the instances to be upgraded  in the format of  `dcdbt-ow728lmc`, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type IsolateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID list of the instances to be upgraded  in the format of  `dcdbt-ow728lmc`, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *IsolateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateHourDCDBInstanceResponseParams struct {
	// IDs of isolated instances
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil" name:"SuccessInstanceIds"`

	// IDs of instances failed to be isolated
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil" name:"FailedInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of session IDs
	SessionId []*int64 `json:"SessionId,omitnil" name:"SessionId"`

	// Shard ID. Either `ShardId` or `ShardSerialId` is required.
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Shard sequence ID. Either `ShardId` or `ShardSerialId` is required.
	ShardSerialId *string `json:"ShardSerialId,omitnil" name:"ShardSerialId"`
}

type KillSessionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of session IDs
	SessionId []*int64 `json:"SessionId,omitnil" name:"SessionId"`

	// Shard ID. Either `ShardId` or `ShardSerialId` is required.
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// Shard sequence ID. Either `ShardId` or `ShardSerialId` is required.
	ShardSerialId *string `json:"ShardSerialId,omitnil" name:"ShardSerialId"`
}

func (r *KillSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	delete(f, "ShardId")
	delete(f, "ShardSerialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type KillSessionResponse struct {
	*tchttp.BaseResponse
	Response *KillSessionResponseParams `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {
	// Last modified time of a log
	Mtime *uint64 `json:"Mtime,omitnil" name:"Mtime"`

	// File length
	Length *uint64 `json:"Length,omitnil" name:"Length"`

	// Uniform resource identifier (URI) used during log download
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// Filename
	FileName *string `json:"FileName,omitnil" name:"FileName"`
}

// Predefined struct for user
type ModifyAccountConfigRequestParams struct {
	// Instance ID in the format of  `tdsqlshard-kpkvq5oj`, which is the same as the one displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Account name
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Account domain name
	Host *string `json:"Host,omitnil" name:"Host"`

	// Configuration list. Each element in the list is a pair of `Config` and `Value`.
	Configs []*ConfigValue `json:"Configs,omitnil" name:"Configs"`
}

type ModifyAccountConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of  `tdsqlshard-kpkvq5oj`, which is the same as the one displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Account name
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Account domain name
	Host *string `json:"Host,omitnil" name:"Host"`

	// Configuration list. Each element in the list is a pair of `Config` and `Value`.
	Configs []*ConfigValue `json:"Configs,omitnil" name:"Configs"`
}

func (r *ModifyAccountConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Configs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccountConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountConfigResponseParams `json:"Response"`
}

func (r *ModifyAccountConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
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
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyAccountPrivilegesRequestParams struct {
	// Instance ID in the format of tdsql-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database account, including username and host address.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: If the parameter is left empty, no change will be made to the granted global permissions. To clear the granted global permissions, set the parameter to an empty array.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// Database permission. Value range: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".	
	// Note: If the parameter is not passed in, no change will be made to the granted stored procedure permissions. To clear the granted database permissions, set `Privileges` to an empty array.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil" name:"DatabasePrivileges"`

	// Database table permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: If the parameter is not passed in, no change will be made to the granted view permissions. To clear the granted table permissions, set `Privileges` to an empty array.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil" name:"TablePrivileges"`

	// Column permission in the table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: If the parameter is not passed in, no change will be made to the granted column permissions. To clear the granted column permissions, set `Privileges` to an empty array.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil" name:"ColumnPrivileges"`

	// Database view permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: If the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted view permissions, set `Privileges` to an empty array.
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil" name:"ViewPrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsql-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database account, including username and host address.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: If the parameter is left empty, no change will be made to the granted global permissions. To clear the granted global permissions, set the parameter to an empty array.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// Database permission. Value range: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".	
	// Note: If the parameter is not passed in, no change will be made to the granted stored procedure permissions. To clear the granted database permissions, set `Privileges` to an empty array.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil" name:"DatabasePrivileges"`

	// Database table permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: If the parameter is not passed in, no change will be made to the granted view permissions. To clear the granted table permissions, set `Privileges` to an empty array.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil" name:"TablePrivileges"`

	// Column permission in the table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: If the parameter is not passed in, no change will be made to the granted column permissions. To clear the granted column permissions, set `Privileges` to an empty array.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil" name:"ColumnPrivileges"`

	// Database view permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: If the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted view permissions, set `Privileges` to an empty array.
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil" name:"ViewPrivileges"`
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
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ViewPrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// Async task ID, which can be used in the [DescribeFlow](https://www.tencentcloud.com/document/product/237/16177) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyDBEncryptAttributesRequestParams struct {
	// Instance ID in the format of `tdsqlshard-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to enable the data encryption (Once enabled, it can’t be disabled). Valid values: `1` (Yes), `0` (No. Default).
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil" name:"EncryptEnabled"`
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsqlshard-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to enable the data encryption (Once enabled, it can’t be disabled). Valid values: `1` (Yes), `0` (No. Default).
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil" name:"EncryptEnabled"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EncryptEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// Instance ID in the format of tdsql-hdaprz0v
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsql-hdaprz0v
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `dcdb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
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
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type ModifyDBInstancesProjectRequestParams struct {
	// List of IDs of instances to be modified. Instance ID is in the format of dcdbt-ow728lmc.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of instances to be modified. Instance ID is in the format of dcdbt-ow728lmc.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter list. Each element is a combination of `Param` and `Value`.
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter list. Each element is a combination of `Param` and `Value`.
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter modification result
	Result []*ParamModifyResult `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeRequestParams struct {
	// ID of the instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
	SyncMode *int64 `json:"SyncMode,omitnil" name:"SyncMode"`
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
	SyncMode *int64 `json:"SyncMode,omitnil" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSyncModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeResponseParams struct {
	// Async task ID. The task status can be queried through the `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSyncModeResponseParams `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The ID of the desired VPC network
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The subnet ID of the desired VPC network
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The field is required to specify VIP.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// The field is required to specify VIPv6.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The ID of the desired VPC network
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The subnet ID of the desired VPC network
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The field is required to specify VIP.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// The field is required to specify VIPv6.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	delete(f, "Vipv6")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkResponseParams struct {
	// Async task ID, which can be used to query the task status through `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNetworkResponseParams `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPv6 flag
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPv6 flag
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vip")
	delete(f, "Ipv6Flag")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVipResponseParams `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance Vport
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`
}

type ModifyInstanceVportRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance Vport
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`
}

func (r *ModifyInstanceVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVportResponseParams `json:"Response"`
}

func (r *ModifyInstanceVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node role. Valid values: `master`, `slave`.
	Role *string `json:"Role,omitnil" name:"Role"`
}

type ParamConstraint struct {
	// Constraint type, such as `enum` and `section`
	Type *string `json:"Type,omitnil" name:"Type"`

	// List of valid values when constraint type is `enum`
	Enum *string `json:"Enum,omitnil" name:"Enum"`

	// Range when constraint type is `section`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Range *ConstraintRange `json:"Range,omitnil" name:"Range"`

	// List of valid values when constraint type is `string`
	String *string `json:"String,omitnil" name:"String"`
}

type ParamDesc struct {
	// Parameter name
	Param *string `json:"Param,omitnil" name:"Param"`

	// Current parameter value
	Value *string `json:"Value,omitnil" name:"Value"`

	// Previously set value, which is the same as `value` after the parameter takes effect. If no value has been set, this field will not be returned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetValue *string `json:"SetValue,omitnil" name:"SetValue"`

	// Default value
	Default *string `json:"Default,omitnil" name:"Default"`

	// Parameter constraint
	Constraint *ParamConstraint `json:"Constraint,omitnil" name:"Constraint"`

	// Whether a value has been set. false: no, true: yes.
	HaveSetValue *bool `json:"HaveSetValue,omitnil" name:"HaveSetValue"`

	// Whether restart is required. false: no;
	// true: yes.
	NeedRestart *bool `json:"NeedRestart,omitnil" name:"NeedRestart"`
}

type ParamModifyResult struct {
	// Renames parameter
	Param *string `json:"Param,omitnil" name:"Param"`

	// Result of parameter modification. 0: success; -1: failure; -2: invalid parameter value
	Code *int64 `json:"Code,omitnil" name:"Code"`
}

type ReservedNetResource struct {
	// VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Reserved private network IP under `VpcId` and `SubnetId`
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Port under `Vip`
	Vports []*int64 `json:"Vports,omitnil" name:"Vports"`

	// Valid hours of VIP	
	RecycleTime *string `json:"RecycleTime,omitnil" name:"RecycleTime"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitnil" name:"Password"`
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
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil" name:"SecurityGroupRemark"`

	// Inbound rule
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil" name:"Inbound"`

	// Outbound rule
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil" name:"Outbound"`
}

type SecurityGroupBound struct {
	// Policy, which can be `ACCEPT` or `DROP`
	Action *string `json:"Action,omitnil" name:"Action"`

	// Source IP or source IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`
}

type ShardBriefInfo struct {
	// Shard serial ID
	ShardSerialId *string `json:"ShardSerialId,omitnil" name:"ShardSerialId"`

	// Shard ID, such as shard-7vg1o339.
	ShardInstanceId *string `json:"ShardInstanceId,omitnil" name:"ShardInstanceId"`

	// Shard running status
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Description of shard running status
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Shard creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Shard memory size in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Shard disk size in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Log disk space size of a shard in GB
	LogDisk *int64 `json:"LogDisk,omitnil" name:"LogDisk"`

	// Number of shard nodes
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Disk space utilization of a shard
	StorageUsage *float64 `json:"StorageUsage,omitnil" name:"StorageUsage"`

	// Version information of the shard proxy
	ProxyVersion *string `json:"ProxyVersion,omitnil" name:"ProxyVersion"`

	// Source AZ of a shard
	ShardMasterZone *string `json:"ShardMasterZone,omitnil" name:"ShardMasterZone"`

	// Replica AZ of a shard
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitnil" name:"ShardSlaveZones"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodesInfo []*NodeInfo `json:"NodesInfo,omitnil" name:"NodesInfo"`
}

type ShardInfo struct {
	// Shard ID
	ShardInstanceId *string `json:"ShardInstanceId,omitnil" name:"ShardInstanceId"`

	// Shard set ID
	ShardSerialId *string `json:"ShardSerialId,omitnil" name:"ShardSerialId"`

	// Status. 0: creating; 1: processing; 2: running; 3: shard not initialized; -2: shard deleted
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Creation time
	Createtime *string `json:"Createtime,omitnil" name:"Createtime"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Numeric ID of a shard
	ShardId *int64 `json:"ShardId,omitnil" name:"ShardId"`

	// Number of nodes. 2: one primary and one secondary; 3: one primary and two secondaries
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Product type ID (this field is obsolete and should not be depended on)
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`
}

type SlowLogData struct {
	// Statement checksum for querying details
	CheckSum *string `json:"CheckSum,omitnil" name:"CheckSum"`

	// Database name
	Db *string `json:"Db,omitnil" name:"Db"`

	// Abstracted SQL statement
	FingerPrint *string `json:"FingerPrint,omitnil" name:"FingerPrint"`

	// Average lock duration
	LockTimeAvg *string `json:"LockTimeAvg,omitnil" name:"LockTimeAvg"`

	// Maximum lock duration
	LockTimeMax *string `json:"LockTimeMax,omitnil" name:"LockTimeMax"`

	// Minimum lock duration
	LockTimeMin *string `json:"LockTimeMin,omitnil" name:"LockTimeMin"`

	// Sum of lock durations
	LockTimeSum *string `json:"LockTimeSum,omitnil" name:"LockTimeSum"`

	// Number of queries
	QueryCount *string `json:"QueryCount,omitnil" name:"QueryCount"`

	// Average query duration
	QueryTimeAvg *string `json:"QueryTimeAvg,omitnil" name:"QueryTimeAvg"`

	// Maximum query duration
	QueryTimeMax *string `json:"QueryTimeMax,omitnil" name:"QueryTimeMax"`

	// Minimum query duration
	QueryTimeMin *string `json:"QueryTimeMin,omitnil" name:"QueryTimeMin"`

	// Sum of query durations
	QueryTimeSum *string `json:"QueryTimeSum,omitnil" name:"QueryTimeSum"`

	// Number of scanned rows
	RowsExaminedSum *string `json:"RowsExaminedSum,omitnil" name:"RowsExaminedSum"`

	// Number of sent rows
	RowsSentSum *string `json:"RowsSentSum,omitnil" name:"RowsSentSum"`

	// Last execution time
	TsMax *string `json:"TsMax,omitnil" name:"TsMax"`

	// First execution time
	TsMin *string `json:"TsMin,omitnil" name:"TsMin"`

	// Account
	User *string `json:"User,omitnil" name:"User"`

	// Sample SQL
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExampleSql *string `json:"ExampleSql,omitnil" name:"ExampleSql"`

	// Host address of account
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil" name:"Host"`
}

type SplitShardConfig struct {
	// Shard IDs in array
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil" name:"ShardInstanceIds"`

	// Data split ratio at 50% (fixed)
	SplitRate *int64 `json:"SplitRate,omitnil" name:"SplitRate"`

	// Shard memory capacity in GB
	ShardMemory *int64 `json:"ShardMemory,omitnil" name:"ShardMemory"`

	// Shard storage capacity in GB
	ShardStorage *int64 `json:"ShardStorage,omitnil" name:"ShardStorage"`
}

// Predefined struct for user
type SwitchDBInstanceHARequestParams struct {
	// Instance ID in the format of tdsql-ow728lmc
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to source node.
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

type SwitchDBInstanceHARequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsql-ow728lmc
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to source node.
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceHARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceHARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceHAResponseParams struct {
	// Async task ID
	FlowId *uint64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchDBInstanceHAResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstanceHAResponseParams `json:"Response"`
}

func (r *SwitchDBInstanceHAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceHAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableColumn struct {
	// Column name
	Col *string `json:"Col,omitnil" name:"Col"`

	// Column type
	Type *string `json:"Type,omitnil" name:"Type"`
}

type TablePrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

// Predefined struct for user
type TerminateDedicatedDBInstanceRequestParams struct {
	// Instance ID in the format of `dcdbt-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type TerminateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `dcdbt-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *TerminateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDedicatedDBInstanceResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *TerminateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeHourDCDBInstanceRequestParams struct {
	// Instance ID to be upgraded in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Upgrade type. Valid values: 
	// <li> `ADD`: Add a new shard </li> 
	//  <li> `EXPAND`: Upgrade the existing shads</li> 
	//  <li> `SPLIT`: Split data of the existing shads to the new ones</li>
	UpgradeType *string `json:"UpgradeType,omitnil" name:"UpgradeType"`

	// Add shards when `UpgradeType` is `ADD`.
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil" name:"AddShardConfig"`

	// Expand shard when `UpgradeType` is `EXPAND`.
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil" name:"ExpandShardConfig"`

	// Split shard when `UpgradeType` is `SPLIT`.
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil" name:"SplitShardConfig"`

	// Switch start time in the format of "2019-12-12 07:00:00", which is no less than one hour and within 3 days from the current time.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil" name:"SwitchStartTime"`

	// Switch end time in the format of "2019-12-12 07:15:00", which must be later than the start time.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil" name:"SwitchEndTime"`

	// Whether to retry automatically. Valid values: `0` (no), `1` (yes).
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil" name:"SwitchAutoRetry"`

	// The list of new AZs specified in deployment modification. The first one is the source AZ, and the rest are replica AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`
}

type UpgradeHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID to be upgraded in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Upgrade type. Valid values: 
	// <li> `ADD`: Add a new shard </li> 
	//  <li> `EXPAND`: Upgrade the existing shads</li> 
	//  <li> `SPLIT`: Split data of the existing shads to the new ones</li>
	UpgradeType *string `json:"UpgradeType,omitnil" name:"UpgradeType"`

	// Add shards when `UpgradeType` is `ADD`.
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil" name:"AddShardConfig"`

	// Expand shard when `UpgradeType` is `EXPAND`.
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil" name:"ExpandShardConfig"`

	// Split shard when `UpgradeType` is `SPLIT`.
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil" name:"SplitShardConfig"`

	// Switch start time in the format of "2019-12-12 07:00:00", which is no less than one hour and within 3 days from the current time.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil" name:"SwitchStartTime"`

	// Switch end time in the format of "2019-12-12 07:15:00", which must be later than the start time.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil" name:"SwitchEndTime"`

	// Whether to retry automatically. Valid values: `0` (no), `1` (yes).
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil" name:"SwitchAutoRetry"`

	// The list of new AZs specified in deployment modification. The first one is the source AZ, and the rest are replica AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`
}

func (r *UpgradeHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "SwitchAutoRetry")
	delete(f, "Zones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeHourDCDBInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewPrivileges struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// View name
	View *string `json:"View,omitnil" name:"View"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}