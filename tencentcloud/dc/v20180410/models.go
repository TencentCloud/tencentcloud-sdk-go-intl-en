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

package v20180410

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptDirectConnectTunnelRequestParams struct {
	// The connection owner accepts an application for sharing the dedicated tunnel
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// The connection owner accepts an application for sharing the dedicated tunnel
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

func (r *AcceptDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptDirectConnectTunnelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AcceptDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *AcceptDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *AcceptDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessPoint struct {
	// Access point name.
	AccessPointName *string `json:"AccessPointName,omitnil" name:"AccessPointName"`

	// Unique access point ID.
	AccessPointId *string `json:"AccessPointId,omitnil" name:"AccessPointId"`

	// Access point status. Valid values: available, unavailable.
	State *string `json:"State,omitnil" name:"State"`

	// Access point location.
	Location *string `json:"Location,omitnil" name:"Location"`

	// List of ISPs supported by access point.
	LineOperator []*string `json:"LineOperator,omitnil" name:"LineOperator"`

	// ID of the region that manages the access point.
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// Available port type at the access point. Valid values: 1000BASE-T: gigabit electrical port; 1000BASE-LX: 10 km gigabit single-mode optical port; 1000BASE-ZX: 80 km gigabit single-mode optical port; 10GBASE-LR: 10 km 10-gigabit single-mode optical port; 10GBASE-ZR: 80 km 10-gigabit single-mode optical port; 10GBASE-LH: 40 km 10-gigabit single-mode optical port; 100GBASE-LR4: 10 km 100-gigabit single-mode optical portfiber optic port.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AvailablePortType []*string `json:"AvailablePortType,omitnil" name:"AvailablePortType"`

	// Latitude and longitude of the access point
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Coordinate *Coordinate `json:"Coordinate,omitnil" name:"Coordinate"`

	// City where the access point is located
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	City *string `json:"City,omitnil" name:"City"`

	// Access point region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Area *string `json:"Area,omitnil" name:"Area"`

	// Access point type. Valid values: `VXLAN`, `QCPL`, and `QCAR`.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AccessPointType *string `json:"AccessPointType,omitnil" name:"AccessPointType"`
}

// Predefined struct for user
type ApplyInternetAddressRequestParams struct {
	// Mask length of a CIDR block
	MaskLen *int64 `json:"MaskLen,omitnil" name:"MaskLen"`

	// Address type. Valid values: 0: BGP
	// 1: China Telecom
	// 2: China Mobile
	// 3: China Unicom
	AddrType *int64 `json:"AddrType,omitnil" name:"AddrType"`

	// Address protocol. Valid values: 0: IPv4
	// 1: IPv6
	AddrProto *int64 `json:"AddrProto,omitnil" name:"AddrProto"`
}

type ApplyInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// Mask length of a CIDR block
	MaskLen *int64 `json:"MaskLen,omitnil" name:"MaskLen"`

	// Address type. Valid values: 0: BGP
	// 1: China Telecom
	// 2: China Mobile
	// 3: China Unicom
	AddrType *int64 `json:"AddrType,omitnil" name:"AddrType"`

	// Address protocol. Valid values: 0: IPv4
	// 1: IPv6
	AddrProto *int64 `json:"AddrProto,omitnil" name:"AddrProto"`
}

func (r *ApplyInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaskLen")
	delete(f, "AddrType")
	delete(f, "AddrProto")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyInternetAddressResponseParams struct {
	// ID of the internet tunnel’s public IP address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *ApplyInternetAddressResponseParams `json:"Response"`
}

func (r *ApplyInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BFDInfo struct {
	// Number of health checks
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitnil" name:"ProbeFailedTimes"`

	// Health check interval
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`
}

type BgpPeer struct {
	// User-side BGP Asn.
	Asn *int64 `json:"Asn,omitnil" name:"Asn"`

	// User-side BGP key.
	AuthKey *string `json:"AuthKey,omitnil" name:"AuthKey"`
}

type Coordinate struct {
	// Latitude
	Lat *float64 `json:"Lat,omitnil" name:"Lat"`

	// Longitude
	Lng *float64 `json:"Lng,omitnil" name:"Lng"`
}

// Predefined struct for user
type CreateDirectConnectRequestParams struct {
	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitnil" name:"DirectConnectName"`

	// Access point of connection.
	// You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
	AccessPointId *string `json:"AccessPointId,omitnil" name:"AccessPointId"`

	// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
	LineOperator *string `json:"LineOperator,omitnil" name:"LineOperator"`

	// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
	PortType *string `json:"PortType,omitnil" name:"PortType"`

	// Circuit code of a connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitnil" name:"CircuitCode"`

	// Local IDC location.
	Location *string `json:"Location,omitnil" name:"Location"`

	// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// ID of redundant connection.
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil" name:"RedundantDirectConnectId"`

	// VLAN for connection debugging, which is enabled and automatically assigned by default.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// Tencent-side IP address for connection debugging, which is automatically assigned by default.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address for connection debugging, which is automatically assigned by default.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitnil" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitnil" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil" name:"FaultReportContactNumber"`

	// Whether the connection applicant has signed the service agreement. Default value: true.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitnil" name:"DirectConnectName"`

	// Access point of connection.
	// You can call `DescribeAccessPoints` to get the region ID. The selected access point must exist and be available.
	AccessPointId *string `json:"AccessPointId,omitnil" name:"AccessPointId"`

	// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
	LineOperator *string `json:"LineOperator,omitnil" name:"LineOperator"`

	// Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
	PortType *string `json:"PortType,omitnil" name:"PortType"`

	// Circuit code of a connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitnil" name:"CircuitCode"`

	// Local IDC location.
	Location *string `json:"Location,omitnil" name:"Location"`

	// Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// ID of redundant connection.
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil" name:"RedundantDirectConnectId"`

	// VLAN for connection debugging, which is enabled and automatically assigned by default.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// Tencent-side IP address for connection debugging, which is automatically assigned by default.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address for connection debugging, which is automatically assigned by default.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitnil" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitnil" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil" name:"FaultReportContactNumber"`

	// Whether the connection applicant has signed the service agreement. Default value: true.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`
}

func (r *CreateDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectName")
	delete(f, "AccessPointId")
	delete(f, "LineOperator")
	delete(f, "PortType")
	delete(f, "CircuitCode")
	delete(f, "Location")
	delete(f, "Bandwidth")
	delete(f, "RedundantDirectConnectId")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "CustomerName")
	delete(f, "CustomerContactMail")
	delete(f, "CustomerContactNumber")
	delete(f, "FaultReportContactPerson")
	delete(f, "FaultReportContactNumber")
	delete(f, "SignLaw")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectResponseParams struct {
	// Connection ID.
	DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitnil" name:"DirectConnectIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectResponseParams `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectTunnelRequestParams struct {
	// Direct Connect ID, such as `dc-kd7d06of`.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil" name:"DirectConnectTunnelName"`

	// Connection owner, who is the current customer by default.
	// The developer account ID should be entered for shared connections.
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil" name:"DirectConnectOwnerAccount"`

	// Network type. Valid values: VPC, BMVPC, CCN. Default value: VPC.
	// VPC: Virtual Private Cloud.
	// BMVPC: BM VPC.
	// CCN: Cloud Connect Network.
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// Network region.
	NetworkRegion *string `json:"NetworkRegion,omitnil" name:"NetworkRegion"`

	// Unified VPC ID or BMVPC ID.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Direct connect gateway ID, such as `dcg-d545ddf`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// Direct Connect bandwidth in Mbps.
	// Default value: connection bandwidth value.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// BGP: BGP routing.
	// STATIC: Static routing.
	// Default value: BGP routing.
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// BgpPeer, which is BGP information on the user side and includes Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil" name:"BgpPeer"`

	// Static routing, i.e., IP range of the user's IDC.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil" name:"RouteFilterPrefixes"`

	// VLAN. Value range: 0-3,000.
	// 0: sub-interface not enabled.
	// Default value: Non-zero.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// TencentAddress: Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// CustomerAddress: User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// TencentBackupAddress, i.e., Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil" name:"TencentBackupAddress"`

	// Cloud Attached Connection Service ID
	CloudAttachId *string `json:"CloudAttachId,omitnil" name:"CloudAttachId"`

	// Whether to enable BFD
	BfdEnable *int64 `json:"BfdEnable,omitnil" name:"BfdEnable"`

	// Whether to enable NQA
	NqaEnable *int64 `json:"NqaEnable,omitnil" name:"NqaEnable"`

	// BFD configuration information
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil" name:"BfdInfo"`

	// NQA configuration information
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil" name:"NqaInfo"`
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// Direct Connect ID, such as `dc-kd7d06of`.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil" name:"DirectConnectTunnelName"`

	// Connection owner, who is the current customer by default.
	// The developer account ID should be entered for shared connections.
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil" name:"DirectConnectOwnerAccount"`

	// Network type. Valid values: VPC, BMVPC, CCN. Default value: VPC.
	// VPC: Virtual Private Cloud.
	// BMVPC: BM VPC.
	// CCN: Cloud Connect Network.
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// Network region.
	NetworkRegion *string `json:"NetworkRegion,omitnil" name:"NetworkRegion"`

	// Unified VPC ID or BMVPC ID.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Direct connect gateway ID, such as `dcg-d545ddf`.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// Direct Connect bandwidth in Mbps.
	// Default value: connection bandwidth value.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// BGP: BGP routing.
	// STATIC: Static routing.
	// Default value: BGP routing.
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// BgpPeer, which is BGP information on the user side and includes Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil" name:"BgpPeer"`

	// Static routing, i.e., IP range of the user's IDC.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil" name:"RouteFilterPrefixes"`

	// VLAN. Value range: 0-3,000.
	// 0: sub-interface not enabled.
	// Default value: Non-zero.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// TencentAddress: Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// CustomerAddress: User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// TencentBackupAddress, i.e., Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil" name:"TencentBackupAddress"`

	// Cloud Attached Connection Service ID
	CloudAttachId *string `json:"CloudAttachId,omitnil" name:"CloudAttachId"`

	// Whether to enable BFD
	BfdEnable *int64 `json:"BfdEnable,omitnil" name:"BfdEnable"`

	// Whether to enable NQA
	NqaEnable *int64 `json:"NqaEnable,omitnil" name:"NqaEnable"`

	// BFD configuration information
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil" name:"BfdInfo"`

	// NQA configuration information
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil" name:"NqaInfo"`
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	delete(f, "DirectConnectTunnelName")
	delete(f, "DirectConnectOwnerAccount")
	delete(f, "NetworkType")
	delete(f, "NetworkRegion")
	delete(f, "VpcId")
	delete(f, "DirectConnectGatewayId")
	delete(f, "Bandwidth")
	delete(f, "RouteType")
	delete(f, "BgpPeer")
	delete(f, "RouteFilterPrefixes")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "TencentBackupAddress")
	delete(f, "CloudAttachId")
	delete(f, "BfdEnable")
	delete(f, "NqaEnable")
	delete(f, "BfdInfo")
	delete(f, "NqaInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectTunnelResponseParams struct {
	// Dedicated tunnel ID.
	DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitnil" name:"DirectConnectTunnelIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectRequestParams struct {
	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`
}

func (r *DeleteDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectTunnelRequestParams struct {
	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

func (r *DeleteDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectTunnelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessPointsRequestParams struct {
	// Access point region, which can be queried through `DescribeRegions`.
	// 
	// You can call `DescribeRegions` to get the region ID.
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest
	
	// Access point region, which can be queried through `DescribeRegions`.
	// 
	// You can call `DescribeRegions` to get the region ID.
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessPointsResponseParams struct {
	// Access point information.
	AccessPointSet []*AccessPoint `json:"AccessPointSet,omitnil" name:"AccessPointSet"`

	// Number of eligible access points.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessPointsResponseParams `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelsRequestParams struct {
	// Filter conditions:
	// This parameter does not support specifying `DirectConnectTunnelIds` and `Filters` at the same time.
	// <li> direct-connect-tunnel-name: Dedicated tunnel name.</li>
	// <li> direct-connect-tunnel-id: Dedicated tunnel instance ID, such as `dcx-abcdefgh`.</li>
	// <li>direct-connect-id: Connection instance ID, such as `dc-abcdefgh`.</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Array of dedicated tunnel IDs.
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitnil" name:"DirectConnectTunnelIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions:
	// This parameter does not support specifying `DirectConnectTunnelIds` and `Filters` at the same time.
	// <li> direct-connect-tunnel-name: Dedicated tunnel name.</li>
	// <li> direct-connect-tunnel-id: Dedicated tunnel instance ID, such as `dcx-abcdefgh`.</li>
	// <li>direct-connect-id: Connection instance ID, such as `dc-abcdefgh`.</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Array of dedicated tunnel IDs.
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitnil" name:"DirectConnectTunnelIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "DirectConnectTunnelIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectTunnelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelsResponseParams struct {
	// List of dedicated tunnels.
	DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitnil" name:"DirectConnectTunnelSet"`

	// Number of eligible dedicated tunnels.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectTunnelsResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectsRequestParams struct {
	// Filter conditions:
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Array of connection IDs.
	DirectConnectIds []*string `json:"DirectConnectIds,omitnil" name:"DirectConnectIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions:
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Array of connection IDs.
	DirectConnectIds []*string `json:"DirectConnectIds,omitnil" name:"DirectConnectIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "DirectConnectIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectsResponseParams struct {
	// List of connections.
	DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitnil" name:"DirectConnectSet"`

	// Number of eligible connection lists.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Whether all connections under the account have the service agreement signed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AllSignLaw *bool `json:"AllSignLaw,omitnil" name:"AllSignLaw"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectsResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressQuotaRequestParams struct {

}

type DescribeInternetAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressQuotaResponseParams struct {
	// Minimum prefix length allowed for a public IPv6 address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ipv6PrefixLen *int64 `json:"Ipv6PrefixLen,omitnil" name:"Ipv6PrefixLen"`

	// Quota of BGP IPv4 addresses
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ipv4BgpQuota *int64 `json:"Ipv4BgpQuota,omitnil" name:"Ipv4BgpQuota"`

	// Quota of non-BGP IPv4 addresses
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ipv4OtherQuota *int64 `json:"Ipv4OtherQuota,omitnil" name:"Ipv4OtherQuota"`

	// Used number of BGP IPv4 addresses
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ipv4BgpNum *int64 `json:"Ipv4BgpNum,omitnil" name:"Ipv4BgpNum"`

	// Used number of non-BGP IPv4 addresses
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ipv4OtherNum *int64 `json:"Ipv4OtherNum,omitnil" name:"Ipv4OtherNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInternetAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressQuotaResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressRequestParams struct {
	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter conditions:
	// <li>AddrType, address type. Valid values: 0: BGP; 1: China Telecom; 2: China Mobile; 3: China Unicom</li>
	// <li>AddrProto, address protocol. Valid values: 0: IPv4; 1: IPv6</li>
	// <li>Status, address status. Valid values: 0: in use; 1: disabled; 2: returned</li>
	// <li>Subnet, public IP address array</li>
	// <InstanceIds>Public IP address ID array</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter conditions:
	// <li>AddrType, address type. Valid values: 0: BGP; 1: China Telecom; 2: China Mobile; 3: China Unicom</li>
	// <li>AddrProto, address protocol. Valid values: 0: IPv4; 1: IPv6</li>
	// <li>Status, address status. Valid values: 0: in use; 1: disabled; 2: returned</li>
	// <li>Subnet, public IP address array</li>
	// <InstanceIds>Public IP address ID array</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressResponseParams struct {
	// Number of public IP addresses for internet tunnels
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of the public IP addresses for internet tunnels
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Subnets []*InternetAddressDetail `json:"Subnets,omitnil" name:"Subnets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressStatisticsRequestParams struct {

}

type DescribeInternetAddressStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetAddressStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressStatisticsResponseParams struct {
	// Number of public IP address statistics for internet tunnels
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of the public IP address statistics for internet tunnels
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InternetAddressStatistics []*InternetAddressStatistics `json:"InternetAddressStatistics,omitnil" name:"InternetAddressStatistics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInternetAddressStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressStatisticsResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirectConnect struct {
	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitnil" name:"DirectConnectName"`

	// Access point ID of a connection.
	AccessPointId *string `json:"AccessPointId,omitnil" name:"AccessPointId"`

	// Connection status.
	// PENDING: Applying. 
	// REJECTED: Application rejected.   
	// TOPAY: Payment pending. 
	// PAID: Paid. 
	// ALLOCATED: Constructing.   
	// AVAILABLE: Available.  
	// DELETING: Deleting.
	// DELETED: Deleted.
	State *string `json:"State,omitnil" name:"State"`

	// Connection creation time.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Connection activation time.
	EnabledTime *string `json:"EnabledTime,omitnil" name:"EnabledTime"`

	// ISP that provides connections. Valid values: ChinaTelecom (China Telecom), ChinaMobile (China Mobile), ChinaUnicom (China Unicom), In-houseWiring (in-house wiring), ChinaOther (other Chinese ISPs), InternationalOperator (international ISPs).
	LineOperator *string `json:"LineOperator,omitnil" name:"LineOperator"`

	// Location of a local IDC.
	Location *string `json:"Location,omitnil" name:"Location"`

	// Connection port bandwidth in Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// User-side port type of a connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface; it is the default value), 1000Base-LX (1-Gigabit single-mode optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-mode optical Ethernet interface; 10 KM).
	PortType *string `json:"PortType,omitnil" name:"PortType"`

	// Circuit code of a connection, which is provided by the ISP or service provider.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CircuitCode *string `json:"CircuitCode,omitnil" name:"CircuitCode"`

	// ID of a redundant connection.
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil" name:"RedundantDirectConnectId"`

	// VLAN for connection debugging, which is enabled and automatically assigned by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// Tencent-side IP address for connection debugging.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address for connection debugging.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Name of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerName *string `json:"CustomerName,omitnil" name:"CustomerName"`

	// Email address of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerContactMail *string `json:"CustomerContactMail,omitnil" name:"CustomerContactMail"`

	// Contact number of the connection applicant, which is obtained from the account system by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil" name:"CustomerContactNumber"`

	// Connection expiration time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// Connection billing mode. NON_RECURRING_CHARGE: One-time charge for accessing service
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// Fault reporting contact person.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil" name:"FaultReportContactNumber"`

	// Tag key-value pair
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// Access point type of a connection.
	AccessPointType *string `json:"AccessPointType,omitnil" name:"AccessPointType"`

	// IDC city.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdcCity *string `json:"IdcCity,omitnil" name:"IdcCity"`

	// Billing status
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChargeState *string `json:"ChargeState,omitnil" name:"ChargeState"`

	// Connection activation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Whether the connection has the service agreement signed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`

	// Whether the connection is an edge zone.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LocalZone *bool `json:"LocalZone,omitnil" name:"LocalZone"`

	// Number of dedicated tunnels with disabled VLAN in the connection
	// Note: this field may return `null`, indicating that no valid value can be found.
	VlanZeroDirectConnectTunnelCount *uint64 `json:"VlanZeroDirectConnectTunnelCount,omitnil" name:"VlanZeroDirectConnectTunnelCount"`

	// Number of dedicated tunnels with enabled VLAN in the connection
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OtherVlanDirectConnectTunnelCount *uint64 `json:"OtherVlanDirectConnectTunnelCount,omitnil" name:"OtherVlanDirectConnectTunnelCount"`

	// Minimum bandwidth of the connection
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MinBandwidth *uint64 `json:"MinBandwidth,omitnil" name:"MinBandwidth"`
}

type DirectConnectTunnel struct {
	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`

	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Dedicated tunnel status.
	// AVAILABLE: Ready or connected.
	// PENDING: Applying.
	// ALLOCATING: Configuring.
	// ALLOCATED: Configured.
	// ALTERING: Modifying.
	// DELETING: Deleting.
	// DELETED: Deleted.
	// COMFIRMING: To be accepted.
	// REJECTED: Rejected.
	State *string `json:"State,omitnil" name:"State"`

	// Connection owner, i.e., developer account ID.
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil" name:"DirectConnectOwnerAccount"`

	// Dedicated tunnel owner, i.e., developer account ID.
	OwnerAccount *string `json:"OwnerAccount,omitnil" name:"OwnerAccount"`

	// Network type. Valid values: VPC, BMVPC, CCN.
	//  VPC: Virtual Private Cloud; BMVPC: BM VPC; CCN: Cloud Connect Network.
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// Network of the VPC region, such as `ap-guangzhou`.
	NetworkRegion *string `json:"NetworkRegion,omitnil" name:"NetworkRegion"`

	// Unified VPC ID or BMVPC ID.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Direct connect gateway ID.
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// BGP: BGP routing; STATIC: Static routing. Default value: BGP routing.
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// User-side BGP, including Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil" name:"BgpPeer"`

	// User-side IP range.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil" name:"RouteFilterPrefixes"`

	// VLAN of a dedicated tunnel.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// TencentAddress: Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// CustomerAddress: User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil" name:"DirectConnectTunnelName"`

	// Creation time of a dedicated tunnel.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Bandwidth value of a dedicated tunnel.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// Tag value of a dedicated tunnel.
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// Associated custom network probe ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// BGP community switch
	// Note: this field may return null, indicating that no valid values can be obtained.
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil" name:"EnableBGPCommunity"`

	// Whether it is a NAT tunnel
	// Note: this field may return null, indicating that no valid values can be obtained.
	NatType *int64 `json:"NatType,omitnil" name:"NatType"`

	// VPC region abbreviation, such as `gz`, `cd`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcRegion *string `json:"VpcRegion,omitnil" name:"VpcRegion"`

	// Whether to enable BFD
	// Note: this field may return null, indicating that no valid values can be obtained.
	BfdEnable *int64 `json:"BfdEnable,omitnil" name:"BfdEnable"`

	// Access point type of a dedicated tunnel.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessPointType *string `json:"AccessPointType,omitnil" name:"AccessPointType"`

	// Direct connect gateway name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// VPC name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// Backup IP address on the Tencent side.
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil" name:"TencentBackupAddress"`

	// Whether the connection associated with the dedicated tunnel has the service agreement signed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`

	// Cloud Attached Connection Service ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CloudAttachId *string `json:"CloudAttachId,omitnil" name:"CloudAttachId"`
}

// Predefined struct for user
type DisableInternetAddressRequestParams struct {
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DisableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DisableInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInternetAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisableInternetAddressResponseParams `json:"Response"`
}

func (r *DisableInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInternetAddressRequestParams struct {
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type EnableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *EnableInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInternetAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *EnableInternetAddressResponseParams `json:"Response"`
}

func (r *EnableInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Fields to be filtered.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Filter values of the field.
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type InternetAddressDetail struct {
	// Internet tunnel’s IP address ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Internet tunnel’s network address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Subnet *string `json:"Subnet,omitnil" name:"Subnet"`

	// Mask length of a network address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaskLen *int64 `json:"MaskLen,omitnil" name:"MaskLen"`

	// Address type. Valid values: 0: BGP
	// 1: China Telecom
	// 2: China Mobile
	// 3: China Unicom
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AddrType *int64 `json:"AddrType,omitnil" name:"AddrType"`

	// Address status. Valid values: 0: in use
	// 1: disabled
	// 2: returned
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Applied at
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ApplyTime *string `json:"ApplyTime,omitnil" name:"ApplyTime"`

	// Disabled at
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StopTime *string `json:"StopTime,omitnil" name:"StopTime"`

	// Returned at
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReleaseTime *string `json:"ReleaseTime,omitnil" name:"ReleaseTime"`

	// Region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// User ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Address protocol. Valid values: 0: IPv4; 1: IPv6
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AddrProto *int64 `json:"AddrProto,omitnil" name:"AddrProto"`

	// Retention period of a released IP address, in days
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReserveTime *int64 `json:"ReserveTime,omitnil" name:"ReserveTime"`
}

type InternetAddressStatistics struct {
	// Region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Number of public IP addresses for internet tunnels
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetNum *int64 `json:"SubnetNum,omitnil" name:"SubnetNum"`
}

// Predefined struct for user
type ModifyDirectConnectAttributeRequestParams struct {
	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitnil" name:"DirectConnectName"`

	// Circuit code of a connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitnil" name:"CircuitCode"`

	// VLAN for connection debugging.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// Tencent-side IP address for connection debugging.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address for connection debugging.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitnil" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitnil" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil" name:"FaultReportContactNumber"`

	// Whether the connection applicant has signed the service agreement.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`

	// Connection’s bandwidth
	Bandwidth *uint64 `json:"Bandwidth,omitnil" name:"Bandwidth"`
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Connection ID.
	DirectConnectId *string `json:"DirectConnectId,omitnil" name:"DirectConnectId"`

	// Connection name.
	DirectConnectName *string `json:"DirectConnectName,omitnil" name:"DirectConnectName"`

	// Circuit code of a connection, which is provided by the ISP or connection provider.
	CircuitCode *string `json:"CircuitCode,omitnil" name:"CircuitCode"`

	// VLAN for connection debugging.
	Vlan *int64 `json:"Vlan,omitnil" name:"Vlan"`

	// Tencent-side IP address for connection debugging.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address for connection debugging.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Name of connection applicant, which is obtained from the account system by default.
	CustomerName *string `json:"CustomerName,omitnil" name:"CustomerName"`

	// Email address of connection applicant, which is obtained from the account system by default.
	CustomerContactMail *string `json:"CustomerContactMail,omitnil" name:"CustomerContactMail"`

	// Contact number of connection applicant, which is obtained from the account system by default.
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil" name:"CustomerContactNumber"`

	// Fault reporting contact person.
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil" name:"FaultReportContactPerson"`

	// Fault reporting contact number.
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil" name:"FaultReportContactNumber"`

	// Whether the connection applicant has signed the service agreement.
	SignLaw *bool `json:"SignLaw,omitnil" name:"SignLaw"`

	// Connection’s bandwidth
	Bandwidth *uint64 `json:"Bandwidth,omitnil" name:"Bandwidth"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	delete(f, "DirectConnectName")
	delete(f, "CircuitCode")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "CustomerName")
	delete(f, "CustomerContactMail")
	delete(f, "CustomerContactNumber")
	delete(f, "FaultReportContactPerson")
	delete(f, "FaultReportContactNumber")
	delete(f, "SignLaw")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelAttributeRequestParams struct {
	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil" name:"DirectConnectTunnelName"`

	// User-side BGP, including Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil" name:"BgpPeer"`

	// User-side IP range.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil" name:"RouteFilterPrefixes"`

	// Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Bandwidth value of a dedicated tunnel in Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil" name:"TencentBackupAddress"`
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Dedicated tunnel ID.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`

	// Dedicated tunnel name.
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil" name:"DirectConnectTunnelName"`

	// User-side BGP, including Asn and AuthKey.
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil" name:"BgpPeer"`

	// User-side IP range.
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil" name:"RouteFilterPrefixes"`

	// Tencent-side IP address.
	TencentAddress *string `json:"TencentAddress,omitnil" name:"TencentAddress"`

	// User-side IP address.
	CustomerAddress *string `json:"CustomerAddress,omitnil" name:"CustomerAddress"`

	// Bandwidth value of a dedicated tunnel in Mbps.
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// Tencent-side standby IP address
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil" name:"TencentBackupAddress"`
}

func (r *ModifyDirectConnectTunnelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	delete(f, "DirectConnectTunnelName")
	delete(f, "BgpPeer")
	delete(f, "RouteFilterPrefixes")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "Bandwidth")
	delete(f, "TencentBackupAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectTunnelAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectTunnelAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NQAInfo struct {
	// Number of health checks
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitnil" name:"ProbeFailedTimes"`

	// Health check interval
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`

	// IP address for the health check
	DestinationIp *string `json:"DestinationIp,omitnil" name:"DestinationIp"`
}

// Predefined struct for user
type RejectDirectConnectTunnelRequestParams struct {
	// None.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

type RejectDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// None.
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil" name:"DirectConnectTunnelId"`
}

func (r *RejectDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectDirectConnectTunnelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RejectDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *RejectDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *RejectDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseInternetAddressRequestParams struct {
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ReleaseInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// ID of the internet tunnel’s public IP address
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *ReleaseInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseInternetAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReleaseInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseInternetAddressResponseParams `json:"Response"`
}

func (r *ReleaseInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteFilterPrefix struct {
	// User-side IP range.
	Cidr *string `json:"Cidr,omitnil" name:"Cidr"`
}

type Tag struct {
	// Tag key
	// Note: this field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil" name:"Key"`

	// Tag value
	// Note: this field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil" name:"Value"`
}