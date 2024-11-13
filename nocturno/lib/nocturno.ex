defmodule Nocturno do
  use GenServer

  def read_config(config_path) do
    case File.read(config_path) do
      {:ok, content} ->
        case Jason.decode(content) do
          {:ok, %{"paths" => paths, "notifications" => notifications}} ->
            {:ok, paths, notifications}

          _ ->
            {:error, "Invalid configuration file format"}
        end

      {:error, reason} ->
        {:error, "Failed to read config file: #{reason}"}
    end
  end

  def start_link(paths, notifications) do
    GenServer.start_link(__MODULE__, {paths, notifications}, name: __MODULE__)
  end

  @impl true
  def init({paths, notifications}) do
    {:ok, watcher_pid} = FileSystem.start_link(dirs: paths)
    FileSystem.subscribe(watcher_pid)

    {:ok, %{notifications: notifications}}
  end

  @impl true
  def handle_info({:file_event, _pid, {path, events}}, state) do
    IO.puts("File changed: #{path}, events: #{inspect(events)}")

    if state[:notifications] do
      send_notification(path, events)
    end

    {:noreply, state}
  end

  defp send_notification(path, _events) do
    System.cmd("notify-send", ["File changed", "File #{path} has been modified"])
  end
end
