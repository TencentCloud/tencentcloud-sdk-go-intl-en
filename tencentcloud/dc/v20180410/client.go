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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-10"

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


func NewAcceptDirectConnectTunnelRequest() (request *AcceptDirectConnectTunnelRequest) {
    request = &AcceptDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "AcceptDirectConnectTunnel")
    
    
    return
}

func NewAcceptDirectConnectTunnelResponse() (response *AcceptDirectConnectTunnelResponse) {
    response = &AcceptDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AcceptDirectConnectTunnel
// This API is used to accept an application for a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) AcceptDirectConnectTunnel(request *AcceptDirectConnectTunnelRequest) (response *AcceptDirectConnectTunnelResponse, err error) {
    return c.AcceptDirectConnectTunnelWithContext(context.Background(), request)
}

// AcceptDirectConnectTunnel
// This API is used to accept an application for a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) AcceptDirectConnectTunnelWithContext(ctx context.Context, request *AcceptDirectConnectTunnelRequest) (response *AcceptDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewAcceptDirectConnectTunnelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptDirectConnectTunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewApplyInternetAddressRequest() (request *ApplyInternetAddressRequest) {
    request = &ApplyInternetAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ApplyInternetAddress")
    
    
    return
}

func NewApplyInternetAddressResponse() (response *ApplyInternetAddressResponse) {
    response = &ApplyInternetAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyInternetAddress
// This API is used to apply for an internet tunnel’s CIDR block.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ApplyInternetAddress(request *ApplyInternetAddressRequest) (response *ApplyInternetAddressResponse, err error) {
    return c.ApplyInternetAddressWithContext(context.Background(), request)
}

// ApplyInternetAddress
// This API is used to apply for an internet tunnel’s CIDR block.
//
// error code that may be returned:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ApplyInternetAddressWithContext(ctx context.Context, request *ApplyInternetAddressRequest) (response *ApplyInternetAddressResponse, err error) {
    if request == nil {
        request = NewApplyInternetAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyInternetAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyInternetAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectRequest() (request *CreateDirectConnectRequest) {
    request = &CreateDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateDirectConnect")
    
    
    return
}

func NewCreateDirectConnectResponse() (response *CreateDirectConnectResponse) {
    response = &CreateDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDirectConnect
// This API is used to apply for a connection.
//
// When calling this API, please note that:
//
// You need to complete identity verification for your account; otherwise, you cannot apply for a connection;
//
// If there is any connection in arrears under your account, you cannot apply for more connections.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTLIMITEXCEEDED = "LimitExceeded.DirectConnectLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDirectConnect(request *CreateDirectConnectRequest) (response *CreateDirectConnectResponse, err error) {
    return c.CreateDirectConnectWithContext(context.Background(), request)
}

// CreateDirectConnect
// This API is used to apply for a connection.
//
// When calling this API, please note that:
//
// You need to complete identity verification for your account; otherwise, you cannot apply for a connection;
//
// If there is any connection in arrears under your account, you cannot apply for more connections.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTLIMITEXCEEDED = "LimitExceeded.DirectConnectLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDirectConnectWithContext(ctx context.Context, request *CreateDirectConnectRequest) (response *CreateDirectConnectResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDirectConnect require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectTunnelRequest() (request *CreateDirectConnectTunnelRequest) {
    request = &CreateDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateDirectConnectTunnel")
    
    
    return
}

