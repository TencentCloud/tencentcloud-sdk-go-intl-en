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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

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


func NewAddLoginWhiteListRequest() (request *AddLoginWhiteListRequest) {
    request = &AddLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddLoginWhiteList")
    
    
    return
}

func NewAddLoginWhiteListResponse() (response *AddLoginWhiteListResponse) {
    response = &AddLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLoginWhiteList
// This API is used to add a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) AddLoginWhiteList(request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    return c.AddLoginWhiteListWithContext(context.Background(), request)
}

// AddLoginWhiteList
// This API is used to add a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) AddLoginWhiteListWithContext(ctx context.Context, request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddMachineTagRequest() (request *AddMachineTagRequest) {
    request = &AddMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddMachineTag")
    
    
    return
}

func NewAddMachineTagResponse() (response *AddMachineTagResponse) {
    response = &AddMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddMachineTag
// This API is used to add a tag to a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddMachineTag(request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    return c.AddMachineTagWithContext(context.Background(), request)
}

// AddMachineTag
// This API is used to add a tag to a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddMachineTagWithContext(ctx context.Context, request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    if request == nil {
        request = NewAddMachineTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMachineTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
    request = &CloseProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CloseProVersion")
    
    
    return
}

func NewCloseProVersionResponse() (response *CloseProVersionResponse) {
    response = &CloseProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseProVersion
// This API is used to deactivate CWP Pro.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOSEPROVERSION = "FailedOperation.CloseProVersion"
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  FAILEDOPERATION_PREPAYMODE = "FailedOperation.PrePayMode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CloseProVersion(request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    return c.CloseProVersionWithContext(context.Background(), request)
}

// CloseProVersion
// This API is used to deactivate CWP Pro.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOSEPROVERSION = "FailedOperation.CloseProVersion"
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  FAILEDOPERATION_PREPAYMODE = "FailedOperation.PrePayMode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CloseProVersionWithContext(ctx context.Context, request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    if request == nil {
        request = NewCloseProVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenPortTaskRequest() (request *CreateOpenPortTaskRequest) {
    request = &CreateOpenPortTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateOpenPortTask")
    
    
    return
}

func NewCreateOpenPortTaskResponse() (response *CreateOpenPortTaskResponse) {
    response = &CreateOpenPortTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenPortTask
// This API is used to create a real-time port acquisition task.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEOPENPORTTASK = "FailedOperation.CreateOpenPortTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOpenPortTask(request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    return c.CreateOpenPortTaskWithContext(context.Background(), request)
}

// CreateOpenPortTask
// This API is used to create a real-time port acquisition task.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEOPENPORTTASK = "FailedOperation.CreateOpenPortTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOpenPortTaskWithContext(ctx context.Context, request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    if request == nil {
        request = NewCreateOpenPortTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenPortTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenPortTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessTaskRequest() (request *CreateProcessTaskRequest) {
    request = &CreateProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateProcessTask")
    
    
    return
}

func NewCreateProcessTaskResponse() (response *CreateProcessTaskResponse) {
    response = &CreateProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProcessTask
// This API is used to create a real-time process pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPROCESSTASK = "FailedOperation.CreateProcessTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProcessTask(request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    return c.CreateProcessTaskWithContext(context.Background(), request)
}

// CreateProcessTask
// This API is used to create a real-time process pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPROCESSTASK = "FailedOperation.CreateProcessTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProcessTaskWithContext(ctx context.Context, request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateProcessTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcessTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsualLoginPlacesRequest() (request *CreateUsualLoginPlacesRequest) {
    request = &CreateUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateUsualLoginPlaces")
    
    
    return
}

func NewCreateUsualLoginPlacesResponse() (response *CreateUsualLoginPlacesResponse) {
    response = &CreateUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUsualLoginPlaces
// This API is used to add one or more usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUsualLoginPlaces(request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    return c.CreateUsualLoginPlacesWithContext(context.Background(), request)
}

// CreateUsualLoginPlaces
// This API is used to add one or more usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUsualLoginPlacesWithContext(ctx context.Context, request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewCreateUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
    request = &DeleteBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBruteAttacks")
    
    
    return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
    response = &DeleteBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBruteAttacks
// This API is used to delete brute force attack records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    return c.DeleteBruteAttacksWithContext(context.Background(), request)
}

// DeleteBruteAttacks
// This API is used to delete brute force attack records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteBruteAttacksWithContext(ctx context.Context, request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDeleteBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
    request = &DeleteLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteLoginWhiteList")
    
    
    return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
    response = &DeleteLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoginWhiteList
// This API is used to delete a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    return c.DeleteLoginWhiteListWithContext(context.Background(), request)
}

// DeleteLoginWhiteList
// This API is used to delete a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteLoginWhiteListWithContext(ctx context.Context, request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
    request = &DeleteMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachine")
    
    
    return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
    response = &DeleteMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMachine
