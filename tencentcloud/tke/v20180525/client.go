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

package v20180525

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-05-25"

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


func NewAcquireClusterAdminRoleRequest() (request *AcquireClusterAdminRoleRequest) {
    request = &AcquireClusterAdminRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AcquireClusterAdminRole")
    
    
    return
}

func NewAcquireClusterAdminRoleResponse() (response *AcquireClusterAdminRoleResponse) {
    response = &AcquireClusterAdminRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcquireClusterAdminRole
// This API can be called to acquire the ClusterRole tke:admin. By setting a CAM policy, you can grant permission of this API to a sub-account that has higher permission in CAM. In this way, this sub-account can call this API directly to acquire the admin role of a Kubernetes cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) AcquireClusterAdminRole(request *AcquireClusterAdminRoleRequest) (response *AcquireClusterAdminRoleResponse, err error) {
    return c.AcquireClusterAdminRoleWithContext(context.Background(), request)
}

// AcquireClusterAdminRole
// This API can be called to acquire the ClusterRole tke:admin. By setting a CAM policy, you can grant permission of this API to a sub-account that has higher permission in CAM. In this way, this sub-account can call this API directly to acquire the admin role of a Kubernetes cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) AcquireClusterAdminRoleWithContext(ctx context.Context, request *AcquireClusterAdminRoleRequest) (response *AcquireClusterAdminRoleResponse, err error) {
    if request == nil {
        request = NewAcquireClusterAdminRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcquireClusterAdminRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcquireClusterAdminRoleResponse()
    err = c.Send(request, response)
    return
}

func NewAddExistedInstancesRequest() (request *AddExistedInstancesRequest) {
    request = &AddExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddExistedInstances")
    
    
    return
}

func NewAddExistedInstancesResponse() (response *AddExistedInstancesResponse) {
    response = &AddExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddExistedInstances
// This API is used to add one or more existing instances to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_NODEEXISTSSECONDARYNETWORKINTERFACE = "FailedOperation.NodeExistsSecondaryNetworkInterface"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) AddExistedInstances(request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
    return c.AddExistedInstancesWithContext(context.Background(), request)
}

// AddExistedInstances
// This API is used to add one or more existing instances to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_NODEEXISTSSECONDARYNETWORKINTERFACE = "FailedOperation.NodeExistsSecondaryNetworkInterface"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) AddExistedInstancesWithContext(ctx context.Context, request *AddExistedInstancesRequest) (response *AddExistedInstancesResponse, err error) {
    if request == nil {
        request = NewAddExistedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddExistedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddNodeToNodePoolRequest() (request *AddNodeToNodePoolRequest) {
    request = &AddNodeToNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddNodeToNodePool")
    
    
    return
}

func NewAddNodeToNodePoolResponse() (response *AddNodeToNodePoolResponse) {
    response = &AddNodeToNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNodeToNodePool
// This API is used to move nodes in a cluster to a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddNodeToNodePool(request *AddNodeToNodePoolRequest) (response *AddNodeToNodePoolResponse, err error) {
    return c.AddNodeToNodePoolWithContext(context.Background(), request)
}

// AddNodeToNodePool
// This API is used to move nodes in a cluster to a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddNodeToNodePoolWithContext(ctx context.Context, request *AddNodeToNodePoolRequest) (response *AddNodeToNodePoolResponse, err error) {
    if request == nil {
        request = NewAddNodeToNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNodeToNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNodeToNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewAddVpcCniSubnetsRequest() (request *AddVpcCniSubnetsRequest) {
    request = &AddVpcCniSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "AddVpcCniSubnets")
    
    
    return
}

func NewAddVpcCniSubnetsResponse() (response *AddVpcCniSubnetsResponse) {
    response = &AddVpcCniSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddVpcCniSubnets
// This API is used to add subnets in the container network for a VPC-CNI cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDVPCCNISUBNETSFAILED = "FailedOperation.AddVpcCniSubnetsFailed"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddVpcCniSubnets(request *AddVpcCniSubnetsRequest) (response *AddVpcCniSubnetsResponse, err error) {
    return c.AddVpcCniSubnetsWithContext(context.Background(), request)
}

// AddVpcCniSubnets
// This API is used to add subnets in the container network for a VPC-CNI cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDVPCCNISUBNETSFAILED = "FailedOperation.AddVpcCniSubnetsFailed"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) AddVpcCniSubnetsWithContext(ctx context.Context, request *AddVpcCniSubnetsRequest) (response *AddVpcCniSubnetsResponse, err error) {
    if request == nil {
        request = NewAddVpcCniSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddVpcCniSubnets require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddVpcCniSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckEdgeClusterCIDRRequest() (request *CheckEdgeClusterCIDRRequest) {
    request = &CheckEdgeClusterCIDRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CheckEdgeClusterCIDR")
    
    
    return
}

func NewCheckEdgeClusterCIDRResponse() (response *CheckEdgeClusterCIDRResponse) {
    response = &CheckEdgeClusterCIDRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckEdgeClusterCIDR
// This API is used to check if the CIDR block of a TKE Edge cluster conflicts with other CIDR blocks.
//
// error code that may be returned:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INTERNALERROR_VSTATIONERROR = "InternalError.VstationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckEdgeClusterCIDR(request *CheckEdgeClusterCIDRRequest) (response *CheckEdgeClusterCIDRResponse, err error) {
    return c.CheckEdgeClusterCIDRWithContext(context.Background(), request)
}

// CheckEdgeClusterCIDR
// This API is used to check if the CIDR block of a TKE Edge cluster conflicts with other CIDR blocks.
//
// error code that may be returned:
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_CMDTIMEOUT = "InternalError.CmdTimeout"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INTERNALERROR_VSTATIONERROR = "InternalError.VstationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckEdgeClusterCIDRWithContext(ctx context.Context, request *CheckEdgeClusterCIDRRequest) (response *CheckEdgeClusterCIDRResponse, err error) {
    if request == nil {
        request = NewCheckEdgeClusterCIDRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckEdgeClusterCIDR require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckEdgeClusterCIDRResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstancesUpgradeAbleRequest() (request *CheckInstancesUpgradeAbleRequest) {
    request = &CheckInstancesUpgradeAbleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CheckInstancesUpgradeAble")
    
    
    return
}

func NewCheckInstancesUpgradeAbleResponse() (response *CheckInstancesUpgradeAbleResponse) {
    response = &CheckInstancesUpgradeAbleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckInstancesUpgradeAble
// This API is used to query nodes eligible for an upgrade in the given node list.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) CheckInstancesUpgradeAble(request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
    return c.CheckInstancesUpgradeAbleWithContext(context.Background(), request)
}

// CheckInstancesUpgradeAble
// This API is used to query nodes eligible for an upgrade in the given node list.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) CheckInstancesUpgradeAbleWithContext(ctx context.Context, request *CheckInstancesUpgradeAbleRequest) (response *CheckInstancesUpgradeAbleResponse, err error) {
    if request == nil {
        request = NewCheckInstancesUpgradeAbleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstancesUpgradeAble require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckInstancesUpgradeAbleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupStorageLocationRequest() (request *CreateBackupStorageLocationRequest) {
    request = &CreateBackupStorageLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateBackupStorageLocation")
    
    
    return
}

func NewCreateBackupStorageLocationResponse() (response *CreateBackupStorageLocationResponse) {
    response = &CreateBackupStorageLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupStorageLocation
// This API is used to create a backup repository. You can specify the storage type (such as COS), the bucket region and the name. Up to 100 repositories can be created. Note that the settings of this API apply globally. You only need to create one backup repository, and back up TKE clusters in different regions in it.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateBackupStorageLocation(request *CreateBackupStorageLocationRequest) (response *CreateBackupStorageLocationResponse, err error) {
    return c.CreateBackupStorageLocationWithContext(context.Background(), request)
}

// CreateBackupStorageLocation
// This API is used to create a backup repository. You can specify the storage type (such as COS), the bucket region and the name. Up to 100 repositories can be created. Note that the settings of this API apply globally. You only need to create one backup repository, and back up TKE clusters in different regions in it.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateBackupStorageLocationWithContext(ctx context.Context, request *CreateBackupStorageLocationRequest) (response *CreateBackupStorageLocationResponse, err error) {
    if request == nil {
        request = NewCreateBackupStorageLocationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupStorageLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupStorageLocationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSLogConfigRequest() (request *CreateCLSLogConfigRequest) {
    request = &CreateCLSLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateCLSLogConfig")
    
    
    return
}

func NewCreateCLSLogConfigResponse() (response *CreateCLSLogConfigResponse) {
    response = &CreateCLSLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSLogConfig
// This API is used to create log collection configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSCLIENT = "FailedOperation.CreateClsClient"
//  FAILEDOPERATION_CREATECLSCONFIG = "FailedOperation.CreateClsConfig"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSLOGSET = "FailedOperation.CreateClsLogSet"
//  FAILEDOPERATION_CREATECLSMACHINEGROUP = "FailedOperation.CreateClsMachineGroup"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSCONFIG = "FailedOperation.GetClsConfig"
//  FAILEDOPERATION_GETCLSCONFIGMACHINEGROUPS = "FailedOperation.GetClsConfigMachineGroups"
//  FAILEDOPERATION_GETCLSINDEX = "FailedOperation.GetClsIndex"
//  FAILEDOPERATION_GETCLSLOGSET = "FailedOperation.GetClsLogSet"
//  FAILEDOPERATION_GETCLSMACHINEGROUP = "FailedOperation.GetClsMachineGroup"
//  FAILEDOPERATION_GETCLSMACHINEGROUPCONFIGS = "FailedOperation.GetClsMachineGroupConfigs"
//  FAILEDOPERATION_GETCLSTOPIC = "FailedOperation.GetClsTopic"
//  FAILEDOPERATION_K8SCLIENTBUILDERROR = "FailedOperation.K8sClientBuildError"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_MODIFYCLSCONFIG = "FailedOperation.ModifyClsConfig"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  FAILEDOPERATION_MODIFYCLSTOPIC = "FailedOperation.ModifyClsTopic"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreateCLSLogConfig(request *CreateCLSLogConfigRequest) (response *CreateCLSLogConfigResponse, err error) {
    return c.CreateCLSLogConfigWithContext(context.Background(), request)
}

// CreateCLSLogConfig
// This API is used to create log collection configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSCLIENT = "FailedOperation.CreateClsClient"
//  FAILEDOPERATION_CREATECLSCONFIG = "FailedOperation.CreateClsConfig"
//  FAILEDOPERATION_CREATECLSINDEX = "FailedOperation.CreateClsIndex"
//  FAILEDOPERATION_CREATECLSLOGSET = "FailedOperation.CreateClsLogSet"
//  FAILEDOPERATION_CREATECLSMACHINEGROUP = "FailedOperation.CreateClsMachineGroup"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSCONFIG = "FailedOperation.GetClsConfig"
//  FAILEDOPERATION_GETCLSCONFIGMACHINEGROUPS = "FailedOperation.GetClsConfigMachineGroups"
//  FAILEDOPERATION_GETCLSINDEX = "FailedOperation.GetClsIndex"
//  FAILEDOPERATION_GETCLSLOGSET = "FailedOperation.GetClsLogSet"
//  FAILEDOPERATION_GETCLSMACHINEGROUP = "FailedOperation.GetClsMachineGroup"
//  FAILEDOPERATION_GETCLSMACHINEGROUPCONFIGS = "FailedOperation.GetClsMachineGroupConfigs"
//  FAILEDOPERATION_GETCLSTOPIC = "FailedOperation.GetClsTopic"
//  FAILEDOPERATION_K8SCLIENTBUILDERROR = "FailedOperation.K8sClientBuildError"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_MODIFYCLSCONFIG = "FailedOperation.ModifyClsConfig"
//  FAILEDOPERATION_MODIFYCLSINDEX = "FailedOperation.ModifyClsIndex"
//  FAILEDOPERATION_MODIFYCLSTOPIC = "FailedOperation.ModifyClsTopic"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreateCLSLogConfigWithContext(ctx context.Context, request *CreateCLSLogConfigRequest) (response *CreateCLSLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateCLSLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// This API is used to create a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_ACCOUNTUSERNOTAUTHENTICATED = "FailedOperation.AccountUserNotAuthenticated"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_CVMNUMBERNOTMATCH = "FailedOperation.CvmNumberNotMatch"
//  FAILEDOPERATION_CVMVPCIDNOTMATCH = "FailedOperation.CvmVpcidNotMatch"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_DFWGETUSGQUOTA = "FailedOperation.DfwGetUSGQuota"
//  FAILEDOPERATION_NODEEXISTSSECONDARYNETWORKINTERFACE = "FailedOperation.NodeExistsSecondaryNetworkInterface"
//  FAILEDOPERATION_OSNOTSUPPORT = "FailedOperation.OsNotSupport"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXCLSLIMIT = "FailedOperation.QuotaMaxClsLimit"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  FAILEDOPERATION_QUOTAUSGLIMIT = "FailedOperation.QuotaUSGLimit"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  FAILEDOPERATION_WHITELISTUNEXPECTEDERROR = "FailedOperation.WhitelistUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNUMBERNOTMATCH = "InternalError.CvmNumberNotMatch"
//  INTERNALERROR_CVMSTATUS = "InternalError.CvmStatus"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAUSGLIMIT = "InternalError.QuotaUSGLimit"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_CIDRINVALI = "InvalidParameter.CidrInvali"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_INVALIDPRIVATENETWORKCIDR = "InvalidParameter.InvalidPrivateNetworkCIDR"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// This API is used to create a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_ACCOUNTUSERNOTAUTHENTICATED = "FailedOperation.AccountUserNotAuthenticated"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_CVMNUMBERNOTMATCH = "FailedOperation.CvmNumberNotMatch"
//  FAILEDOPERATION_CVMVPCIDNOTMATCH = "FailedOperation.CvmVpcidNotMatch"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_DFWGETUSGQUOTA = "FailedOperation.DfwGetUSGQuota"
//  FAILEDOPERATION_NODEEXISTSSECONDARYNETWORKINTERFACE = "FailedOperation.NodeExistsSecondaryNetworkInterface"
//  FAILEDOPERATION_OSNOTSUPPORT = "FailedOperation.OsNotSupport"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXCLSLIMIT = "FailedOperation.QuotaMaxClsLimit"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  FAILEDOPERATION_QUOTAUSGLIMIT = "FailedOperation.QuotaUSGLimit"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  FAILEDOPERATION_WHITELISTUNEXPECTEDERROR = "FailedOperation.WhitelistUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_CIDRINVALI = "InternalError.CidrInvali"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNUMBERNOTMATCH = "InternalError.CvmNumberNotMatch"
//  INTERNALERROR_CVMSTATUS = "InternalError.CvmStatus"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAUSGLIMIT = "InternalError.QuotaUSGLimit"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRMASKSIZEOUTOFRANGE = "InvalidParameter.CIDRMaskSizeOutOfRange"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERCLUSTER = "InvalidParameter.CidrConflictWithOtherCluster"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCGLOBALROUTE = "InvalidParameter.CidrConflictWithVpcGlobalRoute"
//  INVALIDPARAMETER_CIDRINVALI = "InvalidParameter.CidrInvali"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_INVALIDPRIVATENETWORKCIDR = "InvalidParameter.InvalidPrivateNetworkCIDR"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
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

func NewCreateClusterEndpointRequest() (request *CreateClusterEndpointRequest) {
    request = &CreateClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpoint")
    
    
    return
}

func NewCreateClusterEndpointResponse() (response *CreateClusterEndpointResponse) {
    response = &CreateClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterEndpoint
// This API is used to create a cluster access endpoint.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_EMPTYCLUSTERNOTSUPPORT = "InternalError.EmptyClusterNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpoint(request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
    return c.CreateClusterEndpointWithContext(context.Background(), request)
}

// CreateClusterEndpoint
// This API is used to create a cluster access endpoint.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_EMPTYCLUSTERNOTSUPPORT = "InternalError.EmptyClusterNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointWithContext(ctx context.Context, request *CreateClusterEndpointRequest) (response *CreateClusterEndpointResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterEndpointVipRequest() (request *CreateClusterEndpointVipRequest) {
    request = &CreateClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterEndpointVip")
    
    
    return
}

func NewCreateClusterEndpointVipResponse() (response *CreateClusterEndpointVipResponse) {
    response = &CreateClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterEndpointVip
// This API is used to create a public network access port for a managed cluster. Note: This API will be disused soon. Please call `CreateClusterEndpoint` instead.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointVip(request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
    return c.CreateClusterEndpointVipWithContext(context.Background(), request)
}

// CreateClusterEndpointVip
// This API is used to create a public network access port for a managed cluster. Note: This API will be disused soon. Please call `CreateClusterEndpoint` instead.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterEndpointVipWithContext(ctx context.Context, request *CreateClusterEndpointVipRequest) (response *CreateClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewCreateClusterEndpointVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterEndpointVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterInstancesRequest() (request *CreateClusterInstancesRequest) {
    request = &CreateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterInstances")
    
    
    return
}

func NewCreateClusterInstancesResponse() (response *CreateClusterInstancesResponse) {
    response = &CreateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterInstances
// This API is used to create one or more nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterInstances(request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    return c.CreateClusterInstancesWithContext(context.Background(), request)
}

// CreateClusterInstances
// This API is used to create one or more nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTCOMMON = "FailedOperation.AccountCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_NETWORKSCALEERROR = "FailedOperation.NetworkScaleError"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_QUOTAMAXNODLIMIT = "FailedOperation.QuotaMaxNodLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTCOMMON = "InternalError.AccountCommon"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateClusterInstancesWithContext(ctx context.Context, request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewCreateClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterNodePoolRequest() (request *CreateClusterNodePoolRequest) {
    request = &CreateClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterNodePool")
    
    
    return
}

func NewCreateClusterNodePoolResponse() (response *CreateClusterNodePoolResponse) {
    response = &CreateClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterNodePool
// This API is used to create a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_OSNOTSUPPORT = "InvalidParameter.OsNotSupport"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) CreateClusterNodePool(request *CreateClusterNodePoolRequest) (response *CreateClusterNodePoolResponse, err error) {
    return c.CreateClusterNodePoolWithContext(context.Background(), request)
}

// CreateClusterNodePool
// This API is used to create a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_OSNOTSUPPORT = "InvalidParameter.OsNotSupport"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) CreateClusterNodePoolWithContext(ctx context.Context, request *CreateClusterNodePoolRequest) (response *CreateClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewCreateClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRouteTableRequest() (request *CreateClusterRouteTableRequest) {
    request = &CreateClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterRouteTable")
    
    
    return
}

func NewCreateClusterRouteTableResponse() (response *CreateClusterRouteTableResponse) {
    response = &CreateClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterRouteTable
// This API is used to create a cluster route table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_RESOURCEEXISTALREADY = "InternalError.ResourceExistAlready"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
func (c *Client) CreateClusterRouteTable(request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
    return c.CreateClusterRouteTableWithContext(context.Background(), request)
}

// CreateClusterRouteTable
// This API is used to create a cluster route table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERCLUSTER = "InternalError.CidrConflictWithOtherCluster"
//  INTERNALERROR_CIDRCONFLICTWITHOTHERROUTE = "InternalError.CidrConflictWithOtherRoute"
//  INTERNALERROR_CIDRCONFLICTWITHVPCCIDR = "InternalError.CidrConflictWithVpcCidr"
//  INTERNALERROR_CIDRCONFLICTWITHVPCGLOBALROUTE = "InternalError.CidrConflictWithVpcGlobalRoute"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INTERNALERROR_RESOURCEEXISTALREADY = "InternalError.ResourceExistAlready"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CIDRCONFLICTWITHOTHERROUTE = "InvalidParameter.CidrConflictWithOtherRoute"
//  INVALIDPARAMETER_CIDRCONFLICTWITHVPCCIDR = "InvalidParameter.CidrConflictWithVpcCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
func (c *Client) CreateClusterRouteTableWithContext(ctx context.Context, request *CreateClusterRouteTableRequest) (response *CreateClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateClusterRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterRouteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterVirtualNodeRequest() (request *CreateClusterVirtualNodeRequest) {
    request = &CreateClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNode")
    
    
    return
}

func NewCreateClusterVirtualNodeResponse() (response *CreateClusterVirtualNodeResponse) {
    response = &CreateClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterVirtualNode
// This API is used to create the Pay-as-you-go Super Node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNode(request *CreateClusterVirtualNodeRequest) (response *CreateClusterVirtualNodeResponse, err error) {
    return c.CreateClusterVirtualNodeWithContext(context.Background(), request)
}

// CreateClusterVirtualNode
// This API is used to create the Pay-as-you-go Super Node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodeWithContext(ctx context.Context, request *CreateClusterVirtualNodeRequest) (response *CreateClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewCreateClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterVirtualNodePoolRequest() (request *CreateClusterVirtualNodePoolRequest) {
    request = &CreateClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateClusterVirtualNodePool")
    
    
    return
}

func NewCreateClusterVirtualNodePoolResponse() (response *CreateClusterVirtualNodePoolResponse) {
    response = &CreateClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterVirtualNodePool
// This API is used to create the Super Node Pool.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  INVALIDPARAMETER_SUBNETNOTEXIST = "InvalidParameter.SubnetNotExist"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodePool(request *CreateClusterVirtualNodePoolRequest) (response *CreateClusterVirtualNodePoolResponse, err error) {
    return c.CreateClusterVirtualNodePoolWithContext(context.Background(), request)
}

// CreateClusterVirtualNodePool
// This API is used to create the Super Node Pool.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_SUBNETINVALIDERROR = "InvalidParameter.SubnetInvalidError"
//  INVALIDPARAMETER_SUBNETNOTEXIST = "InvalidParameter.SubnetNotExist"
//  RESOURCEINUSE_SUBNETALREADYEXIST = "ResourceInUse.SubnetAlreadyExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTALLVIRTUALKUBELET = "UnsupportedOperation.NotSupportInstallVirtualKubelet"
func (c *Client) CreateClusterVirtualNodePoolWithContext(ctx context.Context, request *CreateClusterVirtualNodePoolRequest) (response *CreateClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewCreateClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateECMInstancesRequest() (request *CreateECMInstancesRequest) {
    request = &CreateECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateECMInstances")
    
    
    return
}

func NewCreateECMInstancesResponse() (response *CreateECMInstancesResponse) {
    response = &CreateECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateECMInstances
// This API is used to create an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateECMInstances(request *CreateECMInstancesRequest) (response *CreateECMInstancesResponse, err error) {
    return c.CreateECMInstancesWithContext(context.Background(), request)
}

// CreateECMInstances
// This API is used to create an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateECMInstancesWithContext(ctx context.Context, request *CreateECMInstancesRequest) (response *CreateECMInstancesResponse, err error) {
    if request == nil {
        request = NewCreateECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeCVMInstancesRequest() (request *CreateEdgeCVMInstancesRequest) {
    request = &CreateEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeCVMInstances")
    
    
    return
}

func NewCreateEdgeCVMInstancesResponse() (response *CreateEdgeCVMInstancesResponse) {
    response = &CreateEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeCVMInstances
// This API is used to create CVM instances in the specified TKE edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeCVMInstances(request *CreateEdgeCVMInstancesRequest) (response *CreateEdgeCVMInstancesResponse, err error) {
    return c.CreateEdgeCVMInstancesWithContext(context.Background(), request)
}

// CreateEdgeCVMInstances
// This API is used to create CVM instances in the specified TKE edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEdgeCVMInstancesWithContext(ctx context.Context, request *CreateEdgeCVMInstancesRequest) (response *CreateEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewCreateEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeLogConfigRequest() (request *CreateEdgeLogConfigRequest) {
    request = &CreateEdgeLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEdgeLogConfig")
    
    
    return
}

func NewCreateEdgeLogConfigResponse() (response *CreateEdgeLogConfigResponse) {
    response = &CreateEdgeLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeLogConfig
// This API is used to create log collection configuration for a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEdgeLogConfig(request *CreateEdgeLogConfigRequest) (response *CreateEdgeLogConfigResponse, err error) {
    return c.CreateEdgeLogConfigWithContext(context.Background(), request)
}

// CreateEdgeLogConfig
// This API is used to create log collection configuration for a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEdgeLogConfigWithContext(ctx context.Context, request *CreateEdgeLogConfigRequest) (response *CreateEdgeLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateEdgeLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEksLogConfigRequest() (request *CreateEksLogConfigRequest) {
    request = &CreateEksLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateEksLogConfig")
    
    
    return
}

func NewCreateEksLogConfigResponse() (response *CreateEksLogConfigResponse) {
    response = &CreateEksLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEksLogConfig
// This API is used to create Log Collection Configuration for Elastic Cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSTOPIC = "FailedOperation.GetClsTopic"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEksLogConfig(request *CreateEksLogConfigRequest) (response *CreateEksLogConfigResponse, err error) {
    return c.CreateEksLogConfigWithContext(context.Background(), request)
}

// CreateEksLogConfig
// This API is used to create Log Collection Configuration for Elastic Cluster.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_CREATECLSTOPIC = "FailedOperation.CreateClsTopic"
//  FAILEDOPERATION_GETCLSTOPIC = "FailedOperation.GetClsTopic"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESCREATEOPERATIONERROR = "FailedOperation.KubernetesCreateOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEksLogConfigWithContext(ctx context.Context, request *CreateEksLogConfigRequest) (response *CreateEksLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateEksLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEksLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEksLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertRuleRequest() (request *CreatePrometheusAlertRuleRequest) {
    request = &CreatePrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreatePrometheusAlertRule")
    
    
    return
}

func NewCreatePrometheusAlertRuleResponse() (response *CreatePrometheusAlertRuleResponse) {
    response = &CreatePrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAlertRule
// This API is used to create an alarm rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
    return c.CreatePrometheusAlertRuleWithContext(context.Background(), request)
}

// CreatePrometheusAlertRule
// This API is used to create an alarm rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertRuleWithContext(ctx context.Context, request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTKEEdgeClusterRequest() (request *CreateTKEEdgeClusterRequest) {
    request = &CreateTKEEdgeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateTKEEdgeCluster")
    
    
    return
}

