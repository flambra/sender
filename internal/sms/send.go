package sms

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/sender/internal/domain"
	"github.com/gofiber/fiber/v2"

	smsTemplate "github.com/flambra/sender/internal/sms/template"
)

func Send(c *fiber.Ctx) error {
	var request domain.SmsRequest
	var template domain.SMSTemplate
	repo := hRepository.New(hDb.Get(), &template, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"name": request.TemplateName})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Template not found")
	}

	message, err := smsTemplate.Process(template, request.RecipientName)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "Failed to process template")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return err
	}

	client := sns.NewFromConfig(cfg)

	input := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(request.To),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String("MySenderID"),
			},
		},
	}

	_, err = client.Publish(context.TODO(), input)
	if err != nil {
		return err
	}

	return hResp.SuccessResponse(c, "SMS sent successfully")
}
