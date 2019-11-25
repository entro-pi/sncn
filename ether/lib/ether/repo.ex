defmodule Ether.Repo do
use Ecto.Repo,
	otp_app: :ether,
	adapter: Ecto.Adapters.Postgres
end
