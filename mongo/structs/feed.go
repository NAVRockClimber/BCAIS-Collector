package mongoschema

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type FeedDocument struct {
	// Supported capabilities of a feed.
	Capabilities *FeedCapabilities `deepcopier:"Capabilities" bson:"capabilities,omitempty"`
	// This will either be the feed GUID or the feed GUID and view GUID depending on how the feed was accessed.
	FullyQualifiedId *string `deepcopier:"FullyQualifiedId" bson:"fullyQualifiedId,omitempty"`
	// Full name of the view, in feed@view format.
	FullyQualifiedName *string `deepcopier:"FullyQualifiedName" bson:"fullyQualifiedName,omitempty"`
	// A GUID that uniquely identifies this feed.
	Id *uuid.UUID `deepcopier:"Id" bson:"id,omitempty"`
	// If set, all packages in the feed are immutable.  It is important to note that feed views are immutable; therefore, this flag will always be set for views.
	IsReadOnly *bool `deepcopier:"IsReadOnly" bson:"isReadOnly,omitempty"`
	// A name for the feed. feed names must follow these rules: <list type="bullet"><item><description> Must not exceed 64 characters </description></item><item><description> Must not contain whitespaces </description></item><item><description> Must not start with an underscore or a period </description></item><item><description> Must not end with a period </description></item><item><description> Must not contain any of the following illegal characters: <![CDATA[ @, ~, ;, {, }, \, +, =, <, >, |, /, \\, ?, :, &, $, *, \", #, [, ] ]]></description></item></list>
	Name *string `deepcopier:"Name" bson:"name,omitempty"`
	// The project that this feed is associated with.
	Project *ProjectReference `deepcopier:"Project" bson:"project,omitempty"`
	// OBSOLETE: This should always be true.  Setting to false will override all sources in UpstreamSources.
	UpstreamEnabled *bool `deepcopier:"UpstreamEnabled" bson:"upstreamEnabled,omitempty"`
	// A list of sources that this feed will fetch packages from.  An empty list indicates that this feed will not search any additional sources for packages.
	UpstreamSources *[]UpstreamSource `deepcopier:"UpstreamSources" bson:"upstreamSources,omitempty"`
	// Definition of the view.
	View *FeedView `deepcopier:"View" bson:"view,omitempty"`
	// View Id.
	ViewId *uuid.UUID `deepcopier:"ViewId" bson:"viewId,omitempty"`
	// View name.
	ViewName *string `deepcopier:"ViewName" bson:"viewName,omitempty"`
	// Related REST links.
	Links interface{} `deepcopier:"Links" bson:"_links,omitempty"`
	// If set, this feed supports generation of package badges.
	BadgesEnabled *bool `deepcopier:"BadgesEnabled" bson:"badgesEnabled,omitempty"`
	// The view that the feed administrator has indicated is the default experience for readers.
	DefaultViewId *uuid.UUID `deepcopier:"DefaultViewId" bson:"defaultViewId,omitempty"`
	// The date that this feed was deleted.
	DeletedDate *azuredevops.Time `deepcopier:"DeletedDate" bson:"deletedDate,omitempty"`
	// A description for the feed.  Descriptions must not exceed 255 characters.
	Description *string `deepcopier:"Description" bson:"description,omitempty"`
	// If set, the feed will hide all deleted/unpublished versions
	HideDeletedPackageVersions *bool `deepcopier:"HideDeletedPackageVersions" bson:"hideDeletedPackageVersions,omitempty"`
	// Explicit permissions for the feed.
	Permissions *[]FeedPermission `deepcopier:"Permissions" bson:"permissions,omitempty"`
	// If set, time that the UpstreamEnabled property was changed. Will be null if UpstreamEnabled was never changed after Feed creation.
	UpstreamEnabledChangedDate *azuredevops.Time `deepcopier:"UpstreamEnabledChangedDate" bson:"upstreamEnabledChangedDate,omitempty"`
	// The URL of the base feed in GUID form.
	Url *string `deepcopier:"Url" bson:"url,omitempty"`
}

