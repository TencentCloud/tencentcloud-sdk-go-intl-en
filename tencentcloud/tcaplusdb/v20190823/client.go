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

package v20190823

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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


func NewClearTablesRequest() (request *ClearTablesRequest) {
    request = &ClearTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ClearTables")
    
    
    return
}

func NewClearTablesResponse() (response *ClearTablesResponse) {
    response = &ClearTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearTables
// This API is used to clear table data based on the specified table information.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ClearTables(request *ClearTablesRequest) (response *ClearTablesResponse, err error) {
    return c.ClearTablesWithContext(context.Background(), request)
}

// ClearTables
// This API is used to clear table data based on the specified table information.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ClearTablesWithContext(ctx context.Context, request *ClearTablesRequest) (response *ClearTablesResponse, err error) {
    if request == nil {
        request = NewClearTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearTablesResponse()
    err = c.Send(request, response)
    return
}

func NewCompareIdlFilesRequest() (request *CompareIdlFilesRequest) {
    request = &CompareIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CompareIdlFiles")
    
    
    return
}

func NewCompareIdlFilesResponse() (response *CompareIdlFilesResponse) {
    response = &CompareIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CompareIdlFiles
// This API is used to select a target table, upload and verify the table modification file, and return the result of whether the table structure is allowed to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CompareIdlFiles(request *CompareIdlFilesRequest) (response *CompareIdlFilesResponse, err error) {
    return c.CompareIdlFilesWithContext(context.Background(), request)
}

// CompareIdlFiles
// This API is used to select a target table, upload and verify the table modification file, and return the result of whether the table structure is allowed to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CompareIdlFilesWithContext(ctx context.Context, request *CompareIdlFilesRequest) (response *CompareIdlFilesResponse, err error) {
    if request == nil {
        request = NewCompareIdlFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompareIdlFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompareIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateBackup")
    
    
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackup
// This API is used to create a backup task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// This API is used to create a backup task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// This API is used to create a TcaplusDB cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  INVALIDPARAMETERVALUE_UNSUPPORTIDLTYPE = "InvalidParameterValue.UnsupportIdlType"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCEINSUFFICIENT_NOAVAILABLECLUSTER = "ResourceInsufficient.NoAvailableCluster"
//  RESOURCEINSUFFICIENT_NOENOUGHVIPINVPC = "ResourceInsufficient.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_DUPLICATECLUSTERNAME = "ResourceUnavailable.DuplicateClusterName"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// This API is used to create a TcaplusDB cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  INVALIDPARAMETERVALUE_UNSUPPORTIDLTYPE = "InvalidParameterValue.UnsupportIdlType"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCEINSUFFICIENT_NOAVAILABLECLUSTER = "ResourceInsufficient.NoAvailableCluster"
//  RESOURCEINSUFFICIENT_NOENOUGHVIPINVPC = "ResourceInsufficient.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_DUPLICATECLUSTERNAME = "ResourceUnavailable.DuplicateClusterName"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotsRequest() (request *CreateSnapshotsRequest) {
    request = &CreateSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateSnapshots")
    
    
    return
}

func NewCreateSnapshotsResponse() (response *CreateSnapshotsResponse) {
    response = &CreateSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSnapshots
// This API is used to create one or more table snapshots at a specified past time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSnapshots(request *CreateSnapshotsRequest) (response *CreateSnapshotsResponse, err error) {
    return c.CreateSnapshotsWithContext(context.Background(), request)
}

// CreateSnapshots
// This API is used to create one or more table snapshots at a specified past time point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSnapshotsWithContext(ctx context.Context, request *CreateSnapshotsRequest) (response *CreateSnapshotsResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableGroupRequest() (request *CreateTableGroupRequest) {
    request = &CreateTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTableGroup")
    
    
    return
}

