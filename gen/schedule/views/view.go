// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// schedule views
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa"
)

// SchedulePayloadCollection is the viewed result type that is projected based
// on a view.
type SchedulePayloadCollection struct {
	// Type to project
	Projected SchedulePayloadCollectionView
	// View to render
	View string
}

// SchedulePayloadCollectionView is a type that runs validations on a projected
// type.
type SchedulePayloadCollectionView []*SchedulePayloadView

// SchedulePayloadView is a type that runs validations on a projected type.
type SchedulePayloadView struct {
	// ID is the unique id of the schedule.
	ID *string
	// Descriptive Name
	Name *string
	// Valid cron string
	Cron *string
	// color to set
	Color *string
}

// Validate runs the validations defined on the viewed result type
// SchedulePayloadCollection.
func (result SchedulePayloadCollection) Validate() (err error) {
	switch result.View {
	case "default", "":
		err = result.Projected.Validate()
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// Validate runs the validations defined on SchedulePayloadCollectionView using
// the "default" view.
func (result SchedulePayloadCollectionView) Validate() (err error) {
	for _, item := range result {
		if err2 := item.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Validate runs the validations defined on SchedulePayloadView using the
// "default" view.
func (result *SchedulePayloadView) Validate() (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Cron == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cron", "result"))
	}
	if result.Color == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("color", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	if result.Color != nil {
		if !(*result.Color == "red" || *result.Color == "yellow" || *result.Color == "green" || *result.Color == "off") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.color", *result.Color, []interface{}{"red", "yellow", "green", "off"}))
		}
	}
	return
}
