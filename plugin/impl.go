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

	// Set card color to green or red
	themeColor := "96FF33"
	if p.pipeline.Build.Status == "failure" {
		themeColor = "FF5733"
	}

	// Create rich message card body
	card := MessageCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: themeColor,
		Summary:    p.pipeline.Build.Status,
		Sections: []MessageCardSection{{
			ActivityTitle:    "Build status -> " + p.pipeline.Build.Status,
			ActivitySubtitle: "Repo Name -> " + p.pipeline.Repo.Link,
			ActivityImage:    "https://github.com/jdamata/drone-teams/raw/master/drone.png",
			Markdown:         true,
			Facts: []MessageCardSectionFact{
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
				},
			},
		}},
	}
	log.Info("Generated card: ", card)

	// Make MS teams webhook post
	jsonValue, _ := json.Marshal(card)
	_, err := http.Post(p.settings.Webhook, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Error("Failed to send request to teams webhook")
		return err
	}
	return nil
}
