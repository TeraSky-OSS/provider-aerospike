/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"aerospike_role": roleNameConfig(),
	"aerospike_user": userNameConfig(),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

func userNameConfig() config.ExternalName {
	e := config.ParameterAsIdentifier("user_name")

	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, setup map[string]any) (string, error) {
		user, ok := parameters["user_name"]
		if !ok {
			user = ""
		}
		return user.(string), nil
	}

	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["user_name"]
		if !ok {
			return "", errors.New("user_name attribute missing from state file")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("value of user_name needs to be string")
		}
		return idStr, nil
	}

	return e
}

func roleNameConfig() config.ExternalName {
	e := config.ParameterAsIdentifier("role_name")

	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, setup map[string]any) (string, error) {
		role, ok := parameters["role_name"]
		if !ok {
			role = ""
		}
		return role.(string), nil
	}

	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["role_name"]
		if !ok {
			return "", errors.New("role_name attribute missing from state file")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("value of role_name needs to be string")
		}
		return idStr, nil
	}

	return e
}
