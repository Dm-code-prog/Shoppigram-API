/**
This file implements the logic for sending notifications
about changes to marketplaces, to be precise:
- When a new marketplace is created
- When a marketplace is verified
*/

package notifications

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// RunNewMarketplaceNotifier starts a job that batch loads new marketplaces
// and sends notifications to the reviewers of marketplaces
func (s *Service) RunNewMarketplaceNotifier() error {
	ticker := time.NewTicker(s.newMarketplaceProcessingTimer)
	for {
		select {
		case <-ticker.C:
			err := s.runNewMarketplaceNotifier()
			if err != nil {
				s.log.Error("runNewMarketplaceNotifier failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runNewMarketplaceNotifier executes one iteration of loading a batch of new
// marketplaces and sending notifications to the reviewers of marketplaces
func (s *Service) runNewMarketplaceNotifier() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, newMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForNewMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForNewMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}
	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for new marketplaces")

	err = s.sendNewMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		if strings.Contains(err.Error(), "chat not found") {
			s.log.With(
				zap.String("method", "s.sendNewMarketplaceNotifications"),
				zap.String("user_id", strconv.FormatInt(marketplaceNotifications[0].OwnerExternalID, 10)),
			).Warn("chat not found, skipping notification sending")
		} else {
			return errors.Wrap(err, "s.sendNewMarketplaceNotifications")
		}
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            newMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// RunVerifiedMarketplaceNotifier starts a job that batch loads verified marketplaces
// and sends notifications to the owners of those marketplaces
func (s *Service) RunVerifiedMarketplaceNotifier() error {
	ticker := time.NewTicker(s.verifiedMarketplaceProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runVerifiedMarketplaceNotifierOnce()
			if err != nil {
				s.log.Error("runVerifiedMarketplaceNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runVerifiedMarketplaceNotifierOnce executes one iteration of loading a batch of
// verified marketplaces and sending notifications to the owners of those marketplaces
func (s *Service) runVerifiedMarketplaceNotifierOnce() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, verifiedMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for verified marketplaces")
	err = s.sendVerifiedMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendVerifiedMarketplaceNotifications")
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.VerifiedAt,
		LastProcessedID: lastElem.ID,
		Name:            verifiedMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// sendNewMarketplaceNotifications sends batch of notifications for new marketplaces
func (s *Service) sendNewMarketplaceNotifications(marketplaceNotifications []NewMarketplaceNotification) error {
	reviewers, err := s.repo.GetReviewersNotificationList(s.ctx)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetReviewersNotificationList")
	}
	for _, n := range marketplaceNotifications {
		n.ImageBaseUrl = s.bucketUrl

		ownerLang := s.checkAndGetLangCode(n.OwnerLanguage)
		onVerificationMsgTxt, err := n.BuildMessageAdmin(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(n.OwnerExternalID, onVerificationMsgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLink, err := s.getTelegramLink(n.ID.String())
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}
		buttonTextContactSupport := getTranslation(ownerLang, "contact-support")
		buttonTextViewStore := getTranslation(ownerLang, "view-store")

		addTelegramButtonsToMessage(&msg,
			telegramButtonData{buttonTextContactSupport, supportContactUrl},
			telegramButtonData{buttonTextViewStore, tgLink},
		)

		_, err = s.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "bot.Send to chat:"+strconv.FormatInt(n.OwnerExternalID, 10))
		}

		for _, r := range reviewers {
			msgTxt, err := n.BuildMessageShoppigram("en")
			if err != nil {
				return errors.Wrap(err, "a.BuildMessageShoppigram")
			}
			err = s.sendMessageToChat(r, msgTxt)
			if err != nil {
				return errors.Wrap(err, "sendMessageToChat")
			}
		}

	}

	return nil
}

// sendVerifiedMarketplaceNotifications sends batch of notifications for verified marketplaces
func (s *Service) sendVerifiedMarketplaceNotifications(marketplaceNotifications []VerifiedMarketplaceNotification) error {
	for _, notification := range marketplaceNotifications {
		ownerLang := s.checkAndGetLangCode(notification.OwnerLanguage)
		msgTxt, err := notification.BuildMessage(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(notification.OwnerExternalUserID, msgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLinkPath := notification.ID.String()
		tgLink, err := s.getTelegramLink(tgLinkPath)
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}

		buttonText := getTranslation(ownerLang, "continue-setting-up")
		addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

		_, err = s.bot.Send(msg)
		return s.handleTelegramSendError(err, notification.OwnerExternalUserID)
	}

	return nil
}
