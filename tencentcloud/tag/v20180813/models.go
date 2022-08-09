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

package v20180813

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddResourceTagRequestParams struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type AddResourceTagRequest struct {
	*tchttp.BaseRequest
	
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *AddResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddResourceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddResourceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddResourceTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *AddResourceTagResponseParams `json:"Response"`
}

func (r *AddResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachResourcesTagRequestParams struct {
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type AttachResourcesTagRequest struct {
	*tchttp.BaseRequest
	
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *AttachResourcesTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachResourcesTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachResourcesTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachResourcesTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachResourcesTagResponse struct {
	*tchttp.BaseResponse
	Response *AttachResourcesTagResponseParams `json:"Response"`
}

func (r *AttachResourcesTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachResourcesTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRequestParams struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type CreateTagRequest struct {
	*tchttp.BaseRequest
	
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *CreateTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagResponseParams `json:"Response"`
}

func (r *CreateTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagsRequestParams struct {
	// Tag list.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateTagsRequest struct {
	*tchttp.BaseRequest
	
	// Tag list.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagsResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagsResponseParams `json:"Response"`
}

func (r *CreateTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceTagRequestParams struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type DeleteResourceTagRequest struct {
	*tchttp.BaseRequest
	
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *DeleteResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceTagResponseParams `json:"Response"`
}

func (r *DeleteResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRequestParams struct {
	// The tag key to be deleted.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// The tag value to be deleted.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DeleteTagRequest struct {
	*tchttp.BaseRequest
	
	// The tag key to be deleted.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// The tag value to be deleted.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *DeleteTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagResponseParams `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsRequestParams struct {
	// Tag list.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type DeleteTagsRequest struct {
	*tchttp.BaseRequest
	
	// Tag list.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagsResponseParams `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsRequestParams struct {
	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Array of resource IDs (up to 50)
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// The resource's region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourceTagsByResourceIdsRequest struct {
	*tchttp.BaseRequest
	
	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Array of resource IDs (up to 50)
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// The resource's region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceTagsByResourceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceIds")
	delete(f, "ResourceRegion")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByResourceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list.
	Tags []*TagResource `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByResourceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByResourceIdsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsSeqRequestParams struct {
	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeResourceTagsByResourceIdsSeqRequest struct {
	*tchttp.BaseRequest
	
	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceTagsByResourceIdsSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceIds")
	delete(f, "ResourceRegion")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByResourceIdsSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByResourceIdsSeqResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list
	Tags []*TagResource `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByResourceIdsSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByResourceIdsSeqResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByResourceIdsSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByTagKeysRequestParams struct {
	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Unique resource ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Resource tag key
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Number of entries per page. Default value: 400
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeResourceTagsByTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Resource region
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Unique resource ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Resource tag key
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Number of entries per page. Default value: 400
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourceTagsByTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceRegion")
	delete(f, "ResourceIds")
	delete(f, "TagKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsByTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsByTagKeysResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource tag
	Rows []*ResourceIdTag `json:"Rows,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsByTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsByTagKeysResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsByTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsByTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsRequestParams struct {
	// Creator `uin`
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Resource region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID. Queries with `ResourceId` only may be slow or fail to return results. We recommend you also enter `ServiceType`, `ResourcePrefix`, and `ResourceRegion` (which can be ignored for resources that don't have the region attribute) when entering `ResourceId`.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether it is a COS resource (0 or 1). This parameter is required when the entered `ResourceId` is a COS resource.
	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

type DescribeResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// Creator `uin`
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Resource region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID. Queries with `ResourceId` only may be slow or fail to return results. We recommend you also enter `ServiceType`, `ResourcePrefix`, and `ResourceRegion` (which can be ignored for resources that don't have the region attribute) when entering `ResourceId`.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether it is a COS resource (0 or 1). This parameter is required when the entered `ResourceId` is a COS resource.
	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateUin")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CosResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource tag
	Rows []*TagResource `json:"Rows,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsRequestParams struct {
	// Tag filtering arrays.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// Tag creator uin.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 15.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// The resource's region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeResourcesByTagsRequest struct {
	*tchttp.BaseRequest
	
	// Tag filtering arrays.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// Tag creator uin.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 15.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// The resource's region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesByTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagFilters")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page.
	// Note: This field may return null, indicating that no valid value is found.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource tag.
	Rows []*ResourceTag `json:"Rows,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByTagsResponseParams `json:"Response"`
}

func (r *DescribeResourcesByTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsUnionRequestParams struct {
	// Tag filtering arrays.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// Tag creator uin.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 15.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// The resource’s region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type DescribeResourcesByTagsUnionRequest struct {
	*tchttp.BaseRequest
	
	// Tag filtering arrays.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// Tag creator uin.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 15.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// The resource’s region.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesByTagsUnionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsUnionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagFilters")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourcePrefix")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesByTagsUnionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesByTagsUnionResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The size of each page.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Resource tag.
	Rows []*ResourceTag `json:"Rows,omitempty" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourcesByTagsUnionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesByTagsUnionResponseParams `json:"Response"`
}

func (r *DescribeResourcesByTagsUnionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesByTagsUnionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagKeysRequestParams struct {
	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to show project
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to show project
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagKeysResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list.
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagKeysResponseParams `json:"Response"`
}

func (r *DescribeTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesRequestParams struct {
	// Tag key list.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// Tag key list.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesResponseParams `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesSeqRequestParams struct {
	// Tag key list
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Creator `Uin`. If this parameter is blank or left empty, only `Uin` will be used as a condition for query
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagValuesSeqRequest struct {
	*tchttp.BaseRequest
	
	// Tag key list
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Creator `Uin`. If this parameter is blank or left empty, only `Uin` will be used as a condition for query
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagValuesSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "CreateUin")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesSeqResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagValuesSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesSeqResponseParams `json:"Response"`
}

func (r *DescribeTagValuesSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// Tag key. Either exists or does not exist alongside the tag value. If it does not exist, all of the user's tags will be queried.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value. Either exists or does not exist alongside the tag key. If it does not exist, all of the user's tags will be queried.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag key array, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried. If it is passed in together with `TagKey`, it will be used and the `TagKey` will be ignored.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Whether to show project tag
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// Tag key. Either exists or does not exist alongside the tag value. If it does not exist, all of the user's tags will be queried.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value. Either exists or does not exist alongside the tag key. If it does not exist, all of the user's tags will be queried.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag key array, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried. If it is passed in together with `TagKey`, it will be used and the `TagKey` will be ignored.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Whether to show project tag
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreateUin")
	delete(f, "TagKeys")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// Total number of results.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list.
	Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsResponseParams `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsSeqRequestParams struct {
	// Tag key, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value, which either exists or does not exist with the tag key. If it does not exist, all tags of the user will be queried
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Creator `Uin`. If this parameter is blank or left empty, only `Uin` will be used as a condition for query
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag key array, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried. If it is passed in together with `TagKey`, it will be used and the `TagKey` will be ignored.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Whether to show project tag
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

type DescribeTagsSeqRequest struct {
	*tchttp.BaseRequest
	
	// Tag key, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value, which either exists or does not exist with the tag key. If it does not exist, all tags of the user will be queried
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Creator `Uin`. If this parameter is blank or left empty, only `Uin` will be used as a condition for query
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag key array, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried. If it is passed in together with `TagKey`, it will be used and the `TagKey` will be ignored.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Whether to show project tag
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagsSeqRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsSeqRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreateUin")
	delete(f, "TagKeys")
	delete(f, "ShowProject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsSeqRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsSeqResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Data offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag list
	Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsSeqResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsSeqResponseParams `json:"Response"`
}

func (r *DescribeTagsSeqResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsSeqResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachResourcesTagRequestParams struct {
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key to be unbound
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type DetachResourcesTagRequest struct {
	*tchttp.BaseRequest
	
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key to be unbound
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *DetachResourcesTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachResourcesTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachResourcesTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachResourcesTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachResourcesTagResponse struct {
	*tchttp.BaseResponse
	Response *DetachResourcesTagResponseParams `json:"Response"`
}

func (r *DetachResourcesTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachResourcesTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedResource struct {
	// Six-segment descriptions of failed resources
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Error code
	Code *string `json:"Code,omitempty" name:"Code"`

	// Error message
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type GetResourcesRequestParams struct {
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// If this parameter is passed in, the list of all matching resources will be returned, and the specified `MaxResults` will become invalid.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key and value.
	// If multiple tags are specified, resources bound to all such tags will be queried.
	// Value range of N: 0–5
	// There can be up to 10 `TagValues` in each `TagFilters`.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 200).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// If this parameter is passed in, the list of all matching resources will be returned, and the specified `MaxResults` will become invalid.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key and value.
	// If multiple tags are specified, resources bound to all such tags will be queried.
	// Value range of N: 0–5
	// There can be up to 10 `TagValues` in each `TagFilters`.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 200).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "TagFilters")
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourcesResponseParams struct {
	// Token value obtained for the next page
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// List of resources and their associated tags (key-value pairs)
	ResourceTagMappingList []*ResourceTagMapping `json:"ResourceTagMappingList,omitempty" name:"ResourceTagMappingList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetResourcesResponse struct {
	*tchttp.BaseResponse
	Response *GetResourcesResponseParams `json:"Response"`
}

func (r *GetResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagKeysRequestParams struct {
	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetTagKeysRequest struct {
	*tchttp.BaseRequest
	
	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagKeysResponseParams struct {
	// Token value obtained for the next page
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Tag key information.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *GetTagKeysResponseParams `json:"Response"`
}

func (r *GetTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagValuesRequestParams struct {
	// Tag key.
	// All tag values corresponding to the list of tag keys.
	// Maximum length: 20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type GetTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// Tag key.
	// All tag values corresponding to the list of tag keys.
	// Maximum length: 20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *GetTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagValuesResponseParams struct {
	// Token value obtained for the next page
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *GetTagValuesResponseParams `json:"Response"`
}

func (r *GetTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagsRequestParams struct {
	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// Tag key.
	// All tags corresponding to the list of tag keys.
	// Maximum length: 20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

type GetTagsRequest struct {
	*tchttp.BaseRequest
	
	// The token value of the next page obtained from the response of the previous page.
	// Leave it empty for the first request.
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Number of data entries to return per page (up to 1,000).
	// Default value: 50.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// Tag key.
	// All tags corresponding to the list of tag keys.
	// Maximum length: 20
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *GetTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PaginationToken")
	delete(f, "MaxResults")
	delete(f, "TagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTagsResponseParams struct {
	// Token value obtained for the next page
	PaginationToken *string `json:"PaginationToken,omitempty" name:"PaginationToken"`

	// Tag list.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTagsResponse struct {
	*tchttp.BaseResponse
	Response *GetTagsResponseParams `json:"Response"`
}

func (r *GetTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsRequestParams struct {
	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// The tags to be added or modified. If the resource described by `Resource` is not associated with the input tag keys, an association will be added. If the tag keys are already associated, the values corresponding to the associated tag keys will be modified to the input values. This API must contain either `ReplaceTags` or `DeleteTag`, and these two parameters cannot include the same tag keys. This parameter can be omitted, but it cannot be an empty array.
	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// The tags to be disassociated. This API must contain either `ReplaceTags` or `DeleteTag`, and these two parameters cannot include the same tag keys. This parameter can be omitted, but it cannot be an empty array.
	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// The tags to be added or modified. If the resource described by `Resource` is not associated with the input tag keys, an association will be added. If the tag keys are already associated, the values corresponding to the associated tag keys will be modified to the input values. This API must contain either `ReplaceTags` or `DeleteTag`, and these two parameters cannot include the same tag keys. This parameter can be omitted, but it cannot be an empty array.
	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// The tags to be disassociated. This API must contain either `ReplaceTags` or `DeleteTag`, and these two parameters cannot include the same tag keys. This parameter can be omitted, but it cannot be an empty array.
	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Resource")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceTagsResponseParams `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagValueRequestParams struct {
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

type ModifyResourcesTagValueRequest struct {
	*tchttp.BaseRequest
	
	// Resource service name (the third segment in the six-segment resource description)
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource ID array, which can contain up to 50 resources
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource region. If resources have the region attribute, this field is required; otherwise, it is optional.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Resource prefix (the part before "/" in the last segment in the six-segment resource description), which is optional for COS buckets but required for other Tencent Cloud resources.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`
}

func (r *ModifyResourcesTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "ResourceIds")
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "ResourceRegion")
	delete(f, "ResourcePrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcesTagValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagValueResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourcesTagValueResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcesTagValueResponseParams `json:"Response"`
}

func (r *ModifyResourcesTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceIdTag struct {
	// Unique resource ID
	// Note: this field may return null, indicating that no valid values can be obtained
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tag key-value pair
	// Note: this field may return null, indicating that no valid values can be obtained
	TagKeyValues []*Tag `json:"TagKeyValues,omitempty" name:"TagKeyValues"`
}

type ResourceTag struct {
	// The resource's region.
	// Note: This field may return null, indicating that no valid value is found.
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// Service type.
	// Note: This field may return null, indicating that no valid value is found.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix.
	// Note: This field may return null, indicating that no valid value is found.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	// Note: This field may return null, indicating that no valid value is found.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Resource tag.
	// Note: This field may return null, indicating that no valid value is found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ResourceTagMapping struct {
	// Six-segment resource description. Tencent Cloud uses a six-segment value to describe a resource.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// List of tags associated with the resource
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Tag struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value array. '**OR**' relation if multiple values.
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagKeyObject struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

type TagResource struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource ID.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tag key MD5 value.
	TagKeyMd5 *string `json:"TagKeyMd5,omitempty" name:"TagKeyMd5"`

	// Tag value MD5 value.
	TagValueMd5 *string `json:"TagValueMd5,omitempty" name:"TagValueMd5"`

	// Resource type
	// Note: this field may return null, indicating that no valid values found.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

// Predefined struct for user
type TagResourcesRequestParams struct {
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource. For more information, see [CAM](https://intl.cloud.tencent.com/document/product/598/67350?from_cn_redirect=1) > Overview > API List > Six-Segment Resource Information.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key and value.
	// If multiple tags are specified, all such tags will be created and bound to the specified resources.
	// For each resource, each tag key can have only one value. If you try to add an existing tag key, the corresponding tag value will be updated to the new value.
	// Non-existent tags will be automatically created.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type TagResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource. For more information, see [CAM](https://intl.cloud.tencent.com/document/product/598/67350?from_cn_redirect=1) > Overview > API List > Six-Segment Resource Information.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key and value.
	// If multiple tags are specified, all such tags will be created and bound to the specified resources.
	// For each resource, each tag key can have only one value. If you try to add an existing tag key, the corresponding tag value will be updated to the new value.
	// Non-existent tags will be automatically created.
	// Value range of N: 0–9
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *TagResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TagResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TagResourcesResponseParams struct {
	// Information of failed resources.
	// When tag creating and binding succeeds, the returned `FailedResources` will be empty.
	// When tag creating and binding partially or completely fails, the returned `FailedResources` will display the details of failed resources.
	FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TagResourcesResponse struct {
	*tchttp.BaseResponse
	Response *TagResourcesResponseParams `json:"Response"`
}

func (r *TagResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagWithDelete struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// If deletion is allowed.
	CanDelete *uint64 `json:"CanDelete,omitempty" name:"CanDelete"`
}

// Predefined struct for user
type UnTagResourcesRequestParams struct {
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource. For more information, see [CAM](https://intl.cloud.tencent.com/document/product/598/67350?from_cn_redirect=1) > Overview > API List > Six-Segment Resource Information.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key.
	// Value range: 0–9
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

type UnTagResourcesRequest struct {
	*tchttp.BaseRequest
	
	// Six-segment resource description list. Tencent Cloud uses a six-segment value to describe a resource. For more information, see [CAM](https://intl.cloud.tencent.com/document/product/598/67350?from_cn_redirect=1) > Overview > API List > Six-Segment Resource Information.
	// For example, ResourceList.1 = qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}.
	// Value range of N: 0–9
	ResourceList []*string `json:"ResourceList,omitempty" name:"ResourceList"`

	// Tag key.
	// Value range: 0–9
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

func (r *UnTagResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnTagResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceList")
	delete(f, "TagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnTagResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnTagResourcesResponseParams struct {
	// Information of failed resources.
	// When tag unbinding succeeds, the returned `FailedResources` will be empty.
	// When tag unbinding partially or completely fails, the returned `FailedResources` will display the details of failed resources.
	FailedResources []*FailedResource `json:"FailedResources,omitempty" name:"FailedResources"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnTagResourcesResponse struct {
	*tchttp.BaseResponse
	Response *UnTagResourcesResponseParams `json:"Response"`
}

func (r *UnTagResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnTagResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceTagValueRequestParams struct {
	// Tag key associated with the resource.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Modified tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

type UpdateResourceTagValueRequest struct {
	*tchttp.BaseRequest
	
	// Tag key associated with the resource.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Modified tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// [Six-segment resource description](https://intl.cloud.tencent.com/document/product/598/10606?from_cn_redirect=1)
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *UpdateResourceTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceTagValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "TagValue")
	delete(f, "Resource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceTagValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceTagValueResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateResourceTagValueResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceTagValueResponseParams `json:"Response"`
}

func (r *UpdateResourceTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceTagValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}