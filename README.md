# Overview

Welcome to Tencent Cloud Software Development Kit (SDK), a companion tool for the TencentCloud API 3.0 platform.

Tencent Cloud SDK for Go helps Go developers debug and use TencentCloud APIs with ease. This document describes Tencent Cloud SDK for Go and how to quickly use it with code examples provided.

# Dependent Environment

1. Go 1.9 or higher. And, the necessary environment variables such as `GOPATH` have to be set properly.
2. Before using, make sure to activate the corresponding product in the Tencent Cloud Console.
3. Get the `SecretID` and `SecretKey` on the [CAM](https://console.cloud.tencent.com/cam/capi) page in the Tencent Cloud Console.

# Installation

Before installing Tencent Cloud SDK for Go and using TencentCloud API, apply for security credentials in the Tencent Cloud Console. Security credential consists of `SecretID` and `SecretKey`. `SecretID` is for identifying the API requester. `SecretKey` is a key used for signature string encryption and authentication by the server. Please keep your `SecretKey` private and do not disclose it to others.

## Installing via go get (recommended)

You are recommended to install the SDK by using the tool that comes with the language:

    go get -u github.com/tencentcloud/tencentcloud-sdk-go-intl-en


## Installing via source code

Go to the [Github code hosting page](https://github.com/tencentcloud/tencentcloud-sdk-go-intl-en) to download the latest code, decompress, and install it in the `$GOPATH/src/github.com/tencentcloud` directory.

# Example

Each API has a corresponding request structure and a response structure. For example, the `DescribeInstances` API for querying CVM instance list has a corresponding request structure `DescribeInstancesRequest` and a response structure `DescribeInstancesResponse`.

The following uses the API for querying CVM instance list as an example to describe the basic usage of the SDK. For the purpose of demonstration, some nonessential items have been added to show the common functions of the SDK, which makes the example look bloated. When using the SDK to write code, you are recommended to keep it simple.

```
package main

import (
        "fmt"
        "os"

        "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
        "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
        cvm "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/cvm/v20170312"
)

func main() {
        // Essential steps:
        // Instantiate an authentication object. The Tencent Cloud account key pair `secretId` and `secretKey` need to be passed in as the input parameters.
        // The example here uses the way to read from the environment variable, so you need to set these two values in the environment variable first.
        // You can also write the key pair directly into the code, but be careful not to copy, upload, or share the code to others;
        // otherwise, the key pair may be leaked, causing damage to your properties.
        credential := common.NewCredential(
                os.Getenv("TENCENTCLOUD_SECRET_ID"),
                os.Getenv("TENCENTCLOUD_SECRET_KEY"),
        )

        // Nonessential steps
        // Instantiate a client configuration object. You can specify the timeout period and other configuration items
        cpf := profile.NewClientProfile()
        // The SDK uses the POST method by default.
        // If you have to use the GET method, you can set it here, but the GET method cannot handle some large requests.
        // Do not modify the default settings unless absolutely necessary.
        //cpf.HttpProfile.ReqMethod = "GET"
        // The SDK has a default timeout period. Do not adjust it unless absolutely necessary.
        // If needed, check in the code to get the latest default value.
        //cpf.HttpProfile.ReqTimeout = 10
        // The SDK automatically specifies the domain name. Generally, you don't need to specify a domain name, but if you are accessing a service in a finance availability zone,
        // you have to manually specify the domain name, such as cvm.ap-shanghai-fsi.tencentcloudapi.com for the Shanghai Finance availability zone in the CVM
        //cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
        // The SDK uses HmacSHA256 for signing by default, which is more secure but slightly reduces the performance.
        // Do not modify the default settings unless absolutely necessary.
        //cpf.SignMethod = "HmacSHA1"
        // The SDK uses `zh-CN` calls to return Chinese content by default. You can also set it to `en-US` to return English content.
        // However, most products or APIs do not fully support returns in English.
        // Do not modify the default settings unless absolutely necessary.
        //cpf.Language = "en-US"

        // Instantiate the client object of the requested product (with CVM as an example)
        // The second parameter is the region information. You can enter the string `ap-guangzhou` directly or import the preset constant
        client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
        // Instantiate a request object. You can further set the request parameters according to the API called and actual conditions
        // You can check the SDK source code directly to determine which attributes of `DescribeInstancesRequest` can be set.
        // An attribute may be of a basic type or import another data structure.
        // You are recommended to use the IDE for development where you can redirect to and view the documentation of each API and data structure easily
        request := cvm.NewDescribeInstancesRequest()

        // Settings of a basic parameter.
        // This API allows setting the number of instances returned, which is specified as only one here.
        // The SDK uses the pointer style to specify parameters, so even for basic parameters, you need to use pointers to assign values to them.
        // The SDK provides encapsulation functions for importing the pointers of basic parameters
        request.Limit = common.Int64Ptr(1)

        // Settings of an array.
        // This API allows filtering by specified instance ID; however, as it conflicts with the `Filter` parameter to be demonstrated next, it is commented out here.
        // request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

        // Settings of a complex object.
        // In this API, `Filters` is an array whose elements are complex objects `Filter`, and the member `Values` of `Filter` are string arrays.
        request.Filters = []*cvm.Filter{
            &cvm.Filter{
                Name: common.StringPtr("zone"),
                Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
            },
        }

        // Use a JSON string to set a request. Note that this is actually an update request, that is, `Limit=1` will be retained,
        // and the `zone` of the filter will be changed to `ap-guangzhou-2`.
        // If you need a new request, you should create it with `cvm.NewDescribeInstancesRequest()`.
        err := request.FromJsonString(`{"Filters":[{"Name":"zone","Values":["ap-guangzhou-2"]}]}`)
        if err != nil {
                panic(err)
        }
        // Call the API you want to access through the client object. You need to pass in the request object
        response, err := client.DescribeInstances(request)
        // Handle the exception
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        // This is a direct failure instead of SDK exception. You can add other troubleshooting measures in the real code.
        if err != nil {
                panic(err)
        }
        // Print the returned JSON string
        fmt.Printf("%s", response.ToJsonString())
}
```

For more examples, please see the [examples](https://github.com/TencentCloud/tencentcloud-sdk-go-intl-en/tree/master/examples) directory. For an example of initializing a request for a complicated API, please see `examples/cvm/v20170312/run_instances.go`. For an example of initializing a request by using a JSON string, please see `examples/cvm/v20170312/describe_instances.go`.

# Relevant Configuration

## Proxy

If there is a proxy in your environment, you need to set the system environment variable `https_proxy`; otherwise, it may not be called normally, and a connection timeout exception will be thrown.

## Enabling DNS cache

Currently, the SDK for Go always requests the DNS server without using the cache of nscd. You can export the environment variable `GODEBUG=netdns=cgo` or specify the `-tags 'netcgo'` parameter when compiling `go build` so as to get the nscd cache.

## Ignoring server certificate verification

When the SDK is used to call a public cloud service, the server certificate must be verified to identify forged servers and ensure request security.
However, in some extreme cases such as testing, you may need to ignore self-signed server certificates.
Here is one of the possible methods:

```
import "crypto/tls"
...
    client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client.WithHttpTransport(tr)
...
```

Again, unless you know what you are doing and understand the risks involved, do not try to disable server certificate verification.
