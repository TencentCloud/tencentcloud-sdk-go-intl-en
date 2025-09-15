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

package v20190103

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-01-03"

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


func NewAddMetricScaleStrategyRequest() (request *AddMetricScaleStrategyRequest) {
    request = &AddMetricScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AddMetricScaleStrategy")
    
    
    return
}

func NewAddMetricScaleStrategyResponse() (response *AddMetricScaleStrategyResponse) {
    response = &AddMetricScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddMetricScaleStrategy
// This API is used to add scaling rules by load and time.
//
// error code that may be returned:
//  FAILEDOPERATION_MORESTRATEGYNOTALLOWED = "FailedOperation.MoreStrategyNotAllowed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDCOMPAREMETHOD = "InvalidParameter.InvalidCompareMethod"
//  INVALIDPARAMETER_INVALIDCONDITIONNUM = "InvalidParameter.InvalidConditionNum"
//  INVALIDPARAMETER_INVALIDPARAMTERINVALIDSOFTINFO = "InvalidParameter.InvalidParamterInvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDPROCESSMETHOD = "InvalidParameter.InvalidProcessMethod"
//  INVALIDPARAMETER_INVALIDSCALEACTION = "InvalidParameter.InvalidScaleAction"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddMetricScaleStrategy(request *AddMetricScaleStrategyRequest) (response *AddMetricScaleStrategyResponse, err error) {
    return c.AddMetricScaleStrategyWithContext(context.Background(), request)
}

// AddMetricScaleStrategy
// This API is used to add scaling rules by load and time.
//
// error code that may be returned:
//  FAILEDOPERATION_MORESTRATEGYNOTALLOWED = "FailedOperation.MoreStrategyNotAllowed"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDCOMPAREMETHOD = "InvalidParameter.InvalidCompareMethod"
//  INVALIDPARAMETER_INVALIDCONDITIONNUM = "InvalidParameter.InvalidConditionNum"
//  INVALIDPARAMETER_INVALIDPARAMTERINVALIDSOFTINFO = "InvalidParameter.InvalidParamterInvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDPROCESSMETHOD = "InvalidParameter.InvalidProcessMethod"
//  INVALIDPARAMETER_INVALIDSCALEACTION = "InvalidParameter.InvalidScaleAction"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddMetricScaleStrategyWithContext(ctx context.Context, request *AddMetricScaleStrategyRequest) (response *AddMetricScaleStrategyResponse, err error) {
    if request == nil {
        request = NewAddMetricScaleStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "AddMetricScaleStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMetricScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMetricScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewAddNodeResourceConfigRequest() (request *AddNodeResourceConfigRequest) {
    request = &AddNodeResourceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AddNodeResourceConfig")
    
    
    return
}

func NewAddNodeResourceConfigResponse() (response *AddNodeResourceConfigResponse) {
    response = &AddNodeResourceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNodeResourceConfig
// This API is used to add node specifications of the current cluster.
//
// error code that may be returned:
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  RESOURCEUNAVAILABLE_REPEATSPEC = "ResourceUnavailable.RepeatSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddNodeResourceConfig(request *AddNodeResourceConfigRequest) (response *AddNodeResourceConfigResponse, err error) {
    return c.AddNodeResourceConfigWithContext(context.Background(), request)
}

// AddNodeResourceConfig
// This API is used to add node specifications of the current cluster.
//
// error code that may be returned:
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  RESOURCEUNAVAILABLE_REPEATSPEC = "ResourceUnavailable.RepeatSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) AddNodeResourceConfigWithContext(ctx context.Context, request *AddNodeResourceConfigRequest) (response *AddNodeResourceConfigResponse, err error) {
    if request == nil {
        request = NewAddNodeResourceConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "AddNodeResourceConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNodeResourceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNodeResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewAddUsersForUserManagerRequest() (request *AddUsersForUserManagerRequest) {
    request = &AddUsersForUserManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AddUsersForUserManager")
    
    
    return
}

func NewAddUsersForUserManagerResponse() (response *AddUsersForUserManagerResponse) {
    response = &AddUsersForUserManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUsersForUserManager
// This API is available for clusters with OpenLDAP components configured.
//
// This API is used to add user lists (user management).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBEXCEPTION = "FailedOperation.DBException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  LIMITEXCEEDED_REQUESTBACKLOGEXCEEDSLIMIT = "LimitExceeded.RequestBacklogExceedsLimit"
//  LIMITEXCEEDED_USERCOUNTEXCEEDSLIMIT = "LimitExceeded.UserCountExceedsLimit"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) AddUsersForUserManager(request *AddUsersForUserManagerRequest) (response *AddUsersForUserManagerResponse, err error) {
    return c.AddUsersForUserManagerWithContext(context.Background(), request)
}

// AddUsersForUserManager
// This API is available for clusters with OpenLDAP components configured.
//
// This API is used to add user lists (user management).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBEXCEPTION = "FailedOperation.DBException"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  LIMITEXCEEDED_REQUESTBACKLOGEXCEEDSLIMIT = "LimitExceeded.RequestBacklogExceedsLimit"
//  LIMITEXCEEDED_USERCOUNTEXCEEDSLIMIT = "LimitExceeded.UserCountExceedsLimit"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) AddUsersForUserManagerWithContext(ctx context.Context, request *AddUsersForUserManagerRequest) (response *AddUsersForUserManagerResponse, err error) {
    if request == nil {
        request = NewAddUsersForUserManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "AddUsersForUserManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUsersForUserManagerResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachDisks
// This API is used to mount cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    return c.AttachDisksWithContext(context.Background(), request)
}

// AttachDisks
// This API is used to mount cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
func (c *Client) AttachDisksWithContext(ctx context.Context, request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    if request == nil {
        request = NewAttachDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "AttachDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewConvertPreToPostClusterRequest() (request *ConvertPreToPostClusterRequest) {
    request = &ConvertPreToPostClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ConvertPreToPostCluster")
    
    
    return
}

func NewConvertPreToPostClusterResponse() (response *ConvertPreToPostClusterResponse) {
    response = &ConvertPreToPostClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConvertPreToPostCluster
// This API is used to convert a monthly subscription cluster to a pay-as-you-go cluster (excluding cdb).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ConvertPreToPostCluster(request *ConvertPreToPostClusterRequest) (response *ConvertPreToPostClusterResponse, err error) {
    return c.ConvertPreToPostClusterWithContext(context.Background(), request)
}

// ConvertPreToPostCluster
// This API is used to convert a monthly subscription cluster to a pay-as-you-go cluster (excluding cdb).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ConvertPreToPostClusterWithContext(ctx context.Context, request *ConvertPreToPostClusterRequest) (response *ConvertPreToPostClusterResponse, err error) {
    if request == nil {
        request = NewConvertPreToPostClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ConvertPreToPostCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConvertPreToPostCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewConvertPreToPostClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// This API is used to create an EMR cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// This API is used to create an EMR cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "CreateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupsSTDRequest() (request *CreateGroupsSTDRequest) {
    request = &CreateGroupsSTDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateGroupsSTD")
    
    
    return
}

func NewCreateGroupsSTDResponse() (response *CreateGroupsSTDResponse) {
    response = &CreateGroupsSTDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroupsSTD
// This API is used to create user groups in batches under User Management.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateGroupsSTD(request *CreateGroupsSTDRequest) (response *CreateGroupsSTDResponse, err error) {
    return c.CreateGroupsSTDWithContext(context.Background(), request)
}

// CreateGroupsSTD
// This API is used to create user groups in batches under User Management.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDALLNODERESOURCESPEC = "InvalidParameter.InvalidAllNodeResourceSpec"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOREDISKTYPE = "InvalidParameter.InvalidCoreDiskType"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPRODUCTVERSION = "InvalidParameter.InvalidProductVersion"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSCRIPTBOOTSTRAPACTIONCONFIG = "InvalidParameter.InvalidScriptBootstrapActionConfig"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_KERBEROSSUPPORT = "InvalidParameter.KerberosSupport"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) CreateGroupsSTDWithContext(ctx context.Context, request *CreateGroupsSTDRequest) (response *CreateGroupsSTDResponse, err error) {
    if request == nil {
        request = NewCreateGroupsSTDRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "CreateGroupsSTD")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupsSTD require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupsSTDResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// This API is used to create an EMR cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETADATAJDBCURL = "InvalidParameter.InvalidMetaDataJdbcUrl"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSECURITYSUPPORT = "InvalidParameter.InvalidSecuritySupport"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SECURITYGROUPNUMLIMITEXCEEDED = "LimitExceeded.SecurityGroupNumLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// This API is used to create an EMR cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOSBUCKET = "InvalidParameter.InvalidCosBucket"
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOADBALANCER = "InvalidParameter.InvalidLoadBalancer"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETADATAJDBCURL = "InvalidParameter.InvalidMetaDataJdbcUrl"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSECURITYSUPPORT = "InvalidParameter.InvalidSecuritySupport"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SECURITYGROUPNUMLIMITEXCEEDED = "LimitExceeded.SecurityGroupNumLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "CreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSLInstanceRequest() (request *CreateSLInstanceRequest) {
    request = &CreateSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "CreateSLInstance")
    
    
    return
}

func NewCreateSLInstanceResponse() (response *CreateSLInstanceResponse) {
    response = &CreateSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSLInstance
// This API is used to create a Serverless HBase instance.- If the API call is successful, a Serverless HBase instance will be created. If the instance creation request is successful, the InstanceId of the created instance and the RequestID of the request will be returned.- This is an asynchronous API. The operation is not completed immediately when the API call returns. The instance operation result can be viewed by calling DescribeInstancesList to view the StatusDesc status of the current instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDUINNUM = "InvalidParameter.InvalidUinNum"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) CreateSLInstance(request *CreateSLInstanceRequest) (response *CreateSLInstanceResponse, err error) {
    return c.CreateSLInstanceWithContext(context.Background(), request)
}

// CreateSLInstance
// This API is used to create a Serverless HBase instance.- If the API call is successful, a Serverless HBase instance will be created. If the instance creation request is successful, the InstanceId of the created instance and the RequestID of the request will be returned.- This is an asynchronous API. The operation is not completed immediately when the API call returns. The instance operation result can be viewed by calling DescribeInstancesList to view the StatusDesc status of the current instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDUINNUM = "InvalidParameter.InvalidUinNum"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) CreateSLInstanceWithContext(ctx context.Context, request *CreateSLInstanceRequest) (response *CreateSLInstanceResponse, err error) {
    if request == nil {
        request = NewCreateSLInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "CreateSLInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScaleStrategyRequest() (request *DeleteAutoScaleStrategyRequest) {
    request = &DeleteAutoScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeleteAutoScaleStrategy")
    
    
    return
}

