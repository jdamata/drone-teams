// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"github.com/jdamata/drone-teams/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "webhook",
			Usage:       "MS teams connector webhook endpoint",
			EnvVars:     []string{"PLUGIN_WEBHOOK"},
			Destination: &settings.Webhook,
		},
		&cli.StringFlag{
			Name:        "status",
			Usage:       "Overwrite the status value",
			EnvVars:     []string{"PLUGIN_STATUS"},
			Destination: &settings.Status,
		},
	}
}
