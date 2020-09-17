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

package v20190823

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ClearTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be cleared
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *ClearTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of cleared tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of table clearing results
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Cluster data description language type, such as `PROTO`, `TDR`, or `MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// Network type
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// ID of the VPC instance with which a cluster is associated
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the subnet instance with which a cluster is associated
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Cluster password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Password status
	PasswordStatus *string `json:"PasswordStatus,omitempty" name:"PasswordStatus"`

	// TcaplusDB SDK connection parameter: access ID
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// TcaplusDB SDK connection parameter: access address
	ApiAccessIp *string `json:"ApiAccessIp,omitempty" name:"ApiAccessIp"`

	// TcaplusDB SDK connection parameter: access port
	ApiAccessPort *int64 `json:"ApiAccessPort,omitempty" name:"ApiAccessPort"`

	// If `PasswordStatus` is `unmodifiable`, the old password has not expired, and this field will display its expiration time; otherwise, this field will be empty
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// TcaplusDB SDK connection parameter for accessing IPv6 addresses
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAccessIpv6 *string `json:"ApiAccessIpv6,omitempty" name:"ApiAccessIpv6"`
}

type CompareIdlFilesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// Selected list of uploaded IDL files. Either this parameter or `NewIdlFiles` must be selected
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles" list`

	// List of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be selected
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles" list`
}

func (r *CompareIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompareIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information list of all IDL files uploaded and verified in this request
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// Number of tables verified to be valid in this request
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Verification result parsed from the selected table after the IDL description file is read
		TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompareIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be backed up resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information list of tables to be backed up
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of IDs of created backup tasks
		TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster data description language type, such as `PROTO`, `TDR`, or `MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// Cluster name, which can contain up to 32 letters and digits
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of the VPC instance bound to a cluster in the format of `vpc-f49l6u0z`
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ID of the subnet instance bound to a cluster in the format of `subnet-pxir56ns`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Cluster access password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	Password *string `json:"Password,omitempty" name:"Password"`

	// Cluster tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags" list`

	// Whether to enable IPv6 address access for clusters
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
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

type CreateTableGroupRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group name, which can contain up to 32 letters and digits
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// Table group ID, which can be customized but must be unique in one cluster. If it is not specified, the auto-increment mode will be used.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table group tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags" list`
}

func (r *CreateTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of table group successfully created
		TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table creation IDL file list selected by user
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// Information list of tables to be created
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// Table tag list
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags" list`
}

