defmodule Ether.MixProject do
  use Mix.Project

  def project do
    [
      app: :ether,
	ecto_repos: [Ether.Repo],
      version: "0.1.0",
      elixir: "~> 1.8",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      application: [
	:amqp, :amqp_client, :logger, :gossip,
	:jason, :poison, :json, :ecto_sql, :postgrex
	],
	mod: {Ether.Application, []}

    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
	#{:gossip, "~> 0.6"},
	{:ecto_sql, "~> 3.2"},
	{:postgrex, ">= 0.0.0"},
	{:poison, ">= 0.0.0"},
	{:json, ">= 0.0.0"},
	{:jason, ">= 0.0.0"},
	{:type_struct, ">= 0.0.0"},
	{:amqp_client, "~> 3.7.11"},
	{:amqp, "~> 1.3"}
      # {:dep_from_hexpm, "~> 0.3.0"},
      # {:dep_from_git, git: "https://github.com/elixir-lang/my_dep.git", tag: "0.1.0"}
    ]
  end
end
