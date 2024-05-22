// This file was auto-generated by Fern from our API Definition.

package client

import (
	accountdetails "github.com/merge-api/merge-go-client/accounting/accountdetails"
	accountingperiods "github.com/merge-api/merge-go-client/accounting/accountingperiods"
	accounts "github.com/merge-api/merge-go-client/accounting/accounts"
	accounttoken "github.com/merge-api/merge-go-client/accounting/accounttoken"
	addresses "github.com/merge-api/merge-go-client/accounting/addresses"
	asyncpassthrough "github.com/merge-api/merge-go-client/accounting/asyncpassthrough"
	attachments "github.com/merge-api/merge-go-client/accounting/attachments"
	audittrail "github.com/merge-api/merge-go-client/accounting/audittrail"
	availableactions "github.com/merge-api/merge-go-client/accounting/availableactions"
	balancesheets "github.com/merge-api/merge-go-client/accounting/balancesheets"
	cashflowstatements "github.com/merge-api/merge-go-client/accounting/cashflowstatements"
	companyinfo "github.com/merge-api/merge-go-client/accounting/companyinfo"
	contacts "github.com/merge-api/merge-go-client/accounting/contacts"
	creditnotes "github.com/merge-api/merge-go-client/accounting/creditnotes"
	deleteaccount "github.com/merge-api/merge-go-client/accounting/deleteaccount"
	expenses "github.com/merge-api/merge-go-client/accounting/expenses"
	fieldmapping "github.com/merge-api/merge-go-client/accounting/fieldmapping"
	forceresync "github.com/merge-api/merge-go-client/accounting/forceresync"
	generatekey "github.com/merge-api/merge-go-client/accounting/generatekey"
	incomestatements "github.com/merge-api/merge-go-client/accounting/incomestatements"
	invoices "github.com/merge-api/merge-go-client/accounting/invoices"
	issues "github.com/merge-api/merge-go-client/accounting/issues"
	items "github.com/merge-api/merge-go-client/accounting/items"
	journalentries "github.com/merge-api/merge-go-client/accounting/journalentries"
	linkedaccounts "github.com/merge-api/merge-go-client/accounting/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/accounting/linktoken"
	passthrough "github.com/merge-api/merge-go-client/accounting/passthrough"
	payments "github.com/merge-api/merge-go-client/accounting/payments"
	phonenumbers "github.com/merge-api/merge-go-client/accounting/phonenumbers"
	purchaseorders "github.com/merge-api/merge-go-client/accounting/purchaseorders"
	regeneratekey "github.com/merge-api/merge-go-client/accounting/regeneratekey"
	scopes "github.com/merge-api/merge-go-client/accounting/scopes"
	syncstatus "github.com/merge-api/merge-go-client/accounting/syncstatus"
	taxrates "github.com/merge-api/merge-go-client/accounting/taxrates"
	trackingcategories "github.com/merge-api/merge-go-client/accounting/trackingcategories"
	transactions "github.com/merge-api/merge-go-client/accounting/transactions"
	vendorcredits "github.com/merge-api/merge-go-client/accounting/vendorcredits"
	webhookreceivers "github.com/merge-api/merge-go-client/accounting/webhookreceivers"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	AccountDetails     *accountdetails.Client
	AccountToken       *accounttoken.Client
	AccountingPeriods  *accountingperiods.Client
	Accounts           *accounts.Client
	Addresses          *addresses.Client
	AsyncPassthrough   *asyncpassthrough.Client
	Attachments        *attachments.Client
	AuditTrail         *audittrail.Client
	AvailableActions   *availableactions.Client
	BalanceSheets      *balancesheets.Client
	CashFlowStatements *cashflowstatements.Client
	CompanyInfo        *companyinfo.Client
	Contacts           *contacts.Client
	CreditNotes        *creditnotes.Client
	Scopes             *scopes.Client
	DeleteAccount      *deleteaccount.Client
	Expenses           *expenses.Client
	FieldMapping       *fieldmapping.Client
	GenerateKey        *generatekey.Client
	IncomeStatements   *incomestatements.Client
	Invoices           *invoices.Client
	Issues             *issues.Client
	Items              *items.Client
	JournalEntries     *journalentries.Client
	LinkToken          *linktoken.Client
	LinkedAccounts     *linkedaccounts.Client
	Passthrough        *passthrough.Client
	Payments           *payments.Client
	PhoneNumbers       *phonenumbers.Client
	PurchaseOrders     *purchaseorders.Client
	RegenerateKey      *regeneratekey.Client
	SyncStatus         *syncstatus.Client
	ForceResync        *forceresync.Client
	TaxRates           *taxrates.Client
	TrackingCategories *trackingcategories.Client
	Transactions       *transactions.Client
	VendorCredits      *vendorcredits.Client
	WebhookReceivers   *webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:            options.BaseURL,
		caller:             core.NewCaller(options.HTTPClient),
		header:             options.ToHeader(),
		AccountDetails:     accountdetails.NewClient(opts...),
		AccountToken:       accounttoken.NewClient(opts...),
		AccountingPeriods:  accountingperiods.NewClient(opts...),
		Accounts:           accounts.NewClient(opts...),
		Addresses:          addresses.NewClient(opts...),
		AsyncPassthrough:   asyncpassthrough.NewClient(opts...),
		Attachments:        attachments.NewClient(opts...),
		AuditTrail:         audittrail.NewClient(opts...),
		AvailableActions:   availableactions.NewClient(opts...),
		BalanceSheets:      balancesheets.NewClient(opts...),
		CashFlowStatements: cashflowstatements.NewClient(opts...),
		CompanyInfo:        companyinfo.NewClient(opts...),
		Contacts:           contacts.NewClient(opts...),
		CreditNotes:        creditnotes.NewClient(opts...),
		Scopes:             scopes.NewClient(opts...),
		DeleteAccount:      deleteaccount.NewClient(opts...),
		Expenses:           expenses.NewClient(opts...),
		FieldMapping:       fieldmapping.NewClient(opts...),
		GenerateKey:        generatekey.NewClient(opts...),
		IncomeStatements:   incomestatements.NewClient(opts...),
		Invoices:           invoices.NewClient(opts...),
		Issues:             issues.NewClient(opts...),
		Items:              items.NewClient(opts...),
		JournalEntries:     journalentries.NewClient(opts...),
		LinkToken:          linktoken.NewClient(opts...),
		LinkedAccounts:     linkedaccounts.NewClient(opts...),
		Passthrough:        passthrough.NewClient(opts...),
		Payments:           payments.NewClient(opts...),
		PhoneNumbers:       phonenumbers.NewClient(opts...),
		PurchaseOrders:     purchaseorders.NewClient(opts...),
		RegenerateKey:      regeneratekey.NewClient(opts...),
		SyncStatus:         syncstatus.NewClient(opts...),
		ForceResync:        forceresync.NewClient(opts...),
		TaxRates:           taxrates.NewClient(opts...),
		TrackingCategories: trackingcategories.NewClient(opts...),
		Transactions:       transactions.NewClient(opts...),
		VendorCredits:      vendorcredits.NewClient(opts...),
		WebhookReceivers:   webhookreceivers.NewClient(opts...),
	}
}
