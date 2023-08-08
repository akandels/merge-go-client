// This file was auto-generated by Fern from our API Definition.

package ats

import (
	time "time"
)

type UpdateApplicationStageRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool `json:"-"`
	// The interview stage to move the application to.
	JobInterviewStage *string `json:"job_interview_stage,omitempty"`
	// <span style="white-space: nowrap">`non-empty`</span>
	RemoteUserId *string `json:"remote_user_id,omitempty"`
}

type ApplicationEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync     *bool               `json:"-"`
	Model        *ApplicationRequest `json:"model,omitempty"`
	RemoteUserId string              `json:"remote_user_id"`
}

type ApplicationsListRequest struct {
	// If provided, will only return applications for this candidate.
	CandidateId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// If provided, will only return applications credited to this user.
	CreditedToId *string `json:"-"`
	// If provided, will only return applications at this interview stage.
	CurrentStageId *string `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ApplicationsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, will only return applications for this job.
	JobId *string `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// If provided, will only return applications with this reject reason.
	RejectReasonId *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return applications with this source.
	Source *string `json:"-"`
}

type ApplicationsMetaPostRetrieveRequest struct {
	// The template ID associated with the nested application in the request.
	ApplicationRemoteTemplateId *string `json:"-"`
}

type ApplicationsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ApplicationsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
