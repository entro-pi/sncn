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

defmodule Object do
        defstruct name: "A nyancat",
         longname: "A poptart kitten",
         vnum: 2, zone: "zem",
         ownerhash: "1234",
        worth: 1,
         slot: 0,
         x: 0,
        y: 0,
         owned: false
end

defmodule EquipmentItem do
	defstruct item: Object, number: 0
end

defmodule InventoryItem do
	defstruct item: Object, number: 0
end

defmodule InventoryBank do
          defstruct slotone: InventoryItem, slotoneamount: 0,
                        slottwo: InventoryItem, slottwoamount: 0,
                        slotthree: InventoryItem, slotthreeamount: 0,
                        slotfour: InventoryItem, slotfouramount: 0,
                        slotfive: InventoryItem, slotfiveamount: 0,
                        slotsix: InventoryItem, slotsixamount: 0,
                        slotseven: InventoryItem, slotsevenamount: 0,
                        sloteight: InventoryItem, sloteightamount: 0,
                        slotnine: InventoryItem, slotnineamount: 0,
                        slotten: InventoryItem, slottenamount: 0
end


defmodule EquipmentBank do
          defstruct slotone: EquipmentItem, slotoneamount: 0,
                        slottwo: EquipmentItem, slottwoamount: 0,
                        slotthree: EquipmentItem, slotthreeamount: 0,
                        slotfour: EquipmentItem, slotfouramount: 0,
                        slotfive: EquipmentItem, slotfiveamount: 0,
                        slotsix: EquipmentItem, slotsixamount: 0,
                        slotseven: EquipmentItem, slotsevenamount: 0,
                        sloteight: EquipmentItem, sloteightamount: 0,
                        slotnine: EquipmentItem, slotnineamount: 0,
                        slotten: EquipmentItem, slottenamount: 0
end



defmodule Player do	
	defstruct name: "dorp", hostname: "dev.snowcrash.network", title: "The Unknown",
		itembank: InventoryBank, equipmentbank: EquipmentBank,
		str: 1, int: 1, dex: 1,
		wis: 1, con: 1, cha: 1,
		inventory: [InventoryItem] , Equipped: [EquipmentItem] ,
		coreboard: "", plaincoreboard: "",
		currentroom: nil, playerhash: "",
		classes: [], level: 0.0,
		target: "", targetlong: "",
		eslotspell: nil, eslotskill: nil,
		qslotspell: nil, qslotskill: nil,
		tobuy: 0, bankaccount: nil,
		tarx: 0, tary: 0,
		oldx: 0, oldy: 0,
		cpu: "", coreshow: false,
		channels: [], battling: false,
		battlingmob: nil, session: "",
		attack: 0, defend: 0,
		slain: 0, hoarded: 0,
		maxrezz: 10, rezz: 10,
		maxtech: 5, tech: 5,
		fights: nil, won: 0, found: 0
	


	def init do
		play = %Player{}
	end
        def init(name) do
                play = %Player{name: name}
        end
	def get_value do
		{:ok, jsonPlayer} = JSON.encode([name: "dorp", hostname: "dev.snowcrash.network", title: "The Unknown",
                itembank: InventoryBank, equipmentbank: EquipmentBank,
                str: 1, int: 1, dex: 1,
                wis: 1, con: 1, cha: 1,
                inventory: [InventoryItem] , Equipped: [EquipmentItem] ,
                coreboard: "", plaincoreboard: "",
                currentroom: nil, playerhash: "",
                classes: [], level: 0.0,
                target: "", targetlong: "",
                eslotspell: nil, eslotskill: nil,
                qslotspell: nil, qslotskill: nil,
                tobuy: 0, bankaccount: nil,
                tarx: 0, tary: 0,
                oldx: 0, oldy: 0,
                cpu: "", coreshow: false,
                channels: [], battling: false,
                battlingmob: nil, session: "",
                attack: 0, defend: 0,
                slain: 0, hoarded: 0,
                maxrezz: 10, rezz: 10,
                maxtech: 5, tech: 5,
                fights: nil, won: 0, found: 0])
	end	


end


defmodule PlayerWatcher do
	def doThings do
		IO.puts("do things to the player!")
	end
	def start_connection do
		IO.puts("This is where the connection is started and passed along")
		IO.puts("spawn connection")
		{:ok, conn} = Mongo.start_link(database: "pfiles")
		conn
	end
	def add_p_file(conn, play) do
		{:ok, file} = JSON.encode(play)
		{:ok, new_file} = JSON.decode(file)
		data = BSON.encode(new_file)
		Mongo.insert_many(conn, "Players", data)
	end
	def save_p_file(conn, play) do
		{:ok, file} = JSON.encode(play)
		{:ok, new_file} = JSON.decode(file)
		{:ok, new_file_encoded} = BSON.encode(new_file)
		Mongo.upsert(conn, "Players", new_file)
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
		weasel = Player.init("weasel")
		conn = PlayerWatcher.start_connection
		PlayerWatcher.add_p_file(conn, weasel)
	end
end
