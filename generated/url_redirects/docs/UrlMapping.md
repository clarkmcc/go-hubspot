# UrlMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique ID of this URL redirect. | [default to null]
**PortalId** | **int32** |  | [default to null]
**Created** | **int64** |  | [default to null]
**Updated** | **int64** |  | [default to null]
**CreatedById** | **int32** |  | [default to null]
**UpdatedById** | **int32** |  | [default to null]
**RoutePrefix** | **string** | The target incoming URL, path, or pattern to match for redirection. | [default to null]
**Destination** | **string** | The destination URL, where the target URL should be redirected if it matches the &#x60;routePrefix&#x60;. | [default to null]
**RedirectStyle** | **int32** | The type of redirect to create. Options include: 301 (permanent), 302 (temporary), or 305 (proxy). Find more details [here](https://knowledge.hubspot.com/cos-general/how-to-redirect-a-hubspot-page). | [default to null]
**ContentGroupId** | **int64** |  | [default to null]
**IsOnlyAfterNotFound** | **bool** | Whether the URL redirect mapping should apply only if a live page on the URL isn&#x27;t found. If False, the URL redirect mapping will take precedence over any existing page. | [default to null]
**IsRegex** | **bool** |  | [default to null]
**IsMatchFullUrl** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL, including the domain. | [default to null]
**IsMatchQueryString** | **bool** | Whether the &#x60;routePrefix&#x60; should match on the entire URL path, including the query string. | [default to null]
**IsPattern** | **bool** | Whether the &#x60;routePrefix&#x60; should match based on pattern. | [default to null]
**IsTrailingSlashOptional** | **bool** | Whether a trailing slash will be ignored. | [default to null]
**IsProtocolAgnostic** | **bool** | Whether the &#x60;routePrefix&#x60; should match both HTTP and HTTPS protocols. | [default to null]
**Name** | **string** |  | [default to null]
**Precedence** | **int32** | Used to prioritize URL redirection. If a given URL matches more than one redirect, the one with the **lower** precedence will be used. | [default to null]
**DeletedAt** | **int64** |  | [default to null]
**Note** | **string** |  | [default to null]
**Label** | **string** |  | [default to null]
**InternallyCreated** | **bool** |  | [default to null]
**CosObjectType** | **string** |  | [default to null]
**CdnPurgeEmbargoTime** | **int64** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

