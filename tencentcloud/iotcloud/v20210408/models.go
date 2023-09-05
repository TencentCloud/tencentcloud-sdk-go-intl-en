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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Attribute struct {
	// Attribute list
	Tags []*DeviceTag `json:"Tags,omitnil" name:"Tags"`
}

type BindProductInfo struct {
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`
}

type CertInfo struct {
	// Certificate name
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// Hex sequence number of a certificate
	CertSN *string `json:"CertSN,omitnil" name:"CertSN"`

	// Certificate issuer
	IssuerName *string `json:"IssuerName,omitnil" name:"IssuerName"`

	// Certificate subject
	Subject *string `json:"Subject,omitnil" name:"Subject"`

	// Certificate creation time (timestamp in milliseconds)
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Certificate effective time (timestamp in milliseconds)
	EffectiveTime *uint64 `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// Certificate expiration time (timestamp in milliseconds)
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// X.509 certificate content
	CertText *string `json:"CertText,omitnil" name:"CertText"`
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
	// Product ID, globally unique ID assigned by Tencent Cloud during product creation
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name. It is a string of 1 to 48 characters. Letters, digits, and :_- are allowed.
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Device attribute
	Attribute *Attribute `json:"Attribute,omitnil" name:"Attribute"`

	// Whether to use custom PSK, no by default
	DefinedPsk *string `json:"DefinedPsk,omitnil" name:"DefinedPsk"`

	// ISP, required for a NB-IoT product. `1`: China Telecom; `2`: China Mobile; `3`: China Unicom
	Isp *uint64 `json:"Isp,omitnil" name:"Isp"`

	// IMEI, required for a NB-IoT product
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// DevEUI of a LoRa device, required when you create a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitnil" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitnil" name:"LoraMoteType"`

	// Skey, required when you create a LoRa device
	Skey *string `json:"Skey,omitnil" name:"Skey"`

	// AppKey of a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitnil" name:"LoraAppKey"`

	// Private CA certificate
	TlsCrt *string `json:"TlsCrt,omitnil" name:"TlsCrt"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Product ID, globally unique ID assigned by Tencent Cloud during product creation
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name. It is a string of 1 to 48 characters. Letters, digits, and :_- are allowed.
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Device attribute
	Attribute *Attribute `json:"Attribute,omitnil" name:"Attribute"`

	// Whether to use custom PSK, no by default
	DefinedPsk *string `json:"DefinedPsk,omitnil" name:"DefinedPsk"`

	// ISP, required for a NB-IoT product. `1`: China Telecom; `2`: China Mobile; `3`: China Unicom
	Isp *uint64 `json:"Isp,omitnil" name:"Isp"`

	// IMEI, required for a NB-IoT product
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// DevEUI of a LoRa device, required when you create a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitnil" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitnil" name:"LoraMoteType"`

	// Skey, required when you create a LoRa device
	Skey *string `json:"Skey,omitnil" name:"Skey"`

	// AppKey of a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitnil" name:"LoraAppKey"`

	// Private CA certificate
	TlsCrt *string `json:"TlsCrt,omitnil" name:"TlsCrt"`
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
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Base64-encoded symmetric encryption key, which is returned if symmetric encryption is used
	DevicePsk *string `json:"DevicePsk,omitnil" name:"DevicePsk"`

	// Device certificate, which authenticates client identity during TLS connection establishment and is returned if asymmetric encryption is used
	DeviceCert *string `json:"DeviceCert,omitnil" name:"DeviceCert"`

	// Device private key, which authenticates client identity during TLS connection establishment and is returned if asymmetric encryption is used. Tencent Cloud does not store the key. Please store it by yourself properly.
	DevicePrivateKey *string `json:"DevicePrivateKey,omitnil" name:"DevicePrivateKey"`

	// DevEUI of a LoRa device, which is returned for a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitnil" name:"LoraDevEui"`

	// MoteType of a LoRa device, which is returned for a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitnil" name:"LoraMoteType"`

	// AppKey of a LoRa device, which is returned for a LoRa device
	LoraAppKey *string `json:"LoraAppKey,omitnil" name:"LoraAppKey"`

	// NwkKey of a LoRa device, which is returned for a LoRa device
	LoraNwkKey *string `json:"LoraNwkKey,omitnil" name:"LoraNwkKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitnil" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitnil" name:"VerifyCertText"`
}

type CreatePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// CA certificate name
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitnil" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitnil" name:"VerifyCertText"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitnil" name:"ProductProperties"`

	// Skey, which is required to create a CLAA product.
	Skey *string `json:"Skey,omitnil" name:"Skey"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// Product name, which cannot be same as that of an existing product. Naming rule: [a-zA-Z0-9:_-]{1,32}.
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitnil" name:"ProductProperties"`

	// Skey, which is required to create a CLAA product.
	Skey *string `json:"Skey,omitnil" name:"Skey"`
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
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Product ID, the globally unique ID assigned by Tencent Cloud.
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitnil" name:"ProductProperties"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Name of the device to delete
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Skey, which is required to delete a LoRa device or LoRa gateway device
	Skey *string `json:"Skey,omitnil" name:"Skey"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Name of the device to delete
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Skey, which is required to delete a LoRa device or LoRa gateway device
	Skey *string `json:"Skey,omitnil" name:"Skey"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
}

type DeleteDeviceShadowRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CertName *string `json:"CertName,omitnil" name:"CertName"`
}

type DeletePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// Private CA certificate name
	CertName *string `json:"CertName,omitnil" name:"CertName"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Skey, which is required to delete a LoRa product
	Skey *string `json:"Skey,omitnil" name:"Skey"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to delete
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Skey, which is required to delete a LoRa product
	Skey *string `json:"Skey,omitnil" name:"Skey"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`
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
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Whether the device is online. `0`: offline; `1`: online
	Online *uint64 `json:"Online,omitnil" name:"Online"`

	// Device login time
	LoginTime *uint64 `json:"LoginTime,omitnil" name:"LoginTime"`

	// Device firmware version
	Version *string `json:"Version,omitnil" name:"Version"`

	// Last updated time of the device
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`

	// Device certificate
	DeviceCert *string `json:"DeviceCert,omitnil" name:"DeviceCert"`

	// Device key
	DevicePsk *string `json:"DevicePsk,omitnil" name:"DevicePsk"`

	// Device attribute
	Tags []*DeviceTag `json:"Tags,omitnil" name:"Tags"`

	// Device type
	DeviceType *uint64 `json:"DeviceType,omitnil" name:"DeviceType"`

	// International Mobile Equipment Identity (IMEI)
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// ISP
	Isp *uint64 `json:"Isp,omitnil" name:"Isp"`

	// IP address
	ConnIP *uint64 `json:"ConnIP,omitnil" name:"ConnIP"`

	// Device ID at the NB-IoT ISP
	NbiotDeviceID *string `json:"NbiotDeviceID,omitnil" name:"NbiotDeviceID"`

	// DevEUI of a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitnil" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitnil" name:"LoraMoteType"`

	// SDK log level of the device
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogLevel *uint64 `json:"LogLevel,omitnil" name:"LogLevel"`

	// The first time when the device went online
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitnil" name:"FirstOnlineTime"`

	// The last time when the device went offline
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitnil" name:"LastOfflineTime"`

	// Device creation time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whether the device certificate has been obtained. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CertState *uint64 `json:"CertState,omitnil" name:"CertState"`

	// Whether the device is enabled
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EnableState *uint64 `json:"EnableState,omitnil" name:"EnableState"`

	// Device tags
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Labels []*DeviceLabel `json:"Labels,omitnil" name:"Labels"`

	// IP address of the MQTT client
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ClientIP *string `json:"ClientIP,omitnil" name:"ClientIP"`

	// Firmware update time of the device
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitnil" name:"FirmwareUpdateTime"`

	// Account ID of the creator
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateUserId *uint64 `json:"CreateUserId,omitnil" name:"CreateUserId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Offset, which starts from 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Page size. Value range: 10-250
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Device firmware version. If no value is passed in, devices of all firmware versions are returned. If `None-FirmwareVersion` is passed in, devices without version numbers are returned.
	FirmwareVersion *string `json:"FirmwareVersion,omitnil" name:"FirmwareVersion"`

	// Device name to query
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Whether to query enabled or disabled devices. `0`: disabled devices; `1`: enabled devices. By default, both enabled and disabled devices are queried.
	EnableState *uint64 `json:"EnableState,omitnil" name:"EnableState"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product whose devices are queried
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Offset, which starts from 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Page size. Value range: 10-250
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Device firmware version. If no value is passed in, devices of all firmware versions are returned. If `None-FirmwareVersion` is passed in, devices without version numbers are returned.
	FirmwareVersion *string `json:"FirmwareVersion,omitnil" name:"FirmwareVersion"`

	// Device name to query
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Whether to query enabled or disabled devices. `0`: disabled devices; `1`: enabled devices. By default, both enabled and disabled devices are queried.
	EnableState *uint64 `json:"EnableState,omitnil" name:"EnableState"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of device details
	Devices []*DeviceInfo `json:"Devices,omitnil" name:"Devices"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// Offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Maximum number of records to return, which is 20 by default and cannot exceed 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribePrivateCABindedProductsRequest struct {
	*tchttp.BaseRequest
	
	// Certificate name
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// Offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Maximum number of records to return, which is 20 by default and cannot exceed 200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	Products []*BindProductInfo `json:"Products,omitnil" name:"Products"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CertName *string `json:"CertName,omitnil" name:"CertName"`
}

type DescribePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// Name of the private CA certificate to query
	CertName *string `json:"CertName,omitnil" name:"CertName"`
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
	CA *CertInfo `json:"CA,omitnil" name:"CA"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CAs []*CertInfo `json:"CAs,omitnil" name:"CAs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`
}

type DescribeProductCARequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`
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
	CAs []*CertInfo `json:"CAs,omitnil" name:"CAs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Product metadata
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitnil" name:"ProductMetadata"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitnil" name:"ProductProperties"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of entries returned per page. Valid range: 10–250.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// Offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of entries returned per page. Valid range: 10–250.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of product details
	Products []*ProductInfo `json:"Products,omitnil" name:"Products"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Whether the device is online. `0`: offline; `1`: online
	Online *uint64 `json:"Online,omitnil" name:"Online"`

	// Device login time
	LoginTime *uint64 `json:"LoginTime,omitnil" name:"LoginTime"`

	// Device version
	Version *string `json:"Version,omitnil" name:"Version"`

	// Device certificate, which is returned for devices that use certificates for authentication
	DeviceCert *string `json:"DeviceCert,omitnil" name:"DeviceCert"`

	// Device key, which is returned for devices that use keys for authentication
	DevicePsk *string `json:"DevicePsk,omitnil" name:"DevicePsk"`

	// Device attribute
	Tags []*DeviceTag `json:"Tags,omitnil" name:"Tags"`

	// Device type
	DeviceType *uint64 `json:"DeviceType,omitnil" name:"DeviceType"`

	// International Mobile Equipment Identity (IMEI)
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// ISP
	Isp *uint64 `json:"Isp,omitnil" name:"Isp"`

	// Device ID at the NB-IoT ISP
	NbiotDeviceID *string `json:"NbiotDeviceID,omitnil" name:"NbiotDeviceID"`

	// IP address
	ConnIP *uint64 `json:"ConnIP,omitnil" name:"ConnIP"`

	// Last updated time of the device
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`

	// DevEUI of a LoRa device
	LoraDevEui *string `json:"LoraDevEui,omitnil" name:"LoraDevEui"`

	// MoteType of a LoRa device
	LoraMoteType *uint64 `json:"LoraMoteType,omitnil" name:"LoraMoteType"`

	// The first time when the device went online
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitnil" name:"FirstOnlineTime"`

	// The last time when the device went offline
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitnil" name:"LastOfflineTime"`

	// Device creation time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Device log level
	// Note: this field may return `null`, indicating that no valid value is obtained.
	LogLevel *uint64 `json:"LogLevel,omitnil" name:"LogLevel"`

	// Whether the device certificate has been obtained. `0`: no; `1`: yes
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CertState *uint64 `json:"CertState,omitnil" name:"CertState"`

	// Whether the device is enabled. `0`: disabled; `1`: enabled
	// Note: this field may return `null`, indicating that no valid value is obtained.
	EnableState *uint64 `json:"EnableState,omitnil" name:"EnableState"`

	// Device tags
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Labels []*DeviceLabel `json:"Labels,omitnil" name:"Labels"`

	// IP address of the MQTT client
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ClientIP *string `json:"ClientIP,omitnil" name:"ClientIP"`

	// Time of last OTA update
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitnil" name:"FirmwareUpdateTime"`
}

type DeviceLabel struct {
	// Tag key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DeviceTag struct {
	// Attribute name
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// Attribute value type. `1`: integer; `2`: string
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// Attribute value
	Value *string `json:"Value,omitnil" name:"Value"`

	// Attribute description
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Name *string `json:"Name,omitnil" name:"Name"`
}

type ProductInfo struct {
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Product name
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Product metadata
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitnil" name:"ProductMetadata"`

	// Product properties
	ProductProperties *ProductProperties `json:"ProductProperties,omitnil" name:"ProductProperties"`
}

type ProductMetadata struct {
	// Product creation time
	CreationDate *uint64 `json:"CreationDate,omitnil" name:"CreationDate"`
}

