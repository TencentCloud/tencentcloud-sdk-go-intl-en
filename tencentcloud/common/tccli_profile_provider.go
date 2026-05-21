package common

import (
	"os"
	"path/filepath"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

const (
	EnvTencentCloudProfile = "TENCENTCLOUD_PROFILE"
)

type TccliProfileProvider struct {
	profile string
}

func DefaultTccliProfileProvider() *TccliProfileProvider {
	profile, ok := os.LookupEnv(EnvTencentCloudProfile)
	if !ok {
		profile = "default"
	}

	return &TccliProfileProvider{profile: profile}
}

func (p *TccliProfileProvider) checkCredentialFile() (string, error) {
	path := filepath.Join(getHomePath(), ".tccli", p.profile+".credential")
	if path == "" {
		return path, nil
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return path, nil
}

func (p *TccliProfileProvider) GetCredential() (CredentialIface, error) {
	path, err := p.checkCredentialFile()
	if err != nil {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "Failed to find profile file, "+err.Error(), "")
	}
	if path == "" {
		return nil, fileDoseNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tccliCred := struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
		Token     string `json:"token"`
	}{}
	if err := json.Unmarshal(data, &tccliCred); err != nil {
		return nil, err
	}

	if tccliCred.SecretId == "" || tccliCred.SecretKey == "" {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "Failed to parse profile file,please confirm whether it contains \"secretId\" and \"secretKey\"", "")
	}
	return &Credential{
		SecretId:  tccliCred.SecretId,
		SecretKey: tccliCred.SecretKey,
		Token:     tccliCred.Token,
	}, nil
}
