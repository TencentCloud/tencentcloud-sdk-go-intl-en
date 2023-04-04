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

package v20220817

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type AddGroupMemberRequestParams struct {
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type AddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *AddGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddGroupMemberResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *AddGroupMemberResponseParams `json:"Response"`
}

func (r *AddGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppCustomContent struct {
	// Multiple scenarios can be set for an application.
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// Logo URL
	LogoUrl *string `json:"LogoUrl,omitempty" name:"LogoUrl"`

	// Homepage URL, which can be used for redirection
	HomeUrl *string `json:"HomeUrl,omitempty" name:"HomeUrl"`

	// Custom JS URL
	JsUrl *string `json:"JsUrl,omitempty" name:"JsUrl"`

	// Custom CSS URL
	CssUrl *string `json:"CssUrl,omitempty" name:"CssUrl"`
}

type BackgroundPictureConfig struct {
	// The URL of the background image.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`
}

// Predefined struct for user
type BatchAddGroupMemberRequestParams struct {
	// The target group IDs. Array length limit: 100.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users to add. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchAddGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// The target group IDs. Array length limit: 100.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users to add. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchAddGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAddGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddGroupMemberResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchAddGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *BatchAddGroupMemberResponseParams `json:"Response"`
}

func (r *BatchAddGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateGroupWithMembersRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The information of the groups to create. Array length limit: 256.
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitempty" name:"GroupBaseInfos"`

	// The group members. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchCreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The information of the groups to create. Array length limit: 256.
	GroupBaseInfos []*GroupBaseInfo `json:"GroupBaseInfos,omitempty" name:"GroupBaseInfos"`

	// The group members. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchCreateGroupWithMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateGroupWithMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "GroupBaseInfos")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateGroupWithMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateGroupWithMembersResponseParams struct {
	// The IDs of the groups created, which are in the same order as the elements in the request parameter `GroupBaseInfos.N`.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchCreateGroupWithMembersResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateGroupWithMembersResponseParams `json:"Response"`
}

func (r *BatchCreateGroupWithMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateGroupWithMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateRoomRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The information of the rooms to create.
	RoomInfos []*RoomInfo `json:"RoomInfos,omitempty" name:"RoomInfos"`
}

type BatchCreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The information of the rooms to create.
	RoomInfos []*RoomInfo `json:"RoomInfos,omitempty" name:"RoomInfos"`
}

