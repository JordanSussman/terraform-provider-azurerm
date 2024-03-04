package staticsites

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteLinkedBackendARMResourceProperties struct {
	BackendResourceId *string `json:"backendResourceId,omitempty"`
	CreatedOn         *string `json:"createdOn,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Region            *string `json:"region,omitempty"`
}

func (o *StaticSiteLinkedBackendARMResourceProperties) GetCreatedOnAsTime() (*time.Time, error) {
	if o.CreatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *StaticSiteLinkedBackendARMResourceProperties) SetCreatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedOn = &formatted
}
