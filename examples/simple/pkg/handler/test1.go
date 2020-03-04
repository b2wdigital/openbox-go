package handler

import (
	"context"

	"github.com/b2wdigital/openbox-go/examples/simple/internal/pkg/model/response"
	"github.com/cloudevents/sdk-go"
)

func Test1(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	r := cloudevents.Event{
		Context: cloudevents.EventContextV1{
			Source: *cloudevents.ParseURIRef("/mod3"),
			Type:   "samples.http.mod3",
		}.AsV1(),
		Data: response.Default{
			Message: "Test 1!!",
		},
	}

	resp.Event = &r

	return nil
}
