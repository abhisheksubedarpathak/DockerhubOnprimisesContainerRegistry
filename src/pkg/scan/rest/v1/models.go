// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"encoding/json"
	"fmt"

	"github.com/goharbor/harbor/src/lib/errors"
)

const (
	supportVulnerability = "support_vulnerability"
	supportSBOM          = "support_sbom"
)

var supportedMimeTypes = []string{
	MimeTypeNativeReport,
	MimeTypeGenericVulnerabilityReport,
	MimeTypeSBOMReport,
}

// Scanner represents metadata of a Scanner Adapter which allow Harbor to lookup a scanner capable of
// scanning a given Artifact stored in its registry and making sure that it can interpret a
// returned result.
type Scanner struct {
	// The name of the scanner.
	Name string `json:"name"`
	// The name of the scanner's provider.
	Vendor string `json:"vendor"`
	// The version of the scanner.
	Version string `json:"version"`
}

// ScannerCapability consists of the set of recognized artifact MIME types and the set of scanner
// report MIME types. For example, a scanner capable of analyzing Docker images and producing
// a vulnerabilities report recognizable by Harbor web console might be represented with the
// following capability:
//   - type: vulnerability
//   - consumes MIME types:
//     -- application/vnd.oci.image.manifest.v1+json
//     -- application/vnd.docker.distribution.manifest.v2+json
//   - produces MIME types
//     -- application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0
//     -- application/vnd.scanner.adapter.vuln.report.raw
type ScannerCapability struct {
	// The type of the scanner capability, vulnerability or sbom
	Type string `json:"type"`
	// The set of MIME types of the artifacts supported by the scanner to produce the reports
	// specified in the "produces_mime_types". A given mime type should only be present in one
	// capability item.
	ConsumesMimeTypes []string `json:"consumes_mime_types"`
	// The set of MIME types of reports generated by the scanner for the consumes_mime_types of
	// the same capability record.
	ProducesMimeTypes []string `json:"produces_mime_types"`
}

// ScannerProperties is a set of custom properties that can further describe capabilities of a given scanner.
type ScannerProperties map[string]string

// ScannerAdapterMetadata represents metadata of a Scanner Adapter which allows Harbor to lookup
// a scanner capable of scanning a given Artifact stored in its registry and making sure that it
// can interpret a returned result.
type ScannerAdapterMetadata struct {
	Scanner      *Scanner             `json:"scanner"`
	Capabilities []*ScannerCapability `json:"capabilities"`
	Properties   ScannerProperties    `json:"properties"`
}

// Validate validate the metadata
func (md *ScannerAdapterMetadata) Validate() error {
	// Validate the required properties
	if md.Scanner == nil ||
		len(md.Scanner.Name) == 0 ||
		len(md.Scanner.Version) == 0 ||
		len(md.Scanner.Vendor) == 0 {
		return errors.New("invalid scanner in metadata")
	}

	if len(md.Capabilities) == 0 {
		return errors.New("invalid capabilities in metadata")
	}

	for _, ca := range md.Capabilities {
		// v1.MimeTypeDockerArtifact is required now
		found := false
		for _, cm := range ca.ConsumesMimeTypes {
			if cm == MimeTypeDockerArtifact {
				found = true
				break
			}
		}
		if !found {
			return errors.Errorf("missing %s in consumes_mime_types", MimeTypeDockerArtifact)
		}

		// either of v1.MimeTypeNativeReport OR v1.MimeTypeGenericVulnerabilityReport is required
		found = false
		for _, pm := range ca.ProducesMimeTypes {
			if isSupportedMimeType(pm) {
				found = true
				break
			}
		}

		if !found {
			return errors.Errorf("missing %s or %s in produces_mime_types", MimeTypeNativeReport, MimeTypeGenericVulnerabilityReport)
		}
	}

	return nil
}

func isSupportedMimeType(mimeType string) bool {
	for _, mt := range supportedMimeTypes {
		if mt == mimeType {
			return true
		}
	}
	return false
}

// HasCapability returns true when mine type of the artifact support by the scanner
func (md *ScannerAdapterMetadata) HasCapability(mimeType string) bool {
	for _, capability := range md.Capabilities {
		for _, mt := range capability.ConsumesMimeTypes {
			if mt == mimeType {
				return true
			}
		}
	}

	return false
}

// GetCapability returns capability for the mime type
func (md *ScannerAdapterMetadata) GetCapability(mimeType string) *ScannerCapability {
	for _, capability := range md.Capabilities {
		for _, mt := range capability.ConsumesMimeTypes {
			if mt == mimeType {
				return capability
			}
		}
	}

	return nil
}

