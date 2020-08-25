package controllers

import (
	"services/auth_data"

	"github.com/go-openapi/runtime/middleware"

	"restapi/operations/auth_data"
)

// QueryAuthSubsDataHandlerFunc handles GET /subscription-data/{ueId}/authentication-data/authentication-subscription
func QueryAuthSubsDataHandlerFunc(params auth_data.QueryAuthSubsDataParams) middleware.Responder {

	service := auth_data.NewAuthenticationSubscriptionService(params.HTTPRequest.Context())

	payload, problem := service.AuthenticationSubscriptionSearch(params.UeID)
	if problem != nil {
		return auth_data.NewQueryAuthSubsDataDefault(int(problem.Status)).WithPayload(problem)
	}

	return auth_data.NewQueryAuthSubsDataOK().WithPayload(payload)
}