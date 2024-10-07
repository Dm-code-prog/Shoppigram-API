package main

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/internal/admin"
	"github.com/shoppigram-com/marketplace-api/internal/notifications"
	"github.com/shoppigram-com/marketplace-api/internal/webhooks"
)

// notificationsAdminAdapter is an adapter for the notifications service
// meant to be used by the admins service
type notificationsAdminAdapter struct {
	notifier *notifications.Service
}

func (a *notificationsAdminAdapter) AddUserToNewOrderNotifications(ctx context.Context, req admin.AddUserToNewOrderNotificationsParams) error {
	return a.notifier.AddUserToNewOrderNotifications(ctx, notifications.AddUserToNewOrderNotificationsRequest(req))
}

func (a *notificationsAdminAdapter) SendMarketplaceBanner(_ context.Context, params admin.SendMarketplaceBannerParams) (message int64, err error) {
	return a.notifier.SendMarketplaceBanner(context.Background(), notifications.SendMarketplaceBannerParams(params))
}

func (a *notificationsAdminAdapter) PinNotification(ctx context.Context, req admin.PinNotificationParams) error {
	return a.notifier.PinNotification(ctx, notifications.PinNotificationParams(req))
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

func (a *notificationsWebhooksAdapter) NotifyGreetings(ctx context.Context, req webhooks.NotifyGreetingsRequest) error {
	return a.notifier.NotifyGreetings(ctx, notifications.NotifyGreetingsRequest(req))
}
