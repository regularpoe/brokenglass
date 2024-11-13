defmodule Nocturno.CLI do
  def main(args) do
    {opts, _, _} = OptionParser.parse(args, switches: [config: :string], aliases: [c: :config])

    case opts[:config] do
      nil ->
        IO.puts("Please provide a config file with the -c or --config flag")

      config_path ->
        case Nocturno.read_config(config_path) do
          {:ok, paths, notifications} ->
            {:ok, _pid} = Nocturno.start_link(paths, notifications)
            IO.puts("\nnocturno started...")
            :timer.sleep(:infinity)

          {:error, reason} ->
            IO.puts("Error: #{reason}")
        end
    end
  end
end
