package service

import (
	"context"
	"net/http"

	"github.com/deepcode-ai/hermes/domain"
	"github.com/deepcode-ai/hermes/provider"
	"github.com/deepcode-ai/hermes/provider/jira"
	"github.com/deepcode-ai/hermes/provider/slack"

	log "github.com/sirupsen/logrus"
)

type ProviderService interface {
	GetProvider(context.Context, *GetProviderReqeuest) (*GetProviderResponse, domain.IError)
}

type GetProviderReqeuest struct {
	Token string              `header:"X-Notifier-Token"`
	Type  domain.ProviderType `param:"provider"`
}

type GetProviderResponse struct {
	Type   string                 `json:"type"`
	Values map[string]interface{} `json:"values"`
}

type providerService struct{}

func NewProviderService() ProviderService {
	return &providerService{}
}

func (service *providerService) GetProvider(ctx context.Context, request *GetProviderReqeuest) (*GetProviderResponse, domain.IError) {
	provider := newProvider(request.Type)
	response, err := provider.GetOptValues(ctx, &domain.NotifierSecret{Token: request.Token})
	if err != nil {
		log.Errorf("Failed to get options values: %v", err)
		return nil, errUnprocessable(err.Error())
	}
	return &GetProviderResponse{Type: string(request.Type), Values: response}, nil
}

func newProvider(providerType domain.ProviderType) provider.Provider {
	switch providerType {
	case slack.ProviderType:
		return slack.NewSlackProvider(http.DefaultClient)
	case jira.ProviderType:
		return jira.NewJIRAProvider(http.DefaultClient)

	}
	return nil
}
