// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Settings for the plugin.
type Settings struct {
	Webhook string
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Validation of the settings.
	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {

	// Set card color to green or red
	themeColor := "96FF33"
	if p.pipeline.Build.Status == "Failed" {
		themeColor = "FF5733"
	}

	// Create rich message card body
	card := MessageCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: themeColor,
		Summary:    p.pipeline.Build.Status,
		Sections: []MessageCardSection{{
			ActivityTitle:    p.pipeline.Build.Action,
			ActivitySubtitle: p.pipeline.Repo.Name,
			ActivityImage:    "https://github.com/jdamata/drone-teams/raw/master/drone.png",
			Markdown:         true,
			Facts: []MessageCardSectionFact{
				{
					Name:  "Author",
					Value: p.pipeline.Commit.Author,
				},
				{
					Name:  "Commit",
					Value: p.pipeline.Commit.Message,
				},
				{
					Name:  "Link",
					Value: p.pipeline.Commit.Link,
				},
			},
		}},
	}

	// Make MS teams webhook post
	jsonValue, _ := json.Marshal(card)
	log.Info(p.settings.Webhook)
	_, err := http.Post(p.settings.Webhook, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Error("Failed to send request to teams webhook")
		return err
	}
	return nil
}
