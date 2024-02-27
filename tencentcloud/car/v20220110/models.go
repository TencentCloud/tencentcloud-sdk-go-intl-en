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

package v20220110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyConcurrentRequestParams struct {
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP of user’s application client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// The project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Application version ID. If the application of the current version is requested, you do not need to fill in this field. If the application of other versions is requested, you need to specify the version through this field.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application ID, which is used only by the multi-application project to specify applications. For a single-application project, this parameter is ignored, and the application bound to the project will be used.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

type ApplyConcurrentRequest struct {
	*tchttp.BaseRequest
	
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP of user’s application client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// The project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Application version ID. If the application of the current version is requested, you do not need to fill in this field. If the application of other versions is requested, you need to specify the version through this field.
	ApplicationVersionId *string `json:"ApplicationVersionId,omitnil,omitempty" name:"ApplicationVersionId"`

	// Application ID, which is used only by the multi-application project to specify applications. For a single-application project, this parameter is ignored, and the application bound to the project will be used.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`
}

func (r *ApplyConcurrentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ProjectId")
	delete(f, "ApplicationVersionId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConcurrentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyConcurrentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyConcurrentResponse struct {
	*tchttp.BaseResponse
	Response *ApplyConcurrentResponseParams `json:"Response"`
}

func (r *ApplyConcurrentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionRequestParams struct {
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP of user’s application client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// The client-side session data, which is obtained from the SDK. If `RunMode` is `RunWithoutClient`, this parameter can be null.
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// The on-cloud running mode.
	// `RunWithoutClient`: Keep the application running on the cloud even when there are no client connections.
	// Empty string (default): Keep the application running on the cloud only when there are client connections.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Application startup parameter.
	// If the user requests a multi-application project or a prelaunch-disabled single-application project, this parameter takes effect.
	//  
	// If the user requests a prelaunch-enabled single-application project, this parameter is invalid.
	// 
	// Note: When this parameter takes effect, the `ApplicationParameters` parameter will be appended to the end of the application startup parameter. The application startup parameter is set in the application or project configuration in the console.
	// For example, for a prelaunch-disabled single-application project, if its application startup parameter `bar` is `0` and the `ApplicationParameters` parameter `foo` is `1`, the actual application startup parameters will be `bar=0 foo=1`.
	ApplicationParameters *string `json:"ApplicationParameters,omitnil,omitempty" name:"ApplicationParameters"`

	// The user ID of the host in **multi-person interaction** scenarios, which is required.
	// If the current user is the host, `HostUserId` must be the same as their `UserId`; otherwise, `HostUserId` should be the host's `UserId`.
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// The role in **multi-person interaction** scenarios. Valid values:
	// `Player`: A user who can operate an application by using a keyboard and mouse
	// `Viewer`: A user who can only watch the video in the room but cannot operate the application
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type CreateSessionRequest struct {
	*tchttp.BaseRequest
	
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Public IP of user’s application client, which is used for nearby scheduling.
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// The client-side session data, which is obtained from the SDK. If `RunMode` is `RunWithoutClient`, this parameter can be null.
	ClientSession *string `json:"ClientSession,omitnil,omitempty" name:"ClientSession"`

	// The on-cloud running mode.
	// `RunWithoutClient`: Keep the application running on the cloud even when there are no client connections.
	// Empty string (default): Keep the application running on the cloud only when there are client connections.
	RunMode *string `json:"RunMode,omitnil,omitempty" name:"RunMode"`

	// Application startup parameter.
	// If the user requests a multi-application project or a prelaunch-disabled single-application project, this parameter takes effect.
	//  
	// If the user requests a prelaunch-enabled single-application project, this parameter is invalid.
	// 
	// Note: When this parameter takes effect, the `ApplicationParameters` parameter will be appended to the end of the application startup parameter. The application startup parameter is set in the application or project configuration in the console.
	// For example, for a prelaunch-disabled single-application project, if its application startup parameter `bar` is `0` and the `ApplicationParameters` parameter `foo` is `1`, the actual application startup parameters will be `bar=0 foo=1`.
	ApplicationParameters *string `json:"ApplicationParameters,omitnil,omitempty" name:"ApplicationParameters"`

	// The user ID of the host in **multi-person interaction** scenarios, which is required.
	// If the current user is the host, `HostUserId` must be the same as their `UserId`; otherwise, `HostUserId` should be the host's `UserId`.
	HostUserId *string `json:"HostUserId,omitnil,omitempty" name:"HostUserId"`

	// The role in **multi-person interaction** scenarios. Valid values:
	// `Player`: A user who can operate an application by using a keyboard and mouse
	// `Viewer`: A user who can only watch the video in the room but cannot operate the application
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ClientSession")
	delete(f, "RunMode")
	delete(f, "ApplicationParameters")
	delete(f, "HostUserId")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionResponseParams struct {
	// The server-side session data, which is returned to the SDK.
	ServerSession *string `json:"ServerSession,omitnil,omitempty" name:"ServerSession"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionResponseParams `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentCountRequestParams struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeConcurrentCountRequest struct {
	*tchttp.BaseRequest
	
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeConcurrentCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrentCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentCountResponseParams struct {
	// Total Concurrency Count
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The number of concurrent executions, including those in pre-launch, connected, waiting for reconnection, to be cleaned up or recovered, and all non-idle concurrent executions. Therefore, refreshing projects or disconnecting user connections with concurrency packages will affect this value.
	Running *uint64 `json:"Running,omitnil,omitempty" name:"Running"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrentCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrentCountResponseParams `json:"Response"`
}

func (r *DescribeConcurrentCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionRequestParams struct {
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DestroySessionRequest struct {
	*tchttp.BaseRequest
	
	// The user’s unique ID. Tencent Cloud does not parse the ID. You need to manage your own user IDs. Based on your needs, you can either define unique IDs for users or use timestamps to generate random IDs. Make sure the same ID is used when a user reconnects to your application.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DestroySessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroySessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroySessionResponse struct {
	*tchttp.BaseResponse
	Response *DestroySessionResponseParams `json:"Response"`
}

func (r *DestroySessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamRequestParams struct {
	// Unique user ID, which is customized by you and is not understood by CAR. It will be used as the `StreamId` for pushing streams. For example, if the bound push domain is **abc.livepush.myqcloud.com**, the push address will be **rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx**.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Push parameter, which is a custom parameter carried during stream pushing.
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

type StartPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not understood by CAR. It will be used as the `StreamId` for pushing streams. For example, if the bound push domain is **abc.livepush.myqcloud.com**, the push address will be **rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx**.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Push parameter, which is a custom parameter carried during stream pushing.
	PublishStreamArgs *string `json:"PublishStreamArgs,omitnil,omitempty" name:"PublishStreamArgs"`
}

func (r *StartPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishStreamArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamResponseParams `json:"Response"`
}

func (r *StartPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamWithURLRequestParams struct {
	// Unique user ID, which is customized by you and is not understood by CAR.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Push address. Only RTMP is supported for push currently.
	PublishStreamURL *string `json:"PublishStreamURL,omitnil,omitempty" name:"PublishStreamURL"`
}

type StartPublishStreamWithURLRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not understood by CAR.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Push address. Only RTMP is supported for push currently.
	PublishStreamURL *string `json:"PublishStreamURL,omitnil,omitempty" name:"PublishStreamURL"`
}

func (r *StartPublishStreamWithURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamWithURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PublishStreamURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamWithURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamWithURLResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishStreamWithURLResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamWithURLResponseParams `json:"Response"`
}

func (r *StartPublishStreamWithURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamWithURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamRequestParams struct {
	// Unique user ID, which is customized by you and is not understood by CAR. It can also be randomly generated using the timestamp and should be kept unchanged during user reconnection.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type StopPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// Unique user ID, which is customized by you and is not understood by CAR. It can also be randomly generated using the timestamp and should be kept unchanged during user reconnection.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *StopPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishStreamResponseParams `json:"Response"`
}

func (r *StopPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}