func NewDeleteAutoScaleStrategyResponse() (response *DeleteAutoScaleStrategyResponse) {
    response = &DeleteAutoScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoScaleStrategy
// This API is used to delete automatic scaling rules. Nodes scaled based on these rules are destroyed in the background.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteAutoScaleStrategy(request *DeleteAutoScaleStrategyRequest) (response *DeleteAutoScaleStrategyResponse, err error) {
    return c.DeleteAutoScaleStrategyWithContext(context.Background(), request)
}

// DeleteAutoScaleStrategy
// This API is used to delete automatic scaling rules. Nodes scaled based on these rules are destroyed in the background.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteAutoScaleStrategyWithContext(ctx context.Context, request *DeleteAutoScaleStrategyRequest) (response *DeleteAutoScaleStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScaleStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DeleteAutoScaleStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupsSTDRequest() (request *DeleteGroupsSTDRequest) {
    request = &DeleteGroupsSTDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeleteGroupsSTD")
    
    
    return
}

func NewDeleteGroupsSTDResponse() (response *DeleteGroupsSTDResponse) {
    response = &DeleteGroupsSTDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroupsSTD
// This API is used to delete user groups in batches.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteGroupsSTD(request *DeleteGroupsSTDRequest) (response *DeleteGroupsSTDResponse, err error) {
    return c.DeleteGroupsSTDWithContext(context.Background(), request)
}

// DeleteGroupsSTD
// This API is used to delete user groups in batches.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteGroupsSTDWithContext(ctx context.Context, request *DeleteGroupsSTDRequest) (response *DeleteGroupsSTDResponse, err error) {
    if request == nil {
        request = NewDeleteGroupsSTDRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DeleteGroupsSTD")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroupsSTD require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupsSTDResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNodeResourceConfigRequest() (request *DeleteNodeResourceConfigRequest) {
    request = &DeleteNodeResourceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeleteNodeResourceConfig")
    
    
    return
}

func NewDeleteNodeResourceConfigResponse() (response *DeleteNodeResourceConfigResponse) {
    response = &DeleteNodeResourceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNodeResourceConfig
// This API is used to delete the node specifications of the current cluster.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteNodeResourceConfig(request *DeleteNodeResourceConfigRequest) (response *DeleteNodeResourceConfigResponse, err error) {
    return c.DeleteNodeResourceConfigWithContext(context.Background(), request)
}

// DeleteNodeResourceConfig
// This API is used to delete the node specifications of the current cluster.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeleteNodeResourceConfigWithContext(ctx context.Context, request *DeleteNodeResourceConfigRequest) (response *DeleteNodeResourceConfigResponse, err error) {
    if request == nil {
        request = NewDeleteNodeResourceConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DeleteNodeResourceConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNodeResourceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNodeResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeployYarnConfRequest() (request *DeployYarnConfRequest) {
    request = &DeployYarnConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DeployYarnConf")
    
    
    return
}

func NewDeployYarnConfResponse() (response *DeployYarnConfResponse) {
    response = &DeployYarnConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployYarnConf
// This API is used to bring the configuration into effect in YARN resource scheduling after deployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeployYarnConf(request *DeployYarnConfRequest) (response *DeployYarnConfResponse, err error) {
    return c.DeployYarnConfWithContext(context.Background(), request)
}

// DeployYarnConf
// This API is used to bring the configuration into effect in YARN resource scheduling after deployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DeployYarnConfWithContext(ctx context.Context, request *DeployYarnConfRequest) (response *DeployYarnConfResponse, err error) {
    if request == nil {
        request = NewDeployYarnConfRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DeployYarnConf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployYarnConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployYarnConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleGroupGlobalConfRequest() (request *DescribeAutoScaleGroupGlobalConfRequest) {
    request = &DescribeAutoScaleGroupGlobalConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleGroupGlobalConf")
    
    
    return
}

func NewDescribeAutoScaleGroupGlobalConfResponse() (response *DescribeAutoScaleGroupGlobalConfResponse) {
    response = &DescribeAutoScaleGroupGlobalConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleGroupGlobalConf
// This API is used to access the global configuration of automatic scaling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeAutoScaleGroupGlobalConf(request *DescribeAutoScaleGroupGlobalConfRequest) (response *DescribeAutoScaleGroupGlobalConfResponse, err error) {
    return c.DescribeAutoScaleGroupGlobalConfWithContext(context.Background(), request)
}

// DescribeAutoScaleGroupGlobalConf
// This API is used to access the global configuration of automatic scaling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeAutoScaleGroupGlobalConfWithContext(ctx context.Context, request *DescribeAutoScaleGroupGlobalConfRequest) (response *DescribeAutoScaleGroupGlobalConfResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleGroupGlobalConfRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeAutoScaleGroupGlobalConf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleGroupGlobalConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleGroupGlobalConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleRecordsRequest() (request *DescribeAutoScaleRecordsRequest) {
    request = &DescribeAutoScaleRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleRecords")
    
    
    return
}

func NewDescribeAutoScaleRecordsResponse() (response *DescribeAutoScaleRecordsResponse) {
    response = &DescribeAutoScaleRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleRecords
// This API is used to inquiry detailed records of cluster autoscaling.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDSTARTTIMEORENDTIME = "InvalidParameter.InvalidStartTimeOrEndTime"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleRecords(request *DescribeAutoScaleRecordsRequest) (response *DescribeAutoScaleRecordsResponse, err error) {
    return c.DescribeAutoScaleRecordsWithContext(context.Background(), request)
}

// DescribeAutoScaleRecords
// This API is used to inquiry detailed records of cluster autoscaling.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDSTARTTIMEORENDTIME = "InvalidParameter.InvalidStartTimeOrEndTime"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleRecordsWithContext(ctx context.Context, request *DescribeAutoScaleRecordsRequest) (response *DescribeAutoScaleRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeAutoScaleRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScaleStrategiesRequest() (request *DescribeAutoScaleStrategiesRequest) {
    request = &DescribeAutoScaleStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeAutoScaleStrategies")
    
    
    return
}

func NewDescribeAutoScaleStrategiesResponse() (response *DescribeAutoScaleStrategiesResponse) {
    response = &DescribeAutoScaleStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoScaleStrategies
// This API is used to access automatic scaling rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleStrategies(request *DescribeAutoScaleStrategiesRequest) (response *DescribeAutoScaleStrategiesResponse, err error) {
    return c.DescribeAutoScaleStrategiesWithContext(context.Background(), request)
}

// DescribeAutoScaleStrategies
// This API is used to access automatic scaling rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoScaleStrategiesWithContext(ctx context.Context, request *DescribeAutoScaleStrategiesRequest) (response *DescribeAutoScaleStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScaleStrategiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeAutoScaleStrategies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterFlowStatusDetailRequest() (request *DescribeClusterFlowStatusDetailRequest) {
    request = &DescribeClusterFlowStatusDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeClusterFlowStatusDetail")
    
    
    return
}

func NewDescribeClusterFlowStatusDetailResponse() (response *DescribeClusterFlowStatusDetailResponse) {
    response = &DescribeClusterFlowStatusDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterFlowStatusDetail
// This API is used to query the EMR task running details status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeClusterFlowStatusDetail(request *DescribeClusterFlowStatusDetailRequest) (response *DescribeClusterFlowStatusDetailResponse, err error) {
    return c.DescribeClusterFlowStatusDetailWithContext(context.Background(), request)
}

// DescribeClusterFlowStatusDetail
// This API is used to query the EMR task running details status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeClusterFlowStatusDetailWithContext(ctx context.Context, request *DescribeClusterFlowStatusDetailRequest) (response *DescribeClusterFlowStatusDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterFlowStatusDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeClusterFlowStatusDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterFlowStatusDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterFlowStatusDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodesRequest() (request *DescribeClusterNodesRequest) {
    request = &DescribeClusterNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeClusterNodes")
    
    
    return
}

func NewDescribeClusterNodesResponse() (response *DescribeClusterNodesResponse) {
    response = &DescribeClusterNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterNodes
// This API is used to query the information of nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
    return c.DescribeClusterNodesWithContext(context.Background(), request)
}

// DescribeClusterNodes
// This API is used to query the information of nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterNodesWithContext(ctx context.Context, request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeClusterNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDAGInfoRequest() (request *DescribeDAGInfoRequest) {
    request = &DescribeDAGInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeDAGInfo")
    
    
    return
}

func NewDescribeDAGInfoResponse() (response *DescribeDAGInfoResponse) {
    response = &DescribeDAGInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDAGInfo
// This API is used to query DAG information.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeDAGInfo(request *DescribeDAGInfoRequest) (response *DescribeDAGInfoResponse, err error) {
    return c.DescribeDAGInfoWithContext(context.Background(), request)
}

// DescribeDAGInfo
// This API is used to query DAG information.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeDAGInfoWithContext(ctx context.Context, request *DescribeDAGInfoRequest) (response *DescribeDAGInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDAGInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeDAGInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDAGInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDAGInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmrApplicationStaticsRequest() (request *DescribeEmrApplicationStaticsRequest) {
    request = &DescribeEmrApplicationStaticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeEmrApplicationStatics")
    
    
    return
}

func NewDescribeEmrApplicationStaticsResponse() (response *DescribeEmrApplicationStaticsResponse) {
    response = &DescribeEmrApplicationStaticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEmrApplicationStatics
// This API is used to query the YARN application statistics API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStatics(request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    return c.DescribeEmrApplicationStaticsWithContext(context.Background(), request)
}

// DescribeEmrApplicationStatics
// This API is used to query the YARN application statistics API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStaticsWithContext(ctx context.Context, request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    if request == nil {
        request = NewDescribeEmrApplicationStaticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeEmrApplicationStatics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmrApplicationStatics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmrApplicationStaticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmrOverviewMetricsRequest() (request *DescribeEmrOverviewMetricsRequest) {
    request = &DescribeEmrOverviewMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeEmrOverviewMetrics")
    
    
    return
}

func NewDescribeEmrOverviewMetricsResponse() (response *DescribeEmrOverviewMetricsResponse) {
    response = &DescribeEmrOverviewMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEmrOverviewMetrics
// This API is used to query the metric data on the monitoring overview page.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeEmrOverviewMetrics(request *DescribeEmrOverviewMetricsRequest) (response *DescribeEmrOverviewMetricsResponse, err error) {
    return c.DescribeEmrOverviewMetricsWithContext(context.Background(), request)
}

