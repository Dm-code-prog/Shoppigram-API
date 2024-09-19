package main

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/internal/admins"
	"github.com/shoppigram-com/marketplace-api/internal/notifications"
	"github.com/shoppigram-com/marketplace-api/internal/webhooks"
)

// notificationsAdminAdapter is an adapter for the notifications service
// meant to be used by the admins service
type notificationsAdminAdapter struct {
	notifier *notifications.Service
}

func (a *notificationsAdminAdapter) AddUserToNewOrderNotifications(ctx context.Context, req admins.AddUserToNewOrderNotificationsParams) error {
	return a.notifier.AddUserToNewOrderNotifications(ctx, notifications.AddUserToNewOrderNotificationsRequest(req))
}

func (a *notificationsAdminAdapter) SendMarketplaceBanner(_ context.Context, params admins.SendMarketplaceBannerParams) (message int64, err error) {
	return a.notifier.SendMarketplaceBanner(context.Background(), notifications.SendMarketplaceBannerParams(params))
}

func (a *notificationsAdminAdapter) PinNotification(ctx context.Context, req admins.PinNotificationParams) error {
	return a.notifier.PinNotification(ctx, notifications.PinNotificationParams(req))
}

// adminWebhooksAdapter is an adapter for the admin service
// meant to be used by the webhooks service
type adminWebhooksAdapter struct {
	admin admins.Service
}

func (a *adminWebhooksAdapter) CreateOrUpdateTelegramChannel(ctx context.Context, req webhooks.CreateOrUpdateTelegramChannelRequest) error {
	return a.admin.CreateOrUpdateTelegramChannel(ctx, admins.CreateOrUpdateTelegramChannelRequest(req))
}

func (a *adminWebhooksAdapter) GetTelegramChannelOwner(ctx context.Context, req webhooks.GetTelegramChannelOwnerRequest) (webhooks.GetTelegramChannelOwnerResponse, error) {
	resp, err := a.admin.GetTelegramChannelOwner(ctx, admins.GetTelegramChannelOwnerRequest{ChannelChatId: req.ChannelChatId})
	return webhooks.GetTelegramChannelOwnerResponse(resp), err
}

type notificationsWebhooksAdapter struct {
	notifier *notifications.Service
}

func (a *notificationsWebhooksAdapter) NotifyChannelIntegrationSuccess(ctx context.Context, req webhooks.NotifyChannelIntegrationSuccessRequest) error {
	return a.notifier.NotifyChannelIntegrationSuccess(ctx, notifications.NotifyChannelIntegrationSuccessRequest(req))
}

func (a *notificationsWebhooksAdapter) NotifyChannelIntegrationFailure(ctx context.Context, req webhooks.NotifyChannelIntegrationFailureRequest) error {
	return a.notifier.NotifyChannelIntegrationFailure(ctx, notifications.NotifyChannelIntegrationFailureRequest(req))
}

func (a *notificationsWebhooksAdapter) NotifyBotRemovedFromChannel(ctx context.Context, req webhooks.NotifyBotRemovedFromChannelRequest) error {
	return a.notifier.NotifyBotRemovedFromChannel(ctx, notifications.NotifyBotRemovedFromChannelRequest(req))
}

func (a *notificationsWebhooksAdapter) NotifyGreetings(ctx context.Context, req webhooks.NotifyGreetingsRequest) error {
	return a.notifier.NotifyGreetings(ctx, notifications.NotifyGreetingsRequest(req))
}
