defmodule Pfiles.Repo do
  use Ecto.Repo,
    otp_app: :ether,
    adapter: Ecto.Adapters.Postgres
end