func NewCreateDirectConnectTunnelResponse() (response *CreateDirectConnectTunnelResponse) {
    response = &CreateDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDirectConnectTunnel
// This API is used to create a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ADDRESSERROR = "InvalidParameter.AddressError"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETER_UINISNOTEXIST = "InvalidParameter.UinIsNotExist"
//  INVALIDPARAMETER_VLANCONFLICT = "InvalidParameter.VlanConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VLANCONFLICT = "InvalidParameterValue.VlanConfLict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTLIMITEXCEEDED = "LimitExceeded.DirectConnectLimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTTUNNELLIMITEXCEEDED = "LimitExceeded.DirectConnectTunnelLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DCVPCISEXIST = "ResourceInUse.DcVpcIsExist"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CROSSBORDERDIRECTCONNECTTUNNEL = "UnsupportedOperation.CrossBorderDirectConnectTunnel"
func (c *Client) CreateDirectConnectTunnel(request *CreateDirectConnectTunnelRequest) (response *CreateDirectConnectTunnelResponse, err error) {
    return c.CreateDirectConnectTunnelWithContext(context.Background(), request)
}

// CreateDirectConnectTunnel
// This API is used to create a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ADDRESSERROR = "InvalidParameter.AddressError"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETER_UINISNOTEXIST = "InvalidParameter.UinIsNotExist"
//  INVALIDPARAMETER_VLANCONFLICT = "InvalidParameter.VlanConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VLANCONFLICT = "InvalidParameterValue.VlanConfLict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTLIMITEXCEEDED = "LimitExceeded.DirectConnectLimitExceeded"
//  LIMITEXCEEDED_DIRECTCONNECTTUNNELLIMITEXCEEDED = "LimitExceeded.DirectConnectTunnelLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DCVPCISEXIST = "ResourceInUse.DcVpcIsExist"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CROSSBORDERDIRECTCONNECTTUNNEL = "UnsupportedOperation.CrossBorderDirectConnectTunnel"
func (c *Client) CreateDirectConnectTunnelWithContext(ctx context.Context, request *CreateDirectConnectTunnelRequest) (response *CreateDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectTunnelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDirectConnectTunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectRequest() (request *DeleteDirectConnectRequest) {
    request = &DeleteDirectConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteDirectConnect")
    
    
    return
}

func NewDeleteDirectConnectResponse() (response *DeleteDirectConnectResponse) {
    response = &DeleteDirectConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDirectConnect
// This API is used to delete a connection.
//
// Only connected connections can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) DeleteDirectConnect(request *DeleteDirectConnectRequest) (response *DeleteDirectConnectResponse, err error) {
    return c.DeleteDirectConnectWithContext(context.Background(), request)
}

// DeleteDirectConnect
// This API is used to delete a connection.
//
// Only connected connections can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) DeleteDirectConnectWithContext(ctx context.Context, request *DeleteDirectConnectRequest) (response *DeleteDirectConnectResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDirectConnect require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDirectConnectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectTunnelRequest() (request *DeleteDirectConnectTunnelRequest) {
    request = &DeleteDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteDirectConnectTunnel")
    
    
    return
}

func NewDeleteDirectConnectTunnelResponse() (response *DeleteDirectConnectTunnelResponse) {
    response = &DeleteDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDirectConnectTunnel
// This API is used to delete a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) DeleteDirectConnectTunnel(request *DeleteDirectConnectTunnelRequest) (response *DeleteDirectConnectTunnelResponse, err error) {
    return c.DeleteDirectConnectTunnelWithContext(context.Background(), request)
}

// DeleteDirectConnectTunnel
// This API is used to delete a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) DeleteDirectConnectTunnelWithContext(ctx context.Context, request *DeleteDirectConnectTunnelRequest) (response *DeleteDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectTunnelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDirectConnectTunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessPointsRequest() (request *DescribeAccessPointsRequest) {
    request = &DescribeAccessPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeAccessPoints")
    
    
    return
}

func NewDescribeAccessPointsResponse() (response *DescribeAccessPointsResponse) {
    response = &DescribeAccessPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessPoints
// This API is used to query connection access points.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
    return c.DescribeAccessPointsWithContext(context.Background(), request)
}

// DescribeAccessPoints
// This API is used to query connection access points.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessPointsWithContext(ctx context.Context, request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessPointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessPoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectTunnelsRequest() (request *DescribeDirectConnectTunnelsRequest) {
    request = &DescribeDirectConnectTunnelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnectTunnels")
    
    
    return
}