func NewCreateTableGroupResponse() (response *CreateTableGroupResponse) {
    response = &CreateTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTableGroup
// This API is used to create a table group in a TcaplusDB cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPINFO = "ResourceUnavailable.DuplicateTableGroupInfo"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPNAME = "ResourceUnavailable.DuplicateTableGroupName"
//  RESOURCEUNAVAILABLE_NOAVAILABLETABLEGROUP = "ResourceUnavailable.NoAvailableTableGroup"
func (c *Client) CreateTableGroup(request *CreateTableGroupRequest) (response *CreateTableGroupResponse, err error) {
    return c.CreateTableGroupWithContext(context.Background(), request)
}

// CreateTableGroup
// This API is used to create a table group in a TcaplusDB cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPINFO = "ResourceUnavailable.DuplicateTableGroupInfo"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPNAME = "ResourceUnavailable.DuplicateTableGroupName"
//  RESOURCEUNAVAILABLE_NOAVAILABLETABLEGROUP = "ResourceUnavailable.NoAvailableTableGroup"
func (c *Client) CreateTableGroupWithContext(ctx context.Context, request *CreateTableGroupRequest) (response *CreateTableGroupResponse, err error) {
    if request == nil {
        request = NewCreateTableGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTableGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTablesRequest() (request *CreateTablesRequest) {
    request = &CreateTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTables")
    
    
    return
}

func NewCreateTablesResponse() (response *CreateTablesResponse) {
    response = &CreateTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTables
// This API is used to create tables in batches based on the selected IDL file list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTables(request *CreateTablesRequest) (response *CreateTablesResponse, err error) {
    return c.CreateTablesWithContext(context.Background(), request)
}

// CreateTables
// This API is used to create tables in batches based on the selected IDL file list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTablesWithContext(ctx context.Context, request *CreateTablesRequest) (response *CreateTablesResponse, err error) {
    if request == nil {
        request = NewCreateTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRecordsRequest() (request *DeleteBackupRecordsRequest) {
    request = &DeleteBackupRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteBackupRecords")
    
    
    return
}

func NewDeleteBackupRecordsResponse() (response *DeleteBackupRecordsResponse) {
    response = &DeleteBackupRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackupRecords
// This API is used to delete a manual backup.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBackupRecords(request *DeleteBackupRecordsRequest) (response *DeleteBackupRecordsResponse, err error) {
    return c.DeleteBackupRecordsWithContext(context.Background(), request)
}

// DeleteBackupRecords
// This API is used to delete a manual backup.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBackupRecordsWithContext(ctx context.Context, request *DeleteBackupRecordsRequest) (response *DeleteBackupRecordsResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCluster
// This API is used to delete a TcaplusDB cluster, which will succeed only after all resources (including table groups and tables) in the cluster are released.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// This API is used to delete a TcaplusDB cluster, which will succeed only after all resources (including table groups and tables) in the cluster are released.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIdlFilesRequest() (request *DeleteIdlFilesRequest) {
    request = &DeleteIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteIdlFiles")
    
    
    return
}

func NewDeleteIdlFilesResponse() (response *DeleteIdlFilesResponse) {
    response = &DeleteIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIdlFiles
// This API is used to delete a target IDL file by specifying the cluster ID and information of the file to be deleted. If the file is associated with a table, deletion will fail.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteIdlFiles(request *DeleteIdlFilesRequest) (response *DeleteIdlFilesResponse, err error) {
    return c.DeleteIdlFilesWithContext(context.Background(), request)
}

// DeleteIdlFiles
// This API is used to delete a target IDL file by specifying the cluster ID and information of the file to be deleted. If the file is associated with a table, deletion will fail.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteIdlFilesWithContext(ctx context.Context, request *DeleteIdlFilesRequest) (response *DeleteIdlFilesResponse, err error) {
    if request == nil {
        request = NewDeleteIdlFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIdlFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSnapshots
// This API is used to delete one or more table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// This API is used to delete one or more table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSnapshotsWithContext(ctx context.Context, request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableDataFlowRequest() (request *DeleteTableDataFlowRequest) {
    request = &DeleteTableDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableDataFlow")
    
    
    return
}

func NewDeleteTableDataFlowResponse() (response *DeleteTableDataFlowResponse) {
    response = &DeleteTableDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTableDataFlow
// This API is used to disable data subscription for tables.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableDataFlow(request *DeleteTableDataFlowRequest) (response *DeleteTableDataFlowResponse, err error) {
    return c.DeleteTableDataFlowWithContext(context.Background(), request)
}

// DeleteTableDataFlow
// This API is used to disable data subscription for tables.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableDataFlowWithContext(ctx context.Context, request *DeleteTableDataFlowRequest) (response *DeleteTableDataFlowResponse, err error) {
    if request == nil {
        request = NewDeleteTableDataFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTableDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableGroupRequest() (request *DeleteTableGroupRequest) {
    request = &DeleteTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableGroup")
    
    
    return
}

func NewDeleteTableGroupResponse() (response *DeleteTableGroupResponse) {
    response = &DeleteTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTableGroup
// This API is used to delete a table group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableGroup(request *DeleteTableGroupRequest) (response *DeleteTableGroupResponse, err error) {
    return c.DeleteTableGroupWithContext(context.Background(), request)
}

// DeleteTableGroup
// This API is used to delete a table group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableGroupWithContext(ctx context.Context, request *DeleteTableGroupRequest) (response *DeleteTableGroupResponse, err error) {
    if request == nil {
        request = NewDeleteTableGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTableGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableIndexRequest() (request *DeleteTableIndexRequest) {
    request = &DeleteTableIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableIndex")
    
    
    return
}

func NewDeleteTableIndexResponse() (response *DeleteTableIndexResponse) {
    response = &DeleteTableIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTableIndex
// This API is used to delete the global index from a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableIndex(request *DeleteTableIndexRequest) (response *DeleteTableIndexResponse, err error) {
    return c.DeleteTableIndexWithContext(context.Background(), request)
}

// DeleteTableIndex
// This API is used to delete the global index from a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTableIndexWithContext(ctx context.Context, request *DeleteTableIndexRequest) (response *DeleteTableIndexResponse, err error) {
    if request == nil {
        request = NewDeleteTableIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTableIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTablesRequest() (request *DeleteTablesRequest) {
    request = &DeleteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTables")
    
    
    return
}

func NewDeleteTablesResponse() (response *DeleteTablesResponse) {
    response = &DeleteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTables
// This API is used to drop a specified table. Calling this API for the first time means to move the table to the recycle bin, while calling it again means to drop the table completely from the recycle bin.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTables(request *DeleteTablesRequest) (response *DeleteTablesResponse, err error) {
    return c.DeleteTablesWithContext(context.Background(), request)
}

// DeleteTables
// This API is used to drop a specified table. Calling this API for the first time means to move the table to the recycle bin, while calling it again means to drop the table completely from the recycle bin.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTablesWithContext(ctx context.Context, request *DeleteTablesRequest) (response *DeleteTablesResponse, err error) {
    if request == nil {
        request = NewDeleteTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupRecordsRequest() (request *DescribeBackupRecordsRequest) {
    request = &DescribeBackupRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeBackupRecords")
    
    
    return
}

func NewDescribeBackupRecordsResponse() (response *DescribeBackupRecordsResponse) {
    response = &DescribeBackupRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupRecords
// This API is used to query backup records.
//
// 
//
// When querying the cluster level, set `TableGroupId` to `-1` and `TableName` to `-1`.
//
// When querying the cluster and table group levels, set `TableName` to `-1`.
//
// When querying the cluster, table group, and table levels, both `TableGroupId` and `TableName` cannot be set to `-1`.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBackupRecords(request *DescribeBackupRecordsRequest) (response *DescribeBackupRecordsResponse, err error) {
    return c.DescribeBackupRecordsWithContext(context.Background(), request)
}

// DescribeBackupRecords
// This API is used to query backup records.
//
// 
//
// When querying the cluster level, set `TableGroupId` to `-1` and `TableName` to `-1`.
//
// When querying the cluster and table group levels, set `TableName` to `-1`.
//
// When querying the cluster, table group, and table levels, both `TableGroupId` and `TableName` cannot be set to `-1`.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBackupRecordsWithContext(ctx context.Context, request *DescribeBackupRecordsRequest) (response *DescribeBackupRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterTagsRequest() (request *DescribeClusterTagsRequest) {
    request = &DescribeClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeClusterTags")
    
    
    return
}

func NewDescribeClusterTagsResponse() (response *DescribeClusterTagsResponse) {
    response = &DescribeClusterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterTags
// This API is used to get the associated tag list of a cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClusterTags(request *DescribeClusterTagsRequest) (response *DescribeClusterTagsResponse, err error) {
    return c.DescribeClusterTagsWithContext(context.Background(), request)
}

// DescribeClusterTags
// This API is used to get the associated tag list of a cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClusterTagsWithContext(ctx context.Context, request *DescribeClusterTagsRequest) (response *DescribeClusterTagsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// This API is used to query the TcaplusDB cluster list, including cluster details.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// This API is used to query the TcaplusDB cluster list, including cluster details.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdlFileInfosRequest() (request *DescribeIdlFileInfosRequest) {
    request = &DescribeIdlFileInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeIdlFileInfos")
    
    
    return
}

func NewDescribeIdlFileInfosResponse() (response *DescribeIdlFileInfosResponse) {
    response = &DescribeIdlFileInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdlFileInfos
// This API is used to query table description file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdlFileInfos(request *DescribeIdlFileInfosRequest) (response *DescribeIdlFileInfosResponse, err error) {
    return c.DescribeIdlFileInfosWithContext(context.Background(), request)
}

// DescribeIdlFileInfos
// This API is used to query table description file details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdlFileInfosWithContext(ctx context.Context, request *DescribeIdlFileInfosRequest) (response *DescribeIdlFileInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIdlFileInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdlFileInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdlFileInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRequest() (request *DescribeMachineRequest) {
    request = &DescribeMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeMachine")
    
    
    return
}

func NewDescribeMachineResponse() (response *DescribeMachineResponse) {
    response = &DescribeMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachine
// This API is used to query the available machines in a dedicated cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMachine(request *DescribeMachineRequest) (response *DescribeMachineResponse, err error) {
    return c.DescribeMachineWithContext(context.Background(), request)
}

// DescribeMachine
// This API is used to query the available machines in a dedicated cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMachineWithContext(ctx context.Context, request *DescribeMachineRequest) (response *DescribeMachineResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// This API is used to query the list of regions supported by the TcaplusDB service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to query the list of regions supported by the TcaplusDB service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshots
// This API is used to query the list of table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPNAME = "InvalidParameterValue.InvalidAppName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// This API is used to query the list of table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPNAME = "InvalidParameterValue.InvalidAppName"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableGroupTagsRequest() (request *DescribeTableGroupTagsRequest) {
    request = &DescribeTableGroupTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableGroupTags")
    
    
    return
}

func NewDescribeTableGroupTagsResponse() (response *DescribeTableGroupTagsResponse) {
    response = &DescribeTableGroupTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableGroupTags
// This API is used to get the associated tag list of a table group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableGroupTags(request *DescribeTableGroupTagsRequest) (response *DescribeTableGroupTagsResponse, err error) {
    return c.DescribeTableGroupTagsWithContext(context.Background(), request)
}

// DescribeTableGroupTags
// This API is used to get the associated tag list of a table group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableGroupTagsWithContext(ctx context.Context, request *DescribeTableGroupTagsRequest) (response *DescribeTableGroupTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTableGroupTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableGroupTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableGroupTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableGroupsRequest() (request *DescribeTableGroupsRequest) {
    request = &DescribeTableGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableGroups")
    
    
    return
}

func NewDescribeTableGroupsResponse() (response *DescribeTableGroupsResponse) {
    response = &DescribeTableGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableGroups
// This API is used to query the table group list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableGroups(request *DescribeTableGroupsRequest) (response *DescribeTableGroupsResponse, err error) {
    return c.DescribeTableGroupsWithContext(context.Background(), request)
}

// DescribeTableGroups
// This API is used to query the table group list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableGroupsWithContext(ctx context.Context, request *DescribeTableGroupsRequest) (response *DescribeTableGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTableGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableTagsRequest() (request *DescribeTableTagsRequest) {
    request = &DescribeTableTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableTags")
    
    
    return
}

func NewDescribeTableTagsResponse() (response *DescribeTableTagsResponse) {
    response = &DescribeTableTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableTags
// This API is used to get table tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableTags(request *DescribeTableTagsRequest) (response *DescribeTableTagsResponse, err error) {
    return c.DescribeTableTagsWithContext(context.Background(), request)
}

// DescribeTableTags
// This API is used to get table tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableTagsWithContext(ctx context.Context, request *DescribeTableTagsRequest) (response *DescribeTableTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTableTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTables
// This API is used to query table details.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// This API is used to query table details.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesInRecycleRequest() (request *DescribeTablesInRecycleRequest) {
    request = &DescribeTablesInRecycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTablesInRecycle")
    
    
    return
}