func NewCreateTKEEdgeClusterResponse() (response *CreateTKEEdgeClusterResponse) {
    response = &CreateTKEEdgeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTKEEdgeCluster
// This API is used to create a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTKEEdgeCluster(request *CreateTKEEdgeClusterRequest) (response *CreateTKEEdgeClusterResponse, err error) {
    return c.CreateTKEEdgeClusterWithContext(context.Background(), request)
}

// CreateTKEEdgeCluster
// This API is used to create a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTKEEdgeClusterWithContext(ctx context.Context, request *CreateTKEEdgeClusterRequest) (response *CreateTKEEdgeClusterResponse, err error) {
    if request == nil {
        request = NewCreateTKEEdgeClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTKEEdgeCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTKEEdgeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddonRequest() (request *DeleteAddonRequest) {
    request = &DeleteAddonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteAddon")
    
    
    return
}

func NewDeleteAddonResponse() (response *DeleteAddonResponse) {
    response = &DeleteAddonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAddon
// This API is used to delete an add-on.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAddon(request *DeleteAddonRequest) (response *DeleteAddonResponse, err error) {
    return c.DeleteAddonWithContext(context.Background(), request)
}

// DeleteAddon
// This API is used to delete an add-on.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAddonWithContext(ctx context.Context, request *DeleteAddonRequest) (response *DeleteAddonResponse, err error) {
    if request == nil {
        request = NewDeleteAddonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddon require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAddonResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupStorageLocationRequest() (request *DeleteBackupStorageLocationRequest) {
    request = &DeleteBackupStorageLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteBackupStorageLocation")
    
    
    return
}

func NewDeleteBackupStorageLocationResponse() (response *DeleteBackupStorageLocationResponse) {
    response = &DeleteBackupStorageLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackupStorageLocation
// This API is used to delete a backup repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteBackupStorageLocation(request *DeleteBackupStorageLocationRequest) (response *DeleteBackupStorageLocationResponse, err error) {
    return c.DeleteBackupStorageLocationWithContext(context.Background(), request)
}

// DeleteBackupStorageLocation
// This API is used to delete a backup repository.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteBackupStorageLocationWithContext(ctx context.Context, request *DeleteBackupStorageLocationRequest) (response *DeleteBackupStorageLocationResponse, err error) {
    if request == nil {
        request = NewDeleteBackupStorageLocationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupStorageLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupStorageLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCluster
// This API is used to delete a cluster. (Cloud API v3).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_CLUSTERFORBIDDENTODELETE = "FailedOperation.ClusterForbiddenToDelete"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_TAGCOMMON = "FailedOperation.TagCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_CLUSTERINDELETIONPROTECTION = "OperationDenied.ClusterInDeletionProtection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// This API is used to delete a cluster. (Cloud API v3).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_CLUSTERFORBIDDENTODELETE = "FailedOperation.ClusterForbiddenToDelete"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_TAGCOMMON = "FailedOperation.TagCommon"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_CLUSTERINDELETIONPROTECTION = "OperationDenied.ClusterInDeletionProtection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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

func NewDeleteClusterAsGroupsRequest() (request *DeleteClusterAsGroupsRequest) {
    request = &DeleteClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterAsGroups")
    
    
    return
}

func NewDeleteClusterAsGroupsResponse() (response *DeleteClusterAsGroupsResponse) {
    response = &DeleteClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterAsGroups
// Delete a cluster scaling group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteClusterAsGroups(request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
    return c.DeleteClusterAsGroupsWithContext(context.Background(), request)
}

// DeleteClusterAsGroups
// Delete a cluster scaling group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_QUOTAMAXRTLIMIT = "InternalError.QuotaMaxRtLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteClusterAsGroupsWithContext(ctx context.Context, request *DeleteClusterAsGroupsRequest) (response *DeleteClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteClusterAsGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterAsGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointRequest() (request *DeleteClusterEndpointRequest) {
    request = &DeleteClusterEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpoint")
    
    
    return
}

func NewDeleteClusterEndpointResponse() (response *DeleteClusterEndpointResponse) {
    response = &DeleteClusterEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterEndpoint
// This API is used to delete a cluster access endpoint.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpoint(request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
    return c.DeleteClusterEndpointWithContext(context.Background(), request)
}

// DeleteClusterEndpoint
// This API is used to delete a cluster access endpoint.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointWithContext(ctx context.Context, request *DeleteClusterEndpointRequest) (response *DeleteClusterEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterEndpointVipRequest() (request *DeleteClusterEndpointVipRequest) {
    request = &DeleteClusterEndpointVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterEndpointVip")
    
    
    return
}

func NewDeleteClusterEndpointVipResponse() (response *DeleteClusterEndpointVipResponse) {
    response = &DeleteClusterEndpointVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterEndpointVip
// Delete the external network access port of the managed cluster (the old way, only the external network port of the managed cluster is supported)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointVip(request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
    return c.DeleteClusterEndpointVipWithContext(context.Background(), request)
}

// DeleteClusterEndpointVip
// Delete the external network access port of the managed cluster (the old way, only the external network port of the managed cluster is supported)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterEndpointVipWithContext(ctx context.Context, request *DeleteClusterEndpointVipRequest) (response *DeleteClusterEndpointVipResponse, err error) {
    if request == nil {
        request = NewDeleteClusterEndpointVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterEndpointVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterEndpointVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
    request = &DeleteClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterInstances")
    
    
    return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
    response = &DeleteClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterInstances
// This API is used to delete one or more nodes from a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMDELETIONPROTECTION = "FailedOperation.CvmDeletionProtection"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    return c.DeleteClusterInstancesWithContext(context.Background(), request)
}

// DeleteClusterInstances
// This API is used to delete one or more nodes from a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMDELETIONPROTECTION = "FailedOperation.CvmDeletionProtection"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteClusterInstancesWithContext(ctx context.Context, request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterNodePoolRequest() (request *DeleteClusterNodePoolRequest) {
    request = &DeleteClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterNodePool")
    
    
    return
}

