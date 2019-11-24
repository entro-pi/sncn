package main

import (
	"github.com/SolarLune/dngn"

	"time"
)


type StatusPayload struct {
	Game string
	Players []string
}

type Status struct {
	Event string
	Ref string
	Payload StatusPayload
}


type SignOutPayload struct {
	Name string
	Game string
}

type SignOut struct {
	Event string
	Ref string
	Payload SignOutPayload
}

type SignIn struct {
	Event string
	Ref string
	Payload SignInPayload
}
type InventoryItem struct {
	Item Object
	Number int
}
type SignInPayload struct {
	Name string
	Game string
}
type Object struct {
	Name string
	LongName string
	Vnum int
	Zone string
	OwnerHash string
	Worth int
	Slot int
	X int
	Y int
	Owned bool
}

type BroadcastPayload struct {
  Channel string
  Message string
  Game string
  Name string
	Row int
	Col int
	Selected bool
	BigMessage string
	CoreBoard string
	Fights Fight
	PlainCoreBoard string
	CPU string
	ID int
	Transaction OnlineTransaction
	Store OnlineStore
}
type Butler struct {
	Employer string
	Funds Account
	ToBuy Object
}
type OnlineStore struct {
	Owner string
	Float float64
	Inventory []OnlineTransaction
}

type OnlineTransaction struct {
	ItemHash string
	Item Object
	Sold bool
	To Account
	Price float64
}
type Bank struct {
	Clientele Client
	Owner string
}
type Client struct {
		User string
		TotalAmount float64
		Accounts []float64
}
type Account struct {
	Owner string
	Income OnlineStore
	Amount float64
}

type Broadcast struct {
    Event string
    Ref string
    Payload BroadcastPayload
}

type Descriptions struct {
	BATTLESPAM int
	ROOMDESC int
	PLAYERDESC int
	ROOMTITLE int
}
type Chat struct {
	User Player
	Message string
	Time time.Time
}
type Space struct{
	Room dngn.Room
	Vnums string
	Zone string
	ZonePos []int
	ZoneMap [][]int
	Vnum int
	Desc string
	Mobiles []int
	MobilesInRoom []Mobile
	Items []int
	CoreBoard string
	Exits Exit
	Altered bool
}
type Exit struct {
	North int
	South int
	East int
	West int
	NorthWest int
	NorthEast int
	SouthWest int
	SouthEast int
	Up int
	Down int
}
type EquipmentItem struct {
	Item Object
	Number int
}
type EquipmentBank struct {

  SlotOne EquipmentItem
  SlotOneAmount int

  SlotTwo EquipmentItem
  SlotTwoAmount int

  SlotThree EquipmentItem
  SlotThreeAmount int

  SlotFour EquipmentItem
  SlotFourAmount int

  SlotFive EquipmentItem
  SlotFiveAmount int

  SlotSix EquipmentItem
  SlotSixAmount int

  SlotSeven EquipmentItem
  SlotSevenAmount int

  SlotEight EquipmentItem
  SlotEightAmount int

  SlotNine EquipmentItem
  SlotNineAmount int

  SlotTen EquipmentItem
  SlotTenAmount int
}
type InventoryBank struct {

  SlotOne InventoryItem
  SlotOneAmount int

  SlotTwo InventoryItem
  SlotTwoAmount int

  SlotThree InventoryItem
  SlotThreeAmount int

  SlotFour InventoryItem
  SlotFourAmount int

  SlotFive InventoryItem
  SlotFiveAmount int

  SlotSix InventoryItem
  SlotSixAmount int

  SlotSeven InventoryItem
  SlotSevenAmount int

  SlotEight InventoryItem
  SlotEightAmount int

  SlotNine InventoryItem
  SlotNineAmount int

  SlotTen InventoryItem
  SlotTenAmount int
}
type Action struct {
	Affects Player
	From Mobile
	Damage int
	DamMsg string
}
type Player struct {
	Hostname string
	Name string
	Title string
	ItemBank InventoryBank
	EqBank EquipmentBank
	Inventory []InventoryItem
	Equipped []EquipmentItem
	CoreBoard string
	PlainCoreBoard string
	CurrentRoom Space
	PlayerHash string
	Classes []Class
	Level float64
	Target string
	TargetLong string
	ESlotSpell Spell
	ESlotSkill Skill
	QSlotSpell Spell
	QSlotSkill Skill
	//ToBuy will have to either be an ItemHash
	//Or a vnum
	ToBuy int
	BankAccount Account
	TarX int
	TarY int
	OldX int
	OldY int
	CPU string
	CoreShow bool
	Channels []string
	Battling bool
	BattlingMob Mobile
	Profile string
	Session string
	Attack int
	Defend int

	Slain int
	Hoarded int


	MaxRezz int
	Rezz int
	Tech int
	Fights Fight
	Won int
	Found int

	Str int
	Int int
	Dex int
	Wis int
	Con int
	Cha int
}
type Fight struct {
	Won int
	Found int
	Oppose []Mobile
	Former []Player
	Treasure []Object
}
type Mobile struct {
	Name string
	Corpse string
	LongName string
	ItemSpawn []int
	Rep string
	MaxRezz int
	AC int
	Rezz int
	Tech int
	Aggro bool
	Align int
	Attack float64
	Defend float64
	Vnum int
	X int
	Y int
	Char string
}
