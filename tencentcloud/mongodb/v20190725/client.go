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

package v20190725

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-25"

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


func NewAssignProjectRequest() (request *AssignProjectRequest) {
    request = &AssignProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "AssignProject")
    
    
    return
}

func NewAssignProjectResponse() (response *AssignProjectResponse) {
    response = &AssignProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssignProject
// This API is used to specify the project of a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
func (c *Client) AssignProject(request *AssignProjectRequest) (response *AssignProjectResponse, err error) {
    return c.AssignProjectWithContext(context.Background(), request)
}

// AssignProject
// This API is used to specify the project of a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
func (c *Client) AssignProjectWithContext(ctx context.Context, request *AssignProjectRequest) (response *AssignProjectResponse, err error) {
    if request == nil {
        request = NewAssignProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "AssignProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssignProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupDBInstanceRequest() (request *CreateBackupDBInstanceRequest) {
    request = &CreateBackupDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateBackupDBInstance")
    
    
    return
}

func NewCreateBackupDBInstanceResponse() (response *CreateBackupDBInstanceResponse) {
    response = &CreateBackupDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupDBInstance
// This API is used to back up an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSPARENTDATAENCRYPTIONALREADYOPEN = "FailedOperation.TransparentDataEncryptionAlreadyOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) CreateBackupDBInstance(request *CreateBackupDBInstanceRequest) (response *CreateBackupDBInstanceResponse, err error) {
    return c.CreateBackupDBInstanceWithContext(context.Background(), request)
}

// CreateBackupDBInstance
// This API is used to back up an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSPARENTDATAENCRYPTIONALREADYOPEN = "FailedOperation.TransparentDataEncryptionAlreadyOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) CreateBackupDBInstanceWithContext(ctx context.Context, request *CreateBackupDBInstanceRequest) (response *CreateBackupDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateBackupDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateBackupDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupDownloadTaskRequest() (request *CreateBackupDownloadTaskRequest) {
    request = &CreateBackupDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateBackupDownloadTask")
    
    
    return
}

func NewCreateBackupDownloadTaskResponse() (response *CreateBackupDownloadTaskResponse) {
    response = &CreateBackupDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupDownloadTask
// This API is used to create a backup download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateBackupDownloadTask(request *CreateBackupDownloadTaskRequest) (response *CreateBackupDownloadTaskResponse, err error) {
    return c.CreateBackupDownloadTaskWithContext(context.Background(), request)
}

// CreateBackupDownloadTask
// This API is used to create a backup download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateBackupDownloadTaskWithContext(ctx context.Context, request *CreateBackupDownloadTaskRequest) (response *CreateBackupDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateBackupDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateBackupDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstance")
    
    
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstance
// This API is used to create a yearly/monthly subscription TencentDB for MongoDB instance. The [DescribeSpecInfo](https://www.tencentcloud.comom/document/product/240/35767?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SECONDARYNUMERROR = "InvalidParameterValue.SecondaryNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    return c.CreateDBInstanceWithContext(context.Background(), request)
}

// CreateDBInstance
// This API is used to create a yearly/monthly subscription TencentDB for MongoDB instance. The [DescribeSpecInfo](https://www.tencentcloud.comom/document/product/240/35767?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SECONDARYNUMERROR = "InvalidParameterValue.SecondaryNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceHourRequest() (request *CreateDBInstanceHourRequest) {
    request = &CreateDBInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstanceHour")
    
    
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstanceHour
// This API is used to create a pay-as-you-go TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    return c.CreateDBInstanceHourWithContext(context.Background(), request)
}

// CreateDBInstanceHour
// This API is used to create a pay-as-you-go TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceHourWithContext(ctx context.Context, request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstanceHour")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsyncRequestInfo
// This API is used to query the asynchronous task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// This API is used to query the asynchronous task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAsyncRequestInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadTaskRequest() (request *DescribeBackupDownloadTaskRequest) {
    request = &DescribeBackupDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeBackupDownloadTask")
    
    
    return
}

func NewDescribeBackupDownloadTaskResponse() (response *DescribeBackupDownloadTaskResponse) {
    response = &DescribeBackupDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadTask
// This API is used to query information about the backup download task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeBackupDownloadTask(request *DescribeBackupDownloadTaskRequest) (response *DescribeBackupDownloadTaskResponse, err error) {
    return c.DescribeBackupDownloadTaskWithContext(context.Background(), request)
}

