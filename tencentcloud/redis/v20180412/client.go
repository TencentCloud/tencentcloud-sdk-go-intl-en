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

package v20180412

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-12"

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


func NewAllocateWanAddressRequest() (request *AllocateWanAddressRequest) {
    request = &AllocateWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "AllocateWanAddress")
    
    
    return
}

func NewAllocateWanAddressResponse() (response *AllocateWanAddressResponse) {
    response = &AllocateWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateWanAddress
// This API is used to enable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) AllocateWanAddress(request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    return c.AllocateWanAddressWithContext(context.Background(), request)
}

// AllocateWanAddress
// This API is used to enable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) AllocateWanAddressWithContext(ctx context.Context, request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    if request == nil {
        request = NewAllocateWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateWanAddressResponse()
    err = c.Send(request, response)
    return
}

func NewChangeReplicaToMasterRequest() (request *ChangeReplicaToMasterRequest) {
    request = &ChangeReplicaToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ChangeReplicaToMaster")
    
    
    return
}

func NewChangeReplicaToMasterResponse() (response *ChangeReplicaToMasterResponse) {
    response = &ChangeReplicaToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeReplicaToMaster
// This API is used to promote a replica node group of a multi-AZ deployed instance to master node group or a replica node of a single-AZ deployed instance to master node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMaster(request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    return c.ChangeReplicaToMasterWithContext(context.Background(), request)
}

// ChangeReplicaToMaster
// This API is used to promote a replica node group of a multi-AZ deployed instance to master node group or a replica node of a single-AZ deployed instance to master node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMasterWithContext(ctx context.Context, request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    if request == nil {
        request = NewChangeReplicaToMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeReplicaToMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeReplicaToMasterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAccountRequest() (request *DescribeInstanceAccountRequest) {
    request = &DescribeInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceAccount")
    
    
    return
}

func NewDescribeInstanceAccountResponse() (response *DescribeInstanceAccountResponse) {
    response = &DescribeInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceAccount
// This API is used to view instance sub-account information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccount(request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    return c.DescribeInstanceAccountWithContext(context.Background(), request)
}

// DescribeInstanceAccount
// This API is used to view instance sub-account information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccountWithContext(ctx context.Context, request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyRequest() (request *DescribeInstanceMonitorBigKeyRequest) {
    request = &DescribeInstanceMonitorBigKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKey")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyResponse() (response *DescribeInstanceMonitorBigKeyResponse) {
    response = &DescribeInstanceMonitorBigKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKey
// This API is used to query the big key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKey(request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKey
// This API is used to query the big key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistRequest() (request *DescribeInstanceMonitorBigKeySizeDistRequest) {
    request = &DescribeInstanceMonitorBigKeySizeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeySizeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistResponse() (response *DescribeInstanceMonitorBigKeySizeDistResponse) {
    response = &DescribeInstanceMonitorBigKeySizeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKeySizeDist
// This API is used to query the big key size distribution of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDist(request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeySizeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeySizeDist
// This API is used to query the big key size distribution of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeySizeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeySizeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeySizeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistRequest() (request *DescribeInstanceMonitorBigKeyTypeDistRequest) {
    request = &DescribeInstanceMonitorBigKeyTypeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeyTypeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistResponse() (response *DescribeInstanceMonitorBigKeyTypeDistResponse) {
    response = &DescribeInstanceMonitorBigKeyTypeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKeyTypeDist
// This API is used to query the big key type distribution of an instance
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDist(request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyTypeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeyTypeDist
// This API is used to query the big key type distribution of an instance
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyTypeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeyTypeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyTypeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorHotKeyRequest() (request *DescribeInstanceMonitorHotKeyRequest) {
    request = &DescribeInstanceMonitorHotKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorHotKey")
    
    
    return
}

func NewDescribeInstanceMonitorHotKeyResponse() (response *DescribeInstanceMonitorHotKeyResponse) {
    response = &DescribeInstanceMonitorHotKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorHotKey
// This API is used to query the hot key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKey(request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    return c.DescribeInstanceMonitorHotKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorHotKey
// This API is used to query the hot key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorHotKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorHotKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorHotKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorSIPRequest() (request *DescribeInstanceMonitorSIPRequest) {
    request = &DescribeInstanceMonitorSIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorSIP")
    
    
    return
}

func NewDescribeInstanceMonitorSIPResponse() (response *DescribeInstanceMonitorSIPResponse) {
    response = &DescribeInstanceMonitorSIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorSIP
// This API is used to query the access source information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIP(request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    return c.DescribeInstanceMonitorSIPWithContext(context.Background(), request)
}

// DescribeInstanceMonitorSIP
// This API is used to query the access source information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIPWithContext(ctx context.Context, request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorSIPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorSIP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorSIPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTookDistRequest() (request *DescribeInstanceMonitorTookDistRequest) {
    request = &DescribeInstanceMonitorTookDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTookDist")
    
    
    return
}

func NewDescribeInstanceMonitorTookDistResponse() (response *DescribeInstanceMonitorTookDistResponse) {
    response = &DescribeInstanceMonitorTookDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTookDist
// This API is used to query the distribution of instance access duration.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDist(request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    return c.DescribeInstanceMonitorTookDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTookDist
// This API is used to query the distribution of instance access duration.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDistWithContext(ctx context.Context, request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTookDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTookDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTookDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdRequest() (request *DescribeInstanceMonitorTopNCmdRequest) {
    request = &DescribeInstanceMonitorTopNCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmd")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdResponse() (response *DescribeInstanceMonitorTopNCmdResponse) {
    response = &DescribeInstanceMonitorTopNCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTopNCmd
// This API is used to query an instance access command.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmd(request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmd
// This API is used to query an instance access command.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdTookRequest() (request *DescribeInstanceMonitorTopNCmdTookRequest) {
    request = &DescribeInstanceMonitorTopNCmdTookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmdTook")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdTookResponse() (response *DescribeInstanceMonitorTopNCmdTookResponse) {
    response = &DescribeInstanceMonitorTopNCmdTookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTopNCmdTook
// This API is used to query the instance CPU time.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTook(request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdTookWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmdTook
// This API is used to query the instance CPU time.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTookWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdTookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmdTook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdTookResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodeInfoRequest() (request *DescribeInstanceNodeInfoRequest) {
    request = &DescribeInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceNodeInfo")
    
    
    return
}

func NewDescribeInstanceNodeInfoResponse() (response *DescribeInstanceNodeInfoResponse) {
    response = &DescribeInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceNodeInfo
// This API is used to query instance node information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfo(request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    return c.DescribeInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodeInfo
// This API is used to query instance node information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfoWithContext(ctx context.Context, request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceWindowRequest() (request *DescribeMaintenanceWindowRequest) {
    request = &DescribeMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeMaintenanceWindow")
    
    
    return
}

func NewDescribeMaintenanceWindowResponse() (response *DescribeMaintenanceWindowResponse) {
    response = &DescribeMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaintenanceWindow
// This API is used to query instance maintenance window. The maintenance window specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeMaintenanceWindow(request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    return c.DescribeMaintenanceWindowWithContext(context.Background(), request)
}

// DescribeMaintenanceWindow
// This API is used to query instance maintenance window. The maintenance window specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeMaintenanceWindowWithContext(ctx context.Context, request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
    request = &DescribeProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProductInfo")
    
    
    return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
    response = &DescribeProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductInfo
// This API is used to query the purchasable capacity specifications of Redis instances in the specified AZ and instance type. If you are not in the allowlist for the AZ or instance type, you cannot view the details of the capacity specifications. To apply for the eligibility, please submit a ticket.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    return c.DescribeProductInfoWithContext(context.Background(), request)
}

// DescribeProductInfo
// This API is used to query the purchasable capacity specifications of Redis instances in the specified AZ and instance type. If you are not in the allowlist for the AZ or instance type, you cannot view the details of the capacity specifications. To apply for the eligibility, please submit a ticket.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfoWithContext(ctx context.Context, request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProductInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationGroupRequest() (request *DescribeReplicationGroupRequest) {
    request = &DescribeReplicationGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeReplicationGroup")
    
    
    return
}

