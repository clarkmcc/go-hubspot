package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	directoryUrl = "https://api.hubspot.com/api-catalog-public/v1/apis"
	replacer     = strings.NewReplacer(" ", "_", "-", "_")
)

//preProcessMap contains a grouped list of operations to rename before generating typings.
var preProcessMap = map[string][]preProcessEntry{
	"Accounting": {
		{old: "post-/crm/.*/extensions/accounting/callback/customer-create/{requestId}_createCustomer", new: "CallbackCreateCustomer"},
		{old: "post-/crm/.*/extensions/accounting/callback/customer-search/{requestId}_doCustomerSearch", new: "CallbackDoSearchCustomer"},
		{old: "post-/crm/.*/extensions/accounting/callback/exchange-rate/{requestId}_createExchangeRate", new: "CallbackCreateExchangeRate"},
		{old: "post-/crm/.*/extensions/accounting/callback/invoice-create/{requestId}_createInvoice", new: "CallbackCreateInvoice"},
		{old: "post-/crm/.*/extensions/accounting/callback/invoice-pdf/{requestId}_invoicePdf", new: "CallbackInvoicePDF"},
		{old: "post-/crm/.*/extensions/accounting/callback/invoice-search/{requestId}_doInvoiceSearch", new: "CallbackDoInvoiceSearch"},
		{old: "post-/crm/.*/extensions/accounting/callback/invoices/{requestId}_getById", new: "CallbackGetByID"},
		{old: "post-/crm/.*/extensions/accounting/callback/product-search/{requestId}_doProductSearch", new: "CallbackDoProductSearch"},
		{old: "post-/crm/.*/extensions/accounting/callback/tax-search/{requestId}_doTaxSearch", new: "CallbackDoTaxSearch"},
		{old: "post-/crm/.*/extensions/accounting/callback/terms/{requestId}_createTerm", new: "CallbackCreateTerm"},
		{old: "get-/crm/.*/extensions/accounting/invoice/{invoiceId}_getById", new: "InvoiceGetByID"},
		{old: "patch-/crm/.*/extensions/accounting/invoice/{invoiceId}_update", new: "InvoiceUpdate"},
		{old: "post-/crm/.*/extensions/accounting/invoice/{invoiceId}/payment_createPayment", new: "InvoiceCreatePayment"},
		{old: "get-/crm/.*/extensions/accounting/settings/{appId}_getById", new: "SettingsGetByID"},
		{old: "put-/crm/.*/extensions/accounting/settings/{appId}_replace", new: "SettingsReplace"},
		{old: "post-/crm/.*/extensions/accounting/sync/{appId}/contacts_createContact", new: "SyncCreateContact"},
		{old: "post-/crm/.*/extensions/accounting/sync/{appId}/products_createProduct", new: "SyncCreateProduct"},
		{old: "put-/crm/.*/extensions/accounting/user-accounts_replace", new: "UserAccountsReplace"},
		{old: "delete-/crm/.*/extensions/accounting/user-accounts/{accountId}_archive", new: "UserAccountsArchive"},
	},
	"Actions": {
		{old: "post-/automation/.*/actions/callbacks/complete_completeBatch", new: "CallbackCompleteBatch"},
		{old: "post-/automation/.*/actions/callbacks/{callbackId}/complete_complete", new: "CallbackComplete"},
		{old: "get-/automation/.*/actions/{appId}_getPage", new: "GetPage"},
		{old: "post-/automation/.*/actions/{appId}_create", new: "Create"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}_getById", new: "GetByID"},
		{old: "delete-/automation/.*/actions/{appId}/{definitionId}_archive", new: "Archive"},
		{old: "patch-/automation/.*/actions/{appId}/{definitionId}_update", new: "Update"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}/functions_getPage", new: "FunctionsGetPage"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}_getByFunctionType", new: "FunctionsGetByType"},
		{old: "put-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}_createOrReplaceByFunctionType", new: "FunctionsCreateOrReplaceByType"},
		{old: "delete-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}_archiveByFunctionType", new: "FunctionsArchiveByType"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}_getById", new: "FunctionsGetByID"},
		{old: "put-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}_createOrReplace", new: "FunctionsCreateOrReplace"},
		{old: "delete-/automation/.*/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}_archive", new: "FunctionsArchive"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}/revisions_getPage", new: "RevisionsGetPage"},
		{old: "get-/automation/.*/actions/{appId}/{definitionId}/revisions/{revisionId}_getById", new: "RevisionsGetByID"},
	},
	"Associations": {
		{old: "post-/crm/.*/associations/{fromObjectType}/{toObjectType}/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/associations/{fromObjectType}/{toObjectType}/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/associations/{fromObjectType}/{toObjectType}/batch/read_read", new: "BatchRead"},
		{old: "get-/crm/.*/associations/{fromObjectType}/{toObjectType}/types_getAll", new: "GetAll"},
	},
	"Audit-logs": {
		{old: "get-/cms/.*/audit-logs/_getPage", new: "GetPage"},
	},
	"Authors": {
		{old: "get-/cms/.*/blogs/authors_getPage", new: "GetPage"},
		{old: "post-/cms/.*/blogs/authors_create", new: "Create"},
		{old: "post-/cms/.*/blogs/authors/batch/archive_archiveBatch", new: "BatchArchive"},
		{old: "post-/cms/.*/blogs/authors/batch/create_createBatch", new: "BatchCreate"},
		{old: "post-/cms/.*/blogs/authors/batch/read_readBatch", new: "BatchRead"},
		{old: "post-/cms/.*/blogs/authors/batch/update_updateBatch", new: "BatchUpdate"},
		{old: "post-/cms/.*/blogs/authors/multi-language/attach-to-lang-group_attachToLangGroup", new: "AttachToLanguageGroup"},
		{old: "post-/cms/.*/blogs/authors/multi-language/create-language-variation_createLangVariation", new: "CreateLanguageVariation"},
		{old: "post-/cms/.*/blogs/authors/multi-language/detach-from-lang-group_detachFromLangGroup", new: "DetachFromLanguageGroup"},
		{old: "put-/cms/.*/blogs/authors/multi-language/set-new-lang-primary_setLangPrimary", new: "SetLanguagePrimary"},
		{old: "post-/cms/.*/blogs/authors/multi-language/update-languages_updateLangs", new: "UpdateLanguages"},
		{old: "get-/cms/.*/blogs/authors/{objectId}_getById", new: "GetByID"},
		{old: "delete-/cms/.*/blogs/authors/{objectId}_archive", new: "Archive"},
		{old: "patch-/cms/.*/blogs/authors/{objectId}_update", new: "Update"},
	},
	"Blog-posts": {
		{old: "get-/cms/.*/blogs/posts_getPage", new: "GetPage"},
		{old: "post-/cms/.*/blogs/posts_create", new: "Create"},
		{old: "post-/cms/.*/blogs/posts/batch/archive_archiveBatch", new: "BatchArchive"},
		{old: "post-/cms/.*/blogs/posts/batch/create_createBatch", new: "BatchCreate"},
		{old: "post-/cms/.*/blogs/posts/batch/read_readBatch", new: "BatchRead"},
		{old: "post-/cms/.*/blogs/posts/batch/update_updateBatch", new: "BatchUpdate"},
		{old: "post-/cms/.*/blogs/posts/clone_clone", new: "Clone"},
		{old: "post-/cms/.*/blogs/posts/multi-language/attach-to-lang-group_attachToLangGroup", new: "AttachToLanguageGroup"},
		{old: "post-/cms/.*/blogs/posts/multi-language/create-language-variation_createLangVariation", new: "CreateLanguageVariation"},
		{old: "post-/cms/.*/blogs/posts/multi-language/detach-from-lang-group_detachFromLangGroup", new: "DetachFromLanguageGroup"},
		{old: "put-/cms/.*/blogs/posts/multi-language/set-new-lang-primary_setLangPrimary", new: "SetLanguagePrimary"},
		{old: "post-/cms/.*/blogs/posts/multi-language/update-languages_updateLangs", new: "UpdateLanguages"},
		{old: "post-/cms/.*/blogs/posts/schedule_schedule", new: "Schedule"},
		{old: "get-/cms/.*/blogs/posts/{objectId}_getById", new: "GetByID"},
		{old: "delete-/cms/.*/blogs/posts/{objectId}_archive", new: "Archive"},
		{old: "patch-/cms/.*/blogs/posts/{objectId}_update", new: "Update"},
		{old: "get-/cms/.*/blogs/posts/{objectId}/draft_getDraftById", new: "GetDraftByID"},
		{old: "patch-/cms/.*/blogs/posts/{objectId}/draft_updateDraft", new: "UpdateDraft"},
		{old: "post-/cms/.*/blogs/posts/{objectId}/draft/push-live_pushLive", new: "PushLive"},
		{old: "post-/cms/.*/blogs/posts/{objectId}/draft/reset_resetDraft", new: "ResetDraft"},
		{old: "get-/cms/.*/blogs/posts/{objectId}/revisions_getPreviousVersions", new: "GetPreviousVersions"},
		{old: "get-/cms/.*/blogs/posts/{objectId}/revisions/{revisionId}_getPreviousVersion", new: "GetPreviousVersion"},
		{old: "post-/cms/.*/blogs/posts/{objectId}/revisions/{revisionId}/restore_restorePreviousVersion", new: "RestorePreviousVersion"},
		{old: "post-/cms/.*/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft_restorePreviousVersionToDraft", new: "RestorePreviousVersionToDraft"},
	},
	"Calling": {
		{old: "get-/crm/.*/extensions/calling/{appId}/settings_getById", new: "GetByID"},
		{old: "post-/crm/.*/extensions/calling/{appId}/settings_create", new: "Create"},
		{old: "delete-/crm/.*/extensions/calling/{appId}/settings_archive", new: "Archive"},
		{old: "patch-/crm/.*/extensions/calling/{appId}/settings_update", new: "Update"},
	},
	"Communications-status": {
		{old: "get-/communication-preferences/.*/definitions_getPage", new: "GetPage"},
		{old: "get-/communication-preferences/.*/status/email/{emailAddress}_getEmailStatus", new: "GetEmailStatus"},
		{old: "post-/communication-preferences/.*/subscribe_subscribe", new: "Subscribe"},
		{old: "post-/communication-preferences/.*/unsubscribe_unsubscribe", new: "Unsubscribe"},
	},
	"Companies": {
		{old: "get-/crm/.*/objects/companies_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/companies_create", new: "Create"},
		{old: "post-/crm/.*/objects/companies/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/companies/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/companies/{companyId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/companies/{companyId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/companies/{companyId}_update", new: "Update"},
		{old: "post-/crm/.*/objects/companies/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/companies/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/companies/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/companies/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/companies/{companyId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Contacts": {
		{old: "get-/crm/.*/objects/contacts_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/contacts_create", new: "Create"},
		{old: "post-/crm/.*/objects/contacts/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/.*/objects/contacts/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/contacts/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/contacts/{contactId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/contacts/{contactId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/contacts/{contactId}_update", new: "Update"},
		{old: "post-/crm/.*/objects/contacts/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/contacts/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/contacts/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/contacts/batch/update_update", new: "BatchUpdate"},
		{old: "get-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/contacts/{contactId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Crm Extensions": {
		{old: "get-/crm/.*/extensions/cards/sample-response_getCardsSampleResponse", new: "CardsGetSample"},
		{old: "get-/crm/.*/extensions/cards/{appId}_getAll", new: "CardsGetAll"},
		{old: "post-/crm/.*/extensions/cards/{appId}_create", new: "CardsCreate"},
		{old: "get-/crm/.*/extensions/cards/{appId}/{cardId}_getById", new: "CardsGetByID"},
		{old: "delete-/crm/.*/extensions/cards/{appId}/{cardId}_archive", new: "CardsArchive"},
		{old: "patch-/crm/.*/extensions/cards/{appId}/{cardId}_update", new: "CardsUpdate"},
	},
	"Custom Behavioral Events": {
		{old: "post-/events/.*/send", new: "Send"},
	},
	"Deals": {
		{old: "get-/crm/.*/objects/deals_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/deals_create", new: "Create"},
		{old: "post-/crm/.*/objects/deals/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/deals/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/deals/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/deals/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/deals/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/deals/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/deals/{dealId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/deals/{dealId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/deals/{dealId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/deals/{dealId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/deals/{dealId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Domains": {
		{old: "get-/cms/.*/domains/_getPage", new: "GetPage"},
		{old: "get-/cms/.*/domains/{domainId}_getById", new: "GetByID"},
	},
	"Events": {
		{old: "get-/events/.*/events_getPage", new: "GetPage"},
	},
	"Feedback Submissions": {
		{old: "get-/crm/.*/objects/feedback-submissions_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/feedback-submissions_create", new: "Create"},
		{old: "post-/crm/.*/objects/feedback-submissions/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/feedback-submissions/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/feedback-submissions/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/feedback-submissions/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/feedback-submissions/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/feedback-submissions/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}/associations/{toObjectType}_getPage", new: "AssociationsGetPage"},
		{old: "put-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/feedback-submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Hubdb": {
		{old: "get-/cms/.*/hubdb/tables_getAllTables", new: "GetAllTables"},
		{old: "post-/cms/.*/hubdb/tables_createTable", new: "CreateTable"},
		{old: "get-/cms/.*/hubdb/tables/draft_getAllDraftTables", new: "GetAllDraftTables"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}_getTableDetails", new: "GetTableDetails"},
		{old: "delete-/cms/.*/hubdb/tables/{tableIdOrName}_archiveTable", new: "ArchiveTable"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/draft_getDraftTableDetailsById", new: "GetDraftTableDetailsByID"},
		{old: "patch-/cms/.*/hubdb/tables/{tableIdOrName}/draft_updateDraftTable", new: "UpdateDraftTable"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/draft/clone_cloneDraftTable", new: "CloneDraftTable"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/draft/export_exportDraftTable", new: "ExportDraftTable"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/draft/import_importDraftTable", new: "ImportDraftTable"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/draft/publish_publishDraftTable", new: "PublishDraftTable"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/draft/reset_resetDraftTable", new: "ResetDraftTable"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/export_exportTable", new: "ExportTable"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/rows_getTableRows", new: "GetTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows_createTableRow", new: "CreateTableRow"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/batch/read_batchReadTableRows", new: "BatchReadTableRows"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft_readDraftTableRows", new: "ReadDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone_batchCloneDraftTableRows", new: "BatchCloneDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/create_batchCreateDraftTableRows", new: "BatchCreateDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge_batchPurgeDraftTableRows", new: "BatchPurgeDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/read_batchReadDraftTableRows", new: "BatchReadDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace_batchReplaceDraftTableRows", new: "BatchReplaceDraftTableRows"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/draft/batch/update_batchUpdateDraftTableRows", new: "BatchUpdateDraftTableRows"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}_getTableRow", new: "GetTableRow"},
		{old: "get-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft_getDraftTableRowById", new: "GetDraftTableRowByID"},
		{old: "put-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft_replaceDraftTableRow", new: "ReplaceDraftTableRow"},
		{old: "delete-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft_purgeDraftTableRow", new: "PurgeDraftTableRow"},
		{old: "patch-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft_updateDraftTableRow", new: "UpdateDraftTableRow"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone_cloneDraftTableRow", new: "CloneDraftTableRow"},
		{old: "post-/cms/.*/hubdb/tables/{tableIdOrName}/unpublish_unpublishTable", new: "UnpublishTable"},
	},
	"Imports": {
		{old: "get-/crm/.*/imports/_getPage", new: "GetPage"},
		{old: "post-/crm/.*/imports/_create", new: "Create"},
		{old: "get-/crm/.*/imports/{importId}_getById", new: "GetByID"},
		{old: "post-/crm/.*/imports/{importId}/cancel_cancel", new: "Cancel"},
		{old: "get-/crm/.*/imports/{importId}/errors_getErrors", new: "GetErrors"},
	},
	"Line Items": {
		{old: "get-/crm/.*/objects/line-items_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/line-items_create", new: "Create"},
		{old: "post-/crm/.*/objects/line-items/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/line-items/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/line-items/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/line-items/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/line-items/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/line-items/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/line-items/{lineItemId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/line-items/{lineItemId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/line-items/{lineItemId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/line-items/{lineItemId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/line-items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/line-items/{lineItemId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Marketing-events-beta": {
		{old: "post-/marketing/.*/marketing-events/attendance/{externalEventId}/{subscriberState}/create", new: "AttendanceCreate"},
		{old: "post-/marketing/.*/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create", new: "AttendanceEmailCreate"},
		{old: "post-/marketing/.*/marketing-events/events_create", new: "Create"},
		{old: "post-/marketing/.*/marketing-events/events/delete_archive", new: "Archive"},
		{old: "get-/marketing/.*/marketing-events/events/search_doSearch", new: "Search"},
		{old: "post-/marketing/.*/marketing-events/events/upsert_doUpsert", new: "Upsert"},
		{old: "get-/marketing/.*/marketing-events/events/{externalEventId}_getById", new: "ExternalGetByID"},
		{old: "put-/marketing/.*/marketing-events/events/{externalEventId}_replace", new: "ExternalReplace"},
		{old: "delete-/marketing/.*/marketing-events/events/{externalEventId}_archive", new: "ExternalArchive"},
		{old: "patch-/marketing/.*/marketing-events/events/{externalEventId}_update", new: "ExternalUpdate"},
		{old: "post-/marketing/.*/marketing-events/events/{externalEventId}/cancel_doCancel", new: "ExternalCancel"},
		{old: "post-/marketing/.*/marketing-events/events/{externalEventId}/complete", new: "ExternalComplete"},
		{old: "post-/marketing/.*/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert_doEmailUpsertById", new: "ExternalEmailUpsertByID"},
		{old: "post-/marketing/.*/marketing-events/events/{externalEventId}/{subscriberState}/upsert_doUpsertById", new: "ExternalUpsertByID"},
		{old: "get-/marketing/.*/marketing-events/{appId}/settings_getAll", new: "SettingsGetAll"},
		{old: "post-/marketing/.*/marketing-events/{appId}/settings_create", new: "SettingsCreate"},
	},
	"Oauth": {
		{old: "get-/oauth/.*/access-tokens/{token}_getAccessToken", new: "GetAccessToken"},
		{old: "get-/oauth/.*/refresh-tokens/{token}_getRefreshToken", new: "GetRefreshToken"},
		{old: "delete-/oauth/.*/refresh-tokens/{token}_archiveRefreshToken", new: "ArchiveRefreshToken"},
		{old: "post-/oauth/.*/token_createToken", new: "CreateToken"},
	},
	"Objects": {
		{old: "get-/crm/.*/objects/{objectType}_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/{objectType}_create", new: "Create"},
		{old: "post-/crm/.*/objects/{objectType}/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/{objectType}/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/{objectType}/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/{objectType}/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/{objectType}/gdpr-delete_purge", new: "Delete"},
		{old: "post-/crm/.*/objects/{objectType}/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/{objectType}/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/{objectType}/{objectId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/{objectType}/{objectId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/{objectType}/{objectId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Owners": {
		{old: "get-/crm/.*/owners/_getPage", new: "GetPage"},
		{old: "get-/crm/.*/owners/{ownerId}_getById", new: "GetByID"},
	},
	"Performance": {
		{old: "get-/cms/.*/performance/_getPage", new: "GetPage"},
		{old: "get-/cms/.*/performance/uptime_getUptime", new: "GetUptime"},
	},
	"Pipelines": {
		{old: "get-/crm/.*/pipelines/{objectType}_getAll", new: "GetAll"},
		{old: "post-/crm/.*/pipelines/{objectType}_create", new: "Create"},
		{old: "get-/crm/.*/pipelines/{objectType}/{pipelineId}_getById", new: "GetByID"},
		{old: "put-/crm/.*/pipelines/{objectType}/{pipelineId}_replace", new: "Replace"},
		{old: "delete-/crm/.*/pipelines/{objectType}/{pipelineId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/pipelines/{objectType}/{pipelineId}_update", new: "Update"},
		{old: "get-/crm/.*/pipelines/{objectType}/{pipelineId}/audit_getAudit", new: "GetAudit"},
		{old: "get-/crm/.*/pipelines/{objectType}/{pipelineId}/stages_getAll", new: "StagesGetAll"},
		{old: "post-/crm/.*/pipelines/{objectType}/{pipelineId}/stages_create", new: "StagesCreate"},
		{old: "get-/crm/.*/pipelines/{objectType}/{pipelineId}/stages/{stageId}_getById", new: "StagesGetByID"},
		{old: "put-/crm/.*/pipelines/{objectType}/{pipelineId}/stages/{stageId}_replace", new: "StagesReplace"},
		{old: "delete-/crm/.*/pipelines/{objectType}/{pipelineId}/stages/{stageId}_archive", new: "StagesArchive"},
		{old: "patch-/crm/.*/pipelines/{objectType}/{pipelineId}/stages/{stageId}_update", new: "StagesUpdate"},
		{old: "get-/crm/.*/pipelines/{objectType}/{pipelineId}/stages/{stageId}/audit_getAudit", new: "StagesGetAudit"},
	},
	"Products": {
		{old: "get-/crm/.*/objects/products_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/products_create", new: "Create"},
		{old: "post-/crm/.*/objects/products/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/products/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/products/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/products/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/products/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/products/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/products/{productId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/products/{productId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/products/{productId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/products/{productId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Properties": {
		{old: "get-/crm/.*/properties/{objectType}_getAll", new: "GetAll"},
		{old: "post-/crm/.*/properties/{objectType}_create", new: "Create"},
		{old: "post-/crm/.*/properties/{objectType}/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/properties/{objectType}/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/properties/{objectType}/batch/read_read", new: "BatchRead"},
		{old: "get-/crm/.*/properties/{objectType}/groups_getAll", new: "GroupsGetAll"},
		{old: "post-/crm/.*/properties/{objectType}/groups_create", new: "GroupsCreate"},
		{old: "get-/crm/.*/properties/{objectType}/groups/{groupName}_getByName", new: "GroupsGetByName"},
		{old: "delete-/crm/.*/properties/{objectType}/groups/{groupName}_archive", new: "GroupsArchive"},
		{old: "patch-/crm/.*/properties/{objectType}/groups/{groupName}_update", new: "GroupsUpdate"},
		{old: "get-/crm/.*/properties/{objectType}/{propertyName}_getByName", new: "GetByName"},
		{old: "delete-/crm/.*/properties/{objectType}/{propertyName}_archive", new: "Archive"},
		{old: "patch-/crm/.*/properties/{objectType}/{propertyName}_update", new: "Update"},
	},
	"Quotes": {
		{old: "get-/crm/.*/objects/quotes_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/quotes_create", new: "Create"},
		{old: "post-/crm/.*/objects/quotes/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/quotes/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/quotes/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/quotes/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/quotes/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/quotes/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/quotes/{quoteId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/quotes/{quoteId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/quotes/{quoteId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/quotes/{quoteId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/quotes/{quoteId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/quotes/{quoteId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Schemas": {
		{old: "get-/crm-object-schemas/.*/schemas_getAll", new: "GetAll"},
		{old: "post-/crm-object-schemas/.*/schemas_create", new: "Create"},
		{old: "get-/crm-object-schemas/.*/schemas/{objectType}_getById", new: "GetByID"},
		{old: "delete-/crm-object-schemas/.*/schemas/{objectType}_archive", new: "Archive"},
		{old: "patch-/crm-object-schemas/.*/schemas/{objectType}_update", new: "Update"},
		{old: "post-/crm-object-schemas/.*/schemas/{objectType}/associations_createAssociation", new: "CreateAssociation"},
		{old: "delete-/crm-object-schemas/.*/schemas/{objectType}/associations/{associationIdentifier}_archiveAssociation", new: "ArchiveAssociation"},
		{old: "delete-/crm-object-schemas/.*/schemas/{objectType}/purge_purge", new: "Purge"},
	},
	"Site-search": {
		{old: "get-/cms/.*/site-search/indexed-data/{contentId}_getById", new: "GetByID"},
		{old: "get-/cms/.*/site-search/search_search", new: "Search"},
	},
	"Source Code": {
		{old: "post-/cms/.*/source-code/extract/async_doAsync", new: "ExtractAsync"},
		{old: "get-/cms/.*/source-code/extract/async/tasks/{taskId}/status_getAsyncStatus", new: "ExtractGetAsyncStatus"},
		{old: "post-/cms/.*/source-code/extract/{path}_extractByPath", new: "ExtractByPath"},
		{old: "get-/cms/.*/source-code/{environment}/content/{path}_get", new: "ContentGet"},
		{old: "put-/cms/.*/source-code/{environment}/content/{path}_replace", new: "ContentReplace"},
		{old: "post-/cms/.*/source-code/{environment}/content/{path}_create", new: "ContentCreate"},
		{old: "delete-/cms/.*/source-code/{environment}/content/{path}_archive", new: "ContentArchive"},
		{old: "get-/cms/.*/source-code/{environment}/metadata/{path}_get", new: "MetadataGet"},
		{old: "post-/cms/.*/source-code/{environment}/validate/{path}_doValidate", new: "Validate"},
	},
	"Tags": {
		{old: "get-/cms/.*/blogs/tags_getPage", new: "GetPage"},
		{old: "post-/cms/.*/blogs/tags_create", new: "Create"},
		{old: "post-/cms/.*/blogs/tags/batch/archive_archiveBatch", new: "BatchArchive"},
		{old: "post-/cms/.*/blogs/tags/batch/create_createBatch", new: "BatchCreate"},
		{old: "post-/cms/.*/blogs/tags/batch/read_readBatch", new: "BatchRead"},
		{old: "post-/cms/.*/blogs/tags/batch/update_updateBatch", new: "BatchUpdate"},
		{old: "post-/cms/.*/blogs/tags/multi-language/attach-to-lang-group_attachToLangGroup", new: "AttachToLanguageGroup"},
		{old: "post-/cms/.*/blogs/tags/multi-language/create-language-variation_createLangVariation", new: "CreateLanguageVariation"},
		{old: "post-/cms/.*/blogs/tags/multi-language/detach-from-lang-group_detachFromLangGroup", new: "DetachFromLanguageGroup"},
		{old: "put-/cms/.*/blogs/tags/multi-language/set-new-lang-primary_setLangPrimary", new: "SetLanguagePrimary"},
		{old: "post-/cms/.*/blogs/tags/multi-language/update-languages_updateLangs", new: "UpdateLanguages"},
		{old: "get-/cms/.*/blogs/tags/{objectId}_getById", new: "GetByID"},
		{old: "delete-/cms/.*/blogs/tags/{objectId}_archive", new: "Archive"},
		{old: "patch-/cms/.*/blogs/tags/{objectId}_update", new: "Update"},
	},
	"Tickets": {
		{old: "get-/crm/.*/objects/tickets_getPage", new: "GetPage"},
		{old: "post-/crm/.*/objects/tickets_create", new: "Create"},
		{old: "post-/crm/.*/objects/tickets/batch/archive_archive", new: "BatchArchive"},
		{old: "post-/crm/.*/objects/tickets/batch/create_create", new: "BatchCreate"},
		{old: "post-/crm/.*/objects/tickets/batch/read_read", new: "BatchRead"},
		{old: "post-/crm/.*/objects/tickets/batch/update_update", new: "BatchUpdate"},
		{old: "post-/crm/.*/objects/tickets/merge_merge", new: "Merge"},
		{old: "post-/crm/.*/objects/tickets/search_doSearch", new: "Search"},
		{old: "get-/crm/.*/objects/tickets/{ticketId}_getById", new: "GetByID"},
		{old: "delete-/crm/.*/objects/tickets/{ticketId}_archive", new: "Archive"},
		{old: "patch-/crm/.*/objects/tickets/{ticketId}_update", new: "Update"},
		{old: "get-/crm/.*/objects/tickets/{ticketId}/associations/{toObjectType}_getAll", new: "AssociationsGetAll"},
		{old: "put-/crm/.*/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType}_create", new: "AssociationsCreate"},
		{old: "delete-/crm/.*/objects/tickets/{ticketId}/associations/{toObjectType}/{toObjectId}/{associationType}_archive", new: "AssociationsArchive"},
	},
	"Timeline": {
		{old: "post-/integrators/timeline/.*/events_create", new: "Create"},
		{old: "post-/integrators/timeline/.*/events/batch/create_createBatch", new: "BatchCreate"},
		{old: "get-/integrators/timeline/.*/events/{eventTemplateId}/{eventId}_getById", new: "GetByID"},
		{old: "get-/integrators/timeline/.*/events/{eventTemplateId}/{eventId}/detail_getDetailById", new: "GetDetailByID"},
		{old: "get-/integrators/timeline/.*/events/{eventTemplateId}/{eventId}/render_getRenderById", new: "GetRenderByID"},
		{old: "get-/integrators/timeline/.*/{appId}/event-templates_getAll", new: "TemplateGetAll"},
		{old: "post-/integrators/timeline/.*/{appId}/event-templates_create", new: "TemplateCreate"},
		{old: "get-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}_getById", new: "TemplatesGetByID"},
		{old: "put-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}_update", new: "TemplatesUpdate"},
		{old: "delete-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}_archive", new: "TemplatesArchive"},
		{old: "post-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}/tokens_create", new: "TemplatesTokensCreate"},
		{old: "put-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName}_update", new: "TemplatesTokensUpdate"},
		{old: "delete-/integrators/timeline/.*/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName}_archive", new: "TemplatesTokensArchive"},
	},
	"Transactional": {
		{old: "post-/marketing/.*/transactional/single-email/send_sendEmail", new: "SendEmail"},
		{old: "get-/marketing/.*/transactional/smtp-tokens_getTokensPage", new: "GetTokensPage"},
		{old: "post-/marketing/.*/transactional/smtp-tokens_createToken", new: "CreateToken"},
		{old: "get-/marketing/.*/transactional/smtp-tokens/{tokenId}_getTokenById", new: "GetTokenByID"},
		{old: "delete-/marketing/.*/transactional/smtp-tokens/{tokenId}_archiveToken", new: "ArchiveToken"},
		{old: "post-/marketing/.*/transactional/smtp-tokens/{tokenId}/password-reset_resetPassword", new: "ResetPassword"},
	},
	"Url-redirects": {
		{old: "get-/cms/.*/url-redirects/_getPage", new: "GetPage"},
		{old: "post-/cms/.*/url-redirects/_create", new: "Create"},
		{old: "get-/cms/.*/url-redirects/{urlRedirectId}_getById", new: "GetByID"},
		{old: "delete-/cms/.*/url-redirects/{urlRedirectId}_archive", new: "Archive"},
		{old: "patch-/cms/.*/url-redirects/{urlRedirectId}_update", new: "Update"},
	},
	"Videoconferencing": {
		{old: "get-/crm/.*/extensions/videoconferencing/settings/{appId}_getById", new: "GetByID"},
		{old: "put-/crm/.*/extensions/videoconferencing/settings/{appId}_replace", new: "Replace"},
		{old: "delete-/crm/.*/extensions/videoconferencing/settings/{appId}_archive", new: "Archive"},
	},
	"Visitor Identification": {
		{old: "post-/visitor-identification/.*/tokens/create_generateToken", new: "GenerateToken"},
	},
	"Webhooks": {
		{old: "get-/webhooks/.*/{appId}/settings_getAll", new: "SettingsGetAll"},
		{old: "put-/webhooks/.*/{appId}/settings_configure", new: "SettingsConfigure"},
		{old: "delete-/webhooks/.*/{appId}/settings_clear", new: "SettingsClear"},
		{old: "get-/webhooks/.*/{appId}/subscriptions_getAll", new: "SubscriptionsGetAll"},
		{old: "post-/webhooks/.*/{appId}/subscriptions_create", new: "SubscriptionsCreate"},
		{old: "post-/webhooks/.*/{appId}/subscriptions/batch/update_updateBatch", new: "SubscriptionsBatchUpdate"},
		{old: "get-/webhooks/.*/{appId}/subscriptions/{subscriptionId}_getById", new: "SubscriptionsGetByID"},
		{old: "delete-/webhooks/.*/{appId}/subscriptions/{subscriptionId}_archive", new: "SubscriptionsArchive"},
		{old: "patch-/webhooks/.*/{appId}/subscriptions/{subscriptionId}_update", new: "SubscriptionsUpdate"},
	},
}

// preProcessEntry contains a mapping from an old operation to a new operation.
type preProcessEntry struct {
	old string
	new string
}

// schemaOpenAPI contains the OpenAPI schema definition.
type schemaOpenAPI struct {
	Info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
}

// directory is the schema for the https://api.hubspot.com/api-catalog-public/v1/apis page which lists all
// the publicly available APIs.
type directory struct {
	Results []struct {
		Name     string `json:"name"`
		Features map[string]struct {
			OpenAPI string `json:"openAPI"`
			Stage   string `json:"stage"`
		} `json:"features"`
	} `json:"results"`
}

// retrieveSchema downloads a schema from the given url.
// It returns the schema text as a string.
func retrieveSchema(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	schema, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(schema), nil
}

// preProcessSchema replaces operation names in the given schema and saves the results to file.
// It returns a path to the saved file.
func preProcessSchema(group string, schema string) (string, error) {
	entries := preProcessMap[group]

	for _, entry := range entries {
		pattern := regexp.MustCompile(entry.old)
		schema = pattern.ReplaceAllString(schema, entry.new)
	}

	filename := "./schema/" + group + ".json"
	err := ioutil.WriteFile(filename, []byte(schema), 0644)

	if err != nil {
		return "", err
	}

	return filename, nil
}

// versionFromSchema retrieves the OpenAPI schema version from the input JSON string.
// It returns the API version.
func versionFromSchema(input string) (string, error) {
	var schema schemaOpenAPI
	err := json.Unmarshal([]byte(input), &schema)

	if err != nil {
		return "", err
	}

	return schema.Info.Version, nil
}

func main() {
	res, err := http.Get(directoryUrl)
	if err != nil {
		panic(err)
	}
	var r directory
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_USER_ID", "clarkmcc")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_REPO_ID", "go-hubspot")
	if err != nil {
		panic(err)
	}
	generator, err := exec.LookPath("openapi-generator")
	if err != nil {
		panic(err)
	}

	//i := 0
	//outer:
	for _, group := range r.Results {
		for name, feature := range group.Features {
			schema, err := retrieveSchema(feature.OpenAPI)

			if err != nil {
				panic(err)
			}

			filename, err := preProcessSchema(name, schema)

			if err != nil {
				panic(err)
			}

			version, err := versionFromSchema(schema)

			if err != nil {
				panic(err)
			}

			//if i > 0 {
			//	break outer
			//}
			name = replacer.Replace(strings.ToLower(name))
			fmt.Printf("generating group/feature %s/%s\n", strings.ToLower(group.Name), name)
			_, err = exec.Command(generator, "generate",
				"-i", filename,
				"-g", "go",
				"--package-name", name,
				"--additional-properties=isGoSubmodule=false",
				"--skip-validate-spec",
				"-o", "./generated/"+version+"/"+name,
			).CombinedOutput()
			if err != nil {
				panic(err)
			}
			//i++
		}
	}

	fmt.Println("updating generated files")
	err = filepath.Walk("generated", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// API files check for the presence of an API key in the request context using type
		// assertion. Because we generate each module separately, each module has its own
		// type for storing the API key in the context (contacts.APIKey vs deals.APIKey). In
		// order to allow us to share the same authorizer across packages, we'll do a find
		// and replace on all the package-specific API key structs and replace them with our
		// own global structs.
		if strings.HasSuffix(info.Name(), ".go") {
			// Open the file, read it, then close it
			var b []byte
			err = func() error {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				b, err = ioutil.ReadAll(file)
				if err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}

			for _, method := range findMethods {
				if !bytes.Contains(b, method) {
					continue
				}

				// Replace generated API key auth with custom API key auth
				b = bytes.ReplaceAll(b, method, replaceWith)

				// Add imports for authorization package
				idx := bytes.Index(b, []byte("\"net/url\""))
				if idx >= 0 {
					b = bytes.Join([][]byte{b[:idx], []byte("\t\"github.com/clarkmcc/go-hubspot\""), b[idx:]}, []byte("\n"))
				}

				err = os.WriteFile(path, b, 0666)
				if err != nil {
					return err
				}
				break
			}
		}
		// The client configuration gets generated with a go module for each generated
		// client. We'll go through and delete each submodule so that we can take advantage
		// of the root parent module.
		if info.Name() != "go.mod" && info.Name() != "go.sum" && info.Name() != "git_push.sh" {
			return nil
		}
		return os.Remove(path)
	})
	if err != nil {
		panic(err)
	}
}

var findPrivateAppsAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"private_apps\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarHeaderParams[\"private-app\"] = key\n\t\t\t}\n\t\t}\n\t}")
var findPrivateAppsLegacyAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"private_apps_legacy\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarHeaderParams[\"private-app-legacy\"] = key\n\t\t\t}\n\t\t}\n\t}")
var findDeveloperHapiKeyAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"developer_hapikey\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarQueryParams.Add(\"hapikey\", key)\n\t\t\t}\n\t\t}\n\t}")
var replaceWith = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {\n\t\t\tauth.Apply(hubspot.AuthorizationRequest{\n\t\t\t\tQueryParams: localVarQueryParams,\n\t\t\t\tFormParams:  localVarFormParams,\n\t\t\t\tHeaders:     localVarHeaderParams,\n\t\t\t})\n\t\t}\n\t}")

var findMethods = [][]byte{
	findPrivateAppsAuth,
	findPrivateAppsLegacyAuth,
	findDeveloperHapiKeyAuth,
}