// DescribeBackupDownloadTask
// This API is used to query information about the backup download task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeBackupDownloadTaskWithContext(ctx context.Context, request *DescribeBackupDownloadTaskRequest) (response *DescribeBackupDownloadTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeBackupDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientConnectionsRequest() (request *DescribeClientConnectionsRequest) {
    request = &DescribeClientConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeClientConnections")
    
    
    return
}

func NewDescribeClientConnectionsResponse() (response *DescribeClientConnectionsResponse) {
    response = &DescribeClientConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClientConnections
// This API is used to query the client connection information on an instance, including the IP address for connection and the number of connections.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MONGOVERSIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.MongoVersionNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROXYNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.ProxyNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.RegionNotSupportQueryClient"
func (c *Client) DescribeClientConnections(request *DescribeClientConnectionsRequest) (response *DescribeClientConnectionsResponse, err error) {
    return c.DescribeClientConnectionsWithContext(context.Background(), request)
}

// DescribeClientConnections
// This API is used to query the client connection information on an instance, including the IP address for connection and the number of connections.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MONGOVERSIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.MongoVersionNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROXYNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.ProxyNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.RegionNotSupportQueryClient"
func (c *Client) DescribeClientConnectionsWithContext(ctx context.Context, request *DescribeClientConnectionsRequest) (response *DescribeClientConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeClientConnectionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeClientConnections")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBBackups")
    
    
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBBackups
// This API is used to query the list of instance backups. Currently, only backups created in the last seven days can be queried.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    return c.DescribeDBBackupsWithContext(context.Background(), request)
}

// DescribeDBBackups
// This API is used to query the list of instance backups. Currently, only backups created in the last seven days can be queried.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceDealRequest() (request *DescribeDBInstanceDealRequest) {
    request = &DescribeDBInstanceDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceDeal")
    
    
    return
}

func NewDescribeDBInstanceDealResponse() (response *DescribeDBInstanceDealResponse) {
    response = &DescribeDBInstanceDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceDeal
// This API is used to get order details of purchase, renewal, and specification adjustment of a MongoDB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceDeal(request *DescribeDBInstanceDealRequest) (response *DescribeDBInstanceDealResponse, err error) {
    return c.DescribeDBInstanceDealWithContext(context.Background(), request)
}

// DescribeDBInstanceDeal
// This API is used to get order details of purchase, renewal, and specification adjustment of a MongoDB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceDealWithContext(ctx context.Context, request *DescribeDBInstanceDealRequest) (response *DescribeDBInstanceDealResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDealRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceDeal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceDealResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceNamespaceRequest() (request *DescribeDBInstanceNamespaceRequest) {
    request = &DescribeDBInstanceNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceNamespace")
    
    
    return
}

func NewDescribeDBInstanceNamespaceResponse() (response *DescribeDBInstanceNamespaceResponse) {
    response = &DescribeDBInstanceNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceNamespace
// This API is used to query the table information on a database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceNamespace(request *DescribeDBInstanceNamespaceRequest) (response *DescribeDBInstanceNamespaceResponse, err error) {
    return c.DescribeDBInstanceNamespaceWithContext(context.Background(), request)
}

// DescribeDBInstanceNamespace
// This API is used to query the table information on a database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceNamespaceWithContext(ctx context.Context, request *DescribeDBInstanceNamespaceRequest) (response *DescribeDBInstanceNamespaceResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceNodePropertyRequest() (request *DescribeDBInstanceNodePropertyRequest) {
    request = &DescribeDBInstanceNodePropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceNodeProperty")
    
    
    return
}

func NewDescribeDBInstanceNodePropertyResponse() (response *DescribeDBInstanceNodePropertyResponse) {
    response = &DescribeDBInstanceNodePropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceNodeProperty
// This API is used to query node attributes, such as the AZ, node name, address, role, status, delay between primary and secondary nodes, priority, voting right, and tags.
//
// error code that may be returned:
//  FAILEDOPERATION_KERNELRESPONSETIMEOUT = "FailedOperation.KernelResponseTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceNodeProperty(request *DescribeDBInstanceNodePropertyRequest) (response *DescribeDBInstanceNodePropertyResponse, err error) {
    return c.DescribeDBInstanceNodePropertyWithContext(context.Background(), request)
}

