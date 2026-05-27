package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	jobRunInputCompileAssetsTimeout = time.Minute * 1

	taskInputDirectory = "in"
	unsubstitutedValue = "<no value>"

	typeEnvFileName    = ".env"
	typeSecretFileName = ".secret"

	ISOTimeLayout = time.RFC3339
)

type jobRunInputCommand struct {
	logger     log.Logger
	connection *connection.Insecure

	configFilePath string

	assetOutputDir string
	runType        string
	runName        string
	scheduledAt    string
	projectName    string
	host           string

	keysWithUnsubstitutedValue []string
}

// NewJobRunInputCommand gets compiled assets required for a job run
func NewJobRunInputCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (j *jobRunInputCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags

// Mandatory flags if config is not set

func (j *jobRunInputCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (j *jobRunInputCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// writeInstanceResponse fetches a JobRun from the store (eg, postgres)
// Based on the response, it builds assets like query, env and config
// for the Job Run which is saved into output files.
func (j *jobRunInputCommand) writeInstanceResponse(jobResponse *pb.JobRunInputResponse) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (j *jobRunInputCommand) writeJobResponseSecretToFile(
	jobResponse *pb.JobRunInputResponse, dirPath string,
) error {
	_ = "STUB: not implemented"
	// write all secrets into a file
	return nil
}

func (j *jobRunInputCommand) writeJobResponseEnvToFile(jobResponse *pb.JobRunInputResponse, dirPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *jobRunInputCommand) writeJobAssetsToFiles(
	jobResponse *pb.JobRunInputResponse, dirPath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *jobRunInputCommand) sendJobRunInputRequest(jobName string, jobScheduledTimeProto *timestamppb.Timestamp) (*pb.JobRunInputResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch Instance by calling the optimus API

func (j *jobRunInputCommand) getJobScheduledTimeProto() (*timestamppb.Timestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
