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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) AddUsersForUserManagerWithContext(ctx context.Context, request *AddUsersForUserManagerRequest) (response *AddUsersForUserManagerResponse, err error) {
    if request == nil {
        request = NewAddUsersForUserManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUsersForUserManagerResponse()
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
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
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
//  INVALIDPARAMETER_INVALIDDEPENDSERVICEANDENABLEKERBEROSCONFLICT = "InvalidParameter.InvalidDependServiceAndEnableKerberosConflict"
//  INVALIDPARAMETER_INVALIDDISKNUM = "InvalidParameter.InvalidDiskNum"
//  INVALIDPARAMETER_INVALIDINSTANCECHARGETYPE = "InvalidParameter.InvalidInstanceChargeType"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
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
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
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
//  INVALIDPARAMETER_INVALIDCOSFILEURI = "InvalidParameter.InvalidCosFileURI"
//  INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"
//  INVALIDPARAMETER_INVALIDEXTENDFIELD = "InvalidParameter.InvalidExtendField"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDINSTANCETYPE = "InvalidParameter.InvalidInstanceType"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSLInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSLInstanceResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScaleRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScaleRecordsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodesResponse()
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
//  This API is used to query the Yarn application statistics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStatics(request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    return c.DescribeEmrApplicationStaticsWithContext(context.Background(), request)
}

// DescribeEmrApplicationStatics
//  This API is used to query the Yarn application statistics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeEmrApplicationStaticsWithContext(ctx context.Context, request *DescribeEmrApplicationStaticsRequest) (response *DescribeEmrApplicationStaticsResponse, err error) {
    if request == nil {
        request = NewDescribeEmrApplicationStaticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmrApplicationStatics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmrApplicationStaticsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHiveQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHiveQueriesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesListResponse()
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
// This API is used to query the data of YARN Resource Scheduling.
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
// This API is used to query the data of YARN Resource Scheduling.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceScheduleResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSLInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSLInstanceListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsersForUserManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersForUserManagerResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpdateInstanceResponse()
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
// This API is used to modify the resource configuration of YARN Resource Scheduling.
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
// This API is used to modify the resource configuration of YARN Resource Scheduling.
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
// This API is used to modify the YARN resource scheduler (the change will take effect after you click Apply).
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
// This API is used to modify the YARN resource scheduler (the change will take effect after you click Apply).
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySLInstanceBasic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySLInstanceBasicResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserManagerPwd require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserManagerPwdResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTasksResponse()
    err = c.Send(request, response)
    return
}
