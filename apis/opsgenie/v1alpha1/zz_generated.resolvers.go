/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AlertPolicy.
func (mg *AlertPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Responders); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Responders[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Responders[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Responders[i3].IDSelector,
			To: reference.To{
				List:    &TeamList{},
				Managed: &Team{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Responders[i3].ID")
		}
		mg.Spec.ForProvider.Responders[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Responders[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamIDRef,
		Selector:     mg.Spec.ForProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamID")
	}
	mg.Spec.ForProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ApiIntegration.
func (mg *ApiIntegration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OwnerTeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OwnerTeamIDRef,
		Selector:     mg.Spec.ForProvider.OwnerTeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OwnerTeamID")
	}
	mg.Spec.ForProvider.OwnerTeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OwnerTeamIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Responders); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Responders[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Responders[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Responders[i3].IDSelector,
			To: reference.To{
				List:    &TeamList{},
				Managed: &Team{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Responders[i3].ID")
		}
		mg.Spec.ForProvider.Responders[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Responders[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Escalation.
func (mg *Escalation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OwnerTeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OwnerTeamIDRef,
		Selector:     mg.Spec.ForProvider.OwnerTeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OwnerTeamID")
	}
	mg.Spec.ForProvider.OwnerTeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OwnerTeamIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rules); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Rules[i3].Recipient); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rules[i3].Recipient[i4].ID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Rules[i3].Recipient[i4].IDRef,
				Selector:     mg.Spec.ForProvider.Rules[i3].Recipient[i4].IDSelector,
				To: reference.To{
					List:    &TeamList{},
					Managed: &Team{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Rules[i3].Recipient[i4].ID")
			}
			mg.Spec.ForProvider.Rules[i3].Recipient[i4].ID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Rules[i3].Recipient[i4].IDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this TeamRoutingRule.
func (mg *TeamRoutingRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Notify); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Notify[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Notify[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Notify[i3].IDSelector,
			To: reference.To{
				List:    &EscalationList{},
				Managed: &Escalation{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Notify[i3].ID")
		}
		mg.Spec.ForProvider.Notify[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Notify[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamIDRef,
		Selector:     mg.Spec.ForProvider.TeamIDSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamID")
	}
	mg.Spec.ForProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamIDRef = rsp.ResolvedReference

	return nil
}
