defmodule Heart do
  	def print(lang) do
		IO.puts("Hello!"<>lang)
	end
	def notify(mess, mes) do
		{:ok, file} = File.open("../pot/broadcast", [:read, :write])
		{:ok, oldcontents } = File.read("../pot/broadcast")
		IO.binwrite(file, [mess, "DOOT", mes])
		File.close(file)
	end
	def notify(mess) do
		{:ok, file} = File.open("../pot/broadcast", [:read, :write])
		{:ok, oldcontents } = File.read("../pot/broadcast")
		IO.binwrite(file, oldcontents <> "broadcast:" <> mess <> "\n")
		File.close(file)
	end
	def see do
		File.read("../pot/broadcast")
	end

end
Application.ensure_started(:amqp_client)

defmodule Listener do
	def wait_for_messages(channel) do
		receive do
			{:basic_deliver, payload, meta} ->
			IO.puts " rabbit receivied #{payload}"
			{:ok, file} = File.open("../pot/broadcast", [:read, :write])
			{:ok, oldcontents } = File.read("../pot/broadcast")
			IO.binwrite(file, oldcontents <> "broadcast:" <> payload <> "\n")
			File.close(file)
			IO.puts "written to file"
			AMQP.Basic.ack(channel, meta.delivery_tag)
			wait_for_messages(channel)
		end
	end
	def wait_for_dotdot_messages(channel) do
		receive do
			{:basic_deliver, payload, meta} ->
			IO.puts " [x] Received #{payload}"
			payload
			|> to_char_list
			|> Enum.count(fn x -> x == ?. end)
			|> Kernel.*(1000)
			|> :timer.sleep
			IO.puts " [x] Done."
			AMQP.Basic.ack(channel, meta.delivery_tag)
			wait_for_dotdot_messages(channel)
		end
	end
	def listen do
	
		creds = File.read!("creds")
		
		cred = creds |> String.split("\n")
		userCred = Enum.at(cred, 0)
		passCred = Enum.at(cred, 1)
		hostname = Enum.at(cred, 2)
		{ok, connection} = AMQP.Connection.open(host: hostname, username: userCred, password: passCred)
		{:ok, channel} = AMQP.Channel.open(connection)

		AMQP.Queue.declare(channel, "input", auto_delete: true, durable: true)
		AMQP.Basic.qos(channel, prefetch_count: 1)
		AMQP.Basic.consume(channel, "input", nil, no_ack: false, persistent: true)

		IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"

		Listener.wait_for_messages(channel)
	end
end