func (r *BatchCreateRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateRoomResponseParams struct {
	// The IDs of the rooms created, which are in the same order as they are passed in.
	RoomIds []*uint64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchCreateRoomResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateRoomResponseParams `json:"Response"`
}

func (r *BatchCreateRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteGroupMemberRequestParams struct {
	// The target group IDs. Array length limit: 100.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users to remove. Array length limit: 256.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type BatchDeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// The target group IDs. Array length limit: 100.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users to remove. Array length limit: 256.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *BatchDeleteGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteGroupMemberResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteGroupMemberResponseParams `json:"Response"`
}

func (r *BatchDeleteGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRecordRequestParams struct {
	// The room IDs.
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type BatchDeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// The room IDs.
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *BatchDeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomIds")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRecordResponseParams struct {
	// The IDs of the rooms whose recordings are deleted. Note: This field may return null, indicating that no valid values can be obtained.
	RoomIds []*int64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteRecordResponseParams `json:"Response"`
}

func (r *BatchDeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterRequestParams struct {
	// The information of the users to register.
	Users []*BatchUserRequest `json:"Users,omitempty" name:"Users"`
}

type BatchRegisterRequest struct {
	*tchttp.BaseRequest
	
	// The information of the users to register.
	Users []*BatchUserRequest `json:"Users,omitempty" name:"Users"`
}

func (r *BatchRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterResponseParams struct {
	// The information of the successfully registered users. Note: This field may return null, indicating that no valid values can be obtained.
	Users []*BatchUserInfo `json:"Users,omitempty" name:"Users"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchRegisterResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterResponseParams `json:"Response"`
}

func (r *BatchRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchUserInfo struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

type BatchUserRequest struct {
	// The SDKAppID assigned by LCIC.  Note: This field may return null, indicating that no valid values can be obtained.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The username.  Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The user’s ID in your system, which must be unique across the same application.  Note: This field may return null, indicating that no valid values can be obtained.
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// The user’s profile photo.  Note: This field may return null, indicating that no valid values can be obtained.
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

// Predefined struct for user
type BindDocumentToRoomRequestParams struct {
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Document ID
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// Binding type. The default value is `0`. The backend passes through this parameter to clients so that the clients can implement business logic based on this parameter.
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

type BindDocumentToRoomRequest struct {
	*tchttp.BaseRequest
	
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Document ID
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// Binding type. The default value is `0`. The backend passes through this parameter to clients so that the clients can implement business logic based on this parameter.
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *BindDocumentToRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDocumentToRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "DocumentId")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDocumentToRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDocumentToRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindDocumentToRoomResponse struct {
	*tchttp.BaseResponse
	Response *BindDocumentToRoomResponseParams `json:"Response"`
}

func (r *BindDocumentToRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDocumentToRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Document URL	
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// Document name	
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// Document owner ID	
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Transcoding type. Valid values: `0`: No transcoding required (default); `1`: Documents that need to be transcoded: ppt, pptx, pdf, doc, docx; `2`: Videos that need to be transcoded: mp4, 3pg, mpeg, avi, flv, wmv, rm, h264, etc.; `2`: Audio that needs to be transcoded: mp3, wav, wma, aac, flac, opus
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// Permission. Valid values: `0`: Private document (default); `1`: Public document
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// Document extension
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// Document size, in bytes
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Document URL	
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// Document name	
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// Document owner ID	
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Transcoding type. Valid values: `0`: No transcoding required (default); `1`: Documents that need to be transcoded: ppt, pptx, pdf, doc, docx; `2`: Videos that need to be transcoded: mp4, 3pg, mpeg, avi, flv, wmv, rm, h264, etc.; `2`: Audio that needs to be transcoded: mp3, wav, wma, aac, flac, opus
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// Permission. Valid values: `0`: Private document (default); `1`: Public document
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// Document extension
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// Document size, in bytes
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`
}

func (r *CreateDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "DocumentUrl")
	delete(f, "DocumentName")
	delete(f, "Owner")
	delete(f, "TranscodeType")
	delete(f, "Permission")
	delete(f, "DocumentType")
	delete(f, "DocumentSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// Document ID
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDocumentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocumentResponseParams `json:"Response"`
}

func (r *CreateDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithMembersRequestParams struct {
	// The group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The group members. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type CreateGroupWithMembersRequest struct {
	*tchttp.BaseRequest
	
	// The group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The group members. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *CreateGroupWithMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "SdkAppId")
	delete(f, "TeacherId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupWithMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithMembersResponseParams struct {
	// The ID of the successfully created group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupWithMembersResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupWithMembersResponseParams `json:"Response"`
}

func (r *CreateGroupWithMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithSubGroupRequestParams struct {
	// The group name after merging.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The IDs of the groups to merge. Duplicate group IDs are not allowed. Array length limit: 40.
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

type CreateGroupWithSubGroupRequest struct {
	*tchttp.BaseRequest
	
	// The group name after merging.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The IDs of the groups to merge. Duplicate group IDs are not allowed. Array length limit: 40.
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

func (r *CreateGroupWithSubGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithSubGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "SdkAppId")
	delete(f, "SubGroupIds")
	delete(f, "TeacherId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupWithSubGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupWithSubGroupResponseParams struct {
	// The ID of the merged group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupWithSubGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupWithSubGroupResponseParams `json:"Response"`
}

func (r *CreateGroupWithSubGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupWithSubGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomRequestParams struct {
	// Room name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Reserved room start time, in UNIX timestamp format
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Reserved room end time, in UNIX timestamp format
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 	Resolution. Valid values: 1: SD; 2: HD; 3: FHD
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// Maximum number of mic-on users (excluding teachers). Value range: [0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// The room subtype. Valid values: videodoc: Document + Video; video: Video only.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// The user ID of the teacher. User IDs are returned by the user registration APIs. The user specified will have teacher permissions in the room created.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// Whether to automatically turn the mic on when the user enters a room. Valid values: 0: No (default value); 1: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to enable the high audio quality mode. Valid values: 0: No (default value); 1: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// Whether to disable auto recording. Valid values: 0: No (default); 1: Yes. If this parameter is 0, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// The user IDs of the teaching assistants. User IDs are returned by the user registration APIs. The users specified will have teaching assistant permissions in the room created.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// Recording layout
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// The ID of the group to bind. If you specify this parameter, only members of the group can enter this room.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type CreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// Room name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Reserved room start time, in UNIX timestamp format
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Reserved room end time, in UNIX timestamp format
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 	Resolution. Valid values: 1: SD; 2: HD; 3: FHD
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// Maximum number of mic-on users (excluding teachers). Value range: [0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// The room subtype. Valid values: videodoc: Document + Video; video: Video only.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// The user ID of the teacher. User IDs are returned by the user registration APIs. The user specified will have teacher permissions in the room created.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// Whether to automatically turn the mic on when the user enters a room. Valid values: 0: No (default value); 1: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to enable the high audio quality mode. Valid values: 0: No (default value); 1: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// Whether to disable auto recording. Valid values: 0: No (default); 1: Yes. If this parameter is 0, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// The user IDs of the teaching assistants. User IDs are returned by the user registration APIs. The users specified will have teaching assistant permissions in the room created.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// Recording layout
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// The ID of the group to bind. If you specify this parameter, only members of the group can enter this room.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "Resolution")
	delete(f, "MaxMicNumber")
	delete(f, "SubType")
	delete(f, "TeacherId")
	delete(f, "AutoMic")
	delete(f, "AudioQuality")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "RecordLayout")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomResponseParams struct {
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoomResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoomResponseParams `json:"Response"`
}

func (r *CreateRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user IDs.
	Users []*string `json:"Users,omitempty" name:"Users"`
}

type CreateSupervisorRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user IDs.
	Users []*string `json:"Users,omitempty" name:"Users"`
}

func (r *CreateSupervisorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSupervisorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSupervisorResponse struct {
	*tchttp.BaseResponse
	Response *CreateSupervisorResponseParams `json:"Response"`
}

func (r *CreateSupervisorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppCustomContentRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The custom elements (for which a scene has been configured) to delete. If this is empty, all custom elements will be deleted.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

type DeleteAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The custom elements (for which a scene has been configured) to delete. If this is empty, all custom elements will be deleted.
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

func (r *DeleteAppCustomContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppCustomContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Scenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppCustomContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppCustomContentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAppCustomContentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppCustomContentResponseParams `json:"Response"`
}

func (r *DeleteAppCustomContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppCustomContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocumentRequestParams struct {
	// The document ID.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type DeleteDocumentRequest struct {
	*tchttp.BaseRequest
	
	// The document ID.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DeleteDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocumentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDocumentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocumentResponseParams `json:"Response"`
}

func (r *DeleteDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupMemberRequestParams struct {
	// The group ID. You cannot remove members from a merged group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

type DeleteGroupMemberRequest struct {
	*tchttp.BaseRequest
	
	// The group ID. You cannot remove members from a merged group.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The users. Array length limit: 200.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`
}

func (r *DeleteGroupMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "MemberIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupMemberResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupMemberResponseParams `json:"Response"`
}

func (r *DeleteGroupMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// The IDs of the groups to delete.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of the groups to delete.
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRequestParams struct {
	// The room ID.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// The room ID.
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordResponseParams `json:"Response"`
}

func (r *DeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomRequestParams struct {
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

type DeleteRoomRequest struct {
	*tchttp.BaseRequest
	
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DeleteRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoomResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoomResponseParams `json:"Response"`
}

func (r *DeleteRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentMemberListRequestParams struct {
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The page to return records from. Pagination starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCurrentMemberListRequest struct {
	*tchttp.BaseRequest
	
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The page to return records from. Pagination starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCurrentMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentMemberListResponseParams struct {
	// The total number of records.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The user list.
	MemberRecords []*MemberRecord `json:"MemberRecords,omitempty" name:"MemberRecords"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCurrentMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentMemberListResponseParams `json:"Response"`
}

func (r *DescribeCurrentMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeveloperRequestParams struct {

}

type DescribeDeveloperRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeveloperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeveloperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeveloperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeveloperResponseParams struct {

	DeveloperId *string `json:"DeveloperId,omitempty" name:"DeveloperId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeveloperResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeveloperResponseParams `json:"Response"`
}

func (r *DescribeDeveloperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeveloperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentRequestParams struct {
	// The (unique) document ID.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type DescribeDocumentRequest struct {
	*tchttp.BaseRequest
	
	// The (unique) document ID.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DescribeDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentResponseParams struct {
	// The document ID.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// The document’s original URL.
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// The document title.
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// The user ID of the document’s owner.
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The document access type.
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// The transcoding result. If the file is not transcoded, this parameter will be empty. If it is successfully transcoded, this parameter will be the URL of the transcoded file. If transcoding fails, this parameter will indicate the error code.
	TranscodeResult *string `json:"TranscodeResult,omitempty" name:"TranscodeResult"`

	// The transcoding type.
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// The transcoding progress. Value range: 0-100.
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitempty" name:"TranscodeProgress"`

	// The transcoding status. 0: The file is not transcoded. 1: The file is being transcoded. 2: Transcoding failed. 3: Transcoding is successful.
	TranscodeState *uint64 `json:"TranscodeState,omitempty" name:"TranscodeState"`

	// The error message for failed transcoding.
	TranscodeInfo *string `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// The document type.
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// The document size (bytes).
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`

	// The time (Unix timestamp) when the document was last updated.
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDocumentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocumentResponseParams `json:"Response"`
}

func (r *DescribeDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentsByRoomRequestParams struct {
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. Pagination starts from 1, which is also the default value of this parameter.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records to return per page. The maximum value can be 1000. The default value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The document access type. [0]: The private documents of the owner. [1]: The public documents of the owner. [0,1]: The private and public documents of the owner. [2]: The private and public documents of all users (including the owner). Default value: [2].
	Permission []*uint64 `json:"Permission,omitempty" name:"Permission"`

	// The user ID of the document owner. If you do not specify this, the information of all documents under the application will be returned.
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type DescribeDocumentsByRoomRequest struct {
	*tchttp.BaseRequest
	
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. Pagination starts from 1, which is also the default value of this parameter.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records to return per page. The maximum value can be 1000. The default value is 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The document access type. [0]: The private documents of the owner. [1]: The public documents of the owner. [0,1]: The private and public documents of the owner. [2]: The private and public documents of all users (including the owner). Default value: [2].
	Permission []*uint64 `json:"Permission,omitempty" name:"Permission"`

	// The user ID of the document owner. If you do not specify this, the information of all documents under the application will be returned.
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *DescribeDocumentsByRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsByRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Permission")
	delete(f, "Owner")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocumentsByRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocumentsByRoomResponseParams struct {
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Documents []*DocumentInfo `json:"Documents,omitempty" name:"Documents"`

	// The total number of records that meet the conditions.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDocumentsByRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocumentsByRoomResponseParams `json:"Response"`
}

func (r *DescribeDocumentsByRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocumentsByRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupListRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. Pagination starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000 and is 20 by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The user ID of the teacher, which is used as the filter. This parameter and MemberId are mutually exclusive. If both are specified, only this parameter will take effect.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The user ID of a member, which is used as the filter. This parameter and TeacherId are mutually exclusive.
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
}

type DescribeGroupListRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. Pagination starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000 and is 20 by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The user ID of the teacher, which is used as the filter. This parameter and MemberId are mutually exclusive. If both are specified, only this parameter will take effect.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The user ID of a member, which is used as the filter. This parameter and TeacherId are mutually exclusive.
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
}

func (r *DescribeGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "TeacherId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupListResponseParams struct {
	// The total number of groups that meet the conditions.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupListResponseParams `json:"Response"`
}

func (r *DescribeGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupMemberListRequestParams struct {
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. The default value is 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000 and is 20 by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGroupMemberListRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. The default value is 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records per page. The value of this parameter cannot exceed 1000 and is 20 by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupMemberListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupMemberListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupMemberListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupMemberListResponseParams struct {
	// The total number of records that meet the conditions.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberIds []*string `json:"MemberIds,omitempty" name:"MemberIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupMemberListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupMemberListResponseParams `json:"Response"`
}

func (r *DescribeGroupMemberListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupMemberListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupRequestParams struct {
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The group type. 0: Ordinary group. 1: Merged group. If the group queried is a merged group, the IDs of the sub-groups will be returned.
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubGroupIds []*string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupResponseParams `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomRequestParams struct {
	// Room ID	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// Room ID	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomResponseParams struct {
	// Room name	
	Name *string `json:"Name,omitempty" name:"Name"`

	// Reserved room start time, in UNIX timestamp format	
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// Reserved room end time, in UNIX timestamp format	
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// Teacher ID	
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// LCIC SdkAppId	
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Resolution. Valid values: 1: SD; 2: HD; 3: FHD
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// Maximum number of mic-on users (excluding teachers). Value range: [0, 16]
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// Whether to automatically turn the mic on when the user enters a room. Valid values: 0: No (default value); 1: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to enable the high audio quality mode. Valid values: 0: No (default value); 1: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// The room subtype. Valid values: videodoc: Document + Video; video: Video only.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// Whether to disable auto recording. Valid values: 0: No (default); 1: Yes. If this parameter is 0, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// Assistant ID list Note: This field may return null, indicating that no valid values can be obtained.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// Recording URL. This parameter exists only after a room is ended. Note: This field may return null, indicating that no valid values can be obtained.
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`

	// The class status. 0: The class has not started. 1: The class has started. 2: The class ended. 3: The class expired. Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomResponseParams `json:"Response"`
}

func (r *DescribeRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomStatisticsRequestParams struct {
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Current page in pagination, which starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of data entries to return per page. Maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRoomStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Room ID
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Current page in pagination, which starts from 1.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of data entries to return per page. Maximum value: 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRoomStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomStatisticsResponseParams struct {
	// Peak number of online members
	PeakMemberNumber *uint64 `json:"PeakMemberNumber,omitempty" name:"PeakMemberNumber"`

	// Accumulated number of online members
	MemberNumber *uint64 `json:"MemberNumber,omitempty" name:"MemberNumber"`

	// Total number of records, including members who entered the room and members who should attend the class but did not
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// Member record list
	MemberRecords []*MemberRecord `json:"MemberRecords,omitempty" name:"MemberRecords"`

	// The actual start time of the room, in Unix timestamp, accurate to seconds. Note: This field may return null, indicating that no valid values can be obtained.
	RealStartTime *uint64 `json:"RealStartTime,omitempty" name:"RealStartTime"`

	// The actual end time of the room, in Unix timestamp, accurate to seconds. Note: This field may return null, indicating that no valid values can be obtained.
	RealEndTime *uint64 `json:"RealEndTime,omitempty" name:"RealEndTime"`

	// The total message count of the room.
	MessageCount *uint64 `json:"MessageCount,omitempty" name:"MessageCount"`

	// The total number of mic-on students in the room.
	MicCount *uint64 `json:"MicCount,omitempty" name:"MicCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomStatisticsResponseParams `json:"Response"`
}

func (r *DescribeRoomStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSdkAppIdUsersRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. The default value is 1.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records to return per page. The default value is 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSdkAppIdUsersRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The page to return records from. The default value is 1.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records to return per page. The default value is 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSdkAppIdUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSdkAppIdUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSdkAppIdUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSdkAppIdUsersResponseParams struct {
	// The total number of users.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The information of the users. Note: This field may return null, indicating that no valid values can be obtained.
	Users []*UserInfo `json:"Users,omitempty" name:"Users"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSdkAppIdUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSdkAppIdUsersResponseParams `json:"Response"`
}

func (r *DescribeSdkAppIdUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSdkAppIdUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// User ID	
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID	
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// The application ID.	
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User ID	
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username	
	Name *string `json:"Name,omitempty" name:"Name"`

	// URL of user profile photo.	
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentInfo struct {
	// The document ID. Note: This field may return null, indicating that no valid values can be obtained.
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// The document’s original URL. Note: This field may return null, indicating that no valid values can be obtained.
	DocumentUrl *string `json:"DocumentUrl,omitempty" name:"DocumentUrl"`

	// The document title. Note: This field may return null, indicating that no valid values can be obtained.
	DocumentName *string `json:"DocumentName,omitempty" name:"DocumentName"`

	// The user ID of the document’s owner. Note: This field may return null, indicating that no valid values can be obtained.
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// The application ID. Note: This field may return null, indicating that no valid values can be obtained.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The document access type. 0: Private; 1: Public. Note: This field may return null, indicating that no valid values can be obtained.
	Permission *uint64 `json:"Permission,omitempty" name:"Permission"`

	// The transcoding result. If the file is not transcoded, this parameter will be empty. If it is successfully transcoded, this parameter will be the URL of the transcoded file. If transcoding fails, this parameter will indicate the error code. Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeResult *string `json:"TranscodeResult,omitempty" name:"TranscodeResult"`

	// The transcoding type. Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeType *uint64 `json:"TranscodeType,omitempty" name:"TranscodeType"`

	// The transcoding progress. Value range: 0-100. Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeProgress *uint64 `json:"TranscodeProgress,omitempty" name:"TranscodeProgress"`

	// The transcoding status. 0: The file is not transcoded. 1: The file is being transcoded. 2: Transcoding failed. 3: Transcoding is successful. Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeState *uint64 `json:"TranscodeState,omitempty" name:"TranscodeState"`

	// The error message for failed transcoding. Note: This field may return null, indicating that no valid values can be obtained.
	TranscodeInfo *string `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// The document type. Note: This field may return null, indicating that no valid values can be obtained.
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// The document size (bytes). Note: This field may return null, indicating that no valid values can be obtained.
	DocumentSize *uint64 `json:"DocumentSize,omitempty" name:"DocumentSize"`

	// The time (Unix timestamp) when the document was last updated. Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type EventDataInfo struct {
	// The room ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The ID of the user to whom the event occurred.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type EventInfo struct {
	// The Unix timestamp (seconds) when the event occurred.
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The event type. Valid values:
	// `RoomStart`: The class started. `RoomEnd`: The class ended. `MemberJoin`: A user joined. `MemberQuit`: A user left. `RecordFinish`: Recording is finished.
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// The details of the event, including the room ID and the user to whom the event occurred.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventData *EventDataInfo `json:"EventData,omitempty" name:"EventData"`
}

// Predefined struct for user
type GetRoomEventRequestParams struct {
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The starting page. Pagination starts from 1. This parameter is valid only if `keyword` is empty.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records (up to 200) per page. This parameter is valid only if `keyword` is empty.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The type of events to query. Valid values:
	// `RoomStart`: The class started.
	// `RoomEnd`: The class ended.
	// `MemberJoin`: A user joined.
	// `MemberQuit`: A user left.
	// `RecordFinish`: Recording is finished.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type GetRoomEventRequest struct {
	*tchttp.BaseRequest
	
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The starting page. Pagination starts from 1. This parameter is valid only if `keyword` is empty.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// The maximum number of records (up to 200) per page. This parameter is valid only if `keyword` is empty.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The type of events to query. Valid values:
	// `RoomStart`: The class started.
	// `RoomEnd`: The class ended.
	// `MemberJoin`: A user joined.
	// `MemberQuit`: A user left.
	// `RecordFinish`: Recording is finished.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *GetRoomEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomEventResponseParams struct {
	// The total number of events for the room. The value of this parameter is not affected by `keyword`.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The event details, including the type and time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Events []*EventInfo `json:"Events,omitempty" name:"Events"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRoomEventResponse struct {
	*tchttp.BaseResponse
	Response *GetRoomEventResponseParams `json:"Response"`
}

func (r *GetRoomEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomMessageRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The starting message sequence. Messages before this sequence will be returned (not including the message whose sequence is `seq`).
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// The maximum number of messages to return. The value of this parameter cannot exceed the maximum message count allowed by your package.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetRoomMessageRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The starting message sequence. Messages before this sequence will be returned (not including the message whose sequence is `seq`).
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// The maximum number of messages to return. The value of this parameter cannot exceed the maximum message count allowed by your package.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetRoomMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "Seq")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoomMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoomMessageResponseParams struct {
	// The message list.
	Messages []*MessageList `json:"Messages,omitempty" name:"Messages"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRoomMessageResponse struct {
	*tchttp.BaseResponse
	Response *GetRoomMessageResponseParams `json:"Response"`
}

func (r *GetRoomMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoomMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWatermarkRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type GetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *GetWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWatermarkResponseParams struct {
	// The watermark settings for the teacher’s video. Note: This field may return null, indicating that no valid values can be obtained.
	TeacherLogo *WatermarkConfig `json:"TeacherLogo,omitempty" name:"TeacherLogo"`

	// The watermark settings for the whiteboard. Note: This field may return null, indicating that no valid values can be obtained.
	BoardLogo *WatermarkConfig `json:"BoardLogo,omitempty" name:"BoardLogo"`

	// The background image. Note: This field may return null, indicating that no valid values can be obtained.
	BackgroundPicture *BackgroundPictureConfig `json:"BackgroundPicture,omitempty" name:"BackgroundPicture"`

	// The watermark text. Note: This field may return null, indicating that no valid values can be obtained.
	Text *TextMarkConfig `json:"Text,omitempty" name:"Text"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *GetWatermarkResponseParams `json:"Response"`
}

func (r *GetWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupBaseInfo struct {
	// The group names. Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// The user ID of the teacher. Note: This field may return null, indicating that no valid values can be obtained.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`
}

type GroupInfo struct {
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubGroupIds *string `json:"SubGroupIds,omitempty" name:"SubGroupIds"`
}

// Predefined struct for user
type LoginOriginIdRequestParams struct {
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User's ID in the customer system, which should be unique under the same application
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

type LoginOriginIdRequest struct {
	*tchttp.BaseRequest
	
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// User's ID in the customer system, which should be unique under the same application
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

func (r *LoginOriginIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "OriginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginOriginIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOriginIdResponseParams struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Login status token returned after successful login or registration. The token is valid for seven days.
	Token *string `json:"Token,omitempty" name:"Token"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginOriginIdResponse struct {
	*tchttp.BaseResponse
	Response *LoginOriginIdResponseParams `json:"Response"`
}

func (r *LoginOriginIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserRequestParams struct {
	// User ID obtained during registration
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type LoginUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID obtained during registration
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *LoginUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserResponseParams struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Login status token returned after successful login or registration. The token is valid for seven days.
	Token *string `json:"Token,omitempty" name:"Token"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginUserResponse struct {
	*tchttp.BaseResponse
	Response *LoginUserResponseParams `json:"Response"`
}

func (r *LoginUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberRecord struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Online duration, in seconds
	PresentTime *uint64 `json:"PresentTime,omitempty" name:"PresentTime"`

	// Whether the camera is enabled
	Camera *uint64 `json:"Camera,omitempty" name:"Camera"`

	// Whether the mic is enabled
	Mic *uint64 `json:"Mic,omitempty" name:"Mic"`

	// Whether the user is muted
	Silence *uint64 `json:"Silence,omitempty" name:"Silence"`

	// Number of questions answered by the user
	AnswerQuestions *uint64 `json:"AnswerQuestions,omitempty" name:"AnswerQuestions"`

	// Number of hand raising times
	HandUps *uint64 `json:"HandUps,omitempty" name:"HandUps"`

	// First time that the user entered the room, in UNIX timestamp format
	FirstJoinTimestamp *uint64 `json:"FirstJoinTimestamp,omitempty" name:"FirstJoinTimestamp"`

	// Last time that the user left the room, in UNIX timestamp format
	LastQuitTimestamp *uint64 `json:"LastQuitTimestamp,omitempty" name:"LastQuitTimestamp"`

	// Number of rewards received
	Rewords *uint64 `json:"Rewords,omitempty" name:"Rewords"`

	// The user’s IP address.
	IPAddress *string `json:"IPAddress,omitempty" name:"IPAddress"`

	// The user’s location.
	Location *string `json:"Location,omitempty" name:"Location"`

	// The user’s device type. 0: Unknown; 1: Windows; 2: macOS; 3: Android; 4: iOS; 5: Web; 6: Mobile webpage; 7: Weixin Mini Program.
	Device *int64 `json:"Device,omitempty" name:"Device"`

	// The number of times a user turned their mic on.
	PerMemberMicCount *int64 `json:"PerMemberMicCount,omitempty" name:"PerMemberMicCount"`

	// The number of messages sent by a user.
	PerMemberMessageCount *int64 `json:"PerMemberMessageCount,omitempty" name:"PerMemberMessageCount"`
}

type MessageItem struct {
	// The message type. `0`: Text; `1`: Image.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MessageType *int64 `json:"MessageType,omitempty" name:"MessageType"`

	// The text. This parameter is valid if `MessageType` is `0`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TextMessage *string `json:"TextMessage,omitempty" name:"TextMessage"`

	// The image URL. This parameter is valid if `MessageType` is `1`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ImageMessage *string `json:"ImageMessage,omitempty" name:"ImageMessage"`
}

type MessageList struct {
	// The message timestamp.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The sender.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FromAccount *string `json:"FromAccount,omitempty" name:"FromAccount"`

	// The message sequence, which is unique across a class. The earlier a message is sent, the lower the sequence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// The message content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MessageBody []*MessageItem `json:"MessageBody,omitempty" name:"MessageBody"`
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL. Currently, only port 80 and port 443 are supported.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// The callback key.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// LCIC SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Callback URL. Currently, only port 80 and port 443 are supported.
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// The callback key.
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *ModifyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppResponseParams `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// The ID of the group to modify.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The new group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the group to modify.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The user ID of the teacher.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The new group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SdkAppId")
	delete(f, "TeacherId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomRequestParams struct {
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room start time (Unix timestamp).
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// The room end time (Unix timestamp).
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// The user ID of the teacher. User IDs are returned by the user registration APIs.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The room name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The resolution. Valid values: 1: SD; 2: HD; 3: FHD.
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// The maximum number of mic-on users (excluding the teacher). Value range: 0-16.
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// Whether to automatically turn the mic on when a user enters the room. Valid values: 0: No (default value); 1: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to enable the high audio quality mode. Valid values: 0: No (default value); 1: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// The room subtype. Valid values: videodoc: Document + Video; video: Video only; coteaching: Dual-teacher.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// Whether to disable auto recording. Valid values: 0: No (default); 1: Yes. If this parameter is 0, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// The user IDs of the teacher assistants. User IDs are returned by the user registration APIs.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// The ID of the group to bind.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ModifyRoomRequest struct {
	*tchttp.BaseRequest
	
	// The room ID.
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The room start time (Unix timestamp).
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// The room end time (Unix timestamp).
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// The user ID of the teacher. User IDs are returned by the user registration APIs.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// The room name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The resolution. Valid values: 1: SD; 2: HD; 3: FHD.
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// The maximum number of mic-on users (excluding the teacher). Value range: 0-16.
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// Whether to automatically turn the mic on when a user enters the room. Valid values: 0: No (default value); 1: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to enable the high audio quality mode. Valid values: 0: No (default value); 1: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// The room subtype. Valid values: videodoc: Document + Video; video: Video only; coteaching: Dual-teacher.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// Whether to disable auto recording. Valid values: 0: No (default); 1: Yes. If this parameter is 0, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// The user IDs of the teacher assistants. User IDs are returned by the user registration APIs.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// The ID of the group to bind.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TeacherId")
	delete(f, "Name")
	delete(f, "Resolution")
	delete(f, "MaxMicNumber")
	delete(f, "AutoMic")
	delete(f, "AudioQuality")
	delete(f, "SubType")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRoomResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoomResponseParams `json:"Response"`
}

func (r *ModifyRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserProfileRequestParams struct {
	// The ID of the user whose profile will be modified.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The new username to use.
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// The URL of the new profile photo.
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type ModifyUserProfileRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the user whose profile will be modified.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The new username to use.
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// The URL of the new profile photo.
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *ModifyUserProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Nickname")
	delete(f, "Avatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserProfileResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserProfileResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserProfileResponseParams `json:"Response"`
}

func (r *ModifyUserProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserRequestParams struct {
	// LCIC SdkAppId	
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Username	
	Name *string `json:"Name,omitempty" name:"Name"`

	// User's ID in the customer system, which should be unique under the same application	
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// User's profile photo	
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type RegisterUserRequest struct {
	*tchttp.BaseRequest
	
	// LCIC SdkAppId	
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// Username	
	Name *string `json:"Name,omitempty" name:"Name"`

	// User's ID in the customer system, which should be unique under the same application	
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// User's profile photo	
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *RegisterUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "OriginId")
	delete(f, "Avatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserResponseParams struct {
	// User ID	
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Login status token returned after successful login or registration. The token is valid for seven days.	
	Token *string `json:"Token,omitempty" name:"Token"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterUserResponse struct {
	*tchttp.BaseResponse
	Response *RegisterUserResponseParams `json:"Response"`
}

func (r *RegisterUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomInfo struct {
	// The room name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The room start time (Unix timestamp).
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// The room end time (Unix timestamp).
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// The resolution. Valid values: `1`: SD; `2`: HD; `3`: FHD.
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// The maximum number of mic-on users (excluding the teacher). Value range: 0-16.
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// The room subtype. Valid values: `videodoc`: Document + Video; `video`: Video only; `coteaching`: Dual-teacher.
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// The user ID of the teacher. User IDs are returned by the user registration APIs.
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// Whether to automatically turn the mic on when a user enters the room. Valid values: `0` (default): No; `1`: Yes.
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// Whether to disconnect communication after audio/video permissions are revoked. Valid values: `0` (default): Yes; `1`: No.
	TurnOffMic *uint64 `json:"TurnOffMic,omitempty" name:"TurnOffMic"`

	// Whether to enable the high audio quality mode. Valid values: `0` (default): No; `1`: Yes.
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// Whether to disable auto recording. Valid values: `0` (default): No; `1`: Yes. If this parameter is `0`, recording will start when the class starts and stops when the class ends.
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// The user IDs of the teacher assistants. User IDs are returned by the user registration APIs.
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// The number of RTC users.
	RTCAudienceNumber *uint64 `json:"RTCAudienceNumber,omitempty" name:"RTCAudienceNumber"`

	// The audience type.
	AudienceType *uint64 `json:"AudienceType,omitempty" name:"AudienceType"`

	// The recording layout.
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`

	// The ID of the group to bind. Note: This field may return null, indicating that no valid values can be obtained.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

// Predefined struct for user
type SetAppCustomContentRequestParams struct {
	// Custom content
	CustomContent []*AppCustomContent `json:"CustomContent,omitempty" name:"CustomContent"`

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type SetAppCustomContentRequest struct {
	*tchttp.BaseRequest
	
	// Custom content
	CustomContent []*AppCustomContent `json:"CustomContent,omitempty" name:"CustomContent"`

	// Application ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *SetAppCustomContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAppCustomContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomContent")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAppCustomContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAppCustomContentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetAppCustomContentResponse struct {
	*tchttp.BaseResponse
	Response *SetAppCustomContentResponseParams `json:"Response"`
}

func (r *SetAppCustomContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAppCustomContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWatermarkRequestParams struct {
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The URL of the watermark for the teacher’s video. If you pass in an empty string, the teacher’s video will not have a watermark.
	TeacherUrl *string `json:"TeacherUrl,omitempty" name:"TeacherUrl"`

	// The URL of the watermark for the whiteboard. If you pass in an empty string, the whiteboard video will not have a watermark.
	BoardUrl *string `json:"BoardUrl,omitempty" name:"BoardUrl"`

	// The image displayed when there is no video. If you pass in an empty string, no images will be displayed.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// The width of the whiteboard’s watermark, which is expressed as a percentage of the video width. The value range is 0-100, and the default value is 0.
	BoardW *float64 `json:"BoardW,omitempty" name:"BoardW"`

	// The height of the whiteboard’s watermark, which is expressed as a percentage of the video height. The value range is 0-100, and the default value is 0.
	BoardH *float64 `json:"BoardH,omitempty" name:"BoardH"`

	// The horizontal offset of the whiteboard’s watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle horizontally. Value range: 0-100.
	BoardX *float64 `json:"BoardX,omitempty" name:"BoardX"`

	// The vertical offset of the whiteboard’s watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle vertically. Value range: 0-100.
	BoardY *float64 `json:"BoardY,omitempty" name:"BoardY"`

	// The width of the watermark for the teacher’s video, which is expressed as a percentage of the video width. The value range is 0-100, and the default value is 0.
	TeacherW *float64 `json:"TeacherW,omitempty" name:"TeacherW"`

	// The height of the watermark for the teacher’s video, which is expressed as a percentage of the video height. The value range is 0-100, and the default value is 0.
	TeacherH *float64 `json:"TeacherH,omitempty" name:"TeacherH"`

	// The horizontal offset of the watermark for the teacher’s video, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle horizontally. Value range: 0-100.
	TeacherX *float64 `json:"TeacherX,omitempty" name:"TeacherX"`

	// The vertical offset of the watermark for the teacher’s video, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle vertically. Value range: 0-100.
	TeacherY *float64 `json:"TeacherY,omitempty" name:"TeacherY"`

	// The watermark text. If you pass in an empty string, there will be no text.
	Text *string `json:"Text,omitempty" name:"Text"`

	// The watermark text color.
	TextColor *string `json:"TextColor,omitempty" name:"TextColor"`
}

type SetWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppID assigned by LCIC.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// The URL of the watermark for the teacher’s video. If you pass in an empty string, the teacher’s video will not have a watermark.
	TeacherUrl *string `json:"TeacherUrl,omitempty" name:"TeacherUrl"`

	// The URL of the watermark for the whiteboard. If you pass in an empty string, the whiteboard video will not have a watermark.
	BoardUrl *string `json:"BoardUrl,omitempty" name:"BoardUrl"`

	// The image displayed when there is no video. If you pass in an empty string, no images will be displayed.
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// The width of the whiteboard’s watermark, which is expressed as a percentage of the video width. The value range is 0-100, and the default value is 0.
	BoardW *float64 `json:"BoardW,omitempty" name:"BoardW"`

	// The height of the whiteboard’s watermark, which is expressed as a percentage of the video height. The value range is 0-100, and the default value is 0.
	BoardH *float64 `json:"BoardH,omitempty" name:"BoardH"`

	// The horizontal offset of the whiteboard’s watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle horizontally. Value range: 0-100.
	BoardX *float64 `json:"BoardX,omitempty" name:"BoardX"`

	// The vertical offset of the whiteboard’s watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle vertically. Value range: 0-100.
	BoardY *float64 `json:"BoardY,omitempty" name:"BoardY"`

	// The width of the watermark for the teacher’s video, which is expressed as a percentage of the video width. The value range is 0-100, and the default value is 0.
	TeacherW *float64 `json:"TeacherW,omitempty" name:"TeacherW"`

	// The height of the watermark for the teacher’s video, which is expressed as a percentage of the video height. The value range is 0-100, and the default value is 0.
	TeacherH *float64 `json:"TeacherH,omitempty" name:"TeacherH"`

	// The horizontal offset of the watermark for the teacher’s video, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle horizontally. Value range: 0-100.
	TeacherX *float64 `json:"TeacherX,omitempty" name:"TeacherX"`

	// The vertical offset of the watermark for the teacher’s video, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle vertically. Value range: 0-100.
	TeacherY *float64 `json:"TeacherY,omitempty" name:"TeacherY"`

	// The watermark text. If you pass in an empty string, there will be no text.
	Text *string `json:"Text,omitempty" name:"Text"`

	// The watermark text color.
	TextColor *string `json:"TextColor,omitempty" name:"TextColor"`
}

func (r *SetWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TeacherUrl")
	delete(f, "BoardUrl")
	delete(f, "VideoUrl")
	delete(f, "BoardW")
	delete(f, "BoardH")
	delete(f, "BoardX")
	delete(f, "BoardY")
	delete(f, "TeacherW")
	delete(f, "TeacherH")
	delete(f, "TeacherX")
	delete(f, "TeacherY")
	delete(f, "Text")
	delete(f, "TextColor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWatermarkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *SetWatermarkResponseParams `json:"Response"`
}

func (r *SetWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextMarkConfig struct {
	// The watermark text. Note: This field may return null, indicating that no valid values can be obtained.
	Text *string `json:"Text,omitempty" name:"Text"`

	// The watermark text color. Note: This field may return null, indicating that no valid values can be obtained.
	Color *string `json:"Color,omitempty" name:"Color"`
}

// Predefined struct for user
type UnbindDocumentFromRoomRequestParams struct {
	// Room ID	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Document ID	
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

type UnbindDocumentFromRoomRequest struct {
	*tchttp.BaseRequest
	
	// Room ID	
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// Document ID	
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *UnbindDocumentFromRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDocumentFromRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "DocumentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindDocumentFromRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDocumentFromRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindDocumentFromRoomResponse struct {
	*tchttp.BaseResponse
	Response *UnbindDocumentFromRoomResponseParams `json:"Response"`
}

func (r *UnbindDocumentFromRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDocumentFromRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type WatermarkConfig struct {
	// The URL of the watermark image. Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitempty" name:"Url"`

	// The watermark width, which is expressed as a percentage of the video width. Note: This field may return null, indicating that no valid values can be obtained.
	Width *float64 `json:"Width,omitempty" name:"Width"`

	// The watermark height, which is expressed as a percentage of the video height. Note: This field may return null, indicating that no valid values can be obtained.
	Height *float64 `json:"Height,omitempty" name:"Height"`

	// The horizontal offset of the watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle horizontally. Value range: 0-100. Note: This field may return null, indicating that no valid values can be obtained.
	LocationX *float64 `json:"LocationX,omitempty" name:"LocationX"`

	// The vertical offset of the watermark, which is expressed as a percentage of the video width. For example, 50 indicates that the watermark will appear in the middle vertically. Value range: 0-100. Note: This field may return null, indicating that no valid values can be obtained.
	LocationY *float64 `json:"LocationY,omitempty" name:"LocationY"`
}