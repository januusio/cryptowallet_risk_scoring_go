/*
 * Risk scoring API
 *
 * This is an API that provides whitebox risk scoring
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Reason for the risk score
type Reason struct {
	Explanation string `json:"explanation"`
	Label string `json:"label"`
	Offsets *Offsets `json:"offsets"`
	RiskElaboration *RiskElaboration `json:"riskElaboration"`
}
