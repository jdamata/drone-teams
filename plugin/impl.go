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
	"strings"

	log "github.com/sirupsen/logrus"
)

// Settings for the plugin.
type Settings struct {
	Webhook string
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Verify the source url
	webhook := p.settings.Webhook
	if webhook == "" {
		return fmt.Errorf("no webhook endpoint provided")
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
			Name:  "Repo Link",
			Value: p.pipeline.Repo.Link,
		},
		{
			Name:  "Git Author",
			Value: p.pipeline.Commit.Author,
		},
		{
			Name:  "Commit Message",
			Value: p.pipeline.Commit.Message,
		},
		{
			Name:  "Commit Link",
			Value: p.pipeline.Commit.Link,
		}}

	// If build has failed, change card details
	if p.pipeline.Build.Status == "failure" {
		themeColor = "FF5733"
		facts = append(facts, MessageCardSectionFact{
			Name:  "Failed Build Steps",
			Value: strings.Join(p.pipeline.Build.FailedSteps, " "),
		})
	}

	// Create rich message card body
	card := MessageCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: themeColor,
		Summary:    p.pipeline.Repo.Slug,
		Sections: []MessageCardSection{{
			ActivityTitle: p.pipeline.Repo.Slug + " - " + strings.ToUpper(p.pipeline.Build.Status),
			// ActivitySubtitle: strings.ToUpper(p.pipeline.Build.Status),
			ActivityImage: "https://github.com/jdamata/drone-teams/raw/master/drone.png",
			Markdown:      false,
			Facts:         facts,
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
