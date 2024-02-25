// This file was auto-generated by Fern from our API Definition.

package ats

type CreateFieldMappingRequest struct {
	// The name of the target field you want this remote field to map to.
	TargetFieldName string `json:"target_field_name"`
	// The description of the target field you want this remote field to map to.
	TargetFieldDescription string `json:"target_field_description"`
	// The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
	RemoteFieldTraversalPath []interface{} `json:"remote_field_traversal_path,omitempty"`
	// The method of the remote endpoint where the remote field is coming from.
	RemoteMethod string `json:"remote_method"`
	// The path of the remote endpoint where the remote field is coming from.
	RemoteUrlPath string `json:"remote_url_path"`
	// The name of the Common Model that the remote field corresponds to in a given category.
	CommonModelName string `json:"common_model_name"`
}

type PatchedEditFieldMappingRequest struct {
	// The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
	RemoteFieldTraversalPath []interface{} `json:"remote_field_traversal_path,omitempty"`
	// The method of the remote endpoint where the remote field is coming from.
	RemoteMethod *string `json:"remote_method,omitempty"`
	// The path of the remote endpoint where the remote field is coming from.
	RemoteUrlPath *string `json:"remote_url_path,omitempty"`
}

type RemoteFieldsRetrieveRequest struct {
	// A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
	CommonModels *string `json:"-"`
	// If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
	IncludeExampleValues *string `json:"-"`
}