func NewDescribeReplicationGroupResponse() (response *DescribeReplicationGroupResponse) {
    response = &DescribeReplicationGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReplicationGroup
// This API is used to query the global replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroup(request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    return c.DescribeReplicationGroupWithContext(context.Background(), request)
}

// DescribeReplicationGroup
// This API is used to query the global replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroupWithContext(ctx context.Context, request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogRequest() (request *DescribeSlowLogRequest) {
    request = &DescribeSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "DescribeSlowLog")
    
    
    return
}

func NewDescribeSlowLogResponse() (response *DescribeSlowLogResponse) {
    response = &DescribeSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowLog
// This API is used to query the slow log.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLog(request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    return c.DescribeSlowLogWithContext(context.Background(), request)
}

// DescribeSlowLog
// This API is used to query the slow log.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLogWithContext(ctx context.Context, request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceCreateInstance")
    
    
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateInstance
// This API is used to query the price for purchasing an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    return c.InquiryPriceCreateInstanceWithContext(context.Background(), request)
}

// InquiryPriceCreateInstance
// This API is used to query the price for purchasing an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstanceWithContext(ctx context.Context, request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeInstanceRequest() (request *InquiryPriceUpgradeInstanceRequest) {
    request = &InquiryPriceUpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceUpgradeInstance")
    
    
    return
}

func NewInquiryPriceUpgradeInstanceResponse() (response *InquiryPriceUpgradeInstanceResponse) {
    response = &InquiryPriceUpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceUpgradeInstance
// This API is used to query the price for scaling an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstance(request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    return c.InquiryPriceUpgradeInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeInstance
// This API is used to query the price for scaling an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModfiyInstancePasswordRequest() (request *ModfiyInstancePasswordRequest) {
    request = &ModfiyInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ModfiyInstancePassword")
    
    
    return
}

func NewModfiyInstancePasswordResponse() (response *ModfiyInstancePasswordResponse) {
    response = &ModfiyInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModfiyInstancePassword
// This API is used to change the Redis password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePassword(request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    return c.ModfiyInstancePasswordWithContext(context.Background(), request)
}

// ModfiyInstancePassword
// This API is used to change the Redis password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePasswordWithContext(ctx context.Context, request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    if request == nil {
        request = NewModfiyInstancePasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModfiyInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModfiyInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseWanAddressRequest() (request *ReleaseWanAddressRequest) {
    request = &ReleaseWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("redis", APIVersion, "ReleaseWanAddress")
    
    
    return
}

func NewReleaseWanAddressResponse() (response *ReleaseWanAddressResponse) {
    response = &ReleaseWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseWanAddress
// This API is used to disable public network access.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddress(request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    return c.ReleaseWanAddressWithContext(context.Background(), request)
}

// ReleaseWanAddress
// This API is used to disable public network access.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddressWithContext(ctx context.Context, request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    if request == nil {
        request = NewReleaseWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseWanAddressResponse()
    err = c.Send(request, response)
    return
}
