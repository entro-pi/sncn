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
                creds = File.read!("postcreds")

                cred = creds |> String.split("\n")
                dataCred = Enum.at(cred, 0)
		hostCred = Enum.at(cred, 1)
                userCred = Enum.at(cred, 2)
                passCred = Enum.at(cred, 3)
		

		IO.puts("This is where the connection is started and passed along")
		IO.puts("spawn connection")
		{:ok, conn} = Postgrex.start_link(database: dataCred, hostname: hostCred, username: userCred, password: passCred)
	end
	def create_table_broadcasts(conn) do
		Postgrex.query!(conn, "CREATE TABLE broadcasts (message varchar(255));", [])
	end
	def create_table_pfile(conn) do
		Postgrex.query!(conn, "CREATE TABLE players (ID serial NOT NULL PRIMARY KEY, pfile json NOT NULL);", [])
	end
	def insert_broadcast(conn, msg) do
		Postgrex.query!(conn, "INSERT INTO broadcasts (message) VALUES ('"<>msg<>"')", [])
	end
	def get_all_broadcasts(conn) do
		Postgrex.query!(conn, "SELECT * FROM broadcasts", [])
	end
	def get_all_chars(conn) do
		{:ok, result} = Postgrex.query!(conn, "SELECT pfile FROM players", [])
		IO.puts Enum.at(result.row, 0).name
	end
	def add_p_file(conn, play) do
		{:ok, encodedPlayer} = JSON.encode(play)
		Postgrex.query!(conn, "INSERT INTO players (pfile) VALUES ('"<>encodedPlayer<>"');", [])
	end
	def find_p_files(conn) do
		query = "select array_to_json(array_agg(pfile))::text from players"
		{:ok, resultJson} = Postgrex.query(conn, query, [])
		resultString = List.flatten(hd(resultJson.rows))
		resultDecoded = Poison.decode!(resultString, as: [%Pfiles.Pfile{}])
		IO.puts(Enum.at(resultDecoded, 0).name)
	end
	def change_p_file(conn, change) do
		Postgrex.query!(conn, "UPDATE players SET name = '"<>change<>"'", [])
	end
	def find_p_file(conn, name) do
		Postgrex.query!(conn, "SELECT * FROM players WHERE name = '"<>name<>"'", [])
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
	def wait_for_messages_not_basic(channel) do
		receive do
			{:deliver, payload, meta} ->
			IO.puts " rabbit receivied #{payload}"
			{:ok, file} = File.open("../pot/broadcast", [:read, :write])
			{:ok, oldcontents } = File.read("../pot/broadcast")
			IO.binwrite(file, oldcontents <> "broadcast:" <> payload <> "\n")
			File.close(file)
			IO.puts "written to file"
			AMQP.Basic.ack(channel, meta.delivery_tag)
			wait_for_messages_not_basic(channel)
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
	def listenFanOut do
	
		creds = File.read!("creds")
		
		cred = creds |> String.split("\n")
		userCred = Enum.at(cred, 0)
		passCred = Enum.at(cred, 1)
		hostname = Enum.at(cred, 2)
		vhost = Enum.at(cred, 3)
		{:ok, connection} = AMQP.Connection.open(virtual_host: vhost, host: hostname, username: userCred, password: passCred)
		{:ok, channel} = AMQP.Channel.open(connection)
		AMQP.Exchange.declare(channel, "broadcasts", :fanout)
		{:ok, %{queue: queue_name}} = AMQP.Queue.declare(channel, "", auto_delete: false, durable: true, exclusive: false)
		AMQP.Queue.bind(channel, queue_name, "broadcasts")
		AMQP.Queue.bind(channel, queue_name, "broadcastsRight")
		AMQP.Queue.bind(channel, queue_name, "broadcastsLeft")
		AMQP.Basic.consume(channel, queue_name, nil, no_ack: false)

		IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"
		spawn(Listener.wait_for_messages(channel))
		spawn(Listener.tick(channel))
	end
	def listenFanOutLeft do
	
		creds = File.read!("creds")
		
		cred = creds |> String.split("\n")
		userCred = Enum.at(cred, 0)
		passCred = Enum.at(cred, 1)
		hostname = Enum.at(cred, 2)
		vhost = Enum.at(cred, 3)
		{:ok, connection} = AMQP.Connection.open(virtual_host: vhost, host: hostname, username: userCred, password: passCred)
		{:ok, channel} = AMQP.Channel.open(connection)
		AMQP.Exchange.declare(channel, "broadcastsLeft", :direct)
		AMQP.Queue.declare(channel, "left", auto_delete: false, durable: true, exclusive: false)
		AMQP.Queue.bind(channel, "left", "broadcastsLeft")
		AMQP.Basic.consume(channel, "left", nil, no_ack: false)

		IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"
		spawn(Listener.wait_for_messages(channel))
		spawn(Listener.tick(channel))
	end
	def listenFanOutRight do
	
		creds = File.read!("creds")
		
		cred = creds |> String.split("\n")
		userCred = Enum.at(cred, 0)
		passCred = Enum.at(cred, 1)
		hostname = Enum.at(cred, 2)
		vhost = Enum.at(cred, 3)
		{:ok, connection} = AMQP.Connection.open(virtual_host: vhost, host: hostname, username: userCred, password: passCred)
		{:ok, channel} = AMQP.Channel.open(connection)
		AMQP.Exchange.declare(channel, "broadcastsRight", :direct)
		AMQP.Queue.declare(channel, "right", auto_delete: false, durable: true, exclusive: false)
		AMQP.Queue.bind(channel, "right", "broadcastsRight")
		AMQP.Basic.consume(channel, "right", nil, no_ack: false)

		IO.puts " [*] Waiting for messages. To exit press CTRL+C, CTRL+C"
		spawn(Listener.wait_for_messages(channel))
		spawn(Listener.tick(channel))
	end
