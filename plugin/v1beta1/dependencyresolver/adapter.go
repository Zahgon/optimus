package dependencyresolver

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/plugins/v1beta1"
	"github.com/raystack/optimus/sdk/plugin"
)

func adaptConfigsToProto(c plugin.Configs) *pb.Configs { _ = "STUB: not implemented"; return nil }

func adaptConfigsFromProto(a *pb.Configs) plugin.Configs {
	_ = "STUB: not implemented"
	return *new(plugin.Configs)
}

func adaptAssetsToProto(a plugin.Assets) *pb.Assets { _ = "STUB: not implemented"; return nil }

func adaptAssetsFromProto(a *pb.Assets) plugin.Assets {
	_ = "STUB: not implemented"
	return *new(plugin.Assets)
}
