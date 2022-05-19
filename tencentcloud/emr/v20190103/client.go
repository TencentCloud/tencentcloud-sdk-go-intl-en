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
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
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
//  INVALIDPARAMETER_INCORRECTCOMMONCOUNT = "InvalidParameter.IncorrectCommonCount"
//  INVALIDPARAMETER_INCORRECTMASTERCOUNT = "InvalidParameter.IncorrectMasterCount"
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
// This API is used to query the information of a hardware node.
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
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterNodes(request *DescribeClusterNodesRequest) (response *DescribeClusterNodesResponse, err error) {
    return c.DescribeClusterNodesWithContext(context.Background(), request)
}

// DescribeClusterNodes
// This API is used to query the information of a hardware node.
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
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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
// This API is used to query EMR instances.
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
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query EMR instances.
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
//  INVALIDPARAMETER_DISPLAYSTRATEGYNOTMATCH = "InvalidParameter.DisplayStrategyNotMatch"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDINSTANCENAME = "InvalidParameter.InvalidInstanceName"
//  INVALIDPARAMETER_ORDERFIELDNOTMATCH = "InvalidParameter.OrderFieldNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_HARDWAREINFONOTFOUND = "ResourceNotFound.HardwareInfoNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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
// This API is used to get data from the YARN Resource Scheduling page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeResourceSchedule(request *DescribeResourceScheduleRequest) (response *DescribeResourceScheduleResponse, err error) {
    return c.DescribeResourceScheduleWithContext(context.Background(), request)
}

// DescribeResourceSchedule
// This API is used to get data from the YARN Resource Scheduling page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) InquiryPriceRenewInstance(request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    return c.InquiryPriceRenewInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewInstance
// This API is used to query the price for renewal.
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyResourceScheduleConfig(request *ModifyResourceScheduleConfigRequest) (response *ModifyResourceScheduleConfigResponse, err error) {
    return c.ModifyResourceScheduleConfigWithContext(context.Background(), request)
}

// ModifyResourceScheduleConfig
// This API is used to modify the resource configuration of YARN Resource Scheduling.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_WOODSERVERERROR = "InternalError.WoodServerError"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyResourceScheduler(request *ModifyResourceSchedulerRequest) (response *ModifyResourceSchedulerResponse, err error) {
    return c.ModifyResourceSchedulerWithContext(context.Background(), request)
}

// ModifyResourceScheduler
// This API is used to modify the YARN resource scheduler (the change will take effect after you click Apply).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