func NewDeleteClusterNodePoolResponse() (response *DeleteClusterNodePoolResponse) {
    response = &DeleteClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterNodePool
// This API is used to delete a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_CVMDELETIONPROTECTION = "FailedOperation.CvmDeletionProtection"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteClusterNodePool(request *DeleteClusterNodePoolRequest) (response *DeleteClusterNodePoolResponse, err error) {
    return c.DeleteClusterNodePoolWithContext(context.Background(), request)
}

// DeleteClusterNodePool
// This API is used to delete a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_CVMDELETIONPROTECTION = "FailedOperation.CvmDeletionProtection"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteClusterNodePoolWithContext(ctx context.Context, request *DeleteClusterNodePoolRequest) (response *DeleteClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewDeleteClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteRequest() (request *DeleteClusterRouteRequest) {
    request = &DeleteClusterRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRoute")
    
    
    return
}

func NewDeleteClusterRouteResponse() (response *DeleteClusterRouteResponse) {
    response = &DeleteClusterRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterRoute
// This API is used to delete a cluster route.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteClusterRoute(request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
    return c.DeleteClusterRouteWithContext(context.Background(), request)
}

// DeleteClusterRoute
// This API is used to delete a cluster route.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_VPCCOMMON = "FailedOperation.VpcCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_ROUTETABLENOTFOUND = "InternalError.RouteTableNotFound"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteClusterRouteWithContext(ctx context.Context, request *DeleteClusterRouteRequest) (response *DeleteClusterRouteResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRouteTableRequest() (request *DeleteClusterRouteTableRequest) {
    request = &DeleteClusterRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterRouteTable")
    
    
    return
}

func NewDeleteClusterRouteTableResponse() (response *DeleteClusterRouteTableResponse) {
    response = &DeleteClusterRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterRouteTable
// This API is used to delete cluster a route table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ROUTETABLENOTEMPTY = "InternalError.RouteTableNotEmpty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
func (c *Client) DeleteClusterRouteTable(request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
    return c.DeleteClusterRouteTableWithContext(context.Background(), request)
}

// DeleteClusterRouteTable
// This API is used to delete cluster a route table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ROUTETABLENOTEMPTY = "InternalError.RouteTableNotEmpty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
func (c *Client) DeleteClusterRouteTableWithContext(ctx context.Context, request *DeleteClusterRouteTableRequest) (response *DeleteClusterRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterRouteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterVirtualNodeRequest() (request *DeleteClusterVirtualNodeRequest) {
    request = &DeleteClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNode")
    
    
    return
}

func NewDeleteClusterVirtualNodeResponse() (response *DeleteClusterVirtualNodeResponse) {
    response = &DeleteClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterVirtualNode
// This API is used to delete the super node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNode(request *DeleteClusterVirtualNodeRequest) (response *DeleteClusterVirtualNodeResponse, err error) {
    return c.DeleteClusterVirtualNodeWithContext(context.Background(), request)
}

// DeleteClusterVirtualNode
// This API is used to delete the super node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodeWithContext(ctx context.Context, request *DeleteClusterVirtualNodeRequest) (response *DeleteClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDeleteClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterVirtualNodePoolRequest() (request *DeleteClusterVirtualNodePoolRequest) {
    request = &DeleteClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterVirtualNodePool")
    
    
    return
}

func NewDeleteClusterVirtualNodePoolResponse() (response *DeleteClusterVirtualNodePoolResponse) {
    response = &DeleteClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterVirtualNodePool
// This API is used to delete the Super Node Pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodePool(request *DeleteClusterVirtualNodePoolRequest) (response *DeleteClusterVirtualNodePoolResponse, err error) {
    return c.DeleteClusterVirtualNodePoolWithContext(context.Background(), request)
}

// DeleteClusterVirtualNodePool
// This API is used to delete the Super Node Pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCEINUSE_EXISTRUNNINGPOD = "ResourceInUse.ExistRunningPod"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DeleteClusterVirtualNodePoolWithContext(ctx context.Context, request *DeleteClusterVirtualNodePoolRequest) (response *DeleteClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewDeleteClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteECMInstancesRequest() (request *DeleteECMInstancesRequest) {
    request = &DeleteECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteECMInstances")
    
    
    return
}

func NewDeleteECMInstancesResponse() (response *DeleteECMInstancesResponse) {
    response = &DeleteECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteECMInstances
// This API is used to delete one or more ECM instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteECMInstances(request *DeleteECMInstancesRequest) (response *DeleteECMInstancesResponse, err error) {
    return c.DeleteECMInstancesWithContext(context.Background(), request)
}

// DeleteECMInstances
// This API is used to delete one or more ECM instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteECMInstancesWithContext(ctx context.Context, request *DeleteECMInstancesRequest) (response *DeleteECMInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeCVMInstancesRequest() (request *DeleteEdgeCVMInstancesRequest) {
    request = &DeleteEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeCVMInstances")
    
    
    return
}

func NewDeleteEdgeCVMInstancesResponse() (response *DeleteEdgeCVMInstancesResponse) {
    response = &DeleteEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEdgeCVMInstances
// This API is used to delete one or more edge CVM instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeCVMInstances(request *DeleteEdgeCVMInstancesRequest) (response *DeleteEdgeCVMInstancesResponse, err error) {
    return c.DeleteEdgeCVMInstancesWithContext(context.Background(), request)
}

// DeleteEdgeCVMInstances
// This API is used to delete one or more edge CVM instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeCVMInstancesWithContext(ctx context.Context, request *DeleteEdgeCVMInstancesRequest) (response *DeleteEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeClusterInstancesRequest() (request *DeleteEdgeClusterInstancesRequest) {
    request = &DeleteEdgeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteEdgeClusterInstances")
    
    
    return
}

func NewDeleteEdgeClusterInstancesResponse() (response *DeleteEdgeClusterInstancesResponse) {
    response = &DeleteEdgeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEdgeClusterInstances
// This API is used to delete one or more edge compute instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeClusterInstances(request *DeleteEdgeClusterInstancesRequest) (response *DeleteEdgeClusterInstancesResponse, err error) {
    return c.DeleteEdgeClusterInstancesWithContext(context.Background(), request)
}

// DeleteEdgeClusterInstances
// This API is used to delete one or more edge compute instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEdgeClusterInstancesWithContext(ctx context.Context, request *DeleteEdgeClusterInstancesRequest) (response *DeleteEdgeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogConfigsRequest() (request *DeleteLogConfigsRequest) {
    request = &DeleteLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteLogConfigs")
    
    
    return
}

