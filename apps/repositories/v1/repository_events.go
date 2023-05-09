// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    repositoryEvents, err := UnmarshalRepositoryEvents(bytes)
//    bytes, err = repositoryEvents.Marshal()

package repositories

import "encoding/json"

func UnmarshalRepositoryEvents(data []byte) (RepositoryEvents, error) {
	var r RepositoryEvents
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RepositoryEvents) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Event data for Repository Events.
type RepositoryEvents struct {
	// List of repositories affected by the event               
	Repositories                                 []Repositories `json:"repositories"`
}

type Repositories struct {
	DistributionArch             *string  `json:"distribution_arch,omitempty"`
	DistributionVersions         []string `json:"distribution_versions,omitempty"`
	FailedIntrospectionsCount    *int64   `json:"failed_introspections_count,omitempty"`
	GPGKey                       *string  `json:"gpg_key,omitempty"`
	LastIntrospectionError       *string  `json:"last_introspection_error,omitempty"`
	LastIntrospectionTime        *string  `json:"last_introspection_time,omitempty"`
	LastSuccessIntrospectionTime *string  `json:"last_success_introspection_time,omitempty"`
	LastUpdateIntrospectionTime  *string  `json:"last_update_introspection_time,omitempty"`
	MetadataVerification         *bool    `json:"metadata_verification,omitempty"`
	Name                         string   `json:"name"`
	PackageCount                 *int64   `json:"package_count,omitempty"`
	Status                       *string  `json:"status,omitempty"`
	URL                          string   `json:"url"`
	UUID                         string   `json:"uuid"`
}
