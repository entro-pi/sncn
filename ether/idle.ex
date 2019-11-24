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

defmodule Player do
	use TypeStruct

	defstruct Object,
	 name: String.t \\ "A nyancat",
	 longname: String.t \\ "A poptart kitten",
	 vnum: integer \\ 2, zone: String.t \\ "zem",
	 ownerhash: String.t \\ "1234",
	worth: integer \\ 1,
	 slot: integer \\ 0,
	 x: integer \\ 0, 
	y: integer \\ 0,
	 owned: bool \\ false

	defstruct EquipmentItem, item: Object, number: integer \\ 0
	defstruct InventoryItem, item: Object, number: integer \\ 0
	
	defstruct InventoryBank, slotone: InventoryItem, slotoneamount: integer \\ 0,
			slottwo: InventoryItem, slottwoamount: integer \\ 0,
			slotthree: InventoryItem, slotthreeamount: integer \\ 0,
			slotfour: InventoryItem, slotfouramount: integer \\ 0,
			slotfive: InventoryItem, slotfiveamount: integer \\ 0,
			slotsix: InventoryItem, slotsixamount: integer \\ 0,
			slotseven: InventoryItem, slotsevenamount: integer \\ 0,
			sloteight: InventoryItem, sloteightamount: integer \\ 0,
			slotnine: InventoryItem, slotnineamount: integer \\ 0,
			slotten: InventoryItem, slottenamount: integer \\ 0
        defstruct EquipmentBank, slotone: EquipmentItem, slotoneamount: integer \\ 0,
                	slottwo: EquipmentItem, slottwoamount: integer \\ 0,
                	slotthree: EquipmentItem, slotthreeamount: integer \\ 0,
                	slotfour: EquipmentItem, slotfouramount: integer \\ 0,
                	slotfive: EquipmentItem, slotfiveamount: integer \\ 0,
                	slotsix: EquipmentItem, slotsixamount: integer \\ 0,
                	slotseven: EquipmentItem, slotsevenamount: integer \\ 0,
                	sloteight: EquipmentItem, sloteightamount: integer \\ 0,
                	slotnine: EquipmentItem, slotnineamount: integer \\ 0,
                	slotten: EquipmentItem, slottenamount: integer \\ 0
	defstruct Play, name: String.t \\ "dorp", hostname: String.t \\ "dev.snowcrash.network", title: String.t \\ "The Unknown",
		itembank: Player.InventoryBank, equipmentbank: Player.EquipmentBank,
		str: integer \\ 1, int: integer \\ 1, dex: integer \\ 1,
		wis: integer \\ 1, con: integer \\ 1, cha: integer \\ 1,
		inventory: [InventoryItem] \\ [], Equipped: [EquipmentItem] \\ [],
		coreboard: String.t \\ "", plaincoreboard: String.t \\ "",
		currentroom: Space \\ nil, playerhash: String.t \\ "",
		classes: [Classes] \\ [], level: Float.t \\ 0.0,
		target: String.t \\ "", targetlong: String.t \\ "",
		eslotspell: Spell \\ nil, eslotskill: Skill \\ nil,
		qslotspell: Spell \\ nil, qslotskill: Skill \\ nil,
		tobuy: integer \\ 0, bankaccount: Account \\ nil,
		tarx: integer \\ 0, tary: integer \\ 0,
		oldx: integer \\ 0, oldy: integer \\ 0,
		cpu: String.t \\ "", coreshow: bool \\ false,
		channels: [String.t] \\ [], battling: bool \\ false,
		battlingmob: Mobile \\ nil, session: String.t \\ "",
		attack: integer \\ 0, defend: integer \\ 0,
		slain: integer \\ 0, hoarded: integer \\ 0,
		maxrezz: integer \\ 10, rezz: integer \\ 10,
		maxtech: integer \\ 5, tech: integer \\ 5,
		fights: Fight \\ nil, won: integer \\ 0, found: integer \\ 0

	def init do
		nyan = %Object{}
		ibank = %InventoryBank{slotone: nyan, slottwo: nyan, slotthree: nyan, slotfour: nyan, slotfive: nyan,
					slotsix: nyan, slotseven: nyan, sloteight: nyan, slotnine: nyan, slotten: nyan}
		
                eqbank = %EquipmentBank{slotone: nyan, slottwo: nyan, slotthree: nyan, slotfour: nyan, slotfive: nyan,
                                        slotsix: nyan, slotseven: nyan, sloteight: nyan, slotnine: nyan, slotten: nyan}
		play = %Player.Play{itembank: ibank, equipmentbank: eqbank}
	end
        def init(name) do
                nyan = %Object{}
                ibank = %InventoryBank{slotone: nyan, slottwo: nyan, slotthree: nyan, slotfour: nyan, slotfive: nyan,
                                        slotsix: nyan, slotseven: nyan, sloteight: nyan, slotnine: nyan, slotten: nyan}

                eqbank = %EquipmentBank{slotone: nyan, slottwo: nyan, slotthree: nyan, slotfour: nyan, slotfive: nyan,
                                        slotsix: nyan, slotseven: nyan, sloteight: nyan, slotnine: nyan, slotten: nyan}
                play = %Player.Play{name: name, itembank: ibank, equipmentbank: eqbank}
        end
	


end


defmodule PlayerWatcher do
	def doThings do
		IO.puts("do things to the player!")
	end
	def start_connection do
		IO.puts("This is where the connection is started and passed along")
		IO.puts("spawn connection")
		{:ok, conn} = Mongo.start_link(name: :mongo, database: "pfiles", pool_size: 2)
		PlayerWatcher.watch(conn, "Weasel")
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
	def unwatch do
		IO.puts("Save the pFile, and close the connection")
		
	end
end
PlayerWatcher.start_connection



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
