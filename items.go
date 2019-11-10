package main


func lookupObject(vnum int) Object {
  var obj Object
  switch vnum {
  case 1:
    obj.Vnum = 1
    obj.Name = "a red rose"
    obj.LongName = "A rose floats here, slowly rotating."
    obj.Zone = "zem"
    obj.Value = 1
    obj.Owned = false
    return obj
  case 2:
    obj.Vnum = 2
    obj.Name = "a nyancat"
    obj.LongName = "A poptart kitten happily miaos in a circle."
    obj.Zone = "zem"
    obj.Value = 100
    obj.Owned = false
  default:
    obj.Vnum = 0
    obj.Name = "nothing"
    obj.LongName = "The lack of things is...lacking"
    obj.Zone = "zem"
    obj.Value = 0
    obj.Owned = false
  }
  return obj
}
