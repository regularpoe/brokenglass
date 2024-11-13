defmodule YureiTest do
  use ExUnit.Case
  doctest Yurei

  test "greets the world" do
    assert Yurei.hello() == :world
  end
end