func NewDeleteLogConfigsResponse() (response *DeleteLogConfigsResponse) {
    response = &DeleteLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogConfigs
// This API is used to delete collection rules within the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_K8SCLIENTBUILDERROR = "FailedOperation.K8sClientBuildError"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_KUBERNETESRESOURCENOTFOUND = "FailedOperation.KubernetesResourceNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
func (c *Client) DeleteLogConfigs(request *DeleteLogConfigsRequest) (response *DeleteLogConfigsResponse, err error) {
    return c.DeleteLogConfigsWithContext(context.Background(), request)
}

// DeleteLogConfigs
// This API is used to delete collection rules within the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_K8SCLIENTBUILDERROR = "FailedOperation.K8sClientBuildError"
//  FAILEDOPERATION_KUBERNETESDELETEOPERATIONERROR = "FailedOperation.KubernetesDeleteOperationError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_KUBERNETESRESOURCENOTFOUND = "FailedOperation.KubernetesResourceNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
func (c *Client) DeleteLogConfigsWithContext(ctx context.Context, request *DeleteLogConfigsRequest) (response *DeleteLogConfigsResponse, err error) {
    if request == nil {
        request = NewDeleteLogConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertRuleRequest() (request *DeletePrometheusAlertRuleRequest) {
    request = &DeletePrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeletePrometheusAlertRule")
    
    
    return
}

func NewDeletePrometheusAlertRuleResponse() (response *DeletePrometheusAlertRuleResponse) {
    response = &DeletePrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusAlertRule
// This API is used to delete an alarm rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
    return c.DeletePrometheusAlertRuleWithContext(context.Background(), request)
}

// DeletePrometheusAlertRule
// This API is used to delete an alarm rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertRuleWithContext(ctx context.Context, request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTKEEdgeClusterRequest() (request *DeleteTKEEdgeClusterRequest) {
    request = &DeleteTKEEdgeClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteTKEEdgeCluster")
    
    
    return
}

func NewDeleteTKEEdgeClusterResponse() (response *DeleteTKEEdgeClusterResponse) {
    response = &DeleteTKEEdgeClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTKEEdgeCluster
// This API is used to delete a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTKEEdgeCluster(request *DeleteTKEEdgeClusterRequest) (response *DeleteTKEEdgeClusterResponse, err error) {
    return c.DeleteTKEEdgeClusterWithContext(context.Background(), request)
}

// DeleteTKEEdgeCluster
// This API is used to delete a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTKEEdgeClusterWithContext(ctx context.Context, request *DeleteTKEEdgeClusterRequest) (response *DeleteTKEEdgeClusterResponse, err error) {
    if request == nil {
        request = NewDeleteTKEEdgeClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTKEEdgeCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTKEEdgeClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddonRequest() (request *DescribeAddonRequest) {
    request = &DescribeAddonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAddon")
    
    
    return
}

func NewDescribeAddonResponse() (response *DescribeAddonResponse) {
    response = &DescribeAddonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddon
// This API is used to query the list of add-ons.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddon(request *DescribeAddonRequest) (response *DescribeAddonResponse, err error) {
    return c.DescribeAddonWithContext(context.Background(), request)
}

// DescribeAddon
// This API is used to query the list of add-ons.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddonWithContext(ctx context.Context, request *DescribeAddonRequest) (response *DescribeAddonResponse, err error) {
    if request == nil {
        request = NewDescribeAddonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddon require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddonResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddonValuesRequest() (request *DescribeAddonValuesRequest) {
    request = &DescribeAddonValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAddonValues")
    
    
    return
}

func NewDescribeAddonValuesResponse() (response *DescribeAddonValuesResponse) {
    response = &DescribeAddonValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddonValues
// This API is used to query parameters of an add-on.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeAddonValues(request *DescribeAddonValuesRequest) (response *DescribeAddonValuesResponse, err error) {
    return c.DescribeAddonValuesWithContext(context.Background(), request)
}

// DescribeAddonValues
// This API is used to query parameters of an add-on.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeAddonValuesWithContext(ctx context.Context, request *DescribeAddonValuesRequest) (response *DescribeAddonValuesResponse, err error) {
    if request == nil {
        request = NewDescribeAddonValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddonValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddonValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableClusterVersionRequest() (request *DescribeAvailableClusterVersionRequest) {
    request = &DescribeAvailableClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableClusterVersion")
    
    
    return
}

func NewDescribeAvailableClusterVersionResponse() (response *DescribeAvailableClusterVersionResponse) {
    response = &DescribeAvailableClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableClusterVersion
// This API is used to obtain all versions that the cluster can upgrade to.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
func (c *Client) DescribeAvailableClusterVersion(request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
    return c.DescribeAvailableClusterVersionWithContext(context.Background(), request)
}

// DescribeAvailableClusterVersion
// This API is used to obtain all versions that the cluster can upgrade to.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
func (c *Client) DescribeAvailableClusterVersionWithContext(ctx context.Context, request *DescribeAvailableClusterVersionRequest) (response *DescribeAvailableClusterVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableTKEEdgeVersionRequest() (request *DescribeAvailableTKEEdgeVersionRequest) {
    request = &DescribeAvailableTKEEdgeVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeAvailableTKEEdgeVersion")
    
    
    return
}

func NewDescribeAvailableTKEEdgeVersionResponse() (response *DescribeAvailableTKEEdgeVersionResponse) {
    response = &DescribeAvailableTKEEdgeVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableTKEEdgeVersion
// This API is used to check the edge component versions and K8s versions supported by TKE Edge.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAvailableTKEEdgeVersion(request *DescribeAvailableTKEEdgeVersionRequest) (response *DescribeAvailableTKEEdgeVersionResponse, err error) {
    return c.DescribeAvailableTKEEdgeVersionWithContext(context.Background(), request)
}

// DescribeAvailableTKEEdgeVersion
// This API is used to check the edge component versions and K8s versions supported by TKE Edge.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAvailableTKEEdgeVersionWithContext(ctx context.Context, request *DescribeAvailableTKEEdgeVersionRequest) (response *DescribeAvailableTKEEdgeVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableTKEEdgeVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableTKEEdgeVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableTKEEdgeVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupStorageLocationsRequest() (request *DescribeBackupStorageLocationsRequest) {
    request = &DescribeBackupStorageLocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeBackupStorageLocations")
    
    
    return
}

func NewDescribeBackupStorageLocationsResponse() (response *DescribeBackupStorageLocationsResponse) {
    response = &DescribeBackupStorageLocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupStorageLocations
// This API is used to query backup repositories.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBackupStorageLocations(request *DescribeBackupStorageLocationsRequest) (response *DescribeBackupStorageLocationsResponse, err error) {
    return c.DescribeBackupStorageLocationsWithContext(context.Background(), request)
}

// DescribeBackupStorageLocations
// This API is used to query backup repositories.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBackupStorageLocationsWithContext(ctx context.Context, request *DescribeBackupStorageLocationsRequest) (response *DescribeBackupStorageLocationsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupStorageLocationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupStorageLocations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupStorageLocationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchModifyTagsStatusRequest() (request *DescribeBatchModifyTagsStatusRequest) {
    request = &DescribeBatchModifyTagsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeBatchModifyTagsStatus")
    
    
    return
}

func NewDescribeBatchModifyTagsStatusResponse() (response *DescribeBatchModifyTagsStatusResponse) {
    response = &DescribeBatchModifyTagsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchModifyTagsStatus
// This API is used to query batch modification Tag status.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
func (c *Client) DescribeBatchModifyTagsStatus(request *DescribeBatchModifyTagsStatusRequest) (response *DescribeBatchModifyTagsStatusResponse, err error) {
    return c.DescribeBatchModifyTagsStatusWithContext(context.Background(), request)
}

// DescribeBatchModifyTagsStatus
// This API is used to query batch modification Tag status.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
func (c *Client) DescribeBatchModifyTagsStatusWithContext(ctx context.Context, request *DescribeBatchModifyTagsStatusRequest) (response *DescribeBatchModifyTagsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBatchModifyTagsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchModifyTagsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchModifyTagsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupOptionRequest() (request *DescribeClusterAsGroupOptionRequest) {
    request = &DescribeClusterAsGroupOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroupOption")
    
    
    return
}

func NewDescribeClusterAsGroupOptionResponse() (response *DescribeClusterAsGroupOptionResponse) {
    response = &DescribeClusterAsGroupOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAsGroupOption
// Cluster auto scaling configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterAsGroupOption(request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
    return c.DescribeClusterAsGroupOptionWithContext(context.Background(), request)
}

// DescribeClusterAsGroupOption
// Cluster auto scaling configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterAsGroupOptionWithContext(ctx context.Context, request *DescribeClusterAsGroupOptionRequest) (response *DescribeClusterAsGroupOptionResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupOptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAsGroupOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAsGroupOptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAsGroupsRequest() (request *DescribeClusterAsGroupsRequest) {
    request = &DescribeClusterAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAsGroups")
    
    
    return
}

func NewDescribeClusterAsGroupsResponse() (response *DescribeClusterAsGroupsResponse) {
    response = &DescribeClusterAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAsGroups
// Cluster-associated scaling group list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PODNOTFOUND = "InternalError.PodNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
func (c *Client) DescribeClusterAsGroups(request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
    return c.DescribeClusterAsGroupsWithContext(context.Background(), request)
}

// DescribeClusterAsGroups
// Cluster-associated scaling group list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PODNOTFOUND = "InternalError.PodNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCPEERNOTFOUND = "InternalError.VpcPeerNotFound"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
func (c *Client) DescribeClusterAsGroupsWithContext(ctx context.Context, request *DescribeClusterAsGroupsRequest) (response *DescribeClusterAsGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAsGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAsGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAuthenticationOptionsRequest() (request *DescribeClusterAuthenticationOptionsRequest) {
    request = &DescribeClusterAuthenticationOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterAuthenticationOptions")
    
    
    return
}

func NewDescribeClusterAuthenticationOptionsResponse() (response *DescribeClusterAuthenticationOptionsResponse) {
    response = &DescribeClusterAuthenticationOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAuthenticationOptions
// This API is used to query cluster authentication configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterAuthenticationOptions(request *DescribeClusterAuthenticationOptionsRequest) (response *DescribeClusterAuthenticationOptionsResponse, err error) {
    return c.DescribeClusterAuthenticationOptionsWithContext(context.Background(), request)
}

// DescribeClusterAuthenticationOptions
// This API is used to query cluster authentication configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterAuthenticationOptionsWithContext(ctx context.Context, request *DescribeClusterAuthenticationOptionsRequest) (response *DescribeClusterAuthenticationOptionsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAuthenticationOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAuthenticationOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAuthenticationOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterCommonNamesRequest() (request *DescribeClusterCommonNamesRequest) {
    request = &DescribeClusterCommonNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterCommonNames")
    
    
    return
}

func NewDescribeClusterCommonNamesResponse() (response *DescribeClusterCommonNamesResponse) {
    response = &DescribeClusterCommonNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterCommonNames
// This API is used to obtain the CommonName from the kube-apiserver client certificate that corresponding to the sub-account in RBAC authorization mode. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeClusterCommonNames(request *DescribeClusterCommonNamesRequest) (response *DescribeClusterCommonNamesResponse, err error) {
    return c.DescribeClusterCommonNamesWithContext(context.Background(), request)
}

// DescribeClusterCommonNames
// This API is used to obtain the CommonName from the kube-apiserver client certificate that corresponding to the sub-account in RBAC authorization mode. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeClusterCommonNamesWithContext(ctx context.Context, request *DescribeClusterCommonNamesRequest) (response *DescribeClusterCommonNamesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterCommonNamesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterCommonNames require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterCommonNamesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointStatusRequest() (request *DescribeClusterEndpointStatusRequest) {
    request = &DescribeClusterEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointStatus")
    
    
    return
}

func NewDescribeClusterEndpointStatusResponse() (response *DescribeClusterEndpointStatusResponse) {
    response = &DescribeClusterEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterEndpointStatus
// Query cluster access port status (intranet / extranet access is enabled for independent clusters, and intranet access is supported for managed clusters)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_KUBERNETESINTERNAL = "InternalError.KubernetesInternal"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointStatus(request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
    return c.DescribeClusterEndpointStatusWithContext(context.Background(), request)
}

// DescribeClusterEndpointStatus
// Query cluster access port status (intranet / extranet access is enabled for independent clusters, and intranet access is supported for managed clusters)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBECLIENTCONNECTION = "FailedOperation.KubeClientConnection"
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_KUBERNETESINTERNAL = "InternalError.KubernetesInternal"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointStatusWithContext(ctx context.Context, request *DescribeClusterEndpointStatusRequest) (response *DescribeClusterEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpointStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointVipStatusRequest() (request *DescribeClusterEndpointVipStatusRequest) {
    request = &DescribeClusterEndpointVipStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpointVipStatus")
    
    
    return
}

func NewDescribeClusterEndpointVipStatusResponse() (response *DescribeClusterEndpointVipStatusResponse) {
    response = &DescribeClusterEndpointVipStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterEndpointVipStatus
// Query cluster open port process status (only supports external ports of the managed cluster)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointVipStatus(request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
    return c.DescribeClusterEndpointVipStatusWithContext(context.Background(), request)
}

// DescribeClusterEndpointVipStatus
// Query cluster open port process status (only supports external ports of the managed cluster)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_DFWGETUSGCOUNT = "InternalError.DfwGetUSGCount"
//  INTERNALERROR_DFWGETUSGQUOTA = "InternalError.DfwGetUSGQuota"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterEndpointVipStatusWithContext(ctx context.Context, request *DescribeClusterEndpointVipStatusRequest) (response *DescribeClusterEndpointVipStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointVipStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpointVipStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointVipStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointsRequest() (request *DescribeClusterEndpointsRequest) {
    request = &DescribeClusterEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterEndpoints")
    
    
    return
}

func NewDescribeClusterEndpointsResponse() (response *DescribeClusterEndpointsResponse) {
    response = &DescribeClusterEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterEndpoints
// This API is used to query cluster access addresses, including private network address, public network address, public network domain name, and security policy for public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterEndpoints(request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    return c.DescribeClusterEndpointsWithContext(context.Background(), request)
}

// DescribeClusterEndpoints
// This API is used to query cluster access addresses, including private network address, public network address, public network domain name, and security policy for public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterEndpointsWithContext(ctx context.Context, request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterExtraArgsRequest() (request *DescribeClusterExtraArgsRequest) {
    request = &DescribeClusterExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterExtraArgs")
    
    
    return
}

func NewDescribeClusterExtraArgsResponse() (response *DescribeClusterExtraArgsResponse) {
    response = &DescribeClusterExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterExtraArgs
// This API is used to query custom parameters of a cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterExtraArgs(request *DescribeClusterExtraArgsRequest) (response *DescribeClusterExtraArgsResponse, err error) {
    return c.DescribeClusterExtraArgsWithContext(context.Background(), request)
}

// DescribeClusterExtraArgs
// This API is used to query custom parameters of a cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterExtraArgsWithContext(ctx context.Context, request *DescribeClusterExtraArgsRequest) (response *DescribeClusterExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterExtraArgsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterExtraArgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstances
// This API is used to query information of node instances in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    return c.DescribeClusterInstancesWithContext(context.Background(), request)
}

// DescribeClusterInstances
// This API is used to query information of node instances in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterKubeconfigRequest() (request *DescribeClusterKubeconfigRequest) {
    request = &DescribeClusterKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterKubeconfig")
    
    
    return
}

func NewDescribeClusterKubeconfigResponse() (response *DescribeClusterKubeconfigResponse) {
    response = &DescribeClusterKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterKubeconfig
// This API is used to obtain the cluster kubeconfig file. Different sub-accounts have their own kubeconfig files. The kubeconfig file contains the kube-apiserver client certificate of the corresponding sub-account. By default, the client certificate is created when this API is called for the first time, and the certificate is valid for 20 years with no permissions granted. For the cluster owner or primary account, the cluster-admin permission is granted by default.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterKubeconfig(request *DescribeClusterKubeconfigRequest) (response *DescribeClusterKubeconfigResponse, err error) {
    return c.DescribeClusterKubeconfigWithContext(context.Background(), request)
}

// DescribeClusterKubeconfig
// This API is used to obtain the cluster kubeconfig file. Different sub-accounts have their own kubeconfig files. The kubeconfig file contains the kube-apiserver client certificate of the corresponding sub-account. By default, the client certificate is created when this API is called for the first time, and the certificate is valid for 20 years with no permissions granted. For the cluster owner or primary account, the cluster-admin permission is granted by default.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBECLIENTCONNECTION = "InternalError.KubeClientConnection"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESCREATEOPERATIONERROR = "InternalError.KubernetesCreateOperationError"
//  INTERNALERROR_KUBERNETESDELETEOPERATIONERROR = "InternalError.KubernetesDeleteOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_WHITELISTUNEXPECTEDERROR = "InternalError.WhitelistUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERNETESRESOURCENOTFOUND = "ResourceNotFound.KubernetesResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterKubeconfigWithContext(ctx context.Context, request *DescribeClusterKubeconfigRequest) (response *DescribeClusterKubeconfigResponse, err error) {
    if request == nil {
        request = NewDescribeClusterKubeconfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterLevelAttributeRequest() (request *DescribeClusterLevelAttributeRequest) {
    request = &DescribeClusterLevelAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelAttribute")
    
    
    return
}

func NewDescribeClusterLevelAttributeResponse() (response *DescribeClusterLevelAttributeResponse) {
    response = &DescribeClusterLevelAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterLevelAttribute
// This API is used to obtain the cluster model.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelAttribute(request *DescribeClusterLevelAttributeRequest) (response *DescribeClusterLevelAttributeResponse, err error) {
    return c.DescribeClusterLevelAttributeWithContext(context.Background(), request)
}

// DescribeClusterLevelAttribute
// This API is used to obtain the cluster model.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelAttributeWithContext(ctx context.Context, request *DescribeClusterLevelAttributeRequest) (response *DescribeClusterLevelAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterLevelAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterLevelAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterLevelAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterLevelChangeRecordsRequest() (request *DescribeClusterLevelChangeRecordsRequest) {
    request = &DescribeClusterLevelChangeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterLevelChangeRecords")
    
    
    return
}

func NewDescribeClusterLevelChangeRecordsResponse() (response *DescribeClusterLevelChangeRecordsResponse) {
    response = &DescribeClusterLevelChangeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterLevelChangeRecords
// This API is used to query the cluster model adjustment history.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelChangeRecords(request *DescribeClusterLevelChangeRecordsRequest) (response *DescribeClusterLevelChangeRecordsResponse, err error) {
    return c.DescribeClusterLevelChangeRecordsWithContext(context.Background(), request)
}

// DescribeClusterLevelChangeRecords
// This API is used to query the cluster model adjustment history.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeClusterLevelChangeRecordsWithContext(ctx context.Context, request *DescribeClusterLevelChangeRecordsRequest) (response *DescribeClusterLevelChangeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterLevelChangeRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterLevelChangeRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterLevelChangeRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodePoolDetailRequest() (request *DescribeClusterNodePoolDetailRequest) {
    request = &DescribeClusterNodePoolDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePoolDetail")
    
    
    return
}

func NewDescribeClusterNodePoolDetailResponse() (response *DescribeClusterNodePoolDetailResponse) {
    response = &DescribeClusterNodePoolDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterNodePoolDetail
// This API is used to query detailed information of a node pool.
//
// error code that may be returned:
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolDetail(request *DescribeClusterNodePoolDetailRequest) (response *DescribeClusterNodePoolDetailResponse, err error) {
    return c.DescribeClusterNodePoolDetailWithContext(context.Background(), request)
}

// DescribeClusterNodePoolDetail
// This API is used to query detailed information of a node pool.
//
// error code that may be returned:
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolDetailWithContext(ctx context.Context, request *DescribeClusterNodePoolDetailRequest) (response *DescribeClusterNodePoolDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodePoolDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodePoolDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodePoolDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNodePoolsRequest() (request *DescribeClusterNodePoolsRequest) {
    request = &DescribeClusterNodePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterNodePools")
    
    
    return
}

func NewDescribeClusterNodePoolsResponse() (response *DescribeClusterNodePoolsResponse) {
    response = &DescribeClusterNodePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterNodePools
// This API is used to query the node pool list
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePools(request *DescribeClusterNodePoolsRequest) (response *DescribeClusterNodePoolsResponse, err error) {
    return c.DescribeClusterNodePoolsWithContext(context.Background(), request)
}

// DescribeClusterNodePools
// This API is used to query the node pool list
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterNodePoolsWithContext(ctx context.Context, request *DescribeClusterNodePoolsRequest) (response *DescribeClusterNodePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNodePoolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNodePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRouteTablesRequest() (request *DescribeClusterRouteTablesRequest) {
    request = &DescribeClusterRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRouteTables")
    
    
    return
}

func NewDescribeClusterRouteTablesResponse() (response *DescribeClusterRouteTablesResponse) {
    response = &DescribeClusterRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterRouteTables
// This API is used to query one or more cluster route tables.
//
// error code that may be returned:
//  INTERNALERROR_DB = "InternalError.Db"
func (c *Client) DescribeClusterRouteTables(request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
    return c.DescribeClusterRouteTablesWithContext(context.Background(), request)
}

// DescribeClusterRouteTables
// This API is used to query one or more cluster route tables.
//
// error code that may be returned:
//  INTERNALERROR_DB = "InternalError.Db"
func (c *Client) DescribeClusterRouteTablesWithContext(ctx context.Context, request *DescribeClusterRouteTablesRequest) (response *DescribeClusterRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRouteTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRouteTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRoutesRequest() (request *DescribeClusterRoutesRequest) {
    request = &DescribeClusterRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterRoutes")
    
    
    return
}

func NewDescribeClusterRoutesResponse() (response *DescribeClusterRoutesResponse) {
    response = &DescribeClusterRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterRoutes
// This API is used to query cluster routes.
//
// error code that may be returned:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeClusterRoutes(request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
    return c.DescribeClusterRoutesWithContext(context.Background(), request)
}

// DescribeClusterRoutes
// This API is used to query cluster routes.
//
// error code that may be returned:
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeClusterRoutesWithContext(ctx context.Context, request *DescribeClusterRoutesRequest) (response *DescribeClusterRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSecurityRequest() (request *DescribeClusterSecurityRequest) {
    request = &DescribeClusterSecurityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterSecurity")
    
    
    return
}

func NewDescribeClusterSecurityResponse() (response *DescribeClusterSecurityResponse) {
    response = &DescribeClusterSecurityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterSecurity
// This API is used to query the key information of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_LBCOMMON = "InternalError.LbCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterSecurity(request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
    return c.DescribeClusterSecurityWithContext(context.Background(), request)
}

// DescribeClusterSecurity
// This API is used to query the key information of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_COMPONENTCLIENTHTTP = "InternalError.ComponentClientHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_KUBECOMMON = "InternalError.KubeCommon"
//  INTERNALERROR_LBCOMMON = "InternalError.LbCommon"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CIDRINVALID = "InvalidParameter.CidrInvalid"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCENOTFOUND_KUBERESOURCENOTFOUND = "ResourceNotFound.KubeResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERINABNORMALSTAT = "ResourceUnavailable.ClusterInAbnormalStat"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterSecurityWithContext(ctx context.Context, request *DescribeClusterSecurityRequest) (response *DescribeClusterSecurityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSecurityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterSecurity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterSecurityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterStatusRequest() (request *DescribeClusterStatusRequest) {
    request = &DescribeClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterStatus")
    
    
    return
}

func NewDescribeClusterStatusResponse() (response *DescribeClusterStatusResponse) {
    response = &DescribeClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterStatus
// This API is used to query the information of clusters under the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
    return c.DescribeClusterStatusWithContext(context.Background(), request)
}

// DescribeClusterStatus
// This API is used to query the information of clusters under the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMNOAUTH = "FailedOperation.CamNoAuth"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterStatusWithContext(ctx context.Context, request *DescribeClusterStatusRequest) (response *DescribeClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClusterStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterVirtualNodeRequest() (request *DescribeClusterVirtualNodeRequest) {
    request = &DescribeClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNode")
    
    
    return
}

func NewDescribeClusterVirtualNodeResponse() (response *DescribeClusterVirtualNodeResponse) {
    response = &DescribeClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterVirtualNode
// This API is used to view the Super Node list.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  INTERNALERROR_KUBERNETESLISTOPERATIONERROR = "InternalError.KubernetesListOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterVirtualNode(request *DescribeClusterVirtualNodeRequest) (response *DescribeClusterVirtualNodeResponse, err error) {
    return c.DescribeClusterVirtualNodeWithContext(context.Background(), request)
}

// DescribeClusterVirtualNode
// This API is used to view the Super Node list.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  INTERNALERROR_KUBERNETESLISTOPERATIONERROR = "InternalError.KubernetesListOperationError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterVirtualNodeWithContext(ctx context.Context, request *DescribeClusterVirtualNodeRequest) (response *DescribeClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterVirtualNodePoolsRequest() (request *DescribeClusterVirtualNodePoolsRequest) {
    request = &DescribeClusterVirtualNodePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterVirtualNodePools")
    
    
    return
}

func NewDescribeClusterVirtualNodePoolsResponse() (response *DescribeClusterVirtualNodePoolsResponse) {
    response = &DescribeClusterVirtualNodePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterVirtualNodePools
// This API is used to view the Super Node Pool list.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterVirtualNodePools(request *DescribeClusterVirtualNodePoolsRequest) (response *DescribeClusterVirtualNodePoolsResponse, err error) {
    return c.DescribeClusterVirtualNodePoolsWithContext(context.Background(), request)
}

// DescribeClusterVirtualNodePools
// This API is used to view the Super Node Pool list.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusterVirtualNodePoolsWithContext(ctx context.Context, request *DescribeClusterVirtualNodePoolsRequest) (response *DescribeClusterVirtualNodePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterVirtualNodePoolsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterVirtualNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterVirtualNodePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// This API is used to query clusters list.
//
// error code that may be returned:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// This API is used to query clusters list.
//
// error code that may be returned:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeECMInstancesRequest() (request *DescribeECMInstancesRequest) {
    request = &DescribeECMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeECMInstances")
    
    
    return
}

func NewDescribeECMInstancesResponse() (response *DescribeECMInstancesResponse) {
    response = &DescribeECMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeECMInstances
// This API is used to obtain the ECM instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeECMInstances(request *DescribeECMInstancesRequest) (response *DescribeECMInstancesResponse, err error) {
    return c.DescribeECMInstancesWithContext(context.Background(), request)
}

// DescribeECMInstances
// This API is used to obtain the ECM instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeECMInstancesWithContext(ctx context.Context, request *DescribeECMInstancesRequest) (response *DescribeECMInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeECMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeECMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeECMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeAvailableExtraArgsRequest() (request *DescribeEdgeAvailableExtraArgsRequest) {
    request = &DescribeEdgeAvailableExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeAvailableExtraArgs")
    
    
    return
}

func NewDescribeEdgeAvailableExtraArgsResponse() (response *DescribeEdgeAvailableExtraArgsResponse) {
    response = &DescribeEdgeAvailableExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeAvailableExtraArgs
// This API is used to query the custom parameters available for an edge cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeEdgeAvailableExtraArgs(request *DescribeEdgeAvailableExtraArgsRequest) (response *DescribeEdgeAvailableExtraArgsResponse, err error) {
    return c.DescribeEdgeAvailableExtraArgsWithContext(context.Background(), request)
}

// DescribeEdgeAvailableExtraArgs
// This API is used to query the custom parameters available for an edge cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeEdgeAvailableExtraArgsWithContext(ctx context.Context, request *DescribeEdgeAvailableExtraArgsRequest) (response *DescribeEdgeAvailableExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeAvailableExtraArgsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeAvailableExtraArgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeAvailableExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeCVMInstancesRequest() (request *DescribeEdgeCVMInstancesRequest) {
    request = &DescribeEdgeCVMInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeCVMInstances")
    
    
    return
}

func NewDescribeEdgeCVMInstancesResponse() (response *DescribeEdgeCVMInstancesResponse) {
    response = &DescribeEdgeCVMInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeCVMInstances
// This API is used to obtain the edge CVM instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeCVMInstances(request *DescribeEdgeCVMInstancesRequest) (response *DescribeEdgeCVMInstancesResponse, err error) {
    return c.DescribeEdgeCVMInstancesWithContext(context.Background(), request)
}

// DescribeEdgeCVMInstances
// This API is used to obtain the edge CVM instance information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeCVMInstancesWithContext(ctx context.Context, request *DescribeEdgeCVMInstancesRequest) (response *DescribeEdgeCVMInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeCVMInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeCVMInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeCVMInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterExtraArgsRequest() (request *DescribeEdgeClusterExtraArgsRequest) {
    request = &DescribeEdgeClusterExtraArgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterExtraArgs")
    
    
    return
}

func NewDescribeEdgeClusterExtraArgsResponse() (response *DescribeEdgeClusterExtraArgsResponse) {
    response = &DescribeEdgeClusterExtraArgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeClusterExtraArgs
// This API is used to query custom parameters of an edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterExtraArgs(request *DescribeEdgeClusterExtraArgsRequest) (response *DescribeEdgeClusterExtraArgsResponse, err error) {
    return c.DescribeEdgeClusterExtraArgsWithContext(context.Background(), request)
}

// DescribeEdgeClusterExtraArgs
// This API is used to query custom parameters of an edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterExtraArgsWithContext(ctx context.Context, request *DescribeEdgeClusterExtraArgsRequest) (response *DescribeEdgeClusterExtraArgsResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterExtraArgsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterExtraArgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterExtraArgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterInstancesRequest() (request *DescribeEdgeClusterInstancesRequest) {
    request = &DescribeEdgeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterInstances")
    
    
    return
}

func NewDescribeEdgeClusterInstancesResponse() (response *DescribeEdgeClusterInstancesResponse) {
    response = &DescribeEdgeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeClusterInstances
// This API is used to query the TKE Edge cluster node information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterInstances(request *DescribeEdgeClusterInstancesRequest) (response *DescribeEdgeClusterInstancesResponse, err error) {
    return c.DescribeEdgeClusterInstancesWithContext(context.Background(), request)
}

// DescribeEdgeClusterInstances
// This API is used to query the TKE Edge cluster node information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeClusterInstancesWithContext(ctx context.Context, request *DescribeEdgeClusterInstancesRequest) (response *DescribeEdgeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeClusterUpgradeInfoRequest() (request *DescribeEdgeClusterUpgradeInfoRequest) {
    request = &DescribeEdgeClusterUpgradeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeClusterUpgradeInfo")
    
    
    return
}

