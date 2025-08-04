package main

import (
	"log"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/cvm/v20170312"
)

// This function is expected to fail for the first 5 times and succeed for the next 5 times.
func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 2
	// An incorrect domain will cause the request to fail, but it's expected to succeed after the 5th attempt.
	cpf.HttpProfile.Endpoint = "cvm.ap-shanghai1.tencentcloudapi.com"
	cpf.BackupEndPoint = "ap-beijing.tencentcloudapi.com"

	// Enable the circuit breaker
	cpf.DisableRegionBreaker = false
	client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
	for i := 0; i < 10; i++ {
		request := cvm.NewDescribeZonesRequest()
		response, err := client.DescribeZones(request)
		log.Println(i+1, "/", 10)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			log.Printf("An API error has returned: %s", err)
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("%s\n", response.ToJsonString())
	}
}
