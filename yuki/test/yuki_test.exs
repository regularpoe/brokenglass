defmodule YukiTest do
  use ExUnit.Case
  doctest Yuki

  test "greets the world" do
    assert Yuki.hello() == :world
  end
end
