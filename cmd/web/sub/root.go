package sub

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	listenPort string
)

func init() {
	cmdRoot.Flags().StringVar(&listenPort, "port", ":4000", "port for listening(include colon as well)")
}

// env is a struct that is needed for dependency injection into handlers.
type env struct {
	router fasthttp.RequestHandler
}

// newEnv is a helper function to initialize(construct) env type.
func newEnv() (*env, error) {
	e := &env{}

	e.routes()

	return e, nil
}

var cmdRoot = &cobra.Command{
	Short: "s3-stream is a program to stream data to Amazon S3",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var op = "sub.RunE"

		// init an environment for dependecy injection
		e, err := newEnv()
		if err != nil {
			err = errors.Wrapf(err, "(%s): initializing env", op)
			return
		}

		// setup a server
		var server = &fasthttp.Server{
			Handler:      e.router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
		}

		// listen and serve connections
		errChan := make(chan error)
		go func(errChan chan<- error) {
			log.Infof("listening for incoming connections on: %s PORT", listenPort)
			if err := server.ListenAndServe(listenPort); err != nil && err != http.ErrServerClosed {
				errChan <- errors.Wrapf(err, "(%s): listen", op)
				return
			}
		}(errChan)

		// deal a CTRL + C signal
		log.Info("waiting for SIGINT signal")
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// catch channels
		select {
		case <-errChan:
			err = <-errChan
			break
		case <-quit:
			// shutdown gracefully
			log.Info("shutting down server gracefully")
			if err = server.Shutdown(); err != nil {
				err = errors.Wrapf(err, "(%s): server shutdown", op)
				break
			}
			break
		}

		return err
	},
}

// Execute is function that starts the programs
func Execute() error {
	return cmdRoot.Execute()
}
