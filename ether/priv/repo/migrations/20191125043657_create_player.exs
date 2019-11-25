defmodule Ether.Repo.Migrations.CreatePlayer do
  use Ecto.Migration
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


end


  def change do
	create table(:players) do
		add :player, %Player{}
	end
	end
end
