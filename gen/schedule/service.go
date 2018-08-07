// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// schedule service
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package schedule

import (
	"context"

	scheduleviews "github.com/jaredwarren/rg/gen/schedule/views"
)

// The Alarm schedule service.
type Service interface {
	// Alarm Schedule Home
	Home(context.Context) (err error)
	// List all stored bottles
	List(context.Context) (res SchedulePayloadCollection, err error)
	// create new cron schedule
	Schedule(context.Context, *SchedulePayload) (res *Schedule, err error)
	// Remove cron schedule
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "schedule"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"home", "list", "schedule", "remove"}

// SchedulePayloadCollection is the result type of the schedule service list
// method.
type SchedulePayloadCollection []*SchedulePayload

// SchedulePayload is the payload type of the schedule service schedule method.
type SchedulePayload struct {
	// ID is the unique id of the schedule.
	ID string
	// Descriptive Name
	Name string
	// Valid cron string
	Cron string
	// color to set
	Color string
}

// Schedule is the result type of the schedule service schedule method.
type Schedule struct {
	// Descriptive Name
	Name *string
	// Valid cron string
	Cron string
	// color to set
	Color string
}

// RemovePayload is the payload type of the schedule service remove method.
type RemovePayload struct {
	// ID of bottle to remove
	ID string
}

// NotFound is the type returned when attempting to show or delete a bottle
// that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing bottle
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a bottle that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewSchedulePayloadCollection initializes result type
// SchedulePayloadCollection from viewed result type SchedulePayloadCollection.
func NewSchedulePayloadCollection(vres scheduleviews.SchedulePayloadCollection) SchedulePayloadCollection {
	var res SchedulePayloadCollection
	switch vres.View {
	case "default", "":
		res = newSchedulePayloadCollection(vres.Projected)
	}
	return res
}

// NewViewedSchedulePayloadCollection initializes viewed result type
// SchedulePayloadCollection from result type SchedulePayloadCollection using
// the given view.
func NewViewedSchedulePayloadCollection(res SchedulePayloadCollection, view string) scheduleviews.SchedulePayloadCollection {
	var vres scheduleviews.SchedulePayloadCollection
	switch view {
	case "default", "":
		p := newSchedulePayloadCollectionView(res)
		vres = scheduleviews.SchedulePayloadCollection{p, "default"}
	}
	return vres
}

// newSchedulePayloadCollection converts projected type
// SchedulePayloadCollection to service type SchedulePayloadCollection.
func newSchedulePayloadCollection(vres scheduleviews.SchedulePayloadCollectionView) SchedulePayloadCollection {
	res := make(SchedulePayloadCollection, len(vres))
	for i, n := range vres {
		res[i] = newSchedulePayload(n)
	}
	return res
}

// newSchedulePayloadCollectionView projects result type
// SchedulePayloadCollection into projected type SchedulePayloadCollectionView
// using the "default" view.
func newSchedulePayloadCollectionView(res SchedulePayloadCollection) scheduleviews.SchedulePayloadCollectionView {
	vres := make(scheduleviews.SchedulePayloadCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSchedulePayloadView(n)
	}
	return vres
}

// newSchedulePayload converts projected type SchedulePayload to service type
// SchedulePayload.
func newSchedulePayload(vres *scheduleviews.SchedulePayloadView) *SchedulePayload {
	res := &SchedulePayload{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Cron != nil {
		res.Cron = *vres.Cron
	}
	if vres.Color != nil {
		res.Color = *vres.Color
	}
	return res
}

// newSchedulePayloadView projects result type SchedulePayload into projected
// type SchedulePayloadView using the "default" view.
func newSchedulePayloadView(res *SchedulePayload) *scheduleviews.SchedulePayloadView {
	vres := &scheduleviews.SchedulePayloadView{
		ID:    &res.ID,
		Name:  &res.Name,
		Cron:  &res.Cron,
		Color: &res.Color,
	}
	return vres
}
