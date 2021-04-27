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
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// Unique ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `Uuid`
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Account name.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Account group.
	Groups *string `json:"Groups,omitempty" name:"Groups"`

	// Account type.
	// <li>ORDINARY: ordinary account</li>
	// <li>SUPPER: super admin account</li>
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Account creation time.
	AccountCreateTime *string `json:"AccountCreateTime,omitempty" name:"AccountCreateTime"`

	// Account last login time.
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
}

type AccountStatistics struct {

	// Username.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type AddLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *AddLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("AddLoginWhiteListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMachineTagRequest struct {
	*tchttp.BaseRequest

	// Server ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// Tag ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`

	// Server region
	MRegion *string `json:"MRegion,omitempty" name:"MRegion"`

	// Server type (`CVM` or `BM`)
	MArea *string `json:"MArea,omitempty" name:"MArea"`
}

func (r *AddMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("AddMachineTagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentVul struct {

	// Vulnerability ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Vulnerability name.
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// Vulnerability severity level.
	// <li>HIGH: high</li>
	// <li>MIDDLE: medium</li>
	// <li>LOW: low</li>
	// <li>NOTICE: notice</li>
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// Last scanned time.
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// Vulnerability description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// Vulnerability status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type BruteAttack struct {

	// Event ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Brute force attack event status
	// <li>BRUTEATTACK_FAIL_ACCOUNT: brute force attack event - failure (the account exists)</li>
	// <li>BRUTEATTACK_FAIL_NOACCOUNT: brute force attack event - failure (the account does not exist)</li>
	// <li>BRUTEATTACK_SUCCESS: brute force attack event - success </li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// City ID.
	City *uint64 `json:"City,omitempty" name:"City"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Number of attempts.
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Occurrence time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Whether the server enables CWP Pro.
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// Whether the server is banned.
	BanStatus *string `json:"BanStatus,omitempty" name:"BanStatus"`

	// Server `UUID`
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest

	// Server `Uuid`.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	if len(f) > 0 {
		return errors.New("CloseProVersionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {

	// Unique ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Component version number.
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`

	// Component type.
	// <li>SYSTEM: system component</li>
	// <li>WEB: web component</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// Component name.
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// Component detection update time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type ComponentStatistics struct {

	// Component ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

	// Component name.
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// Component type.
	// <li>WEB: web component</li>
	// <li>SYSTEM: system component</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// Component description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateOpenPortTaskRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateOpenPortTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("CreateOpenPortTaskRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenPortTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOpenPortTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenPortTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessTaskRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *CreateProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("CreateProcessTaskRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// CWP agent `UUID` array.
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids" list`

	// Login region information array.
	Places []*Place `json:"Places,omitempty" name:"Places" list`
}

func (r *CreateUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuids")
	delete(f, "Places")
	if len(f) > 0 {
		return errors.New("CreateUsualLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// Brute force attack event ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("DeleteBruteAttacksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// Whitelist ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("DeleteLoginWhiteListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("DeleteMachineRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest

	// Associated tag ID
	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rid")
	if len(f) > 0 {
		return errors.New("DeleteMachineTagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// Malicious request record ID array. Maximum value: 100 entries.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("DeleteMaliciousRequestsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest

	// Trojan record ID array
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("DeleteMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("DeleteNonlocalLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Added usual login city ID array
	CityIds []*uint64 `json:"CityIds,omitempty" name:"CityIds" list`
}

func (r *DeleteUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "CityIds")
	if len(f) > 0 {
		return errors.New("DeleteUsualLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account username</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAccountStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in account statistics list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Account statistics list.
		AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitempty" name:"AccountStatistics" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// CWP agent `Uuid`. Either `Username` or `Uuid` must be specified. If `Username` is specified, it indicates to query the information list under the specified username.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	// <li>Privilege - String - Required: No - Account name (ORDINARY: ordinary account, SUPPER: super admin account)</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAccountsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in account list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Account data list.
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentVulsRequest struct {
	*tchttp.BaseRequest

	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// Agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeAgentVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAgentVulsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Server vulnerability information
		AgentVuls []*AgentVul `json:"AgentVuls,omitempty" name:"AgentVuls" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmAttributeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeAlarmAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CWP deactivation alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
		Offline *string `json:"Offline,omitempty" name:"Offline"`

		// Trojan discovery alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
		Malware *string `json:"Malware,omitempty" name:"Malware"`

		// Unusual login location discovery alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
		NonlocalLogin *string `json:"NonlocalLogin,omitempty" name:"NonlocalLogin"`

		// Brute force attack success alarm status:
	// <li>OPEN: alarm enabled</li>
	// <li>CLOSE: alarm disabled</li>
		CrackSuccess *string `json:"CrackSuccess,omitempty" name:"CrackSuccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Query status (FAILED: brute force attack failed, SUCCESS: brute force attack succeeded)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeBruteAttacksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of events
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Brute force attack event list
		BruteAttacks []*BruteAttack `json:"BruteAttacks,omitempty" name:"BruteAttacks" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoRequest struct {
	*tchttp.BaseRequest

	// Component ID.
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`
}

func (r *DescribeComponentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ComponentId")
	if len(f) > 0 {
		return errors.New("DescribeComponentInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Component ID.
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// Component name.
		ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

		// Component type.
	// <li>WEB: web component</li>
	// <li>SYSTEM: system component</li>
		ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

		// Component's official website.
		Homepage *string `json:"Homepage,omitempty" name:"Homepage"`

		// Component description.
		Description *string `json:"Description,omitempty" name:"Description"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// ComponentName - String - Required: No - Component name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeComponentStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in component statistics list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Component statistics list data array.
		ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitempty" name:"ComponentStatistics" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`. Either `Uuid` or `ComponentId` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Component ID. Either `Uuid` or `ComponentId` must be specified. If `ComponentId` is specified, it indicates to query the information list under the specified component.
	ComponentId *uint64 `json:"ComponentId,omitempty" name:"ComponentId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>ComponentVersion - String - Required: No - Component version number</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeComponentsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in component list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Component list data.
		Components []*Component `json:"Components,omitempty" name:"Components" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Username - String - Required: No - Account name</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeHistoryAccountsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in account change history list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Account change history data array.
		HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitempty" name:"HistoryAccounts" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeImpactedHostsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Affected server list array
		ImpactedHosts []*ImpactedHost `json:"ImpactedHosts,omitempty" name:"ImpactedHosts" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeLoginWhiteListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Login allowlist array
		LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("DescribeMachineInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Server IP.
		MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

		// Days under protection by CWP
		ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`

		// OS.
		MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

		// Server name.
		MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

		// Status.
	// <li>ONLINE: online</li>
	// <li>OFFLINE: offline</li>
		MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

		// Unique ID of CVM or BM instance.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Public IP of server.
		MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

		// CVM or BM instance `Uuid`.
		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

		// CWP agent `Uuid`.
		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

		// Whether CWP Pro is activated.
	// <li>true: yes</li>
	// <li>false: no</li>
		IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

		// CWP Pro activation time.
		ProVersionOpenDate *string `json:"ProVersionOpenDate,omitempty" name:"ProVersionOpenDate"`

		// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

		// Server region, such as ap-guangzhou or ap-shanghai
		MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

		// Server status.
	// <li>POSTPAY: post-paid, i.e., pay-as-you-go </li>
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// Number of trojans left for free scan.
		FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitempty" name:"FreeMalwaresLeft"`

		// Number of vulnerability left for free scan.
		FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitempty" name:"FreeVulsLeft"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Server region, such as ap-guangzhou or ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: no - Query keywords </li>
	// <li>Status - String - Required: no - CWP client status (valid values: OFFLINE, ONLINE, UNINSTALLED)</li>
	// <li>Version - String - Required: no - Current CWP version (valid values: PRO_VERSION, BASIC_VERSION)</li>
	// Each filter can have only one value but does not support "OR" queries with multiple values
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeMachinesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Server list
		Machines []*Machine `json:"Machines,omitempty" name:"Machines" list`

		// Number of servers
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, TRUSTED: trusted, UN_TRUSTED: untrusted)</li>
	// <li>Domain - String - Required: No - Malicious request domain name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeMaliciousRequestsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Malicious request record array.
		MaliciousRequests []*MaliciousRequest `json:"MaliciousRequests,omitempty" name:"MaliciousRequests" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresRequest struct {
	*tchttp.BaseRequest

	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Trojan status (UN_OPERATED: not processed, SEGREGATED: isolated, TRUSTED: trusted)</li>
	// Each filter supports only one value. Query with multiple values in "OR" relationship is not supported for the time being.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of trojans.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Malware array.
		Malwares []*Malware `json:"Malwares,omitempty" name:"Malwares" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// Agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Keywords - String - Required: No - Query keywords</li>
	// <li>Status - String - Required: No - Login status (NON_LOCAL_LOGIN: unusual login location, NORMAL_LOGIN: intended login)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeNonlocalLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Unusual login location information array.
		NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeOpenPortStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in port statistics list
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Port statistics list
		OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitempty" name:"OpenPortStatistics" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeOpenPortTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("DescribeOpenPortTaskStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task status.
	// <li>COMPLETE: completed (at this point, you can call the `DescribeOpenPorts` API to get the list of real-time processes) </li>
	// <li>AGENT_OFFLINE: CWP agent is offline</li>
	// <li>COLLECTING: getting port</li>
	// <li>FAILED: failed to get processes</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`. Either `Port` or `Uuid` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Open port number. Either `Port` or `Uuid` must be specified. If `Port` is specified, it indicates to query the information list under the specified port.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Port - Uint64 - Required: No - Port number</li>
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeOpenPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeOpenPortsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in port list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Port list.
		OpenPorts []*OpenPort `json:"OpenPorts,omitempty" name:"OpenPorts" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeOverviewStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of online servers.
		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

		// Number of servers activated CWP Pro.
		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

		// Number of trojan files.
		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

		// Number of unusual login locations.
		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

		// Number of successful brute force attacks.
		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

		// Number of vulnerabilities.
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// Security baseline number
		BaseLineNum *uint64 `json:"BaseLineNum,omitempty" name:"BaseLineNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeProVersionInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Fee on yesterday (pay-as-you-go)
		PostPayCost *uint64 `json:"PostPayCost,omitempty" name:"PostPayCost"`

		// Whether CWP Pro is activated for new servers
		IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitempty" name:"IsAutoOpenProVersion"`

		// Number of servers on CWP Pro
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeProcessStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in process statistics list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Process statistics list array.
		ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitempty" name:"ProcessStatistics" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessTaskStatusRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeProcessTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("DescribeProcessTaskStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task status.
	// <li>COMPLETE: completed (at this point, you can call the `DescribeProcesses` API to get the list of real-time processes)</li>
	// <li>AGENT_OFFLINE: CWP agent is offline</li>
	// <li>COLLECTING: getting process</li>
	// <li>FAILED: failed to get processes</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesRequest struct {
	*tchttp.BaseRequest

	// CWP agent `Uuid`. Either `Uuid` or `ProcessName` must be specified. If `Uuid` is specified, it indicates to query the information list under the specified server.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Process name. Either `Uuid` or `ProcessName` must be specified. If `ProcessName` is specified, it indicates to query the information list under the specified process.
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>ProcessName - String - Required: No - Process name</li>
	// <li>MachineIp - String - Required: No - Private IP of server</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProcessesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeProcessesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records in process list.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Process list data array.
		Processes []*Process `json:"Processes,omitempty" name:"Processes" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeSecurityDynamicsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Security event message array.
		SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest

	// Start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// End time.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return errors.New("DescribeSecurityTrendsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Trojan event statistics array.
		Malwares []*SecurityTrend `json:"Malwares,omitempty" name:"Malwares" list`

		// Unusual login location event statistics array.
		NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces" list`

		// Brute force attack event statistics array.
		BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitempty" name:"BruteAttacks" list`

		// Vulnerability statistics array.
		Vuls []*SecurityTrend `json:"Vuls,omitempty" name:"Vuls" list`

		// Baseline statistics array.
		BaseLines []*SecurityTrend `json:"BaseLines,omitempty" name:"BaseLines" list`

		// Statistics array of malicious requests.
		MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitempty" name:"MaliciousRequests" list`

		// Statistics array of high-risk commands.
		HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitempty" name:"HighRiskBashs" list`

		// Statistics array of reverse shells.
		ReverseShells []*SecurityTrend `json:"ReverseShells,omitempty" name:"ReverseShells" list`

		// Statistics array of local privilege escalations.
		PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitempty" name:"PrivilegeEscalations" list`

		// Statistics array of network attacks.
		CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitempty" name:"CyberAttacks" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest

	// Tag ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DescribeTagMachinesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List data
		List []*TagMachine `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// CVM instance type.
	// <li>CVM: CVM</li>
	// <li>BM: CPM</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Server region, such as `ap-guangzhou` and `ap-shanghai`
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	if len(f) > 0 {
		return errors.New("DescribeTagsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List information
		List []*Tag `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// CWP agent `UUID`
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return errors.New("DescribeUsualLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Usual login location array
		UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitempty" name:"UsualLoginPlaces" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoRequest struct {
	*tchttp.BaseRequest

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`
}

func (r *DescribeVulInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulId")
	if len(f) > 0 {
		return errors.New("DescribeVulInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Vulnerability category ID.
		VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

		// Vulnerability name.
		VulName *string `json:"VulName,omitempty" name:"VulName"`

		// Vulnerability level.
		VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

		// Vulnerability type.
		VulType *string `json:"VulType,omitempty" name:"VulType"`

		// Vulnerability description.
		Description *string `json:"Description,omitempty" name:"Description"`

		// Repair plan.
		RepairPlan *string `json:"RepairPlan,omitempty" name:"RepairPlan"`

		// Vulnerability CVE.
		CveId *string `json:"CveId,omitempty" name:"CveId"`

		// Reference link.
		Reference *string `json:"Reference,omitempty" name:"Reference"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanResultRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVulScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeVulScanResultRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulScanResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of vulnerabilities.
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// Number of servers activated CWP Pro
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

		// Number of affected activated CWP Pro.
		ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

		// Total number of servers.
		HostNum *uint64 `json:"HostNum,omitempty" name:"HostNum"`

		// Number of servers on CWP Basic.
		BasicVersionNum *uint64 `json:"BasicVersionNum,omitempty" name:"BasicVersionNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest

	// Vulnerability type.
	// <li>WEB: web application vulnerability</li>
	// <li>SYSTEM: system component vulnerability</li>
	// <li>BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter.
	// <li>Status - String - Required: No - Filter by status (UN_OPERATED: to be processed, FIXED: fixed)
	// 
	// Only one value is allowed for the `Status` filter, and "OR" logic is not supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeVulsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of vulnerabilities.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Vulnerability list array.
		Vuls []*Vul `json:"Vuls,omitempty" name:"Vuls" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeWeeklyReportBruteAttacksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Brute force attack array in weekly CWP Pro report.
		WeeklyReportBruteAttacks []*WeeklyReportBruteAttack `json:"WeeklyReportBruteAttacks,omitempty" name:"WeeklyReportBruteAttacks" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportInfoRequest struct {
	*tchttp.BaseRequest

	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`
}

func (r *DescribeWeeklyReportInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	if len(f) > 0 {
		return errors.New("DescribeWeeklyReportInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account owner name.
		CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

		// Total number of servers.
		MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

		// Number of online CWP agents
		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

		// Number of offline CWP agents.
		OfflineMachineNum *uint64 `json:"OfflineMachineNum,omitempty" name:"OfflineMachineNum"`

		// Number of servers on CWP Pro.
		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

		// Weekly report start time
		BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

		// Weekly report end time
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// Security level
	// <li>HIGH: high</li>
	// <li>MIDDLE: medium</li>
	// <li>LOW: low</li>
		Level *string `json:"Level,omitempty" name:"Level"`

		// Number of trojan records.
		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

		// Number of unusual login locations.
		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

		// Number of successful brute force attacks.
		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

		// Number of vulnerabilities.
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// Download address for exported file.
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresRequest struct {
	*tchttp.BaseRequest

	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeWeeklyReportMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Trojan data in weekly CWP Pro report.
		WeeklyReportMalwares []*WeeklyReportMalware `json:"WeeklyReportMalwares,omitempty" name:"WeeklyReportMalwares" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeWeeklyReportNonlocalLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unusual login location data in weekly CWP Pro report
		WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces,omitempty" name:"WeeklyReportNonlocalLoginPlaces" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsRequest struct {
	*tchttp.BaseRequest

	// Weekly CWP Pro report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeWeeklyReportVulsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportVulsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Vulnerability data array in weekly CWP Pro report.
		WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls,omitempty" name:"WeeklyReportVuls" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportsRequest struct {
	*tchttp.BaseRequest

	// Number of results to be returned. Default value: 10. Maximum value: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWeeklyReportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeWeeklyReportsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWeeklyReportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Weekly CWP Pro report list array.
		WeeklyReports []*WeeklyReport `json:"WeeklyReports,omitempty" name:"WeeklyReports" list`

		// Total number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWeeklyReportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeeklyReportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsRequest struct {
	*tchttp.BaseRequest

	// Tag name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CVM instance ID
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids" list`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("EditTagsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ExportBruteAttacksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download address for exported file.
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ExportMaliciousRequestsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download address for exported file.
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ExportMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download address for exported file.
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ExportNonlocalLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Download address for exported file.
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// Export task ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Filter key name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type HistoryAccount struct {

	// Unique ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Account name.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Account change type.
	// <li>CREATE: creates account</li>
	// <li>MODIFY: modifies account</li>
	// <li>DELETE: deletes account</li>
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// Change time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// Vulnerability ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("IgnoreImpactedHostsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImpactedHost struct {

	// Vulnerability ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Last detection time.
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// Vulnerability status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>SCANING: scanning</li>
	// <li>FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Vulnerability description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Vulnerability category ID.
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// Whether it is the CWP Pro.
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`
}

type LoginWhiteLists struct {

	// Record ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Whitelisted location
	Places []*Place `json:"Places,omitempty" name:"Places" list`

	// Whitelisted users (separated by commas)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whitelisted IPs (separated by commas)
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Whether this rule is applied to all servers under the current account
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// Whitelist creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whitelist modification time
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// Server name
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Server IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type LoginWhiteListsRule struct {

	// Whitelisted location
	Places []*Place `json:"Places,omitempty" name:"Places" list`

	// Whitelisted IPs (separated by commas). This parameter can be an IP range.
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Whitelisted usernames (separated by commas)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether this rule is applied to all servers under the current account
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// Server for which the allowlist takes effect
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// Rule ID, used for rule updating
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Machine struct {

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Server OS.
	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

	// Server status.
	// <li>OFFLINE: offline</li>
	// <li>ONLINE: online</li>
	// <li>MACHINE_STOPPED: shut down</li>
	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

	// CWP agent `Uuid`. If the agent is offline for a long time, a null character will be returned.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// CVM or BM instance `Uuid`.
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// Number of vulnerabilities.
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Whether the server has enabled CWP Pro.
	// <li>true: yes</li>
	// <li>false: no</li>
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// Public IP of server.
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// Server status.
	// <li>POSTPAY: post-paid, i.e., pay-as-you-go </li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Number of trojans.
	MalwareNum *int64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// Tag information
	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag" list`

	// Number of baseline risks.
	BaselineNum *int64 `json:"BaselineNum,omitempty" name:"BaselineNum"`

	// Number of network risks.
	CyberAttackNum *int64 `json:"CyberAttackNum,omitempty" name:"CyberAttackNum"`

	// Risk status.
	// <li>SAFE: safe</li>
	// <li>RISK: at risk</li>
	// <li>UNKNOWN: unknown</li>
	SecurityStatus *string `json:"SecurityStatus,omitempty" name:"SecurityStatus"`

	// Number of intrusions
	InvasionNum *int64 `json:"InvasionNum,omitempty" name:"InvasionNum"`

	// Region information
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
}

type MachineTag struct {

	// Associated tag ID
	Rid *int64 `json:"Rid,omitempty" name:"Rid"`

	// Tag name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Tag ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

type MaliciousRequest struct {

	// Record ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Malicious request domain name.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Number of malicious requests.
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Process name.
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// Record status.
	// <li>UN_OPERATED: to be processed</li>
	// <li>TRUSTED: trusted</li>
	// <li>UN_TRUSTED: untrusted</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Malicious request domain name description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Reference address.
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// Discovery time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Record merge time.
	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`

	// Process MD5
	// Value.
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// Executed command line.
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// Process `PID`.
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
}

type Malware struct {

	// Event ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Current trojan status.
	// <li>UN_OPERATED: not processed</li><li>SEGREGATED: isolated</li><li>TRUSTED: trusted</li>
	// <li>SEPARATING: isolating</li><li>RECOVERING: recovering</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Trojan path.
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// Trojan description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Trojan file creation time.
	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`

	// Trojan file modification time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type MisAlarmNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// Unusual login location event ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *MisAlarmNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("MisAlarmNonlocalLoginPlacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MisAlarmNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MisAlarmNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeRequest struct {
	*tchttp.BaseRequest

	// Alarm item.
	// <li>Offline: CWP is offline</li>
	// <li>Malware: trojan event</li>
	// <li>NonlocalLogin: unusual login location discovered</li>
	// <li>CrackSuccess: brute force attack succeeded</li>
	Attribute *string `json:"Attribute,omitempty" name:"Attribute"`

	// Alarm item attributes.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Attribute")
	delete(f, "Value")
	if len(f) > 0 {
		return errors.New("ModifyAlarmAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest

	// Auto-Activation status.
	// <li>CLOSE: disabled</li>
	// <li>OPEN: enabled</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return errors.New("ModifyAutoOpenProVersionConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// Whitelist rule
	Rules *LoginWhiteListsRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return errors.New("ModifyLoginWhiteListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Auto-renewal flag. Valid values:
	// <li>NOTIFY_AND_AUTO_RENEW: notifies of expiration and auto-renews</li>
	// <li>NOTIFY_AND_MANUAL_RENEW: notifies of expiration but does not auto-renew</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW: does not notify of expiration or auto-renew</li>
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Unique server ID, corresponding to `uuid` for CVM or `instanceId` for BM.
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RenewFlag")
	delete(f, "Quuid")
	if len(f) > 0 {
		return errors.New("ModifyProVersionRenewFlagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NonLocalLoginPlace struct {

	// Event ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Login status
	// <li>NON_LOCAL_LOGIN: unusual login location</li>
	// <li>NORMAL_LOGIN: intended login</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// Username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// City ID.
	City *uint64 `json:"City,omitempty" name:"City"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// Login IP.
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Login time.
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`

	// CWP agent `Uuid`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type OpenPort struct {

	// Unique ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Open port number.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Process name corresponding to port.
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// Process `Pid` corresponding to port.
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// Record creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Record update time.
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type OpenPortStatistics struct {

	// Port number
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Number of servers
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest

	// Server type.
	// <li>CVM: CVM</li>
	// <li>BM: BM</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// Server region
	// Examples: ap-guangzhou, ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// Server `Uuid` array.
	// `InstanceId` for BM or `Uuid` for CVM
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids" list`

	// Event ID.
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("OpenProVersionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Place struct {

	// City ID.
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`

	// Province/State ID.
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// Country/Region ID. Currently, only `1` (Mainland China) is supported.
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`
}

type Process struct {

	// Unique ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Server name.
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Process `Pid`.
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// Process `Ppid`.
	Ppid *uint64 `json:"Ppid,omitempty" name:"Ppid"`

	// Process name.
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// Process username.
	Username *string `json:"Username,omitempty" name:"Username"`

	// OS.
	// <li>WIN32: Windows 32-bit</li>
	// <li>WIN64: Windows 64-bit</li>
	// <li>LINUX32: Linux 32-bit</li>
	// <li>LINUX64: Linux 64-bit</li>
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Process path.
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProcessStatistics struct {

	// Process name.
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// Number of servers.
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest

	// Trojan ID array. Up to 200 IDs can be deleted at a time
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("RecoverMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of IDs of successfully recovered trojans.
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds" list`

		// Array of IDs of trojans failed to be recovered.
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region, such as `ap-guangzhou`, `ap-shanghai` and `ap-beijing`
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name, such as `South China (Guangzhou)`, `East China (Shanghai)`, and `North China (Beijing)`
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region code, such as `gz`, `sh`, and `bj`
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest

	// Vulnerability ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("RescanImpactedHostRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityDynamic struct {

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Security event occurrence time.
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// Security event type.
	// <li>MALWARE: trojan event</li>
	// <li>NON_LOCAL_LOGIN: unusual login location</li>
	// <li>BRUTEATTACK_SUCCESS: brute force attack succeeded</li>
	// <li>VUL: vulnerability</li>
	// <li>BASELINE: security baseline</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// Security event message.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Security event level.
	// <li>RISK: severe</li>
	// <li>HIGH: high</li>
	// <li>NORMAL: medium</li>
	// <li>LOW: low</li>
	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
}

type SecurityTrend struct {

	// Event time.
	Date *string `json:"Date,omitempty" name:"Date"`

	// Number of events.
	EventNum *uint64 `json:"EventNum,omitempty" name:"EventNum"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest

	// Trojan event ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("SeparateMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of IDs of successfully isolated trojans.
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds" list`

		// Array of IDs of trojans failed to be isolated.
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// Tag ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Tag name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Number of servers
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type TagMachine struct {

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Server ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// Server name
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// Private IP of server
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Public IP of server
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// Server region
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// Server region type
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type TrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// Malicious request record ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *TrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("TrustMaliciousRequestRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// Trojan ID array.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("TrustMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestRequest struct {
	*tchttp.BaseRequest

	// Trusted record ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UntrustMaliciousRequestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("UntrustMaliciousRequestRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMaliciousRequestResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMaliciousRequestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMaliciousRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// Trojan event ID array. Up to 200 IDs can be processed at a time.
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids" list`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return errors.New("UntrustMalwaresRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {

	// ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CWP agent `UUID`.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Country/Region ID.
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`

	// Province/State ID.
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// City ID.
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
}

type Vul struct {

	// Vulnerability category ID
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// Vulnerability name
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// Vulnerability severity level:
	// HIGH: high
	// MIDDLE: medium
	// LOW: low
	// NOTICE: notice
	VulLevel *string `json:"VulLevel,omitempty" name:"VulLevel"`

	// Last scanned time
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// Number of affected servers
	ImpactedHostNum *uint64 `json:"ImpactedHostNum,omitempty" name:"ImpactedHostNum"`

	// Vulnerability status
	// * UN_OPERATED: to be processed
	// * FIXED: fixed
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`
}

type WeeklyReport struct {

	// Weekly report start time.
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// Weekly report end time.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type WeeklyReportBruteAttack struct {

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Hacked username.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Number of attempts.
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// Attack time.
	AttackTime *string `json:"AttackTime,omitempty" name:"AttackTime"`
}

type WeeklyReportMalware struct {

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Trojan file path.
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// Trojan file MD5 value.
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// Trojan discovery time.
	FindTime *string `json:"FindTime,omitempty" name:"FindTime"`

	// Current trojan status.
	// <li>UN_OPERATED: not processed</li>
	// <li>SEGREGATED: isolated</li>
	// <li>TRUSTED: trusted</li>
	// <li>SEPARATING: isolating</li>
	// <li>RECOVERING: recovering</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type WeeklyReportNonlocalLoginPlace struct {

	// Server IP.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Username.
	Username *string `json:"Username,omitempty" name:"Username"`

	// Source IP.
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// Country/Region ID.
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// Province/State ID.
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// City ID.
	City *uint64 `json:"City,omitempty" name:"City"`

	// Login time.
	LoginTime *string `json:"LoginTime,omitempty" name:"LoginTime"`
}

type WeeklyReportVul struct {

	// Private IP of server.
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// Vulnerability name.
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// Vulnerability type.
	// <li> WEB: web vulnerability</li>
	// <li> SYSTEM: system component vulnerability</li>
	// <li> BASELINE: security baseline</li>
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// Vulnerability description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Vulnerability status.
	// <li> UN_OPERATED: to be processed</li>
	// <li> SCANING: scanning</li>
	// <li> FIXED: fixed</li>
	VulStatus *string `json:"VulStatus,omitempty" name:"VulStatus"`

	// Last scanned time.
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`
}
