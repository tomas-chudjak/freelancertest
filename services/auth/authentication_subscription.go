package auth

import "freelancertest/models"

// AuthenticationSubscriptionService is an implementation of service working through LDAP
type AuthenticationSubscriptionService struct {
	// repositories.LdapRepository
}

// NewAuthenticationSubscriptionService constructor
func NewAuthenticationSubscriptionService() *AuthenticationSubscriptionService {

	return &AuthenticationSubscriptionService{
		// LdapRepository: services.CreateRepository(ctx, "subscriptionData.AuthData.AuthenticationSubscription"),
	}
}


// AuthenticationSubscriptionSearch returns object by it's UeID
func (s *AuthenticationSubscriptionService) AuthenticationSubscriptionSearch(ueID string) (*models.AuthenticationSubscription, *models.ProblemDetails) {

	return nil, nil
}