// DescribeEmrOverviewMetrics
// This API is used to query the metric data on the monitoring overview page.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeEmrOverviewMetricsWithContext(ctx context.Context, request *DescribeEmrOverviewMetricsRequest) (response *DescribeEmrOverviewMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeEmrOverviewMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeEmrOverviewMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmrOverviewMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmrOverviewMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalConfigRequest() (request *DescribeGlobalConfigRequest) {
    request = &DescribeGlobalConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeGlobalConfig")
    
    
    return
}

func NewDescribeGlobalConfigResponse() (response *DescribeGlobalConfigResponse) {
    response = &DescribeGlobalConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalConfig
// This API is used to query the global configurations of YARN Resource Scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeGlobalConfig(request *DescribeGlobalConfigRequest) (response *DescribeGlobalConfigResponse, err error) {
    return c.DescribeGlobalConfigWithContext(context.Background(), request)
}

// DescribeGlobalConfig
// This API is used to query the global configurations of YARN Resource Scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeGlobalConfigWithContext(ctx context.Context, request *DescribeGlobalConfigRequest) (response *DescribeGlobalConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeGlobalConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsSTDRequest() (request *DescribeGroupsSTDRequest) {
    request = &DescribeGroupsSTDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeGroupsSTD")
    
    
    return
}

func NewDescribeGroupsSTDResponse() (response *DescribeGroupsSTDResponse) {
    response = &DescribeGroupsSTDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupsSTD
// This API is used to query a user group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeGroupsSTD(request *DescribeGroupsSTDRequest) (response *DescribeGroupsSTDResponse, err error) {
    return c.DescribeGroupsSTDWithContext(context.Background(), request)
}

// DescribeGroupsSTD
// This API is used to query a user group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeGroupsSTDWithContext(ctx context.Context, request *DescribeGroupsSTDRequest) (response *DescribeGroupsSTDResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsSTDRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeGroupsSTD")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupsSTD require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupsSTDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHBaseTableOverviewRequest() (request *DescribeHBaseTableOverviewRequest) {
    request = &DescribeHBaseTableOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeHBaseTableOverview")
    
    
    return
}

func NewDescribeHBaseTableOverviewResponse() (response *DescribeHBaseTableOverviewResponse) {
    response = &DescribeHBaseTableOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHBaseTableOverview
// This API is used to access the overview of HBase table-level monitoring data.
//
// error code that may be returned:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHBaseTableOverview(request *DescribeHBaseTableOverviewRequest) (response *DescribeHBaseTableOverviewResponse, err error) {
    return c.DescribeHBaseTableOverviewWithContext(context.Background(), request)
}

// DescribeHBaseTableOverview
// This API is used to access the overview of HBase table-level monitoring data.
//
// error code that may be returned:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHBaseTableOverviewWithContext(ctx context.Context, request *DescribeHBaseTableOverviewRequest) (response *DescribeHBaseTableOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeHBaseTableOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeHBaseTableOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHBaseTableOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHBaseTableOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHDFSStorageInfoRequest() (request *DescribeHDFSStorageInfoRequest) {
    request = &DescribeHDFSStorageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeHDFSStorageInfo")
    
    
    return
}

func NewDescribeHDFSStorageInfoResponse() (response *DescribeHDFSStorageInfoResponse) {
    response = &DescribeHDFSStorageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHDFSStorageInfo
// This API is used to query information of file(s) stored in HDFS.
//
// error code that may be returned:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHDFSStorageInfo(request *DescribeHDFSStorageInfoRequest) (response *DescribeHDFSStorageInfoResponse, err error) {
    return c.DescribeHDFSStorageInfoWithContext(context.Background(), request)
}

// DescribeHDFSStorageInfo
// This API is used to query information of file(s) stored in HDFS.
//
// error code that may be returned:
//  INTERNALERROR_DOOPENTSDBREQUESTEXCEPTION = "InternalError.DoOpenTSDBRequestException"
//  INTERNALERROR_OPENTSDBHTTPRETURNCODENOTOK = "InternalError.OpenTSDBHttpReturnCodeNotOK"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeHDFSStorageInfoWithContext(ctx context.Context, request *DescribeHDFSStorageInfoRequest) (response *DescribeHDFSStorageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeHDFSStorageInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeHDFSStorageInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHDFSStorageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHDFSStorageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHiveQueriesRequest() (request *DescribeHiveQueriesRequest) {
    request = &DescribeHiveQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeHiveQueries")
    
    
    return
}

func NewDescribeHiveQueriesResponse() (response *DescribeHiveQueriesResponse) {
    response = &DescribeHiveQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHiveQueries
// This API is used to inquiry Hive query data.
//
// error code that may be returned:
//  INTERNALERROR_DBQUERYEXCEPTION = "InternalError.DBQueryException"
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeHiveQueries(request *DescribeHiveQueriesRequest) (response *DescribeHiveQueriesResponse, err error) {
    return c.DescribeHiveQueriesWithContext(context.Background(), request)
}

// DescribeHiveQueries
// This API is used to inquiry Hive query data.
//
// error code that may be returned:
//  INTERNALERROR_DBQUERYEXCEPTION = "InternalError.DBQueryException"
//  INVALIDPARAMETER_IMPALAQUERYEXCEPTION = "InvalidParameter.ImpalaQueryException"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeHiveQueriesWithContext(ctx context.Context, request *DescribeHiveQueriesRequest) (response *DescribeHiveQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeHiveQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeHiveQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHiveQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHiveQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInsightListRequest() (request *DescribeInsightListRequest) {
    request = &DescribeInsightListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInsightList")
    
    
    return
}

func NewDescribeInsightListResponse() (response *DescribeInsightListResponse) {
    response = &DescribeInsightListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInsightList
// This API is used to obtain insight result information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeInsightList(request *DescribeInsightListRequest) (response *DescribeInsightListResponse, err error) {
    return c.DescribeInsightListWithContext(context.Background(), request)
}

// DescribeInsightList
// This API is used to obtain insight result information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeInsightListWithContext(ctx context.Context, request *DescribeInsightListRequest) (response *DescribeInsightListResponse, err error) {
    if request == nil {
        request = NewDescribeInsightListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeInsightList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInsightList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInsightListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInspectionTaskResultRequest() (request *DescribeInspectionTaskResultRequest) {
    request = &DescribeInspectionTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInspectionTaskResult")
    
    
    return
}

func NewDescribeInspectionTaskResultResponse() (response *DescribeInspectionTaskResultResponse) {
    response = &DescribeInspectionTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInspectionTaskResult
// This API is used to obtain the inspection task result list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeInspectionTaskResult(request *DescribeInspectionTaskResultRequest) (response *DescribeInspectionTaskResultResponse, err error) {
    return c.DescribeInspectionTaskResultWithContext(context.Background(), request)
}

// DescribeInspectionTaskResult
// This API is used to obtain the inspection task result list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeInspectionTaskResultWithContext(ctx context.Context, request *DescribeInspectionTaskResultRequest) (response *DescribeInspectionTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeInspectionTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeInspectionTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInspectionTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInspectionTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to query the information of instances in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query the information of instances in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesListRequest() (request *DescribeInstancesListRequest) {
    request = &DescribeInstancesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeInstancesList")
    
    
    return
}

func NewDescribeInstancesListResponse() (response *DescribeInstancesListResponse) {
    response = &DescribeInstancesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesList
// This API is used to query the cluster list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeInstancesList(request *DescribeInstancesListRequest) (response *DescribeInstancesListResponse, err error) {
    return c.DescribeInstancesListWithContext(context.Background(), request)
}

// DescribeInstancesList
// This API is used to query the cluster list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBERESOURCETAGSFAILED = "FailedOperation.DescribeResourceTagsFailed"
//  FAILEDOPERATION_GETCAMROLEFAILED = "FailedOperation.GetCamRoleFailed"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeInstancesListWithContext(ctx context.Context, request *DescribeInstancesListRequest) (response *DescribeInstancesListResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeInstancesList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKyuubiQueryInfoRequest() (request *DescribeKyuubiQueryInfoRequest) {
    request = &DescribeKyuubiQueryInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeKyuubiQueryInfo")
    
    
    return
}

func NewDescribeKyuubiQueryInfoResponse() (response *DescribeKyuubiQueryInfoResponse) {
    response = &DescribeKyuubiQueryInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKyuubiQueryInfo
// This API is used to query Kyuubi query information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeKyuubiQueryInfo(request *DescribeKyuubiQueryInfoRequest) (response *DescribeKyuubiQueryInfoResponse, err error) {
    return c.DescribeKyuubiQueryInfoWithContext(context.Background(), request)
}

