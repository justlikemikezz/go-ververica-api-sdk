package vapi


type Affinity struct {
  NodeAffinity *NodeAffinity `json:"nodeAffinity,omitonempty"`
}

type NodeAffinity struct {
  RequiredDuringSchedulingIgnoredDuringExecution *RequiredDuringSchedulingIgnoredDuringExecution `json:"requiredDuringScedulingIgnoredDuringExecution,omitemtpy`
}

type RequiredDuringSchedulingIgnoredDuringExecution struct {
  NodeSelectorTerms []NodeSelectorTerms `json:"nodeSelectorTerms,omitempty"`
}
type NodeSelectorTerms struct {
  MatchExpressions *MatchExpressions `json:"matchExpressions,omitempty"`
}

type MatchExpressions struct {
  Key string `json:"key,omitempty"`
  Operator string `json:"operator,omitempty:"`
  Values []string `json:"values,omitempty"`
}

