package main

import (
  "fmt"
  "math"
  "strconv"
)

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
  DamType string
	Consumed bool
	Name string
  Sound int
  Dam int
}

type Skill struct {
	Name string
	DamType string
	Level int
	Usage rune
	Dam int
  Sound int
}

func listClasses() []Class {
  var totalClasses []Class

  //add classes here to populate the game logic with them
  totalClasses = append(totalClasses, Brutalizer())
  totalClasses = append(totalClasses, Shaman())
  return totalClasses
}
func listMyClasses(play Player) string {
  out := ""
  for i := 0;i < len(play.Classes);i++ {
    if len(play.Classes[i].Name) > 1 {
      out += fmt.Sprintf("\033[38:2:150:0:150m"+play.Classes[i].Name+" %.2f\033[0m\n", play.Classes[i].Level)
    }else {
      out += fmt.Sprintln("\033[38:2:150:0:150mUnassigned "+strconv.Itoa(int(math.Floor(play.Classes[i].Level)))+"\033[0m")

    }
  }
  return out
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
  stab.Sound = 17
  brutal.Skills[0] = stab

  return brutal
}
func Shaman() Class {
  var shaman Class
  shaman.Level = 0
  shaman.Name = "Shaman"
  shaman.Skills = make([]Skill, 10, 10)
  shaman.Spells = make([]Spell, 10, 10)

  var shake Spell
  shake.Name = "shake"
  shake.DamType = "blud"
  shake.Level = 1
  shake.Usage = 'q'
  shake.Dam = 15
  shake.Sound = 12
  shaman.Spells[0] = shake

  return shaman
}