// DescribeDBInstanceNodeProperty
// This API is used to query node attributes, such as the AZ, node name, address, role, status, delay between primary and secondary nodes, priority, voting right, and tags.
//
// error code that may be returned:
//  FAILEDOPERATION_KERNELRESPONSETIMEOUT = "FailedOperation.KernelResponseTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceNodePropertyWithContext(ctx context.Context, request *DescribeDBInstanceNodePropertyRequest) (response *DescribeDBInstanceNodePropertyResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceNodePropertyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceNodeProperty")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceNodeProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceNodePropertyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstances
// This API is used to query the list of TencentDB for MongoDB instances. It supports filtering primary instances, disaster recovery instances, and read-only instances by project ID, instance ID, instance status, and other conditions.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  LIMITEXCEEDED_TOOMANYREQUESTS = "LimitExceeded.TooManyRequests"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// This API is used to query the list of TencentDB for MongoDB instances. It supports filtering primary instances, disaster recovery instances, and read-only instances by project ID, instance ID, instance status, and other conditions.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  LIMITEXCEEDED_TOOMANYREQUESTS = "LimitExceeded.TooManyRequests"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailedSlowLogsRequest() (request *DescribeDetailedSlowLogsRequest) {
    request = &DescribeDetailedSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDetailedSlowLogs")
    
    
    return
}

func NewDescribeDetailedSlowLogsResponse() (response *DescribeDetailedSlowLogsResponse) {
    response = &DescribeDetailedSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDetailedSlowLogs
// This API is used to query slow log details of the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITPARAOUTOFRANGE = "InvalidParameterValue.LimitParaOutOfRange"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OFFSETPARAOUTOFRANGE = "InvalidParameterValue.OffsetParaOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeDetailedSlowLogs(request *DescribeDetailedSlowLogsRequest) (response *DescribeDetailedSlowLogsResponse, err error) {
    return c.DescribeDetailedSlowLogsWithContext(context.Background(), request)
}

// DescribeDetailedSlowLogs
// This API is used to query slow log details of the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITPARAOUTOFRANGE = "InvalidParameterValue.LimitParaOutOfRange"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OFFSETPARAOUTOFRANGE = "InvalidParameterValue.OffsetParaOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeDetailedSlowLogsWithContext(ctx context.Context, request *DescribeDetailedSlowLogsRequest) (response *DescribeDetailedSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDetailedSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDetailedSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetailedSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetailedSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParams
// This API is used to query the list of parameters that can be modified for the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CURRENTINSTANCENOTSUPPORTMODIFYPARAMS = "InvalidParameter.CurrentInstanceNotSupportModifyParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// This API is used to query the list of parameters that can be modified for the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CURRENTINSTANCENOTSUPPORTMODIFYPARAMS = "InvalidParameter.CurrentInstanceNotSupportModifyParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
    request = &DescribeSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSecurityGroup")
    
    
    return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
    response = &DescribeSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroup
// This API is used to query security groups bound to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
    return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

// DescribeSecurityGroup
// This API is used to query security groups bound to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogPatternsRequest() (request *DescribeSlowLogPatternsRequest) {
    request = &DescribeSlowLogPatternsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogPatterns")
    
    
    return
}

func NewDescribeSlowLogPatternsResponse() (response *DescribeSlowLogPatternsResponse) {
    response = &DescribeSlowLogPatternsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogPatterns
// This API is used to get the slow log statistics of a database instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogPatterns(request *DescribeSlowLogPatternsRequest) (response *DescribeSlowLogPatternsResponse, err error) {
    return c.DescribeSlowLogPatternsWithContext(context.Background(), request)
}

// DescribeSlowLogPatterns
// This API is used to get the slow log statistics of a database instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogPatternsWithContext(ctx context.Context, request *DescribeSlowLogPatternsRequest) (response *DescribeSlowLogPatternsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogPatternsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSlowLogPatterns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogPatterns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogPatternsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogs")
    
    
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogs
// This API is used to get the slow log information of a TencentDB instance. Only slow logs for the last 7 days can be queried.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// This API is used to get the slow log information of a TencentDB instance. Only slow logs for the last 7 days can be queried.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecInfoRequest() (request *DescribeSpecInfoRequest) {
    request = &DescribeSpecInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSpecInfo")
    
    
    return
}

func NewDescribeSpecInfoResponse() (response *DescribeSpecInfoResponse) {
    response = &DescribeSpecInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpecInfo
// This API is used to query the sales specification of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfo(request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    return c.DescribeSpecInfoWithContext(context.Background(), request)
}