// DescribeKyuubiQueryInfo
// This API is used to query Kyuubi query information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeKyuubiQueryInfoWithContext(ctx context.Context, request *DescribeKyuubiQueryInfoRequest) (response *DescribeKyuubiQueryInfoResponse, err error) {
    if request == nil {
        request = NewDescribeKyuubiQueryInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeKyuubiQueryInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKyuubiQueryInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKyuubiQueryInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeDataDisksRequest() (request *DescribeNodeDataDisksRequest) {
    request = &DescribeNodeDataDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeNodeDataDisks")
    
    
    return
}

func NewDescribeNodeDataDisksResponse() (response *DescribeNodeDataDisksResponse) {
    response = &DescribeNodeDataDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeDataDisks
// This API is used to query data disk information of nodes.
//
// error code that may be returned:
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
func (c *Client) DescribeNodeDataDisks(request *DescribeNodeDataDisksRequest) (response *DescribeNodeDataDisksResponse, err error) {
    return c.DescribeNodeDataDisksWithContext(context.Background(), request)
}

// DescribeNodeDataDisks
// This API is used to query data disk information of nodes.
//
// error code that may be returned:
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
func (c *Client) DescribeNodeDataDisksWithContext(ctx context.Context, request *DescribeNodeDataDisksRequest) (response *DescribeNodeDataDisksResponse, err error) {
    if request == nil {
        request = NewDescribeNodeDataDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeNodeDataDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeDataDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeDataDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeResourceConfigFastRequest() (request *DescribeNodeResourceConfigFastRequest) {
    request = &DescribeNodeResourceConfigFastRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeNodeResourceConfigFast")
    
    
    return
}

func NewDescribeNodeResourceConfigFastResponse() (response *DescribeNodeResourceConfigFastResponse) {
    response = &DescribeNodeResourceConfigFastResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeResourceConfigFast
// This API is used to quickly obtain node specifications of the current cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTRESOURCETYPE = "ResourceUnavailable.NotSupportResourceType"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeNodeResourceConfigFast(request *DescribeNodeResourceConfigFastRequest) (response *DescribeNodeResourceConfigFastResponse, err error) {
    return c.DescribeNodeResourceConfigFastWithContext(context.Background(), request)
}

// DescribeNodeResourceConfigFast
// This API is used to quickly obtain node specifications of the current cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTRESOURCETYPE = "ResourceUnavailable.NotSupportResourceType"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeNodeResourceConfigFastWithContext(ctx context.Context, request *DescribeNodeResourceConfigFastRequest) (response *DescribeNodeResourceConfigFastResponse, err error) {
    if request == nil {
        request = NewDescribeNodeResourceConfigFastRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeNodeResourceConfigFast")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeResourceConfigFast require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeResourceConfigFastResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeSpecRequest() (request *DescribeNodeSpecRequest) {
    request = &DescribeNodeSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeNodeSpec")
    
    
    return
}

func NewDescribeNodeSpecResponse() (response *DescribeNodeSpecResponse) {
    response = &DescribeNodeSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodeSpec
// This API is used to query node specifications.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  FAILEDOPERATION_SPECDELETEDENYFORAUTOSCALESTRATEGIES = "FailedOperation.SpecDeleteDenyForAutoScaleStrategies"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_CPUTYPE = "InvalidParameter.CpuType"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDVENDORTYPE = "InvalidParameter.InvalidVendorType"
//  INVALIDPARAMETER_INVALIDVOLUMETYPE = "InvalidParameter.InvalidVolumeType"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_LESSCOMMONCOUNT = "InvalidParameter.LessCommonCount"
//  INVALIDPARAMETER_LESSTASKCOUNT = "InvalidParameter.LessTaskCount"
//  INVALIDPARAMETER_MOREMAXLIMITNUM = "InvalidParameter.MoreMaxlimitNum"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  INVALIDPARAMETER_RESOURCEPROVIDERTYPE = "InvalidParameter.ResourceProviderType"
//  INVALIDPARAMETER_RESTARTSERVICEUNSUPPORTED = "InvalidParameter.RestartServiceUnsupported"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCR = "InvalidParameter.SoftwareNotInProducr"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNSATISFIEDSOFTDEPENDECY = "InvalidParameter.UnsatisfiedSoftDependecy"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGCORERESOURCE = "MissingParameter.MissingCoreResource"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_SPECNOTFOUND = "ResourceNotFound.SpecNotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTNODETYPE = "ResourceUnavailable.NotSupportNodeType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeNodeSpec(request *DescribeNodeSpecRequest) (response *DescribeNodeSpecResponse, err error) {
    return c.DescribeNodeSpecWithContext(context.Background(), request)
}

// DescribeNodeSpec
// This API is used to query node specifications.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  FAILEDOPERATION_SPECDELETEDENYFORAUTOSCALESTRATEGIES = "FailedOperation.SpecDeleteDenyForAutoScaleStrategies"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_CPUTYPE = "InvalidParameter.CpuType"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDVENDORTYPE = "InvalidParameter.InvalidVendorType"
//  INVALIDPARAMETER_INVALIDVOLUMETYPE = "InvalidParameter.InvalidVolumeType"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_LESSCOMMONCOUNT = "InvalidParameter.LessCommonCount"
//  INVALIDPARAMETER_LESSTASKCOUNT = "InvalidParameter.LessTaskCount"
//  INVALIDPARAMETER_MOREMAXLIMITNUM = "InvalidParameter.MoreMaxlimitNum"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  INVALIDPARAMETER_RESOURCEPROVIDERTYPE = "InvalidParameter.ResourceProviderType"
//  INVALIDPARAMETER_RESTARTSERVICEUNSUPPORTED = "InvalidParameter.RestartServiceUnsupported"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCR = "InvalidParameter.SoftwareNotInProducr"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNSATISFIEDSOFTDEPENDECY = "InvalidParameter.UnsatisfiedSoftDependecy"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGCORERESOURCE = "MissingParameter.MissingCoreResource"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_SPECNOTFOUND = "ResourceNotFound.SpecNotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTNODETYPE = "ResourceUnavailable.NotSupportNodeType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhiteList"
func (c *Client) DescribeNodeSpecWithContext(ctx context.Context, request *DescribeNodeSpecRequest) (response *DescribeNodeSpecResponse, err error) {
    if request == nil {
        request = NewDescribeNodeSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeNodeSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodeSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceScheduleRequest() (request *DescribeResourceScheduleRequest) {
    request = &DescribeResourceScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeResourceSchedule")
    
    
    return
}

func NewDescribeResourceScheduleResponse() (response *DescribeResourceScheduleResponse) {
    response = &DescribeResourceScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceSchedule
// This API is used to query YARN resource scheduling information. It has been deprecated. You can use the DescribeYarnQueue API to query queue information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceSchedule(request *DescribeResourceScheduleRequest) (response *DescribeResourceScheduleResponse, err error) {
    return c.DescribeResourceScheduleWithContext(context.Background(), request)
}

// DescribeResourceSchedule
// This API is used to query YARN resource scheduling information. It has been deprecated. You can use the DescribeYarnQueue API to query queue information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleWithContext(ctx context.Context, request *DescribeResourceScheduleRequest) (response *DescribeResourceScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeResourceScheduleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeResourceSchedule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceScheduleDiffDetailRequest() (request *DescribeResourceScheduleDiffDetailRequest) {
    request = &DescribeResourceScheduleDiffDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeResourceScheduleDiffDetail")
    
    
    return
}

func NewDescribeResourceScheduleDiffDetailResponse() (response *DescribeResourceScheduleDiffDetailResponse) {
    response = &DescribeResourceScheduleDiffDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceScheduleDiffDetail
// This API is used to query change details in YARN resource scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleDiffDetail(request *DescribeResourceScheduleDiffDetailRequest) (response *DescribeResourceScheduleDiffDetailResponse, err error) {
    return c.DescribeResourceScheduleDiffDetailWithContext(context.Background(), request)
}

// DescribeResourceScheduleDiffDetail
// This API is used to query change details in YARN resource scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceScheduleDiffDetailWithContext(ctx context.Context, request *DescribeResourceScheduleDiffDetailRequest) (response *DescribeResourceScheduleDiffDetailResponse, err error) {
    if request == nil {
        request = NewDescribeResourceScheduleDiffDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeResourceScheduleDiffDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceScheduleDiffDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceScheduleDiffDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSLInstanceRequest() (request *DescribeSLInstanceRequest) {
    request = &DescribeSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSLInstance")
    
    
    return
}

func NewDescribeSLInstanceResponse() (response *DescribeSLInstanceResponse) {
    response = &DescribeSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSLInstance
// This API is used to query the basic information of Serverless HBase instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeSLInstance(request *DescribeSLInstanceRequest) (response *DescribeSLInstanceResponse, err error) {
    return c.DescribeSLInstanceWithContext(context.Background(), request)
}

// DescribeSLInstance
// This API is used to query the basic information of Serverless HBase instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeSLInstanceWithContext(ctx context.Context, request *DescribeSLInstanceRequest) (response *DescribeSLInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeSLInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeSLInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSLInstanceListRequest() (request *DescribeSLInstanceListRequest) {
    request = &DescribeSLInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSLInstanceList")
    
    
    return
}

func NewDescribeSLInstanceListResponse() (response *DescribeSLInstanceListResponse) {
    response = &DescribeSLInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSLInstanceList
// This API is used to query the detailed information of the Serverless HBase instance list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSLInstanceList(request *DescribeSLInstanceListRequest) (response *DescribeSLInstanceListResponse, err error) {
    return c.DescribeSLInstanceListWithContext(context.Background(), request)
}

// DescribeSLInstanceList
// This API is used to query the detailed information of the Serverless HBase instance list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSLInstanceListWithContext(ctx context.Context, request *DescribeSLInstanceListRequest) (response *DescribeSLInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeSLInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeSLInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSLInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSLInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceConfGroupInfosRequest() (request *DescribeServiceConfGroupInfosRequest) {
    request = &DescribeServiceConfGroupInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeServiceConfGroupInfos")
    
    
    return
}

func NewDescribeServiceConfGroupInfosResponse() (response *DescribeServiceConfGroupInfosResponse) {
    response = &DescribeServiceConfGroupInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceConfGroupInfos
// This API is used to describe service configuration group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServiceConfGroupInfos(request *DescribeServiceConfGroupInfosRequest) (response *DescribeServiceConfGroupInfosResponse, err error) {
    return c.DescribeServiceConfGroupInfosWithContext(context.Background(), request)
}

// DescribeServiceConfGroupInfos
// This API is used to describe service configuration group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDAUTORENEW = "InvalidParameter.InvalidAutoRenew"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCOMPONENT = "InvalidParameter.InvalidComponent"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDLOGINSETTING = "InvalidParameter.InvalidLoginSetting"
//  INVALIDPARAMETER_INVALIDMETATYPE = "InvalidParameter.InvalidMetaType"
//  INVALIDPARAMETER_INVALIDPASSWORD = "InvalidParameter.InvalidPassword"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDPREEXECUTEDFILE = "InvalidParameter.InvalidPreExecutedFile"
//  INVALIDPARAMETER_INVALIDPRODUCTID = "InvalidParameter.InvalidProductId"
//  INVALIDPARAMETER_INVALIDPROJECTID = "InvalidParameter.InvalidProjectId"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENAME = "InvalidParameter.InvalidServiceName"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDSOFTINFO = "InvalidParameter.InvalidSoftInfo"
//  INVALIDPARAMETER_INVALIDSOFTWARE = "InvalidParameter.InvalidSoftWare"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDSOFTWAREVERSION = "InvalidParameter.InvalidSoftWareVersion"
//  INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"
//  INVALIDPARAMETER_INVALIDSUPPORTHA = "InvalidParameter.InvalidSupportHA"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETER_UNGRANTEDPOLICY = "InvalidParameter.UngrantedPolicy"
//  INVALIDPARAMETER_UNGRANTEDROLE = "InvalidParameter.UngrantedRole"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeServiceConfGroupInfosWithContext(ctx context.Context, request *DescribeServiceConfGroupInfosRequest) (response *DescribeServiceConfGroupInfosResponse, err error) {
    if request == nil {
        request = NewDescribeServiceConfGroupInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeServiceConfGroupInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceConfGroupInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceConfGroupInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceNodeInfosRequest() (request *DescribeServiceNodeInfosRequest) {
    request = &DescribeServiceNodeInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeServiceNodeInfos")
    
    
    return
}

func NewDescribeServiceNodeInfosResponse() (response *DescribeServiceNodeInfosResponse) {
    response = &DescribeServiceNodeInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceNodeInfos
// This API is used to query service process information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeServiceNodeInfos(request *DescribeServiceNodeInfosRequest) (response *DescribeServiceNodeInfosResponse, err error) {
    return c.DescribeServiceNodeInfosWithContext(context.Background(), request)
}

// DescribeServiceNodeInfos
// This API is used to query service process information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) DescribeServiceNodeInfosWithContext(ctx context.Context, request *DescribeServiceNodeInfosRequest) (response *DescribeServiceNodeInfosResponse, err error) {
    if request == nil {
        request = NewDescribeServiceNodeInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeServiceNodeInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceNodeInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceNodeInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkApplicationsRequest() (request *DescribeSparkApplicationsRequest) {
    request = &DescribeSparkApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSparkApplications")
    
    
    return
}

func NewDescribeSparkApplicationsResponse() (response *DescribeSparkApplicationsResponse) {
    response = &DescribeSparkApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkApplications
// This API is used to obtain a Spark application list.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeSparkApplications(request *DescribeSparkApplicationsRequest) (response *DescribeSparkApplicationsResponse, err error) {
    return c.DescribeSparkApplicationsWithContext(context.Background(), request)
}

// DescribeSparkApplications
// This API is used to obtain a Spark application list.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeSparkApplicationsWithContext(ctx context.Context, request *DescribeSparkApplicationsRequest) (response *DescribeSparkApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSparkApplicationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeSparkApplications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkQueriesRequest() (request *DescribeSparkQueriesRequest) {
    request = &DescribeSparkQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeSparkQueries")
    
    
    return
}

func NewDescribeSparkQueriesResponse() (response *DescribeSparkQueriesResponse) {
    response = &DescribeSparkQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkQueries
// This API is used to query the Spark query information list.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeSparkQueries(request *DescribeSparkQueriesRequest) (response *DescribeSparkQueriesResponse, err error) {
    return c.DescribeSparkQueriesWithContext(context.Background(), request)
}

// DescribeSparkQueries
// This API is used to query the Spark query information list.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeSparkQueriesWithContext(ctx context.Context, request *DescribeSparkQueriesRequest) (response *DescribeSparkQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeSparkQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeSparkQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStarRocksQueryInfoRequest() (request *DescribeStarRocksQueryInfoRequest) {
    request = &DescribeStarRocksQueryInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeStarRocksQueryInfo")
    
    
    return
}

func NewDescribeStarRocksQueryInfoResponse() (response *DescribeStarRocksQueryInfoResponse) {
    response = &DescribeStarRocksQueryInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStarRocksQueryInfo
// This API is used to query StarRocks information.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeStarRocksQueryInfo(request *DescribeStarRocksQueryInfoRequest) (response *DescribeStarRocksQueryInfoResponse, err error) {
    return c.DescribeStarRocksQueryInfoWithContext(context.Background(), request)
}

// DescribeStarRocksQueryInfo
// This API is used to query StarRocks information.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeStarRocksQueryInfoWithContext(ctx context.Context, request *DescribeStarRocksQueryInfoRequest) (response *DescribeStarRocksQueryInfoResponse, err error) {
    if request == nil {
        request = NewDescribeStarRocksQueryInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeStarRocksQueryInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStarRocksQueryInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStarRocksQueryInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrinoQueryInfoRequest() (request *DescribeTrinoQueryInfoRequest) {
    request = &DescribeTrinoQueryInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeTrinoQueryInfo")
    
    
    return
}

func NewDescribeTrinoQueryInfoResponse() (response *DescribeTrinoQueryInfoResponse) {
    response = &DescribeTrinoQueryInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrinoQueryInfo
// This API is used to query Trino(PrestoSQL) query information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeTrinoQueryInfo(request *DescribeTrinoQueryInfoRequest) (response *DescribeTrinoQueryInfoResponse, err error) {
    return c.DescribeTrinoQueryInfoWithContext(context.Background(), request)
}

// DescribeTrinoQueryInfo
// This API is used to query Trino(PrestoSQL) query information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_APPIDMISMATCHED = "UnauthorizedOperation.AppIdMismatched"
func (c *Client) DescribeTrinoQueryInfoWithContext(ctx context.Context, request *DescribeTrinoQueryInfoRequest) (response *DescribeTrinoQueryInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTrinoQueryInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeTrinoQueryInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrinoQueryInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrinoQueryInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersForUserManagerRequest() (request *DescribeUsersForUserManagerRequest) {
    request = &DescribeUsersForUserManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeUsersForUserManager")
    
    
    return
}

func NewDescribeUsersForUserManagerResponse() (response *DescribeUsersForUserManagerResponse) {
    response = &DescribeUsersForUserManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsersForUserManager
// This API is available for clusters with OpenLDAP components configured.
//
// This API is used to export users in batches. For a Kerberos cluster, set `NeedKeytabInfo` to `true` to obtain the download link of the Keytab file. If `SupportDownLoadKeyTab` is `true`, but `DownLoadKeyTabUrl` is null, the Keytab file is not ready yet (being generated) in the backend.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeUsersForUserManager(request *DescribeUsersForUserManagerRequest) (response *DescribeUsersForUserManagerResponse, err error) {
    return c.DescribeUsersForUserManagerWithContext(context.Background(), request)
}

// DescribeUsersForUserManager
// This API is available for clusters with OpenLDAP components configured.
//
// This API is used to export users in batches. For a Kerberos cluster, set `NeedKeytabInfo` to `true` to obtain the download link of the Keytab file. If `SupportDownLoadKeyTab` is `true`, but `DownLoadKeyTabUrl` is null, the Keytab file is not ready yet (being generated) in the backend.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeUsersForUserManagerWithContext(ctx context.Context, request *DescribeUsersForUserManagerRequest) (response *DescribeUsersForUserManagerResponse, err error) {
    if request == nil {
        request = NewDescribeUsersForUserManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeUsersForUserManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersForUserManagerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYarnQueueRequest() (request *DescribeYarnQueueRequest) {
    request = &DescribeYarnQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeYarnQueue")
    
    
    return
}

func NewDescribeYarnQueueResponse() (response *DescribeYarnQueueResponse) {
    response = &DescribeYarnQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeYarnQueue
// This API is used to obtain queue information in resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeYarnQueue(request *DescribeYarnQueueRequest) (response *DescribeYarnQueueResponse, err error) {
    return c.DescribeYarnQueueWithContext(context.Background(), request)
}

// DescribeYarnQueue
// This API is used to obtain queue information in resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) DescribeYarnQueueWithContext(ctx context.Context, request *DescribeYarnQueueRequest) (response *DescribeYarnQueueResponse, err error) {
    if request == nil {
        request = NewDescribeYarnQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeYarnQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYarnQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYarnQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeYarnScheduleHistoryRequest() (request *DescribeYarnScheduleHistoryRequest) {
    request = &DescribeYarnScheduleHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "DescribeYarnScheduleHistory")
    
    
    return
}

func NewDescribeYarnScheduleHistoryResponse() (response *DescribeYarnScheduleHistoryResponse) {
    response = &DescribeYarnScheduleHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeYarnScheduleHistory
// This API is used to view the YARN resource scheduling history. It has been deprecated. You can use the Process Center to view the history records.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
func (c *Client) DescribeYarnScheduleHistory(request *DescribeYarnScheduleHistoryRequest) (response *DescribeYarnScheduleHistoryResponse, err error) {
    return c.DescribeYarnScheduleHistoryWithContext(context.Background(), request)
}

// DescribeYarnScheduleHistory
// This API is used to view the YARN resource scheduling history. It has been deprecated. You can use the Process Center to view the history records.
//
// error code that may be returned:
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
func (c *Client) DescribeYarnScheduleHistoryWithContext(ctx context.Context, request *DescribeYarnScheduleHistoryRequest) (response *DescribeYarnScheduleHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeYarnScheduleHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "DescribeYarnScheduleHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeYarnScheduleHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeYarnScheduleHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceCreateInstance")
    
    
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceCreateInstance
// This API is used to query price of instance creation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDCOMMONDISKTYPE = "InvalidParameter.InvalidCommonDiskType"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    return c.InquiryPriceCreateInstanceWithContext(context.Background(), request)
}

// InquiryPriceCreateInstance
// This API is used to query price of instance creation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCAMSERVERFAILED = "FailedOperation.GetCamServerFailed"
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_HALESSMASTERCOUNT = "InvalidParameter.HALessMasterCount"
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
//  INVALIDPARAMETER_INVALIDCOMMONDISKTYPE = "InvalidParameter.InvalidCommonDiskType"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMASTERDISKTYPE = "InvalidParameter.InvalidMasterDiskType"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSOFTWARENAME = "InvalidParameter.InvalidSoftWareName"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDUNIFYMETA = "InvalidParameter.InvalidUnifyMeta"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_NOTCONTAINMUSTSELECTSOFTWARE = "InvalidParameter.NotContainMustSelectSoftware"
//  INVALIDPARAMETER_SOFTWARENOTINPRODUCT = "InvalidParameter.SoftwareNotInProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_SUBNETNOTFOUND = "ResourceNotFound.SubnetNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceCreateInstanceWithContext(ctx context.Context, request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "InquiryPriceCreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewInstanceRequest() (request *InquiryPriceRenewInstanceRequest) {
    request = &InquiryPriceRenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceRenewInstance")
    
    
    return
}

func NewInquiryPriceRenewInstanceResponse() (response *InquiryPriceRenewInstanceResponse) {
    response = &InquiryPriceRenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewInstance
// This API is used to query the price for renewal.
//
// error code that may be returned:
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceRenewInstance(request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    return c.InquiryPriceRenewInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewInstance
// This API is used to query the price for renewal.
//
// error code that may be returned:
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_PAYMODERESOURCENOTMATCH = "InvalidParameter.PayModeResourceNotMatch"
//  INVALIDPARAMETER_PROJECTRESOURCENOTMATCH = "InvalidParameter.ProjectResourceNotMatch"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceRenewInstanceWithContext(ctx context.Context, request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "InquiryPriceRenewInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceScaleOutInstanceRequest() (request *InquiryPriceScaleOutInstanceRequest) {
    request = &InquiryPriceScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceScaleOutInstance")
    
    
    return
}

func NewInquiryPriceScaleOutInstanceResponse() (response *InquiryPriceScaleOutInstanceResponse) {
    response = &InquiryPriceScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceScaleOutInstance
// This API is used to query price of scale-out.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) InquiryPriceScaleOutInstance(request *InquiryPriceScaleOutInstanceRequest) (response *InquiryPriceScaleOutInstanceResponse, err error) {
    return c.InquiryPriceScaleOutInstanceWithContext(context.Background(), request)
}

// InquiryPriceScaleOutInstance
// This API is used to query price of scale-out.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) InquiryPriceScaleOutInstanceWithContext(ctx context.Context, request *InquiryPriceScaleOutInstanceRequest) (response *InquiryPriceScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceScaleOutInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "InquiryPriceScaleOutInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpdateInstanceRequest() (request *InquiryPriceUpdateInstanceRequest) {
    request = &InquiryPriceUpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "InquiryPriceUpdateInstance")
    
    
    return
}

func NewInquiryPriceUpdateInstanceResponse() (response *InquiryPriceUpdateInstanceResponse) {
    response = &InquiryPriceUpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceUpdateInstance
// This API is used to query price of scaling.
//
// error code that may be returned:
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceUpdateInstance(request *InquiryPriceUpdateInstanceRequest) (response *InquiryPriceUpdateInstanceResponse, err error) {
    return c.InquiryPriceUpdateInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpdateInstance
// This API is used to query price of scaling.
//
// error code that may be returned:
//  FAILEDOPERATION_GETTRADESERVERFAILED = "FailedOperation.GetTradeServerFailed"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CHECKQUOTAERR = "InternalError.CheckQuotaErr"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  INVALIDPARAMETER_ZONERESOURCENOTMATCH = "InvalidParameter.ZoneResourceNotMatch"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceUpdateInstanceWithContext(ctx context.Context, request *InquiryPriceUpdateInstanceRequest) (response *InquiryPriceUpdateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpdateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "InquiryPriceUpdateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoRenewFlag
// This API is used to introduce the prerequisite prepaid clusters.
//
// This API is used to enable or disable automatic renewal at the resource level.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// This API is used to introduce the prerequisite prepaid clusters.
//
// This API is used to enable or disable automatic renewal at the resource level.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyAutoRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoScaleStrategyRequest() (request *ModifyAutoScaleStrategyRequest) {
    request = &ModifyAutoScaleStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyAutoScaleStrategy")
    
    
    return
}

func NewModifyAutoScaleStrategyResponse() (response *ModifyAutoScaleStrategyResponse) {
    response = &ModifyAutoScaleStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoScaleStrategy
// This API is used to modify automatic scaling rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTATISTICPERIODORTRIGGERTHRESHOLD = "InvalidParameter.InvalidStatisticPeriodOrTriggerThreshold"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYPRIORITY = "InvalidParameter.InvalidStrategyPriority"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyAutoScaleStrategy(request *ModifyAutoScaleStrategyRequest) (response *ModifyAutoScaleStrategyResponse, err error) {
    return c.ModifyAutoScaleStrategyWithContext(context.Background(), request)
}

// ModifyAutoScaleStrategy
// This API is used to modify automatic scaling rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDSTATISTICPERIODORTRIGGERTHRESHOLD = "InvalidParameter.InvalidStatisticPeriodOrTriggerThreshold"
//  INVALIDPARAMETER_INVALIDSTRATEGY = "InvalidParameter.InvalidStrategy"
//  INVALIDPARAMETER_INVALIDSTRATEGYPRIORITY = "InvalidParameter.InvalidStrategyPriority"
//  INVALIDPARAMETER_INVALIDSTRATEGYSPEC = "InvalidParameter.InvalidStrategySpec"
//  INVALIDPARAMETER_INVALIDSTRATEGYTYPE = "InvalidParameter.InvalidStrategyType"
//  INVALIDPARAMETER_INVALIDTIMELAYOUT = "InvalidParameter.InvalidTimeLayout"
//  INVALIDPARAMETER_REPEATEDEXECUTIONTIME = "InvalidParameter.RepeatedExecutionTime"
//  INVALIDPARAMETER_REPEATEDSTRATEGYNAME = "InvalidParameter.RepeatedStrategyName"
//  RESOURCENOTFOUND_STRATEGYNOTFOUND = "ResourceNotFound.StrategyNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyAutoScaleStrategyWithContext(ctx context.Context, request *ModifyAutoScaleStrategyRequest) (response *ModifyAutoScaleStrategyResponse, err error) {
    if request == nil {
        request = NewModifyAutoScaleStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyAutoScaleStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoScaleStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoScaleStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalConfigRequest() (request *ModifyGlobalConfigRequest) {
    request = &ModifyGlobalConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyGlobalConfig")
    
    
    return
}

func NewModifyGlobalConfigResponse() (response *ModifyGlobalConfigResponse) {
    response = &ModifyGlobalConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalConfig
// This API is used to modify the global configuration of YARN Resource Scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyGlobalConfig(request *ModifyGlobalConfigRequest) (response *ModifyGlobalConfigResponse, err error) {
    return c.ModifyGlobalConfigWithContext(context.Background(), request)
}

// ModifyGlobalConfig
// This API is used to modify the global configuration of YARN Resource Scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyGlobalConfigWithContext(ctx context.Context, request *ModifyGlobalConfigRequest) (response *ModifyGlobalConfigResponse, err error) {
    if request == nil {
        request = NewModifyGlobalConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyGlobalConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInspectionSettingsRequest() (request *ModifyInspectionSettingsRequest) {
    request = &ModifyInspectionSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyInspectionSettings")
    
    
    return
}

func NewModifyInspectionSettingsResponse() (response *ModifyInspectionSettingsResponse) {
    response = &ModifyInspectionSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInspectionSettings
// This API is used to set inspection task configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyInspectionSettings(request *ModifyInspectionSettingsRequest) (response *ModifyInspectionSettingsResponse, err error) {
    return c.ModifyInspectionSettingsWithContext(context.Background(), request)
}

// ModifyInspectionSettings
// This API is used to set inspection task configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyInspectionSettingsWithContext(ctx context.Context, request *ModifyInspectionSettingsRequest) (response *ModifyInspectionSettingsResponse, err error) {
    if request == nil {
        request = NewModifyInspectionSettingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyInspectionSettings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInspectionSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInspectionSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceBasicRequest() (request *ModifyInstanceBasicRequest) {
    request = &ModifyInstanceBasicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyInstanceBasic")
    
    
    return
}

func NewModifyInstanceBasicResponse() (response *ModifyInstanceBasicResponse) {
    response = &ModifyInstanceBasicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceBasic
// This API is used to modify a cluster name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyInstanceBasic(request *ModifyInstanceBasicRequest) (response *ModifyInstanceBasicResponse, err error) {
    return c.ModifyInstanceBasicWithContext(context.Background(), request)
}

// ModifyInstanceBasic
// This API is used to modify a cluster name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyInstanceBasicWithContext(ctx context.Context, request *ModifyInstanceBasicRequest) (response *ModifyInstanceBasicResponse, err error) {
    if request == nil {
        request = NewModifyInstanceBasicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyInstanceBasic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceBasic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceBasicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceRequest() (request *ModifyResourceRequest) {
    request = &ModifyResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResource")
    
    
    return
}

func NewModifyResourceResponse() (response *ModifyResourceResponse) {
    response = &ModifyResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResource
// This API is used to resize an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLASSIFICATION = "InvalidParameter.InvalidClassification"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCEID = "InvalidParameter.InvalidResourceId"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyResource(request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
    return c.ModifyResourceWithContext(context.Background(), request)
}

// ModifyResource
// This API is used to resize an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLASSIFICATION = "InvalidParameter.InvalidClassification"
//  INVALIDPARAMETER_INVALIDMODIFYSPEC = "InvalidParameter.InvalidModifySpec"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCEID = "InvalidParameter.InvalidResourceId"
//  INVALIDPARAMETER_INVALIDVPCID = "InvalidParameter.InvalidVpcId"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ModifyResourceWithContext(ctx context.Context, request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
    if request == nil {
        request = NewModifyResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceScheduleConfigRequest() (request *ModifyResourceScheduleConfigRequest) {
    request = &ModifyResourceScheduleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourceScheduleConfig")
    
    
    return
}

func NewModifyResourceScheduleConfigResponse() (response *ModifyResourceScheduleConfigResponse) {
    response = &ModifyResourceScheduleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceScheduleConfig
// This API is deprecated. Use ModifyYarnQueueV2 to modify queue configuration. No related logs exist in the past one year.
//
// 
//
// This API is used to modify the resource configuration of YARN Resource Scheduling. It has been deprecated. Use the ModifyYarnQueueV2 API to modify the queue configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourceScheduleConfig(request *ModifyResourceScheduleConfigRequest) (response *ModifyResourceScheduleConfigResponse, err error) {
    return c.ModifyResourceScheduleConfigWithContext(context.Background(), request)
}

// ModifyResourceScheduleConfig
// This API is deprecated. Use ModifyYarnQueueV2 to modify queue configuration. No related logs exist in the past one year.
//
// 
//
// This API is used to modify the resource configuration of YARN Resource Scheduling. It has been deprecated. Use the ModifyYarnQueueV2 API to modify the queue configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyResourceScheduleConfigWithContext(ctx context.Context, request *ModifyResourceScheduleConfigRequest) (response *ModifyResourceScheduleConfigResponse, err error) {
    if request == nil {
        request = NewModifyResourceScheduleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyResourceScheduleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceScheduleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceScheduleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceSchedulerRequest() (request *ModifyResourceSchedulerRequest) {
    request = &ModifyResourceSchedulerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourceScheduler")
    
    
    return
}

func NewModifyResourceSchedulerResponse() (response *ModifyResourceSchedulerResponse) {
    response = &ModifyResourceSchedulerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceScheduler
// This API is used to modify a YARN resource scheduler. After the modification, you can click Deploy to bring it into effect.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyResourceScheduler(request *ModifyResourceSchedulerRequest) (response *ModifyResourceSchedulerResponse, err error) {
    return c.ModifyResourceSchedulerWithContext(context.Background(), request)
}

// ModifyResourceScheduler
// This API is used to modify a YARN resource scheduler. After the modification, you can click Deploy to bring it into effect.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) ModifyResourceSchedulerWithContext(ctx context.Context, request *ModifyResourceSchedulerRequest) (response *ModifyResourceSchedulerResponse, err error) {
    if request == nil {
        request = NewModifyResourceSchedulerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyResourceScheduler")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceScheduler require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceSchedulerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcesTagsRequest() (request *ModifyResourcesTagsRequest) {
    request = &ModifyResourcesTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyResourcesTags")
    
    
    return
}

func NewModifyResourcesTagsResponse() (response *ModifyResourcesTagsResponse) {
    response = &ModifyResourcesTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcesTags
// This API is used to forcibly modify tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourcesTags(request *ModifyResourcesTagsRequest) (response *ModifyResourcesTagsResponse, err error) {
    return c.ModifyResourcesTagsWithContext(context.Background(), request)
}

// ModifyResourcesTags
// This API is used to forcibly modify tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourcesTagsWithContext(ctx context.Context, request *ModifyResourcesTagsRequest) (response *ModifyResourcesTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourcesTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyResourcesTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcesTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcesTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySLInstanceRequest() (request *ModifySLInstanceRequest) {
    request = &ModifySLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifySLInstance")
    
    
    return
}

func NewModifySLInstanceResponse() (response *ModifySLInstanceResponse) {
    response = &ModifySLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySLInstance
// This API is used to resize a Serverless HBase instance.- If the API call is successful, a Serverless HBase instance will be created. If the instance creation request is successful, the RequestID of the request will be returned.- This is an asynchronous API. The operation is not completed immediately when the API call returns. The instance operation result can be viewed by calling DescribeInstancesList to view the StatusDesc status of the current instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) ModifySLInstance(request *ModifySLInstanceRequest) (response *ModifySLInstanceResponse, err error) {
    return c.ModifySLInstanceWithContext(context.Background(), request)
}

// ModifySLInstance
// This API is used to resize a Serverless HBase instance.- If the API call is successful, a Serverless HBase instance will be created. If the instance creation request is successful, the RequestID of the request will be returned.- This is an asynchronous API. The operation is not completed immediately when the API call returns. The instance operation result can be viewed by calling DescribeInstancesList to view the StatusDesc status of the current instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODECOUNT = "InvalidParameter.InvalidNodeCount"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) ModifySLInstanceWithContext(ctx context.Context, request *ModifySLInstanceRequest) (response *ModifySLInstanceResponse, err error) {
    if request == nil {
        request = NewModifySLInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifySLInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifySLInstanceBasicRequest() (request *ModifySLInstanceBasicRequest) {
    request = &ModifySLInstanceBasicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifySLInstanceBasic")
    
    
    return
}

func NewModifySLInstanceBasicResponse() (response *ModifySLInstanceBasicResponse) {
    response = &ModifySLInstanceBasicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySLInstanceBasic
// This API is used to modify the Serverless HBase instance name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifySLInstanceBasic(request *ModifySLInstanceBasicRequest) (response *ModifySLInstanceBasicResponse, err error) {
    return c.ModifySLInstanceBasicWithContext(context.Background(), request)
}

// ModifySLInstanceBasic
// This API is used to modify the Serverless HBase instance name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifySLInstanceBasicWithContext(ctx context.Context, request *ModifySLInstanceBasicRequest) (response *ModifySLInstanceBasicResponse, err error) {
    if request == nil {
        request = NewModifySLInstanceBasicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifySLInstanceBasic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySLInstanceBasic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySLInstanceBasicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserGroupRequest() (request *ModifyUserGroupRequest) {
    request = &ModifyUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyUserGroup")
    
    
    return
}

func NewModifyUserGroupResponse() (response *ModifyUserGroupResponse) {
    response = &ModifyUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserGroup
// This API is used to modify user groups under User Management.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    return c.ModifyUserGroupWithContext(context.Background(), request)
}

// ModifyUserGroup
// This API is used to modify user groups under User Management.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserGroupWithContext(ctx context.Context, request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
    if request == nil {
        request = NewModifyUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserManagerPwdRequest() (request *ModifyUserManagerPwdRequest) {
    request = &ModifyUserManagerPwdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyUserManagerPwd")
    
    
    return
}

func NewModifyUserManagerPwdResponse() (response *ModifyUserManagerPwdResponse) {
    response = &ModifyUserManagerPwdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserManagerPwd
// This API is used to change user password (user management).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserManagerPwd(request *ModifyUserManagerPwdRequest) (response *ModifyUserManagerPwdResponse, err error) {
    return c.ModifyUserManagerPwdWithContext(context.Background(), request)
}

// ModifyUserManagerPwd
// This API is used to change user password (user management).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUserManagerPwdWithContext(ctx context.Context, request *ModifyUserManagerPwdRequest) (response *ModifyUserManagerPwdResponse, err error) {
    if request == nil {
        request = NewModifyUserManagerPwdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyUserManagerPwd")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserManagerPwd require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserManagerPwdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsersOfGroupSTDRequest() (request *ModifyUsersOfGroupSTDRequest) {
    request = &ModifyUsersOfGroupSTDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyUsersOfGroupSTD")
    
    
    return
}

func NewModifyUsersOfGroupSTDResponse() (response *ModifyUsersOfGroupSTDResponse) {
    response = &ModifyUsersOfGroupSTDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUsersOfGroupSTD
// This API is used to change the user information of user groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUsersOfGroupSTD(request *ModifyUsersOfGroupSTDRequest) (response *ModifyUsersOfGroupSTDResponse, err error) {
    return c.ModifyUsersOfGroupSTDWithContext(context.Background(), request)
}

// ModifyUsersOfGroupSTD
// This API is used to change the user information of user groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyUsersOfGroupSTDWithContext(ctx context.Context, request *ModifyUsersOfGroupSTDRequest) (response *ModifyUsersOfGroupSTDResponse, err error) {
    if request == nil {
        request = NewModifyUsersOfGroupSTDRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyUsersOfGroupSTD")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUsersOfGroupSTD require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUsersOfGroupSTDResponse()
    err = c.Send(request, response)
    return
}

func NewModifyYarnDeployRequest() (request *ModifyYarnDeployRequest) {
    request = &ModifyYarnDeployRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyYarnDeploy")
    
    
    return
}

func NewModifyYarnDeployResponse() (response *ModifyYarnDeployResponse) {
    response = &ModifyYarnDeployResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyYarnDeploy
// This API is deprecated. Use DeployYarnConf to bring configurations into effect after deployment.
//
// 
//
// This API is used to bring configurations into effect after deployment. It has been deprecated. Use the DeployYarnConf API to bring configurations into effect after deployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyYarnDeploy(request *ModifyYarnDeployRequest) (response *ModifyYarnDeployResponse, err error) {
    return c.ModifyYarnDeployWithContext(context.Background(), request)
}

// ModifyYarnDeploy
// This API is deprecated. Use DeployYarnConf to bring configurations into effect after deployment.
//
// 
//
// This API is used to bring configurations into effect after deployment. It has been deprecated. Use the DeployYarnConf API to bring configurations into effect after deployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyYarnDeployWithContext(ctx context.Context, request *ModifyYarnDeployRequest) (response *ModifyYarnDeployResponse, err error) {
    if request == nil {
        request = NewModifyYarnDeployRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyYarnDeploy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyYarnDeploy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyYarnDeployResponse()
    err = c.Send(request, response)
    return
}

func NewModifyYarnQueueV2Request() (request *ModifyYarnQueueV2Request) {
    request = &ModifyYarnQueueV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ModifyYarnQueueV2")
    
    
    return
}

func NewModifyYarnQueueV2Response() (response *ModifyYarnQueueV2Response) {
    response = &ModifyYarnQueueV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyYarnQueueV2
// This API is used to modify queue information in resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyYarnQueueV2(request *ModifyYarnQueueV2Request) (response *ModifyYarnQueueV2Response, err error) {
    return c.ModifyYarnQueueV2WithContext(context.Background(), request)
}

// ModifyYarnQueueV2
// This API is used to modify queue information in resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ModifyYarnQueueV2WithContext(ctx context.Context, request *ModifyYarnQueueV2Request) (response *ModifyYarnQueueV2Response, err error) {
    if request == nil {
        request = NewModifyYarnQueueV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ModifyYarnQueueV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyYarnQueueV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyYarnQueueV2Response()
    err = c.Send(request, response)
    return
}

func NewResetYarnConfigRequest() (request *ResetYarnConfigRequest) {
    request = &ResetYarnConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ResetYarnConfig")
    
    
    return
}

func NewResetYarnConfigResponse() (response *ResetYarnConfigResponse) {
    response = &ResetYarnConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetYarnConfig
// This API is used to modify the resource configuration of YARN resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResetYarnConfig(request *ResetYarnConfigRequest) (response *ResetYarnConfigResponse, err error) {
    return c.ResetYarnConfigWithContext(context.Background(), request)
}

// ResetYarnConfig
// This API is used to modify the resource configuration of YARN resource scheduling.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResetYarnConfigWithContext(ctx context.Context, request *ResetYarnConfigRequest) (response *ResetYarnConfigResponse, err error) {
    if request == nil {
        request = NewResetYarnConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ResetYarnConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetYarnConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetYarnConfigResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDataDisksRequest() (request *ResizeDataDisksRequest) {
    request = &ResizeDataDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ResizeDataDisks")
    
    
    return
}

func NewResizeDataDisksResponse() (response *ResizeDataDisksResponse) {
    response = &ResizeDataDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeDataDisks
// This API is used to scale out the cloud data disk.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResizeDataDisks(request *ResizeDataDisksRequest) (response *ResizeDataDisksResponse, err error) {
    return c.ResizeDataDisksWithContext(context.Background(), request)
}

// ResizeDataDisks
// This API is used to scale out the cloud data disk.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) ResizeDataDisksWithContext(ctx context.Context, request *ResizeDataDisksRequest) (response *ResizeDataDisksResponse, err error) {
    if request == nil {
        request = NewResizeDataDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ResizeDataDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDataDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDataDisksResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutClusterRequest() (request *ScaleOutClusterRequest) {
    request = &ScaleOutClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ScaleOutCluster")
    
    
    return
}

func NewScaleOutClusterResponse() (response *ScaleOutClusterResponse) {
    response = &ScaleOutClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutCluster
// This API is used to scale out a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
func (c *Client) ScaleOutCluster(request *ScaleOutClusterRequest) (response *ScaleOutClusterResponse, err error) {
    return c.ScaleOutClusterWithContext(context.Background(), request)
}

// ScaleOutCluster
// This API is used to scale out a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
func (c *Client) ScaleOutClusterWithContext(ctx context.Context, request *ScaleOutClusterRequest) (response *ScaleOutClusterResponse, err error) {
    if request == nil {
        request = NewScaleOutClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ScaleOutCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutClusterResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// This API is used to scale out instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKIFSUPPORTPODSTRETCH = "FailedOperation.CheckIfSupportPodStretch"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  FAILEDOPERATION_NOTSUPPORTPOD = "FailedOperation.NotSupportPod"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLICKHOUSECLUSTER = "InvalidParameter.InvalidClickHouseCluster"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNT = "InvalidParameter.InvalidCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDCUSTOMIZEDPODPARAM = "InvalidParameter.InvalidCustomizedPodParam"
//  INVALIDPARAMETER_INVALIDEKSINSTANCE = "InvalidParameter.InvalidEksInstance"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENODEINFO = "InvalidParameter.InvalidServiceNodeInfo"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDTASKCOUNT = "InvalidParameter.InvalidTaskCount"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDTKEINSTANCE = "InvalidParameter.InvalidTkeInstance"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TKEPRECONDITIONNOTFOUND = "ResourceNotFound.TKEPreconditionNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// This API is used to scale out instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKIFSUPPORTPODSTRETCH = "FailedOperation.CheckIfSupportPodStretch"
//  FAILEDOPERATION_DUPLICATEORDERNOTALLOWED = "FailedOperation.DuplicateOrderNotAllowed"
//  FAILEDOPERATION_GETCVMCONFIGQUOTAFAILED = "FailedOperation.GetCvmConfigQuotaFailed"
//  FAILEDOPERATION_GETCVMSERVERFAILED = "FailedOperation.GetCvmServerFailed"
//  FAILEDOPERATION_NOTSUPPORTPOD = "FailedOperation.NotSupportPod"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_EKSERROR = "InternalError.EKSError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TKEERROR = "InternalError.TKEError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLICKHOUSECLUSTER = "InvalidParameter.InvalidClickHouseCluster"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDCORECOUNT = "InvalidParameter.InvalidCoreCount"
//  INVALIDPARAMETER_INVALIDCOUNT = "InvalidParameter.InvalidCount"
//  INVALIDPARAMETER_INVALIDCOUNTNUM = "InvalidParameter.InvalidCountNum"
//  INVALIDPARAMETER_INVALIDCUSTOMIZEDPODPARAM = "InvalidParameter.InvalidCustomizedPodParam"
//  INVALIDPARAMETER_INVALIDEKSINSTANCE = "InvalidParameter.InvalidEksInstance"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPaymode"
//  INVALIDPARAMETER_INVALIDRESOURCESPEC = "InvalidParameter.InvalidResourceSpec"
//  INVALIDPARAMETER_INVALIDSERCURITYGRPUPID = "InvalidParameter.InvalidSercurityGrpupId"
//  INVALIDPARAMETER_INVALIDSERVICENODEINFO = "InvalidParameter.InvalidServiceNodeInfo"
//  INVALIDPARAMETER_INVALIDSOFTDEPLOYINFO = "InvalidParameter.InvalidSoftDeployInfo"
//  INVALIDPARAMETER_INVALIDTASKCOUNT = "InvalidParameter.InvalidTaskCount"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_INVALIDTKEINSTANCE = "InvalidParameter.InvalidTkeInstance"
//  INVALIDPARAMETERVALUE_INVALIDTKEINSTANCE = "InvalidParameterValue.InvalidTkeInstance"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_DISKINSUFFICIENT = "ResourceInsufficient.DiskInsufficient"
//  RESOURCEINSUFFICIENT_INSTANCEINSUFFICIENT = "ResourceInsufficient.InstanceInsufficient"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TKEPRECONDITIONNOTFOUND = "ResourceNotFound.TKEPreconditionNotFound"
//  RESOURCENOTFOUND_TAGSNOTFOUND = "ResourceNotFound.TagsNotFound"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTDEFAULTSPEC = "ResourceUnavailable.ResourceSpecNotDefaultSpec"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_CBSSOLDOUT = "ResourcesSoldOut.CbsSoldOut"
//  RESOURCESSOLDOUT_CVMSOLDOUT = "ResourcesSoldOut.CvmSoldOut"
func (c *Client) ScaleOutInstanceWithContext(ctx context.Context, request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "ScaleOutInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetNodeResourceConfigDefaultRequest() (request *SetNodeResourceConfigDefaultRequest) {
    request = &SetNodeResourceConfigDefaultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "SetNodeResourceConfigDefault")
    
    
    return
}

func NewSetNodeResourceConfigDefaultResponse() (response *SetNodeResourceConfigDefaultResponse) {
    response = &SetNodeResourceConfigDefaultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNodeResourceConfigDefault
// This API is used to set specifications for a node in the current cluster to default or not.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) SetNodeResourceConfigDefault(request *SetNodeResourceConfigDefaultRequest) (response *SetNodeResourceConfigDefaultResponse, err error) {
    return c.SetNodeResourceConfigDefaultWithContext(context.Background(), request)
}

// SetNodeResourceConfigDefault
// This API is used to set specifications for a node in the current cluster to default or not.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_RESOURCESPECNOTEXIST = "ResourceUnavailable.ResourceSpecNotExist"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
func (c *Client) SetNodeResourceConfigDefaultWithContext(ctx context.Context, request *SetNodeResourceConfigDefaultRequest) (response *SetNodeResourceConfigDefaultResponse, err error) {
    if request == nil {
        request = NewSetNodeResourceConfigDefaultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "SetNodeResourceConfigDefault")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNodeResourceConfigDefault require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNodeResourceConfigDefaultResponse()
    err = c.Send(request, response)
    return
}

func NewStartStopServiceOrMonitorRequest() (request *StartStopServiceOrMonitorRequest) {
    request = &StartStopServiceOrMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "StartStopServiceOrMonitor")
    
    
    return
}

func NewStartStopServiceOrMonitorResponse() (response *StartStopServiceOrMonitorResponse) {
    response = &StartStopServiceOrMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStopServiceOrMonitor
// This API is used to start, stop, or restart services.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) StartStopServiceOrMonitor(request *StartStopServiceOrMonitorRequest) (response *StartStopServiceOrMonitorResponse, err error) {
    return c.StartStopServiceOrMonitorWithContext(context.Background(), request)
}

// StartStopServiceOrMonitor
// This API is used to start, stop, or restart services.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) StartStopServiceOrMonitorWithContext(ctx context.Context, request *StartStopServiceOrMonitorRequest) (response *StartStopServiceOrMonitorResponse, err error) {
    if request == nil {
        request = NewStartStopServiceOrMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "StartStopServiceOrMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStopServiceOrMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStopServiceOrMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateClusterNodesRequest() (request *TerminateClusterNodesRequest) {
    request = &TerminateClusterNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateClusterNodes")
    
    
    return
}

func NewTerminateClusterNodesResponse() (response *TerminateClusterNodesResponse) {
    response = &TerminateClusterNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateClusterNodes
// This API is used to terminate cluster nodes.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CVMINSTANCENOTFOUND = "ResourceNotFound.CvmInstanceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) TerminateClusterNodes(request *TerminateClusterNodesRequest) (response *TerminateClusterNodesResponse, err error) {
    return c.TerminateClusterNodesWithContext(context.Background(), request)
}

// TerminateClusterNodes
// This API is used to terminate cluster nodes.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDNODEFLAG = "InvalidParameter.InvalidNodeFlag"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CVMINSTANCENOTFOUND = "ResourceNotFound.CvmInstanceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) TerminateClusterNodesWithContext(ctx context.Context, request *TerminateClusterNodesRequest) (response *TerminateClusterNodesResponse, err error) {
    if request == nil {
        request = NewTerminateClusterNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "TerminateClusterNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateClusterNodesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstanceRequest() (request *TerminateInstanceRequest) {
    request = &TerminateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateInstance")
    
    
    return
}

func NewTerminateInstanceResponse() (response *TerminateInstanceResponse) {
    response = &TerminateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstance
// This API is used to terminate EMR instances. It is only supported in the official paid edition of EMR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
    return c.TerminateInstanceWithContext(context.Background(), request)
}

// TerminateInstance
// This API is used to terminate EMR instances. It is only supported in the official paid edition of EMR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateInstanceWithContext(ctx context.Context, request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "TerminateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateSLInstanceRequest() (request *TerminateSLInstanceRequest) {
    request = &TerminateSLInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateSLInstance")
    
    
    return
}

func NewTerminateSLInstanceResponse() (response *TerminateSLInstanceResponse) {
    response = &TerminateSLInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateSLInstance
// This API is used to terminate a Serverless HBase instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateSLInstance(request *TerminateSLInstanceRequest) (response *TerminateSLInstanceResponse, err error) {
    return c.TerminateSLInstanceWithContext(context.Background(), request)
}

// TerminateSLInstance
// This API is used to terminate a Serverless HBase instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REFUNDCVMFAILED = "FailedOperation.RefundCvmFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_CHECKCAMAUTH = "UnauthorizedOperation.CheckCamAuth"
//  UNSUPPORTEDOPERATION_SERVICENOTSUPPORT = "UnsupportedOperation.ServiceNotSupport"
func (c *Client) TerminateSLInstanceWithContext(ctx context.Context, request *TerminateSLInstanceRequest) (response *TerminateSLInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateSLInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "TerminateSLInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateSLInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTasksRequest() (request *TerminateTasksRequest) {
    request = &TerminateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("emr", APIVersion, "TerminateTasks")
    
    
    return
}

func NewTerminateTasksResponse() (response *TerminateTasksResponse) {
    response = &TerminateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateTasks
// This API is used to terminate a task node.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateTasks(request *TerminateTasksRequest) (response *TerminateTasksResponse, err error) {
    return c.TerminateTasksWithContext(context.Background(), request)
}

// TerminateTasks
// This API is used to terminate a task node.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCGWERROR = "InternalError.AccountCgwError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_CAMERROR = "InternalError.CamError"
//  INTERNALERROR_CBSCGWERROR = "InternalError.CbsCgwError"
//  INTERNALERROR_CBSERROR = "InternalError.CbsError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_CONFIGCGWERROR = "InternalError.ConfigCgwError"
//  INTERNALERROR_CVMERROR = "InternalError.CvmError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_PROJECTCGWERROR = "InternalError.ProjectCgwError"
//  INTERNALERROR_SGERROR = "InternalError.SgError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_TRADECGWERROR = "InternalError.TradeCgwError"
//  INTERNALERROR_VPCCGWERROR = "InternalError.VpcCgwError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDJOBFLOW = "InvalidParameter.InvalidJobFlow"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateTasksWithContext(ctx context.Context, request *TerminateTasksRequest) (response *TerminateTasksResponse, err error) {
    if request == nil {
        request = NewTerminateTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "emr", APIVersion, "TerminateTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTasksResponse()
    err = c.Send(request, response)
    return
}
