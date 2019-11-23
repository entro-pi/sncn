defmodule Heart do
  	def print(lang) do
		IO.puts("Hello!"<>lang)
	end
	def notify(mess, mes) do
		{:ok, file} = File.open("../pot/broadcasts", [:read, :write])
		{:ok, oldcontents } = File.read("../pot/broadcasts")
		IO.binwrite(file, [mess, "DOOT", mes])
		File.close(file)
	end
	def notify(mess) do
		{:ok, file} = File.open("../pot/broadcasts", [:read, :write])
		{:ok, oldcontents } = File.read("../pot/broadcasts")
		IO.binwrite(file, oldcontents <> mess)
		File.close(file)
	end
	def see do
		File.read("../pot/broadcasts")
	end

end
Application.ensure_started(:amqp_client)

defmodule Listen do
	def wait_for_messages do
		receive do
			{:basic_deliver, payload, _meta} ->
			IO.puts " [x] Received #{payload}"

			wait_for_messages()
		end
	end
end


{:ok, connection} = AMQP.Connection.open("amqp://guest:guest@dev.snowcrash.network")
{:ok, channel} = AMQP.Channel.open(connection)

AMQP.Queue.declare(channel, "hello")
AMQP.Basic.consume(channel, "hello", nil, no_ack: true)

IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"

Listen.wait_for_messages()