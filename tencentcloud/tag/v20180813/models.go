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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddResourceTagRequest struct {
	*tchttp.BaseRequest

	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource description in six-piece format.
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *AddResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddResourceTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddResourceTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceTagRequest struct {
	*tchttp.BaseRequest

	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Resource description in six-piece format.
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *DeleteResourceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteResourceTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteResourceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteResourceTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByResourceIdsRequest struct {
	*tchttp.BaseRequest

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Resource prefix.
	ResourcePrefix *string `json:"ResourcePrefix,omitempty" name:"ResourcePrefix"`

	// Unique resource ID.
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// The resource’s region.
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

func (r *DescribeResourceTagsByResourceIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByResourceIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Page size.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Tag list.
		Tags []*TagResource `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsByResourceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceTagsByResourceIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// Resource tag key
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

	// Number of entries per page. Default value: 400
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourceTagsByTagKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceTagsByTagKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsByTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Number of entries per page
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Resource tag
		Rows []*ResourceIdTag `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsByTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceTagsByTagKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// Unique resource ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Data offset. Default value: 0. It must be an integer multiple of the `Limit` parameter
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 15
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether it is a COS resource ID
	CosResourceId *uint64 `json:"CosResourceId,omitempty" name:"CosResourceId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Number of entries per page.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Resource tag
		Rows []*TagResource `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagsRequest struct {
	*tchttp.BaseRequest

	// Tag filtering arrays.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters" list`

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

	// Service type.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *DescribeResourcesByTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesByTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesByTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Number of entries per page.
	// Note: This field may return null, indicating that no valid value is found.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Resource tag.
		Rows []*ResourceTag `json:"Rows,omitempty" name:"Rows" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesByTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourcesByTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeTagKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Page size.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Tag list.
		Tags []*string `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest

	// Tag key list.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

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

func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Page size.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Tag list.
		Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// Tag key. Either exists or does not exist alongside the tag value. If it does not exist, all of the user’s tags will be queried.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value. Either exists or does not exist alongside the tag key. If it does not exist, all of the user’s tags will be queried.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Data offset. The default value is 0. Must be an integral multiple of the `Limit` parameter.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. The default value is 0.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Creator `Uin`. If not specified, `Uin` is only used as the query condition.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tag key array, which either exists or does not exist with the tag value. If it does not exist, all tags of the user will be queried. If it is passed in together with `TagKey`, it will be used and the `TagKey` will be ignored
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

	// Whether to show project tag
	ShowProject *uint64 `json:"ShowProject,omitempty" name:"ShowProject"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of results.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Data offset.
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// Page size.
		Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

		// Tag list.
		Tags []*TagWithDelete `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest

	// Resource description in six-piece format.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// The tags to be added or modified. If the resource described by `Resource` is not associated with the input tag keys, an association will be added. If the tag keys are already associated, the values corresponding to the associated tag keys will be modified to the input values. This API must contain either `ReplaceTags` or `DeleteTag`. And these two parameters cannot include the same tag keys.
	ReplaceTags []*Tag `json:"ReplaceTags,omitempty" name:"ReplaceTags" list`

	// The tags to be unassociated. This API must contain either `ReplaceTags` or `DeleteTag`. And these two parameters cannot include the same tag keys.
	DeleteTags []*TagKeyObject `json:"DeleteTags,omitempty" name:"DeleteTags" list`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResourceIdTag struct {

	// Unique resource ID
	// Note: this field may return null, indicating that no valid values can be obtained
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Tag key-value pair
	// Note: this field may return null, indicating that no valid values can be obtained
	TagKeyValues []*Tag `json:"TagKeyValues,omitempty" name:"TagKeyValues" list`
}

type ResourceTag struct {

	// The resource’s region.
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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
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

	// Tag value array. “**OR**” relation if multiple values.
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue" list`
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

type TagWithDelete struct {

	// Tag key.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// If deletion is allowed.
	CanDelete *uint64 `json:"CanDelete,omitempty" name:"CanDelete"`
}

type UpdateResourceTagValueRequest struct {
	*tchttp.BaseRequest

	// Tag key associated with the resource.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Modified tag value.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Resource description in six-piece format.
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

func (r *UpdateResourceTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateResourceTagValueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceTagValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResourceTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateResourceTagValueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
