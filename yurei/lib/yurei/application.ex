defmodule Yurei.Application do
  require Logger
  use Application

  def start(_type, _args) do
    port = String.to_integer(System.get_env("PORT") || "9078")

    Logger.info("\yurei starting on port #{port}..")

    children = [
      {Task.Supervisor, name: Yurei.TaskSupervisor},
      Supervisor.child_spec({Task, fn -> Yurei.accept(port) end}, restart: :permanent)
    ]

    opts = [strategy: :one_for_one, name: Yurei.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
