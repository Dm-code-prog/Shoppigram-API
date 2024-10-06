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
	s.log.
		With(zap.String("count", strconv.Itoa(len(marketplaceNotifications)))).
		Info("sending notifications for new marketplaces")

	err = s.sendNewMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendNewMarketplaceNotifications")
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

		ownerLang := checkAndGetLangCode(n.OwnerLanguage)
		onVerificationMsgTxt, err := n.BuildMessageAdmin(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(n.OwnerExternalID, onVerificationMsgTxt)

		tgLink, err := s.makeMiniAppLink(n.ID.String())
		if err != nil {
			return errors.Wrap(err, "makeMiniAppLink()")
		}

		addButtonsToMessage(&msg,
			telegramButtonData{
				getTranslation(ownerLang, "contact-support"),
				supportContactUrl,
			},
			telegramButtonData{
				getTranslation(ownerLang, "view-store"),
				tgLink,
			},
		)

		_, err = s.SendMessage(msg)
		if err != nil {
			return err
		}

		for _, r := range reviewers {
			msgTxt, err := n.BuildMessageShoppigram("en")
			if err != nil {
				return errors.Wrap(err, "a.BuildMessageShoppigram")
			}

			msg := tgbotapi.NewMessage(r, msgTxt)
			_, err = s.SendMessage(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// sendVerifiedMarketplaceNotifications sends batch of notifications for verified marketplaces
func (s *Service) sendVerifiedMarketplaceNotifications(marketplaceNotifications []VerifiedMarketplaceNotification) error {
	for _, n := range marketplaceNotifications {
		ownerLang := checkAndGetLangCode(n.OwnerLanguage)
		msgTxt, err := n.BuildMessage(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}
		tgLink, err := s.makeMiniAppLink(n.ID.String())
		if err != nil {
			return errors.Wrap(err, "makeMiniAppLink()")
		}

		msg := tgbotapi.NewMessage(n.OwnerExternalUserID, msgTxt)
		addButtonsToMessage(
			&msg,
			telegramButtonData{
				getTranslation(ownerLang, "continue-setting-up"),
				tgLink,
			})
		_, err = s.SendMessage(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