// This API is used to uninstall the CWP agent.
//
// error code that may be returned:
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    return c.DeleteMachineWithContext(context.Background(), request)
}

// DeleteMachine
// This API is used to uninstall the CWP agent.
//
// error code that may be returned:
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMachineWithContext(ctx context.Context, request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    if request == nil {
        request = NewDeleteMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineTagRequest() (request *DeleteMachineTagRequest) {
    request = &DeleteMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachineTag")
    
    
    return
}

func NewDeleteMachineTagResponse() (response *DeleteMachineTagResponse) {
    response = &DeleteMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMachineTag
// This API is used to remove a tag from a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineTag(request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    return c.DeleteMachineTagWithContext(context.Background(), request)
}

// DeleteMachineTag
// This API is used to remove a tag from a server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineTagWithContext(ctx context.Context, request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    if request == nil {
        request = NewDeleteMachineTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMaliciousRequestsRequest() (request *DeleteMaliciousRequestsRequest) {
    request = &DeleteMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMaliciousRequests")
    
    
    return
}

func NewDeleteMaliciousRequestsResponse() (response *DeleteMaliciousRequestsResponse) {
    response = &DeleteMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMaliciousRequests
// This API is used to delete malicious request records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    return c.DeleteMaliciousRequestsWithContext(context.Background(), request)
}

// DeleteMaliciousRequests
// This API is used to delete malicious request records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMaliciousRequestsWithContext(ctx context.Context, request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
    request = &DeleteMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMalwares")
    
    
    return
}

func NewDeleteMalwaresResponse() (response *DeleteMalwaresResponse) {
    response = &DeleteMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMalwares
// This API is used to delete trojan records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    return c.DeleteMalwaresWithContext(context.Background(), request)
}

// DeleteMalwares
// This API is used to delete trojan records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMalwaresWithContext(ctx context.Context, request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    if request == nil {
        request = NewDeleteMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNonlocalLoginPlacesRequest() (request *DeleteNonlocalLoginPlacesRequest) {
    request = &DeleteNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteNonlocalLoginPlaces")
    
    
    return
}

func NewDeleteNonlocalLoginPlacesResponse() (response *DeleteNonlocalLoginPlacesResponse) {
    response = &DeleteNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNonlocalLoginPlaces
// This API is used to delete unusual login location records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    return c.DeleteNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DeleteNonlocalLoginPlaces
// This API is used to delete unusual login location records.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteNonlocalLoginPlacesWithContext(ctx context.Context, request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsualLoginPlacesRequest() (request *DeleteUsualLoginPlacesRequest) {
    request = &DeleteUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteUsualLoginPlaces")
    
    
    return
}

func NewDeleteUsualLoginPlacesResponse() (response *DeleteUsualLoginPlacesResponse) {
    response = &DeleteUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsualLoginPlaces
// This API is used to delete one or more usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteUsualLoginPlaces(request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    return c.DeleteUsualLoginPlacesWithContext(context.Background(), request)
}

// DeleteUsualLoginPlaces
// This API is used to delete one or more usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteUsualLoginPlacesWithContext(ctx context.Context, request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
    request = &DescribeAccountStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccountStatistics")
    
    
    return
}

func NewDescribeAccountStatisticsResponse() (response *DescribeAccountStatisticsResponse) {
    response = &DescribeAccountStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountStatistics
// This API is used to get the account statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountStatistics(request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    return c.DescribeAccountStatisticsWithContext(context.Background(), request)
}

// DescribeAccountStatistics
// This API is used to get the account statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountStatisticsWithContext(ctx context.Context, request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// This API is used to get the account list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// This API is used to get the account list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVulsRequest() (request *DescribeAgentVulsRequest) {
    request = &DescribeAgentVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAgentVuls")
    
    
    return
}

func NewDescribeAgentVulsResponse() (response *DescribeAgentVulsResponse) {
    response = &DescribeAgentVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentVuls
// This API is used to get the list of vulnerabilities on a single server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentVuls(request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    return c.DescribeAgentVulsWithContext(context.Background(), request)
}

// DescribeAgentVuls
// This API is used to get the list of vulnerabilities on a single server.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentVulsWithContext(ctx context.Context, request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
    request = &DescribeAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAlarmAttribute")
    
    
    return
}

func NewDescribeAlarmAttributeResponse() (response *DescribeAlarmAttributeResponse) {
    response = &DescribeAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmAttribute
// This API is used to get the alarm settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAlarmAttribute(request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    return c.DescribeAlarmAttributeWithContext(context.Background(), request)
}

// DescribeAlarmAttribute
// This API is used to get the alarm settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAlarmAttributeWithContext(ctx context.Context, request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttacksRequest() (request *DescribeBruteAttacksRequest) {
    request = &DescribeBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBruteAttacks")
    
    
    return
}

func NewDescribeBruteAttacksResponse() (response *DescribeBruteAttacksResponse) {
    response = &DescribeBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBruteAttacks
// This API is used to get the brute force attack event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeBruteAttacks(request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    return c.DescribeBruteAttacksWithContext(context.Background(), request)
}

