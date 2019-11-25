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
	import Ecto.Changeset
	def changeset(player, params \\ %{}) do
		player
		|>cast(params, [:name])
		|>validate_required([:name])
	end

end
