defmodule Yuki.Router do
  use Plug.Router

  plug(:match)
  plug(:dispatch)

  get "/json" do
    json_data = :ets.lookup_element(:json_data, :data, 2)
    rsp(conn, 200, Jason.encode!(json_data))
  end

  match _ do
    error_response = %{"error" => "Not found"}
    rsp(conn, 404, Jason.encode!(error_response))
  end

  def update_json(new_json) do
    :ets.insert(:json_data, {:data, new_json})
  end

  defp rsp(conn, status, body) do
    conn
    |> put_resp_content_type("application/json")
    |> Plug.Conn.send_resp(status, body)
  end
end