func NewDescribeDirectConnectTunnelsResponse() (response *DescribeDirectConnectTunnelsResponse) {
    response = &DescribeDirectConnectTunnelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDirectConnectTunnels
// This API is used to query the list of dedicated tunnels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
func (c *Client) DescribeDirectConnectTunnels(request *DescribeDirectConnectTunnelsRequest) (response *DescribeDirectConnectTunnelsResponse, err error) {
    return c.DescribeDirectConnectTunnelsWithContext(context.Background(), request)
}

// DescribeDirectConnectTunnels
// This API is used to query the list of dedicated tunnels.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
func (c *Client) DescribeDirectConnectTunnelsWithContext(ctx context.Context, request *DescribeDirectConnectTunnelsRequest) (response *DescribeDirectConnectTunnelsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectTunnelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDirectConnectTunnels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDirectConnectTunnelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectsRequest() (request *DescribeDirectConnectsRequest) {
    request = &DescribeDirectConnectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnects")
    
    
    return
}

func NewDescribeDirectConnectsResponse() (response *DescribeDirectConnectsResponse) {
    response = &DescribeDirectConnectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDirectConnects
// This API is used to query the list of connections.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnects(request *DescribeDirectConnectsRequest) (response *DescribeDirectConnectsResponse, err error) {
    return c.DescribeDirectConnectsWithContext(context.Background(), request)
}

// DescribeDirectConnects
// This API is used to query the list of connections.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnectsWithContext(ctx context.Context, request *DescribeDirectConnectsRequest) (response *DescribeDirectConnectsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDirectConnects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDirectConnectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternetAddressRequest() (request *DescribeInternetAddressRequest) {
    request = &DescribeInternetAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeInternetAddress")
    
    
    return
}

func NewDescribeInternetAddressResponse() (response *DescribeInternetAddressResponse) {
    response = &DescribeInternetAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternetAddress
// This API is used to obtain the public IP address of an internet tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddress(request *DescribeInternetAddressRequest) (response *DescribeInternetAddressResponse, err error) {
    return c.DescribeInternetAddressWithContext(context.Background(), request)
}

// DescribeInternetAddress
// This API is used to obtain the public IP address of an internet tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddressWithContext(ctx context.Context, request *DescribeInternetAddressRequest) (response *DescribeInternetAddressResponse, err error) {
    if request == nil {
        request = NewDescribeInternetAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternetAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInternetAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternetAddressQuotaRequest() (request *DescribeInternetAddressQuotaRequest) {
    request = &DescribeInternetAddressQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeInternetAddressQuota")
    
    
    return
}

func NewDescribeInternetAddressQuotaResponse() (response *DescribeInternetAddressQuotaResponse) {
    response = &DescribeInternetAddressQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternetAddressQuota
// This API is used to obtain the public IP quota of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddressQuota(request *DescribeInternetAddressQuotaRequest) (response *DescribeInternetAddressQuotaResponse, err error) {
    return c.DescribeInternetAddressQuotaWithContext(context.Background(), request)
}

// DescribeInternetAddressQuota
// This API is used to obtain the public IP quota of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddressQuotaWithContext(ctx context.Context, request *DescribeInternetAddressQuotaRequest) (response *DescribeInternetAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeInternetAddressQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternetAddressQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInternetAddressQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternetAddressStatisticsRequest() (request *DescribeInternetAddressStatisticsRequest) {
    request = &DescribeInternetAddressStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeInternetAddressStatistics")
    
    
    return
}

func NewDescribeInternetAddressStatisticsResponse() (response *DescribeInternetAddressStatisticsResponse) {
    response = &DescribeInternetAddressStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternetAddressStatistics
// This API is used to obtain the public IP address assignment statistics of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddressStatistics(request *DescribeInternetAddressStatisticsRequest) (response *DescribeInternetAddressStatisticsResponse, err error) {
    return c.DescribeInternetAddressStatisticsWithContext(context.Background(), request)
}

// DescribeInternetAddressStatistics
// This API is used to obtain the public IP address assignment statistics of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternetAddressStatisticsWithContext(ctx context.Context, request *DescribeInternetAddressStatisticsRequest) (response *DescribeInternetAddressStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeInternetAddressStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternetAddressStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInternetAddressStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableInternetAddressRequest() (request *DisableInternetAddressRequest) {
    request = &DisableInternetAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DisableInternetAddress")
    
    
    return
}

func NewDisableInternetAddressResponse() (response *DisableInternetAddressResponse) {
    response = &DisableInternetAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableInternetAddress
// This API is used to disable a public IP address of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableInternetAddress(request *DisableInternetAddressRequest) (response *DisableInternetAddressResponse, err error) {
    return c.DisableInternetAddressWithContext(context.Background(), request)
}

// DisableInternetAddress
// This API is used to disable a public IP address of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableInternetAddressWithContext(ctx context.Context, request *DisableInternetAddressRequest) (response *DisableInternetAddressResponse, err error) {
    if request == nil {
        request = NewDisableInternetAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableInternetAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableInternetAddressResponse()
    err = c.Send(request, response)
    return
}

func NewEnableInternetAddressRequest() (request *EnableInternetAddressRequest) {
    request = &EnableInternetAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "EnableInternetAddress")
    
    
    return
}

func NewEnableInternetAddressResponse() (response *EnableInternetAddressResponse) {
    response = &EnableInternetAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableInternetAddress
// This API is used to enable a public IP address for internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableInternetAddress(request *EnableInternetAddressRequest) (response *EnableInternetAddressResponse, err error) {
    return c.EnableInternetAddressWithContext(context.Background(), request)
}