func NewDescribeEdgeClusterUpgradeInfoResponse() (response *DescribeEdgeClusterUpgradeInfoResponse) {
    response = &DescribeEdgeClusterUpgradeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeClusterUpgradeInfo
// This API is used to query the upgrade information of an edge cluster, including the upgradeable components, the current upgrade status, and errors occur during the upgrade.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEdgeClusterUpgradeInfo(request *DescribeEdgeClusterUpgradeInfoRequest) (response *DescribeEdgeClusterUpgradeInfoResponse, err error) {
    return c.DescribeEdgeClusterUpgradeInfoWithContext(context.Background(), request)
}

// DescribeEdgeClusterUpgradeInfo
// This API is used to query the upgrade information of an edge cluster, including the upgradeable components, the current upgrade status, and errors occur during the upgrade.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeEdgeClusterUpgradeInfoWithContext(ctx context.Context, request *DescribeEdgeClusterUpgradeInfoRequest) (response *DescribeEdgeClusterUpgradeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeClusterUpgradeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeClusterUpgradeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeClusterUpgradeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeLogSwitchesRequest() (request *DescribeEdgeLogSwitchesRequest) {
    request = &DescribeEdgeLogSwitchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEdgeLogSwitches")
    
    
    return
}

func NewDescribeEdgeLogSwitchesResponse() (response *DescribeEdgeLogSwitchesResponse) {
    response = &DescribeEdgeLogSwitchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeLogSwitches
// This API is used to query the status of event storage, cluster auditing and logging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeLogSwitches(request *DescribeEdgeLogSwitchesRequest) (response *DescribeEdgeLogSwitchesResponse, err error) {
    return c.DescribeEdgeLogSwitchesWithContext(context.Background(), request)
}

// DescribeEdgeLogSwitches
// This API is used to query the status of event storage, cluster auditing and logging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEdgeLogSwitchesWithContext(ctx context.Context, request *DescribeEdgeLogSwitchesRequest) (response *DescribeEdgeLogSwitchesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeLogSwitchesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeLogSwitches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeLogSwitchesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnableVpcCniProgressRequest() (request *DescribeEnableVpcCniProgressRequest) {
    request = &DescribeEnableVpcCniProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEnableVpcCniProgress")
    
    
    return
}

func NewDescribeEnableVpcCniProgressResponse() (response *DescribeEnableVpcCniProgressResponse) {
    response = &DescribeEnableVpcCniProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnableVpcCniProgress
// This API is used to query the task progress of enabling VPC-CNI mode.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEnableVpcCniProgress(request *DescribeEnableVpcCniProgressRequest) (response *DescribeEnableVpcCniProgressResponse, err error) {
    return c.DescribeEnableVpcCniProgressWithContext(context.Background(), request)
}

// DescribeEnableVpcCniProgress
// This API is used to query the task progress of enabling VPC-CNI mode.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEnableVpcCniProgressWithContext(ctx context.Context, request *DescribeEnableVpcCniProgressRequest) (response *DescribeEnableVpcCniProgressResponse, err error) {
    if request == nil {
        request = NewDescribeEnableVpcCniProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnableVpcCniProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnableVpcCniProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEncryptionStatusRequest() (request *DescribeEncryptionStatusRequest) {
    request = &DescribeEncryptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeEncryptionStatus")
    
    
    return
}

func NewDescribeEncryptionStatusResponse() (response *DescribeEncryptionStatusResponse) {
    response = &DescribeEncryptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEncryptionStatus
// This API is used to query the encryption status of etcd data.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEncryptionStatus(request *DescribeEncryptionStatusRequest) (response *DescribeEncryptionStatusResponse, err error) {
    return c.DescribeEncryptionStatusWithContext(context.Background(), request)
}

// DescribeEncryptionStatus
// This API is used to query the encryption status of etcd data.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEncryptionStatusWithContext(ctx context.Context, request *DescribeEncryptionStatusRequest) (response *DescribeEncryptionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeEncryptionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEncryptionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEncryptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExistedInstancesRequest() (request *DescribeExistedInstancesRequest) {
    request = &DescribeExistedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeExistedInstances")
    
    
    return
}

func NewDescribeExistedInstancesResponse() (response *DescribeExistedInstancesResponse) {
    response = &DescribeExistedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExistedInstances
// This API is used to query one or more existing node and determine whether they can be added to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExistedInstances(request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
    return c.DescribeExistedInstancesWithContext(context.Background(), request)
}

// DescribeExistedInstances
// This API is used to query one or more existing node and determine whether they can be added to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ASCOMMON = "FailedOperation.AsCommon"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_CLUSTERSTATE = "FailedOperation.ClusterState"
//  FAILEDOPERATION_CVMCOMMON = "FailedOperation.CvmCommon"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CREATEMASTERFAILED = "InternalError.CreateMasterFailed"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCCOMMON = "InternalError.VpcCommon"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExistedInstancesWithContext(ctx context.Context, request *DescribeExistedInstancesRequest) (response *DescribeExistedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeExistedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExistedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExistedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalNodeSupportConfigRequest() (request *DescribeExternalNodeSupportConfigRequest) {
    request = &DescribeExternalNodeSupportConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeExternalNodeSupportConfig")
    
    
    return
}

