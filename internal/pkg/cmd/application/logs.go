/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package application

import (
	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func NewLogsCommand(cliContext runtime.CliContext) *cobra.Command {

	cmd := &cobra.Command{
		Use:   cmdLogs,
		Short: "Manage application logs",
		Args:  cobra.MaximumNArgs(1),
	}
	cmd.AddCommand(NewShowLogsCommand(cliContext))
	return cmd
}
