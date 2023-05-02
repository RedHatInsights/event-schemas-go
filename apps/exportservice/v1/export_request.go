// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    exportRequest, err := UnmarshalExportRequest(bytes)
//    bytes, err = exportRequest.Marshal()

package exportservice

import "encoding/json"

func UnmarshalExportRequest(data []byte) (ExportRequest, error) {
	var r ExportRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Event data for data export requests.
type ExportRequest struct {
	// A request for data to be exported                   
	ExportRequest                       ExportRequestClass `json:"exportRequest"`
}

// A request for data to be exported
type ExportRequestClass struct {
	// The application being requested                                                              
	Application                                                              string                 `json:"application"`
	// The filters to be applied to the data                                                        
	Filters                                                                  map[string]interface{} `json:"filters,omitempty"`
	// The format of the data to be exported                                                        
	Format                                                                   Format                 `json:"format"`
	// The resource to be exported                                                                  
	Resource                                                                 string                 `json:"resource"`
	// A unique identifier for the request                                                          
	UUID                                                                     string                 `json:"uuid"`
	// The Base64-encoded JSON identity header of the user making the request                       
	XRhIdentity                                                              string                 `json:"x-rh-identity"`
}

// The format of the data to be exported
type Format string
const (
	CSV Format = "csv"
	JSON Format = "json"
)
