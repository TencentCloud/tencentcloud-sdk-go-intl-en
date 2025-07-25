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

package v20240718

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type CreateIncrementalMigrationStrategyRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// Incremental migration strategy name. The name length should not exceed 100 characters. Allowed characters include: Chinese characters, English characters, `0-9`,` _` and ` -`.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Source type. Valid values: 
	// <li>HTTP: the source type is HTTP.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Incremental migration HTTP origin source configuration. This field is required when the OriginType value is `HTTP`.
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

type CreateIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// Incremental migration strategy name. The name length should not exceed 100 characters. Allowed characters include: Chinese characters, English characters, `0-9`,` _` and ` -`.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Source type. Valid values: 
	// <li>HTTP: the source type is HTTP.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Incremental migration HTTP origin source configuration. This field is required when the OriginType value is `HTTP`.
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

func (r *CreateIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyName")
	delete(f, "OriginType")
	delete(f, "HttpOriginConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIncrementalMigrationStrategyResponseParams struct {
	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *CreateIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageCredentialsRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The policy string serialized by URL Encode is used as the input parameter Policy. The server will URL Decode the Policy value and grant temporary access credentials according to the parsed policy. Please pass in parameters according to the specification.
	// Note:
	// 1.The policy syntax refers to [Cloud Access Management](/document/product/598/10603).
	// 2.The policy cannot contain the principal element.
	// 3.The policy actions include: 
	// <li>name/vod:ListObjects;</li>
	// <li>name/vod:ListObjectsV2;</li>
	// <li>name/vod:HeadObject;</li>
	// <li>name/vod:PutObject;</li>
	// <li>name/vod:ListParts;</li>
	// <li>name/vod:PostObject;</li>
	// <li>name/vod:CreateMultipartUpload;</li>
	// <li>name/vod:UploadPart;</li>
	// <li>name/vod:CompleteMultipartUpload;</li>
	// <li>name/vod:AbortMultipartUpload;</li>
	// <li>name/vod:ListMultipartUploads;</li>
	// <li>name/vod:CopyObject;</li>
	// <li>name/vod:RestoreObject;</li>
	// <li>name/vod:DeleteObjects;</li>
	// <li>name/vod:DeleteObject;</li>
	// <li>name/vod:UploadPartCopy.</li>4.The resource format of the policy is `qcs::vod:[region]:uid/[AppID]:prefix//[SubAppId]/[BucketId]/[Path]`, where AppID, SubAppId, BucketId and Path need to be filled in as required.
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// Specifies the validity period of credentials in seconds. The default value is 1800 seconds and the maximum value is 129600 seconds.
	DurationSeconds *uint64 `json:"DurationSeconds,omitnil,omitempty" name:"DurationSeconds"`
}

type CreateStorageCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The policy string serialized by URL Encode is used as the input parameter Policy. The server will URL Decode the Policy value and grant temporary access credentials according to the parsed policy. Please pass in parameters according to the specification.
	// Note:
	// 1.The policy syntax refers to [Cloud Access Management](/document/product/598/10603).
	// 2.The policy cannot contain the principal element.
	// 3.The policy actions include: 
	// <li>name/vod:ListObjects;</li>
	// <li>name/vod:ListObjectsV2;</li>
	// <li>name/vod:HeadObject;</li>
	// <li>name/vod:PutObject;</li>
	// <li>name/vod:ListParts;</li>
	// <li>name/vod:PostObject;</li>
	// <li>name/vod:CreateMultipartUpload;</li>
	// <li>name/vod:UploadPart;</li>
	// <li>name/vod:CompleteMultipartUpload;</li>
	// <li>name/vod:AbortMultipartUpload;</li>
	// <li>name/vod:ListMultipartUploads;</li>
	// <li>name/vod:CopyObject;</li>
	// <li>name/vod:RestoreObject;</li>
	// <li>name/vod:DeleteObjects;</li>
	// <li>name/vod:DeleteObject;</li>
	// <li>name/vod:UploadPartCopy.</li>4.The resource format of the policy is `qcs::vod:[region]:uid/[AppID]:prefix//[SubAppId]/[BucketId]/[Path]`, where AppID, SubAppId, BucketId and Path need to be filled in as required.
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// Specifies the validity period of credentials in seconds. The default value is 1800 seconds and the maximum value is 129600 seconds.
	DurationSeconds *uint64 `json:"DurationSeconds,omitnil,omitempty" name:"DurationSeconds"`
}