type FeedCapabilities string

type ProjectReference struct {
	// Gets or sets id of the project.
	Id *uuid.UUID `deepcopier:"Id" bson:"id,omitempty"`
	// Gets or sets name of the project.
	Name *string `deepcopier:"Name" bson:"name,omitempty"`
	// Gets or sets visibility of the project.
	Visibility *string `deepcopier:"Visibility" bson:"visibility,omitempty"`
}

type FeedPermission struct {
	// Display name for the identity.
	DisplayName *string `deepcopier:"DisplayName" bson:"displayName,omitempty"`
	// Identity associated with this role.
	IdentityDescriptor *string `deepcopier:"IdentityDescriptor" bson:"identityDescriptor,omitempty"`
	// Id of the identity associated with this role.
	IdentityId *uuid.UUID `deepcopier:"IdentityId" bson:"identityId,omitempty"`
	// The role for this identity on a feed.
	Role *FeedRole `deepcopier:"Role" bson:"role,omitempty"`
}

type FeedRole string

// Upstream source definition, including its Identity, package type, and other associated information.
type UpstreamSource struct {
	// UTC date that this upstream was deleted.
	DeletedDate *azuredevops.Time `deepcopier:"DeletedDate" bson:"deletedDate,omitempty"`
	// Locator for connecting to the upstream source in a user friendly format, that may potentially change over time
	DisplayLocation *string `deepcopier:"DisplayLocation" bson:"displayLocation,omitempty"`
	// Identity of the upstream source.
	Id *uuid.UUID `deepcopier:"Id" bson:"id,omitempty"`
	// For an internal upstream type, track the Azure DevOps organization that contains it.
	InternalUpstreamCollectionId *uuid.UUID `deepcopier:"InternalUpstreamCollectionId" bson:"internalUpstreamCollectionId,omitempty"`
	// For an internal upstream type, track the feed id being referenced.
	InternalUpstreamFeedId *uuid.UUID `deepcopier:"InternalUpstreamFeedId" bson:"internalUpstreamFeedId,omitempty"`
	// For an internal upstream type, track the view of the feed being referenced.
	InternalUpstreamViewId *uuid.UUID `deepcopier:"InternalUpstreamViewId" bson:"internalUpstreamViewId,omitempty"`
	// Consistent locator for connecting to the upstream source.
	Location *string `deepcopier:"Location" bson:"location,omitempty"`
	// Display name.
	Name *string `deepcopier:"Name" bson:"name,omitempty"`
	// Package type associated with the upstream source.
	Protocol *string `deepcopier:"Protocol" bson:"protocol,omitempty"`
	// Source type, such as Public or Internal.
	UpstreamSourceType *UpstreamSourceType `deepcopier:"UpstreamSourceType" bson:"upstreamSourceType,omitempty"`
}

// Type of an upstream source, such as Public or Internal.
type UpstreamSourceType string

// A view on top of a feed.
type FeedView struct {
	// Related REST links.
	Links interface{} `deepcopier:"Links" bson:"_links,omitempty"`
	// Id of the view.
	Id *uuid.UUID `deepcopier:"Id" bson:"id,omitempty"`
	// Name of the view.
	Name *string `deepcopier:"Name" bson:"name,omitempty"`
	// Type of view.
	Type *FeedViewType `deepcopier:"Type" bson:"type,omitempty"`
	// Url of the view.
	Url *string `deepcopier:"Url" bson:"url,omitempty"`
	// Visibility status of the view.
	Visibility *FeedVisibility `deepcopier:"Visibility" bson:"visibility,omitempty"`
}

// The type of view, often used to control capabilities and exposure to options such as promote.  Implicit views are internally created only.
type FeedViewType string

// Feed visibility controls the scope in which a certain feed is accessible by a particular user
type FeedVisibility string
