// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/merge-api/merge-go-client/core"
	accountdetails "github.com/merge-api/merge-go-client/crm/accountdetails"
	accounts "github.com/merge-api/merge-go-client/crm/accounts"
	accounttoken "github.com/merge-api/merge-go-client/crm/accounttoken"
	associations "github.com/merge-api/merge-go-client/crm/associations"
	associationtypes "github.com/merge-api/merge-go-client/crm/associationtypes"
	asyncpassthrough "github.com/merge-api/merge-go-client/crm/asyncpassthrough"
	audittrail "github.com/merge-api/merge-go-client/crm/audittrail"
	availableactions "github.com/merge-api/merge-go-client/crm/availableactions"
	contacts "github.com/merge-api/merge-go-client/crm/contacts"
	customobjectclasses "github.com/merge-api/merge-go-client/crm/customobjectclasses"
	customobjects "github.com/merge-api/merge-go-client/crm/customobjects"
	deleteaccount "github.com/merge-api/merge-go-client/crm/deleteaccount"
	engagements "github.com/merge-api/merge-go-client/crm/engagements"
	engagementtypes "github.com/merge-api/merge-go-client/crm/engagementtypes"
	fieldmapping "github.com/merge-api/merge-go-client/crm/fieldmapping"
	forceresync "github.com/merge-api/merge-go-client/crm/forceresync"
	generatekey "github.com/merge-api/merge-go-client/crm/generatekey"
	issues "github.com/merge-api/merge-go-client/crm/issues"
	leads "github.com/merge-api/merge-go-client/crm/leads"
	linkedaccounts "github.com/merge-api/merge-go-client/crm/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/crm/linktoken"
	notes "github.com/merge-api/merge-go-client/crm/notes"
	opportunities "github.com/merge-api/merge-go-client/crm/opportunities"
	passthrough "github.com/merge-api/merge-go-client/crm/passthrough"
	regeneratekey "github.com/merge-api/merge-go-client/crm/regeneratekey"
	scopes "github.com/merge-api/merge-go-client/crm/scopes"
	stages "github.com/merge-api/merge-go-client/crm/stages"
	syncstatus "github.com/merge-api/merge-go-client/crm/syncstatus"
	tasks "github.com/merge-api/merge-go-client/crm/tasks"
	users "github.com/merge-api/merge-go-client/crm/users"
	webhookreceivers "github.com/merge-api/merge-go-client/crm/webhookreceivers"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	AccountDetails      *accountdetails.Client
	AccountToken        *accounttoken.Client
	Accounts            *accounts.Client
	AsyncPassthrough    *asyncpassthrough.Client
	AuditTrail          *audittrail.Client
	AvailableActions    *availableactions.Client
	Contacts            *contacts.Client
	CustomObjectClasses *customobjectclasses.Client
	AssociationTypes    *associationtypes.Client
	CustomObjects       *customobjects.Client
	Associations        *associations.Client
	Scopes              *scopes.Client
	DeleteAccount       *deleteaccount.Client
	EngagementTypes     *engagementtypes.Client
	Engagements         *engagements.Client
	FieldMapping        *fieldmapping.Client
	GenerateKey         *generatekey.Client
	Issues              *issues.Client
	Leads               *leads.Client
	LinkToken           *linktoken.Client
	LinkedAccounts      *linkedaccounts.Client
	Notes               *notes.Client
	Opportunities       *opportunities.Client
	Passthrough         *passthrough.Client
	RegenerateKey       *regeneratekey.Client
	Stages              *stages.Client
	SyncStatus          *syncstatus.Client
	ForceResync         *forceresync.Client
	Tasks               *tasks.Client
	Users               *users.Client
	WebhookReceivers    *webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:             options.BaseURL,
		caller:              core.NewCaller(options.HTTPClient),
		header:              options.ToHeader(),
		AccountDetails:      accountdetails.NewClient(opts...),
		AccountToken:        accounttoken.NewClient(opts...),
		Accounts:            accounts.NewClient(opts...),
		AsyncPassthrough:    asyncpassthrough.NewClient(opts...),
		AuditTrail:          audittrail.NewClient(opts...),
		AvailableActions:    availableactions.NewClient(opts...),
		Contacts:            contacts.NewClient(opts...),
		CustomObjectClasses: customobjectclasses.NewClient(opts...),
		AssociationTypes:    associationtypes.NewClient(opts...),
		CustomObjects:       customobjects.NewClient(opts...),
		Associations:        associations.NewClient(opts...),
		Scopes:              scopes.NewClient(opts...),
		DeleteAccount:       deleteaccount.NewClient(opts...),
		EngagementTypes:     engagementtypes.NewClient(opts...),
		Engagements:         engagements.NewClient(opts...),
		FieldMapping:        fieldmapping.NewClient(opts...),
		GenerateKey:         generatekey.NewClient(opts...),
		Issues:              issues.NewClient(opts...),
		Leads:               leads.NewClient(opts...),
		LinkToken:           linktoken.NewClient(opts...),
		LinkedAccounts:      linkedaccounts.NewClient(opts...),
		Notes:               notes.NewClient(opts...),
		Opportunities:       opportunities.NewClient(opts...),
		Passthrough:         passthrough.NewClient(opts...),
		RegenerateKey:       regeneratekey.NewClient(opts...),
		Stages:              stages.NewClient(opts...),
		SyncStatus:          syncstatus.NewClient(opts...),
		ForceResync:         forceresync.NewClient(opts...),
		Tasks:               tasks.NewClient(opts...),
		Users:               users.NewClient(opts...),
		WebhookReceivers:    webhookreceivers.NewClient(opts...),
	}
}
