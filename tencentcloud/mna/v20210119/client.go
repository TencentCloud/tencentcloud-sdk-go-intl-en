// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20210119

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-01-19"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewActivateHardwareRequest() (request *ActivateHardwareRequest) {
    request = &ActivateHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ActivateHardware")
    
    
    return
}

func NewActivateHardwareResponse() (response *ActivateHardwareResponse) {
    response = &ActivateHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateHardware
// Activate hardware device
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) ActivateHardware(request *ActivateHardwareRequest) (response *ActivateHardwareResponse, err error) {
    return c.ActivateHardwareWithContext(context.Background(), request)
}

// ActivateHardware
// Activate hardware device
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) ActivateHardwareWithContext(ctx context.Context, request *ActivateHardwareRequest) (response *ActivateHardwareResponse, err error) {
    if request == nil {
        request = NewActivateHardwareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "ActivateHardware")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewAddApplicationRequest() (request *AddApplicationRequest) {
    request = &AddApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddApplication")
    
    
    return
}

func NewAddApplicationResponse() (response *AddApplicationResponse) {
    response = &AddApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddApplication
// This API is used to create an application
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddApplication(request *AddApplicationRequest) (response *AddApplicationResponse, err error) {
    return c.AddApplicationWithContext(context.Background(), request)
}

// AddApplication
// This API is used to create an application
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddApplicationWithContext(ctx context.Context, request *AddApplicationRequest) (response *AddApplicationResponse, err error) {
    if request == nil {
        request = NewAddApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "AddApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewAddDeviceRequest() (request *AddDeviceRequest) {
    request = &AddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddDevice")
    
    
    return
}

func NewAddDeviceResponse() (response *AddDeviceResponse) {
    response = &AddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddDevice
// Create new device records
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDATAKEY = "InternalError.DuplicateDataKey"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INTERNALERROR_UNDEFINEDENCRYPTEDKEY = "InternalError.UndefinedEncryptedKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddDevice(request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    return c.AddDeviceWithContext(context.Background(), request)
}

// AddDevice
// Create new device records
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDATAKEY = "InternalError.DuplicateDataKey"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INTERNALERROR_UNDEFINEDENCRYPTEDKEY = "InternalError.UndefinedEncryptedKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddDeviceWithContext(ctx context.Context, request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    if request == nil {
        request = NewAddDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "AddDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAddGroupRequest() (request *AddGroupRequest) {
    request = &AddGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddGroup")
    
    
    return
}

func NewAddGroupResponse() (response *AddGroupResponse) {
    response = &AddGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddGroup
// Create group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddGroup(request *AddGroupRequest) (response *AddGroupResponse, err error) {
    return c.AddGroupWithContext(context.Background(), request)
}

// AddGroup
// Create group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddGroupWithContext(ctx context.Context, request *AddGroupRequest) (response *AddGroupResponse, err error) {
    if request == nil {
        request = NewAddGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "AddGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAddHardwareRequest() (request *AddHardwareRequest) {
    request = &AddHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddHardware")
    
    
    return
}

func NewAddHardwareResponse() (response *AddHardwareResponse) {
    response = &AddHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddHardware
// Add hardware devices, generate inactive hardware devices, and support batch addition
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddHardware(request *AddHardwareRequest) (response *AddHardwareResponse, err error) {
    return c.AddHardwareWithContext(context.Background(), request)
}

// AddHardware
// Add hardware devices, generate inactive hardware devices, and support batch addition
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddHardwareWithContext(ctx context.Context, request *AddHardwareRequest) (response *AddHardwareResponse, err error) {
    if request == nil {
        request = NewAddHardwareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "AddHardware")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewAddL3ConnRequest() (request *AddL3ConnRequest) {
    request = &AddL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddL3Conn")
    
    
    return
}

func NewAddL3ConnResponse() (response *AddL3ConnResponse) {
    response = &AddL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddL3Conn
// Create an interconnection rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_L3CIDROVERLAP = "OperationDenied.L3CidrOverLap"
//  OPERATIONDENIED_L3CONNECTIONOVERSIZE = "OperationDenied.L3ConnectionOverSize"
func (c *Client) AddL3Conn(request *AddL3ConnRequest) (response *AddL3ConnResponse, err error) {
    return c.AddL3ConnWithContext(context.Background(), request)
}

// AddL3Conn
// Create an interconnection rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_L3CIDROVERLAP = "OperationDenied.L3CidrOverLap"
//  OPERATIONDENIED_L3CONNECTIONOVERSIZE = "OperationDenied.L3ConnectionOverSize"
func (c *Client) AddL3ConnWithContext(ctx context.Context, request *AddL3ConnRequest) (response *AddL3ConnResponse, err error) {
    if request == nil {
        request = NewAddL3ConnRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "AddL3Conn")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddL3ConnResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEncryptedKeyRequest() (request *CreateEncryptedKeyRequest) {
    request = &CreateEncryptedKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "CreateEncryptedKey")
    
    
    return
}

func NewCreateEncryptedKeyResponse() (response *CreateEncryptedKeyResponse) {
    response = &CreateEncryptedKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEncryptedKey
// This API is used to configure and refresh preset keys.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateEncryptedKey(request *CreateEncryptedKeyRequest) (response *CreateEncryptedKeyResponse, err error) {
    return c.CreateEncryptedKeyWithContext(context.Background(), request)
}

// CreateEncryptedKey
// This API is used to configure and refresh preset keys.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateEncryptedKeyWithContext(ctx context.Context, request *CreateEncryptedKeyRequest) (response *CreateEncryptedKeyResponse, err error) {
    if request == nil {
        request = NewCreateEncryptedKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "CreateEncryptedKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEncryptedKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEncryptedKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplication
// This API is used to delete applications
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
}

// DeleteApplication
// This API is used to delete applications
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DeleteApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteDevice")
    
    
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDevice
// Delete device info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// Delete device info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDeviceWithContext(ctx context.Context, request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DeleteDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// Delete group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// Delete group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DeleteGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL3ConnRequest() (request *DeleteL3ConnRequest) {
    request = &DeleteL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteL3Conn")
    
    
    return
}

func NewDeleteL3ConnResponse() (response *DeleteL3ConnResponse) {
    response = &DeleteL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL3Conn
// Delete an interconnection rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteL3Conn(request *DeleteL3ConnRequest) (response *DeleteL3ConnResponse, err error) {
    return c.DeleteL3ConnWithContext(context.Background(), request)
}

// DeleteL3Conn
// Delete an interconnection rule
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteL3ConnWithContext(ctx context.Context, request *DeleteL3ConnRequest) (response *DeleteL3ConnResponse, err error) {
    if request == nil {
        request = NewDeleteL3ConnRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DeleteL3Conn")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL3ConnResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRegionsRequest() (request *DescribeAccessRegionsRequest) {
    request = &DescribeAccessRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DescribeAccessRegions")
    
    
    return
}

func NewDescribeAccessRegionsResponse() (response *DescribeAccessRegionsResponse) {
    response = &DescribeAccessRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessRegions
// Query the access region list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAccessRegions(request *DescribeAccessRegionsRequest) (response *DescribeAccessRegionsResponse, err error) {
    return c.DescribeAccessRegionsWithContext(context.Background(), request)
}

// DescribeAccessRegions
// Query the access region list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAccessRegionsWithContext(ctx context.Context, request *DescribeAccessRegionsRequest) (response *DescribeAccessRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DescribeAccessRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadActiveDeviceCountRequest() (request *DownloadActiveDeviceCountRequest) {
    request = &DownloadActiveDeviceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DownloadActiveDeviceCount")
    
    
    return
}

func NewDownloadActiveDeviceCountResponse() (response *DownloadActiveDeviceCountResponse) {
    response = &DownloadActiveDeviceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadActiveDeviceCount
// Download the number of active devices statistics
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DownloadActiveDeviceCount(request *DownloadActiveDeviceCountRequest) (response *DownloadActiveDeviceCountResponse, err error) {
    return c.DownloadActiveDeviceCountWithContext(context.Background(), request)
}

// DownloadActiveDeviceCount
// Download the number of active devices statistics
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DownloadActiveDeviceCountWithContext(ctx context.Context, request *DownloadActiveDeviceCountRequest) (response *DownloadActiveDeviceCountResponse, err error) {
    if request == nil {
        request = NewDownloadActiveDeviceCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "DownloadActiveDeviceCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadActiveDeviceCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadActiveDeviceCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetActiveDeviceCountRequest() (request *GetActiveDeviceCountRequest) {
    request = &GetActiveDeviceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetActiveDeviceCount")
    
    
    return
}

func NewGetActiveDeviceCountResponse() (response *GetActiveDeviceCountResponse) {
    response = &GetActiveDeviceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetActiveDeviceCount
// Number of active devices statistics
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetActiveDeviceCount(request *GetActiveDeviceCountRequest) (response *GetActiveDeviceCountResponse, err error) {
    return c.GetActiveDeviceCountWithContext(context.Background(), request)
}

// GetActiveDeviceCount
// Number of active devices statistics
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetActiveDeviceCountWithContext(ctx context.Context, request *GetActiveDeviceCountRequest) (response *GetActiveDeviceCountResponse, err error) {
    if request == nil {
        request = NewGetActiveDeviceCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetActiveDeviceCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetActiveDeviceCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetActiveDeviceCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetApplicationRequest() (request *GetApplicationRequest) {
    request = &GetApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetApplication")
    
    
    return
}

func NewGetApplicationResponse() (response *GetApplicationResponse) {
    response = &GetApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetApplication
// This API is used to query applications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetApplication(request *GetApplicationRequest) (response *GetApplicationResponse, err error) {
    return c.GetApplicationWithContext(context.Background(), request)
}

// GetApplication
// This API is used to query applications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetApplicationWithContext(ctx context.Context, request *GetApplicationRequest) (response *GetApplicationResponse, err error) {
    if request == nil {
        request = NewGetApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewGetDestIPByNameRequest() (request *GetDestIPByNameRequest) {
    request = &GetDestIPByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDestIPByName")
    
    
    return
}

func NewGetDestIPByNameResponse() (response *GetDestIPByNameResponse) {
    response = &GetDestIPByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDestIPByName
// Statistics of a single device accessing the target IP address
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDestIPByName(request *GetDestIPByNameRequest) (response *GetDestIPByNameResponse, err error) {
    return c.GetDestIPByNameWithContext(context.Background(), request)
}

// GetDestIPByName
// Statistics of a single device accessing the target IP address
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDestIPByNameWithContext(ctx context.Context, request *GetDestIPByNameRequest) (response *GetDestIPByNameResponse, err error) {
    if request == nil {
        request = NewGetDestIPByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetDestIPByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDestIPByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDestIPByNameResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceRequest() (request *GetDeviceRequest) {
    request = &GetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevice")
    
    
    return
}

func NewGetDeviceResponse() (response *GetDeviceResponse) {
    response = &GetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDevice
// This API is used to search device details by specified device ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevice(request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    return c.GetDeviceWithContext(context.Background(), request)
}

// GetDevice
// This API is used to search device details by specified device ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDeviceWithContext(ctx context.Context, request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    if request == nil {
        request = NewGetDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicePayModeRequest() (request *GetDevicePayModeRequest) {
    request = &GetDevicePayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevicePayMode")
    
    
    return
}

func NewGetDevicePayModeResponse() (response *GetDevicePayModeResponse) {
    response = &GetDevicePayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDevicePayMode
// This API is used to obtain the payment mode of a device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
func (c *Client) GetDevicePayMode(request *GetDevicePayModeRequest) (response *GetDevicePayModeResponse, err error) {
    return c.GetDevicePayModeWithContext(context.Background(), request)
}

// GetDevicePayMode
// This API is used to obtain the payment mode of a device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
func (c *Client) GetDevicePayModeWithContext(ctx context.Context, request *GetDevicePayModeRequest) (response *GetDevicePayModeResponse, err error) {
    if request == nil {
        request = NewGetDevicePayModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetDevicePayMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevicePayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDevicePayModeResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicesRequest() (request *GetDevicesRequest) {
    request = &GetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevices")
    
    
    return
}

func NewGetDevicesResponse() (response *GetDevicesResponse) {
    response = &GetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDevices
// This API is used to get device information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevices(request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    return c.GetDevicesWithContext(context.Background(), request)
}

// GetDevices
// This API is used to get device information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevicesWithContext(ctx context.Context, request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    if request == nil {
        request = NewGetDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowAlarmInfoRequest() (request *GetFlowAlarmInfoRequest) {
    request = &GetFlowAlarmInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowAlarmInfo")
    
    
    return
}

func NewGetFlowAlarmInfoResponse() (response *GetFlowAlarmInfoResponse) {
    response = &GetFlowAlarmInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowAlarmInfo
// This API is used to query user traffic alarm settings based on AppId, including threshold, callback url and key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) GetFlowAlarmInfo(request *GetFlowAlarmInfoRequest) (response *GetFlowAlarmInfoResponse, err error) {
    return c.GetFlowAlarmInfoWithContext(context.Background(), request)
}

// GetFlowAlarmInfo
// This API is used to query user traffic alarm settings based on AppId, including threshold, callback url and key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) GetFlowAlarmInfoWithContext(ctx context.Context, request *GetFlowAlarmInfoRequest) (response *GetFlowAlarmInfoResponse, err error) {
    if request == nil {
        request = NewGetFlowAlarmInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowAlarmInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowAlarmInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowAlarmInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowPackagesRequest() (request *GetFlowPackagesRequest) {
    request = &GetFlowPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowPackages")
    
    
    return
}

func NewGetFlowPackagesResponse() (response *GetFlowPackagesResponse) {
    response = &GetFlowPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowPackages
// Retrieve the data transfer plan list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFlowPackages(request *GetFlowPackagesRequest) (response *GetFlowPackagesResponse, err error) {
    return c.GetFlowPackagesWithContext(context.Background(), request)
}

// GetFlowPackages
// Retrieve the data transfer plan list
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFlowPackagesWithContext(ctx context.Context, request *GetFlowPackagesRequest) (response *GetFlowPackagesResponse, err error) {
    if request == nil {
        request = NewGetFlowPackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowPackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticRequest() (request *GetFlowStatisticRequest) {
    request = &GetFlowStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatistic")
    
    
    return
}

func NewGetFlowStatisticResponse() (response *GetFlowStatisticResponse) {
    response = &GetFlowStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatistic
// Retrieve traffic usage data for a specified device Id at a specified time point.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatistic(request *GetFlowStatisticRequest) (response *GetFlowStatisticResponse, err error) {
    return c.GetFlowStatisticWithContext(context.Background(), request)
}

// GetFlowStatistic
// Retrieve traffic usage data for a specified device Id at a specified time point.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticWithContext(ctx context.Context, request *GetFlowStatisticRequest) (response *GetFlowStatisticResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticByGroupRequest() (request *GetFlowStatisticByGroupRequest) {
    request = &GetFlowStatisticByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatisticByGroup")
    
    
    return
}

func NewGetFlowStatisticByGroupResponse() (response *GetFlowStatisticByGroupResponse) {
    response = &GetFlowStatisticByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatisticByGroup
// Retrieve traffic usage data for the specified group and time period
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByGroup(request *GetFlowStatisticByGroupRequest) (response *GetFlowStatisticByGroupResponse, err error) {
    return c.GetFlowStatisticByGroupWithContext(context.Background(), request)
}

// GetFlowStatisticByGroup
// Retrieve traffic usage data for the specified group and time period
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByGroupWithContext(ctx context.Context, request *GetFlowStatisticByGroupRequest) (response *GetFlowStatisticByGroupResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticByGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowStatisticByGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatisticByGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticByNameRequest() (request *GetFlowStatisticByNameRequest) {
    request = &GetFlowStatisticByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatisticByName")
    
    
    return
}

func NewGetFlowStatisticByNameResponse() (response *GetFlowStatisticByNameResponse) {
    response = &GetFlowStatisticByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatisticByName
// Retrieve traffic usage data for a specified device Id at a specified time point.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByName(request *GetFlowStatisticByNameRequest) (response *GetFlowStatisticByNameResponse, err error) {
    return c.GetFlowStatisticByNameWithContext(context.Background(), request)
}

// GetFlowStatisticByName
// Retrieve traffic usage data for a specified device Id at a specified time point.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByNameWithContext(ctx context.Context, request *GetFlowStatisticByNameRequest) (response *GetFlowStatisticByNameResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowStatisticByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatisticByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticByNameResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticByRegionRequest() (request *GetFlowStatisticByRegionRequest) {
    request = &GetFlowStatisticByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatisticByRegion")
    
    
    return
}

func NewGetFlowStatisticByRegionResponse() (response *GetFlowStatisticByRegionResponse) {
    response = &GetFlowStatisticByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatisticByRegion
// Retrieve traffic usage data for the specified region and time point
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByRegion(request *GetFlowStatisticByRegionRequest) (response *GetFlowStatisticByRegionResponse, err error) {
    return c.GetFlowStatisticByRegionWithContext(context.Background(), request)
}

// GetFlowStatisticByRegion
// Retrieve traffic usage data for the specified region and time point
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByRegionWithContext(ctx context.Context, request *GetFlowStatisticByRegionRequest) (response *GetFlowStatisticByRegionResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticByRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetFlowStatisticByRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatisticByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupDetailRequest() (request *GetGroupDetailRequest) {
    request = &GetGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetGroupDetail")
    
    
    return
}

func NewGetGroupDetailResponse() (response *GetGroupDetailResponse) {
    response = &GetGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroupDetail
// View group details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupDetail(request *GetGroupDetailRequest) (response *GetGroupDetailResponse, err error) {
    return c.GetGroupDetailWithContext(context.Background(), request)
}

// GetGroupDetail
// View group details
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupDetailWithContext(ctx context.Context, request *GetGroupDetailRequest) (response *GetGroupDetailResponse, err error) {
    if request == nil {
        request = NewGetGroupDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetGroupDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetGroupList")
    
    
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroupList
// This API is used to obtain a group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    return c.GetGroupListWithContext(context.Background(), request)
}

// GetGroupList
// This API is used to obtain a group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupListWithContext(ctx context.Context, request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetHardwareInfoRequest() (request *GetHardwareInfoRequest) {
    request = &GetHardwareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetHardwareInfo")
    
    
    return
}

func NewGetHardwareInfoResponse() (response *GetHardwareInfoResponse) {
    response = &GetHardwareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetHardwareInfo
// This API is used to get hardware device information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
func (c *Client) GetHardwareInfo(request *GetHardwareInfoRequest) (response *GetHardwareInfoResponse, err error) {
    return c.GetHardwareInfoWithContext(context.Background(), request)
}

// GetHardwareInfo
// This API is used to get hardware device information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
func (c *Client) GetHardwareInfoWithContext(ctx context.Context, request *GetHardwareInfoRequest) (response *GetHardwareInfoResponse, err error) {
    if request == nil {
        request = NewGetHardwareInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetHardwareInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetHardwareInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetHardwareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetHardwareListRequest() (request *GetHardwareListRequest) {
    request = &GetHardwareListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetHardwareList")
    
    
    return
}

func NewGetHardwareListResponse() (response *GetHardwareListResponse) {
    response = &GetHardwareListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetHardwareList
// This API is used to get the hardware list of the manufacturer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetHardwareList(request *GetHardwareListRequest) (response *GetHardwareListResponse, err error) {
    return c.GetHardwareListWithContext(context.Background(), request)
}

// GetHardwareList
// This API is used to get the hardware list of the manufacturer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetHardwareListWithContext(ctx context.Context, request *GetHardwareListRequest) (response *GetHardwareListResponse, err error) {
    if request == nil {
        request = NewGetHardwareListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetHardwareList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetHardwareList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetHardwareListResponse()
    err = c.Send(request, response)
    return
}

func NewGetL3ConnListRequest() (request *GetL3ConnListRequest) {
    request = &GetL3ConnListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetL3ConnList")
    
    
    return
}

func NewGetL3ConnListResponse() (response *GetL3ConnListResponse) {
    response = &GetL3ConnListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetL3ConnList
// Retrieve the list of interconnection rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetL3ConnList(request *GetL3ConnListRequest) (response *GetL3ConnListResponse, err error) {
    return c.GetL3ConnListWithContext(context.Background(), request)
}

// GetL3ConnList
// Retrieve the list of interconnection rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetL3ConnListWithContext(ctx context.Context, request *GetL3ConnListRequest) (response *GetL3ConnListResponse, err error) {
    if request == nil {
        request = NewGetL3ConnListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetL3ConnList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetL3ConnList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetL3ConnListResponse()
    err = c.Send(request, response)
    return
}

func NewGetMonitorDataByNameRequest() (request *GetMonitorDataByNameRequest) {
    request = &GetMonitorDataByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetMonitorDataByName")
    
    
    return
}

func NewGetMonitorDataByNameResponse() (response *GetMonitorDataByNameResponse) {
    response = &GetMonitorDataByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMonitorDataByName
// This API is used to obtain the download file URL for all monitoring metrics of a single device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_COSREQUESTERROR = "InternalError.CosRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMonitorDataByName(request *GetMonitorDataByNameRequest) (response *GetMonitorDataByNameResponse, err error) {
    return c.GetMonitorDataByNameWithContext(context.Background(), request)
}

// GetMonitorDataByName
// This API is used to obtain the download file URL for all monitoring metrics of a single device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_COSREQUESTERROR = "InternalError.CosRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMonitorDataByNameWithContext(ctx context.Context, request *GetMonitorDataByNameRequest) (response *GetMonitorDataByNameResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetMonitorDataByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMonitorDataByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMonitorDataByNameResponse()
    err = c.Send(request, response)
    return
}

func NewGetMultiFlowStatisticRequest() (request *GetMultiFlowStatisticRequest) {
    request = &GetMultiFlowStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetMultiFlowStatistic")
    
    
    return
}

func NewGetMultiFlowStatisticResponse() (response *GetMultiFlowStatisticResponse) {
    response = &GetMultiFlowStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMultiFlowStatistic
// Obtain batch device traffic statistics curves
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMultiFlowStatistic(request *GetMultiFlowStatisticRequest) (response *GetMultiFlowStatisticResponse, err error) {
    return c.GetMultiFlowStatisticWithContext(context.Background(), request)
}

// GetMultiFlowStatistic
// Obtain batch device traffic statistics curves
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMultiFlowStatisticWithContext(ctx context.Context, request *GetMultiFlowStatisticRequest) (response *GetMultiFlowStatisticResponse, err error) {
    if request == nil {
        request = NewGetMultiFlowStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetMultiFlowStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMultiFlowStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMultiFlowStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewGetNetMonitorRequest() (request *GetNetMonitorRequest) {
    request = &GetNetMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetNetMonitor")
    
    
    return
}

func NewGetNetMonitorResponse() (response *GetNetMonitorResponse) {
    response = &GetNetMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetNetMonitor
// This API is used to obtain real-time traffic statistics per device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMEFUTURE = "InvalidParameterValue.TimeFuture"
//  INVALIDPARAMETERVALUE_TIMESPANEXCEEDED = "InvalidParameterValue.TimeSpanExceeded"
//  INVALIDPARAMETERVALUE_TIMETOOEARLY = "InvalidParameterValue.TimeTooEarly"
func (c *Client) GetNetMonitor(request *GetNetMonitorRequest) (response *GetNetMonitorResponse, err error) {
    return c.GetNetMonitorWithContext(context.Background(), request)
}

// GetNetMonitor
// This API is used to obtain real-time traffic statistics per device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMEFUTURE = "InvalidParameterValue.TimeFuture"
//  INVALIDPARAMETERVALUE_TIMESPANEXCEEDED = "InvalidParameterValue.TimeSpanExceeded"
//  INVALIDPARAMETERVALUE_TIMETOOEARLY = "InvalidParameterValue.TimeTooEarly"
func (c *Client) GetNetMonitorWithContext(ctx context.Context, request *GetNetMonitorRequest) (response *GetNetMonitorResponse, err error) {
    if request == nil {
        request = NewGetNetMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetNetMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetNetMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetNetMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewGetNetMonitorByNameRequest() (request *GetNetMonitorByNameRequest) {
    request = &GetNetMonitorByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetNetMonitorByName")
    
    
    return
}

func NewGetNetMonitorByNameResponse() (response *GetNetMonitorByNameResponse) {
    response = &GetNetMonitorByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetNetMonitorByName
// This API is used to obtain real-time traffic statistics per device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMEFUTURE = "InvalidParameterValue.TimeFuture"
//  INVALIDPARAMETERVALUE_TIMESPANEXCEEDED = "InvalidParameterValue.TimeSpanExceeded"
//  INVALIDPARAMETERVALUE_TIMETOOEARLY = "InvalidParameterValue.TimeTooEarly"
func (c *Client) GetNetMonitorByName(request *GetNetMonitorByNameRequest) (response *GetNetMonitorByNameResponse, err error) {
    return c.GetNetMonitorByNameWithContext(context.Background(), request)
}

// GetNetMonitorByName
// This API is used to obtain real-time traffic statistics per device.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMEFUTURE = "InvalidParameterValue.TimeFuture"
//  INVALIDPARAMETERVALUE_TIMESPANEXCEEDED = "InvalidParameterValue.TimeSpanExceeded"
//  INVALIDPARAMETERVALUE_TIMETOOEARLY = "InvalidParameterValue.TimeTooEarly"
func (c *Client) GetNetMonitorByNameWithContext(ctx context.Context, request *GetNetMonitorByNameRequest) (response *GetNetMonitorByNameResponse, err error) {
    if request == nil {
        request = NewGetNetMonitorByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetNetMonitorByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetNetMonitorByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetNetMonitorByNameResponse()
    err = c.Send(request, response)
    return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
    request = &GetPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetPublicKey")
    
    
    return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
    response = &GetPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetPublicKey
// This API is used to access the public key for signature verification.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    return c.GetPublicKeyWithContext(context.Background(), request)
}

// GetPublicKey
// This API is used to access the public key for signature verification.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetPublicKeyWithContext(ctx context.Context, request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetPublicKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticDataRequest() (request *GetStatisticDataRequest) {
    request = &GetStatisticDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetStatisticData")
    
    
    return
}

func NewGetStatisticDataResponse() (response *GetStatisticDataResponse) {
    response = &GetStatisticDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetStatisticData
// Download traffic data on the usage statistics page
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticData(request *GetStatisticDataRequest) (response *GetStatisticDataResponse, err error) {
    return c.GetStatisticDataWithContext(context.Background(), request)
}

// GetStatisticData
// Download traffic data on the usage statistics page
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticDataWithContext(ctx context.Context, request *GetStatisticDataRequest) (response *GetStatisticDataResponse, err error) {
    if request == nil {
        request = NewGetStatisticDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetStatisticData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetStatisticData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetStatisticDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticDataByNameRequest() (request *GetStatisticDataByNameRequest) {
    request = &GetStatisticDataByNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetStatisticDataByName")
    
    
    return
}

func NewGetStatisticDataByNameResponse() (response *GetStatisticDataByNameResponse) {
    response = &GetStatisticDataByNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetStatisticDataByName
// Download traffic data on the usage statistics page
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticDataByName(request *GetStatisticDataByNameRequest) (response *GetStatisticDataByNameResponse, err error) {
    return c.GetStatisticDataByNameWithContext(context.Background(), request)
}

// GetStatisticDataByName
// Download traffic data on the usage statistics page
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticDataByNameWithContext(ctx context.Context, request *GetStatisticDataByNameRequest) (response *GetStatisticDataByNameResponse, err error) {
    if request == nil {
        request = NewGetStatisticDataByNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetStatisticDataByName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetStatisticDataByName require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetStatisticDataByNameResponse()
    err = c.Send(request, response)
    return
}

func NewGetVendorHardwareRequest() (request *GetVendorHardwareRequest) {
    request = &GetVendorHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetVendorHardware")
    
    
    return
}

func NewGetVendorHardwareResponse() (response *GetVendorHardwareResponse) {
    response = &GetVendorHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetVendorHardware
// This API is used to get the hardware device list of the manufacturer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetVendorHardware(request *GetVendorHardwareRequest) (response *GetVendorHardwareResponse, err error) {
    return c.GetVendorHardwareWithContext(context.Background(), request)
}

// GetVendorHardware
// This API is used to get the hardware device list of the manufacturer.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetVendorHardwareWithContext(ctx context.Context, request *GetVendorHardwareRequest) (response *GetVendorHardwareResponse, err error) {
    if request == nil {
        request = NewGetVendorHardwareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GetVendorHardware")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVendorHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVendorHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewGroupAddDeviceRequest() (request *GroupAddDeviceRequest) {
    request = &GroupAddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GroupAddDevice")
    
    
    return
}

func NewGroupAddDeviceResponse() (response *GroupAddDeviceResponse) {
    response = &GroupAddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupAddDevice
// Add device to already exist group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupAddDevice(request *GroupAddDeviceRequest) (response *GroupAddDeviceResponse, err error) {
    return c.GroupAddDeviceWithContext(context.Background(), request)
}

// GroupAddDevice
// Add device to already exist group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupAddDeviceWithContext(ctx context.Context, request *GroupAddDeviceRequest) (response *GroupAddDeviceResponse, err error) {
    if request == nil {
        request = NewGroupAddDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GroupAddDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupAddDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewGroupDeleteDeviceRequest() (request *GroupDeleteDeviceRequest) {
    request = &GroupDeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GroupDeleteDevice")
    
    
    return
}

func NewGroupDeleteDeviceResponse() (response *GroupDeleteDeviceResponse) {
    response = &GroupDeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupDeleteDevice
// Delete devices in the group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupDeleteDevice(request *GroupDeleteDeviceRequest) (response *GroupDeleteDeviceResponse, err error) {
    return c.GroupDeleteDeviceWithContext(context.Background(), request)
}

// GroupDeleteDevice
// Delete devices in the group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupDeleteDeviceWithContext(ctx context.Context, request *GroupDeleteDeviceRequest) (response *GroupDeleteDeviceResponse, err error) {
    if request == nil {
        request = NewGroupDeleteDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "GroupDeleteDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupDeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceAccessRegionsRequest() (request *ModifyDeviceAccessRegionsRequest) {
    request = &ModifyDeviceAccessRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ModifyDeviceAccessRegions")
    
    
    return
}

func NewModifyDeviceAccessRegionsResponse() (response *ModifyDeviceAccessRegionsResponse) {
    response = &ModifyDeviceAccessRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDeviceAccessRegions
// This API is used to modify device connectivity regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDeviceAccessRegions(request *ModifyDeviceAccessRegionsRequest) (response *ModifyDeviceAccessRegionsResponse, err error) {
    return c.ModifyDeviceAccessRegionsWithContext(context.Background(), request)
}

// ModifyDeviceAccessRegions
// This API is used to modify device connectivity regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDeviceAccessRegionsWithContext(ctx context.Context, request *ModifyDeviceAccessRegionsRequest) (response *ModifyDeviceAccessRegionsResponse, err error) {
    if request == nil {
        request = NewModifyDeviceAccessRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "ModifyDeviceAccessRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceAccessRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceAccessRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPackageRenewFlagRequest() (request *ModifyPackageRenewFlagRequest) {
    request = &ModifyPackageRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ModifyPackageRenewFlag")
    
    
    return
}

func NewModifyPackageRenewFlagResponse() (response *ModifyPackageRenewFlagResponse) {
    response = &ModifyPackageRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPackageRenewFlag
// Auto renewal of data transfer plans can be enabled or disabled, unaffected by ongoing effective plans in the current cycle.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
//  OPERATIONDENIED_MODIFIEDORRENEWED = "OperationDenied.ModifiedOrRenewed"
//  OPERATIONDENIED_TRUNCFLAGON = "OperationDenied.TruncFlagOn"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPackageRenewFlag(request *ModifyPackageRenewFlagRequest) (response *ModifyPackageRenewFlagResponse, err error) {
    return c.ModifyPackageRenewFlagWithContext(context.Background(), request)
}

// ModifyPackageRenewFlag
// Auto renewal of data transfer plans can be enabled or disabled, unaffected by ongoing effective plans in the current cycle.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
//  OPERATIONDENIED_MODIFIEDORRENEWED = "OperationDenied.ModifiedOrRenewed"
//  OPERATIONDENIED_TRUNCFLAGON = "OperationDenied.TruncFlagOn"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPackageRenewFlagWithContext(ctx context.Context, request *ModifyPackageRenewFlagRequest) (response *ModifyPackageRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyPackageRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "ModifyPackageRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPackageRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPackageRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOrderFlowPackageRequest() (request *OrderFlowPackageRequest) {
    request = &OrderFlowPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "OrderFlowPackage")
    
    
    return
}

func NewOrderFlowPackageResponse() (response *OrderFlowPackageResponse) {
    response = &OrderFlowPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OrderFlowPackage
// Purchase a Prepaid Traffic Package
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNOPENEDLIVESERVICE = "UnauthorizedOperation.UnopenedLiveService"
func (c *Client) OrderFlowPackage(request *OrderFlowPackageRequest) (response *OrderFlowPackageResponse, err error) {
    return c.OrderFlowPackageWithContext(context.Background(), request)
}

// OrderFlowPackage
// Purchase a Prepaid Traffic Package
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNOPENEDLIVESERVICE = "UnauthorizedOperation.UnopenedLiveService"
func (c *Client) OrderFlowPackageWithContext(ctx context.Context, request *OrderFlowPackageRequest) (response *OrderFlowPackageResponse, err error) {
    if request == nil {
        request = NewOrderFlowPackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "OrderFlowPackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OrderFlowPackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewOrderFlowPackageResponse()
    err = c.Send(request, response)
    return
}

func NewOrderPerLicenseRequest() (request *OrderPerLicenseRequest) {
    request = &OrderPerLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "OrderPerLicense")
    
    
    return
}

func NewOrderPerLicenseResponse() (response *OrderPerLicenseResponse) {
    response = &OrderPerLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OrderPerLicense
// Purchase a single-use License
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) OrderPerLicense(request *OrderPerLicenseRequest) (response *OrderPerLicenseResponse, err error) {
    return c.OrderPerLicenseWithContext(context.Background(), request)
}

// OrderPerLicense
// Purchase a single-use License
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) OrderPerLicenseWithContext(ctx context.Context, request *OrderPerLicenseRequest) (response *OrderPerLicenseResponse, err error) {
    if request == nil {
        request = NewOrderPerLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "OrderPerLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OrderPerLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewOrderPerLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewReportOrderRequest() (request *ReportOrderRequest) {
    request = &ReportOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ReportOrder")
    
    
    return
}

func NewReportOrderResponse() (response *ReportOrderResponse) {
    response = &ReportOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportOrder
// Users report custom order information, and the Multiple Network Acceleration service saves the information related to.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ReportOrder(request *ReportOrderRequest) (response *ReportOrderResponse, err error) {
    return c.ReportOrderWithContext(context.Background(), request)
}

// ReportOrder
// Users report custom order information, and the Multiple Network Acceleration service saves the information related to.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ReportOrderWithContext(ctx context.Context, request *ReportOrderRequest) (response *ReportOrderResponse, err error) {
    if request == nil {
        request = NewReportOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "ReportOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportOrderResponse()
    err = c.Send(request, response)
    return
}

func NewSetNotifyUrlRequest() (request *SetNotifyUrlRequest) {
    request = &SetNotifyUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "SetNotifyUrl")
    
    
    return
}

func NewSetNotifyUrlResponse() (response *SetNotifyUrlResponse) {
    response = &SetNotifyUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNotifyUrl
// This API is used to set user traffic alarm information. Use this API setting to configure the data transfer plan alarm threshold as well as the callback url and key when an alarm is generated.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) SetNotifyUrl(request *SetNotifyUrlRequest) (response *SetNotifyUrlResponse, err error) {
    return c.SetNotifyUrlWithContext(context.Background(), request)
}

// SetNotifyUrl
// This API is used to set user traffic alarm information. Use this API setting to configure the data transfer plan alarm threshold as well as the callback url and key when an alarm is generated.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) SetNotifyUrlWithContext(ctx context.Context, request *SetNotifyUrlRequest) (response *SetNotifyUrlResponse, err error) {
    if request == nil {
        request = NewSetNotifyUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "SetNotifyUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNotifyUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNotifyUrlResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApplicationInfoRequest() (request *UpdateApplicationInfoRequest) {
    request = &UpdateApplicationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateApplicationInfo")
    
    
    return
}

func NewUpdateApplicationInfoResponse() (response *UpdateApplicationInfoResponse) {
    response = &UpdateApplicationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApplicationInfo
// Update application information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) UpdateApplicationInfo(request *UpdateApplicationInfoRequest) (response *UpdateApplicationInfoResponse, err error) {
    return c.UpdateApplicationInfoWithContext(context.Background(), request)
}

// UpdateApplicationInfo
// Update application information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) UpdateApplicationInfoWithContext(ctx context.Context, request *UpdateApplicationInfoRequest) (response *UpdateApplicationInfoResponse, err error) {
    if request == nil {
        request = NewUpdateApplicationInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateApplicationInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApplicationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApplicationKeyRequest() (request *UpdateApplicationKeyRequest) {
    request = &UpdateApplicationKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateApplicationKey")
    
    
    return
}

func NewUpdateApplicationKeyResponse() (response *UpdateApplicationKeyResponse) {
    response = &UpdateApplicationKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApplicationKey
// Update application key
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) UpdateApplicationKey(request *UpdateApplicationKeyRequest) (response *UpdateApplicationKeyResponse, err error) {
    return c.UpdateApplicationKeyWithContext(context.Background(), request)
}

// UpdateApplicationKey
// Update application key
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) UpdateApplicationKeyWithContext(ctx context.Context, request *UpdateApplicationKeyRequest) (response *UpdateApplicationKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApplicationKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateApplicationKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApplicationKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApplicationKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceRequest() (request *UpdateDeviceRequest) {
    request = &UpdateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateDevice")
    
    
    return
}

func NewUpdateDeviceResponse() (response *UpdateDeviceResponse) {
    response = &UpdateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDevice
// Update device information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateDevice(request *UpdateDeviceRequest) (response *UpdateDeviceResponse, err error) {
    return c.UpdateDeviceWithContext(context.Background(), request)
}

// UpdateDevice
// Update device information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateDeviceWithContext(ctx context.Context, request *UpdateDeviceRequest) (response *UpdateDeviceResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateGroup")
    
    
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGroup
// Update group remark
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    return c.UpdateGroupWithContext(context.Background(), request)
}

// UpdateGroup
// Update group remark
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHardwareRequest() (request *UpdateHardwareRequest) {
    request = &UpdateHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateHardware")
    
    
    return
}

func NewUpdateHardwareResponse() (response *UpdateHardwareResponse) {
    response = &UpdateHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateHardware
// Refresh hardware info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) UpdateHardware(request *UpdateHardwareRequest) (response *UpdateHardwareResponse, err error) {
    return c.UpdateHardwareWithContext(context.Background(), request)
}

// UpdateHardware
// Refresh hardware info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) UpdateHardwareWithContext(ctx context.Context, request *UpdateHardwareRequest) (response *UpdateHardwareResponse, err error) {
    if request == nil {
        request = NewUpdateHardwareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateHardware")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3CidrRequest() (request *UpdateL3CidrRequest) {
    request = &UpdateL3CidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Cidr")
    
    
    return
}

func NewUpdateL3CidrResponse() (response *UpdateL3CidrResponse) {
    response = &UpdateL3CidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Cidr
// Update the interconnection rule CIDR
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Cidr(request *UpdateL3CidrRequest) (response *UpdateL3CidrResponse, err error) {
    return c.UpdateL3CidrWithContext(context.Background(), request)
}

// UpdateL3Cidr
// Update the interconnection rule CIDR
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3CidrWithContext(ctx context.Context, request *UpdateL3CidrRequest) (response *UpdateL3CidrResponse, err error) {
    if request == nil {
        request = NewUpdateL3CidrRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateL3Cidr")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Cidr require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3CidrResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3ConnRequest() (request *UpdateL3ConnRequest) {
    request = &UpdateL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Conn")
    
    
    return
}

func NewUpdateL3ConnResponse() (response *UpdateL3ConnResponse) {
    response = &UpdateL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Conn
// Update the interconnection rule remark
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Conn(request *UpdateL3ConnRequest) (response *UpdateL3ConnResponse, err error) {
    return c.UpdateL3ConnWithContext(context.Background(), request)
}

// UpdateL3Conn
// Update the interconnection rule remark
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3ConnWithContext(ctx context.Context, request *UpdateL3ConnRequest) (response *UpdateL3ConnResponse, err error) {
    if request == nil {
        request = NewUpdateL3ConnRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateL3Conn")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3ConnResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3SwitchRequest() (request *UpdateL3SwitchRequest) {
    request = &UpdateL3SwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Switch")
    
    
    return
}

func NewUpdateL3SwitchResponse() (response *UpdateL3SwitchResponse) {
    response = &UpdateL3SwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Switch
// Update the interconnection rule switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Switch(request *UpdateL3SwitchRequest) (response *UpdateL3SwitchResponse, err error) {
    return c.UpdateL3SwitchWithContext(context.Background(), request)
}

// UpdateL3Switch
// Update the interconnection rule switch
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3SwitchWithContext(ctx context.Context, request *UpdateL3SwitchRequest) (response *UpdateL3SwitchResponse, err error) {
    if request == nil {
        request = NewUpdateL3SwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mna", APIVersion, "UpdateL3Switch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Switch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3SwitchResponse()
    err = c.Send(request, response)
    return
}
