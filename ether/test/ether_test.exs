defmodule EtherTest do
  use ExUnit.Case
  doctest Ether

  test "greets the world" do
    assert Ether.hello() == :world
  end
end