// DescribeBruteAttacks
// This API is used to get the brute force attack event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeBruteAttacksWithContext(ctx context.Context, request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentInfoRequest() (request *DescribeComponentInfoRequest) {
    request = &DescribeComponentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentInfo")
    
    
    return
}

func NewDescribeComponentInfoResponse() (response *DescribeComponentInfoResponse) {
    response = &DescribeComponentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponentInfo
// This API is used to get the component information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeComponentInfo(request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    return c.DescribeComponentInfoWithContext(context.Background(), request)
}

// DescribeComponentInfo
// This API is used to get the component information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeComponentInfoWithContext(ctx context.Context, request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComponentInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponentInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
    request = &DescribeComponentStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentStatistics")
    
    
    return
}

func NewDescribeComponentStatisticsResponse() (response *DescribeComponentStatisticsResponse) {
    response = &DescribeComponentStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponentStatistics
// This API is used to get the component statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    return c.DescribeComponentStatisticsWithContext(context.Background(), request)
}

// DescribeComponentStatistics
// This API is used to get the component statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeComponentStatisticsWithContext(ctx context.Context, request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponentStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
    request = &DescribeComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponents")
    
    
    return
}

func NewDescribeComponentsResponse() (response *DescribeComponentsResponse) {
    response = &DescribeComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponents
// This API is used to get the component list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    return c.DescribeComponentsWithContext(context.Background(), request)
}

// DescribeComponents
// This API is used to get the component list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeComponentsWithContext(ctx context.Context, request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
    request = &DescribeHistoryAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeHistoryAccounts")
    
    
    return
}

func NewDescribeHistoryAccountsResponse() (response *DescribeHistoryAccountsResponse) {
    response = &DescribeHistoryAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryAccounts
// This API is used to get the account change history list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistoryAccounts(request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    return c.DescribeHistoryAccountsWithContext(context.Background(), request)
}

// DescribeHistoryAccounts
// This API is used to get the account change history list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistoryAccountsWithContext(ctx context.Context, request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistoryAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistoryAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImpactedHostsRequest() (request *DescribeImpactedHostsRequest) {
    request = &DescribeImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeImpactedHosts")
    
    
    return
}

func NewDescribeImpactedHostsResponse() (response *DescribeImpactedHostsResponse) {
    response = &DescribeImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImpactedHosts
// This API is used to get the list of servers affected by a vulnerability.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImpactedHosts(request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    return c.DescribeImpactedHostsWithContext(context.Background(), request)
}

// DescribeImpactedHosts
// This API is used to get the list of servers affected by a vulnerability.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImpactedHostsWithContext(ctx context.Context, request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    if request == nil {
        request = NewDescribeImpactedHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImpactedHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
    request = &DescribeLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeLoginWhiteList")
    
    
    return
}

func NewDescribeLoginWhiteListResponse() (response *DescribeLoginWhiteListResponse) {
    response = &DescribeLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoginWhiteList
// This API is used to get the list of login allowlist entries.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeLoginWhiteList(request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    return c.DescribeLoginWhiteListWithContext(context.Background(), request)
}

// DescribeLoginWhiteList
// This API is used to get the list of login allowlist entries.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeLoginWhiteListWithContext(ctx context.Context, request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineInfoRequest() (request *DescribeMachineInfoRequest) {
    request = &DescribeMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachineInfo")
    
    
    return
}

func NewDescribeMachineInfoResponse() (response *DescribeMachineInfoResponse) {
    response = &DescribeMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineInfo
// This API is used to get server details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineInfo(request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    return c.DescribeMachineInfoWithContext(context.Background(), request)
}

// DescribeMachineInfo
// This API is used to get server details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineInfoWithContext(ctx context.Context, request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMachineInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachines")
    
    
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachines
// This API is used to get the list of servers in a specified region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    return c.DescribeMachinesWithContext(context.Background(), request)
}

// DescribeMachines
// This API is used to get the list of servers in a specified region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachinesWithContext(ctx context.Context, request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaliciousRequestsRequest() (request *DescribeMaliciousRequestsRequest) {
    request = &DescribeMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMaliciousRequests")
    
    
    return
}

func NewDescribeMaliciousRequestsResponse() (response *DescribeMaliciousRequestsResponse) {
    response = &DescribeMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaliciousRequests
// This API is used to get malicious request data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeMaliciousRequests(request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    return c.DescribeMaliciousRequestsWithContext(context.Background(), request)
}

// DescribeMaliciousRequests
// This API is used to get malicious request data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeMaliciousRequestsWithContext(ctx context.Context, request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwaresRequest() (request *DescribeMalwaresRequest) {
    request = &DescribeMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMalwares")
    
    
    return
}

func NewDescribeMalwaresResponse() (response *DescribeMalwaresResponse) {
    response = &DescribeMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwares
// This API is used to get the list of trojan events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeMalwares(request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    return c.DescribeMalwaresWithContext(context.Background(), request)
}

