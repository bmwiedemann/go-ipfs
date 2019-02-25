package cmdenv

import (
	"fmt"
	"strings"

	"github.com/ipfs/go-ipfs/commands"
	"github.com/ipfs/go-ipfs/core"

	cmds "gx/ipfs/QmQkW9fnCsg9SLHdViiAh6qfBppodsPZVpU92dZLqYtEfs/go-ipfs-cmds"
	coreiface "gx/ipfs/QmSU6tTh2ng2WnqfaxPggpCvc5qYBWeGgJiQzEWQhotNhi/interface-go-ipfs-core"
	options "gx/ipfs/QmSU6tTh2ng2WnqfaxPggpCvc5qYBWeGgJiQzEWQhotNhi/interface-go-ipfs-core/options"
	config "gx/ipfs/QmUAuYuiafnJRZxDDX7MuruMNsicYNuyub5vUeAcupUBNs/go-ipfs-config"
	logging "gx/ipfs/QmbkT7eMTyXfpeyB3ZMxxcxg7XH8t6uXp49jqzz4HB7BGF/go-log"
)

var log = logging.Logger("core/commands/cmdenv")

// GetNode extracts the node from the environment.
func GetNode(env interface{}) (*core.IpfsNode, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.GetNode()
}

// GetApi extracts CoreAPI instance from the environment.
func GetApi(env cmds.Environment, req *cmds.Request) (coreiface.CoreAPI, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	offline, _ := req.Options["offline"].(bool)
	if !offline {
		offline, _ = req.Options["local"].(bool)
		if offline {
			log.Errorf("Command '%s', --local is deprecated, use --offline instead", strings.Join(req.Path, " "))
		}
	}
	api, err := ctx.GetAPI()
	if err != nil {
		return nil, err
	}
	if offline {
		return api.WithOptions(options.Api.Offline(offline))
	}

	return api, nil
}

// GetConfig extracts the config from the environment.
func GetConfig(env cmds.Environment) (*config.Config, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.GetConfig()
}

// GetConfigRoot extracts the config root from the environment
func GetConfigRoot(env cmds.Environment) (string, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return "", fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.ConfigRoot, nil
}