func (r *CreateStorageCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Policy")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageCredentialsResponseParams struct {
	// Credentials.
	Credentials *Credentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStorageCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageCredentialsResponseParams `json:"Response"`
}

func (r *CreateStorageCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Storage region, which must be a region supported by the system. All storage regions and the regions where storage buckets have already been enabled can be queried using the [DescribeStorageRegions](https://www.tencentcloud.com/document/product/266/46542) API.
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// The name of the storage.
	// <li>Only lowercase English letters, numbers, hyphens "-", and their combinations are supported.</li>
	// <li>The storage name cannot start or end with a "-".</li>
	// <li>The maximum length of the storage name is 64 characters.</li>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`
}

type CreateStorageRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Storage region, which must be a region supported by the system. All storage regions and the regions where storage buckets have already been enabled can be queried using the [DescribeStorageRegions](https://www.tencentcloud.com/document/product/266/46542) API.
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// The name of the storage.
	// <li>Only lowercase English letters, numbers, hyphens "-", and their combinations are supported.</li>
	// <li>The storage name cannot start or end with a "-".</li>
	// <li>The maximum length of the storage name is 64 characters.</li>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`
}

func (r *CreateStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "StorageRegion")
	delete(f, "StorageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageResponseParams struct {
	// The unique identifier ID of the storage bucket.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageResponseParams `json:"Response"`
}

func (r *CreateStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// Access Key ID.
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// Secret Access Key.
	SecretAccessKey *string `json:"SecretAccessKey,omitnil,omitempty" name:"SecretAccessKey"`

	// The session token length depends on the binding policy and is no longer than 4096 bytes.
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// The expiration time of the credentials.
	Expiration *string `json:"Expiration,omitnil,omitempty" name:"Expiration"`
}

// Predefined struct for user
type DeleteIncrementalMigrationStrategyRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DeleteIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DeleteIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIncrementalMigrationStrategyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *DeleteIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationStrategyInfosRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Filter criteria. The maximum number of Filters.Values is `20`. If this parameter is not input, all stategy information under the current SubAppId will be returned. The detailed filter criteria are as follows:
	// <li>BucketId: Filter by the ID of bucket;</li>
	// <li>StrategyId: Filter by the ID of strategy.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. 
	// SortBy.Field Values include: 
	// <li>UpdateTime: (Default) Update time of the strategy.</li>SortBy.Order Values include: 
	// <li>Desc: (Default) The order is descend.</li> 
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIncrementalMigrationStrategyInfosRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Filter criteria. The maximum number of Filters.Values is `20`. If this parameter is not input, all stategy information under the current SubAppId will be returned. The detailed filter criteria are as follows:
	// <li>BucketId: Filter by the ID of bucket;</li>
	// <li>StrategyId: Filter by the ID of strategy.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. 
	// SortBy.Field Values include: 
	// <li>UpdateTime: (Default) Update time of the strategy.</li>SortBy.Order Values include: 
	// <li>Desc: (Default) The order is descend.</li> 
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Offset for paginated queries. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeIncrementalMigrationStrategyInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationStrategyInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementalMigrationStrategyInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationStrategyInfosResponseParams struct {
	// Total count of matched strategies.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of all matched strategies.
	StrategyInfoSet []*IncrementalMigrationStrategyInfo `json:"StrategyInfoSet,omitnil,omitempty" name:"StrategyInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIncrementalMigrationStrategyInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementalMigrationStrategyInfosResponseParams `json:"Response"`
}

func (r *DescribeIncrementalMigrationStrategyInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationStrategyInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Filter criteria. The maximum number of Filters.Values is `20`. If this parameter is not input, all storage under the current SubAppId will be returned. The detailed filter criteria are as follows:
	// <li>BucketId: Filter by the ID of bucket;</li>
	// <li>StorageName: Filter by the name of storage.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. 
	// SortBy.Field Values include: 
	// <li>UpdateTime: (Default) Create time of the storage.</li>SortBy.Order Values include: 
	// <li>Asc: (Default) The order is ascend.</li> 
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Offset for paginated queries. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeStorageRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// Filter criteria. The maximum number of Filters.Values is `20`. If this parameter is not input, all storage under the current SubAppId will be returned. The detailed filter criteria are as follows:
	// <li>BucketId: Filter by the ID of bucket;</li>
	// <li>StorageName: Filter by the name of storage.</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sort the returned results according to this field. 
	// SortBy.Field Values include: 
	// <li>UpdateTime: (Default) Create time of the storage.</li>SortBy.Order Values include: 
	// <li>Asc: (Default) The order is ascend.</li> 
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Offset for paginated queries. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit on paginated queries. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageResponseParams struct {
	// Total count of matched storage.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of all matched storage.
	StorageInfoSet []*StorageInfo `json:"StorageInfoSet,omitnil,omitempty" name:"StorageInfoSet"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageResponseParams `json:"Response"`
}

func (r *DescribeStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Fields to be filtered.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value of the filtered field.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type IncrementalMigrationHttpEndpointInfo struct {
	// Address information, supporting domain names or IP addresses.
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// Backup address information.
	StandbyEndpointSet []*string `json:"StandbyEndpointSet,omitnil,omitempty" name:"StandbyEndpointSet"`
}

type IncrementalMigrationHttpHeader struct {
	// Header Key.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Header Value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type IncrementalMigrationHttpHeaderInfo struct {
	// HTTP Header Passthrough Mode. Values valid:
	// <li>FOLLOW_ALL: Pass through all header information;<\li>
	// <li>FOLLOW_PART: Pass through partial header information;<\li>
	// <li>IGNORE_PART: Ignore partial header information.<\li>
	HeaderFollowMode *string `json:"HeaderFollowMode,omitnil,omitempty" name:"HeaderFollowMode"`

	// Header Key Set for Passthrough. This field is required only when the HeaderFollowMode is set to `FOLLOW_PART`.
	FollowHttpHeaderKeySet []*string `json:"FollowHttpHeaderKeySet,omitnil,omitempty" name:"FollowHttpHeaderKeySet"`

	// Add Header Key-Value Pair Collection.
	NewHttpHeaderSet []*IncrementalMigrationHttpHeader `json:"NewHttpHeaderSet,omitnil,omitempty" name:"NewHttpHeaderSet"`
}

type IncrementalMigrationHttpOriginCondition struct {
	// HTTP code to trigger back-to-source conditions. If not filled, the default value is `404`.
	HttpStatusCode *uint64 `json:"HttpStatusCode,omitnil,omitempty" name:"HttpStatusCode"`

	// Object key prefix to trigger back-to-source conditions
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type IncrementalMigrationHttpOriginConfig struct {
	// Origin information for back-to-source.
	OriginInfo *IncrementalMigrationHttpOriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// Back-to-source parameters.
	OriginParameter *IncrementalMigrationHttpOriginParameter `json:"OriginParameter,omitnil,omitempty" name:"OriginParameter"`

	// Back-to-source mode. Valid values:
	// <li>SYNC: Synchronous back-to-source;</li>
	// <li>ASYNC: Asynchronous back-to-source.</li>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Back-to-source conditions.
	OriginCondition *IncrementalMigrationHttpOriginCondition `json:"OriginCondition,omitnil,omitempty" name:"OriginCondition"`
}

type IncrementalMigrationHttpOriginInfo struct {
	// Incremental migration source address information.
	EndpointInfo *IncrementalMigrationHttpEndpointInfo `json:"EndpointInfo,omitnil,omitempty" name:"EndpointInfo"`

	// Incremental migration source file information.
	FileInfo *IncrementalMigrationOriginFileInfo `json:"FileInfo,omitnil,omitempty" name:"FileInfo"`
}

type IncrementalMigrationHttpOriginParameter struct {
	// HTTP header passthrough information.
	HttpHeaderInfo *IncrementalMigrationHttpHeaderInfo `json:"HttpHeaderInfo,omitnil,omitempty" name:"HttpHeaderInfo"`

	// Back-to-source protocol. Valid values:
	// <li>HTTP: Force HTTP;</li>
	// <li>HTTPS: Force HTTPS;</li>
	// <li>FOLLOW: Follow the request protocol.</li>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Query string passthrough mode. Valid values
	// <li>FOLLOW: Fully passthrough;</li>
	// <li>IGNORE: No passthrough at all.</li>
	QueryStringFollowMode *string `json:"QueryStringFollowMode,omitnil,omitempty" name:"QueryStringFollowMode"`

	// HTTP Code for redirection. Currently, only `301`, `302`, and `307` are supported. The default value is `302`.
	HttpRedirectCode *uint64 `json:"HttpRedirectCode,omitnil,omitempty" name:"HttpRedirectCode"`

	// Origin server redirection follow mode. Valid values:
	// <li>FOLLOW: Follow origin server redirection;</li>
	// <li>IGNORE: Ignore origin server redirection.</li>
	OriginRedirectionFollowMode *string `json:"OriginRedirectionFollowMode,omitnil,omitempty" name:"OriginRedirectionFollowMode"`
}

type IncrementalMigrationOriginFileInfo struct {
	// File prefix configuration.
	PrefixConfig *IncrementalMigrationOriginPrefixConfig `json:"PrefixConfig,omitnil,omitempty" name:"PrefixConfig"`

	// File 	suffix configuration.
	SuffixConfig *IncrementalMigrationOriginSuffixConfig `json:"SuffixConfig,omitnil,omitempty" name:"SuffixConfig"`

	// Fixed file configuration.
	FixedFileConfig *IncrementalMigrationOriginFixedFileConfig `json:"FixedFileConfig,omitnil,omitempty" name:"FixedFileConfig"`
}

type IncrementalMigrationOriginFixedFileConfig struct {
	// Fixed file path: If you fill in `example/test.png`, the origin-pull address will be: `http(s)://<origin domain>/example/test.png`.
	FixedFilePath *string `json:"FixedFilePath,omitnil,omitempty" name:"FixedFilePath"`
}

type IncrementalMigrationOriginPrefixConfig struct {
	// Origin address prefix: If you fill in `test/`, the origin-pull address will be `http(s)://<origin domain>/test/<file name>`.
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type IncrementalMigrationOriginSuffixConfig struct {
	// File suffix.
	// if filled with `.ts`, the origin-pull address will be: `http(s)://<origin domain>/<file name>.ts`.
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

type IncrementalMigrationStrategyInfo struct {
	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// The name of the incremental migration strategy.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// Source type.
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Incremental migration HTTP origin source configuration.
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

// Predefined struct for user
type ModifyIncrementalMigrationStrategyRequestParams struct {
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// The name of the incremental migration strategy.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Source type. Valid values: 
	// <li>HTTP: (Default) the source type is HTTP.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Incremental migration HTTP origin source configuration. If left blank, it will default to no modification.
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

type ModifyIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>The ID of [VOD Professional Application](http://tencentcloud.com/document/product/266/67977).</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// The bucket ID where the strategy takes effect.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The ID of the incremental migration strategy.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// The name of the incremental migration strategy.
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// Source type. Valid values: 
	// <li>HTTP: (Default) the source type is HTTP.</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// Incremental migration HTTP origin source configuration. If left blank, it will default to no modification.
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

func (r *ModifyIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyId")
	delete(f, "StrategyName")
	delete(f, "OriginType")
	delete(f, "HttpOriginConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIncrementalMigrationStrategyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *ModifyIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortBy struct {
	// Sort by field
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Sorting order. Valid values: Asc (ascending), Desc (descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type StorageInfo struct {
	// The ID of bucket.
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// The name of bucket.
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`

	// The region of storage.
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// The status of the internet access domain name is stored. Valid values: <li>ONLINE: Active;</li> <li>DEPLOYING: In deployment.</li>
	InternetAccessDomainStatus *string `json:"InternetAccessDomainStatus,omitnil,omitempty" name:"InternetAccessDomainStatus"`

	// The internet access domain name of storage.
	InternetAccessDomain *string `json:"InternetAccessDomain,omitnil,omitempty" name:"InternetAccessDomain"`

	// The creation time of the storage.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}