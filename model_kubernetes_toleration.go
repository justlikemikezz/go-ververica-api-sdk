package vapi

type Toleration struct {
  Key string `json:"key,omitempty"`
  Operator string `json:"operator,omitempty"`
  Value string `json:"value,omitempty"`
  Effect string `json:"effect,omitempty"`
}