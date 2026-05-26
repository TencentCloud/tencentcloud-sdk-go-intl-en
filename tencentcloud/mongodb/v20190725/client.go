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

func NewCloseAuditServiceRequest() (request *CloseAuditServiceRequest) {
    request = &CloseAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CloseAuditService")
    
    
    return
}

func NewCloseAuditServiceResponse() (response *CloseAuditServiceResponse) {
    response = &CloseAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAuditService
// This API is used to close the audit service.
//
// error code that may be returned:
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// This API is used to close the audit service.
//
// error code that may be returned:
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CloseAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountUserRequest() (request *CreateAccountUserRequest) {
    request = &CreateAccountUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateAccountUser")
    
    
    return
}

func NewCreateAccountUserResponse() (response *CreateAccountUserResponse) {
    response = &CreateAccountUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccountUser
// This API is used to customize an account to access the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_PASSWORDERROR = "InternalError.PasswordError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) CreateAccountUser(request *CreateAccountUserRequest) (response *CreateAccountUserResponse, err error) {
    return c.CreateAccountUserWithContext(context.Background(), request)
}

// CreateAccountUser
// This API is used to customize an account to access the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_PASSWORDERROR = "InternalError.PasswordError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) CreateAccountUserWithContext(ctx context.Context, request *CreateAccountUserRequest) (response *CreateAccountUserResponse, err error) {
    if request == nil {
        request = NewCreateAccountUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateAccountUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccountUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountUserResponse()
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
// This API is used to create a yearly/monthly subscription TencentDB for MongoDB instance. The [DescribeSpecInfo](https://www.tencentcloud.com/document/product/240/34701) API can be called to query and obtain the supported sales specifications.
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
// This API is used to create a yearly/monthly subscription TencentDB for MongoDB instance. The [DescribeSpecInfo](https://www.tencentcloud.com/document/product/240/34701) API can be called to query and obtain the supported sales specifications.
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

func NewCreateDBInstanceParamTplRequest() (request *CreateDBInstanceParamTplRequest) {
    request = &CreateDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstanceParamTpl")
    
    
    return
}

func NewCreateDBInstanceParamTplResponse() (response *CreateDBInstanceParamTplResponse) {
    response = &CreateDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstanceParamTpl
// This API is used to create a parameter template for TencentDB for MongoDB.
//
// **Description:** The CreateDBInstanceParamTpl API is in public beta. During this period, this API is only applicable to beta test participants.
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
func (c *Client) CreateDBInstanceParamTpl(request *CreateDBInstanceParamTplRequest) (response *CreateDBInstanceParamTplResponse, err error) {
    return c.CreateDBInstanceParamTplWithContext(context.Background(), request)
}

// CreateDBInstanceParamTpl
// This API is used to create a parameter template for TencentDB for MongoDB.
//
// **Description:** The CreateDBInstanceParamTpl API is in public beta. During this period, this API is only applicable to beta test participants.
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
func (c *Client) CreateDBInstanceParamTplWithContext(ctx context.Context, request *CreateDBInstanceParamTplRequest) (response *CreateDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogDownloadTaskRequest() (request *CreateLogDownloadTaskRequest) {
    request = &CreateLogDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateLogDownloadTask")
    
    
    return
}

func NewCreateLogDownloadTaskResponse() (response *CreateLogDownloadTaskResponse) {
    response = &CreateLogDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogDownloadTask
// This API is used to create a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateLogDownloadTask(request *CreateLogDownloadTaskRequest) (response *CreateLogDownloadTaskResponse, err error) {
    return c.CreateLogDownloadTaskWithContext(context.Background(), request)
}

// CreateLogDownloadTask
// This API is used to create a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateLogDownloadTaskWithContext(ctx context.Context, request *CreateLogDownloadTaskRequest) (response *CreateLogDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateLogDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateLogDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountUserRequest() (request *DeleteAccountUserRequest) {
    request = &DeleteAccountUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteAccountUser")
    
    
    return
}

func NewDeleteAccountUserResponse() (response *DeleteAccountUserResponse) {
    response = &DeleteAccountUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccountUser
// This API is used to delete a custom account of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteAccountUser(request *DeleteAccountUserRequest) (response *DeleteAccountUserResponse, err error) {
    return c.DeleteAccountUserWithContext(context.Background(), request)
}

// DeleteAccountUser
// This API is used to delete a custom account of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteAccountUserWithContext(ctx context.Context, request *DeleteAccountUserRequest) (response *DeleteAccountUserResponse, err error) {
    if request == nil {
        request = NewDeleteAccountUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteAccountUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccountUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBBackupsRequest() (request *DeleteDBBackupsRequest) {
    request = &DeleteDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteDBBackups")
    
    
    return
}

func NewDeleteDBBackupsResponse() (response *DeleteDBBackupsResponse) {
    response = &DeleteDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBBackups
// This API is used to delete full backups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteDBBackups(request *DeleteDBBackupsRequest) (response *DeleteDBBackupsResponse, err error) {
    return c.DeleteDBBackupsWithContext(context.Background(), request)
}

// DeleteDBBackups
// This API is used to delete full backups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteDBBackupsWithContext(ctx context.Context, request *DeleteDBBackupsRequest) (response *DeleteDBBackupsResponse, err error) {
    if request == nil {
        request = NewDeleteDBBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteDBBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogDownloadTaskRequest() (request *DeleteLogDownloadTaskRequest) {
    request = &DeleteLogDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteLogDownloadTask")
    
    
    return
}

func NewDeleteLogDownloadTaskResponse() (response *DeleteLogDownloadTaskResponse) {
    response = &DeleteLogDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogDownloadTask
// This API is used to delete a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteLogDownloadTask(request *DeleteLogDownloadTaskRequest) (response *DeleteLogDownloadTaskResponse, err error) {
    return c.DeleteLogDownloadTaskWithContext(context.Background(), request)
}

// DeleteLogDownloadTask
// This API is used to delete a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteLogDownloadTaskWithContext(ctx context.Context, request *DeleteLogDownloadTaskRequest) (response *DeleteLogDownloadTaskResponse, err error) {
    if request == nil {
        request = NewDeleteLogDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteLogDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountUsersRequest() (request *DescribeAccountUsersRequest) {
    request = &DescribeAccountUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAccountUsers")
    
    
    return
}

func NewDescribeAccountUsersResponse() (response *DescribeAccountUsersResponse) {
    response = &DescribeAccountUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountUsers
// This API is used to obtain all accounts of the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeAccountUsers(request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    return c.DescribeAccountUsersWithContext(context.Background(), request)
}

// DescribeAccountUsers
// This API is used to obtain all accounts of the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeAccountUsersWithContext(ctx context.Context, request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    if request == nil {
        request = NewDescribeAccountUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAccountUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountUsersResponse()
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

func NewDescribeAuditConfigRequest() (request *DescribeAuditConfigRequest) {
    request = &DescribeAuditConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAuditConfig")
    
    
    return
}

func NewDescribeAuditConfigResponse() (response *DescribeAuditConfigResponse) {
    response = &DescribeAuditConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditConfig
// This API is used to query the service configurations for a TencentDB audit policy, including the audit log retention period.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) DescribeAuditConfig(request *DescribeAuditConfigRequest) (response *DescribeAuditConfigResponse, err error) {
    return c.DescribeAuditConfigWithContext(context.Background(), request)
}

// DescribeAuditConfig
// This API is used to query the service configurations for a TencentDB audit policy, including the audit log retention period.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) DescribeAuditConfigWithContext(ctx context.Context, request *DescribeAuditConfigRequest) (response *DescribeAuditConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAuditConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAuditConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogFilesRequest() (request *DescribeAuditLogFilesRequest) {
    request = &DescribeAuditLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAuditLogFiles")
    
    
    return
}

func NewDescribeAuditLogFilesResponse() (response *DescribeAuditLogFilesResponse) {
    response = &DescribeAuditLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogFiles
// This API is used to query audit log files of a cloud database instance.
//
// error code that may be returned:
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// This API is used to query audit log files of a cloud database instance.
//
// error code that may be returned:
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditLogFilesWithContext(ctx context.Context, request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAuditLogFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogsRequest() (request *DescribeAuditLogsRequest) {
    request = &DescribeAuditLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAuditLogs")
    
    
    return
}

func NewDescribeAuditLogsResponse() (response *DescribeAuditLogsResponse) {
    response = &DescribeAuditLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogs
// This API is used to query database audit logs.
//
// error code that may be returned:
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    return c.DescribeAuditLogsWithContext(context.Background(), request)
}

// DescribeAuditLogs
// This API is used to query database audit logs.
//
// error code that may be returned:
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditLogsWithContext(ctx context.Context, request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAuditLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogsResponse()
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

func NewDescribeBackupRulesRequest() (request *DescribeBackupRulesRequest) {
    request = &DescribeBackupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeBackupRules")
    
    
    return
}

func NewDescribeBackupRulesResponse() (response *DescribeBackupRulesResponse) {
    response = &DescribeBackupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupRules
// This API is used to obtain the automatic backup configuration information of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeBackupRules(request *DescribeBackupRulesRequest) (response *DescribeBackupRulesResponse, err error) {
    return c.DescribeBackupRulesWithContext(context.Background(), request)
}

// DescribeBackupRules
// This API is used to obtain the automatic backup configuration information of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeBackupRulesWithContext(ctx context.Context, request *DescribeBackupRulesRequest) (response *DescribeBackupRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeBackupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupRulesResponse()
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

func NewDescribeCurrentOpRequest() (request *DescribeCurrentOpRequest) {
    request = &DescribeCurrentOpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeCurrentOp")
    
    
    return
}

func NewDescribeCurrentOpResponse() (response *DescribeCurrentOpResponse) {
    response = &DescribeCurrentOpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCurrentOp
// This API is used to query the operation currently being performed on a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYOUTOFRANGE = "InvalidParameterValue.QueryOutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTOPERATION = "InvalidParameterValue.RegionNotSupportOperation"
//  INVALIDPARAMETERVALUE_REPLICANOTFOUND = "InvalidParameterValue.ReplicaNotFound"
func (c *Client) DescribeCurrentOp(request *DescribeCurrentOpRequest) (response *DescribeCurrentOpResponse, err error) {
    return c.DescribeCurrentOpWithContext(context.Background(), request)
}

// DescribeCurrentOp
// This API is used to query the operation currently being performed on a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYOUTOFRANGE = "InvalidParameterValue.QueryOutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTOPERATION = "InvalidParameterValue.RegionNotSupportOperation"
//  INVALIDPARAMETERVALUE_REPLICANOTFOUND = "InvalidParameterValue.ReplicaNotFound"
func (c *Client) DescribeCurrentOpWithContext(ctx context.Context, request *DescribeCurrentOpRequest) (response *DescribeCurrentOpResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentOpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeCurrentOp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentOp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentOpResponse()
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

func NewDescribeDBInstanceParamTplRequest() (request *DescribeDBInstanceParamTplRequest) {
    request = &DescribeDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceParamTpl")
    
    
    return
}

func NewDescribeDBInstanceParamTplResponse() (response *DescribeDBInstanceParamTplResponse) {
    response = &DescribeDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceParamTpl
// This API is used to query ALL MongoDB database parameter templates under the current account.
//
// **Description:** The DescribeDBInstanceParamTpl API is in public beta. During this period, this API is only applicable to beta test participants.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTpl(request *DescribeDBInstanceParamTplRequest) (response *DescribeDBInstanceParamTplResponse, err error) {
    return c.DescribeDBInstanceParamTplWithContext(context.Background(), request)
}

// DescribeDBInstanceParamTpl
// This API is used to query ALL MongoDB database parameter templates under the current account.
//
// **Description:** The DescribeDBInstanceParamTpl API is in public beta. During this period, this API is only applicable to beta test participants.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplWithContext(ctx context.Context, request *DescribeDBInstanceParamTplRequest) (response *DescribeDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceParamTplDetailRequest() (request *DescribeDBInstanceParamTplDetailRequest) {
    request = &DescribeDBInstanceParamTplDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceParamTplDetail")
    
    
    return
}

func NewDescribeDBInstanceParamTplDetailResponse() (response *DescribeDBInstanceParamTplDetailResponse) {
    response = &DescribeDBInstanceParamTplDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceParamTplDetail
// This API is used to query parameter template details of a cloud database instance for MongoDB.
//
// **Description:** The DescribeDBInstanceParamTplDetail API is in public beta. During this period, this interface is only applicable to beta test participants.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplDetail(request *DescribeDBInstanceParamTplDetailRequest) (response *DescribeDBInstanceParamTplDetailResponse, err error) {
    return c.DescribeDBInstanceParamTplDetailWithContext(context.Background(), request)
}

// DescribeDBInstanceParamTplDetail
// This API is used to query parameter template details of a cloud database instance for MongoDB.
//
// **Description:** The DescribeDBInstanceParamTplDetail API is in public beta. During this period, this interface is only applicable to beta test participants.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplDetailWithContext(ctx context.Context, request *DescribeDBInstanceParamTplDetailRequest) (response *DescribeDBInstanceParamTplDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParamTplDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceParamTplDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceParamTplDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceParamTplDetailResponse()
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

func NewDescribeInstanceSSLRequest() (request *DescribeInstanceSSLRequest) {
    request = &DescribeInstanceSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceSSL")
    
    
    return
}

func NewDescribeInstanceSSLResponse() (response *DescribeInstanceSSLResponse) {
    response = &DescribeInstanceSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSSL
// This API is used to view the enabling status of Secure Sockets Layer (SSL) for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (response *DescribeInstanceSSLResponse, err error) {
    return c.DescribeInstanceSSLWithContext(context.Background(), request)
}

// DescribeInstanceSSL
// This API is used to view the enabling status of Secure Sockets Layer (SSL) for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) DescribeInstanceSSLWithContext(ctx context.Context, request *DescribeInstanceSSLRequest) (response *DescribeInstanceSSLResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeInstanceSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSSLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogDownloadTasksRequest() (request *DescribeLogDownloadTasksRequest) {
    request = &DescribeLogDownloadTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeLogDownloadTasks")
    
    
    return
}

func NewDescribeLogDownloadTasksResponse() (response *DescribeLogDownloadTasksResponse) {
    response = &DescribeLogDownloadTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogDownloadTasks
// This API is used to query a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeLogDownloadTasks(request *DescribeLogDownloadTasksRequest) (response *DescribeLogDownloadTasksResponse, err error) {
    return c.DescribeLogDownloadTasksWithContext(context.Background(), request)
}

// DescribeLogDownloadTasks
// This API is used to query a log download task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeLogDownloadTasksWithContext(ctx context.Context, request *DescribeLogDownloadTasksRequest) (response *DescribeLogDownloadTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLogDownloadTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeLogDownloadTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogDownloadTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogDownloadTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMongodbLogsRequest() (request *DescribeMongodbLogsRequest) {
    request = &DescribeMongodbLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongodbLogs")
    
    
    return
}

func NewDescribeMongodbLogsResponse() (response *DescribeMongodbLogsResponse) {
    response = &DescribeMongodbLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMongodbLogs
// This API is used to query running logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeMongodbLogs(request *DescribeMongodbLogsRequest) (response *DescribeMongodbLogsResponse, err error) {
    return c.DescribeMongodbLogsWithContext(context.Background(), request)
}

// DescribeMongodbLogs
// This API is used to query running logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeMongodbLogsWithContext(ctx context.Context, request *DescribeMongodbLogsRequest) (response *DescribeMongodbLogsResponse, err error) {
    if request == nil {
        request = NewDescribeMongodbLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeMongodbLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMongodbLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMongodbLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePasswordRotationRequest() (request *DescribePasswordRotationRequest) {
    request = &DescribePasswordRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribePasswordRotation")
    
    
    return
}

func NewDescribePasswordRotationResponse() (response *DescribePasswordRotationResponse) {
    response = &DescribePasswordRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePasswordRotation
// Retrieve the rotation status info
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribePasswordRotation(request *DescribePasswordRotationRequest) (response *DescribePasswordRotationResponse, err error) {
    return c.DescribePasswordRotationWithContext(context.Background(), request)
}

// DescribePasswordRotation
// Retrieve the rotation status info
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribePasswordRotationWithContext(ctx context.Context, request *DescribePasswordRotationRequest) (response *DescribePasswordRotationResponse, err error) {
    if request == nil {
        request = NewDescribePasswordRotationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribePasswordRotation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePasswordRotation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePasswordRotationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSRVConnectionDomainRequest() (request *DescribeSRVConnectionDomainRequest) {
    request = &DescribeSRVConnectionDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSRVConnectionDomain")
    
    
    return
}

func NewDescribeSRVConnectionDomainResponse() (response *DescribeSRVConnectionDomainResponse) {
    response = &DescribeSRVConnectionDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSRVConnectionDomain
// This API is used to query the current domain information of the MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSRVConnectionDomain(request *DescribeSRVConnectionDomainRequest) (response *DescribeSRVConnectionDomainResponse, err error) {
    return c.DescribeSRVConnectionDomainWithContext(context.Background(), request)
}

// DescribeSRVConnectionDomain
// This API is used to query the current domain information of the MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSRVConnectionDomainWithContext(ctx context.Context, request *DescribeSRVConnectionDomainRequest) (response *DescribeSRVConnectionDomainResponse, err error) {
    if request == nil {
        request = NewDescribeSRVConnectionDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSRVConnectionDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSRVConnectionDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSRVConnectionDomainResponse()
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

func NewDisableSRVConnectionUrlRequest() (request *DisableSRVConnectionUrlRequest) {
    request = &DisableSRVConnectionUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DisableSRVConnectionUrl")
    
    
    return
}

func NewDisableSRVConnectionUrlResponse() (response *DisableSRVConnectionUrlResponse) {
    response = &DisableSRVConnectionUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableSRVConnectionUrl
// This API is used to close the SRV Access Address of the MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DisableSRVConnectionUrl(request *DisableSRVConnectionUrlRequest) (response *DisableSRVConnectionUrlResponse, err error) {
    return c.DisableSRVConnectionUrlWithContext(context.Background(), request)
}

// DisableSRVConnectionUrl
// This API is used to close the SRV Access Address of the MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DisableSRVConnectionUrlWithContext(ctx context.Context, request *DisableSRVConnectionUrlRequest) (response *DisableSRVConnectionUrlResponse, err error) {
    if request == nil {
        request = NewDisableSRVConnectionUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DisableSRVConnectionUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableSRVConnectionUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableSRVConnectionUrlResponse()
    err = c.Send(request, response)
    return
}

func NewEnablePasswordRotationRequest() (request *EnablePasswordRotationRequest) {
    request = &EnablePasswordRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "EnablePasswordRotation")
    
    
    return
}

func NewEnablePasswordRotationResponse() (response *EnablePasswordRotationResponse) {
    response = &EnablePasswordRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnablePasswordRotation
// Enable password rotation
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) EnablePasswordRotation(request *EnablePasswordRotationRequest) (response *EnablePasswordRotationResponse, err error) {
    return c.EnablePasswordRotationWithContext(context.Background(), request)
}

// EnablePasswordRotation
// Enable password rotation
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) EnablePasswordRotationWithContext(ctx context.Context, request *EnablePasswordRotationRequest) (response *EnablePasswordRotationResponse, err error) {
    if request == nil {
        request = NewEnablePasswordRotationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "EnablePasswordRotation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnablePasswordRotation require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnablePasswordRotationResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSRVConnectionUrlRequest() (request *EnableSRVConnectionUrlRequest) {
    request = &EnableSRVConnectionUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "EnableSRVConnectionUrl")
    
    
    return
}

func NewEnableSRVConnectionUrlResponse() (response *EnableSRVConnectionUrlResponse) {
    response = &EnableSRVConnectionUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSRVConnectionUrl
// This API is used to enable the SRV Access Address for MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableSRVConnectionUrl(request *EnableSRVConnectionUrlRequest) (response *EnableSRVConnectionUrlResponse, err error) {
    return c.EnableSRVConnectionUrlWithContext(context.Background(), request)
}

// EnableSRVConnectionUrl
// This API is used to enable the SRV Access Address for MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableSRVConnectionUrlWithContext(ctx context.Context, request *EnableSRVConnectionUrlRequest) (response *EnableSRVConnectionUrlResponse, err error) {
    if request == nil {
        request = NewEnableSRVConnectionUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "EnableSRVConnectionUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSRVConnectionUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableSRVConnectionUrlResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTransparentDataEncryptionRequest() (request *EnableTransparentDataEncryptionRequest) {
    request = &EnableTransparentDataEncryptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "EnableTransparentDataEncryption")
    
    
    return
}

func NewEnableTransparentDataEncryptionResponse() (response *EnableTransparentDataEncryptionResponse) {
    response = &EnableTransparentDataEncryptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableTransparentDataEncryption
// This API is used to enable the transparent data encryption (TDE) capability for TencentDB for MongoDB.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableTransparentDataEncryption(request *EnableTransparentDataEncryptionRequest) (response *EnableTransparentDataEncryptionResponse, err error) {
    return c.EnableTransparentDataEncryptionWithContext(context.Background(), request)
}

// EnableTransparentDataEncryption
// This API is used to enable the transparent data encryption (TDE) capability for TencentDB for MongoDB.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableTransparentDataEncryptionWithContext(ctx context.Context, request *EnableTransparentDataEncryptionRequest) (response *EnableTransparentDataEncryptionResponse, err error) {
    if request == nil {
        request = NewEnableTransparentDataEncryptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "EnableTransparentDataEncryption")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTransparentDataEncryption require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTransparentDataEncryptionResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWanServiceRequest() (request *EnableWanServiceRequest) {
    request = &EnableWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "EnableWanService")
    
    
    return
}

func NewEnableWanServiceResponse() (response *EnableWanServiceResponse) {
    response = &EnableWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableWanService
// This API is used to enable the public network access address of the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYLEVELNOTALLOWOPENWANSERVICE = "InvalidParameterValue.SecurityLevelNotAllowOpenWanService"
func (c *Client) EnableWanService(request *EnableWanServiceRequest) (response *EnableWanServiceResponse, err error) {
    return c.EnableWanServiceWithContext(context.Background(), request)
}

// EnableWanService
// This API is used to enable the public network access address of the current instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYLEVELNOTALLOWOPENWANSERVICE = "InvalidParameterValue.SecurityLevelNotAllowOpenWanService"
func (c *Client) EnableWanServiceWithContext(ctx context.Context, request *EnableWanServiceRequest) (response *EnableWanServiceResponse, err error) {
    if request == nil {
        request = NewEnableWanServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "EnableWanService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableWanService require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableWanServiceResponse()
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

func NewInstanceEnableSSLRequest() (request *InstanceEnableSSLRequest) {
    request = &InstanceEnableSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InstanceEnableSSL")
    
    
    return
}

func NewInstanceEnableSSLResponse() (response *InstanceEnableSSLResponse) {
    response = &InstanceEnableSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstanceEnableSSL
// This API is used to set the SSL status for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) InstanceEnableSSL(request *InstanceEnableSSLRequest) (response *InstanceEnableSSLResponse, err error) {
    return c.InstanceEnableSSLWithContext(context.Background(), request)
}

// InstanceEnableSSL
// This API is used to set the SSL status for an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) InstanceEnableSSLWithContext(ctx context.Context, request *InstanceEnableSSLRequest) (response *InstanceEnableSSLResponse, err error) {
    if request == nil {
        request = NewInstanceEnableSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InstanceEnableSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstanceEnableSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstanceEnableSSLResponse()
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

func NewKillOpsRequest() (request *KillOpsRequest) {
    request = &KillOpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "KillOps")
    
    
    return
}

func NewKillOpsResponse() (response *KillOpsResponse) {
    response = &KillOpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillOps
// This API is used to terminate a specific operation performed on a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOps(request *KillOpsRequest) (response *KillOpsResponse, err error) {
    return c.KillOpsWithContext(context.Background(), request)
}

// KillOps
// This API is used to terminate a specific operation performed on a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOpsWithContext(ctx context.Context, request *KillOpsRequest) (response *KillOpsResponse, err error) {
    if request == nil {
        request = NewKillOpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "KillOps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillOps require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillOpsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupExpireTimeRequest() (request *ModifyBackupExpireTimeRequest) {
    request = &ModifyBackupExpireTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyBackupExpireTime")
    
    
    return
}

func NewModifyBackupExpireTimeResponse() (response *ModifyBackupExpireTimeResponse) {
    response = &ModifyBackupExpireTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupExpireTime
// Modify backup expiration time
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) ModifyBackupExpireTime(request *ModifyBackupExpireTimeRequest) (response *ModifyBackupExpireTimeResponse, err error) {
    return c.ModifyBackupExpireTimeWithContext(context.Background(), request)
}

// ModifyBackupExpireTime
// Modify backup expiration time
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) ModifyBackupExpireTimeWithContext(ctx context.Context, request *ModifyBackupExpireTimeRequest) (response *ModifyBackupExpireTimeResponse, err error) {
    if request == nil {
        request = NewModifyBackupExpireTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyBackupExpireTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupExpireTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupExpireTimeResponse()
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
// This API is used to adjust the TencentDB for MongoDB instance configuration. The [DescribeSpecInfo](https://www.tencentcloud.com/document/product/240/38567?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
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
// This API is used to adjust the TencentDB for MongoDB instance configuration. The [DescribeSpecInfo](https://www.tencentcloud.com/document/product/240/38567?from_cn_redirect=1) API can be called to query and obtain the supported sales specifications.
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

func NewModifyInstanceAzRequest() (request *ModifyInstanceAzRequest) {
    request = &ModifyInstanceAzRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyInstanceAz")
    
    
    return
}

func NewModifyInstanceAzResponse() (response *ModifyInstanceAzResponse) {
    response = &ModifyInstanceAzResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAz
// This API is used to adjust the availability zone distribution of MongoDB cloud database nodes. You can specify the primary AZ and total availability zone distribution info to complete node distribution adjustment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) ModifyInstanceAz(request *ModifyInstanceAzRequest) (response *ModifyInstanceAzResponse, err error) {
    return c.ModifyInstanceAzWithContext(context.Background(), request)
}

// ModifyInstanceAz
// This API is used to adjust the availability zone distribution of MongoDB cloud database nodes. You can specify the primary AZ and total availability zone distribution info to complete node distribution adjustment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) ModifyInstanceAzWithContext(ctx context.Context, request *ModifyInstanceAzRequest) (response *ModifyInstanceAzResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAzRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyInstanceAz")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAz require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAzResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyInstanceParams")
    
    
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceParams
// This API is used to modify the parameter configuration of a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATIONNOTALLOWEDININSTANCELOCKING = "FailedOperation.OperationNotAllowedInInstanceLocking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MODIFYMONGODBPARAMS = "InvalidParameter.ModifyMongodbParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MODIFYMONGODBPARAMS = "InvalidParameterValue.ModifyMongodbParams"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    return c.ModifyInstanceParamsWithContext(context.Background(), request)
}

// ModifyInstanceParams
// This API is used to modify the parameter configuration of a TencentDB for MongoDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATIONNOTALLOWEDININSTANCELOCKING = "FailedOperation.OperationNotAllowedInInstanceLocking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MODIFYMONGODBPARAMS = "InvalidParameter.ModifyMongodbParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MODIFYMONGODBPARAMS = "InvalidParameterValue.ModifyMongodbParams"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyInstanceParamsWithContext(ctx context.Context, request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySRVConnectionUrlRequest() (request *ModifySRVConnectionUrlRequest) {
    request = &ModifySRVConnectionUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifySRVConnectionUrl")
    
    
    return
}

func NewModifySRVConnectionUrlResponse() (response *ModifySRVConnectionUrlResponse) {
    response = &ModifySRVConnectionUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySRVConnectionUrl
// This API is used to modify the TTL duration of the SRV Access Address for a MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) ModifySRVConnectionUrl(request *ModifySRVConnectionUrlRequest) (response *ModifySRVConnectionUrlResponse, err error) {
    return c.ModifySRVConnectionUrlWithContext(context.Background(), request)
}

// ModifySRVConnectionUrl
// This API is used to modify the TTL duration of the SRV Access Address for a MongoDB database.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) ModifySRVConnectionUrlWithContext(ctx context.Context, request *ModifySRVConnectionUrlRequest) (response *ModifySRVConnectionUrlResponse, err error) {
    if request == nil {
        request = NewModifySRVConnectionUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifySRVConnectionUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySRVConnectionUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySRVConnectionUrlResponse()
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

func NewRestoreDBInstanceRequest() (request *RestoreDBInstanceRequest) {
    request = &RestoreDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RestoreDBInstance")
    
    
    return
}

func NewRestoreDBInstanceResponse() (response *RestoreDBInstanceResponse) {
    response = &RestoreDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreDBInstance
// This API is used to rollback the database instance to a specified time point.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) RestoreDBInstance(request *RestoreDBInstanceRequest) (response *RestoreDBInstanceResponse, err error) {
    return c.RestoreDBInstanceWithContext(context.Background(), request)
}

// RestoreDBInstance
// This API is used to rollback the database instance to a specified time point.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) RestoreDBInstanceWithContext(ctx context.Context, request *RestoreDBInstanceRequest) (response *RestoreDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RestoreDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetAccountUserPrivilegeRequest() (request *SetAccountUserPrivilegeRequest) {
    request = &SetAccountUserPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetAccountUserPrivilege")
    
    
    return
}

func NewSetAccountUserPrivilegeResponse() (response *SetAccountUserPrivilegeResponse) {
    response = &SetAccountUserPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetAccountUserPrivilege
// This API is used to set the account permissions of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilege(request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    return c.SetAccountUserPrivilegeWithContext(context.Background(), request)
}

// SetAccountUserPrivilege
// This API is used to set the account permissions of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilegeWithContext(ctx context.Context, request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    if request == nil {
        request = NewSetAccountUserPrivilegeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetAccountUserPrivilege")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAccountUserPrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAccountUserPrivilegeResponse()
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

func NewSetInstanceMaintenanceRequest() (request *SetInstanceMaintenanceRequest) {
    request = &SetInstanceMaintenanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetInstanceMaintenance")
    
    
    return
}

func NewSetInstanceMaintenanceResponse() (response *SetInstanceMaintenanceResponse) {
    response = &SetInstanceMaintenanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetInstanceMaintenance
// This API is used to set the instance maintenance window.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetInstanceMaintenance(request *SetInstanceMaintenanceRequest) (response *SetInstanceMaintenanceResponse, err error) {
    return c.SetInstanceMaintenanceWithContext(context.Background(), request)
}

// SetInstanceMaintenance
// This API is used to set the instance maintenance window.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetInstanceMaintenanceWithContext(ctx context.Context, request *SetInstanceMaintenanceRequest) (response *SetInstanceMaintenanceResponse, err error) {
    if request == nil {
        request = NewSetInstanceMaintenanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetInstanceMaintenance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetInstanceMaintenance require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetInstanceMaintenanceResponse()
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

func NewUpgradeDBInstanceKernelVersionRequest() (request *UpgradeDBInstanceKernelVersionRequest) {
    request = &UpgradeDBInstanceKernelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "UpgradeDBInstanceKernelVersion")
    
    
    return
}

func NewUpgradeDBInstanceKernelVersionResponse() (response *UpgradeDBInstanceKernelVersionResponse) {
    response = &UpgradeDBInstanceKernelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDBInstanceKernelVersion
// This API is used to upgrade the kernel version of the database instance.
//
// error code that may be returned:
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    return c.UpgradeDBInstanceKernelVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceKernelVersion
// This API is used to upgrade the kernel version of the database instance.
//
// error code that may be returned:
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) UpgradeDBInstanceKernelVersionWithContext(ctx context.Context, request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceKernelVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "UpgradeDBInstanceKernelVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceKernelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceKernelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDbInstanceVersionRequest() (request *UpgradeDbInstanceVersionRequest) {
    request = &UpgradeDbInstanceVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "UpgradeDbInstanceVersion")
    
    
    return
}

func NewUpgradeDbInstanceVersionResponse() (response *UpgradeDbInstanceVersionResponse) {
    response = &UpgradeDbInstanceVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDbInstanceVersion
// This API is used to upgrade database version.
//
// **Description**: Upgrade to version 3.6 and above is supported. Only upgrading from a lower version to a higher version step by step is allowed. Cross-version upgrade or version downgrade is not supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATEERROR = "InternalError.DBOperateError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) UpgradeDbInstanceVersion(request *UpgradeDbInstanceVersionRequest) (response *UpgradeDbInstanceVersionResponse, err error) {
    return c.UpgradeDbInstanceVersionWithContext(context.Background(), request)
}

// UpgradeDbInstanceVersion
// This API is used to upgrade database version.
//
// **Description**: Upgrade to version 3.6 and above is supported. Only upgrading from a lower version to a higher version step by step is allowed. Cross-version upgrade or version downgrade is not supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATEERROR = "InternalError.DBOperateError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) UpgradeDbInstanceVersionWithContext(ctx context.Context, request *UpgradeDbInstanceVersionRequest) (response *UpgradeDbInstanceVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDbInstanceVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "UpgradeDbInstanceVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDbInstanceVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDbInstanceVersionResponse()
    err = c.Send(request, response)
    return
}