// DescribeMalwares
// This API is used to get the list of trojan events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeMalwaresWithContext(ctx context.Context, request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNonlocalLoginPlacesRequest() (request *DescribeNonlocalLoginPlacesRequest) {
    request = &DescribeNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeNonlocalLoginPlaces")
    
    
    return
}

func NewDescribeNonlocalLoginPlacesResponse() (response *DescribeNonlocalLoginPlacesResponse) {
    response = &DescribeNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNonlocalLoginPlaces
// This API is used to get unusual login events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeNonlocalLoginPlaces(request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    return c.DescribeNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DescribeNonlocalLoginPlaces
// This API is used to get unusual login events.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeNonlocalLoginPlacesWithContext(ctx context.Context, request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
    request = &DescribeOpenPortStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortStatistics")
    
    
    return
}

func NewDescribeOpenPortStatisticsResponse() (response *DescribeOpenPortStatisticsResponse) {
    response = &DescribeOpenPortStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPortStatistics
// This API is used to get the statistics on port usage
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    return c.DescribeOpenPortStatisticsWithContext(context.Background(), request)
}

// DescribeOpenPortStatistics
// This API is used to get the statistics on port usage
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeOpenPortStatisticsWithContext(ctx context.Context, request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPortStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortTaskStatusRequest() (request *DescribeOpenPortTaskStatusRequest) {
    request = &DescribeOpenPortTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortTaskStatus")
    
    
    return
}

func NewDescribeOpenPortTaskStatusResponse() (response *DescribeOpenPortTaskStatusResponse) {
    response = &DescribeOpenPortTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPortTaskStatus
// This API is used to get the status of a real-time port pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_OPENPORTTASKNOTFOUND = "FailedOperation.OpenPortTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOpenPortTaskStatus(request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    return c.DescribeOpenPortTaskStatusWithContext(context.Background(), request)
}

// DescribeOpenPortTaskStatus
// This API is used to get the status of a real-time port pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_OPENPORTTASKNOTFOUND = "FailedOperation.OpenPortTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOpenPortTaskStatusWithContext(ctx context.Context, request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPortTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortsRequest() (request *DescribeOpenPortsRequest) {
    request = &DescribeOpenPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPorts")
    
    
    return
}

func NewDescribeOpenPortsResponse() (response *DescribeOpenPortsResponse) {
    response = &DescribeOpenPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPorts
// This API is used to get the port list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPorts(request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    return c.DescribeOpenPortsWithContext(context.Background(), request)
}

// DescribeOpenPorts
// This API is used to get the port list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPortsWithContext(ctx context.Context, request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
    request = &DescribeOverviewStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOverviewStatistics")
    
    
    return
}

func NewDescribeOverviewStatisticsResponse() (response *DescribeOverviewStatisticsResponse) {
    response = &DescribeOverviewStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewStatistics
// This API is used to get the overview statistics of CWP under the current account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    return c.DescribeOverviewStatisticsWithContext(context.Background(), request)
}

// DescribeOverviewStatistics
// This API is used to get the overview statistics of CWP under the current account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOverviewStatisticsWithContext(ctx context.Context, request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProVersionInfo")
    
    
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProVersionInfo
// This API is used to get the CWP Pro information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    return c.DescribeProVersionInfoWithContext(context.Background(), request)
}

// DescribeProVersionInfo
// This API is used to get the CWP Pro information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeProVersionInfoWithContext(ctx context.Context, request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProVersionInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProVersionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
    request = &DescribeProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessStatistics")
    
    
    return
}

func NewDescribeProcessStatisticsResponse() (response *DescribeProcessStatisticsResponse) {
    response = &DescribeProcessStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcessStatistics
// This API is used to get the process statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    return c.DescribeProcessStatisticsWithContext(context.Background(), request)
}

// DescribeProcessStatistics
// This API is used to get the process statistics list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessStatisticsWithContext(ctx context.Context, request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessTaskStatusRequest() (request *DescribeProcessTaskStatusRequest) {
    request = &DescribeProcessTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessTaskStatus")
    
    
    return
}

func NewDescribeProcessTaskStatusResponse() (response *DescribeProcessTaskStatusResponse) {
    response = &DescribeProcessTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcessTaskStatus
// This API is used to get the status of a real-time process pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_PROCESSTASKNOTFOUND = "FailedOperation.ProcessTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProcessTaskStatus(request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    return c.DescribeProcessTaskStatusWithContext(context.Background(), request)
}

// DescribeProcessTaskStatus
// This API is used to get the status of a real-time process pulling task.
//
// error code that may be returned:
//  FAILEDOPERATION_PROCESSTASKNOTFOUND = "FailedOperation.ProcessTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProcessTaskStatusWithContext(ctx context.Context, request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProcessTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcesses")
    
    
    return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
    response = &DescribeProcessesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcesses
// This API is used to get the process list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    return c.DescribeProcessesWithContext(context.Background(), request)
}

// DescribeProcesses
// This API is used to get the process list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessesWithContext(ctx context.Context, request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcesses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
    request = &DescribeSecurityDynamicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityDynamics")
    
    
    return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
    response = &DescribeSecurityDynamicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityDynamics
// This API is used to get the security event message data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    return c.DescribeSecurityDynamicsWithContext(context.Background(), request)
}

// DescribeSecurityDynamics
// This API is used to get the security event message data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityDynamicsWithContext(ctx context.Context, request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityDynamicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityDynamics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityDynamicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
    request = &DescribeSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityTrends")
    
    
    return
}

func NewDescribeSecurityTrendsResponse() (response *DescribeSecurityTrendsResponse) {
    response = &DescribeSecurityTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityTrends
// This API is used to get the security event statistics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    return c.DescribeSecurityTrendsWithContext(context.Background(), request)
}

// DescribeSecurityTrends
// This API is used to get the security event statistics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityTrendsWithContext(ctx context.Context, request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
    request = &DescribeTagMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTagMachines")
    
    
    return
}

func NewDescribeTagMachinesResponse() (response *DescribeTagMachinesResponse) {
    response = &DescribeTagMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagMachines
// This API is used to get the information of servers associated with a specified tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagMachines(request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    return c.DescribeTagMachinesWithContext(context.Background(), request)
}

// DescribeTagMachines
// This API is used to get the information of servers associated with a specified tag.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagMachinesWithContext(ctx context.Context, request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeTagMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTags")
    
    
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTags
// This API is used to get all server tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    return c.DescribeTagsWithContext(context.Background(), request)
}

// DescribeTags
// This API is used to get all server tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagsWithContext(ctx context.Context, request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
    request = &DescribeUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeUsualLoginPlaces")
    
    
    return
}

func NewDescribeUsualLoginPlacesResponse() (response *DescribeUsualLoginPlacesResponse) {
    response = &DescribeUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsualLoginPlaces
// This API is used to query usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    return c.DescribeUsualLoginPlacesWithContext(context.Background(), request)
}

// DescribeUsualLoginPlaces
// This API is used to query usual login locations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsualLoginPlacesWithContext(ctx context.Context, request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulInfoRequest() (request *DescribeVulInfoRequest) {
    request = &DescribeVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulInfo")
    
    
    return
}

func NewDescribeVulInfoResponse() (response *DescribeVulInfoResponse) {
    response = &DescribeVulInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulInfo
// This API is used to get vulnerability details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulInfo(request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    return c.DescribeVulInfoWithContext(context.Background(), request)
}

// DescribeVulInfo
// This API is used to get vulnerability details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulInfoWithContext(ctx context.Context, request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanResultRequest() (request *DescribeVulScanResultRequest) {
    request = &DescribeVulScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulScanResult")
    
    
    return
}

func NewDescribeVulScanResultResponse() (response *DescribeVulScanResultResponse) {
    response = &DescribeVulScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulScanResult
// This API is used to get the vulnerability detection result.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulScanResult(request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    return c.DescribeVulScanResultWithContext(context.Background(), request)
}

// DescribeVulScanResult
// This API is used to get the vulnerability detection result.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulScanResultWithContext(ctx context.Context, request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulScanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVuls")
    
    
    return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
    response = &DescribeVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVuls
// This API is used to get the vulnerability list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    return c.DescribeVulsWithContext(context.Background(), request)
}

// DescribeVuls
// This API is used to get the vulnerability list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulsWithContext(ctx context.Context, request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportBruteAttacksRequest() (request *DescribeWeeklyReportBruteAttacksRequest) {
    request = &DescribeWeeklyReportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportBruteAttacks")
    
    
    return
}

func NewDescribeWeeklyReportBruteAttacksResponse() (response *DescribeWeeklyReportBruteAttacksResponse) {
    response = &DescribeWeeklyReportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportBruteAttacks
// This API is used to get the brute force attack data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportBruteAttacks(request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    return c.DescribeWeeklyReportBruteAttacksWithContext(context.Background(), request)
}

// DescribeWeeklyReportBruteAttacks
// This API is used to get the brute force attack data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportBruteAttacksWithContext(ctx context.Context, request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportInfoRequest() (request *DescribeWeeklyReportInfoRequest) {
    request = &DescribeWeeklyReportInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportInfo")
    
    
    return
}

func NewDescribeWeeklyReportInfoResponse() (response *DescribeWeeklyReportInfoResponse) {
    response = &DescribeWeeklyReportInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportInfo
// This API is used to get the details in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportInfo(request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    return c.DescribeWeeklyReportInfoWithContext(context.Background(), request)
}

// DescribeWeeklyReportInfo
// This API is used to get the details in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportInfoWithContext(ctx context.Context, request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportMalwaresRequest() (request *DescribeWeeklyReportMalwaresRequest) {
    request = &DescribeWeeklyReportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportMalwares")
    
    
    return
}

