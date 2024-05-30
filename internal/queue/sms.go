package queue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/flambra/sender/internal/domain"
)

var smsQueue = make(chan domain.SMSTask, 100)

func init() {
	go processSMSTask()
	log.Println("SMS Queue initialized")
}

func EnqueueSMSTask(request domain.SMSRequest, template domain.SMSTemplate, message string) {
	smsQueue <- domain.SMSTask{
		Request:  request,
		Template: template,
		Message:  message,
	}
}

func processSMSTask() {
	for task := range smsQueue {
		sendSMS(task)
	}
}

func sendSMS(task domain.SMSTask) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	client := sns.NewFromConfig(cfg)

	input := &sns.PublishInput{
		Message:     aws.String(task.Message),
		PhoneNumber: aws.String(task.Request.To),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String("MySenderID"),
			},
		},
	}

	_, err = client.Publish(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to send SMS to: %s with error: %v", task.Request.To, err)
	}
}
