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

package v20210408

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Attribute struct {
	// Attribute list
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`
}

type BindProductInfo struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type CertInfo struct {
	// Certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Hex sequence number of a certificate
	CertSN *string `json:"CertSN,omitempty" name:"CertSN"`

	// Certificate issuer
	IssuerName *string `json:"IssuerName,omitempty" name:"IssuerName"`

	// Certificate subject
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// Certificate creation time (timestamp in milliseconds)
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Certificate effective time (timestamp in milliseconds)
	EffectiveTime *uint64 `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// Certificate expiration time (timestamp in milliseconds)
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// X.509 certificate content
	CertText *string `json:"CertText,omitempty" name:"CertText"`
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
	// Product ID, globally unique ID assigned by Tencent Cloud during product creation
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name. It is a string of 1 to 48 characters. Letters, digits, and :_- are allowed.
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Device attribute
	Attribute *Attribute `json:"Attribute,omitempty" name:"Attribute"`

	// Whether to use custom PSK, no by default
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`

	// ISP, required for a NB-IoT product. `1`: China Telecom; `2`: China Mobile; `3`: China Unicom
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// IMEI, required for a NB-IoT product
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// DevEUI of a LoRa device, required when you create a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// Skey, required when you create a LoRa device
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// AppKey of a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

	// Private CA certificate
	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Product ID, globally unique ID assigned by Tencent Cloud during product creation
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name. It is a string of 1 to 48 characters. Letters, digits, and :_- are allowed.
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Device attribute
	Attribute *Attribute `json:"Attribute,omitempty" name:"Attribute"`

	// Whether to use custom PSK, no by default
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`

	// ISP, required for a NB-IoT product. `1`: China Telecom; `2`: China Mobile; `3`: China Unicom
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// IMEI, required for a NB-IoT product
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// DevEUI of a LoRa device, required when you create a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// Skey, required when you create a LoRa device
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// AppKey of a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

	// Private CA certificate
	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Attribute")
	delete(f, "DefinedPsk")
	delete(f, "Isp")
	delete(f, "Imei")
	delete(f, "LoraDevEui")
	delete(f, "LoraMoteType")
	delete(f, "Skey")
	delete(f, "LoraAppKey")
	delete(f, "TlsCrt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceResponseParams struct {
	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Base64-encoded symmetric encryption key, which is returned if symmetric encryption is used
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// Device certificate, which authenticates client identity during TLS connection establishment and is returned if asymmetric encryption is used
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// Device private key, which authenticates client identity during TLS connection establishment and is returned if asymmetric encryption is used. Tencent Cloud does not store the key. Please store it by yourself properly.
	DevicePrivateKey *string `json:"DevicePrivateKey,omitempty" name:"DevicePrivateKey"`

	// DevEUI of a LoRa device, which is returned for a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// MoteType of a LoRa device, which is returned for a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// AppKey of a LoRa device, which is returned for a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

	// NwkKey of a LoRa device, which is returned for a LoRa device
	LoraNwkKey *string `json:"LoraNwkKey,omitempty" name:"LoraNwkKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceResponseParams `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateCARequestParams struct {
	// CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

type CreatePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

func (r *CreatePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "CertText")
	delete(f, "VerifyCertText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateCAResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateCAResponseParams `json:"Response"`
}

func (r *CreatePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductRequestParams struct {
	// Product name, which cannot be same as that of an existing product. Naming rule: [a-zA-Z0-9:_-]{1,32}.
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// Skey, which is required to create a CLAA product.
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// Product name, which cannot be same as that of an existing product. Naming rule: [a-zA-Z0-9:_-]{1,32}.
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// Skey, which is required to create a CLAA product.
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "ProductProperties")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductResponseParams struct {
	// Product name
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Product ID, the globally unique ID assigned by Tencent Cloud.
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateProductResponseParams `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Name of the device to delete
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Skey, which is required to delete a LoRa device or LoRa gateway device
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Name of the device to delete
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Skey, which is required to delete a LoRa device or LoRa gateway device
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceResponseParams `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceShadowRequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DeleteDeviceShadowRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DeleteDeviceShadowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceShadowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceShadowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceShadowResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceShadowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceShadowResponseParams `json:"Response"`
}

func (r *DeleteDeviceShadowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceShadowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateCARequestParams struct {
	// Private CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

type DeletePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// Private CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

func (r *DeletePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateCAResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivateCAResponseParams `json:"Response"`
}

func (r *DeletePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductRequestParams struct {
	// ID of the product to delete
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Skey, which is required to delete a LoRa product
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to delete
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Skey, which is required to delete a LoRa product
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProductResponseParams `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Whether the device is online. `0`: offline; `1`: online
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// Device login time
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// Device firmware version
	Version *string `json:"Version,omitempty" name:"Version"`

	// Last updated time of the device
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// Device certificate
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// Device key
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// Device attribute
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

	// Device type
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// International Mobile Equipment Identity (IMEI)
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// ISP
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// IP address
	ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

	// Device ID at the NB-IoT ISP
	NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

	// DevEUI of a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// SDK log level of the device
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

	// The first time when the device went online
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// The last time when the device went offline
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

	// Device creation time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the device certificate has been obtained. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

	// Whether the device is enabled
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// Device tags
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

	// IP address of the MQTT client
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// Firmware update time of the device
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`

	// Account ID of the creator
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateUserId *uint64 `json:"CreateUserId,omitempty" name:"CreateUserId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResponseParams `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// ID of the product whose devices are queried
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Offset, which starts from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. Value range: 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Device firmware version. If no value is passed in, devices of all firmware versions are returned. If `None-FirmwareVersion` is passed in, devices without version numbers are returned.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// Device name to query
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Whether to query enabled or disabled devices. `0`: disabled devices; `1`: enabled devices. By default, both enabled and disabled devices are queried.
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product whose devices are queried
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Offset, which starts from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Page size. Value range: 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Device firmware version. If no value is passed in, devices of all firmware versions are returned. If `None-FirmwareVersion` is passed in, devices without version numbers are returned.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// Device name to query
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Whether to query enabled or disabled devices. `0`: disabled devices; `1`: enabled devices. By default, both enabled and disabled devices are queried.
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FirmwareVersion")
	delete(f, "DeviceName")
	delete(f, "EnableState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// Total number of the devices returned
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of device details
	Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCABindedProductsRequestParams struct {
	// Certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Offset for query
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of records to return, which is 20 by default and cannot exceed 200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePrivateCABindedProductsRequest struct {
	*tchttp.BaseRequest
	
	// Certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// Offset for query
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of records to return, which is 20 by default and cannot exceed 200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrivateCABindedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCABindedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCABindedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCABindedProductsResponseParams struct {
	// List of the products bound to the private CA certificate
	Products []*BindProductInfo `json:"Products,omitempty" name:"Products"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivateCABindedProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateCABindedProductsResponseParams `json:"Response"`
}

func (r *DescribePrivateCABindedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCABindedProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCARequestParams struct {
	// Name of the private CA certificate to query
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

type DescribePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// Name of the private CA certificate to query
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

func (r *DescribePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCAResponseParams struct {
	// Details of the private CA certificate
	CA *CertInfo `json:"CA,omitempty" name:"CA"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateCAResponseParams `json:"Response"`
}

func (r *DescribePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCAsRequestParams struct {

}

type DescribePrivateCAsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePrivateCAsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCAsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateCAsResponseParams struct {
	// List of private CA certificates
	CAs []*CertInfo `json:"CAs,omitempty" name:"CAs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivateCAsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateCAsResponseParams `json:"Response"`
}

func (r *DescribePrivateCAsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductCARequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductCARequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductCAResponseParams struct {
	// List of CA certificates bound to the product
	CAs []*CertInfo `json:"CAs,omitempty" name:"CAs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductCAResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductCAResponseParams `json:"Response"`
}

func (r *DescribeProductCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductRequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResponseParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Product metadata
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitempty" name:"ProductMetadata"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductResponseParams `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsRequestParams struct {
	// Offset, starting from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries returned per page. Valid range: 10–250.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// Offset, starting from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries returned per page. Valid range: 10–250.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// Total number of products
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of product details
	Products []*ProductInfo `json:"Products,omitempty" name:"Products"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductsResponseParams `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceInfo struct {
	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Whether the device is online. `0`: offline; `1`: online
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// Device login time
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// Device version
	Version *string `json:"Version,omitempty" name:"Version"`

	// Device certificate, which is returned for devices that use certificates for authentication
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// Device key, which is returned for devices that use keys for authentication
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// Device attribute
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

	// Device type
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// International Mobile Equipment Identity (IMEI)
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// ISP
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// Device ID at the NB-IoT ISP
	NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

	// IP address
	ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

	// Last updated time of the device
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// DevEUI of a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// The first time when the device went online
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// The last time when the device went offline
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

	// Device creation time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Device log level
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

	// Whether the device certificate has been obtained. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

	// Whether the device is enabled. `0`: disabled; `1`: enabled
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// Device tags
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

	// IP address of the MQTT client
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// Time of last OTA update
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`
}

type DeviceLabel struct {
	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeviceTag struct {
	// Attribute name
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// Attribute value type. `1`: integer; `2`: string
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Attribute value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Attribute description
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ProductInfo struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// Product metadata
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitempty" name:"ProductMetadata"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`
}

type ProductMetadata struct {
	// Product creation time
	CreationDate *uint64 `json:"CreationDate,omitempty" name:"CreationDate"`
}

type ProductProperties struct {
	// Product description
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// Authentication type. `1` (default): certificate; `2`: signature
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// Product region. Valid value: `gz` (Guangzhou)
	Region *string `json:"Region,omitempty" name:"Region"`

	// Product type. Valid values:
	// `0` (default): general; `2`: NB-IoT; `3`: LoRa gateway; `4`: LoRa; `5`: general gateway
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// Data format. Valid values: `json` (default), `custom`
	Format *string `json:"Format,omitempty" name:"Format"`

	// Platform of the product. Default value: `0`
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// AppEUI at the LoRa product operator, required only for LoRa products
	Appeui *string `json:"Appeui,omitempty" name:"Appeui"`

	// ID of the Thing Specification Language (TSL) model bound to the product. `-1` means no models are bound.
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// Name of the TSL model bound to the product
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// Product key, which is specific to suite products
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// Dynamic registration type. `0`: disable; `1`: preset device names; `2`: generate device names dynamically
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// Dynamic registration product key
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// The maximum number of devices that can be dynamically created when `RegisterType` is set to `2`
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// Original product ID of a transferred product. This parameter is empty for products that are not transferred.
	OriginProductId *string `json:"OriginProductId,omitempty" name:"OriginProductId"`

	// Private CA certificate name
	PrivateCAName *string `json:"PrivateCAName,omitempty" name:"PrivateCAName"`

	// Original user ID of a transferred product. This parameter is empty for products that are not transferred.
	OriginUserId *uint64 `json:"OriginUserId,omitempty" name:"OriginUserId"`
}

// Predefined struct for user
type SetProductsForbiddenStatusRequestParams struct {
	// List of products to enable or disable
	ProductId []*string `json:"ProductId,omitempty" name:"ProductId"`

	// `0`: enable; `1`: disable
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type SetProductsForbiddenStatusRequest struct {
	*tchttp.BaseRequest
	
	// List of products to enable or disable
	ProductId []*string `json:"ProductId,omitempty" name:"ProductId"`

	// `0`: enable; `1`: disable
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetProductsForbiddenStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetProductsForbiddenStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetProductsForbiddenStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetProductsForbiddenStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetProductsForbiddenStatusResponseParams `json:"Response"`
}

func (r *SetProductsForbiddenStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceLogLevelRequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Log level. `0`: disable; `1`: error; `2`: warning; `3`: information; `4`: debugging
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

type UpdateDeviceLogLevelRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// Log level. `0`: disable; `1`: error; `2`: warning; `3`: information; `4`: debugging
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

func (r *UpdateDeviceLogLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceLogLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "LogLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceLogLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceLogLevelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDeviceLogLevelResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceLogLevelResponseParams `json:"Response"`
}

func (r *UpdateDeviceLogLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceLogLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateRequestParams struct {
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device names
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// New status of the devices. `0`: disabled; `1`: enabled
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Device names
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// New status of the devices. `0`: disabled; `1`: enabled
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateDevicesEnableStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicesEnableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDevicesEnableStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDevicesEnableStateResponseParams `json:"Response"`
}

func (r *UpdateDevicesEnableStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrivateCARequestParams struct {
	// CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

type UpdatePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// CA certificate name
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

func (r *UpdatePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "CertText")
	delete(f, "VerifyCertText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrivateCAResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrivateCAResponseParams `json:"Response"`
}

func (r *UpdatePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProductDynamicRegisterRequestParams struct {
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Dynamic registration type. Valid values: 0 - disabled; 1 - pre-create device; 2 - auto-create device.
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

type UpdateProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Dynamic registration type. Valid values: 0 - disabled; 1 - pre-create device; 2 - auto-create device.
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

func (r *UpdateProductDynamicRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "RegisterType")
	delete(f, "RegisterLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProductDynamicRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProductDynamicRegisterResponseParams struct {
	// Dynamic registration type. Valid values: 0 - disabled; 1 - pre-create device; 2 - auto-create device.
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// Product key for dynamic registration
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateProductDynamicRegisterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProductDynamicRegisterResponseParams `json:"Response"`
}

func (r *UpdateProductDynamicRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}