func NewDescribeWeeklyReportMalwaresResponse() (response *DescribeWeeklyReportMalwaresResponse) {
    response = &DescribeWeeklyReportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportMalwares
// This API is used to get the trojan data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportMalwares(request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    return c.DescribeWeeklyReportMalwaresWithContext(context.Background(), request)
}

// DescribeWeeklyReportMalwares
// This API is used to get the trojan data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportMalwaresWithContext(ctx context.Context, request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesRequest() (request *DescribeWeeklyReportNonlocalLoginPlacesRequest) {
    request = &DescribeWeeklyReportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportNonlocalLoginPlaces")
    
    
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesResponse() (response *DescribeWeeklyReportNonlocalLoginPlacesResponse) {
    response = &DescribeWeeklyReportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportNonlocalLoginPlaces
// This API is used to get the unusual login location data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportNonlocalLoginPlaces(request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    return c.DescribeWeeklyReportNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DescribeWeeklyReportNonlocalLoginPlaces
// This API is used to get the unusual login location data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportNonlocalLoginPlacesWithContext(ctx context.Context, request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportVulsRequest() (request *DescribeWeeklyReportVulsRequest) {
    request = &DescribeWeeklyReportVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportVuls")
    
    
    return
}

func NewDescribeWeeklyReportVulsResponse() (response *DescribeWeeklyReportVulsResponse) {
    response = &DescribeWeeklyReportVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportVuls
// This API is used to get the vulnerability data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportVuls(request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    return c.DescribeWeeklyReportVulsWithContext(context.Background(), request)
}

// DescribeWeeklyReportVuls
// This API is used to get the vulnerability data in the weekly CWP Pro report.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportVulsWithContext(ctx context.Context, request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportsRequest() (request *DescribeWeeklyReportsRequest) {
    request = &DescribeWeeklyReportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReports")
    
    
    return
}

func NewDescribeWeeklyReportsResponse() (response *DescribeWeeklyReportsResponse) {
    response = &DescribeWeeklyReportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReports
// This API is used to get the weekly report list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReports(request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    return c.DescribeWeeklyReportsWithContext(context.Background(), request)
}

// DescribeWeeklyReports
// This API is used to get the weekly report list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportsWithContext(ctx context.Context, request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportsResponse()
    err = c.Send(request, response)
    return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
    request = &EditTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditTags")
    
    
    return
}

func NewEditTagsResponse() (response *EditTagsResponse) {
    response = &EditTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditTags
// This API is used to add or edit tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
    return c.EditTagsWithContext(context.Background(), request)
}

// EditTags
// This API is used to add or edit tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditTagsWithContext(ctx context.Context, request *EditTagsRequest) (response *EditTagsResponse, err error) {
    if request == nil {
        request = NewEditTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditTagsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBruteAttacksRequest() (request *ExportBruteAttacksRequest) {
    request = &ExportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBruteAttacks")
    
    
    return
}

func NewExportBruteAttacksResponse() (response *ExportBruteAttacksResponse) {
    response = &ExportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBruteAttacks
// This API is used to export brute force attack records into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    return c.ExportBruteAttacksWithContext(context.Background(), request)
}

// ExportBruteAttacks
// This API is used to export brute force attack records into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportBruteAttacksWithContext(ctx context.Context, request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
    request = &ExportMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMaliciousRequests")
    
    
    return
}

func NewExportMaliciousRequestsResponse() (response *ExportMaliciousRequestsResponse) {
    response = &ExportMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportMaliciousRequests
// This API is used to export the malicious request file into a CSV file for download.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ExportMaliciousRequests(request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    return c.ExportMaliciousRequestsWithContext(context.Background(), request)
}

// ExportMaliciousRequests
// This API is used to export the malicious request file into a CSV file for download.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ExportMaliciousRequestsWithContext(ctx context.Context, request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewExportMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewExportMalwaresRequest() (request *ExportMalwaresRequest) {
    request = &ExportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMalwares")
    
    
    return
}

func NewExportMalwaresResponse() (response *ExportMalwaresResponse) {
    response = &ExportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportMalwares
// This API is used to export trojan records into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    return c.ExportMalwaresWithContext(context.Background(), request)
}

// ExportMalwares
// This API is used to export trojan records into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportMalwaresWithContext(ctx context.Context, request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    if request == nil {
        request = NewExportMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewExportNonlocalLoginPlacesRequest() (request *ExportNonlocalLoginPlacesRequest) {
    request = &ExportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportNonlocalLoginPlaces")
    
    
    return
}

func NewExportNonlocalLoginPlacesResponse() (response *ExportNonlocalLoginPlacesResponse) {
    response = &ExportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportNonlocalLoginPlaces
// This API is used to export unusual login location events into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportNonlocalLoginPlaces(request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    return c.ExportNonlocalLoginPlacesWithContext(context.Background(), request)
}

