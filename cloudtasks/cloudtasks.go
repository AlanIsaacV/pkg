package cloudtasks

import (
	"context"
	"sync"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rs/zerolog/log"

	"app/pkg/gcp"
)

var (
	client     *cloudtasks.Client
	clientOnce sync.Once
	gcpConfig  = gcp.Config()
)

func Client() *cloudtasks.Client {
	clientOnce.Do(
		func() {
			var err error

			client, err = cloudtasks.NewClient(context.Background())
			if err != nil {
				log.Fatal().Err(err).Msg("Error creating cloud tasks client")
			}
		},
	)
	return client
}

func DefaultRequest() *taskspb.CreateTaskRequest {
	c := Config()

	delay := time.Now().Unix() + c.Delay
	return &taskspb.CreateTaskRequest{
		Parent: c.Queue,
		Task: &taskspb.Task{
			ScheduleTime: &timestamppb.Timestamp{Seconds: delay},
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        Config().Url,
				},
			},
		},
	}
}

func GetHttpRequest(request *taskspb.CreateTaskRequest) *taskspb.HttpRequest {
	return request.GetTask().GetMessageType().(*taskspb.Task_HttpRequest).HttpRequest
}

func AddToken(request *taskspb.CreateTaskRequest) *taskspb.CreateTaskRequest {
	GetHttpRequest(request).AuthorizationHeader = &taskspb.HttpRequest_OidcToken{
		OidcToken: &taskspb.OidcToken{ServiceAccountEmail: gcpConfig.Email},
	}

	return request
}