func NewDescribeTablesInRecycleResponse() (response *DescribeTablesInRecycleResponse) {
    response = &DescribeTablesInRecycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTablesInRecycle
// This API is used to query the details of a table in recycle bin.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTablesInRecycle(request *DescribeTablesInRecycleRequest) (response *DescribeTablesInRecycleResponse, err error) {
    return c.DescribeTablesInRecycleWithContext(context.Background(), request)
}

// DescribeTablesInRecycle
// This API is used to query the details of a table in recycle bin.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTablesInRecycleWithContext(ctx context.Context, request *DescribeTablesInRecycleRequest) (response *DescribeTablesInRecycleResponse, err error) {
    if request == nil {
        request = NewDescribeTablesInRecycleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTablesInRecycle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesInRecycleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTasks
// This API is used to query the task list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// This API is used to query the task list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUinInWhitelistRequest() (request *DescribeUinInWhitelistRequest) {
    request = &DescribeUinInWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeUinInWhitelist")
    
    
    return
}

func NewDescribeUinInWhitelistResponse() (response *DescribeUinInWhitelistResponse) {
    response = &DescribeUinInWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUinInWhitelist
// This API is used to query whether the current user is in the allowlist and control whether the user can create TDR-type apps or tables.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUinInWhitelist(request *DescribeUinInWhitelistRequest) (response *DescribeUinInWhitelistResponse, err error) {
    return c.DescribeUinInWhitelistWithContext(context.Background(), request)
}

