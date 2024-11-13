defmodule NocturnoTest do
  use ExUnit.Case
  doctest Nocturno

  test "greets the world" do
    assert Nocturno.hello() == :world
  end
end