// ExportNonlocalLoginPlaces
// This API is used to export unusual login location events into a CSV file.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
func (c *Client) ExportNonlocalLoginPlacesWithContext(ctx context.Context, request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewExportNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
    request = &IgnoreImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "IgnoreImpactedHosts")
    
    
    return
}

func NewIgnoreImpactedHostsResponse() (response *IgnoreImpactedHostsResponse) {
    response = &IgnoreImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IgnoreImpactedHosts
// This API is used to ignore one or more vulnerabilities.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    return c.IgnoreImpactedHostsWithContext(context.Background(), request)
}

// IgnoreImpactedHosts
// This API is used to ignore one or more vulnerabilities.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) IgnoreImpactedHostsWithContext(ctx context.Context, request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    if request == nil {
        request = NewIgnoreImpactedHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreImpactedHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewMisAlarmNonlocalLoginPlacesRequest() (request *MisAlarmNonlocalLoginPlacesRequest) {
    request = &MisAlarmNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "MisAlarmNonlocalLoginPlaces")
    
    
    return
}

func NewMisAlarmNonlocalLoginPlacesResponse() (response *MisAlarmNonlocalLoginPlacesResponse) {
    response = &MisAlarmNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MisAlarmNonlocalLoginPlaces
// This API is used to set the current location as the usual login location.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) MisAlarmNonlocalLoginPlaces(request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    return c.MisAlarmNonlocalLoginPlacesWithContext(context.Background(), request)
}

// MisAlarmNonlocalLoginPlaces
// This API is used to set the current location as the usual login location.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) MisAlarmNonlocalLoginPlacesWithContext(ctx context.Context, request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewMisAlarmNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MisAlarmNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewMisAlarmNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmAttributeRequest() (request *ModifyAlarmAttributeRequest) {
    request = &ModifyAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAlarmAttribute")
    
    
    return
}

func NewModifyAlarmAttributeResponse() (response *ModifyAlarmAttributeResponse) {
    response = &ModifyAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmAttribute
// This API is used to modify alarm settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAlarmAttribute(request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    return c.ModifyAlarmAttributeWithContext(context.Background(), request)
}

// ModifyAlarmAttribute
// This API is used to modify alarm settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAlarmAttributeWithContext(ctx context.Context, request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
    request = &ModifyAutoOpenProVersionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAutoOpenProVersionConfig")
    
    
    return
}

func NewModifyAutoOpenProVersionConfigResponse() (response *ModifyAutoOpenProVersionConfigResponse) {
    response = &ModifyAutoOpenProVersionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoOpenProVersionConfig
// This API is used to set whether to automatically activate CWP Pro for newly added servers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    return c.ModifyAutoOpenProVersionConfigWithContext(context.Background(), request)
}

// ModifyAutoOpenProVersionConfig
// This API is used to set whether to automatically activate CWP Pro for newly added servers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAutoOpenProVersionConfigWithContext(ctx context.Context, request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoOpenProVersionConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoOpenProVersionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteListRequest() (request *ModifyLoginWhiteListRequest) {
    request = &ModifyLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyLoginWhiteList")
    
    
    return
}

func NewModifyLoginWhiteListResponse() (response *ModifyLoginWhiteListResponse) {
    response = &ModifyLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoginWhiteList
// This API is used to edit a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyLoginWhiteList(request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    return c.ModifyLoginWhiteListWithContext(context.Background(), request)
}

// ModifyLoginWhiteList
// This API is used to edit a allowlist rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyLoginWhiteListWithContext(ctx context.Context, request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
    request = &ModifyProVersionRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyProVersionRenewFlag")
    
    
    return
}

func NewModifyProVersionRenewFlagResponse() (response *ModifyProVersionRenewFlagResponse) {
    response = &ModifyProVersionRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProVersionRenewFlag
// This API is used to modify the renewal flag of CWP Pro.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    return c.ModifyProVersionRenewFlagWithContext(context.Background(), request)
}

// ModifyProVersionRenewFlag
// This API is used to modify the renewal flag of CWP Pro.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProVersionRenewFlagWithContext(ctx context.Context, request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProVersionRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProVersionRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
    request = &OpenProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersion")
    
    
    return
}

func NewOpenProVersionResponse() (response *OpenProVersionResponse) {
    response = &OpenProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenProVersion
// This API is used to activate CWP Pro.
//
// error code that may be returned:
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) OpenProVersion(request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    return c.OpenProVersionWithContext(context.Background(), request)
}

