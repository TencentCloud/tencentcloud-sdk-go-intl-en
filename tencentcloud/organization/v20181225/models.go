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

package v20181225

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AcceptOrganizationInvitationRequestParams struct {
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type AcceptOrganizationInvitationRequest struct {
	*tchttp.BaseRequest
	
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *AcceptOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptOrganizationInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptOrganizationInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptOrganizationInvitationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AcceptOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *AcceptOrganizationInvitationResponseParams `json:"Response"`
}

func (r *AcceptOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeRequestParams struct {
	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Organizational unit name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Organizational unit name
	Name *string `json:"Name,omitempty" name:"Name"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeResponseParams struct {
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CancelOrganizationInvitationRequestParams struct {
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type CancelOrganizationInvitationRequest struct {
	*tchttp.BaseRequest
	
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CancelOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationInvitationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationInvitationResponseParams `json:"Response"`
}

func (r *CancelOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationRequestParams struct {
	// Organization type; currently its value is fixed as `1`
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`
}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Organization type; currently its value is fixed as `1`
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`
}

func (r *CreateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationResponseParams struct {
	// Organization ID
	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`

	// Creator's name
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Creator's email address
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Organization type
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationResponseParams `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberFromNodeRequestParams struct {
	// UIN of the member to be deleted
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
}

type DeleteOrganizationMemberFromNodeRequest struct {
	*tchttp.BaseRequest
	
	// UIN of the member to be deleted
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationMemberFromNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberFromNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMemberFromNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberFromNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOrganizationMemberFromNodeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMemberFromNodeResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMemberFromNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberFromNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// List of UINs of members to be deleted
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// List of UINs of members to be deleted
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
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
	delete(f, "Uins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Organizational unit ID list
	NodeIds []*uint64 `json:"NodeIds,omitempty" name:"NodeIds"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// Organizational unit ID list
	NodeIds []*uint64 `json:"NodeIds,omitempty" name:"NodeIds"`
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
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteOrganizationRequestParams struct {

}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DenyOrganizationInvitationRequestParams struct {
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DenyOrganizationInvitationRequest struct {
	*tchttp.BaseRequest
	
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DenyOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DenyOrganizationInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DenyOrganizationInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DenyOrganizationInvitationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DenyOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *DenyOrganizationInvitationResponseParams `json:"Response"`
}

func (r *DenyOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DenyOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOrganizationMemberRequestParams struct {
	// Organization member UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type GetOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// Organization member UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *GetOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOrganizationMemberResponseParams struct {
	// Organization member UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Organization member name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Joining time 
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Organizational unit name
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *GetOrganizationMemberResponseParams `json:"Response"`
}

func (r *GetOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOrganizationRequestParams struct {

}

type GetOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOrganizationResponseParams struct {
	// Organization ID
	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`

	// Creator UIN
	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`

	// Creator's name
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Creator's email address
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Organization type
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

	// Whether the organization is empty or not 
	IsEmpty *uint64 `json:"IsEmpty,omitempty" name:"IsEmpty"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *GetOrganizationResponseParams `json:"Response"`
}

func (r *GetOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationInvitationsRequestParams struct {
	// Whether to list the invitations you received or the invitations you sent. `1`: list the invitations you received; `0`: list the invitations you sent.
	Invited *uint64 `json:"Invited,omitempty" name:"Invited"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListOrganizationInvitationsRequest struct {
	*tchttp.BaseRequest
	
	// Whether to list the invitations you received or the invitations you sent. `1`: list the invitations you received; `0`: list the invitations you sent.
	Invited *uint64 `json:"Invited,omitempty" name:"Invited"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationInvitationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationInvitationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Invited")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationInvitationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationInvitationsResponseParams struct {
	// List of invitations
	Invitations []*OrgInvitation `json:"Invitations,omitempty" name:"Invitations"`

	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListOrganizationInvitationsResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationInvitationsResponseParams `json:"Response"`
}

func (r *ListOrganizationInvitationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationInvitationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationMembersRequestParams struct {
	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationMembersResponseParams struct {
	// Member list
	Members []*OrgMember `json:"Members,omitempty" name:"Members"`

	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationMembersResponseParams `json:"Response"`
}

func (r *ListOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationNodeMembersRequestParams struct {
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest
	
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationNodeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationNodeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationNodeMembersResponseParams struct {
	// Total number of results
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Member list
	Members []*OrgMember `json:"Members,omitempty" name:"Members"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationNodeMembersResponseParams `json:"Response"`
}

func (r *ListOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationNodesRequestParams struct {

}

type ListOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationNodesResponseParams struct {
	// Organizational unit list
	Nodes []*OrgNode `json:"Nodes,omitempty" name:"Nodes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationNodesResponseParams `json:"Response"`
}

func (r *ListOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationMembersToNodeRequestParams struct {
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Member UIN list
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
}

type MoveOrganizationMembersToNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Member UIN list
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`
}

func (r *MoveOrganizationMembersToNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationMembersToNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Uins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveOrganizationMembersToNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationMembersToNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MoveOrganizationMembersToNodeResponse struct {
	*tchttp.BaseResponse
	Response *MoveOrganizationMembersToNodeResponseParams `json:"Response"`
}

func (r *MoveOrganizationMembersToNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationMembersToNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgInvitation struct {
	// Invitation ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// UIN of the invitee
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Creator UIN
	HostUin *uint64 `json:"HostUin,omitempty" name:"HostUin"`

	// Creator's name
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Creator's email address
	HostMail *string `json:"HostMail,omitempty" name:"HostMail"`

	// Invitation status. `-1`: expired; `0`: normal; `1`: accepted; `2`: invalid; `3`: cancelled
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Organization type
	OrgType *uint64 `json:"OrgType,omitempty" name:"OrgType"`

	// Time of invitation
	InviteTime *string `json:"InviteTime,omitempty" name:"InviteTime"`

	// Expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type OrgMember struct {
	// UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Joining time
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`
}

type OrgNode struct {
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Number of members
	MemberCount *uint64 `json:"MemberCount,omitempty" name:"MemberCount"`
}

// Predefined struct for user
type QuitOrganizationRequestParams struct {
	// Organization ID
	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
}

type QuitOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Organization ID
	OrgId *uint64 `json:"OrgId,omitempty" name:"OrgId"`
}

func (r *QuitOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuitOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QuitOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *QuitOrganizationResponseParams `json:"Response"`
}

func (r *QuitOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendOrganizationInvitationRequestParams struct {
	// UIN of the invitee
	InviteUin *uint64 `json:"InviteUin,omitempty" name:"InviteUin"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type SendOrganizationInvitationRequest struct {
	*tchttp.BaseRequest
	
	// UIN of the invitee
	InviteUin *uint64 `json:"InviteUin,omitempty" name:"InviteUin"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *SendOrganizationInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOrganizationInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InviteUin")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendOrganizationInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendOrganizationInvitationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendOrganizationInvitationResponse struct {
	*tchttp.BaseResponse
	Response *SendOrganizationInvitationResponseParams `json:"Response"`
}

func (r *SendOrganizationInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOrganizationInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberRequestParams struct {
	// Member UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN
	MemberUin *uint64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Notes
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeRequestParams struct {
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organizational unit ID
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parent organizational unit ID
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`
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
	delete(f, "ParentNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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