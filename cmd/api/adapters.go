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
	admin *admins.Service
}

func (a *adminWebhooksAdapter) CreateOrUpdateTelegramChannel(ctx context.Context, req webhooks.CreateOrUpdateTelegramChannelRequest) error {
	return a.admin.CreateOrUpdateTelegramChannel(ctx, admins.CreateOrUpdateTelegramChannelRequest(req))
}

type notificationsWebhooksAdapter struct {
	notifier *notifications.Service
}

func (a *notificationsWebhooksAdapter) NotifyChannelIntegrationSuccess(ctx context.Context, req webhooks.NotifyChannelIntegrationSuccessRequest) error {
	return a.notifier.NotifyChannelIntegrationSuccess(ctx, notifications.NotifyChannelIntegrationSuccessRequest(req))
}
