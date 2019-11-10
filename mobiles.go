package main


func lookupMobile(vnum int) Mobile {
  var mob Mobile
  switch vnum {
  case 0:
      mob.Vnum = 0
      mob.Name = "an enormous poptart kitten"
      mob.LongName = "A poptart kitten the size of a bus scoots along here"
      mob.Rep = "1000"
      mob.MaxRezz = 100
      mob.Rezz = 100
      mob.Tech = 100
      mob.Aggro = false
      mob.Align = 0
      return mob
    default:
      mob.Vnum = 0
      mob.Name = "an enormous poptart kitten"
      mob.LongName = "A poptart kitten the size of a bus scoots along here"
      mob.Rep = "1000"
      mob.MaxRezz = 100
      mob.Rezz = 100
      mob.Tech = 100
      mob.Aggro = false
      mob.Align = 0
      return mob
  }
  return mob
}