// ConvertCapability converts the capability to map, used in get scanner API
func (md *ScannerAdapterMetadata) ConvertCapability() map[string]interface{} {
	capabilities := make(map[string]interface{})
	for _, c := range md.Capabilities {
		if c.Type == ScanTypeVulnerability {
			capabilities[supportVulnerability] = true
		}
		if c.Type == ScanTypeSbom {
			capabilities[supportSBOM] = true
		}
	}
	return capabilities
}

// Artifact represents an artifact stored in Registry.
type Artifact struct {
	// ID of the namespace (project). It will not be sent to scanner adapter.
	NamespaceID int64 `json:"namespace_id,omitempty"`
	// The full name of a Harbor repository containing the artifact, including the namespace.
	// For example, `library/oracle/nosql`.
	Repository string `json:"repository"`
	// The info used to identify the version of the artifact,
	Tag string `json:"tag"`
	// The artifact's digest, consisting of an algorithm and hex portion.
	// For example, `sha256:6c3c624b58dbbcd3c0dd82b4c53f04194d1247c6eebdaab7c610cf7d66709b3b`,
	// represents sha256 based digest.
	Digest string `json:"digest"`
	// The mime type of the scanned artifact
	MimeType string `json:"mime_type"`
	// The size the scanned artifact
	Size int64 `json:"size"`
}

// Registry represents Registry connection settings.
type Registry struct {
	// A base URL of the Docker Registry v2 API exposed by Harbor.
	URL string `json:"url"`
	// An optional value of the HTTP Authorization header sent with each request to the Docker Registry for getting or exchanging token.
	// For example, `Basic: Base64(username:password)`.
	Authorization string `json:"authorization"`
}

// ScanRequest represents a structure that is sent to a Scanner Adapter to initiate artifact scanning.
// Conducts all the details required to pull the artifact from a Harbor registry.
type ScanRequest struct {
	// Connection settings for the Docker Registry v2 API exposed by Harbor.
	Registry *Registry `json:"registry"`
	// Artifact to be scanned.
	Artifact *Artifact `json:"artifact"`
	// RequestType
	RequestType []*ScanType `json:"enabled_capabilities"`
}

// ScanType represent the type of the scan request
type ScanType struct {
	// Type sets the type of the scan, it could be sbom or vulnerability, default is vulnerability
	Type string `json:"type"`
	// ProducesMimeTypes defines scanreport should be
	ProducesMimeTypes []string `json:"produces_mime_types"`
	// Parameters extra parameters
	Parameters map[string]interface{} `json:"parameters"`
}

// FromJSON parses ScanRequest from json data
func (s *ScanRequest) FromJSON(jsonData string) error {
	if len(jsonData) == 0 {
		return errors.New("empty json data to parse")
	}

	return json.Unmarshal([]byte(jsonData), s)
}

// ToJSON marshals ScanRequest to JSON data
func (s *ScanRequest) ToJSON() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Validate ScanRequest
func (s *ScanRequest) Validate() error {
	if s.Registry == nil ||
		len(s.Registry.URL) == 0 {
		return errors.New("scan request: invalid registry")
	}

	if s.Artifact == nil ||
		len(s.Artifact.Digest) == 0 ||
		len(s.Artifact.Repository) == 0 ||
		len(s.Artifact.MimeType) == 0 {
		return errors.New("scan request: invalid artifact")
	}

	return nil
}

// ScanResponse represents the response returned by the scanner adapter after scan request successfully
// submitted.
type ScanResponse struct {
	// e.g: 3fa85f64-5717-4562-b3fc-2c963f66afa6
	ID string `json:"id"`
}

// ErrorResponse contains error message when requests are not correctly handled.
type ErrorResponse struct {
	// Error object
	Err *Error `json:"error"`
}

// Error message
type Error struct {
	// Message of the error
	Message string `json:"message"`
}

// Error for ErrorResponse
func (er *ErrorResponse) Error() string {
	if er.Err != nil {
		return er.Err.Message
	}

	return "nil error"
}

// ReportNotReadyError is an error to indicate the scan report is not ready
type ReportNotReadyError struct {
	// Seconds for next retry with seconds
	RetryAfter int
}

// Error for ReportNotReadyError
func (rnr *ReportNotReadyError) Error() string {
	return fmt.Sprintf("report is not ready yet, retry after %d", rnr.RetryAfter)
}
