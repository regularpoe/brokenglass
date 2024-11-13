defmodule Yurei.Command do
  def parse(line) do
    case String.split(line) do
      ["ls" | args] -> {:ok, {:ls, args}}
      _ -> {:error, :unknown_command}
    end
  end

  def run(_command)

  def run({:ls, args}) do
    dir = if Enum.empty?(args), do: ".", else: List.first(args)

    case System.cmd("ls", ["-lah", dir]) do
      {output, 0} -> {:ok, output}
      {error_msg, _} -> {:error, error_msg}
    end
  end
end
