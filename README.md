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

**Use default configurations unless necessary.**

Modify `profile.ClientProfile` fields before client creation for customization.

```go
// Optional
cpf := profile.NewClientProfile()
```

Configuration options:

## Request Method

SDK uses POST by default. Switch to GET only if necessary (**not for large requests**).

```go
cpf.HttpProfile.ReqMethod = "POST"
```

## Timeout

Default timeout (seconds). Check latest defaults in code.

```go
cpf.HttpProfile.ReqTimeout = 30
```

## Custom Endpoint

Auto-set by SDK. Required for financial zone services.

```go
cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
```

## Signature Method

Default: `TC3-HMAC-SHA256` (secure but slightly slower).

```go
cpf.SignMethod = "HmacSHA1"
```

## Debug Mode

Enable for detailed logs during troubleshooting.

```go
cpf.Debug = true
```

## Disable Keep-Alive

Each client uses keep-alive by default. For short-lived connections:

```go
...
    client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
    tp := &http.Transport{
        DisableKeepAlives: true,
    }
    client.WithHttpTransport(tp)
...
```

## Region Failover

Tencent Cloud GO SDK supports region failover:

When:
1. Failures >= 5
2. Failure rate >= 75%

SDK automatically switches to backup region.

Configuration:
```golang
    // Enable
    cpf.DisableRegionBreaker = false
    // Set backup endpoint (service prefix auto-added)
    cpf.BackupEndpoint = "ap-guangzhou.tencentcloudapi.com"
```

Note: Only supports synchronous requests per client.

## Proxy

Set `https_proxy` environment variable or customize Transport with `client.WithHttpTransport()`.

## Enable DNS Cache

Export `GODEBUG=netdns=cgo` or build with `-tags 'netcgo'` to use nscd cache.

## Skip Server Certificate Verification

Although mandatory for security, you might need to skip verification during testing:

```golang
import "crypto/tls"
...
    client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client.WithHttpTransport(tr)
...
```

**Warning: Understand risks before disabling certificate verification.**

# Features

## Credential Management

1. Environment Variables
   ```go
   provider := common.DefaultEnvProvider()
   credential, err := provider.GetCredential()
   ```

2. Config File
   Path priority:
    1. `TENCENTCLOUD_CREDENTIALS_FILE`
    2. Linux/Mac: `~/.tencentcloud/credentials`
    3. Windows: `c:\Users\NAME\.tencentcloud\credentials`

   Format:
   ```ini
   [default]
   secret_id = xxxxx
   secret_key = xxxxx
   ```

   ```go
   provider := common.DefaultProfileProvider()
   credential, err := provider.GetCredential()
   ```

3. Role Assumption
   ```go
   provider := common.DefaultRoleArnProvider(secretId, secretKey, roleArn)
   credential, err := provider.GetCredential()
   ```

4. Instance Role
   ```go
   provider := common.DefaultCvmRoleProvider()
   credential, err := provider.GetCredential()
   ```

5. TKE OIDC Credentials
   Example: [examples/ssm/v20190923/get_secret_value.go](examples/ssm/v20190923/get_secret_value.go)
   ```go
   provider, err := common.DefaultTkeOIDCRoleArnProvider()
   credential, err := provider.GetCredential()
   ```

6. Credential Provider Chain
   Default chain: `Env -> Config File -> Instance Role`
   ```go
   provider := common.DefaultProviderChain()
   credential, err := provider.GetCredential()
   ```

   Custom chain:
   ```go
   provider1 := common.DefaultCvmRoleProvider()
   provider2 := common.DefaultEnvProvider()
   customProviderChain := []common.Provider{provider1, provider2}
   provider := common.NewProviderChain(customProviderChain)
   credential, err := provider.GetCredential()
   ```

   Full example: [provider_chain_test.go](https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/testing/integration/provider_chain_test.go)

## Error Handling

Error codes are defined as constants (IDE-friendly):

```go
response, err := client.DescribeInstances(request)
if terr, ok := err.(*errors.TencentCloudSDKError); ok {
    code := terr.GetCode()
    if code == cvm.FAILEDOPERATION_ILLEGALTAGKEY{
        fmt.Printf("Handling error: FailedOperation.IllegalTagKey,%s", err)
    }else if code == cvm.UNAUTHORIZEDOPERATION{
        fmt.Printf("Handling error: UnauthorizedOperation,%s", err)
    }else{
        fmt.Printf("An API error has returned: %s", err)
    }
    return
}
```

API comments list possible error codes:
```go
// error code that may be returned:
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
// ...
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error){
    ...
}
```

## Common Client

Generic API calls are supported via `common` package.

**Note: You must know exact request parameters.**

Only POST with v3 signature is supported.

Example: [common_client.go](https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/examples/common/common_client.go)

## Custom Headers

[RunInstancesRequest Example](examples/cvm/v20170312/run_instances.go)
```go
    request.SetHeader(map[string]string{
        "X-TC-TraceId": "ffe0c072-8a5d-4e17-8887-a8a60252abca",
    })
```

[CommonRequest Example](examples/common/common_client.go)
```go
    request.SetHeader(map[string]string{
        "X-TC-TraceId": "ffe0c072-8a5d-4e17-8887-a8a60252abca",
    })
```

## HTTP Proxy
[DescribeInstances Example](examples/cvm/v20170312/describe_instances.go)
```go
    // Authenticated
    clientProfile.HttpProfile.Proxy = "http://username:password@127.0.0.1:1080"
    // Unauthenticated
    clientProfile.HttpProfile.Proxy = "http://127.0.0.1:1080"
```

## Request Retry

### Network Error Retry

Configure via `ClientProfile` for automatic retries on temporary failures (disabled by default).

> Only idempotent requests (with `ClientToken` field) are retried.

```golang
prof := profile.NewClientProfile()
prof.NetworkFailureMaxRetries = 3
prof.NetworkFailureRetryDuration = profile.ExponentialBackoff
```

More: [netretry_test.go](https://github.com/TencentCloud/tencentcloud-sdk-go/tree/master/tencentcloud/common/netretry_test.go)

### Rate Limit Retry

Configure for automatic retries on rate limits (disabled by default).

```golang
prof := profile.NewClientProfile()
prof.RateLimitExceededMaxRetries = 3
prof.RateLimitExceededRetryDuration = profile.ExponentialBackoff
```

### Idempotency Token

When retries are enabled, `ClientToken` is auto-injected if empty (skipped if manually set).

> Injected tokens guarantee uniqueness below 100,000 concurrent requests.

## Empty Arrays and omitempty

`omitempty` tag prevented serialization of both nil and empty arrays.

`omitnil` tag allows empty arrays (nil arrays are ignored).

Toggle via `json.OmitBehaviour = json.OmitEmpty` to disable this feature.  
Example: [omitempty.go](https://github.com/TencentCloud/tencentcloud-sdk-go-intl-en/blob/master/examples/common/omitempty.go)

## EOF Errors

If encountering `Code=ClientError.NetworkError, Message=Fail to get response because Post "https://xxx.tencentcloudapi.com/": EOF`, it may relate to keep-alive. Solutions:

- Switch to GET if request payloads are small;
- Enable `ClientProfile.UnsafeRetryOnConnectionFailure` (ensure idempotency);
