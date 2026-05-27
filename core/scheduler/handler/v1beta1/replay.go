package v1beta1

import (
	"github.com/google/uuid"
	"github.com/raystack/salt/log"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type ReplayService interface {
	CreateReplay(ctx context.Context, tenant tenant.Tenant, jobName scheduler.JobName, config *scheduler.ReplayConfig) (replayID uuid.UUID, err error)
	GetReplayList(ctx context.Context, projectName tenant.ProjectName) (replays []*scheduler.Replay, err error)
	GetReplayByID(ctx context.Context, replayID uuid.UUID) (replay *scheduler.ReplayWithRun, err error)
}

type ReplayHandler struct {
	l       log.Logger
	service ReplayService

	pb.UnimplementedReplayServiceServer
}

func (h ReplayHandler) Replay(ctx context.Context, req *pb.ReplayRequest) (*pb.ReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h ReplayHandler) ListReplay(ctx context.Context, req *pb.ListReplayRequest) (*pb.ListReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h ReplayHandler) GetReplay(ctx context.Context, req *pb.GetReplayRequest) (*pb.GetReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replayToProto(replay *scheduler.Replay) *pb.GetReplayResponse {
	_ = "STUB: not implemented"
	return nil
}

func parseJobConfig(jobConfig string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReplayHandler(l log.Logger, service ReplayService) *ReplayHandler {
	_ = "STUB: not implemented"
	return nil
}