// DescribeSpecInfo
// This API is used to query the sales specification of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfoWithContext(ctx context.Context, request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpecInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSpecInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpecInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecInfoResponse()
    err = c.Send(request, response)
    return
}

func NewFlushInstanceRouterConfigRequest() (request *FlushInstanceRouterConfigRequest) {
    request = &FlushInstanceRouterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "FlushInstanceRouterConfig")
    
    
    return
}

func NewFlushInstanceRouterConfigResponse() (response *FlushInstanceRouterConfigResponse) {
    response = &FlushInstanceRouterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FlushInstanceRouterConfig
// This API is used to run the `FlushRouterConfig` command on all mongos instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FlushInstanceRouterConfig(request *FlushInstanceRouterConfigRequest) (response *FlushInstanceRouterConfigResponse, err error) {
    return c.FlushInstanceRouterConfigWithContext(context.Background(), request)
}

// FlushInstanceRouterConfig
// This API is used to run the `FlushRouterConfig` command on all mongos instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FlushInstanceRouterConfigWithContext(ctx context.Context, request *FlushInstanceRouterConfigRequest) (response *FlushInstanceRouterConfigResponse, err error) {
    if request == nil {
        request = NewFlushInstanceRouterConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "FlushInstanceRouterConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlushInstanceRouterConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlushInstanceRouterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDBInstancesRequest() (request *InquirePriceCreateDBInstancesRequest) {
    request = &InquirePriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceCreateDBInstances")
    
    
    return
}

func NewInquirePriceCreateDBInstancesResponse() (response *InquirePriceCreateDBInstancesResponse) {
    response = &InquirePriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateDBInstances
// This API is used to query price of instance creation. The `region` parameter must be passed in this API, otherwise verification will fail. This API allows queries only for purchasable instance specifications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) InquirePriceCreateDBInstances(request *InquirePriceCreateDBInstancesRequest) (response *InquirePriceCreateDBInstancesResponse, err error) {
    return c.InquirePriceCreateDBInstancesWithContext(context.Background(), request)
}

// InquirePriceCreateDBInstances
// This API is used to query price of instance creation. The `region` parameter must be passed in this API, otherwise verification will fail. This API allows queries only for purchasable instance specifications.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) InquirePriceCreateDBInstancesWithContext(ctx context.Context, request *InquirePriceCreateDBInstancesRequest) (response *InquirePriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceCreateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDBInstanceSpecRequest() (request *InquirePriceModifyDBInstanceSpecRequest) {
    request = &InquirePriceModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceModifyDBInstanceSpec")
    
    
    return
}

func NewInquirePriceModifyDBInstanceSpecResponse() (response *InquirePriceModifyDBInstanceSpecResponse) {
    response = &InquirePriceModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceModifyDBInstanceSpec
// This API is used to query the price of instance specification adjustment.
//
// error code that may be returned:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpec(request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    return c.InquirePriceModifyDBInstanceSpecWithContext(context.Background(), request)
}

// InquirePriceModifyDBInstanceSpec
// This API is used to query the price of instance specification adjustment.
//
// error code that may be returned:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpecWithContext(ctx context.Context, request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDBInstanceSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceModifyDBInstanceSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewDBInstancesRequest() (request *InquirePriceRenewDBInstancesRequest) {
    request = &InquirePriceRenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceRenewDBInstances")
    
    
    return
}

func NewInquirePriceRenewDBInstancesResponse() (response *InquirePriceRenewDBInstancesResponse) {
    response = &InquirePriceRenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenewDBInstances
// This API is used to query the renewal price of a monthly subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceRenewDBInstances(request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    return c.InquirePriceRenewDBInstancesWithContext(context.Background(), request)
}

// InquirePriceRenewDBInstances
// This API is used to query the renewal price of a monthly subscription instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceRenewDBInstancesWithContext(ctx context.Context, request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceRenewDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "IsolateDBInstance")
    
    
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBInstance
// This API is used to isolate a pay-as-you-go TencentDB for MongoDB instance. After isolation, the instance is retained in the recycle bin, and data cannot be written into it. After a certain period of isolation, the instance is deleted permanently. For the retention time in the recycle bin, see the pay-as-you-go service terms. The deleted pay-as-you-go instance cannot be recovered. Proceed with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PREPAIDINSTANCEUNABLETOISOLATE = "InvalidParameterValue.PrePaidInstanceUnableToIsolate"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// This API is used to isolate a pay-as-you-go TencentDB for MongoDB instance. After isolation, the instance is retained in the recycle bin, and data cannot be written into it. After a certain period of isolation, the instance is deleted permanently. For the retention time in the recycle bin, see the pay-as-you-go service terms. The deleted pay-as-you-go instance cannot be recovered. Proceed with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PREPAIDINSTANCEUNABLETOISOLATE = "InvalidParameterValue.PrePaidInstanceUnableToIsolate"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "IsolateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNetworkAddressRequest() (request *ModifyDBInstanceNetworkAddressRequest) {
    request = &ModifyDBInstanceNetworkAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceNetworkAddress")
    
    
    return
}

