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

defmodule PlayerWatcher do
	def doThings do
		IO.puts("do things to the player!")
	end
	def start_connection do
		IO.puts("This is where the connection is started and passed along")
		IO.puts("spawn connection")
		{:ok, conn} = Postgrex.start_link(database: "pfiles", hostname: "localhost", username: "postgres")
	end
	def add_p_file(conn, play) do
		Postgrex.query!(conn, "INSERT INTO pfiles (name) VALUES ('weasel')", [])
	end
	def save_p_file(conn, play) do
	end
	def lookup(conn, playerHash) do
		IO.puts("occasionally the poller does stuff")
		cursor = Mongo.find(conn, "Players", %{"$and" =>[%{name: playerHash}]})
		|> Enum.to_list()
	
	end
	def watch(conn, playerHash) do
		IO.puts("This is where the server sets up the player's connection to the database")
		PlayerWatcher.lookup(conn, playerHash)
	end
	def playerTick(conn, play) do
		IO.puts("do things on a tick, like save")		
	end

	def unwatch do
		IO.puts("Save the pFile, and close the connection")
		
	end
end



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
	def tick(channel) do
		determine = [0, 1, 2, 3, 4, 5]
		Enum.random(determine)
		|> Kernel.*(1000)
		|> :timer.sleep
		AMQP.Basic.publish(channel, "", "input", "!:::tick:::!")
		IO.puts "tick!"
		
		Listener.tick(channel)
	end
	def listen do
	
		creds = File.read!("creds")
		
		cred = creds |> String.split("\n")
		userCred = Enum.at(cred, 0)
		passCred = Enum.at(cred, 1)
		hostname = Enum.at(cred, 2)
		vhost = Enum.at(cred, 3)
		{ok, connection} = AMQP.Connection.open(virtual_host: vhost, host: hostname, username: userCred, password: passCred)
		{:ok, channel} = AMQP.Channel.open(connection)

		AMQP.Queue.declare(channel, "input", auto_delete: true, durable: true)
		AMQP.Basic.qos(channel, prefetch_count: 1)
		AMQP.Basic.consume(channel, "input", nil, no_ack: false, persistent: true)

		IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"
		spawn(Listener.wait_for_messages(channel))
		spawn(Listener.tick(channel))
	end
end

defmodule Test do
	def test do
		weasel = %Pfiles.Pfile{name: "Weasel"}
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, weasel)
	end
end
