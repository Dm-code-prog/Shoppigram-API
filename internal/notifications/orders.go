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
	"github.com/shoppigram-com/marketplace-api/packages/logger"
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
				s.log.Error("Failed to send the notifications for order events", logger.SilentError(err))
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
		return nil
	}

	// Update the cursor even if there were errors
	// during sending notifications
	// to avoid sending the same notifications again
	//
	// Ideally, we should break up this function into smaller, atomic functions
	// that can be retried independently
	defer func() {
		// Get the last element of the slice
		// as the last processed notification
		l := notifications[len(notifications)-1]
		err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
			CursorDate:      l.CreatedAt,
			LastProcessedID: l.ID,
			Name:            orderNotifier,
		})
		if err != nil {
			s.log.Error("Failed to update the notifier cursor", logger.SilentError(err))
		}
	}()

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
			} else if n.Status == stateDone {
				message, err = n.MakeDoneNotificationForAdmin(admin.Language)
				if err != nil {
					return errors.Wrap(err, "n.MakeDoneNotificationForAdmin")
				}
			} else {
				continue
			}

			// Add buttons to the message
			msg := tgbotapi.NewMessage(admin.Id, message)
			tgLink, err := s.makeAdminAppURL(n.WebAppID.String() + "/order/" + n.ID.String())
			if err != nil {
				return errors.Wrap(err, "makeAdminAppURL()")
			}

			addButtonsToMessage(
				&msg,
				telegramButtonData{
					getTranslation(admin.Language, "order-management"),
					tgLink,
				},
			)

			_, err = s.SendMessage(msg)
			if err != nil {
				return err
			}
		}

		// Send standard to the buyer

		var message string
		if n.Status == stateConfirmed {
			// send notifications to buyers
			message, err = n.MakeConfirmedNotificationForBuyer(n.BuyerLanguage)
			if err != nil {
				return errors.Wrap(err, "n.MakeConfirmedNotificationForBuyer")
			}
		} else if n.Status == stateDone {
			message, err = n.MakeDoneNotificationForBuyer(n.BuyerLanguage)
			if err != nil {
				return errors.Wrap(err, "n.MakeDoneNotificationForBuyer")
			}
		} else {
			continue
		}

		msg := tgbotapi.NewMessage(n.BuyerExternalID, message)
		tgLink, err := s.makeAdminAppURL(n.WebAppID.String() + "/order/" + n.ID.String())
		if err != nil {
			return errors.Wrap(err, "makeAdminAppURL")
		}

		addButtonsToMessage(
			&msg,
			telegramButtonData{
				getTranslation(n.BuyerLanguage, "view-order"),
				tgLink,
			},
		)

		_, err = s.SendMessage(msg)
		if err != nil {
			return errors.Wrap(err, "s.handleTelegramSendError")
		}

		// Send custom messages and media for products, if any
		for _, product := range n.Products {
			// Send custom message
			customMessage, err := s.repo.GetProductCustomMessage(s.ctx, product.ID, n.Status)
			if err != nil {
				return errors.Wrap(err, "s.repo.GetProductCustomMessage")
			}

			if customMessage != "" {
				_, err = s.SendMessage(tgbotapi.NewMessage(n.BuyerExternalID, customMessage))
				if err != nil {
					return errors.Wrap(err, "s.handleTelegramSendError")
				}
			}

			// Send custom media forwards
			forwardChat, messageID, err := s.repo.GetProductCustomMediaForward(s.ctx, product.ID, n.Status)
			if err != nil {
				return errors.Wrap(err, "s.repo.GetProductCustomMediaForward")
			}

			if forwardChat == 0 || messageID == 0 {
				continue
			}

			tgForward := tgbotapi.NewForward(n.BuyerExternalID, forwardChat, int(messageID))
			_, err = s.SendMessage(tgForward)
			if err != nil {
				return errors.Wrap(err, "s.handleTelegramSendError")
			}
		}
	}

	return nil
}