func NewModifyDBInstanceNetworkAddressResponse() (response *ModifyDBInstanceNetworkAddressResponse) {
    response = &ModifyDBInstanceNetworkAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceNetworkAddress
// This API is used to modify the network information on a TencentDB for MongoDB instance. It supports switching from a basic network to a VPC network or from one VPC network to another VPC network.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWMODIFYADDRAFTEROPENWANSERVICE = "FailedOperation.NotAllowModifyAddrAfterOpenWanService"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVIP = "InvalidParameter.InvalidVip"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
func (c *Client) ModifyDBInstanceNetworkAddress(request *ModifyDBInstanceNetworkAddressRequest) (response *ModifyDBInstanceNetworkAddressResponse, err error) {
    return c.ModifyDBInstanceNetworkAddressWithContext(context.Background(), request)
}

// ModifyDBInstanceNetworkAddress
// This API is used to modify the network information on a TencentDB for MongoDB instance. It supports switching from a basic network to a VPC network or from one VPC network to another VPC network.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWMODIFYADDRAFTEROPENWANSERVICE = "FailedOperation.NotAllowModifyAddrAfterOpenWanService"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVIP = "InvalidParameter.InvalidVip"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
func (c *Client) ModifyDBInstanceNetworkAddressWithContext(ctx context.Context, request *ModifyDBInstanceNetworkAddressRequest) (response *ModifyDBInstanceNetworkAddressResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNetworkAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceNetworkAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetworkAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNetworkAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupRequest() (request *ModifyDBInstanceSecurityGroupRequest) {
    request = &ModifyDBInstanceSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceSecurityGroup")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupResponse() (response *ModifyDBInstanceSecurityGroupResponse) {
    response = &ModifyDBInstanceSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroup
// This API is used to modify security groups bound to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroup(request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroup
// This API is used to modify security groups bound to an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroupWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceSpec")
    
    
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSpec
// This API is used to adjust the TencentDB for MongoDB instance configuration. The [DescribeSpecInfo](https://www.tencentcloud.comom/document/product/240/38567?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_ZONECLOSED = "InvalidParameter.ZoneClosed"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MODIFYMODEERROR = "InvalidParameterValue.ModifyModeError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OPLOGSIZEOUTOFRANGE = "InvalidParameterValue.OplogSizeOutOfRange"
//  INVALIDPARAMETERVALUE_SETDISKLESSTHANUSED = "InvalidParameterValue.SetDiskLessThanUsed"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

// ModifyDBInstanceSpec
// This API is used to adjust the TencentDB for MongoDB instance configuration. The [DescribeSpecInfo](https://www.tencentcloud.comom/document/product/240/38567?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_ZONECLOSED = "InvalidParameter.ZoneClosed"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MODIFYMODEERROR = "InvalidParameterValue.ModifyModeError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OPLOGSIZEOUTOFRANGE = "InvalidParameterValue.OplogSizeOutOfRange"
//  INVALIDPARAMETERVALUE_SETDISKLESSTHANUSED = "InvalidParameterValue.SetDiskLessThanUsed"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedDBInstanceRequest() (request *OfflineIsolatedDBInstanceRequest) {
    request = &OfflineIsolatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "OfflineIsolatedDBInstance")
    
    
    return
}

func NewOfflineIsolatedDBInstanceResponse() (response *OfflineIsolatedDBInstanceResponse) {
    response = &OfflineIsolatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineIsolatedDBInstance
// This API is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALSTATUSTOOFFLINE = "InvalidParameterValue.IllegalStatusToOffline"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) OfflineIsolatedDBInstance(request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    return c.OfflineIsolatedDBInstanceWithContext(context.Background(), request)
}

// OfflineIsolatedDBInstance
// This API is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALSTATUSTOOFFLINE = "InvalidParameterValue.IllegalStatusToOffline"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) OfflineIsolatedDBInstanceWithContext(ctx context.Context, request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "OfflineIsolatedDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
    request = &RenameInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RenameInstance")
    
    
    return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
    response = &RenameInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenameInstance
// This API is used to rename a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenameInstance(request *RenameInstanceRequest) (response *RenameInstanceResponse, err error) {
    return c.RenameInstanceWithContext(context.Background(), request)
}

