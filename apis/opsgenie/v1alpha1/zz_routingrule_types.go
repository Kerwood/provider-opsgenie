/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CriteriaConditionsObservation struct {
}

type CriteriaConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, extra-properties, recipients, teams or priority.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// If field is set as extra-properties, key could be used for key-value pair.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default value is false.
	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// (true) It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty and equals-ignore-whitespace.
	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// The order of the team routing rule within the rules. order value is actually the index of the team routing rule whose minimum value is 0 and whose maximum value is n-1 (number of team routing rules is n)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// List of conditions will be checked before applying team routing rule. This field declaration should be omitted if the criteria type is set to match-all.
	// +kubebuilder:validation:Optional
	Conditions []CriteriaConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NotifyObservation struct {
}

type NotifyParameters struct {

	// +crossplane:generate:reference:type=Escalation
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Escalation to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Escalation to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Name of the team routing rule
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RoutingRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoutingRuleParameters struct {

	// You can refer Criteria for detailed information about criteria and its fields
	// +kubebuilder:validation:Optional
	Criteria []CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Only use when importing default routing rule
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Name of the team routing rule
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target entity of schedule, escalation, or the reserved word none which will be notified in routing rule. The possible values are: schedule, escalation, none
	// +kubebuilder:validation:Required
	Notify []NotifyParameters `json:"notify" tf:"notify,omitempty"`

	// The order of the team routing rule within the rules. order value is actually the index of the team routing rule whose minimum value is 0 and whose maximum value is n-1 (number of team routing rules is n)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Id of the team owning the routing rule
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TimeRestriction []RoutingRuleTimeRestrictionParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`

	// Timezone of team routing rule. If timezone field is not given, account timezone is used as default.You can refer to Supported Locale IDs for available timezones
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RoutingRuleTimeRestrictionObservation struct {
}

type RoutingRuleTimeRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Restriction []TimeRestrictionRestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// +kubebuilder:validation:Optional
	Restrictions []TimeRestrictionRestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type TimeRestrictionRestrictionObservation struct {
}

type TimeRestrictionRestrictionParameters struct {

	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionsObservation struct {
}

type TimeRestrictionRestrictionsParameters struct {

	// +kubebuilder:validation:Required
	EndDay *string `json:"endDay" tf:"end_day,omitempty"`

	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Required
	StartDay *string `json:"startDay" tf:"start_day,omitempty"`

	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

// RoutingRuleSpec defines the desired state of RoutingRule
type RoutingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingRuleParameters `json:"forProvider"`
}

// RoutingRuleStatus defines the observed state of RoutingRule.
type RoutingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingRule is the Schema for the RoutingRules API. Manages a Team Routing Rule within Opsgenie.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type RoutingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoutingRuleSpec   `json:"spec"`
	Status            RoutingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingRuleList contains a list of RoutingRules
type RoutingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingRule `json:"items"`
}

// Repository type metadata.
var (
	RoutingRule_Kind             = "RoutingRule"
	RoutingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingRule_Kind}.String()
	RoutingRule_KindAPIVersion   = RoutingRule_Kind + "." + CRDGroupVersion.String()
	RoutingRule_GroupVersionKind = CRDGroupVersion.WithKind(RoutingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingRule{}, &RoutingRuleList{})
}
