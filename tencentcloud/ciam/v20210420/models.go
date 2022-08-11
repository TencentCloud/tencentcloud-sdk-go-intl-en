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

package v20210420

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type ListUserGroupsRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Number of queried pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Query conditions (user group ID or user group name)
	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

type ListUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Number of queried pages
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Query conditions (user group ID or user group name)
	Condition *string `json:"Condition,omitempty" name:"Condition"`
}

func (r *ListUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Page")
	delete(f, "Size")
	delete(f, "Condition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsResponseParams struct {
	// User group list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content []*UserGroup `json:"Content,omitempty" name:"Content"`

	// Total number
	// Note: this field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Pagination
	// Note: this field may return null, indicating that no valid values can be obtained.
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListUserGroupsResponseParams `json:"Response"`
}

func (r *ListUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Pageable struct {
	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page number
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type UserGroup struct {
	// User group ID
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// User group name
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User group description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Tenant ID
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
}