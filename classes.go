package main


type Class struct {
	Level float64
	Name string
	Skills []Skill
	Spells []Spell
}


type Spell struct {
	TechUsage int
	Usage rune
	Level int
	Consumed bool
	Name string
}

type Skill struct {
	Name string
	DamType string
	Level int
	Usage rune
	Dam int
}

func listClasses() []Class {
  var totalClasses []Class

  //add classes here to populate the game logic with them
  totalClasses = append(totalClasses, Brutalizer())
  return totalClasses
}

func Brutalizer() Class {
  var brutal Class
  brutal.Level = 0
  brutal.Name = "Brutalizer"
  brutal.Skills = make([]Skill, 10, 10)
  brutal.Spells = make([]Spell, 10, 10)

  var stab Skill
  stab.Name = "stab"
  stab.DamType = "slice"
  stab.Level = 1
  stab.Usage = 'e'
  stab.Dam = 15
  brutal.Skills[0] = stab

  return brutal
}