func NewDescribeExternalNodeSupportConfigResponse() (response *DescribeExternalNodeSupportConfigResponse) {
    response = &DescribeExternalNodeSupportConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalNodeSupportConfig
// This API is used to view third-party node pool configuration information.
//
// error code that may be returned:
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeExternalNodeSupportConfig(request *DescribeExternalNodeSupportConfigRequest) (response *DescribeExternalNodeSupportConfigResponse, err error) {
    return c.DescribeExternalNodeSupportConfigWithContext(context.Background(), request)
}

// DescribeExternalNodeSupportConfig
// This API is used to view third-party node pool configuration information.
//
// error code that may be returned:
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) DescribeExternalNodeSupportConfigWithContext(ctx context.Context, request *DescribeExternalNodeSupportConfigRequest) (response *DescribeExternalNodeSupportConfigResponse, err error) {
    if request == nil {
        request = NewDescribeExternalNodeSupportConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalNodeSupportConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalNodeSupportConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPAMDRequest() (request *DescribeIPAMDRequest) {
    request = &DescribeIPAMDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeIPAMD")
    
    
    return
}

func NewDescribeIPAMDResponse() (response *DescribeIPAMDResponse) {
    response = &DescribeIPAMDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPAMD
// This API is used to obtain eniipamd component information.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeIPAMD(request *DescribeIPAMDRequest) (response *DescribeIPAMDResponse, err error) {
    return c.DescribeIPAMDWithContext(context.Background(), request)
}

// DescribeIPAMD
// This API is used to obtain eniipamd component information.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeIPAMDWithContext(ctx context.Context, request *DescribeIPAMDRequest) (response *DescribeIPAMDResponse, err error) {
    if request == nil {
        request = NewDescribeIPAMDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPAMD require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPAMDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImages
// This API is used to get image information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// This API is used to get image information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogConfigsRequest() (request *DescribeLogConfigsRequest) {
    request = &DescribeLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeLogConfigs")
    
    
    return
}

func NewDescribeLogConfigsResponse() (response *DescribeLogConfigsResponse) {
    response = &DescribeLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogConfigs
// This API is used to query the log collection rules.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_KUBERNETESRESOURCENOTFOUND = "FailedOperation.KubernetesResourceNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeLogConfigs(request *DescribeLogConfigsRequest) (response *DescribeLogConfigsResponse, err error) {
    return c.DescribeLogConfigsWithContext(context.Background(), request)
}

// DescribeLogConfigs
// This API is used to query the log collection rules.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESCLIENTBUILDERROR = "FailedOperation.KubernetesClientBuildError"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_KUBERNETESRESOURCENOTFOUND = "FailedOperation.KubernetesResourceNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeLogConfigsWithContext(ctx context.Context, request *DescribeLogConfigsRequest) (response *DescribeLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeLogConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSwitchesRequest() (request *DescribeLogSwitchesRequest) {
    request = &DescribeLogSwitchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeLogSwitches")
    
    
    return
}

func NewDescribeLogSwitchesResponse() (response *DescribeLogSwitchesResponse) {
    response = &DescribeLogSwitchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogSwitches
// This API is used to query Cluster Log (Auditing, Event, Common Log) Switch List.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeLogSwitches(request *DescribeLogSwitchesRequest) (response *DescribeLogSwitchesResponse, err error) {
    return c.DescribeLogSwitchesWithContext(context.Background(), request)
}

// DescribeLogSwitches
// This API is used to query Cluster Log (Auditing, Event, Common Log) Switch List.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBERNETESLISTOPERATIONERROR = "FailedOperation.KubernetesListOperationError"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_KUBERNETESCLIENTBUILDERROR = "InternalError.KubernetesClientBuildError"
//  INTERNALERROR_KUBERNETESGETOPERATIONERROR = "InternalError.KubernetesGetOperationError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeLogSwitchesWithContext(ctx context.Context, request *DescribeLogSwitchesRequest) (response *DescribeLogSwitchesResponse, err error) {
    if request == nil {
        request = NewDescribeLogSwitchesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogSwitches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogSwitchesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePodChargeInfoRequest() (request *DescribePodChargeInfoRequest) {
    request = &DescribePodChargeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePodChargeInfo")
    
    
    return
}

func NewDescribePodChargeInfoResponse() (response *DescribePodChargeInfoResponse) {
    response = &DescribePodChargeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePodChargeInfo
// This API is used to query the billing information of running Pods. You can query a specific Pod by Namespace and Name or batch query by Pod Uid.
//
// error code that may be returned:
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePodChargeInfo(request *DescribePodChargeInfoRequest) (response *DescribePodChargeInfoResponse, err error) {
    return c.DescribePodChargeInfoWithContext(context.Background(), request)
}

// DescribePodChargeInfo
// This API is used to query the billing information of running Pods. You can query a specific Pod by Namespace and Name or batch query by Pod Uid.
//
// error code that may be returned:
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribePodChargeInfoWithContext(ctx context.Context, request *DescribePodChargeInfoRequest) (response *DescribePodChargeInfoResponse, err error) {
    if request == nil {
        request = NewDescribePodChargeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePodChargeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePodChargeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceRequest() (request *DescribePrometheusInstanceRequest) {
    request = &DescribePrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribePrometheusInstance")
    
    
    return
}

func NewDescribePrometheusInstanceResponse() (response *DescribePrometheusInstanceResponse) {
    response = &DescribePrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstance
// This API is used to obtain the instance details.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstance(request *DescribePrometheusInstanceRequest) (response *DescribePrometheusInstanceResponse, err error) {
    return c.DescribePrometheusInstanceWithContext(context.Background(), request)
}

// DescribePrometheusInstance
// This API is used to obtain the instance details.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstanceWithContext(ctx context.Context, request *DescribePrometheusInstanceRequest) (response *DescribePrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// This API is used to obtain all regions supported by TKE.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to obtain all regions supported by TKE.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeReservedInstanceUtilizationRateRequest() (request *DescribeReservedInstanceUtilizationRateRequest) {
    request = &DescribeReservedInstanceUtilizationRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeReservedInstanceUtilizationRate")
    
    
    return
}

func NewDescribeReservedInstanceUtilizationRateResponse() (response *DescribeReservedInstanceUtilizationRateResponse) {
    response = &DescribeReservedInstanceUtilizationRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReservedInstanceUtilizationRate
// This API is used to query the usage rate of various types of Reserved Coupons.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeReservedInstanceUtilizationRate(request *DescribeReservedInstanceUtilizationRateRequest) (response *DescribeReservedInstanceUtilizationRateResponse, err error) {
    return c.DescribeReservedInstanceUtilizationRateWithContext(context.Background(), request)
}

// DescribeReservedInstanceUtilizationRate
// This API is used to query the usage rate of various types of Reserved Coupons.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeReservedInstanceUtilizationRateWithContext(ctx context.Context, request *DescribeReservedInstanceUtilizationRateRequest) (response *DescribeReservedInstanceUtilizationRateResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstanceUtilizationRateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReservedInstanceUtilizationRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReservedInstanceUtilizationRateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
    request = &DescribeResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeResourceUsage")
    
    
    return
}

func NewDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
    response = &DescribeResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceUsage
// This API is used to query the cluster resource usage.
//
// error code that may be returned:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    return c.DescribeResourceUsageWithContext(context.Background(), request)
}

// DescribeResourceUsage
// This API is used to query the cluster resource usage.
//
// error code that may be returned:
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeResourceUsageWithContext(ctx context.Context, request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTableConflictsRequest() (request *DescribeRouteTableConflictsRequest) {
    request = &DescribeRouteTableConflictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeRouteTableConflicts")
    
    
    return
}

func NewDescribeRouteTableConflictsResponse() (response *DescribeRouteTableConflictsResponse) {
    response = &DescribeRouteTableConflictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRouteTableConflicts
// This API is used to query the list of route table conflicts.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRouteTableConflicts(request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
    return c.DescribeRouteTableConflictsWithContext(context.Background(), request)
}

// DescribeRouteTableConflicts
// This API is used to query the list of route table conflicts.
//
// error code that may be returned:
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  FAILEDOPERATION_VPCRECODRNOTFOUND = "FailedOperation.VpcRecodrNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CIDRMASKSIZEOUTOFRANGE = "InternalError.CidrMaskSizeOutOfRange"
//  INTERNALERROR_INVALIDPRIVATENETWORKCIDR = "InternalError.InvalidPrivateNetworkCidr"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_VPCRECODRNOTFOUND = "InternalError.VpcRecodrNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRouteTableConflictsWithContext(ctx context.Context, request *DescribeRouteTableConflictsRequest) (response *DescribeRouteTableConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTableConflictsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRouteTableConflicts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteTableConflictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportedRuntimeRequest() (request *DescribeSupportedRuntimeRequest) {
    request = &DescribeSupportedRuntimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeSupportedRuntime")
    
    
    return
}

func NewDescribeSupportedRuntimeResponse() (response *DescribeSupportedRuntimeResponse) {
    response = &DescribeSupportedRuntimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupportedRuntime
// This API is used to retrieve optional runtime versions based on K8S version.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSupportedRuntime(request *DescribeSupportedRuntimeRequest) (response *DescribeSupportedRuntimeResponse, err error) {
    return c.DescribeSupportedRuntimeWithContext(context.Background(), request)
}

// DescribeSupportedRuntime
// This API is used to retrieve optional runtime versions based on K8S version.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYSERVERCOMMONERROR = "FailedOperation.PolicyServerCommonError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSupportedRuntimeWithContext(ctx context.Context, request *DescribeSupportedRuntimeRequest) (response *DescribeSupportedRuntimeResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedRuntimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedRuntime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportedRuntimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClusterCredentialRequest() (request *DescribeTKEEdgeClusterCredentialRequest) {
    request = &DescribeTKEEdgeClusterCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterCredential")
    
    
    return
}

func NewDescribeTKEEdgeClusterCredentialResponse() (response *DescribeTKEEdgeClusterCredentialResponse) {
    response = &DescribeTKEEdgeClusterCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTKEEdgeClusterCredential
// This API is used to obtain the authentication information of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterCredential(request *DescribeTKEEdgeClusterCredentialRequest) (response *DescribeTKEEdgeClusterCredentialResponse, err error) {
    return c.DescribeTKEEdgeClusterCredentialWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusterCredential
// This API is used to obtain the authentication information of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterCredentialWithContext(ctx context.Context, request *DescribeTKEEdgeClusterCredentialRequest) (response *DescribeTKEEdgeClusterCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClusterCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusterCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClusterCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClusterStatusRequest() (request *DescribeTKEEdgeClusterStatusRequest) {
    request = &DescribeTKEEdgeClusterStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusterStatus")
    
    
    return
}

func NewDescribeTKEEdgeClusterStatusResponse() (response *DescribeTKEEdgeClusterStatusResponse) {
    response = &DescribeTKEEdgeClusterStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTKEEdgeClusterStatus
// This API is used to query the current status and process information of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterStatus(request *DescribeTKEEdgeClusterStatusRequest) (response *DescribeTKEEdgeClusterStatusResponse, err error) {
    return c.DescribeTKEEdgeClusterStatusWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusterStatus
// This API is used to query the current status and process information of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusterStatusWithContext(ctx context.Context, request *DescribeTKEEdgeClusterStatusRequest) (response *DescribeTKEEdgeClusterStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClusterStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusterStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClusterStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeClustersRequest() (request *DescribeTKEEdgeClustersRequest) {
    request = &DescribeTKEEdgeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeClusters")
    
    
    return
}

func NewDescribeTKEEdgeClustersResponse() (response *DescribeTKEEdgeClustersResponse) {
    response = &DescribeTKEEdgeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTKEEdgeClusters
// This API is used to query the list of TKE Edge clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClusters(request *DescribeTKEEdgeClustersRequest) (response *DescribeTKEEdgeClustersResponse, err error) {
    return c.DescribeTKEEdgeClustersWithContext(context.Background(), request)
}

// DescribeTKEEdgeClusters
// This API is used to query the list of TKE Edge clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeClustersWithContext(ctx context.Context, request *DescribeTKEEdgeClustersRequest) (response *DescribeTKEEdgeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeExternalKubeconfigRequest() (request *DescribeTKEEdgeExternalKubeconfigRequest) {
    request = &DescribeTKEEdgeExternalKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeExternalKubeconfig")
    
    
    return
}

func NewDescribeTKEEdgeExternalKubeconfigResponse() (response *DescribeTKEEdgeExternalKubeconfigResponse) {
    response = &DescribeTKEEdgeExternalKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTKEEdgeExternalKubeconfig
// This API is used to obtain the kubeconfig for access to a TKE Edge cluster through the public network.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeExternalKubeconfig(request *DescribeTKEEdgeExternalKubeconfigRequest) (response *DescribeTKEEdgeExternalKubeconfigResponse, err error) {
    return c.DescribeTKEEdgeExternalKubeconfigWithContext(context.Background(), request)
}

// DescribeTKEEdgeExternalKubeconfig
// This API is used to obtain the kubeconfig for access to a TKE Edge cluster through the public network.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeExternalKubeconfigWithContext(ctx context.Context, request *DescribeTKEEdgeExternalKubeconfigRequest) (response *DescribeTKEEdgeExternalKubeconfigResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeExternalKubeconfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeExternalKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeExternalKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTKEEdgeScriptRequest() (request *DescribeTKEEdgeScriptRequest) {
    request = &DescribeTKEEdgeScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeTKEEdgeScript")
    
    
    return
}

func NewDescribeTKEEdgeScriptResponse() (response *DescribeTKEEdgeScriptResponse) {
    response = &DescribeTKEEdgeScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTKEEdgeScript
// This API is used to query the URL of TKE edge script. You can add external nodes to a TKE Edge cluster by downloading the URL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeScript(request *DescribeTKEEdgeScriptRequest) (response *DescribeTKEEdgeScriptResponse, err error) {
    return c.DescribeTKEEdgeScriptWithContext(context.Background(), request)
}

// DescribeTKEEdgeScript
// This API is used to query the URL of TKE edge script. You can add external nodes to a TKE Edge cluster by downloading the URL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTKEEdgeScriptWithContext(ctx context.Context, request *DescribeTKEEdgeScriptRequest) (response *DescribeTKEEdgeScriptResponse, err error) {
    if request == nil {
        request = NewDescribeTKEEdgeScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTKEEdgeScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTKEEdgeScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionsRequest() (request *DescribeVersionsRequest) {
    request = &DescribeVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeVersions")
    
    
    return
}

func NewDescribeVersionsResponse() (response *DescribeVersionsResponse) {
    response = &DescribeVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVersions
// This API is used to query cluster version information.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVersions(request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
    return c.DescribeVersionsWithContext(context.Background(), request)
}

// DescribeVersions
// This API is used to query cluster version information.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLINETHTTP = "FailedOperation.ComponentClinetHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_COMPONENTCLINETHTTP = "InternalError.ComponentClinetHttp"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVersionsWithContext(ctx context.Context, request *DescribeVersionsRequest) (response *DescribeVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcCniPodLimitsRequest() (request *DescribeVpcCniPodLimitsRequest) {
    request = &DescribeVpcCniPodLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeVpcCniPodLimits")
    
    
    return
}

func NewDescribeVpcCniPodLimitsResponse() (response *DescribeVpcCniPodLimitsResponse) {
    response = &DescribeVpcCniPodLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcCniPodLimits
// This API is used to query the maximum number of Pods in the VPC-CNI network mode supported by the models in the specified availability zone of the current user and region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcCniPodLimits(request *DescribeVpcCniPodLimitsRequest) (response *DescribeVpcCniPodLimitsResponse, err error) {
    return c.DescribeVpcCniPodLimitsWithContext(context.Background(), request)
}

// DescribeVpcCniPodLimits
// This API is used to query the maximum number of Pods in the VPC-CNI network mode supported by the models in the specified availability zone of the current user and region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcCniPodLimitsWithContext(ctx context.Context, request *DescribeVpcCniPodLimitsRequest) (response *DescribeVpcCniPodLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcCniPodLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcCniPodLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcCniPodLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClusterDeletionProtectionRequest() (request *DisableClusterDeletionProtectionRequest) {
    request = &DisableClusterDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableClusterDeletionProtection")
    
    
    return
}

func NewDisableClusterDeletionProtectionResponse() (response *DisableClusterDeletionProtectionResponse) {
    response = &DisableClusterDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableClusterDeletionProtection
// This API is used to disable cluster deletion protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableClusterDeletionProtection(request *DisableClusterDeletionProtectionRequest) (response *DisableClusterDeletionProtectionResponse, err error) {
    return c.DisableClusterDeletionProtectionWithContext(context.Background(), request)
}

// DisableClusterDeletionProtection
// This API is used to disable cluster deletion protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableClusterDeletionProtectionWithContext(ctx context.Context, request *DisableClusterDeletionProtectionRequest) (response *DisableClusterDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewDisableClusterDeletionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClusterDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClusterDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDisableEncryptionProtectionRequest() (request *DisableEncryptionProtectionRequest) {
    request = &DisableEncryptionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DisableEncryptionProtection")
    
    
    return
}

func NewDisableEncryptionProtectionResponse() (response *DisableEncryptionProtectionResponse) {
    response = &DisableEncryptionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableEncryptionProtection
// This API is used to disable encryption protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableEncryptionProtection(request *DisableEncryptionProtectionRequest) (response *DisableEncryptionProtectionResponse, err error) {
    return c.DisableEncryptionProtectionWithContext(context.Background(), request)
}

