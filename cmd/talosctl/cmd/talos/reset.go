// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package talos

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/talos-systems/talos/cmd/talosctl/pkg/client"
)

var (
	graceful bool
	reboot   bool
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a node",
	Long:  ``,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return WithClient(func(ctx context.Context, c *client.Client) error {
			if err := c.Reset(ctx, graceful, reboot); err != nil {
				return fmt.Errorf("error executing reset: %s", err)
			}

			return nil
		})
	},
}

func init() {
	resetCmd.Flags().BoolVar(&graceful, "graceful", true, "if true, attempt to cordon/drain node and leave etcd (if applicable)")
	resetCmd.Flags().BoolVar(&reboot, "reboot", false, "if true, reboot the node after resetting instead of shutting down")
	addCommand(resetCmd)
}
