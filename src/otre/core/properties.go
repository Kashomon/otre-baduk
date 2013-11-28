package core

type SgfProperty string

const (
	AB SgfProperty = "AB"
	AE SgfProperty = "AE"
	AN SgfProperty = "AN"
	AP SgfProperty = "AP"
	AR SgfProperty = "AR"
	AS SgfProperty = "AS"
	AW SgfProperty = "AW"
	B  SgfProperty = "B"
	BL SgfProperty = "BL"
	BM SgfProperty = "BM"
	BR SgfProperty = "BR"
	BS SgfProperty = "BS"
	BT SgfProperty = "BT"
	C  SgfProperty = "C"
	CA SgfProperty = "CA"
	CH SgfProperty = "CH"
	CP SgfProperty = "CP"
	CR SgfProperty = "CR"
	DD SgfProperty = "DD"
	DM SgfProperty = "DM"
	DO SgfProperty = "DO"
	DT SgfProperty = "DT"
	EL SgfProperty = "EL"
	EV SgfProperty = "EV"
	EX SgfProperty = "EX"
	FF SgfProperty = "FF"
	FG SgfProperty = "FG"
	GB SgfProperty = "GB"
	GC SgfProperty = "GC"
	GM SgfProperty = "GM"
	GN SgfProperty = "GN"
	GW SgfProperty = "GW"
	HA SgfProperty = "HA"
	HO SgfProperty = "HO"
	ID SgfProperty = "ID"
	IP SgfProperty = "IP"
	IT SgfProperty = "IT"
	IY SgfProperty = "IY"
	KM SgfProperty = "KM"
	KO SgfProperty = "KO"
	L  SgfProperty = "L"
	LB SgfProperty = "LB"
	LN SgfProperty = "LN"
	LT SgfProperty = "LT"
	M  SgfProperty = "M"
	MA SgfProperty = "MA"
	MN SgfProperty = "MN"
	N  SgfProperty = "N"
	OB SgfProperty = "OB"
	OH SgfProperty = "OH"
	OM SgfProperty = "OM"
	ON SgfProperty = "ON"
	OP SgfProperty = "OP"
	OT SgfProperty = "OT"
	OV SgfProperty = "OV"
	OW SgfProperty = "OW"
	PB SgfProperty = "PB"
	PC SgfProperty = "PC"
	PL SgfProperty = "PL"
	PM SgfProperty = "PM"
	PW SgfProperty = "PW"
	RE SgfProperty = "RE"
	RG SgfProperty = "RG"
	RO SgfProperty = "RO"
	RU SgfProperty = "RU"
	SC SgfProperty = "SC"
	SE SgfProperty = "SE"
	SI SgfProperty = "SI"
	SL SgfProperty = "SL"
	SO SgfProperty = "SO"
	SQ SgfProperty = "SQ"
	ST SgfProperty = "ST"
	SU SgfProperty = "SU"
	SZ SgfProperty = "SZ"
	TB SgfProperty = "TB"
	TC SgfProperty = "TC"
	TE SgfProperty = "TE"
	TM SgfProperty = "TM"
	TR SgfProperty = "TR"
	TW SgfProperty = "TW"
	UC SgfProperty = "UC"
	US SgfProperty = "US"
	V  SgfProperty = "V"
	VW SgfProperty = "VW"
	W  SgfProperty = "W"
	WL SgfProperty = "WL"
	WR SgfProperty = "WR"
	WS SgfProperty = "WS"
	WT SgfProperty = "WT"
)

var validProps = map[SgfProperty]bool{
	AB: true,
	AE: true,
	AN: true,
	AP: true,
	AR: true,
	AS: true,
	AW: true,
	B:  true,
	BL: true,
	BM: true,
	BR: true,
	BS: true,
	BT: true,
	C:  true,
	CA: true,
	CH: true,
	CP: true,
	CR: true,
	DD: true,
	DM: true,
	DO: true,
	DT: true,
	EL: true,
	EV: true,
	EX: true,
	FF: true,
	FG: true,
	GB: true,
	GC: true,
	GM: true,
	GN: true,
	GW: true,
	HA: true,
	HO: true,
	ID: true,
	IP: true,
	IT: true,
	IY: true,
	KM: true,
	KO: true,
	L:  true,
	LB: true,
	LN: true,
	LT: true,
	M:  true,
	MA: true,
	MN: true,
	N:  true,
	OB: true,
	OH: true,
	OM: true,
	ON: true,
	OP: true,
	OT: true,
	OV: true,
	OW: true,
	PB: true,
	PC: true,
	PL: true,
	PM: true,
	PW: true,
	RE: true,
	RG: true,
	RO: true,
	RU: true,
	SC: true,
	SE: true,
	SI: true,
	SL: true,
	SO: true,
	SQ: true,
	ST: true,
	SU: true,
	SZ: true,
	TB: true,
	TC: true,
	TE: true,
	TM: true,
	TR: true,
	TW: true,
	UC: true,
	US: true,
	V:  true,
	VW: true,
	W:  true,
	WL: true,
	WR: true,
	WS: true,
	WT: true,
}

func IsValidProperty(prop string) bool {
	_, ok := validProps[SgfProperty(prop)]
	return ok
}
