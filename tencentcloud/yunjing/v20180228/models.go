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

package v20180228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// Unique ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `Uuid`
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Account name.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Account group.
	Groups *string `json:"Groups,omitnil" name:"Groups"`

	// Account type.
	// <li>ORDINARY: ordinary account</li>
	// <li>SUPPER: super admin account</li>
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// Account creation time.
	AccountCreateTime *string `json:"AccountCreateTime,omitnil" name:"AccountCreateTime"`

	// Account last login time.
	LastLoginTime *string `json:"LastLoginTime,omitnil" name:"LastLoginTime"`
}

type AccountStatistics struct {
	// Username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

// Predefined struct for user
type AddLoginWhiteListRequestParams struct {
	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitnil" name:"Rules"`
}

type AddLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitnil" name:"Rules"`
}

func (r *AddLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLoginWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *AddLoginWhiteListResponseParams `json:"Response"`
}

func (r *AddLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineTagRequestParams struct {
	// Server ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// Tag ID
	TagId *uint64 `json:"TagId,omitnil" name:"TagId"`

	// Server region
	MRegion *string `json:"MRegion,omitnil" name:"MRegion"`

	// Server type (`CVM` or `BM`)
	MArea *string `json:"MArea,omitnil" name:"MArea"`
}

type AddMachineTagRequest struct {
	*tchttp.BaseRequest
	
	// Server ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// Tag ID
	TagId *uint64 `json:"TagId,omitnil" name:"TagId"`

	// Server region
	MRegion *string `json:"MRegion,omitnil" name:"MRegion"`

	// Server type (`CVM` or `BM`)
	MArea *string `json:"MArea,omitnil" name:"MArea"`
}

func (r *AddMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	delete(f, "TagId")
	delete(f, "MRegion")
	delete(f, "MArea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *AddMachineTagResponseParams `json:"Response"`
}

func (r *AddMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentVul struct {
	// Vulnerability ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Vulnerability name.
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// Vulnerability severity level.
	// <li>HIGH: high</li>
	// <li>MIDDLE: medium</li>
	// <li>LOW: low</li>
	// <li>NOTICE: notice</li>
	VulLevel *string `json:"VulLevel,omitnil" name:"VulLevel"`

	// Last scanned time.
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// Vulnerability description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Vulnerability status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitnil" name:"VulStatus"`
}

type BruteAttack struct {
	// Event ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Brute force attack event status
	// <li>BRUTEATTACK_FAIL_ACCOUNT: brute force attack event - failure (the account exists)</li>
	// <li>BRUTEATTACK_FAIL_NOACCOUNT: brute force attack event - failure (the account does not exist)</li>
	// <li>BRUTEATTACK_SUCCESS: brute force attack event - success </li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// Username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// City ID.
	City *uint64 `json:"City,omitnil" name:"City"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitnil" name:"Province"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Number of attempts.
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// Occurrence time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Whether the server enables CWP Pro.
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// Whether the server is banned.
	BanStatus *string `json:"BanStatus,omitnil" name:"BanStatus"`

	// Server `UUID`
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

// Predefined struct for user
type CloseProVersionRequestParams struct {
	// Server `Uuid`.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest
	
	// Server `Uuid`.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseProVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse
	Response *CloseProVersionResponseParams `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// Unique ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Component version number.
	ComponentVersion *string `json:"ComponentVersion,omitnil" name:"ComponentVersion"`

	// Component type.
	// <li>SYSTEM: system component</li>
	// <li>WEB: web component</li>
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// Component name.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// Component detection update time.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type ComponentStatistics struct {
	// Component ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`

	// Component name.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// Component type.
	// <li>WEB: web component</li>
	// <li>SYSTEM: system component</li>
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// Component description.
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type CreateOpenPortTaskRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type CreateOpenPortTaskRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *CreateOpenPortTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenPortTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenPortTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateOpenPortTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenPortTaskResponseParams `json:"Response"`
}

func (r *CreateOpenPortTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessTaskRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type CreateProcessTaskRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *CreateProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProcessTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProcessTaskResponseParams `json:"Response"`
}

func (r *CreateProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsualLoginPlacesRequestParams struct {
	// CWP agent `UUID` array.
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// Login region information array.
	Places []*Place `json:"Places,omitnil" name:"Places"`
}

type CreateUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `UUID` array.
	Uuids []*string `json:"Uuids,omitnil" name:"Uuids"`

	// Login region information array.
	Places []*Place `json:"Places,omitnil" name:"Places"`
}

func (r *CreateUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuids")
	delete(f, "Places")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsualLoginPlacesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *CreateUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksRequestParams struct {
	// Brute force attack event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// Brute force attack event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBruteAttacksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBruteAttacksResponseParams `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListRequestParams struct {
	// Whitelist ID
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Whitelist ID
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoginWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoginWhiteListResponseParams `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineResponseParams `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagRequestParams struct {
	// Associated tag ID
	Rid *uint64 `json:"Rid,omitnil" name:"Rid"`
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest
	
	// Associated tag ID
	Rid *uint64 `json:"Rid,omitnil" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineTagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineTagResponseParams `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsRequestParams struct {
	// Malicious request record ID array. Maximum value: 100 entries.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// Malicious request record ID array. Maximum value: 100 entries.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMaliciousRequestsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMaliciousRequestsResponseParams `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresRequestParams struct {
	// Trojan record ID array
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Trojan record ID array
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMalwaresResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMalwaresResponseParams `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesRequestParams struct {
	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNonlocalLoginPlacesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsualLoginPlacesRequestParams struct {
	// CWP agent `Uuid`
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Added usual login city ID array
	CityIds []*uint64 `json:"CityIds,omitnil" name:"CityIds"`
}

type DeleteUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Added usual login city ID array
	CityIds []*uint64 `json:"CityIds,omitnil" name:"CityIds"`
}

func (r *DeleteUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "CityIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsualLoginPlacesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *DeleteUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account username</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account username</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountStatisticsResponseParams struct {
	// Total number of records in account statistics list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Account statistics list.
	AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitnil" name:"AccountStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Username` is specified, it indicates to query the information list under the specified username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	// <li>Privilege - String - Required: No - Account name (ORDINARY: ordinary account, SUPPER: super admin account)</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Username` is specified, it indicates to query the information list under the specified username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	// <li>Privilege - String - Required: No - Account name (ORDINARY: ordinary account, SUPPER: super admin account)</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Username")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Total number of records in account list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Account data list.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVulsRequestParams struct {
	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAgentVulsRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAgentVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulType")
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVulsResponseParams struct {
	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Server vulnerability information
	AgentVuls []*AgentVul `json:"AgentVuls,omitnil" name:"AgentVuls"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAgentVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentVulsResponseParams `json:"Response"`
}

func (r *DescribeAgentVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmAttributeRequestParams struct {

}

type DescribeAlarmAttributeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmAttributeResponseParams struct {
	// CWP deactivation alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
	Offline *string `json:"Offline,omitnil" name:"Offline"`

	// Trojan discovery alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
	Malware *string `json:"Malware,omitnil" name:"Malware"`

	// Unusual login location discovery alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
	NonlocalLogin *string `json:"NonlocalLogin,omitnil" name:"NonlocalLogin"`

	// Brute force attack success alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
	CrackSuccess *string `json:"CrackSuccess,omitnil" name:"CrackSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmAttributeResponseParams `json:"Response"`
}

func (r *DescribeAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttacksRequestParams struct {
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Query status (FAILED: brute force attack failed, SUCCESS: brute force attack succeeded)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Query status (FAILED: brute force attack failed, SUCCESS: brute force attack succeeded)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBruteAttacksResponseParams struct {
	// Number of events
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Brute force attack event list
	BruteAttacks []*BruteAttack `json:"BruteAttacks,omitnil" name:"BruteAttacks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBruteAttacksResponseParams `json:"Response"`
}

func (r *DescribeBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentInfoRequestParams struct {
	// Component ID.
	ComponentId *uint64 `json:"ComponentId,omitnil" name:"ComponentId"`
}

type DescribeComponentInfoRequest struct {
	*tchttp.BaseRequest
	
	// Component ID.
	ComponentId *uint64 `json:"ComponentId,omitnil" name:"ComponentId"`
}

func (r *DescribeComponentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ComponentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentInfoResponseParams struct {
	// Component ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Component name.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// Component type.
	// <li>WEB: web component</li>
	// <li>SYSTEM: system component</li>
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// Component's official website.
	Homepage *string `json:"Homepage,omitnil" name:"Homepage"`

	// Component description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeComponentInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentInfoResponseParams `json:"Response"`
}

func (r *DescribeComponentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// ComponentName - String - Required: No - Component name
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// ComponentName - String - Required: No - Component name
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentStatisticsResponseParams struct {
	// Total number of records in component statistics list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Component statistics list data array.
	ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitnil" name:"ComponentStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentStatisticsResponseParams `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentsRequestParams struct {
	// CWP agent `Uuid`. Either `Uuid` or `ComponentId` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Component ID. Either `Uuid` or `ComponentId` must be specified. If `ComponentId` is specified, it indicates to query the information list under the specified component.
	ComponentId *uint64 `json:"ComponentId,omitnil" name:"ComponentId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ComponentVersion - String - Required: No - Component version number</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeComponentsRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`. Either `Uuid` or `ComponentId` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Component ID. Either `Uuid` or `ComponentId` must be specified. If `ComponentId` is specified, it indicates to query the information list under the specified component.
	ComponentId *uint64 `json:"ComponentId,omitnil" name:"ComponentId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ComponentVersion - String - Required: No - Component version number</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "ComponentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComponentsResponseParams struct {
	// Total number of records in component list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Component list data.
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeComponentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComponentsResponseParams `json:"Response"`
}

func (r *DescribeComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryAccountsResponseParams struct {
	// Total number of records in account change history list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Account change history data array.
	HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitnil" name:"HistoryAccounts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistoryAccountsResponseParams `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpactedHostsRequestParams struct {
	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeImpactedHostsRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpactedHostsResponseParams struct {
	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Affected server list array
	ImpactedHosts []*ImpactedHost `json:"ImpactedHosts,omitnil" name:"ImpactedHosts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImpactedHostsResponseParams `json:"Response"`
}

func (r *DescribeImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginWhiteListResponseParams struct {
	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Login allowlist array
	LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitnil" name:"LoginWhiteLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginWhiteListResponseParams `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineInfoResponseParams struct {
	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Days under protection by CWP
	ProtectDays *uint64 `json:"ProtectDays,omitnil" name:"ProtectDays"`

	// OS.
	MachineOs *string `json:"MachineOs,omitnil" name:"MachineOs"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Status.
	// <li>ONLINE: online</li>
	// <li>OFFLINE: offline</li>
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// Unique ID of CVM or BM instance.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Public IP of server.
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// CVM or BM instance `Uuid`.
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Whether CWP Pro is activated.
	// <li>true: yes</li>
	// <li>false: no</li>
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// CWP Pro activation time.
	ProVersionOpenDate *string `json:"ProVersionOpenDate,omitnil" name:"ProVersionOpenDate"`

	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region, such as ap-guangzhou or ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Server status.
	// <li>POSTPAY: post-paid, i.e., pay-as-you-go </li>
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Number of trojans left for free scan.
	FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitnil" name:"FreeMalwaresLeft"`

	// Number of vulnerability left for free scan.
	FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitnil" name:"FreeVulsLeft"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineInfoResponseParams `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesRequestParams struct {
	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region, such as ap-guangzhou or ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: no - Query keywords </li>
	// <li>Status - String - Required: no - CWP client status (valid values: OFFLINE, ONLINE, UNINSTALLED)</li>
	// <li>Version - String - Required: no - Current CWP version (valid values: PRO_VERSION, BASIC_VERSION)</li>
	// Each filter can have only one value but does not support "OR" queries with multiple values
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest
	
	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region, such as ap-guangzhou or ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: no - Query keywords </li>
	// <li>Status - String - Required: no - CWP client status (valid values: OFFLINE, ONLINE, UNINSTALLED)</li>
	// <li>Version - String - Required: no - Current CWP version (valid values: PRO_VERSION, BASIC_VERSION)</li>
	// Each filter can have only one value but does not support "OR" queries with multiple values
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesResponseParams struct {
	// Server list
	Machines []*Machine `json:"Machines,omitnil" name:"Machines"`

	// Number of servers
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachinesResponseParams `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, TRUSTED: trusted, UN_TRUSTED: untrusted)</li>
	// <li>Domain - String - Required: No - Malicious request domain name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, TRUSTED: trusted, UN_TRUSTED: untrusted)</li>
	// <li>Domain - String - Required: No - Malicious request domain name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaliciousRequestsResponseParams struct {
	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Malicious request record array.
	MaliciousRequests []*MaliciousRequest `json:"MaliciousRequests,omitnil" name:"MaliciousRequests"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaliciousRequestsResponseParams `json:"Response"`
}

func (r *DescribeMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwaresRequestParams struct {
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Trojan status (UN_OPERATED: not processed, SEGREGATED: isolated, TRUSTED: trusted)</li>
	// Each filter supports only one value. Query with multiple values in "OR" relationship is not supported for the time being.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Trojan status (UN_OPERATED: not processed, SEGREGATED: isolated, TRUSTED: trusted)</li>
	// Each filter supports only one value. Query with multiple values in "OR" relationship is not supported for the time being.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMalwaresResponseParams struct {
	// Total number of trojans.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Malware array.
	Malwares []*Malware `json:"Malwares,omitnil" name:"Malwares"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMalwaresResponseParams `json:"Response"`
}

func (r *DescribeMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNonlocalLoginPlacesRequestParams struct {
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Login status (NON_LOCAL_LOGIN: unusual login location, NORMAL_LOGIN: intended login)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Login status (NON_LOCAL_LOGIN: unusual login location, NORMAL_LOGIN: intended login)</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNonlocalLoginPlacesResponseParams struct {
	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Unusual login location information array.
	NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces,omitnil" name:"NonLocalLoginPlaces"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortStatisticsResponseParams struct {
	// Total number of records in port statistics list
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Port statistics list
	OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitnil" name:"OpenPortStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortTaskStatusRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeOpenPortTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeOpenPortTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortTaskStatusResponseParams struct {
	// Task status.
	// <li>COMPLETE: completed (at this point, you can call the `DescribeOpenPorts` API to get the list of real-time processes) </li>
	// <li>AGENT_OFFLINE: CWP agent is offline</li>
	// <li>COLLECTING: getting port</li>
	// <li>FAILED: failed to get processes</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOpenPortTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeOpenPortTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortsRequestParams struct {
	// CWP agent `Uuid`. Either `Port` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Open port number. Either `Port` or `Uuid` must be specified. If `Port` is specified, it indicates to query the information list under the specified port.
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeOpenPortsRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`. Either `Port` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Open port number. Either `Port` or `Uuid` must be specified. If `Port` is specified, it indicates to query the information list under the specified port.
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeOpenPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Port")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPortsResponseParams struct {
	// Total number of records in port list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Port list.
	OpenPorts []*OpenPort `json:"OpenPorts,omitnil" name:"OpenPorts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOpenPortsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPortsResponseParams `json:"Response"`
}

func (r *DescribeOpenPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsRequestParams struct {

}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewStatisticsResponseParams struct {
	// Number of online servers.
	OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitnil" name:"OnlineMachineNum"`

	// Number of servers activated CWP Pro.
	ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitnil" name:"ProVersionMachineNum"`

	// Number of trojan files.
	MalwareNum *uint64 `json:"MalwareNum,omitnil" name:"MalwareNum"`

	// Number of unusual login locations.
	NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitnil" name:"NonlocalLoginNum"`

	// Number of successful brute force attacks.
	BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitnil" name:"BruteAttackSuccessNum"`

	// Number of vulnerabilities.
	VulNum *uint64 `json:"VulNum,omitnil" name:"VulNum"`

	// Security baseline number
	BaseLineNum *uint64 `json:"BaseLineNum,omitnil" name:"BaseLineNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewStatisticsResponseParams `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoRequestParams struct {

}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProVersionInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoResponseParams struct {
	// Fee on yesterday (pay-as-you-go)
	PostPayCost *uint64 `json:"PostPayCost,omitnil" name:"PostPayCost"`

	// Whether CWP Pro is activated for new servers
	IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitnil" name:"IsAutoOpenProVersion"`

	// Number of servers on CWP Pro
	ProVersionNum *uint64 `json:"ProVersionNum,omitnil" name:"ProVersionNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProVersionInfoResponseParams `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessStatisticsResponseParams struct {
	// Total number of records in process statistics list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Process statistics list array.
	ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitnil" name:"ProcessStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessTaskStatusRequestParams struct {
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeProcessTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeProcessTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessTaskStatusResponseParams struct {
	// Task status.
	// <li>COMPLETE: completed (at this point, you can call the `DescribeProcesses` API to get the list of real-time processes)</li>
	// <li>AGENT_OFFLINE: CWP agent is offline</li>
	// <li>COLLECTING: getting process</li>
	// <li>FAILED: failed to get processes</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProcessTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeProcessTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessesRequestParams struct {
	// CWP agent `Uuid`. Either `Uuid` or `ProcessName` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Process name. Either `Uuid` or `ProcessName` must be specified. If `ProcessName` is specified, it indicates to query the information list under the specified process.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeProcessesRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `Uuid`. Either `Uuid` or `ProcessName` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Process name. Either `Uuid` or `ProcessName` must be specified. If `ProcessName` is specified, it indicates to query the information list under the specified process.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "ProcessName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProcessesResponseParams struct {
	// Total number of records in process list.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Process list data array.
	Processes []*Process `json:"Processes,omitnil" name:"Processes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProcessesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProcessesResponseParams `json:"Response"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityDynamicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityDynamicsResponseParams struct {
	// Security event message array.
	SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitnil" name:"SecurityDynamics"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityDynamicsResponseParams `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsRequestParams struct {
	// Start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// End time.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest
	
	// Start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// End time.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityTrendsResponseParams struct {
	// Trojan event statistics array.
	Malwares []*SecurityTrend `json:"Malwares,omitnil" name:"Malwares"`

	// Unusual login location event statistics array.
	NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitnil" name:"NonLocalLoginPlaces"`

	// Brute force attack event statistics array.
	BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitnil" name:"BruteAttacks"`

	// Vulnerability statistics array.
	Vuls []*SecurityTrend `json:"Vuls,omitnil" name:"Vuls"`

	// Baseline statistics array.
	BaseLines []*SecurityTrend `json:"BaseLines,omitnil" name:"BaseLines"`

	// Statistics array of malicious requests.
	MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitnil" name:"MaliciousRequests"`

	// Statistics array of high-risk commands.
	HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitnil" name:"HighRiskBashs"`

	// Statistics array of reverse shells.
	ReverseShells []*SecurityTrend `json:"ReverseShells,omitnil" name:"ReverseShells"`

	// Statistics array of local privilege escalations.
	PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitnil" name:"PrivilegeEscalations"`

	// Statistics array of network attacks.
	CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitnil" name:"CyberAttacks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityTrendsResponseParams `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesRequestParams struct {
	// Tag ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest
	
	// Tag ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagMachinesResponseParams struct {
	// List data
	List []*TagMachine `json:"List,omitnil" name:"List"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagMachinesResponseParams `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// CVM instance type.
	// <li>CVM: CVM</li>
	// <li>BM: CPM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region, such as `ap-guangzhou` and `ap-shanghai`
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// CVM instance type.
	// <li>CVM: CVM</li>
	// <li>BM: CPM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region, such as `ap-guangzhou` and `ap-shanghai`
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`
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
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// List information
	List []*Tag `json:"List,omitnil" name:"List"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeUsualLoginPlacesRequestParams struct {
	// CWP agent `UUID`
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// CWP agent `UUID`
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsualLoginPlacesResponseParams struct {
	// Usual login location array
	UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitnil" name:"UsualLoginPlaces"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsualLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoRequestParams struct {
	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`
}

type DescribeVulInfoRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`
}

func (r *DescribeVulInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulInfoResponseParams struct {
	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Vulnerability name.
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// Vulnerability level.
	VulLevel *string `json:"VulLevel,omitnil" name:"VulLevel"`

	// Vulnerability type.
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Vulnerability description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Repair plan.
	RepairPlan *string `json:"RepairPlan,omitnil" name:"RepairPlan"`

	// Vulnerability CVE.
	CveId *string `json:"CveId,omitnil" name:"CveId"`

	// Reference link.
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulInfoResponseParams `json:"Response"`
}

func (r *DescribeVulInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanResultRequestParams struct {

}

type DescribeVulScanResultRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulScanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanResultResponseParams struct {
	// Number of vulnerabilities.
	VulNum *uint64 `json:"VulNum,omitnil" name:"VulNum"`

	// Number of servers activated CWP Pro
	ProVersionNum *uint64 `json:"ProVersionNum,omitnil" name:"ProVersionNum"`

	// Number of affected activated CWP Pro.
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitnil" name:"ImpactedHostNum"`

	// Total number of servers.
	HostNum *uint64 `json:"HostNum,omitnil" name:"HostNum"`

	// Number of servers on CWP Basic.
	BasicVersionNum *uint64 `json:"BasicVersionNum,omitnil" name:"BasicVersionNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulScanResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulScanResultResponseParams `json:"Response"`
}

func (r *DescribeVulScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsRequestParams struct {
	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	// 
	// Only one value is allowed for the `Status` filter, and "OR" logic is not supported.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	// 
	// Only one value is allowed for the `Status` filter, and "OR" logic is not supported.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsResponseParams struct {
	// Number of vulnerabilities.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Vulnerability list array.
	Vuls []*Vul `json:"Vuls,omitnil" name:"Vuls"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsResponseParams `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportBruteAttacksRequestParams struct {
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWeeklyReportBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWeeklyReportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportBruteAttacksResponseParams struct {
	// Brute force attack array in weekly CWP Pro report.
	WeeklyReportBruteAttacks []*WeeklyReportBruteAttack `json:"WeeklyReportBruteAttacks,omitnil" name:"WeeklyReportBruteAttacks"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportBruteAttacksResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportInfoRequestParams struct {
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`
}

type DescribeWeeklyReportInfoRequest struct {
	*tchttp.BaseRequest
	
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`
}

func (r *DescribeWeeklyReportInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportInfoResponseParams struct {
	// Account owner name.
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// Total number of servers.
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`

	// Number of online CWP agents
	OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitnil" name:"OnlineMachineNum"`

	// Number of offline CWP agents.
	OfflineMachineNum *uint64 `json:"OfflineMachineNum,omitnil" name:"OfflineMachineNum"`

	// Number of servers on CWP Pro.
	ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitnil" name:"ProVersionMachineNum"`

	// Weekly report start time
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Weekly report end time
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// Security level
	// <li>HIGH: high</li>
	// <li>MIDDLE: medium</li>
	// <li>LOW: low</li>
	Level *string `json:"Level,omitnil" name:"Level"`

	// Number of trojan records.
	MalwareNum *uint64 `json:"MalwareNum,omitnil" name:"MalwareNum"`

	// Number of unusual login locations.
	NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitnil" name:"NonlocalLoginNum"`

	// Number of successful brute force attacks.
	BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitnil" name:"BruteAttackSuccessNum"`

	// Number of vulnerabilities.
	VulNum *uint64 `json:"VulNum,omitnil" name:"VulNum"`

	// Download address for exported file.
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportInfoResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportMalwaresRequestParams struct {
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWeeklyReportMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWeeklyReportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportMalwaresResponseParams struct {
	// Trojan data in weekly CWP Pro report.
	WeeklyReportMalwares []*WeeklyReportMalware `json:"WeeklyReportMalwares,omitnil" name:"WeeklyReportMalwares"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportMalwaresResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportNonlocalLoginPlacesRequestParams struct {
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportNonlocalLoginPlacesResponseParams struct {
	// Unusual login location data in weekly CWP Pro report
	WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces,omitnil" name:"WeeklyReportNonlocalLoginPlaces"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportVulsRequestParams struct {
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWeeklyReportVulsRequest struct {
	*tchttp.BaseRequest
	
	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWeeklyReportVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportVulsResponseParams struct {
	// Vulnerability data array in weekly CWP Pro report.
	WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls,omitnil" name:"WeeklyReportVuls"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportVulsResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportsRequestParams struct {
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWeeklyReportsRequest struct {
	*tchttp.BaseRequest
	
	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWeeklyReportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeeklyReportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeeklyReportsResponseParams struct {
	// Weekly CWP Pro report list array.
	WeeklyReports []*WeeklyReport `json:"WeeklyReports,omitnil" name:"WeeklyReports"`

	// Total number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWeeklyReportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeeklyReportsResponseParams `json:"Response"`
}

func (r *DescribeWeeklyReportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsRequestParams struct {
	// Tag name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tag ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CVM instance ID
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

type EditTagsRequest struct {
	*tchttp.BaseRequest
	
	// Tag name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tag ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CVM instance ID
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *EditTagsResponseParams `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksRequestParams struct {

}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBruteAttacksResponseParams struct {
	// Download address for exported file.
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *ExportBruteAttacksResponseParams `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsRequestParams struct {

}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMaliciousRequestsResponseParams struct {
	// Download address for exported file.
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *ExportMaliciousRequestsResponseParams `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresRequestParams struct {

}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportMalwaresResponseParams struct {
	// Download address for exported file.
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *ExportMalwaresResponseParams `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesRequestParams struct {

}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportNonlocalLoginPlacesResponseParams struct {
	// Download address for exported file.
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// Export task ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *ExportNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Filter key name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// One or more filter values.
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type HistoryAccount struct {
	// Unique ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Account name.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Account change type.
	// <li>CREATE: creates account</li>
	// <li>MODIFY: modifies account</li>
	// <li>DELETE: deletes account</li>
	ModifyType *string `json:"ModifyType,omitnil" name:"ModifyType"`

	// Change time.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

// Predefined struct for user
type IgnoreImpactedHostsRequestParams struct {
	// Vulnerability ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreImpactedHostsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreImpactedHostsResponseParams `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImpactedHost struct {
	// Vulnerability ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Last detection time.
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// Vulnerability status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>SCANING: scanning</li>
	// <li>FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitnil" name:"VulStatus"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Vulnerability description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Whether it is the CWP Pro.
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`
}

type LoginWhiteLists struct {
	// Record ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent ID
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Whitelisted location
	Places []*Place `json:"Places,omitnil" name:"Places"`

	// Whitelisted users (separated by commas)
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Whitelisted IPs (separated by commas)
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Whether this rule is applied to all servers under the current account
	IsGlobal *bool `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// Whitelist creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whitelist modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Server name
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Server IP
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type LoginWhiteListsRule struct {
	// Whitelisted location
	Places []*Place `json:"Places,omitnil" name:"Places"`

	// Whitelisted IPs (separated by commas). This parameter can be an IP range.
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Whitelisted usernames (separated by commas)
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Whether this rule is applied to all servers under the current account
	IsGlobal *bool `json:"IsGlobal,omitnil" name:"IsGlobal"`

	// Server for which the allowlist takes effect
	HostIp *string `json:"HostIp,omitnil" name:"HostIp"`

	// Rule ID, used for rule updating
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type Machine struct {
	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Server OS.
	MachineOs *string `json:"MachineOs,omitnil" name:"MachineOs"`

	// Server status.
	// <li>OFFLINE: offline</li>
	// <li>ONLINE: online</li>
	// <li>MACHINE_STOPPED: shut down</li>
	MachineStatus *string `json:"MachineStatus,omitnil" name:"MachineStatus"`

	// CWP agent `Uuid`. If the agent is offline for a long time, a null character will be returned.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// CVM or BM instance `Uuid`.
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// Number of vulnerabilities.
	VulNum *int64 `json:"VulNum,omitnil" name:"VulNum"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Whether the server has enabled CWP Pro.
	// <li>true: yes</li>
	// <li>false: no</li>
	IsProVersion *bool `json:"IsProVersion,omitnil" name:"IsProVersion"`

	// Public IP of server.
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// Server status.
	// <li>POSTPAY: post-paid, i.e., pay-as-you-go </li>
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Number of trojans.
	MalwareNum *int64 `json:"MalwareNum,omitnil" name:"MalwareNum"`

	// Tag information
	Tag []*MachineTag `json:"Tag,omitnil" name:"Tag"`

	// Number of baseline risks.
	BaselineNum *int64 `json:"BaselineNum,omitnil" name:"BaselineNum"`

	// Number of network risks.
	CyberAttackNum *int64 `json:"CyberAttackNum,omitnil" name:"CyberAttackNum"`

	// Risk status.
	// <li>SAFE: safe</li>
	// <li>RISK: at risk</li>
	// <li>UNKNOWN: unknown</li>
	SecurityStatus *string `json:"SecurityStatus,omitnil" name:"SecurityStatus"`

	// Number of intrusions
	InvasionNum *int64 `json:"InvasionNum,omitnil" name:"InvasionNum"`

	// Region information
	RegionInfo *RegionInfo `json:"RegionInfo,omitnil" name:"RegionInfo"`
}

type MachineTag struct {
	// Associated tag ID
	Rid *int64 `json:"Rid,omitnil" name:"Rid"`

	// Tag name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Tag ID
	TagId *uint64 `json:"TagId,omitnil" name:"TagId"`
}

type MaliciousRequest struct {
	// Record ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Malicious request domain name.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Number of malicious requests.
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// Process name.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Record status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>TRUSTED: trusted</li>
	// <li>UN_TRUSTED: untrusted</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// Malicious request domain name description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Reference address.
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// Discovery time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Record merge time.
	MergeTime *string `json:"MergeTime,omitnil" name:"MergeTime"`

	// Process MD5
	// Value.
	ProcessMd5 *string `json:"ProcessMd5,omitnil" name:"ProcessMd5"`

	// Executed command line.
	CmdLine *string `json:"CmdLine,omitnil" name:"CmdLine"`

	// Process `PID`.
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`
}

type Malware struct {
	// Event ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Current trojan status.
	// <li>UN_OPERATED: not processed</li><li>SEGREGATED: isolated</li><li>TRUSTED: trusted</li>
	// <li>SEPARATING: isolating</li><li>RECOVERING: recovering</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// Trojan path.
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// Trojan description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Trojan file creation time.
	FileCreateTime *string `json:"FileCreateTime,omitnil" name:"FileCreateTime"`

	// Trojan file modification time.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

// Predefined struct for user
type MisAlarmNonlocalLoginPlacesRequestParams struct {
	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type MisAlarmNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
	
	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *MisAlarmNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MisAlarmNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MisAlarmNonlocalLoginPlacesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *MisAlarmNonlocalLoginPlacesResponseParams `json:"Response"`
}

func (r *MisAlarmNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmAttributeRequestParams struct {
	// Alarm item.
	// <li>Offline: CWP is offline</li>
	// <li>Malware: trojan event</li>
	// <li>NonlocalLogin: unusual login location discovered</li>
	// <li>CrackSuccess: brute force attack succeeded</li>
	Attribute *string `json:"Attribute,omitnil" name:"Attribute"`

	// Alarm item attributes.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ModifyAlarmAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Alarm item.
	// <li>Offline: CWP is offline</li>
	// <li>Malware: trojan event</li>
	// <li>NonlocalLogin: unusual login location discovered</li>
	// <li>CrackSuccess: brute force attack succeeded</li>
	Attribute *string `json:"Attribute,omitnil" name:"Attribute"`

	// Alarm item attributes.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Value *string `json:"Value,omitnil" name:"Value"`
}

func (r *ModifyAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Attribute")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmAttributeResponseParams `json:"Response"`
}

func (r *ModifyAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigRequestParams struct {
	// Auto-Activation status.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest
	
	// Auto-Activation status.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoOpenProVersionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoOpenProVersionConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoOpenProVersionConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginWhiteListRequestParams struct {
	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitnil" name:"Rules"`
}

type ModifyLoginWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitnil" name:"Rules"`
}

func (r *ModifyLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginWhiteListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoginWhiteListResponseParams `json:"Response"`
}

func (r *ModifyLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProVersionRenewFlagRequestParams struct {
	// Auto-renewal flag. Valid values:
	// <li>NOTIFY_AND_AUTO_RENEW: notifies of expiration and auto-renews</li>
	// <li>NOTIFY_AND_MANUAL_RENEW: notifies of expiration but does not auto-renew</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: does not notify of expiration or auto-renew</li>
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Unique server ID, corresponding to `uuid` for CVM or `instanceId` for BM.
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Auto-renewal flag. Valid values:
	// <li>NOTIFY_AND_AUTO_RENEW: notifies of expiration and auto-renews</li>
	// <li>NOTIFY_AND_MANUAL_RENEW: notifies of expiration but does not auto-renew</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: does not notify of expiration or auto-renew</li>
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Unique server ID, corresponding to `uuid` for CVM or `instanceId` for BM.
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RenewFlag")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProVersionRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProVersionRenewFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProVersionRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NonLocalLoginPlace struct {
	// Event ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Login status
	// <li>NON_LOCAL_LOGIN: unusual login location</li>
	// <li>NORMAL_LOGIN: intended login</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// Username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// City ID.
	City *uint64 `json:"City,omitnil" name:"City"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitnil" name:"Province"`

	// Login IP.
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Login time.
	LoginTime *string `json:"LoginTime,omitnil" name:"LoginTime"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type OpenPort struct {
	// Unique ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Open port number.
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Process name corresponding to port.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Process `Pid` corresponding to port.
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`

	// Record creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Record update time.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type OpenPortStatistics struct {
	// Port number
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Number of servers
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

// Predefined struct for user
type OpenProVersionRequestParams struct {
	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region
	// Examples: ap-guangzhou, ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Server `Uuid` array.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`

	// Event ID.
	ActivityId *uint64 `json:"ActivityId,omitnil" name:"ActivityId"`
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest
	
	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// Server region
	// Examples: ap-guangzhou, ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Server `Uuid` array.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuids []*string `json:"Quuids,omitnil" name:"Quuids"`

	// Event ID.
	ActivityId *uint64 `json:"ActivityId,omitnil" name:"ActivityId"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Quuids")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenProVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse
	Response *OpenProVersionResponseParams `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Place struct {
	// City ID.
	CityId *uint64 `json:"CityId,omitnil" name:"CityId"`

	// Province/State ID.
	ProvinceId *uint64 `json:"ProvinceId,omitnil" name:"ProvinceId"`

	// Country/Region ID. Currently, only `1` (Mainland China) is supported.
	CountryId *uint64 `json:"CountryId,omitnil" name:"CountryId"`
}

type Process struct {
	// Unique ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Process `Pid`.
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`

	// Process `Ppid`.
	Ppid *uint64 `json:"Ppid,omitnil" name:"Ppid"`

	// Process name.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Process username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// OS.
	// <li>WIN32: Windows 32-bit</li>
	// <li>WIN64: Windows 64-bit</li>
	// <li>LINUX32: Linux 32-bit</li>
	// <li>LINUX64: Linux 64-bit</li>
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// Process path.
	FullPath *string `json:"FullPath,omitnil" name:"FullPath"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type ProcessStatistics struct {
	// Process name.
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitnil" name:"MachineNum"`
}

// Predefined struct for user
type RecoverMalwaresRequestParams struct {
	// Trojan ID array. Up to 200 IDs can be deleted at a time
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Trojan ID array. Up to 200 IDs can be deleted at a time
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverMalwaresResponseParams struct {
	// Array of IDs of successfully recovered trojans.
	SuccessIds []*uint64 `json:"SuccessIds,omitnil" name:"SuccessIds"`

	// Array of IDs of trojans failed to be recovered.
	FailedIds []*uint64 `json:"FailedIds,omitnil" name:"FailedIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *RecoverMalwaresResponseParams `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region, such as `ap-guangzhou`, `ap-shanghai` and `ap-beijing`
	Region *string `json:"Region,omitnil" name:"Region"`

	// Region name, such as `South China (Guangzhou)`, `East China (Shanghai)`, and `North China (Beijing)`
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// Region code, such as `gz`, `sh`, and `bj`
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`
}

// Predefined struct for user
type RescanImpactedHostRequestParams struct {
	// Vulnerability ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest
	
	// Vulnerability ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RescanImpactedHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RescanImpactedHostResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse
	Response *RescanImpactedHostResponseParams `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityDynamic struct {
	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Security event occurrence time.
	EventTime *string `json:"EventTime,omitnil" name:"EventTime"`

	// Security event type.
	// <li>MALWARE: trojan event</li>
	// <li>NON_LOCAL_LOGIN: unusual login location</li>
	// <li>BRUTEATTACK_SUCCESS: brute force attack succeeded</li>
	// <li>VUL: vulnerability</li>
	// <li>BASELINE: security baseline</li>
	EventType *string `json:"EventType,omitnil" name:"EventType"`

	// Security event message.
	Message *string `json:"Message,omitnil" name:"Message"`

	// Security event level.
	// <li>RISK: severe</li>
	// <li>HIGH: high</li>
	// <li>NORMAL: medium</li>
	// <li>LOW: low</li>
	SecurityLevel *string `json:"SecurityLevel,omitnil" name:"SecurityLevel"`
}

type SecurityTrend struct {
	// Event time.
	Date *string `json:"Date,omitnil" name:"Date"`

	// Number of events.
	EventNum *uint64 `json:"EventNum,omitnil" name:"EventNum"`
}

// Predefined struct for user
type SeparateMalwaresRequestParams struct {
	// Trojan event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Trojan event ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SeparateMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SeparateMalwaresResponseParams struct {
	// Array of IDs of successfully isolated trojans.
	SuccessIds []*uint64 `json:"SuccessIds,omitnil" name:"SuccessIds"`

	// Array of IDs of trojans failed to be isolated.
	FailedIds []*uint64 `json:"FailedIds,omitnil" name:"FailedIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *SeparateMalwaresResponseParams `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Tag name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Number of servers
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type TagMachine struct {
	// ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Server ID
	Quuid *string `json:"Quuid,omitnil" name:"Quuid"`

	// Server name
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// Private IP of server
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Public IP of server
	MachineWanIp *string `json:"MachineWanIp,omitnil" name:"MachineWanIp"`

	// Server region
	MachineRegion *string `json:"MachineRegion,omitnil" name:"MachineRegion"`

	// Server region type
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`
}

// Predefined struct for user
type TrustMaliciousRequestRequestParams struct {
	// Malicious request record ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type TrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest
	
	// Malicious request record ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *TrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMaliciousRequestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMaliciousRequestResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *TrustMaliciousRequestResponseParams `json:"Response"`
}

func (r *TrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMalwaresRequestParams struct {
	// Trojan ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Trojan ID array.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrustMalwaresResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *TrustMalwaresResponseParams `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMaliciousRequestRequestParams struct {
	// Trusted record ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type UntrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest
	
	// Trusted record ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *UntrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMaliciousRequestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMaliciousRequestResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UntrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *UntrustMaliciousRequestResponseParams `json:"Response"`
}

func (r *UntrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresRequestParams struct {
	// Trojan event ID array. Up to 200 IDs can be processed at a time.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest
	
	// Trojan event ID array. Up to 200 IDs can be processed at a time.
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntrustMalwaresResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *UntrustMalwaresResponseParams `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {
	// ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Country/Region ID.
	CountryId *uint64 `json:"CountryId,omitnil" name:"CountryId"`

	// Province/State ID.
	ProvinceId *uint64 `json:"ProvinceId,omitnil" name:"ProvinceId"`

	// City ID.
	CityId *uint64 `json:"CityId,omitnil" name:"CityId"`
}

type Vul struct {
	// Vulnerability category ID
	VulId *uint64 `json:"VulId,omitnil" name:"VulId"`

	// Vulnerability name
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// Vulnerability severity level:
	// HIGH: high
	// MIDDLE: medium
	// LOW: low
	// NOTICE: notice
	VulLevel *string `json:"VulLevel,omitnil" name:"VulLevel"`

	// Last scanned time
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// Number of affected servers
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitnil" name:"ImpactedHostNum"`

	// Vulnerability status
	// * UN_OPERATED: to be processed
	// * FIXED: fixed
	VulStatus *string `json:"VulStatus,omitnil" name:"VulStatus"`
}

type WeeklyReport struct {
	// Weekly report start time.
	BeginDate *string `json:"BeginDate,omitnil" name:"BeginDate"`

	// Weekly report end time.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`
}

type WeeklyReportBruteAttack struct {
	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Hacked username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Number of attempts.
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// Attack time.
	AttackTime *string `json:"AttackTime,omitnil" name:"AttackTime"`
}

type WeeklyReportMalware struct {
	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Trojan file path.
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// Trojan file MD5 value.
	Md5 *string `json:"Md5,omitnil" name:"Md5"`

	// Trojan discovery time.
	FindTime *string `json:"FindTime,omitnil" name:"FindTime"`

	// Current trojan status.
	// <li>UN_OPERATED: not processed</li>
	// <li>SEGREGATED: isolated</li>
	// <li>TRUSTED: trusted</li>
	// <li>SEPARATING: isolating</li>
	// <li>RECOVERING: recovering</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

type WeeklyReportNonlocalLoginPlace struct {
	// Server IP.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Username.
	Username *string `json:"Username,omitnil" name:"Username"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitnil" name:"Province"`

	// City ID.
	City *uint64 `json:"City,omitnil" name:"City"`

	// Login time.
	LoginTime *string `json:"LoginTime,omitnil" name:"LoginTime"`
}

type WeeklyReportVul struct {
	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitnil" name:"MachineIp"`

	// Vulnerability name.
	VulName *string `json:"VulName,omitnil" name:"VulName"`

	// Vulnerability type.
	// <li> WEB: web vulnerability</li>
	// <li> SYSTEM: system component vulnerability</li>
	// <li> BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitnil" name:"VulType"`

	// Vulnerability description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Vulnerability status.
	// <li> UN_OPERATED: to be processed</li>
	// <li> SCANING: scanning</li>
	// <li> FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitnil" name:"VulStatus"`

	// Last scanned time.
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`
}