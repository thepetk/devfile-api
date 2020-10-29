package devfile

import (
	attributes "github.com/devfile/api/pkg/attributes"
)

// DevfileHeader describes the structure of the devfile-specific top-level fields
// that are not part of the K8S API structures
type DevfileHeader struct {
	// Devfile schema version
	// +kubebuilder:validation:Pattern=^([2-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$
	SchemaVersion string `json:"schemaVersion"`

	// +optional
	// Optional metadata
	Metadata DevfileMetadata `json:"metadata,omitempty"`
}

type DevfileMetadata struct {
	// Optional devfile name
	// +optional
	Name string `json:"name,omitempty"`

	// Optional semver-compatible version
	// +optional
	// +kubebuilder:validation:Pattern=^([0-9]+)\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$
	Version string `json:"version,omitempty"`

	// Map of implementation-dependant free-form YAML attributes.
	// +optional
	Attributes attributes.Attributes `json:"attributes,omitempty"`
}