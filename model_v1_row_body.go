/*
 * Stapl External API
 *
 * API for interacting with the Stapl dashboard's contacts.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1RowBody struct {
	Email string `json:"Email,omitempty"`
	ColumnName string `json:"Column Name,omitempty"`
	Workflows []string `json:"workflows,omitempty"`
}