end

defmodule Test do
	def test_get_all_chars do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.get_all_chars(conn)
	end
	def test_define_types do
		Postgrex.Types.define(Ether.PostgrexTypes, [%Pfiles.Pfile{}], [])
		{:ok, conn} = PlayerWatcher.start_connection(types: %Pfiles.Pfile{})
	end
	def test_create_pfile_table do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.create_table_pfile(conn)
	end
	def test_create_broadcast_table do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.create_table_broadcasts(conn)		
	end
	def test_find_p_files do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.find_p_files(conn)
	end
	def test_insert_broadcast(msg) do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.insert_broadcast(conn, msg)
	end
	def test_find_broadcasts do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.get_all_broadcasts(conn)
	end
	def test_upload do
		weasel = %Pfiles.Pfile{name: "Weasel"}
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, weasel)
		
	end
	def test_upload(playername) do
		weasel = %Pfiles.Pfile{name: playername}
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, weasel)
	end
	def test(name) do
		play = %Pfiles.Pfile{name: name, oldx: 0, oldy: 0, tarx: 0, tary: 0, maxrezz: 20, rezz: 10, maxtech: 10, tech: 5}
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, play)
	end
	def test_change(newname) do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.change_p_file(conn, newname)
	end
	def test_find(name) do
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.find_p_file(conn, name)
	end

	def test do
		weasel = %Pfiles.Pfile{name: "Weasel", oldx: 0, oldy: 0, tarx: 0, tary: 0, maxrezz: 20, rezz: 10, maxtech: 10, tech: 5}
		{:ok, conn} = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, weasel)
	end
end

defmodule Object do
	use Ecto.Schema
        schema "object" do
	field :name, :string
        field :longname, :string
        field :vnum, :integer
	field :zone, :string
        field :ownerhash, :string
        field :worth, :integer
        field :slot, :integer
        field :x, :integer
        field :y, :integer
        field :owned, :boolean
	end
end

defmodule EquipmentItem do
	use Ecto.Schema
	schema "equipmentitem" do
	belongs_to :item, Object
	field :number, :integer
        end
end

defmodule InventoryItem do
	use Ecto.Schema
        schema "inventoryitem"do
	belongs_to :item, Object
	field :number, :integer
	end
end
defmodule InventoryBank do
        use Ecto.Schema
        schema "inventorybank" do
         belongs_to :slotone, InventoryItem
         field :slotoneamount, :integer
         belongs_to :slottwo, InventoryItem
         field :slottwoamount, :integer
         belongs_to :slotthree, InventoryItem
         field :slotthreeamount, :integer
         belongs_to :slotfour, InventoryItem
         field :slotfouramount, :integer
         belongs_to :slotfive, InventoryItem
         field :slotfiveamount, :integer
         belongs_to :slotsix, InventoryItem
         field :slotsixamount, :integer
         belongs_to :slotseven, InventoryItem
         field :slotsevenamount, :integer
         belongs_to :sloteight, InventoryItem
         field :sloteightamount, :integer
         belongs_to :slotnine, InventoryItem
         field :slotnineamount, :integer
         belongs_to :slotten, InventoryItem
         field :slottenamount, :integer
end
end
defmodule EquipmentBank do
        use Ecto.Schema
        schema "equipmentbank" do
         belongs_to :slotone, EquipmentItem
         field :slotoneamount, :integer
         belongs_to :slottwo, EquipmentItem
         field :slottwoamount, :integer
         belongs_to :slotthree, EquipmentItem
         field :slotthreeamount, :integer
         belongs_to :slotfour, EquipmentItem
         field :slotfouramount, :integer
         belongs_to :slotfive, EquipmentItem
         field :slotfiveamount, :integer
         belongs_to :slotsix, EquipmentItem
         field :slotsixamount, :integer
         belongs_to :slotseven, EquipmentItem
         field :slotsevenamount, :integer
         belongs_to :sloteight, EquipmentItem
         field :sloteightamount, :integer
         belongs_to :slotnine, EquipmentItem
         field :slotnineamount, :integer
         belongs_to :slotten, EquipmentItem
         field :slottenamount, :integer
end
end

defmodule Pfiles.Pfile do
	use Ecto.Schema
	schema "pfiles" do
        field :name, :string
	 field :hostname, :string
	 field :title, :string
                belongs_to :itembank, InventoryBank
	belongs_to  :equipmentbank, EquipmentBank
                field :str, :integer
		field :int, :integer
		field :dex, :integer
                field :wis, :integer
		field :con, :integer
		field :cha, :integer
                belongs_to :inventory, Elixir.List
		belongs_to :Equipped, Elixir.List
                field :coreboard, :string
		field :plaincoreboard, :string
                field :currentroom, :string
		field :playerhash, :string
                field :classes, :string
		field :level, :integer 
                field :target, :string
		field  :targetlong, :string
                field :eslotspell, :string
		field :eslotskill, :string
                field :qslotspell, :string
		field :qslotskill, :string
                field :tobuy, :integer
		field :bankaccount, :string
                field :tarx, :integer
		field :tary, :integer
                field :oldx, :integer
		field :oldy, :integer
                field :cpu, :string
		field :coreshow, :boolean
                field :channels, :string
		field :battling, :boolean
                field :battlingmob, :string
		field :session, :string
                field :attack, :integer
		field :defend, :integer
                field :slain, :integer
		field :hoarded, :integer
                field :maxrezz, :integer
		field :rezz, :integer
                field :maxtech, :integer
		field :tech, :integer
                field :fights, :string
		field :won, :integer
		field :found, :integer
	end

end
