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

package v20230202

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Cluster struct {
	// User APP ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// Account ID.Note: This field may return null, indicating that no valid values can be obtained.
	AccountID *string `json:"AccountID,omitnil,omitempty" name:"AccountID"`

	// Customizes the instance name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Region.Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone.Note: This field may return null, indicating that no valid values can be obtained.
	Zones *string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Network information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Networks is deprecated.
	Networks []*Network `json:"Networks,omitnil,omitempty" name:"Networks"`

	// Instance specification.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: Spec is deprecated.
	Spec *Spec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Instance status. 0: running; 1: creating; 16: adjusting configuration; 17: isolating; 18: to be terminated; 19: recovering; 20: shutting down; 21: terminating; 22: terminated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance validity period.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Period *Period `json:"Period,omitnil,omitempty" name:"Period"`

	// Creation time.Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Last modification time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Internal features of the product.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tenant *Tenant `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// Tag.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Security group information.Note: This field may return null, indicating that no valid values can be obtained.
	Security []*string `json:"Security,omitnil,omitempty" name:"Security"`
}

type Database struct {
	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// Database name.Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Cold storage time (days).
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoolDownInDays *int64 `json:"CoolDownInDays,omitnil,omitempty" name:"CoolDownInDays"`

	// Data retention time (days).
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionInDays *int64 `json:"RetentionInDays,omitnil,omitempty" name:"RetentionInDays"`

	// Remarks.Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Status. 0: initializing resources; 1: creating resources; 2: normal status; 3: deleting resources; 4: deleted resources; 5: disabling resources; 6: disabled resources; 7: abnormal resources, and manual operation is required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time.Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Last modification time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Current page number.		
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Page size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query parameter: Filtering and querying by instance ID (cluster_id) and instance name (name) are supported.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sorting parameter: Sorting by the creation time field (created_at) is supported. The value of Type can be set to DESC (descending order) or ASC (ascending order).
	Orders []*Order `json:"Orders,omitnil,omitempty" name:"Orders"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Current page number.		
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Page size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query parameter: Filtering and querying by instance ID (cluster_id) and instance name (name) are supported.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sorting parameter: Sorting by the creation time field (created_at) is supported. The value of Type can be set to DESC (descending order) or ASC (ascending order).
	Orders []*Order `json:"Orders,omitnil,omitempty" name:"Orders"`
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
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "Orders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Total number of records under current conditions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instances meeting the conditions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Clusters []*Cluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`

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
type DescribeDatabasesRequestParams struct {
	// Database parameter.
	Database *Database `json:"Database,omitnil,omitempty" name:"Database"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Pagination page.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Database parameter.
	Database *Database `json:"Database,omitnil,omitempty" name:"Database"`

	// Pagination size.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Pagination page.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
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
	delete(f, "Database")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// Database list.
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Quantity.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type Filter struct {
	// Filter parameter.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter expression.
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// Value involved in filtering.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Network struct {
	// vpc id
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc subnet id
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VPC IP address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// VPC port address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Order struct {
	// Sorting field.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Sorting method.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Period struct {
	// Start time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type Spec struct {
	// 1: yearly/monthly subscription; 2: bill by hour.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Request unit. 0 indicates following the resource configuration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestUnit *uint64 `json:"RequestUnit,omitnil,omitempty" name:"RequestUnit"`

	// Maximum number of CPU cores.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CpuLimit *float64 `json:"CpuLimit,omitnil,omitempty" name:"CpuLimit"`

	// Maximum memory size (Gi).
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryLimit *float64 `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// Maximum number of disks (Gi).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskLimit *uint64 `json:"DiskLimit,omitnil,omitempty" name:"DiskLimit"`

	// Number of business shards.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Shards *uint64 `json:"Shards,omitnil,omitempty" name:"Shards"`

	// Number of business nodes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Replicas *uint64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`
}

type Tag struct {
	// Key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Tenant struct {
	// Whether the password is encrypted.
	IsPasswordEncrypted *bool `json:"IsPasswordEncrypted,omitnil,omitempty" name:"IsPasswordEncrypted"`
}