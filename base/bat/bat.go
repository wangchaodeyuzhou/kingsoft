package bat

type BattleType struct {
	Hello *HelloType
}

type HelloType struct {
	say *SayType
}

type SayType struct {
	team int32
}

func (x *BattleType) GetHelloType() *HelloType {
	if x != nil {
		return x.Hello
	}
	return nil
}

func (x *HelloType) GetSayType() *SayType {
	if x != nil {
		return x.say
	}

	return nil
}

func (x *SayType) GetTeam() int32 {
	if x != nil {
		return x.team
	}
	return 0
}