// RenameInstance
// This API is used to rename a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) (response *RenameInstanceResponse, err error) {
    if request == nil {
        request = NewRenameInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RenameInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenameInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenameInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstancesRequest() (request *RenewDBInstancesRequest) {
    request = &RenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RenewDBInstances")
    
    
    return
}

func NewRenewDBInstancesResponse() (response *RenewDBInstancesResponse) {
    response = &RenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewDBInstances
// This API is used to renew a monthly subscription TencentDB instance. Only monthly subscription instances are supported, while pay-as-you-go instances do not need to be renewed.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenewDBInstances(request *RenewDBInstancesRequest) (response *RenewDBInstancesResponse, err error) {
    return c.RenewDBInstancesWithContext(context.Background(), request)
}

// RenewDBInstances
// This API is used to renew a monthly subscription TencentDB instance. Only monthly subscription instances are supported, while pay-as-you-go instances do not need to be renewed.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenewDBInstancesWithContext(ctx context.Context, request *RenewDBInstancesRequest) (response *RenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewRenewDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RenewDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetDBInstancePasswordRequest() (request *ResetDBInstancePasswordRequest) {
    request = &ResetDBInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ResetDBInstancePassword")
    
    
    return
}

func NewResetDBInstancePasswordResponse() (response *ResetDBInstancePasswordResponse) {
    response = &ResetDBInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDBInstancePassword
// This API is used to reset the instance access password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_USERNOTFOUND = "InvalidParameterValue.UserNotFound"
func (c *Client) ResetDBInstancePassword(request *ResetDBInstancePasswordRequest) (response *ResetDBInstancePasswordResponse, err error) {
    return c.ResetDBInstancePasswordWithContext(context.Background(), request)
}

// ResetDBInstancePassword
// This API is used to reset the instance access password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_USERNOTFOUND = "InvalidParameterValue.UserNotFound"
func (c *Client) ResetDBInstancePasswordWithContext(ctx context.Context, request *ResetDBInstancePasswordRequest) (response *ResetDBInstancePasswordResponse, err error) {
    if request == nil {
        request = NewResetDBInstancePasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ResetDBInstancePassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDBInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDBInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewSetDBInstanceDeletionProtectionRequest() (request *SetDBInstanceDeletionProtectionRequest) {
    request = &SetDBInstanceDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetDBInstanceDeletionProtection")
    
    
    return
}

func NewSetDBInstanceDeletionProtectionResponse() (response *SetDBInstanceDeletionProtectionResponse) {
    response = &SetDBInstanceDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetDBInstanceDeletionProtection
// This API is used to set instance termination protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) SetDBInstanceDeletionProtection(request *SetDBInstanceDeletionProtectionRequest) (response *SetDBInstanceDeletionProtectionResponse, err error) {
    return c.SetDBInstanceDeletionProtectionWithContext(context.Background(), request)
}

// SetDBInstanceDeletionProtection
// This API is used to set instance termination protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) SetDBInstanceDeletionProtectionWithContext(ctx context.Context, request *SetDBInstanceDeletionProtectionRequest) (response *SetDBInstanceDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewSetDBInstanceDeletionProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetDBInstanceDeletionProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDBInstanceDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDBInstanceDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDBInstancesRequest() (request *TerminateDBInstancesRequest) {
    request = &TerminateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "TerminateDBInstances")
    
    
    return
}

func NewTerminateDBInstancesResponse() (response *TerminateDBInstancesResponse) {
    response = &TerminateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateDBInstances
// This API is used to terminate monthly subscription billing instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) TerminateDBInstances(request *TerminateDBInstancesRequest) (response *TerminateDBInstancesResponse, err error) {
    return c.TerminateDBInstancesWithContext(context.Background(), request)
}

// TerminateDBInstances
// This API is used to terminate monthly subscription billing instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) TerminateDBInstancesWithContext(ctx context.Context, request *TerminateDBInstancesRequest) (response *TerminateDBInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "TerminateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDBInstancesResponse()
    err = c.Send(request, response)
    return
}
