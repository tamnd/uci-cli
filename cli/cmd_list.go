package cli

import (
	"github.com/spf13/cobra"
)

func (a *App) listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all UCI ML Repository datasets",
		RunE: func(cmd *cobra.Command, _ []string) error {
			limit := a.effectiveLimit(0)
			datasets, err := a.client.List(cmd.Context(), limit)
			if err != nil {
				return mapFetchErr(err)
			}
			return a.renderOrEmpty(datasets, len(datasets))
		},
	}
	return cmd
}
