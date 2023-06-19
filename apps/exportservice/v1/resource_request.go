// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    resourceRequest, err := UnmarshalResourceRequest(bytes)
//    bytes, err = resourceRequest.Marshal()

package exportservice

import "encoding/json"

func UnmarshalResourceRequest(data []byte) (ResourceRequest, error) {
	var r ResourceRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResourceRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Event data for data export requests
type ResourceRequest struct {
	// A request for data to be exported                     
	ResourceRequest                     ResourceRequestClass `json:"resource_request"`
}

// A request for data to be exported
type ResourceRequestClass struct {
	// The application being requested                                                                       
	Application                                                                       string                 `json:"application"`
	// The unique identifier of the export request that triggered the resource request                       
	ExportRequestUUID                                                                 string                 `json:"export_request_uuid"`
	// The filters to be applied to the data                                                                 
	Filters                                                                           map[string]interface{} `json:"filters,omitempty"`
	// The format of the data to be exported                                                                 
	Format                                                                            Format                 `json:"format"`
	// The resource to be exported                                                                           
	Resource                                                                          string                 `json:"resource"`
	// A unique identifier for the resource request                                                          
	UUID                                                                              string                 `json:"uuid"`
	// The Base64-encoded JSON identity header of the user making the request                                
	XRhIdentity                                                                       string                 `json:"x-rh-identity"`
}

// The format of the data to be exported
type Format string

const (
	CSV  Format = "csv"
	JSON Format = "json"
)