// EnableInternetAddress
// This API is used to enable a public IP address for internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableInternetAddressWithContext(ctx context.Context, request *EnableInternetAddressRequest) (response *EnableInternetAddressResponse, err error) {
    if request == nil {
        request = NewEnableInternetAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableInternetAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableInternetAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectAttributeRequest() (request *ModifyDirectConnectAttributeRequest) {
    request = &ModifyDirectConnectAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectAttribute")
    
    
    return
}

func NewModifyDirectConnectAttributeResponse() (response *ModifyDirectConnectAttributeResponse) {
    response = &ModifyDirectConnectAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDirectConnectAttribute
// This API is used to modify connection attributes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDirectConnectAttribute(request *ModifyDirectConnectAttributeRequest) (response *ModifyDirectConnectAttributeResponse, err error) {
    return c.ModifyDirectConnectAttributeWithContext(context.Background(), request)
}

// ModifyDirectConnectAttribute
// This API is used to modify connection attributes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DIRECTCONNECTIDISNOTUIN = "InvalidParameter.DirectConnectIdIsNotUin"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDirectConnectAttributeWithContext(ctx context.Context, request *ModifyDirectConnectAttributeRequest) (response *ModifyDirectConnectAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDirectConnectAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDirectConnectAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectTunnelAttributeRequest() (request *ModifyDirectConnectTunnelAttributeRequest) {
    request = &ModifyDirectConnectTunnelAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectTunnelAttribute")
    
    
    return
}

func NewModifyDirectConnectTunnelAttributeResponse() (response *ModifyDirectConnectTunnelAttributeResponse) {
    response = &ModifyDirectConnectTunnelAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDirectConnectTunnelAttribute
// This API is used to modify the dedicated tunnel attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) ModifyDirectConnectTunnelAttribute(request *ModifyDirectConnectTunnelAttributeRequest) (response *ModifyDirectConnectTunnelAttributeResponse, err error) {
    return c.ModifyDirectConnectTunnelAttributeWithContext(context.Background(), request)
}

// ModifyDirectConnectTunnelAttribute
// This API is used to modify the dedicated tunnel attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) ModifyDirectConnectTunnelAttributeWithContext(ctx context.Context, request *ModifyDirectConnectTunnelAttributeRequest) (response *ModifyDirectConnectTunnelAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectTunnelAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDirectConnectTunnelAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDirectConnectTunnelAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRejectDirectConnectTunnelRequest() (request *RejectDirectConnectTunnelRequest) {
    request = &RejectDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "RejectDirectConnectTunnel")
    
    
    return
}

func NewRejectDirectConnectTunnelResponse() (response *RejectDirectConnectTunnelResponse) {
    response = &RejectDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RejectDirectConnectTunnel
// This API is used to reject an application for a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) RejectDirectConnectTunnel(request *RejectDirectConnectTunnelRequest) (response *RejectDirectConnectTunnelResponse, err error) {
    return c.RejectDirectConnectTunnelWithContext(context.Background(), request)
}

// RejectDirectConnectTunnel
// This API is used to reject an application for a dedicated tunnel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) RejectDirectConnectTunnelWithContext(ctx context.Context, request *RejectDirectConnectTunnelRequest) (response *RejectDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewRejectDirectConnectTunnelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectDirectConnectTunnel require credential")
    }

    request.SetContext(ctx)
    
    response = NewRejectDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseInternetAddressRequest() (request *ReleaseInternetAddressRequest) {
    request = &ReleaseInternetAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ReleaseInternetAddress")
    
    
    return
}

func NewReleaseInternetAddressResponse() (response *ReleaseInternetAddressResponse) {
    response = &ReleaseInternetAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseInternetAddress
// This API is used to release an IP address of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) ReleaseInternetAddress(request *ReleaseInternetAddressRequest) (response *ReleaseInternetAddressResponse, err error) {
    return c.ReleaseInternetAddressWithContext(context.Background(), request)
}

// ReleaseInternetAddress
// This API is used to release an IP address of internet tunnels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DIRECTCONNECTTUNNELIDISNOTEXIST = "ResourceNotFound.DirectConnectTunnelIdIsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATECONFLICT = "UnsupportedOperation.StateConfLict"
func (c *Client) ReleaseInternetAddressWithContext(ctx context.Context, request *ReleaseInternetAddressRequest) (response *ReleaseInternetAddressResponse, err error) {
    if request == nil {
        request = NewReleaseInternetAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseInternetAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseInternetAddressResponse()
    err = c.Send(request, response)
    return
}
