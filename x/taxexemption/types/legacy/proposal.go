package legacy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeAddTaxExemptionZone       = "AddTaxExemptionZone"
	ProposalTypeRemoveTaxExemptionZone    = "RemoveTaxExemptionZone"
	ProposalTypeModifyTaxExemptionZone    = "ModifyTaxExemptionZone"
	ProposalTypeAddTaxExemptionAddress    = "AddTaxExemptionAddress"
	ProposalTypeRemoveTaxExemptionAddress = "RemoveTaxExemptionAddress"
	RouterKey                             = "taxexemption"
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeAddTaxExemptionZone)
	govv1beta1.RegisterProposalType(ProposalTypeRemoveTaxExemptionZone)
	govv1beta1.RegisterProposalType(ProposalTypeModifyTaxExemptionZone)
	govv1beta1.RegisterProposalType(ProposalTypeAddTaxExemptionAddress)
	govv1beta1.RegisterProposalType(ProposalTypeRemoveTaxExemptionAddress)
}

// ======AddTaxExemptionZoneProposal======

func (p *AddTaxExemptionZoneProposal) GetTitle() string { return p.Title }

func (p *AddTaxExemptionZoneProposal) GetDescription() string { return p.Description }

func (p *AddTaxExemptionZoneProposal) ProposalRoute() string { return RouterKey }

func (p *AddTaxExemptionZoneProposal) ProposalType() string {
	return ProposalTypeAddTaxExemptionZone
}

func (p *AddTaxExemptionZoneProposal) String() string {
	return fmt.Sprintf(`AddTaxExemptionZoneProposal:
	Title:       %s
	Description: %s
	Zone: 	  	 %s
	Outgoing:    %t
	Incoming:    %t
	CrossZone:   %t
	Addresses:   %v
	Authority:   %s
  `, p.Title, p.Description, p.Zone, p.Outgoing, p.Incoming, p.CrossZone, p.Addresses, p.Authority)
}

func (p *AddTaxExemptionZoneProposal) ValidateBasic() error {
	err := govv1beta1.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if p.Zone == "" {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "zone name cannot be empty")
	}

	for _, address := range p.Addresses {
		_, err = sdk.AccAddressFromBech32(address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%s: %s", err, address)
		}
	}

	return nil
}

// ======RemoveTaxExemptionZoneProposal======

func (p *RemoveTaxExemptionZoneProposal) GetTitle() string { return p.Title }

func (p *RemoveTaxExemptionZoneProposal) GetDescription() string { return p.Description }

func (p *RemoveTaxExemptionZoneProposal) ProposalRoute() string { return RouterKey }

func (p *RemoveTaxExemptionZoneProposal) ProposalType() string {
	return ProposalTypeRemoveTaxExemptionZone
}

func (p *RemoveTaxExemptionZoneProposal) String() string {
	return fmt.Sprintf(`RemoveTaxExemptionZoneProposal:
	Title:       %s
	Description: %s
	Zone: 	  	 %s
	Authority:   %s
  `, p.Title, p.Description, p.Zone, p.Authority)
}

func (p *RemoveTaxExemptionZoneProposal) ValidateBasic() error {
	err := govv1beta1.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if p.Zone == "" {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "zone name cannot be empty")
	}

	return nil
}

// ======ModifyTaxExemptionZoneProposal======

func (p *ModifyTaxExemptionZoneProposal) GetTitle() string { return p.Title }

func (p *ModifyTaxExemptionZoneProposal) GetDescription() string { return p.Description }

func (p *ModifyTaxExemptionZoneProposal) ProposalRoute() string { return RouterKey }

func (p *ModifyTaxExemptionZoneProposal) ProposalType() string {
	return ProposalTypeModifyTaxExemptionZone
}

func (p *ModifyTaxExemptionZoneProposal) String() string {
	return fmt.Sprintf(`ModifyTaxExemptionZoneProposal:
	Title:       %s
	Description: %s
	Zone: 	  	 %s
	Outgoing:    %t
	Incoming:    %t
	CrossZone:   %t
	Authority:   %s
  `, p.Title, p.Description, p.Zone, p.Outgoing, p.Incoming, p.CrossZone, p.Authority)
}

func (p *ModifyTaxExemptionZoneProposal) ValidateBasic() error {
	err := govv1beta1.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if p.Zone == "" {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "zone name cannot be empty")
	}

	return nil
}

// ======AddTaxExemptionAddressProposal======

func (p *AddTaxExemptionAddressProposal) GetTitle() string { return p.Title }

func (p *AddTaxExemptionAddressProposal) GetDescription() string { return p.Description }

func (p *AddTaxExemptionAddressProposal) ProposalRoute() string { return RouterKey }

func (p *AddTaxExemptionAddressProposal) ProposalType() string {
	return ProposalTypeAddTaxExemptionAddress
}

func (p *AddTaxExemptionAddressProposal) String() string {
	return fmt.Sprintf(`AddTaxExemptionAddressProposal:
	Title:       %s
	Description: %s
	Zone: 	  	 %s
	Addresses:   %v
	Authority:   %s
  `, p.Title, p.Description, p.Zone, p.Addresses, p.Authority)
}

func (p *AddTaxExemptionAddressProposal) ValidateBasic() error {
	err := govv1beta1.ValidateAbstract(p)
	if err != nil {
		return err
	}

	for _, address := range p.Addresses {
		_, err = sdk.AccAddressFromBech32(address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%s: %s", err, address)
		}
	}

	return nil
}

// ======RemoveTaxExemptionAddressProposal======

func (p *RemoveTaxExemptionAddressProposal) GetTitle() string { return p.Title }

func (p *RemoveTaxExemptionAddressProposal) GetDescription() string { return p.Description }

func (p *RemoveTaxExemptionAddressProposal) ProposalRoute() string { return RouterKey }

func (p *RemoveTaxExemptionAddressProposal) ProposalType() string {
	return ProposalTypeRemoveTaxExemptionAddress
}

func (p *RemoveTaxExemptionAddressProposal) String() string {
	return fmt.Sprintf(`RemoveTaxExemptionAddressProposal:
	Title:       %s
	Description: %s
	Zone: 	  	 %s
	Addresses:   %v
	Authority:   %s
  `, p.Title, p.Description, p.Zone, p.Addresses, p.Authority)
}

func (p *RemoveTaxExemptionAddressProposal) ValidateBasic() error {
	err := govv1beta1.ValidateAbstract(p)
	if err != nil {
		return err
	}

	for _, address := range p.Addresses {
		_, err = sdk.AccAddressFromBech32(address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%s: %s", err, address)
		}
	}

	return nil
}