// DescribeUinInWhitelist
// This API is used to query whether the current user is in the allowlist and control whether the user can create TDR-type apps or tables.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUinInWhitelistWithContext(ctx context.Context, request *DescribeUinInWhitelistRequest) (response *DescribeUinInWhitelistResponse, err error) {
    if request == nil {
        request = NewDescribeUinInWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUinInWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUinInWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRestProxyRequest() (request *DisableRestProxyRequest) {
    request = &DisableRestProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DisableRestProxy")
    
    
    return
}

func NewDisableRestProxyResponse() (response *DisableRestProxyResponse) {
    response = &DisableRestProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableRestProxy
// This API is used to disable the RESTful API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableRestProxy(request *DisableRestProxyRequest) (response *DisableRestProxyResponse, err error) {
    return c.DisableRestProxyWithContext(context.Background(), request)
}

// DisableRestProxy
// This API is used to disable the RESTful API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableRestProxyWithContext(ctx context.Context, request *DisableRestProxyRequest) (response *DisableRestProxyResponse, err error) {
    if request == nil {
        request = NewDisableRestProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableRestProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableRestProxyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRestProxyRequest() (request *EnableRestProxyRequest) {
    request = &EnableRestProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "EnableRestProxy")
    
    
    return
}

func NewEnableRestProxyResponse() (response *EnableRestProxyResponse) {
    response = &EnableRestProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableRestProxy
// This API is used to enable the RESTful API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPNAME = "InvalidParameterValue.InvalidAppName"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEVALUE = "InvalidParameterValue.InvalidTimeValue"
//  INVALIDPARAMETERVALUE_INVALIDZONENAME = "InvalidParameterValue.InvalidZoneName"
//  INVALIDPARAMETERVALUE_UNSUPPORTIDLTYPE = "InvalidParameterValue.UnsupportIdlType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableRestProxy(request *EnableRestProxyRequest) (response *EnableRestProxyResponse, err error) {
    return c.EnableRestProxyWithContext(context.Background(), request)
}

// EnableRestProxy
// This API is used to enable the RESTful API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPNAME = "InvalidParameterValue.InvalidAppName"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEVALUE = "InvalidParameterValue.InvalidTimeValue"
//  INVALIDPARAMETERVALUE_INVALIDZONENAME = "InvalidParameterValue.InvalidZoneName"
//  INVALIDPARAMETERVALUE_UNSUPPORTIDLTYPE = "InvalidParameterValue.UnsupportIdlType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableRestProxyWithContext(ctx context.Context, request *EnableRestProxyRequest) (response *EnableRestProxyResponse, err error) {
    if request == nil {
        request = NewEnableRestProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableRestProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableRestProxyResponse()
    err = c.Send(request, response)
    return
}

func NewImportSnapshotsRequest() (request *ImportSnapshotsRequest) {
    request = &ImportSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ImportSnapshots")
    
    
    return
}

func NewImportSnapshotsResponse() (response *ImportSnapshotsResponse) {
    response = &ImportSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportSnapshots
// This API is used to import a snapshot into a new table or the original table from which the snapshot was created.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportSnapshots(request *ImportSnapshotsRequest) (response *ImportSnapshotsResponse, err error) {
    return c.ImportSnapshotsWithContext(context.Background(), request)
}

// ImportSnapshots
// This API is used to import a snapshot into a new table or the original table from which the snapshot was created.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportSnapshotsWithContext(ctx context.Context, request *ImportSnapshotsRequest) (response *ImportSnapshotsResponse, err error) {
    if request == nil {
        request = NewImportSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewMergeTablesDataRequest() (request *MergeTablesDataRequest) {
    request = &MergeTablesDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "MergeTablesData")
    
    
    return
}

func NewMergeTablesDataResponse() (response *MergeTablesDataResponse) {
    response = &MergeTablesDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MergeTablesData
// This API is used to merge tables.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) MergeTablesData(request *MergeTablesDataRequest) (response *MergeTablesDataResponse, err error) {
    return c.MergeTablesDataWithContext(context.Background(), request)
}

// MergeTablesData
// This API is used to merge tables.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) MergeTablesDataWithContext(ctx context.Context, request *MergeTablesDataRequest) (response *MergeTablesDataResponse, err error) {
    if request == nil {
        request = NewMergeTablesDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MergeTablesData require credential")
    }

    request.SetContext(ctx)
    
    response = NewMergeTablesDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCensorshipRequest() (request *ModifyCensorshipRequest) {
    request = &ModifyCensorshipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyCensorship")
    
    
    return
}

func NewModifyCensorshipResponse() (response *ModifyCensorshipResponse) {
    response = &ModifyCensorshipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCensorship
// This API is used to enable or disable the cluster operation approval feature.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCensorship(request *ModifyCensorshipRequest) (response *ModifyCensorshipResponse, err error) {
    return c.ModifyCensorshipWithContext(context.Background(), request)
}

// ModifyCensorship
// This API is used to enable or disable the cluster operation approval feature.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCensorshipWithContext(ctx context.Context, request *ModifyCensorshipRequest) (response *ModifyCensorshipResponse, err error) {
    if request == nil {
        request = NewModifyCensorshipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCensorship require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCensorshipResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterMachineRequest() (request *ModifyClusterMachineRequest) {
    request = &ModifyClusterMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterMachine")
    
    
    return
}

func NewModifyClusterMachineResponse() (response *ModifyClusterMachineResponse) {
    response = &ModifyClusterMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterMachine
// This API is used to modify the machines of a dedicated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_NOAVAILABLEAPP = "ResourceInsufficient.NoAvailableApp"
//  RESOURCEINSUFFICIENT_NOAVAILABLECLUSTER = "ResourceInsufficient.NoAvailableCluster"
func (c *Client) ModifyClusterMachine(request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
    return c.ModifyClusterMachineWithContext(context.Background(), request)
}

// ModifyClusterMachine
// This API is used to modify the machines of a dedicated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_NOAVAILABLEAPP = "ResourceInsufficient.NoAvailableApp"
//  RESOURCEINSUFFICIENT_NOAVAILABLECLUSTER = "ResourceInsufficient.NoAvailableCluster"
func (c *Client) ModifyClusterMachineWithContext(ctx context.Context, request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
    if request == nil {
        request = NewModifyClusterMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterMachineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterName")
    
    
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterName
// This API is used to rename a specified cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATECLUSTERNAME = "ResourceUnavailable.DuplicateClusterName"
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    return c.ModifyClusterNameWithContext(context.Background(), request)
}

// ModifyClusterName
// This API is used to rename a specified cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATECLUSTERNAME = "ResourceUnavailable.DuplicateClusterName"
func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterPasswordRequest() (request *ModifyClusterPasswordRequest) {
    request = &ModifyClusterPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterPassword")
    
    
    return
}

func NewModifyClusterPasswordResponse() (response *ModifyClusterPasswordResponse) {
    response = &ModifyClusterPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterPassword
// This API is used to change the password of a specified cluster. The backend will allow the TcaplusDB SDK to access databases with both old and new passwords before the old password expires. You cannot submit a new password change request before the old password expires or submit a request to modify the expiration time of the old password after the old password expires.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OLDPASSWORDHASEXPIRED = "FailedOperation.OldPasswordHasExpired"
//  FAILEDOPERATION_OLDPASSWORDINUSE = "FailedOperation.OldPasswordInUse"
//  FAILEDOPERATION_PASSWORDFAILURE = "FailedOperation.PasswordFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTIMEVALUE = "InvalidParameterValue.InvalidTimeValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterPassword(request *ModifyClusterPasswordRequest) (response *ModifyClusterPasswordResponse, err error) {
    return c.ModifyClusterPasswordWithContext(context.Background(), request)
}

// ModifyClusterPassword
// This API is used to change the password of a specified cluster. The backend will allow the TcaplusDB SDK to access databases with both old and new passwords before the old password expires. You cannot submit a new password change request before the old password expires or submit a request to modify the expiration time of the old password after the old password expires.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OLDPASSWORDHASEXPIRED = "FailedOperation.OldPasswordHasExpired"
//  FAILEDOPERATION_OLDPASSWORDINUSE = "FailedOperation.OldPasswordInUse"
//  FAILEDOPERATION_PASSWORDFAILURE = "FailedOperation.PasswordFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTIMEVALUE = "InvalidParameterValue.InvalidTimeValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterPasswordWithContext(ctx context.Context, request *ModifyClusterPasswordRequest) (response *ModifyClusterPasswordResponse, err error) {
    if request == nil {
        request = NewModifyClusterPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterTagsRequest() (request *ModifyClusterTagsRequest) {
    request = &ModifyClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterTags")
    
    
    return
}

func NewModifyClusterTagsResponse() (response *ModifyClusterTagsResponse) {
    response = &ModifyClusterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterTags
// This API is used to modify cluster tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClusterTags(request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
    return c.ModifyClusterTagsWithContext(context.Background(), request)
}

// ModifyClusterTags
// This API is used to modify cluster tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClusterTagsWithContext(ctx context.Context, request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
    if request == nil {
        request = NewModifyClusterTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotsRequest() (request *ModifySnapshotsRequest) {
    request = &ModifySnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifySnapshots")
    
    
    return
}

func NewModifySnapshotsResponse() (response *ModifySnapshotsResponse) {
    response = &ModifySnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySnapshots
// This API is used to modify the expiration time of one or more table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySnapshots(request *ModifySnapshotsRequest) (response *ModifySnapshotsResponse, err error) {
    return c.ModifySnapshotsWithContext(context.Background(), request)
}

// ModifySnapshots
// This API is used to modify the expiration time of one or more table snapshots.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySnapshotsWithContext(ctx context.Context, request *ModifySnapshotsRequest) (response *ModifySnapshotsResponse, err error) {
    if request == nil {
        request = NewModifySnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableGroupNameRequest() (request *ModifyTableGroupNameRequest) {
    request = &ModifyTableGroupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableGroupName")
    
    
    return
}

func NewModifyTableGroupNameResponse() (response *ModifyTableGroupNameResponse) {
    response = &ModifyTableGroupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableGroupName
// This API is used to rename a TcaplusDB table group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPNAME = "ResourceUnavailable.DuplicateTableGroupName"
func (c *Client) ModifyTableGroupName(request *ModifyTableGroupNameRequest) (response *ModifyTableGroupNameResponse, err error) {
    return c.ModifyTableGroupNameWithContext(context.Background(), request)
}

// ModifyTableGroupName
// This API is used to rename a TcaplusDB table group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPNAME = "ResourceUnavailable.DuplicateTableGroupName"
func (c *Client) ModifyTableGroupNameWithContext(ctx context.Context, request *ModifyTableGroupNameRequest) (response *ModifyTableGroupNameResponse, err error) {
    if request == nil {
        request = NewModifyTableGroupNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableGroupName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableGroupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableGroupTagsRequest() (request *ModifyTableGroupTagsRequest) {
    request = &ModifyTableGroupTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableGroupTags")
    
    
    return
}

func NewModifyTableGroupTagsResponse() (response *ModifyTableGroupTagsResponse) {
    response = &ModifyTableGroupTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableGroupTags
// This API is used to modify table group tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableGroupTags(request *ModifyTableGroupTagsRequest) (response *ModifyTableGroupTagsResponse, err error) {
    return c.ModifyTableGroupTagsWithContext(context.Background(), request)
}

// ModifyTableGroupTags
// This API is used to modify table group tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableGroupTagsWithContext(ctx context.Context, request *ModifyTableGroupTagsRequest) (response *ModifyTableGroupTagsResponse, err error) {
    if request == nil {
        request = NewModifyTableGroupTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableGroupTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableGroupTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableMemosRequest() (request *ModifyTableMemosRequest) {
    request = &ModifyTableMemosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableMemos")
    
    
    return
}

func NewModifyTableMemosResponse() (response *ModifyTableMemosResponse) {
    response = &ModifyTableMemosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableMemos
// This API is used to modify table remarks.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableMemos(request *ModifyTableMemosRequest) (response *ModifyTableMemosResponse, err error) {
    return c.ModifyTableMemosWithContext(context.Background(), request)
}

// ModifyTableMemos
// This API is used to modify table remarks.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableMemosWithContext(ctx context.Context, request *ModifyTableMemosRequest) (response *ModifyTableMemosResponse, err error) {
    if request == nil {
        request = NewModifyTableMemosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableMemos require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableMemosResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableQuotasRequest() (request *ModifyTableQuotasRequest) {
    request = &ModifyTableQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableQuotas")
    
    
    return
}

func NewModifyTableQuotasResponse() (response *ModifyTableQuotasResponse) {
    response = &ModifyTableQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableQuotas
// This API is used to scale a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableQuotas(request *ModifyTableQuotasRequest) (response *ModifyTableQuotasResponse, err error) {
    return c.ModifyTableQuotasWithContext(context.Background(), request)
}

// ModifyTableQuotas
// This API is used to scale a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableQuotasWithContext(ctx context.Context, request *ModifyTableQuotasRequest) (response *ModifyTableQuotasResponse, err error) {
    if request == nil {
        request = NewModifyTableQuotasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableQuotas require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableTagsRequest() (request *ModifyTableTagsRequest) {
    request = &ModifyTableTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableTags")
    
    
    return
}

func NewModifyTableTagsResponse() (response *ModifyTableTagsResponse) {
    response = &ModifyTableTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableTags
// This API is used to modify table tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableTags(request *ModifyTableTagsRequest) (response *ModifyTableTagsResponse, err error) {
    return c.ModifyTableTagsWithContext(context.Background(), request)
}

// ModifyTableTags
// This API is used to modify table tags.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTableTagsWithContext(ctx context.Context, request *ModifyTableTagsRequest) (response *ModifyTableTagsResponse, err error) {
    if request == nil {
        request = NewModifyTableTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTablesRequest() (request *ModifyTablesRequest) {
    request = &ModifyTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTables")
    
    
    return
}

func NewModifyTablesResponse() (response *ModifyTablesResponse) {
    response = &ModifyTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTables
// This API is used to modify specified tables in batches based on the selected table definition IDL file.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTables(request *ModifyTablesRequest) (response *ModifyTablesResponse, err error) {
    return c.ModifyTablesWithContext(context.Background(), request)
}

// ModifyTables
// This API is used to modify specified tables in batches based on the selected table definition IDL file.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTablesWithContext(ctx context.Context, request *ModifyTablesRequest) (response *ModifyTablesResponse, err error) {
    if request == nil {
        request = NewModifyTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverRecycleTablesRequest() (request *RecoverRecycleTablesRequest) {
    request = &RecoverRecycleTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RecoverRecycleTables")
    
    
    return
}

func NewRecoverRecycleTablesResponse() (response *RecoverRecycleTablesResponse) {
    response = &RecoverRecycleTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverRecycleTables
// This API is used to recover a dropped table from the recycle bin. It will not work for tables to be released due to arrears.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RecoverRecycleTables(request *RecoverRecycleTablesRequest) (response *RecoverRecycleTablesResponse, err error) {
    return c.RecoverRecycleTablesWithContext(context.Background(), request)
}

// RecoverRecycleTables
// This API is used to recover a dropped table from the recycle bin. It will not work for tables to be released due to arrears.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RecoverRecycleTablesWithContext(ctx context.Context, request *RecoverRecycleTablesRequest) (response *RecoverRecycleTablesResponse, err error) {
    if request == nil {
        request = NewRecoverRecycleTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverRecycleTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverRecycleTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackTablesRequest() (request *RollbackTablesRequest) {
    request = &RollbackTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RollbackTables")
    
    
    return
}

func NewRollbackTablesResponse() (response *RollbackTablesResponse) {
    response = &RollbackTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackTables
// This API is used to roll back table data.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RollbackTables(request *RollbackTablesRequest) (response *RollbackTablesResponse, err error) {
    return c.RollbackTablesWithContext(context.Background(), request)
}

// RollbackTables
// This API is used to roll back table data.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RollbackTablesWithContext(ctx context.Context, request *RollbackTablesRequest) (response *RollbackTablesResponse, err error) {
    if request == nil {
        request = NewRollbackTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackTablesResponse()
    err = c.Send(request, response)
    return
}

func NewSetBackupExpireRuleRequest() (request *SetBackupExpireRuleRequest) {
    request = &SetBackupExpireRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "SetBackupExpireRule")
    
    
    return
}

func NewSetBackupExpireRuleResponse() (response *SetBackupExpireRuleResponse) {
    response = &SetBackupExpireRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetBackupExpireRule
// This API is used to add/delete/modify backup expiration policy. `ClusterId` must be a specific cluster ID (appid).
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetBackupExpireRule(request *SetBackupExpireRuleRequest) (response *SetBackupExpireRuleResponse, err error) {
    return c.SetBackupExpireRuleWithContext(context.Background(), request)
}

// SetBackupExpireRule
// This API is used to add/delete/modify backup expiration policy. `ClusterId` must be a specific cluster ID (appid).
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetBackupExpireRuleWithContext(ctx context.Context, request *SetBackupExpireRuleRequest) (response *SetBackupExpireRuleResponse, err error) {
    if request == nil {
        request = NewSetBackupExpireRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetBackupExpireRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetBackupExpireRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSetTableDataFlowRequest() (request *SetTableDataFlowRequest) {
    request = &SetTableDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "SetTableDataFlow")
    
    
    return
}

func NewSetTableDataFlowResponse() (response *SetTableDataFlowResponse) {
    response = &SetTableDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetTableDataFlow
// This API is used to enable data subscription for tables or modify the feature's configurations.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetTableDataFlow(request *SetTableDataFlowRequest) (response *SetTableDataFlowResponse, err error) {
    return c.SetTableDataFlowWithContext(context.Background(), request)
}

// SetTableDataFlow
// This API is used to enable data subscription for tables or modify the feature's configurations.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetTableDataFlowWithContext(ctx context.Context, request *SetTableDataFlowRequest) (response *SetTableDataFlowResponse, err error) {
    if request == nil {
        request = NewSetTableDataFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTableDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTableDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewSetTableIndexRequest() (request *SetTableIndexRequest) {
    request = &SetTableIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "SetTableIndex")
    
    
    return
}

func NewSetTableIndexResponse() (response *SetTableIndexResponse) {
    response = &SetTableIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetTableIndex
// This API is used to create a global index for a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetTableIndex(request *SetTableIndexRequest) (response *SetTableIndexResponse, err error) {
    return c.SetTableIndexWithContext(context.Background(), request)
}

// SetTableIndex
// This API is used to create a global index for a table.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetTableIndexWithContext(ctx context.Context, request *SetTableIndexRequest) (response *SetTableIndexResponse, err error) {
    if request == nil {
        request = NewSetTableIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTableIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTableIndexResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApplyRequest() (request *UpdateApplyRequest) {
    request = &UpdateApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "UpdateApply")
    
    
    return
}

func NewUpdateApplyResponse() (response *UpdateApplyResponse) {
    response = &UpdateApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApply
// This API is used to update the application status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateApply(request *UpdateApplyRequest) (response *UpdateApplyResponse, err error) {
    return c.UpdateApplyWithContext(context.Background(), request)
}

// UpdateApply
// This API is used to update the application status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateApplyWithContext(ctx context.Context, request *UpdateApplyRequest) (response *UpdateApplyResponse, err error) {
    if request == nil {
        request = NewUpdateApplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApply require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApplyResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyIdlFilesRequest() (request *VerifyIdlFilesRequest) {
    request = &VerifyIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "VerifyIdlFiles")
    
    
    return
}

func NewVerifyIdlFilesResponse() (response *VerifyIdlFilesResponse) {
    response = &VerifyIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyIdlFiles
// This API is used to upload and verify a table creation file and return the definition of tables that are verified to be valid.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyIdlFiles(request *VerifyIdlFilesRequest) (response *VerifyIdlFilesResponse, err error) {
    return c.VerifyIdlFilesWithContext(context.Background(), request)
}

// VerifyIdlFiles
// This API is used to upload and verify a table creation file and return the definition of tables that are verified to be valid.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyIdlFilesWithContext(ctx context.Context, request *VerifyIdlFilesRequest) (response *VerifyIdlFilesResponse, err error) {
    if request == nil {
        request = NewVerifyIdlFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyIdlFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyIdlFilesResponse()
    err = c.Send(request, response)
    return
}
