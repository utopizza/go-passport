package thirdparty

import (
	"log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
)

var (
	accessKeyId     = ""
	accessKeySecret = ""
	domain          = ""
)

type SmsService struct {
	cli *dysmsapi.Client
}

func NewSmsService() (*SmsService, error) {
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &domain,
	}
	cli, err := dysmsapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &SmsService{cli: cli}, nil
}

func (s *SmsService) SendSms(mobile string) error {
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers: &mobile,
		SignName:     nil,
		TemplateCode: nil,
	}

	response, err := s.cli.SendSms(request)
	if err != nil {
		return err
	}
	log.Printf("[SendSms] request:%+v, response:%+v", request, response)

	return nil
}