func (r *CreateTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tables created in batches
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of tables created in batches
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// ID of cluster to be deleted
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

		// Task ID generated by cluster deletion
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

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

type DeleteIdlFilesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where IDL resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of IDL files to be deleted
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`
}

func (r *DeleteIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Deletion result
		IdlFileInfos []*IdlFileInfoWithoutContent `json:"IdlFileInfos,omitempty" name:"IdlFileInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableGroupRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

func (r *DeleteTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID generated by table group deletion
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be dropped resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of information of tables to be dropped
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *DeleteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of dropped tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of details of dropped tables
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterTagsRequest struct {
	*tchttp.BaseRequest

	// The list of cluster IDs
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`
}

func (r *DescribeClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The information list of cluster tags
		Rows []*TagsInfoOfCluster `json:"Rows,omitempty" name:"Rows" list`

		// The number of returned results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// List of IDs of clusters to be queried
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// Query filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to enable IPv6 address access
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
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

		// Number of cluster instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Cluster instance list
		Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters" list`

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

type DescribeIdlFileInfosRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a file resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where a file resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// File ID list
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds" list`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdlFileInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdlFileInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdlFileInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of files
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of file details
		IdlFileInfos []*IdlFileInfo `json:"IdlFileInfos,omitempty" name:"IdlFileInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdlFileInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdlFileInfosResponse) FromJsonString(s string) error {
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

		// Number of queried AZs
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of AZ query results
		RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos" list`

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

type DescribeTableGroupTagsRequest struct {
	*tchttp.BaseRequest

	// The ID of the cluster where table group tags need to be queried
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of IDs of the table groups whose tags need to be queried
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`
}

func (r *DescribeTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The information list of table group tags
		Rows []*TagsInfoOfTableGroup `json:"Rows,omitempty" name:"Rows" list`

		// The number of returned results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableGroupsRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID list
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// Filter. Valid values: TableGroupName, TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTableGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of table groups
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Table group information list
		TableGroups []*TableGroupInfo `json:"TableGroups,omitempty" name:"TableGroups" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableTagsRequest struct {
	*tchttp.BaseRequest

	// The ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table list
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *DescribeTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The total number of returned results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The information list of table tags
		Rows []*TagsInfoOfTable `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesInRecycleRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesInRecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesInRecycleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesInRecycleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Table details result list
		TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesInRecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesInRecycleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be queried resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of IDs of the table groups where the table to be queried resides
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// Information list of tables to be queried
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// Filter. Valid values: TableName, TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Query result offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned query results
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Table details result list
		TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// List of IDs of clusters where the tasks to be queried reside
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// List of IDs of tasks to be queried
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

	// Filter. Valid values: Content, TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Query list offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results in query list
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tasks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of details of queried tasks
		TaskInfos []*TaskInfoNew `json:"TaskInfos,omitempty" name:"TaskInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUinInWhitelistRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUinInWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUinInWhitelistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUinInWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Query result. FALSE: yes, TRUE: no
		Result *string `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUinInWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUinInWhitelistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ErrorInfo struct {

	// Error code
	Code *string `json:"Code,omitempty" name:"Code"`

	// Error message
	Message *string `json:"Message,omitempty" name:"Message"`
}

type Filter struct {

	// Filter field name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter field value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type IdlFileInfo struct {

	// Filename excluding extension
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Data interface description language (IDL) type
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// File extension
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// File ID, which is meaningful for files already uploaded
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// File content, which is meaningful for files to be uploaded in this request
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type IdlFileInfoWithoutContent struct {

	// Filename excluding extension
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Data interface description language (IDL) type
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// File extension
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// File size in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// File ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster to be renamed
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name to be changed to, which can contain up to 32 letters and digits
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterPasswordRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster for which to modify the password
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Old cluster password
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// Expected expiration time of old cluster password
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// New cluster password, which must contain lowercase letters (a-z), uppercase letters (A-Z), and digits (0-9).
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// Update mode. 1: updates password, 2: updates old password expiration time. Default value: 1
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyClusterPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterTagsRequest struct {
	*tchttp.BaseRequest

	// The ID of the cluster whose tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupNameRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a table group resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group to be renamed
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// New table group name, which can contain letters and symbols
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

func (r *ModifyTableGroupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableGroupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupTagsRequest struct {
	*tchttp.BaseRequest

	// The ID of the cluster where table group tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The ID of the table group whose tags need to be modified
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// Tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableMemosRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster instance where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of details of selected tables
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos" list`
}

func (r *ModifyTableMemosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableMemosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableMemosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of tables modified for remarks
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of table remarks modification results
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableMemosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableMemosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableQuotasRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be scaled resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of quotas of tables selected for modification
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas" list`
}

func (r *ModifyTableQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of scaled tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of table scaling results
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableTagsRequest struct {
	*tchttp.BaseRequest

	// The ID of the cluster where table tags need to be modified
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// The list of tables whose tags need to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// The list of tags to add or modify
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// The list of tags to delete
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The total number of returned results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Returned results
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be modified resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Selected table modification IDL files
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// List of tables to be modified
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *ModifyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of modified tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of table modification results
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ParsedTableInfoNew struct {

	// Table description language type. Valid values: PROTO, TDR
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// Table instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type. Valid values: GENERIC, LIST
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Primary key field information
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyFields *string `json:"KeyFields,omitempty" name:"KeyFields"`

	// Old primary key field information, which is valid during verification of table modification
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldKeyFields *string `json:"OldKeyFields,omitempty" name:"OldKeyFields"`

	// Non-primary key field information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValueFields *string `json:"ValueFields,omitempty" name:"ValueFields"`

	// Old non-primary key field information, which is valid during verification of table modification
	// Note: this field may return null, indicating that no valid values can be obtained.
	OldValueFields *string `json:"OldValueFields,omitempty" name:"OldValueFields"`

	// Table group ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Total size of primary key field
	// Note: this field may return null, indicating that no valid values can be obtained.
	SumKeyFieldSize *int64 `json:"SumKeyFieldSize,omitempty" name:"SumKeyFieldSize"`

	// Total size of non-primary key fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SumValueFieldSize *int64 `json:"SumValueFieldSize,omitempty" name:"SumValueFieldSize"`

	// Index key set
	// Note: this field may return null, indicating that no valid values can be obtained.
	IndexKeySet *string `json:"IndexKeySet,omitempty" name:"IndexKeySet"`

	// Shardkey set
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// TDR version number
	// Note: this field may return null, indicating that no valid values can be obtained.
	TdrVersion *int64 `json:"TdrVersion,omitempty" name:"TdrVersion"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Number of LIST-type table elements
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Number of SORTLIST-type table sort fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// Sort order of SORTLIST-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`
}

type RecoverRecycleTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where a table resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Information of tables to be recovered
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *RecoverRecycleTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycleTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecoverRecycleTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of recovered tables
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of information of recovered tables
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverRecycleTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycleTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region `Ap-code`
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region abbreviation
	RegionAbbr *string `json:"RegionAbbr,omitempty" name:"RegionAbbr"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Whether to support IPv6 address access. Valid values: 0 (support), 1 (not support)
	Ipv6Enable *uint64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

type RollbackTablesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where the table to be rolled back resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// List of tables to be rolled back
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// Time to roll back to
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// Rollback mode. `KEYS` is supported
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *RollbackTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of table rollback task results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Table rollback task result list
		TableResults []*TableRollbackResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SelectedTableInfoNew struct {

	// ID of the table group where a table resides
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table instance ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table description language type. Valid values: PROTO, TDR
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// Table data structure type. Valid values: GENERIC, LIST
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Number of LIST-type table elements
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Reserved table capacity in GB
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// Reserved table read QPS
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// Reserved table write QPS
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// Table remarks
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Key rollback filename, which is only used for rollback
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Key rollback file extension, which is only used for rollback
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// Key rollback file size, which is only used for rollback
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Key rollback file content, which is only used for rollback
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type TableGroupInfo struct {

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Table group name
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// Table group creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Number of tables in table group
	TableCount *uint64 `json:"TableCount,omitempty" name:"TableCount"`

	// Total table storage capacity in MB in table group
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type TableInfoNew struct {

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the cluster where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Name of the cluster where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Name of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// JSON string of table's primary key field structure
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyStruct *string `json:"KeyStruct,omitempty" name:"KeyStruct"`

	// JSON string of table non-primary key field structure
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValueStruct *string `json:"ValueStruct,omitempty" name:"ValueStruct"`

	// Table shardkey set, which is valid for PROTO-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// Table index key field set, which is valid for PROTO-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	IndexStruct *string `json:"IndexStruct,omitempty" name:"IndexStruct"`

	// Number of LIST-type table elements
	// Note: this field may return null, indicating that no valid values can be obtained.
	ListElementNum *uint64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// Information list of IDL files associated with table
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// Reserved table capacity in GB
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// Reserved table read QPS
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// Reserved table write QPS
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// Actual table data size in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableSize *int64 `json:"TableSize,omitempty" name:"TableSize"`

	// Table status
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Table creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Table's last modified time
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// Table remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// TcaplusDB SDK data access ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// Number of SORTLIST-type table sort fields
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// Sort order of SORTLIST-type tables
	// Note: this field may return null, indicating that no valid values can be obtained.
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`

	// Distributed index information of table
	DbClusterInfoStruct *string `json:"DbClusterInfoStruct,omitempty" name:"DbClusterInfoStruct"`
}

type TableResultNew struct {

	// Table instance ID in the format of `tcaplus-3be64cbb`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Task ID, which is valid for the API that creates one task
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Task ID list, which is valid for the API that creates multiple tasks
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`
}

type TableRollbackResultNew struct {

	// Table instance ID in the format of `tcaplus-3be64cbb`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Task ID, which is valid for the API that creates one task
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Table name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table data structure type, such as `GENERIC` or `LIST`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// Table data interface description language (IDL) type, such as `PROTO` or `TDR`
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// ID of the table group where a table resides
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Error message
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Task ID list, which is valid for the API that creates multiple tasks
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

	// ID of uploaded key file
	// Note: this field may return null, indicating that no valid values can be obtained.
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// Number of keys successfully verified
	// Note: this field may return null, indicating that no valid values can be obtained.
	SuccKeyNum *uint64 `json:"SuccKeyNum,omitempty" name:"SuccKeyNum"`

	// Total number of keys contained in key file
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalKeyNum *uint64 `json:"TotalKeyNum,omitempty" name:"TotalKeyNum"`
}

type TagInfoUnit struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfCluster struct {

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags" list`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTable struct {

	// Table instance ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// Table name
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags" list`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTableGroup struct {

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Table group ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags" list`

	// Error message
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TaskInfoNew struct {

	// Task ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// Task type
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// ID of TcaplusDB internal transaction associated with task
	TransId *string `json:"TransId,omitempty" name:"TransId"`

	// ID of the cluster where a task resides
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Name of the cluster where a task resides
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Task progress
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// Task creation time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Task last modified time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Task details
	Content *string `json:"Content,omitempty" name:"Content"`
}

type VerifyIdlFilesRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster where to create a table
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ID of the table group where to create a table
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// List of information of uploaded IDL files. Either this parameter or `NewIdlFiles` must be present
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles" list`

	// List of information of IDL files to be uploaded. Either this parameter or `ExistingIdlFiles` must be present
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles" list`
}

func (r *VerifyIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information list of all IDL files uploaded and verified in this request
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// Number of valid tables parsed by reading IDL description file, excluding tables already created
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of valid tables parsed by reading IDL description file, excluding tables already created
		TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