// DisableEncryptionProtection
// This API is used to disable encryption protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableEncryptionProtectionWithContext(ctx context.Context, request *DisableEncryptionProtectionRequest) (response *DisableEncryptionProtectionResponse, err error) {
    if request == nil {
        request = NewDisableEncryptionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableEncryptionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableEncryptionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDrainClusterVirtualNodeRequest() (request *DrainClusterVirtualNodeRequest) {
    request = &DrainClusterVirtualNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DrainClusterVirtualNode")
    
    
    return
}

func NewDrainClusterVirtualNodeResponse() (response *DrainClusterVirtualNodeResponse) {
    response = &DrainClusterVirtualNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DrainClusterVirtualNode
// This API is used to evict the Super Node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DrainClusterVirtualNode(request *DrainClusterVirtualNodeRequest) (response *DrainClusterVirtualNodeResponse, err error) {
    return c.DrainClusterVirtualNodeWithContext(context.Background(), request)
}

// DrainClusterVirtualNode
// This API is used to evict the Super Node.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DrainClusterVirtualNodeWithContext(ctx context.Context, request *DrainClusterVirtualNodeRequest) (response *DrainClusterVirtualNodeResponse, err error) {
    if request == nil {
        request = NewDrainClusterVirtualNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DrainClusterVirtualNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDrainClusterVirtualNodeResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClusterDeletionProtectionRequest() (request *EnableClusterDeletionProtectionRequest) {
    request = &EnableClusterDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableClusterDeletionProtection")
    
    
    return
}

func NewEnableClusterDeletionProtectionResponse() (response *EnableClusterDeletionProtectionResponse) {
    response = &EnableClusterDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableClusterDeletionProtection
// This API is used to enable cluster deletion protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableClusterDeletionProtection(request *EnableClusterDeletionProtectionRequest) (response *EnableClusterDeletionProtectionResponse, err error) {
    return c.EnableClusterDeletionProtectionWithContext(context.Background(), request)
}

// EnableClusterDeletionProtection
// This API is used to enable cluster deletion protection.
//
// error code that may be returned:
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableClusterDeletionProtectionWithContext(ctx context.Context, request *EnableClusterDeletionProtectionRequest) (response *EnableClusterDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewEnableClusterDeletionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClusterDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClusterDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewEnableEncryptionProtectionRequest() (request *EnableEncryptionProtectionRequest) {
    request = &EnableEncryptionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableEncryptionProtection")
    
    
    return
}

func NewEnableEncryptionProtectionResponse() (response *EnableEncryptionProtectionResponse) {
    response = &EnableEncryptionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableEncryptionProtection
// This API is used to enable Encrypted Data Protection, which requires enabling KMS capability and completing KMS authorization.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) EnableEncryptionProtection(request *EnableEncryptionProtectionRequest) (response *EnableEncryptionProtectionResponse, err error) {
    return c.EnableEncryptionProtectionWithContext(context.Background(), request)
}

// EnableEncryptionProtection
// This API is used to enable Encrypted Data Protection, which requires enabling KMS capability and completing KMS authorization.
//
// error code that may be returned:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) EnableEncryptionProtectionWithContext(ctx context.Context, request *EnableEncryptionProtectionRequest) (response *EnableEncryptionProtectionResponse, err error) {
    if request == nil {
        request = NewEnableEncryptionProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableEncryptionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableEncryptionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewEnableVpcCniNetworkTypeRequest() (request *EnableVpcCniNetworkTypeRequest) {
    request = &EnableVpcCniNetworkTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "EnableVpcCniNetworkType")
    
    
    return
}

func NewEnableVpcCniNetworkTypeResponse() (response *EnableVpcCniNetworkTypeResponse) {
    response = &EnableVpcCniNetworkTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableVpcCniNetworkType
// This API is used to enable the VPC-CNI network mode for GR clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITENABLEVPCCNI = "UnsupportedOperation.ClusterNotSuitEnableVPCCNI"
func (c *Client) EnableVpcCniNetworkType(request *EnableVpcCniNetworkTypeRequest) (response *EnableVpcCniNetworkTypeResponse, err error) {
    return c.EnableVpcCniNetworkTypeWithContext(context.Background(), request)
}

// EnableVpcCniNetworkType
// This API is used to enable the VPC-CNI network mode for GR clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_ENABLEVPCCNIFAILED = "FailedOperation.EnableVPCCNIFailed"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  UNSUPPORTEDOPERATION_CLUSTERNOTSUITENABLEVPCCNI = "UnsupportedOperation.ClusterNotSuitEnableVPCCNI"
func (c *Client) EnableVpcCniNetworkTypeWithContext(ctx context.Context, request *EnableVpcCniNetworkTypeRequest) (response *EnableVpcCniNetworkTypeResponse, err error) {
    if request == nil {
        request = NewEnableVpcCniNetworkTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableVpcCniNetworkType require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableVpcCniNetworkTypeResponse()
    err = c.Send(request, response)
    return
}

func NewForwardTKEEdgeApplicationRequestV3Request() (request *ForwardTKEEdgeApplicationRequestV3Request) {
    request = &ForwardTKEEdgeApplicationRequestV3Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ForwardTKEEdgeApplicationRequestV3")
    
    
    return
}

func NewForwardTKEEdgeApplicationRequestV3Response() (response *ForwardTKEEdgeApplicationRequestV3Response) {
    response = &ForwardTKEEdgeApplicationRequestV3Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForwardTKEEdgeApplicationRequestV3
// This API is used to work with the add-ons of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardTKEEdgeApplicationRequestV3(request *ForwardTKEEdgeApplicationRequestV3Request) (response *ForwardTKEEdgeApplicationRequestV3Response, err error) {
    return c.ForwardTKEEdgeApplicationRequestV3WithContext(context.Background(), request)
}

// ForwardTKEEdgeApplicationRequestV3
// This API is used to work with the add-ons of a TKE Edge cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTINWHITELIST = "UnsupportedOperation.NotInWhitelist"
func (c *Client) ForwardTKEEdgeApplicationRequestV3WithContext(ctx context.Context, request *ForwardTKEEdgeApplicationRequestV3Request) (response *ForwardTKEEdgeApplicationRequestV3Response, err error) {
    if request == nil {
        request = NewForwardTKEEdgeApplicationRequestV3Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForwardTKEEdgeApplicationRequestV3 require credential")
    }

    request.SetContext(ctx)
    
    response = NewForwardTKEEdgeApplicationRequestV3Response()
    err = c.Send(request, response)
    return
}

func NewGetClusterLevelPriceRequest() (request *GetClusterLevelPriceRequest) {
    request = &GetClusterLevelPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetClusterLevelPrice")
    
    
    return
}

func NewGetClusterLevelPriceResponse() (response *GetClusterLevelPriceResponse) {
    response = &GetClusterLevelPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetClusterLevelPrice
// Obtaining the price of specified cluster model
//
// error code that may be returned:
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
func (c *Client) GetClusterLevelPrice(request *GetClusterLevelPriceRequest) (response *GetClusterLevelPriceResponse, err error) {
    return c.GetClusterLevelPriceWithContext(context.Background(), request)
}

// GetClusterLevelPrice
// Obtaining the price of specified cluster model
//
// error code that may be returned:
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
func (c *Client) GetClusterLevelPriceWithContext(ctx context.Context, request *GetClusterLevelPriceRequest) (response *GetClusterLevelPriceResponse, err error) {
    if request == nil {
        request = NewGetClusterLevelPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetClusterLevelPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetClusterLevelPriceResponse()
    err = c.Send(request, response)
    return
}

func NewGetTkeAppChartListRequest() (request *GetTkeAppChartListRequest) {
    request = &GetTkeAppChartListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetTkeAppChartList")
    
    
    return
}

func NewGetTkeAppChartListResponse() (response *GetTkeAppChartListResponse) {
    response = &GetTkeAppChartListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTkeAppChartList
// This API is used to retrieve the App List supported by TKE.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTkeAppChartList(request *GetTkeAppChartListRequest) (response *GetTkeAppChartListResponse, err error) {
    return c.GetTkeAppChartListWithContext(context.Background(), request)
}

// GetTkeAppChartList
// This API is used to retrieve the App List supported by TKE.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTkeAppChartListWithContext(ctx context.Context, request *GetTkeAppChartListRequest) (response *GetTkeAppChartListResponse, err error) {
    if request == nil {
        request = NewGetTkeAppChartListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTkeAppChartList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTkeAppChartListResponse()
    err = c.Send(request, response)
    return
}

func NewGetUpgradeInstanceProgressRequest() (request *GetUpgradeInstanceProgressRequest) {
    request = &GetUpgradeInstanceProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "GetUpgradeInstanceProgress")
    
    
    return
}

func NewGetUpgradeInstanceProgressResponse() (response *GetUpgradeInstanceProgressResponse) {
    response = &GetUpgradeInstanceProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUpgradeInstanceProgress
// This API is used to obtain the current progress of node upgrade. If the cluster is not in node upgrade status, the API will report an error: Task not found.
//
// error code that may be returned:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetUpgradeInstanceProgress(request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
    return c.GetUpgradeInstanceProgressWithContext(context.Background(), request)
}

// GetUpgradeInstanceProgress
// This API is used to obtain the current progress of node upgrade. If the cluster is not in node upgrade status, the API will report an error: Task not found.
//
// error code that may be returned:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) GetUpgradeInstanceProgressWithContext(ctx context.Context, request *GetUpgradeInstanceProgressRequest) (response *GetUpgradeInstanceProgressResponse, err error) {
    if request == nil {
        request = NewGetUpgradeInstanceProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUpgradeInstanceProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUpgradeInstanceProgressResponse()
    err = c.Send(request, response)
    return
}

func NewInstallAddonRequest() (request *InstallAddonRequest) {
    request = &InstallAddonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "InstallAddon")
    
    
    return
}

func NewInstallAddonResponse() (response *InstallAddonResponse) {
    response = &InstallAddonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallAddon
// This API is used to install an add-on on the target cluster.
//
// error code that may be returned:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) InstallAddon(request *InstallAddonRequest) (response *InstallAddonResponse, err error) {
    return c.InstallAddonWithContext(context.Background(), request)
}

// InstallAddon
// This API is used to install an add-on on the target cluster.
//
// error code that may be returned:
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) InstallAddonWithContext(ctx context.Context, request *InstallAddonRequest) (response *InstallAddonResponse, err error) {
    if request == nil {
        request = NewInstallAddonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallAddon require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallAddonResponse()
    err = c.Send(request, response)
    return
}

func NewInstallEdgeLogAgentRequest() (request *InstallEdgeLogAgentRequest) {
    request = &InstallEdgeLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "InstallEdgeLogAgent")
    
    
    return
}

func NewInstallEdgeLogAgentResponse() (response *InstallEdgeLogAgentResponse) {
    response = &InstallEdgeLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallEdgeLogAgent
// This API is used to install the log collection add-on on TKE Edge cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) InstallEdgeLogAgent(request *InstallEdgeLogAgentRequest) (response *InstallEdgeLogAgentResponse, err error) {
    return c.InstallEdgeLogAgentWithContext(context.Background(), request)
}

// InstallEdgeLogAgent
// This API is used to install the log collection add-on on TKE Edge cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) InstallEdgeLogAgentWithContext(ctx context.Context, request *InstallEdgeLogAgentRequest) (response *InstallEdgeLogAgentResponse, err error) {
    if request == nil {
        request = NewInstallEdgeLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallEdgeLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallEdgeLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupAttributeRequest() (request *ModifyClusterAsGroupAttributeRequest) {
    request = &ModifyClusterAsGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupAttribute")
    
    
    return
}

func NewModifyClusterAsGroupAttributeResponse() (response *ModifyClusterAsGroupAttributeResponse) {
    response = &ModifyClusterAsGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterAsGroupAttribute
// Modify cluster scaling group attributes
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupAttribute(request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
    return c.ModifyClusterAsGroupAttributeWithContext(context.Background(), request)
}

// ModifyClusterAsGroupAttribute
// Modify cluster scaling group attributes
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ASCOMMONERROR = "InvalidParameter.AsCommonError"
//  INVALIDPARAMETER_CIDROUTOFROUTETABLE = "InvalidParameter.CidrOutOfRouteTable"
//  INVALIDPARAMETER_GATEWAYALREADYASSOCIATEDCIDR = "InvalidParameter.GatewayAlreadyAssociatedCidr"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_ROUTETABLENOTEMPTY = "InvalidParameter.RouteTableNotEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupAttributeWithContext(ctx context.Context, request *ModifyClusterAsGroupAttributeRequest) (response *ModifyClusterAsGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAsGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAsGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAsGroupOptionAttributeRequest() (request *ModifyClusterAsGroupOptionAttributeRequest) {
    request = &ModifyClusterAsGroupOptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAsGroupOptionAttribute")
    
    
    return
}

func NewModifyClusterAsGroupOptionAttributeResponse() (response *ModifyClusterAsGroupOptionAttributeResponse) {
    response = &ModifyClusterAsGroupOptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterAsGroupOptionAttribute
// This API is used to modify cluster auto scaling attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupOptionAttribute(request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
    return c.ModifyClusterAsGroupOptionAttributeWithContext(context.Background(), request)
}

// ModifyClusterAsGroupOptionAttribute
// This API is used to modify cluster auto scaling attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBERNETESINTERNAL = "FailedOperation.KubernetesInternal"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASCOMMON = "InternalError.AsCommon"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_CLUSTERSTATE = "InternalError.ClusterState"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_CVMNOTFOUND = "InternalError.CvmNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterAsGroupOptionAttributeWithContext(ctx context.Context, request *ModifyClusterAsGroupOptionAttributeRequest) (response *ModifyClusterAsGroupOptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAsGroupOptionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAsGroupOptionAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAsGroupOptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAttributeRequest() (request *ModifyClusterAttributeRequest) {
    request = &ModifyClusterAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAttribute")
    
    
    return
}

func NewModifyClusterAttributeResponse() (response *ModifyClusterAttributeResponse) {
    response = &ModifyClusterAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterAttribute
// This API is used to modify cluster attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_TRADEINSUFFICIENTBALANCE = "InternalError.TradeInsufficientBalance"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAttribute(request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
    return c.ModifyClusterAttributeWithContext(context.Background(), request)
}

// ModifyClusterAttribute
// This API is used to modify cluster attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  FAILEDOPERATION_UNEXPECTEDERROR = "FailedOperation.UnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TRADECOMMON = "InternalError.TradeCommon"
//  INTERNALERROR_TRADEINSUFFICIENTBALANCE = "InternalError.TradeInsufficientBalance"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAttributeWithContext(ctx context.Context, request *ModifyClusterAttributeRequest) (response *ModifyClusterAttributeResponse, err error) {
    if request == nil {
        request = NewModifyClusterAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAuthenticationOptionsRequest() (request *ModifyClusterAuthenticationOptionsRequest) {
    request = &ModifyClusterAuthenticationOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterAuthenticationOptions")
    
    
    return
}

func NewModifyClusterAuthenticationOptionsResponse() (response *ModifyClusterAuthenticationOptionsResponse) {
    response = &ModifyClusterAuthenticationOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterAuthenticationOptions
// This API is used to modify cluster authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAuthenticationOptions(request *ModifyClusterAuthenticationOptionsRequest) (response *ModifyClusterAuthenticationOptionsResponse, err error) {
    return c.ModifyClusterAuthenticationOptionsWithContext(context.Background(), request)
}

// ModifyClusterAuthenticationOptions
// This API is used to modify cluster authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) ModifyClusterAuthenticationOptionsWithContext(ctx context.Context, request *ModifyClusterAuthenticationOptionsRequest) (response *ModifyClusterAuthenticationOptionsResponse, err error) {
    if request == nil {
        request = NewModifyClusterAuthenticationOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAuthenticationOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAuthenticationOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterEndpointSPRequest() (request *ModifyClusterEndpointSPRequest) {
    request = &ModifyClusterEndpointSPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterEndpointSP")
    
    
    return
}

func NewModifyClusterEndpointSPResponse() (response *ModifyClusterEndpointSPResponse) {
    response = &ModifyClusterEndpointSPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterEndpointSP
// Modify the security policy of the external port of the managed cluster (the old way, only the external port of the managed cluster is supported)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  FAILEDOPERATION_VPCUNEXPECTEDERROR = "FailedOperation.VPCUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCUNEXPECTEDERROR = "InternalError.VPCUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterEndpointSP(request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
    return c.ModifyClusterEndpointSPWithContext(context.Background(), request)
}

// ModifyClusterEndpointSP
// Modify the security policy of the external port of the managed cluster (the old way, only the external port of the managed cluster is supported)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_KUBERNETESGETOPERATIONERROR = "FailedOperation.KubernetesGetOperationError"
//  FAILEDOPERATION_LBCOMMON = "FailedOperation.LbCommon"
//  FAILEDOPERATION_VPCUNEXPECTEDERROR = "FailedOperation.VPCUnexpectedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INTERNALERROR_VPCUNEXPECTEDERROR = "InternalError.VPCUnexpectedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClusterEndpointSPWithContext(ctx context.Context, request *ModifyClusterEndpointSPRequest) (response *ModifyClusterEndpointSPResponse, err error) {
    if request == nil {
        request = NewModifyClusterEndpointSPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterEndpointSP require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterEndpointSPResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterImageRequest() (request *ModifyClusterImageRequest) {
    request = &ModifyClusterImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterImage")
    
    
    return
}

func NewModifyClusterImageResponse() (response *ModifyClusterImageResponse) {
    response = &ModifyClusterImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterImage
// This API is used to modify the cluster image.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
func (c *Client) ModifyClusterImage(request *ModifyClusterImageRequest) (response *ModifyClusterImageResponse, err error) {
    return c.ModifyClusterImageWithContext(context.Background(), request)
}

// ModifyClusterImage
// This API is used to modify the cluster image.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_PARAM = "FailedOperation.Param"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CVMCOMMON = "InternalError.CvmCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_IMAGEIDNOTFOUND = "InternalError.ImageIdNotFound"
//  INTERNALERROR_OSNOTSUPPORT = "InternalError.OsNotSupport"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VERSIONNOTSUPPORTCGROUPV2 = "InvalidParameter.VersionNotSupportCgroupV2"
func (c *Client) ModifyClusterImageWithContext(ctx context.Context, request *ModifyClusterImageRequest) (response *ModifyClusterImageResponse, err error) {
    if request == nil {
        request = NewModifyClusterImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNodePoolRequest() (request *ModifyClusterNodePoolRequest) {
    request = &ModifyClusterNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterNodePool")
    
    
    return
}

func NewModifyClusterNodePoolResponse() (response *ModifyClusterNodePoolResponse) {
    response = &ModifyClusterNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterNodePool
// This API is used to edit a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  UNSUPPORTEDOPERATION_CAENABLEFAILED = "UnsupportedOperation.CaEnableFailed"
func (c *Client) ModifyClusterNodePool(request *ModifyClusterNodePoolRequest) (response *ModifyClusterNodePoolResponse, err error) {
    return c.ModifyClusterNodePoolWithContext(context.Background(), request)
}

// ModifyClusterNodePool
// This API is used to edit a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ASASGNOTEXIST = "ResourceNotFound.AsAsgNotExist"
//  UNSUPPORTEDOPERATION_CAENABLEFAILED = "UnsupportedOperation.CaEnableFailed"
func (c *Client) ModifyClusterNodePoolWithContext(ctx context.Context, request *ModifyClusterNodePoolRequest) (response *ModifyClusterNodePoolResponse, err error) {
    if request == nil {
        request = NewModifyClusterNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterRuntimeConfigRequest() (request *ModifyClusterRuntimeConfigRequest) {
    request = &ModifyClusterRuntimeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterRuntimeConfig")
    
    
    return
}

func NewModifyClusterRuntimeConfigResponse() (response *ModifyClusterRuntimeConfigResponse) {
    response = &ModifyClusterRuntimeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterRuntimeConfig
// This API is used to modify the latitude runtime configuration of clusters and node pools.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyClusterRuntimeConfig(request *ModifyClusterRuntimeConfigRequest) (response *ModifyClusterRuntimeConfigResponse, err error) {
    return c.ModifyClusterRuntimeConfigWithContext(context.Background(), request)
}

// ModifyClusterRuntimeConfig
// This API is used to modify the latitude runtime configuration of clusters and node pools.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyClusterRuntimeConfigWithContext(ctx context.Context, request *ModifyClusterRuntimeConfigRequest) (response *ModifyClusterRuntimeConfigResponse, err error) {
    if request == nil {
        request = NewModifyClusterRuntimeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterRuntimeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterRuntimeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterTagsRequest() (request *ModifyClusterTagsRequest) {
    request = &ModifyClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterTags")
    
    
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
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyClusterTags(request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
    return c.ModifyClusterTagsWithContext(context.Background(), request)
}

// ModifyClusterTags
// This API is used to modify cluster tags.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"
//  FAILEDOPERATION_TRADECOMMON = "FailedOperation.TradeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKCREATEFAILED = "InternalError.TaskCreateFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
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

func NewModifyClusterVirtualNodePoolRequest() (request *ModifyClusterVirtualNodePoolRequest) {
    request = &ModifyClusterVirtualNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterVirtualNodePool")
    
    
    return
}

func NewModifyClusterVirtualNodePoolResponse() (response *ModifyClusterVirtualNodePoolResponse) {
    response = &ModifyClusterVirtualNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterVirtualNodePool
// This API is used to modify the Super Node Pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
func (c *Client) ModifyClusterVirtualNodePool(request *ModifyClusterVirtualNodePoolRequest) (response *ModifyClusterVirtualNodePoolResponse, err error) {
    return c.ModifyClusterVirtualNodePoolWithContext(context.Background(), request)
}

// ModifyClusterVirtualNodePool
// This API is used to modify the Super Node Pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NODEPOOLSTATENOTNORMAL = "ResourceUnavailable.NodePoolStateNotNormal"
func (c *Client) ModifyClusterVirtualNodePoolWithContext(ctx context.Context, request *ModifyClusterVirtualNodePoolRequest) (response *ModifyClusterVirtualNodePoolResponse, err error) {
    if request == nil {
        request = NewModifyClusterVirtualNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterVirtualNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterVirtualNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodePoolInstanceTypesRequest() (request *ModifyNodePoolInstanceTypesRequest) {
    request = &ModifyNodePoolInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePoolInstanceTypes")
    
    
    return
}

func NewModifyNodePoolInstanceTypesResponse() (response *ModifyNodePoolInstanceTypesResponse) {
    response = &ModifyNodePoolInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNodePoolInstanceTypes
// This API is used to modify the model of instances in a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNodePoolInstanceTypes(request *ModifyNodePoolInstanceTypesRequest) (response *ModifyNodePoolInstanceTypesResponse, err error) {
    return c.ModifyNodePoolInstanceTypesWithContext(context.Background(), request)
}

// ModifyNodePoolInstanceTypes
// This API is used to modify the model of instances in a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNodePoolInstanceTypesWithContext(ctx context.Context, request *ModifyNodePoolInstanceTypesRequest) (response *ModifyNodePoolInstanceTypesResponse, err error) {
    if request == nil {
        request = NewModifyNodePoolInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodePoolInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodePoolInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAlertRuleRequest() (request *ModifyPrometheusAlertRuleRequest) {
    request = &ModifyPrometheusAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyPrometheusAlertRule")
    
    
    return
}

func NewModifyPrometheusAlertRuleResponse() (response *ModifyPrometheusAlertRuleResponse) {
    response = &ModifyPrometheusAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusAlertRule
// This API is used to modify an alert rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertRule(request *ModifyPrometheusAlertRuleRequest) (response *ModifyPrometheusAlertRuleResponse, err error) {
    return c.ModifyPrometheusAlertRuleWithContext(context.Background(), request)
}

// ModifyPrometheusAlertRule
// This API is used to modify an alert rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertRuleWithContext(ctx context.Context, request *ModifyPrometheusAlertRuleRequest) (response *ModifyPrometheusAlertRuleResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAlertRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveNodeFromNodePoolRequest() (request *RemoveNodeFromNodePoolRequest) {
    request = &RemoveNodeFromNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RemoveNodeFromNodePool")
    
    
    return
}

func NewRemoveNodeFromNodePoolResponse() (response *RemoveNodeFromNodePoolResponse) {
    response = &RemoveNodeFromNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveNodeFromNodePool
// This API is used to remove a node from a node pool but retain it in the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveNodeFromNodePool(request *RemoveNodeFromNodePoolRequest) (response *RemoveNodeFromNodePoolResponse, err error) {
    return c.RemoveNodeFromNodePoolWithContext(context.Background(), request)
}

// RemoveNodeFromNodePool
// This API is used to remove a node from a node pool but retain it in the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveNodeFromNodePoolWithContext(ctx context.Context, request *RemoveNodeFromNodePoolRequest) (response *RemoveNodeFromNodePoolResponse, err error) {
    if request == nil {
        request = NewRemoveNodeFromNodePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveNodeFromNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveNodeFromNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewSetNodePoolNodeProtectionRequest() (request *SetNodePoolNodeProtectionRequest) {
    request = &SetNodePoolNodeProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "SetNodePoolNodeProtection")
    
    
    return
}

func NewSetNodePoolNodeProtectionResponse() (response *SetNodePoolNodeProtectionResponse) {
    response = &SetNodePoolNodeProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNodePoolNodeProtection
// This API is used to enable removal protection for the nodes automatically created by the scaling group in a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetNodePoolNodeProtection(request *SetNodePoolNodeProtectionRequest) (response *SetNodePoolNodeProtectionResponse, err error) {
    return c.SetNodePoolNodeProtectionWithContext(context.Background(), request)
}

// SetNodePoolNodeProtection
// This API is used to enable removal protection for the nodes automatically created by the scaling group in a node pool.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ACCOUNTUSERNOTAUTHENTICATED = "InternalError.AccountUserNotAuthenticated"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetNodePoolNodeProtectionWithContext(ctx context.Context, request *SetNodePoolNodeProtectionRequest) (response *SetNodePoolNodeProtectionResponse, err error) {
    if request == nil {
        request = NewSetNodePoolNodeProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNodePoolNodeProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNodePoolNodeProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallEdgeLogAgentRequest() (request *UninstallEdgeLogAgentRequest) {
    request = &UninstallEdgeLogAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UninstallEdgeLogAgent")
    
    
    return
}

func NewUninstallEdgeLogAgentResponse() (response *UninstallEdgeLogAgentResponse) {
    response = &UninstallEdgeLogAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallEdgeLogAgent
// This API is used to uninstall the log collection add-on from TKE Edge cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UninstallEdgeLogAgent(request *UninstallEdgeLogAgentRequest) (response *UninstallEdgeLogAgentResponse, err error) {
    return c.UninstallEdgeLogAgentWithContext(context.Background(), request)
}

// UninstallEdgeLogAgent
// This API is used to uninstall the log collection add-on from TKE Edge cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UninstallEdgeLogAgentWithContext(ctx context.Context, request *UninstallEdgeLogAgentRequest) (response *UninstallEdgeLogAgentResponse, err error) {
    if request == nil {
        request = NewUninstallEdgeLogAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallEdgeLogAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallEdgeLogAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAddonRequest() (request *UpdateAddonRequest) {
    request = &UpdateAddonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateAddon")
    
    
    return
}

func NewUpdateAddonResponse() (response *UpdateAddonResponse) {
    response = &UpdateAddonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAddon
// This API is used to update parameters and version of an add-on.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAddon(request *UpdateAddonRequest) (response *UpdateAddonResponse, err error) {
    return c.UpdateAddonWithContext(context.Background(), request)
}

// UpdateAddon
// This API is used to update parameters and version of an add-on.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAddonWithContext(ctx context.Context, request *UpdateAddonRequest) (response *UpdateAddonResponse, err error) {
    if request == nil {
        request = NewUpdateAddonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAddon require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAddonResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateClusterKubeconfigRequest() (request *UpdateClusterKubeconfigRequest) {
    request = &UpdateClusterKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterKubeconfig")
    
    
    return
}

func NewUpdateClusterKubeconfigResponse() (response *UpdateClusterKubeconfigResponse) {
    response = &UpdateClusterKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateClusterKubeconfig
// This API is used to update the Kubeconfig information of a cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBERNETESPATCHOPERATIONERROR = "InternalError.KubernetesPatchOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateClusterKubeconfig(request *UpdateClusterKubeconfigRequest) (response *UpdateClusterKubeconfigResponse, err error) {
    return c.UpdateClusterKubeconfigWithContext(context.Background(), request)
}

// UpdateClusterKubeconfig
// This API is used to update the Kubeconfig information of a cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_KUBERNETESPATCHOPERATIONERROR = "InternalError.KubernetesPatchOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateClusterKubeconfigWithContext(ctx context.Context, request *UpdateClusterKubeconfigRequest) (response *UpdateClusterKubeconfigResponse, err error) {
    if request == nil {
        request = NewUpdateClusterKubeconfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateClusterKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateClusterKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateClusterVersionRequest() (request *UpdateClusterVersionRequest) {
    request = &UpdateClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateClusterVersion")
    
    
    return
}

func NewUpdateClusterVersionResponse() (response *UpdateClusterVersionResponse) {
    response = &UpdateClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateClusterVersion
// This API is used to upgrade the master component of the cluster to the specified version.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateClusterVersion(request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
    return c.UpdateClusterVersionWithContext(context.Background(), request)
}

// UpdateClusterVersion
// This API is used to upgrade the master component of the cluster to the specified version.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateClusterVersionWithContext(ctx context.Context, request *UpdateClusterVersionRequest) (response *UpdateClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpdateClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEdgeClusterVersionRequest() (request *UpdateEdgeClusterVersionRequest) {
    request = &UpdateEdgeClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpdateEdgeClusterVersion")
    
    
    return
}

func NewUpdateEdgeClusterVersionResponse() (response *UpdateEdgeClusterVersionResponse) {
    response = &UpdateEdgeClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEdgeClusterVersion
// This API is used to upgrade an edge cluster component to a TKE Edge version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateEdgeClusterVersion(request *UpdateEdgeClusterVersionRequest) (response *UpdateEdgeClusterVersionResponse, err error) {
    return c.UpdateEdgeClusterVersionWithContext(context.Background(), request)
}

// UpdateEdgeClusterVersion
// This API is used to upgrade an edge cluster component to a TKE Edge version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERUPGRADENODEVERSION = "FailedOperation.ClusterUpgradeNodeVersion"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLUSTERUPGRADENODEVERSION = "InternalError.ClusterUpgradeNodeVersion"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpdateEdgeClusterVersionWithContext(ctx context.Context, request *UpdateEdgeClusterVersionRequest) (response *UpdateEdgeClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpdateEdgeClusterVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEdgeClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEdgeClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterInstancesRequest() (request *UpgradeClusterInstancesRequest) {
    request = &UpgradeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "UpgradeClusterInstances")
    
    
    return
}

func NewUpgradeClusterInstancesResponse() (response *UpgradeClusterInstancesResponse) {
    response = &UpgradeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeClusterInstances
// This API is used to upgrade work nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKALREADYRUNNING = "FailedOperation.TaskAlreadyRunning"
//  FAILEDOPERATION_TASKLIFESTATEERROR = "FailedOperation.TaskLifeStateError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKALREADYRUNNING = "InternalError.TaskAlreadyRunning"
//  INTERNALERROR_TASKLIFESTATEERROR = "InternalError.TaskLifeStateError"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpgradeClusterInstances(request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
    return c.UpgradeClusterInstancesWithContext(context.Background(), request)
}

// UpgradeClusterInstances
// This API is used to upgrade work nodes in a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKALREADYRUNNING = "FailedOperation.TaskAlreadyRunning"
//  FAILEDOPERATION_TASKLIFESTATEERROR = "FailedOperation.TaskLifeStateError"
//  INTERNALERROR_CLUSTERNOTFOUND = "InternalError.ClusterNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKALREADYRUNNING = "InternalError.TaskAlreadyRunning"
//  INTERNALERROR_TASKLIFESTATEERROR = "InternalError.TaskLifeStateError"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) UpgradeClusterInstancesWithContext(ctx context.Context, request *UpgradeClusterInstancesRequest) (response *UpgradeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}
