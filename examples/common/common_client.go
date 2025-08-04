package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
	"log"
	"os"
)

// Currently only supports Signature V3 + POST
func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	// Create a common client
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf).WithLogger(log.Default())

	// Create a common request, passing in the product name, product version, and API name in order
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")

	// Custom request parameters:
	// The SetActionParameters function supports three data types for input parameters:
	// 1. map[string]interface{}
	// body:=map[string]interface{}{
	// "InstanceId":"crs-xxx",
	// "SpanType":2,
	// }
	body := map[string]interface{}{}

	// // 2. string
	// bodyStr := `{}`

	// // 3. []byte
	// bodyBytes := []byte(bodyStr)

	// Set custom headers
	request.SetHeader(map[string]string{
		"X-TC-TraceId": "ffe0c072-8a5d-4e17-8887-a8a60252abca",
	})

	// Set the request data required by the action
	err := request.SetActionParameters(body)
	if err != nil {
		return
	}

	// Create a common response
	response := tchttp.NewCommonResponse()

	// Send the request
	err = client.Send(request, response)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}

	// Get the response result
	fmt.Println(string(response.GetBody()))
}
