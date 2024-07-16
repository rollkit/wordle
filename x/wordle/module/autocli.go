package wordle

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "wordle/api/wordle/wordle"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "SubmitWordle",
					Use:            "submit-wordle [word]",
					Short:          "Send a submit-wordle tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "word"}},
				},
				{
					RpcMethod:      "SubmitGuess",
					Use:            "submit-guess [word]",
					Short:          "Send a submit-guess tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "word"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
