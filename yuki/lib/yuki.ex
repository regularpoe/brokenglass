defmodule Yuki do
  require Logger
  use Application

  def start(_type, _args) do
    port = String.to_integer(System.get_env("PORT") || "4001")

    :ets.new(:json_data, [:public, :named_table])

    :ets.insert(:json_data, {:data, %{"message" => "Hello, world!", "status" => "initial"}})

    Logger.info("Starting yuki on port #{port}..")

    children = [
      {Plug.Cowboy, scheme: :http, plug: Yuki.Router, options: [port: port]}
    ]

    opts = [strategy: :one_for_one, name: Yuki.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
