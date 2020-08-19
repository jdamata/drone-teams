// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Settings for the plugin.
type Settings struct {
	Webhook string
	Status  string
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Verify the webhook endpoint
	if p.settings.Webhook == "" {
		// If webhook is undefined, check if the ${DRONE_BRANCH}_teams_webhook env var is defined.
		branchWebhook := fmt.Sprintf("%s_teams_webhook", os.Getenv("DRONE_BRANCH"))
		if os.Getenv(branchWebhook) == "" {
			return fmt.Errorf("no webhook endpoint provided")
		}
		// Set webhook setting to ${DRONE_BRANCH}_teams_webhook
		p.settings.Webhook = os.Getenv(branchWebhook)
	}

	if p.settings.Status == "" {
		p.settings.Status = p.pipeline.Build.Status
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {

	// Default card color is green
	themeColor := "96FF33"

	// Create list of card facts
	facts := []MessageCardSectionFact{
		{
			Name:  "Build Number",
			Value: fmt.Sprintf("%d", p.pipeline.Build.Number),
		},
		{
			Name:  "Started",
			Value: p.pipeline.Build.Started.String(),
		},
		{
			Name:  "Repo Link",
			Value: p.pipeline.Repo.Link,
		},
		{
			Name:  "Branch",
			Value: p.pipeline.Build.Branch,
		},
		{
			Name:  "Git Author",
			Value: fmt.Sprintf("%s (%s)", p.pipeline.Commit.Author, p.pipeline.Commit.AuthorEmail),
		},
		{
			Name:  "Commit Message",
			Value: p.pipeline.Commit.Message,
		}}

	// If commit link is not null add to card
	if p.pipeline.Commit.Link != "" {
		facts = append(facts, MessageCardSectionFact{
			Name:  "Commit Link",
			Value: p.pipeline.Commit.Link,
		})
	}

	// If build has failed, change card details
	if p.settings.Status == "failure" {
		themeColor = "FF5733"
		facts = append(facts, MessageCardSectionFact{
			Name:  "Failed Build Steps",
			Value: strings.Join(p.pipeline.Build.FailedSteps, " "),
		})
	} else if p.settings.Status == "building" {
		themeColor = "002BFF"
	}

	// Create rich message card body
	card := MessageCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: themeColor,
		Summary:    p.pipeline.Repo.Slug,
		Sections: []MessageCardSection{{
			ActivityTitle:    p.pipeline.Repo.Slug,
			ActivitySubtitle: strings.ToUpper(p.settings.Status),
			ActivityImage:    "https://github.com/jdamata/drone-teams/raw/master/drone.png",
			Markdown:         false,
			Facts:            facts,
		}},
	}

	log.Info("Generated card: ", card)

	// MS teams webhook post
	jsonValue, _ := json.Marshal(card)
	_, err := http.Post(p.settings.Webhook, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Error("Failed to send request to teams webhook")
		return err
	}
	return nil
}
