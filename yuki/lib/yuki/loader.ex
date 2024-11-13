defmodule Yuki.Loader do
  def load_json_file(path) do
    case File.read(path) do
      {:ok, content} ->
        case Jason.decode(content) do
          {:ok, json_data} ->
            Yuki.Router.update_json(json_data)
            {:ok, "JSON updated successfully"}

          {:error, _reason} ->
            {:error, "Failed to parse JSON"}
        end

      {:error, reason} ->
        {:error, "Failed to read file: #{reason}"}
    end
  end
end
