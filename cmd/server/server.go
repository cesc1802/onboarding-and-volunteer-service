package server

import (
	"context"

	"fmt"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature"
	"github.com/cesc1802/share-module/config"
	"github.com/cesc1802/share-module/system"
	"github.com/spf13/cobra"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono system.Service) error {
	feature.RegisterHandlerV1(mono)
	return nil
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starting HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello World")

		cfg, err := config.LoadAppConfig(".")
		if err != nil {
			return err
		}
		sys := system.New(cfg, cmd.Parent().Name())

		if err := Root(sys.Waiter().Context(), sys); err != nil {
			return err
		}

		sys.Waiter().Add(
			sys.WaitForWeb)

		return sys.Waiter().Wait()
	},
}

func RegisterServer(root *cobra.Command) {
	root.AddCommand(serverCmd)
}