type ProductProperties struct {
	// Product description
	ProductDescription *string `json:"ProductDescription,omitnil" name:"ProductDescription"`

	// Authentication type. `1` (default): certificate; `2`: signature
	EncryptionType *string `json:"EncryptionType,omitnil" name:"EncryptionType"`

	// Product region. Valid value: `gz` (Guangzhou)
	Region *string `json:"Region,omitnil" name:"Region"`

	// Product type. Valid values:
	// `0` (default): general; `2`: NB-IoT; `3`: LoRa gateway; `4`: LoRa; `5`: general gateway
	ProductType *uint64 `json:"ProductType,omitnil" name:"ProductType"`

	// Data format. Valid values: `json` (default), `custom`
	Format *string `json:"Format,omitnil" name:"Format"`

	// Platform of the product. Default value: `0`
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// AppEUI at the LoRa product operator, required only for LoRa products
	Appeui *string `json:"Appeui,omitnil" name:"Appeui"`

	// ID of the Thing Specification Language (TSL) model bound to the product. `-1` means no models are bound.
	ModelId *string `json:"ModelId,omitnil" name:"ModelId"`

	// Name of the TSL model bound to the product
	ModelName *string `json:"ModelName,omitnil" name:"ModelName"`

	// Product key, which is specific to suite products
	ProductKey *string `json:"ProductKey,omitnil" name:"ProductKey"`

	// Dynamic registration type. `0`: disable; `1`: preset device names; `2`: generate device names dynamically
	RegisterType *uint64 `json:"RegisterType,omitnil" name:"RegisterType"`

	// Dynamic registration product key
	ProductSecret *string `json:"ProductSecret,omitnil" name:"ProductSecret"`

	// The maximum number of devices that can be dynamically created when `RegisterType` is set to `2`
	RegisterLimit *uint64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`

	// Original product ID of a transferred product. This parameter is empty for products that are not transferred.
	OriginProductId *string `json:"OriginProductId,omitnil" name:"OriginProductId"`

	// Private CA certificate name
	PrivateCAName *string `json:"PrivateCAName,omitnil" name:"PrivateCAName"`

	// Original user ID of a transferred product. This parameter is empty for products that are not transferred.
	OriginUserId *uint64 `json:"OriginUserId,omitnil" name:"OriginUserId"`
}

// Predefined struct for user
type SetProductsForbiddenStatusRequestParams struct {
	// List of products to enable or disable
	ProductId []*string `json:"ProductId,omitnil" name:"ProductId"`

	// `0`: enable; `1`: disable
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type SetProductsForbiddenStatusRequest struct {
	*tchttp.BaseRequest
	
	// List of products to enable or disable
	ProductId []*string `json:"ProductId,omitnil" name:"ProductId"`

	// `0`: enable; `1`: disable
	Status *uint64 `json:"Status,omitnil" name:"Status"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Log level. `0`: disable; `1`: error; `2`: warning; `3`: information; `4`: debugging
	LogLevel *uint64 `json:"LogLevel,omitnil" name:"LogLevel"`
}

type UpdateDeviceLogLevelRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device name
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// Log level. `0`: disable; `1`: error; `2`: warning; `3`: information; `4`: debugging
	LogLevel *uint64 `json:"LogLevel,omitnil" name:"LogLevel"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device names
	DeviceNames []*string `json:"DeviceNames,omitnil" name:"DeviceNames"`

	// New status of the devices. `0`: disabled; `1`: enabled
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the product to which the device belongs
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Device names
	DeviceNames []*string `json:"DeviceNames,omitnil" name:"DeviceNames"`

	// New status of the devices. `0`: disabled; `1`: enabled
	Status *uint64 `json:"Status,omitnil" name:"Status"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitnil" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitnil" name:"VerifyCertText"`
}

type UpdatePrivateCARequest struct {
	*tchttp.BaseRequest
	
	// CA certificate name
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// CA certificate content
	CertText *string `json:"CertText,omitnil" name:"CertText"`

	// Content verifying the CA certificate
	VerifyCertText *string `json:"VerifyCertText,omitnil" name:"VerifyCertText"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Dynamic registration type. Valid values: 0 - disabled; 1 - pre-create device; 2 - auto-create device.
	RegisterType *uint64 `json:"RegisterType,omitnil" name:"RegisterType"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`
}

type UpdateProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest
	
	// Product ID
	ProductId *string `json:"ProductId,omitnil" name:"ProductId"`

	// Dynamic registration type. Valid values: 0 - disabled; 1 - pre-create device; 2 - auto-create device.
	RegisterType *uint64 `json:"RegisterType,omitnil" name:"RegisterType"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`
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
	RegisterType *uint64 `json:"RegisterType,omitnil" name:"RegisterType"`

	// Product key for dynamic registration
	ProductSecret *string `json:"ProductSecret,omitnil" name:"ProductSecret"`

	// Maximum dynamically registered devices
	RegisterLimit *uint64 `json:"RegisterLimit,omitnil" name:"RegisterLimit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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