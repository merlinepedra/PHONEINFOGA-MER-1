package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/sundowndev/phoneinfoga/v2/web"
)

var httpPort int
var disableClient bool

func init() {
	// Register command
	rootCmd.AddCommand(serveCmd)

	// Register flags
	serveCmd.PersistentFlags().IntVarP(&httpPort, "port", "p", 5000, "HTTP port")
	serveCmd.PersistentFlags().BoolVar(&disableClient, "no-client", false, "Disable web client (REST API only)")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve web client",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()

		_, err := web.Serve(router, disableClient)
		if err != nil {
			log.Fatal(err)
		}

		httpPort := ":" + strconv.Itoa(httpPort)

		srv := &http.Server{
			Addr:    httpPort,
			Handler: router,
		}

		fmt.Printf("Listening on %s\n", httpPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	},
}
