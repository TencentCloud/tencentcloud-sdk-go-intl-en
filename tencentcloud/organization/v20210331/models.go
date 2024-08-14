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

package v20210331

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddOrganizationNodeRequestParams struct {
	// Parent node ID, which can be obtained through the `DescribeOrganizationNodes` API.
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Parent node ID, which can be obtained through the `DescribeOrganizationNodes` API.
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentNodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeResponseParams struct {
	// Node ID.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationNodeResponseParams `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID, which can be obtained through the `DescribeOrganizationMemberPolicies` API.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of sub-account UINs of the organization admin, which can contain up to five UINs.
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID, which can be obtained through the `DescribeOrganizationMemberPolicies` API.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// List of sub-account UINs of the organization admin, which can contain up to five UINs.
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Organization sub-account UIN.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Organization sub-account UIN.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrgServiceAssignRequestParams struct {
	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Uin list of the delegated admins, including up to 20 items.
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// Management scope of the delegated admin. Valid values: 1 (all members), 2 (partial members). Default value: 1.
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// Uin list of the managed members. This parameter is valid when ManagementScope is 2.
	ManagementScopeUins []*int64 `json:"ManagementScopeUins,omitnil,omitempty" name:"ManagementScopeUins"`

	// ID list of the managed departments. This parameter is valid when ManagementScope is 2.
	ManagementScopeNodeIds []*int64 `json:"ManagementScopeNodeIds,omitnil,omitempty" name:"ManagementScopeNodeIds"`
}

type CreateOrgServiceAssignRequest struct {
	*tchttp.BaseRequest
	
	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Uin list of the delegated admins, including up to 20 items.
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// Management scope of the delegated admin. Valid values: 1 (all members), 2 (partial members). Default value: 1.
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// Uin list of the managed members. This parameter is valid when ManagementScope is 2.
	ManagementScopeUins []*int64 `json:"ManagementScopeUins,omitnil,omitempty" name:"ManagementScopeUins"`

	// ID list of the managed departments. This parameter is valid when ManagementScope is 2.
	ManagementScopeNodeIds []*int64 `json:"ManagementScopeNodeIds,omitnil,omitempty" name:"ManagementScopeNodeIds"`
}

func (r *CreateOrgServiceAssignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgServiceAssignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "MemberUins")
	delete(f, "ManagementScope")
	delete(f, "ManagementScopeUins")
	delete(f, "ManagementScopeNodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrgServiceAssignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrgServiceAssignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrgServiceAssignResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrgServiceAssignResponseParams `json:"Response"`
}

func (r *CreateOrgServiceAssignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy name, which can contain up to 128 letters, digits, and symbols `+=,.@_-`.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Member access identity ID, which can be obtained through the `DescribeOrganizationMemberAuthIdentities` API.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy name, which can contain up to 128 letters, digits, and symbols `+=,.@_-`.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Member access identity ID, which can be obtained through the `DescribeOrganizationMemberAuthIdentities` API.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyResponseParams struct {
	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberRequestParams struct {
	// Member name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Relationship policy. Valid value: `Financial`.
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// List of member financial permission IDs. `7` indicates paying, which is the default value.
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// ID of the node of the member's department, which can be obtained through the `DescribeOrganizationNodes` API.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Account name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Member creation record ID, which is required during retry upon creation exception.
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Payer UIN, which is required during paying for a member.
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// List of member access identity IDs, which can be obtained through the `ListOrganizationIdentity` API. `1` indicates supported, which is the default value.
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// Verified entity relationship ID, which is required during creating members for different entities.
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// Member name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Relationship policy. Valid value: `Financial`.
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// List of member financial permission IDs. `7` indicates paying, which is the default value.
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// ID of the node of the member's department, which can be obtained through the `DescribeOrganizationNodes` API.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Account name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Member creation record ID, which is required during retry upon creation exception.
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Payer UIN, which is required during paying for a member.
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// List of member access identity IDs, which can be obtained through the `ListOrganizationIdentity` API. `1` indicates supported, which is the default value.
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// Verified entity relationship ID, which is required during creating members for different entities.
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "NodeId")
	delete(f, "AccountName")
	delete(f, "Remark")
	delete(f, "RecordId")
	delete(f, "PayUin")
	delete(f, "IdentityRoleID")
	delete(f, "AuthRelationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberResponseParams struct {
	// Member UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrgServiceAssignRequestParams struct {
	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Uin of the delegated admin.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteOrgServiceAssignRequest struct {
	*tchttp.BaseRequest
	
	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Uin of the delegated admin.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteOrgServiceAssignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgServiceAssignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrgServiceAssignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrgServiceAssignResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrgServiceAssignResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrgServiceAssignResponseParams `json:"Response"`
}

func (r *DeleteOrgServiceAssignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// List of UINs of the members to be deleted.
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// List of UINs of the members to be deleted.
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesRequestParams struct {
	// List of node IDs.
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// List of node IDs.
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationNodesResponseParams `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsRequestParams struct {
	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned results.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsResponseParams struct {
	// List
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberAuthAccount `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthAccountsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesRequestParams struct {
	// Offset, which is an integer multiple of the value of `Limit`. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit, which defaults to `10`. Value range: 1-50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Organization member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DescribeOrganizationMemberAuthIdentitiesRequest struct {
	*tchttp.BaseRequest
	
	// Offset, which is an integer multiple of the value of `Limit`. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit, which defaults to `10`. Value range: 1-50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Organization member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesResponseParams struct {
	// List of authorizable identities
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberAuthIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthIdentitiesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesRequestParams struct {
	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Search keyword, which can be the policy name or description.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeOrganizationMemberPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Search keyword, which can be the policy name or description.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMemberPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesResponseParams struct {
	// List.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberPoliciesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersRequestParams struct {
	// Offset, which is an integer multiple of the value of `Limit`. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit, which defaults to `10`. Value range: 1-50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// Search by member name or ID.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Entity name.
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// Offset, which is an integer multiple of the value of `Limit`. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit, which defaults to `10`. Value range: 1-50.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// Search by member name or ID.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Entity name.
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Lang")
	delete(f, "SearchKey")
	delete(f, "AuthName")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersResponseParams struct {
	// Member list.
	Items []*OrgMember `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMembersResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesRequestParams struct {
	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesResponseParams struct {
	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgNode `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationNodesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {
	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// Organization ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// Creator UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostUin *int64 `json:"HostUin,omitnil,omitempty" name:"HostUin"`

	// Creator name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Organization type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgType *int64 `json:"OrgType,omitnil,omitempty" name:"OrgType"`

	// Whether the member is the organization admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsManager *bool `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// Policy type. Valid values: `Financial` (finance management).
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// List of member financial permissions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// Organization root node ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootNodeId *int64 `json:"RootNodeId,omitnil,omitempty" name:"RootNodeId"`

	// Organization creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Member joining time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// Payer UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// Payer name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// Whether the member is the trusted service admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// Whether the member is the verified entity admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdentityPolicy struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

// Predefined struct for user
type ListOrgServiceAssignMemberRequestParams struct {
	// Offset. Its value must be an integer multiple of the limit. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit. Value range: 1-50. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

type ListOrgServiceAssignMemberRequest struct {
	*tchttp.BaseRequest
	
	// Offset. Its value must be an integer multiple of the limit. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit. Value range: 1-50. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Organization service ID, which can be obtained through [ListOrganizationService](https://intl.cloud.tencent.com/document/product/850/109561?from_cn_redirect=1).
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

func (r *ListOrgServiceAssignMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrgServiceAssignMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrgServiceAssignMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrgServiceAssignMemberResponseParams struct {
	// Total quantity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List of the delegated admins.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrganizationServiceAssignMember `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrgServiceAssignMemberResponse struct {
	*tchttp.BaseResponse
	Response *ListOrgServiceAssignMemberResponseParams `json:"Response"`
}

func (r *ListOrgServiceAssignMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrgServiceAssignMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityRequestParams struct {
	// Offset.  It must be an integer multiple of the value of `Limit`.  Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit for the number of query results.  Value range:  1-50.  Default value:  `10`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search by name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Search by identity ID.
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity type.  Valid values: `1` (Preset), `2` (Custom).
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

type ListOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Offset.  It must be an integer multiple of the value of `Limit`.  Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The limit for the number of query results.  Value range:  1-50.  Default value:  `10`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search by name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Search by identity ID.
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity type.  Valid values: `1` (Preset), `2` (Custom).
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

func (r *ListOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "IdentityId")
	delete(f, "IdentityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityResponseParams struct {
	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Item details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationIdentityResponseParams `json:"Response"`
}

func (r *ListOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationServiceRequestParams struct {
	// Offset. Its value must be an integer multiple of the limit. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit. Value range: 1-50. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword for search by name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type ListOrganizationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Offset. Its value must be an integer multiple of the limit. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit. Value range: 1-50. Default value: 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword for search by name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *ListOrganizationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationServiceResponseParams struct {
	// Total quantity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Organization service list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrganizationServiceAssign `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationServiceResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationServiceResponseParams `json:"Response"`
}

func (r *ListOrganizationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`
}

type MemberMainInfo struct {
	// Member UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Member name j.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

// Predefined struct for user
type MoveOrganizationNodeMembersRequestParams struct {
	// Organization node ID.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Member UIN list.
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Member UIN list.
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveOrganizationNodeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationNodeMembersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *MoveOrganizationNodeMembersResponseParams `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeMainInfo struct {
	// Department ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Department name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type OrgIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Identity policy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// Identity type. Valid values: `1` (preset); `2` (custom).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgMember struct {
	// Member UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Member name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Member type. Valid values: `Invite` (invited); `Create` (created).
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberType *string `json:"MemberType,omitnil,omitempty" name:"MemberType"`

	// Relationship policy type
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// Relationship policy name
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// Relationship policy permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// Node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node name
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// Payer UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// Payer name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// Management identity
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitnil,omitempty" name:"OrgIdentity"`

	// Security information binding status. Valid values: `Unbound`, `Valid`, `Success`, `Failed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// Member permission status. Valid values: `Confirmed`, `UnConfirmed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PermissionStatus *string `json:"PermissionStatus,omitnil,omitempty" name:"PermissionStatus"`
}

type OrgMemberAuthAccount struct {
	// Organization sub-account UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity role name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// Identity role alias.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Sub-account name
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgSubAccountName *string `json:"OrgSubAccountName,omitnil,omitempty" name:"OrgSubAccountName"`
}

type OrgMemberAuthIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Role name of an identity
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// Role alias of an identity
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Identity type (`1`: Preset; `2`: Custom)
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

type OrgMemberPolicy struct {
	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// Identity role name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// Identity role alias.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgNode struct {
	// Organization node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parent node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParentNodeId *int64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgPermission struct {
	// Permission ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Permission name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type OrganizationServiceAssign struct {
	// Organization service ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Organization service product name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Whether to support delegation. Valid values: 1 (yes), 2 (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAssign *uint64 `json:"IsAssign,omitnil,omitempty" name:"IsAssign"`

	// Organization service description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Number of the current delegated admins.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberNum *string `json:"MemberNum,omitnil,omitempty" name:"MemberNum"`

	// Help documentation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Document *string `json:"Document,omitnil,omitempty" name:"Document"`

	// Console path of the organization service product.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsoleUrl *string `json:"ConsoleUrl,omitnil,omitempty" name:"ConsoleUrl"`

	// Whether to access the usage status. Valid values: 1 (yes), 
	//  2 (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsUsageStatus *uint64 `json:"IsUsageStatus,omitnil,omitempty" name:"IsUsageStatus"`

	// Limit for the number of delegated admins.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanAssignCount *uint64 `json:"CanAssignCount,omitnil,omitempty" name:"CanAssignCount"`

	// Organization service product identifier.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Whether to support organization service authorization. Valid values: 1 (yes), 2 (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceGrant *uint64 `json:"ServiceGrant,omitnil,omitempty" name:"ServiceGrant"`

	// Enabling status of organization service authorization. This field is valid when ServiceGrant is 1. Valid values: Enabled, Disabled. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GrantStatus *string `json:"GrantStatus,omitnil,omitempty" name:"GrantStatus"`

	// Whether to support setting the delegated management scope. Valid values: 1 (yes), 2 (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSetManagementScope *uint64 `json:"IsSetManagementScope,omitnil,omitempty" name:"IsSetManagementScope"`
}

type OrganizationServiceAssignMember struct {
	// Organization service ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Organization service product name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Uin of the delegated admin.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Name of the delegated admin.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`

	// Activation status. Valid values: 0 (the service has no activation status), 1 (activated), 2 (not activated).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageStatus *uint64 `json:"UsageStatus,omitnil,omitempty" name:"UsageStatus"`

	// Delegation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Management scope of the delegated admin. Valid values: 1 (all members), 2 (partial members).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// Uin list of managed members. This parameter is valid when ManagementScope is 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagementScopeMembers []*MemberMainInfo `json:"ManagementScopeMembers,omitnil,omitempty" name:"ManagementScopeMembers"`

	// ID list of the managed departments. This parameter is valid when ManagementScope is 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagementScopeNodes []*NodeMainInfo `json:"ManagementScopeNodes,omitnil,omitempty" name:"ManagementScopeNodes"`
}

// Predefined struct for user
type UpdateOrganizationNodeRequestParams struct {
	// Node ID.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Node ID.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationNodeResponseParams `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}