/**
This file implements the logic for sending notifications
about orders, to be precise:
- When an order is confirmed
- When an order is done
*/

package notifications

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"time"
)

// RunOrdersNotifier starts a job that batch loads new orders
// and sends notifications to the owners of marketplaces
func (s *Service) RunOrdersNotifier() error {
	ticker := time.NewTicker(s.newOrderProcessingTimer)
	for {
		select {
		case <-ticker.C:
			err := s.runOrdersNotifier()
			if err != nil {
				s.log.Error("runOrdersNotifier failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runOrdersNotifier sends notifications for orders
func (s *Service) runOrdersNotifier() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, orderNotifier)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	notifications, err := s.repo.GetNotificationsForOrders(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForOrders")
	}

	if len(notifications) == 0 {
		s.log.Info("no updates of orders")
		return nil
	}

	for _, n := range notifications {
		admins, err := s.repo.GetAdminsNotificationList(s.ctx, n.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}

		// send notifications to admins
		for _, admin := range admins {
			var message string

			// For confirmed orders
			if n.Status == stateConfirmed {
				message, err = n.MakeConfirmedNotificationForAdmin(admin.Language)
				if err != nil {
					return errors.Wrap(err, "n.MakeConfirmedNotificationForAdmin")
				}
			}

			// Add buttons to the message
			tgMessage := tgbotapi.NewMessage(admin.Id, message)
			tgLink, err := s.getTelegramLink(n.WebAppID.String() + "/order/" + n.ID.String())
			if err != nil {
				return errors.Wrap(err, "getTelegramLink()")
			}

			addTelegramButtonsToMessage(
				&tgMessage,
				telegramButtonData{
					getTranslation(admin.Language, "order-management"),
					tgLink,
				},
			)

			_, err = s.bot.Send(tgMessage)
			return s.handleTelegramSendError(err, admin.Id)
		}

		// Send notifications to the buyer

		var message string

		// For confirmed orders
		if n.Status == stateConfirmed {
			// send notifications to buyers
			message, err = n.MakeConfirmedNotificationForBuyer(n.BuyerLanguage)
			if err != nil {
				return errors.Wrap(err, "n.MakeConfirmedNotificationForBuyer")
			}
		}

		// For done orders
		if n.Status == stateDone {
			message, err = n.MakeDoneNotificationForBuyer(n.BuyerLanguage)
			if err != nil {
				return errors.Wrap(err, "n.MakeDoneNotificationForBuyer")
			}
		}

		tgMessage := tgbotapi.NewMessage(n.BuyerExternalID, message)
		tgLink, err := s.getTelegramLink(n.WebAppID.String() + "/order/" + n.ID.String())
		if err != nil {
			return errors.Wrap(err, "getTelegramLink")
		}

		addTelegramButtonsToMessage(
			&tgMessage,
			telegramButtonData{
				getTranslation(n.BuyerLanguage, "view-order"),
				tgLink,
			},
		)

		_, err = s.bot.Send(tgMessage)
		return s.handleTelegramSendError(err, n.BuyerExternalID)
	}

	// Get the last element of the slice
	// as the last processed notification

	l := notifications[len(notifications)-1]
	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      l.CreatedAt,
		LastProcessedID: l.ID,
		Name:            orderNotifier,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}
