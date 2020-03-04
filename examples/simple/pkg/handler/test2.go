package handler

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/openbox-go/examples/simple/internal/pkg/model/response"
	"github.com/b2wdigital/openbox-go/examples/simple/pkg/event/request"
	cloudevents "github.com/cloudevents/sdk-go"
)

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	log := logrus.FromContext(ctx)

	user := &request.User{}
	if err := e.DataAs(user); err != nil {
		log.Printf("Got Data Error: %s\n", err.Error())
	}

	log.Info(user.Name)

	r := cloudevents.Event{
		Context: cloudevents.EventContextV1{
			Source: *cloudevents.ParseURIRef("/mod3"),
			Type:   "samples.http.mod3",
		}.AsV1(),
		Data: response.Default{
			Message: "Test 3!!",
		},
	}

	resp.Event = &r

	return nil
}
