defmodule Pfiles.Repo do
  use Ecto.Repo,
    otp_app: :idlesnow,
    adapter: Ecto.Adapters.Postgres
end
