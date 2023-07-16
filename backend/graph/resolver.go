package graph

import "github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Srv services.Services
	*Loaders
}