// OpenProVersion
// This API is used to activate CWP Pro.
//
// error code that may be returned:
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) OpenProVersionWithContext(ctx context.Context, request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    if request == nil {
        request = NewOpenProVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenProVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
    request = &RecoverMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RecoverMalwares")
    
    
    return
}

func NewRecoverMalwaresResponse() (response *RecoverMalwaresResponse) {
    response = &RecoverMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverMalwares
// This API is used to recover isolated trojan files in a batch.
//
// error code that may be returned:
//  FAILEDOPERATION_RECOVER = "FailedOperation.Recover"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) RecoverMalwares(request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    return c.RecoverMalwaresWithContext(context.Background(), request)
}

// RecoverMalwares
// This API is used to recover isolated trojan files in a batch.
//
// error code that may be returned:
//  FAILEDOPERATION_RECOVER = "FailedOperation.Recover"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) RecoverMalwaresWithContext(ctx context.Context, request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    if request == nil {
        request = NewRecoverMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewRescanImpactedHostRequest() (request *RescanImpactedHostRequest) {
    request = &RescanImpactedHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RescanImpactedHost")
    
    
    return
}

func NewRescanImpactedHostResponse() (response *RescanImpactedHostResponse) {
    response = &RescanImpactedHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RescanImpactedHost
// This API is used to re-scan for vulnerabilities.
//
// error code that may be returned:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"
//  FAILEDOPERATION_RESCANVULPROCESSINUSE = "FailedOperation.RescanVulProcessInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RescanImpactedHost(request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    return c.RescanImpactedHostWithContext(context.Background(), request)
}

// RescanImpactedHost
// This API is used to re-scan for vulnerabilities.
//
// error code that may be returned:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"
//  FAILEDOPERATION_RESCANVULPROCESSINUSE = "FailedOperation.RescanVulProcessInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RescanImpactedHostWithContext(ctx context.Context, request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    if request == nil {
        request = NewRescanImpactedHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RescanImpactedHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewRescanImpactedHostResponse()
    err = c.Send(request, response)
    return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
    request = &SeparateMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SeparateMalwares")
    
    
    return
}

func NewSeparateMalwaresResponse() (response *SeparateMalwaresResponse) {
    response = &SeparateMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SeparateMalwares
// This API is used to isolate trojans.
//
// error code that may be returned:
//  FAILEDOPERATION_PARTSEPARATE = "FailedOperation.PartSeparate"
//  FAILEDOPERATION_SINGLESEPARATE = "FailedOperation.SingleSeparate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) SeparateMalwares(request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    return c.SeparateMalwaresWithContext(context.Background(), request)
}

// SeparateMalwares
// This API is used to isolate trojans.
//
// error code that may be returned:
//  FAILEDOPERATION_PARTSEPARATE = "FailedOperation.PartSeparate"
//  FAILEDOPERATION_SINGLESEPARATE = "FailedOperation.SingleSeparate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) SeparateMalwaresWithContext(ctx context.Context, request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    if request == nil {
        request = NewSeparateMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SeparateMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewSeparateMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMaliciousRequestRequest() (request *TrustMaliciousRequestRequest) {
    request = &TrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMaliciousRequest")
    
    
    return
}

func NewTrustMaliciousRequestResponse() (response *TrustMaliciousRequestResponse) {
    response = &TrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TrustMaliciousRequest
// This API is used to trust a malicious request.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) TrustMaliciousRequest(request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    return c.TrustMaliciousRequestWithContext(context.Background(), request)
}

// TrustMaliciousRequest
// This API is used to trust a malicious request.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) TrustMaliciousRequestWithContext(ctx context.Context, request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewTrustMaliciousRequestRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrustMaliciousRequest require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
    request = &TrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMalwares")
    
    
    return
}

func NewTrustMalwaresResponse() (response *TrustMalwaresResponse) {
    response = &TrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TrustMalwares
// This API is used to trust an identified trojan file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    return c.TrustMalwaresWithContext(context.Background(), request)
}

// TrustMalwares
// This API is used to trust an identified trojan file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) TrustMalwaresWithContext(ctx context.Context, request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrustMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMaliciousRequestRequest() (request *UntrustMaliciousRequestRequest) {
    request = &UntrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMaliciousRequest")
    
    
    return
}

func NewUntrustMaliciousRequestResponse() (response *UntrustMaliciousRequestResponse) {
    response = &UntrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UntrustMaliciousRequest
// This API is used to untrust a malicious request.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMaliciousRequest(request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    return c.UntrustMaliciousRequestWithContext(context.Background(), request)
}

// UntrustMaliciousRequest
// This API is used to untrust a malicious request.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMaliciousRequestWithContext(ctx context.Context, request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewUntrustMaliciousRequestRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UntrustMaliciousRequest require credential")
    }

    request.SetContext(ctx)
    
    response = NewUntrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
    request = &UntrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMalwares")
    
    
    return
}

func NewUntrustMalwaresResponse() (response *UntrustMalwaresResponse) {
    response = &UntrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UntrustMalwares
// This API is used to untrust a trojan file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) UntrustMalwares(request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    return c.UntrustMalwaresWithContext(context.Background(), request)
}

// UntrustMalwares
// This API is used to untrust a trojan file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) UntrustMalwaresWithContext(ctx context.Context, request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    if request == nil {
        request = NewUntrustMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UntrustMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewUntrustMalwaresResponse()
    err = c.Send(request, response)
    return
}
