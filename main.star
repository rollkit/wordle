# This Kurtosis package spins up a wordle rollup that connects to a DA node

# Import the local da kurtosis package
da_node = import_module("github.com/rollkit/local-da/main.star@v0.3.0")

def run(plan):
    # Start the DA node
    da_address = da_node.run(
        plan,
    )
    plan.print("connecting to da layer via {0}".format(da_address))

    # Define the wordle start command
    wordle_start_cmd = [
        "rollkit",
        "start",
        "--rollkit.aggregator",
        "--rollkit.da_address {0}".format(da_address),
    ]
    # Define the jsonrpc ports
    wordle_ports = {
        "jsonrpc": PortSpec(
            number=26657, transport_protocol="TCP", application_protocol="http"
        ),
    }
    # Start the wordle chain
    wordle = plan.add_service(
        name="wordle",
        config=ServiceConfig(
            # Locally built wordle image
            image="wordle",
            cmd=["/bin/sh", "-c", " ".join(wordle_start_cmd)],
            ports=wordle_ports,
            public_ports=wordle_ports,
            ready_conditions=ReadyCondition(
                recipe=ExecRecipe(
                    command=["rollkit", "status"],
                    extract={
                        "output": "fromjson | .node_info.network",
                    },
                ),
                field="extract.output",
                assertion="==",
                target_value="wordle",
                interval="1s",
                timeout="1m",
            ),
        ),
    )
