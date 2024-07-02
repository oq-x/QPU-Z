package specs

var intelFamilyModelMap = map[int]map[int]string{
	15: {
		1: "Netburst (Williamette)",
		2: "Netburst (Northwood)",
		3: "Netburst (Prescott)",
		4: "Netburst (Prescott)",
		6: "Netburst",
	},
	11: {
		0: "Knights Ferry",
		1: "Knights Corner",
	},
	6: {
		183: "Raptor Lake-S",
		186: "Raptor Lake-P",

		151: "Alder Lake-S",
		154: "Alder Lake-P",

		167: "Rocket Lake-S",

		141: "Tiger Lake-H",
		140: "Tiger Lake-U",

		126: "Ice Lake-U/Y",

		165: "Comet Lake-S/H",
		142: "Comet Lake/Coffee Lake/Whiskey Lake/Kaby Lake-U/Amber Lake-Y",

		102: "Cannon Lake-U",

		158: "Coffee Lake-S/H/E",

		94: "Skylake-DT/H/S",
		78: "Skylake-Y/U",

		71: "Broadwell-C/W/H",
		61: "Broadwell-U/Y/S",

		70: "Haswell-GT3E",
		69: "Haswell-ULT",
		60: "Haswell-S",

		58: "Ivy Bridge",

		42: "Sandy Bridge",

		37: "Westmere (Arrandale/Clarkdale)",
		31: "Nahelem (Auburndale/Havendale)",
		30: "Nahelem (Clarksfield)",
		23: "Penyrn/Wolfdale/Yorkfield",
		22: "Merom-L",
		15: "Merom",
		14: "Yonah",
		21: "Tolapai",
		13: "Dothan",
		9:  "Banias",
		11: "Tualatin",
		8:  "Coppermine",
		7:  "Katmai",
	},
}

// [family][model, extmodel] => display name
var amdFamilyModelMap = map[int]map[[2]int]string{
	0x19: {
		{0, 5}: "Zen 3",
		{4, 4}: "Zen 3",
		{0, 4}: "Zen 3",
		{1, 2}: "Zen 3",

		{0, 1}: "Zen 4",

		{8, 0}: "Zen 3",
		{1, 0}: "Zen 3",
		{0, 0}: "Zen 3",
	},
	0x18: {
		{0, 0}: "Zen",
	},
	0x17: {
		{160, 10}: `Zen 2 "Mendocino"`,
		{144, 9}:  `Zen 2 "Van Gogh"`,
		{113, 7}:  `Zen 2 "Matisse"`,
		{104, 6}:  `Zen 2 "Lucienne"`,
		{96, 6}:   `Zen 2 "Renoir/Grey Hawk"`,
		{71, 4}:   `Zen 2 "Xbox Series X"`,
		{49, 3}:   `Zen 2 "Rome/Castle Peak"`,

		{24, 1}: `Zen+ "Picasso"`,
		{8, 0}:  `Zen+ "Colfax/Pinnacle Ridge"`,

		{32, 2}: `Zen "Dali"`,
		{24, 1}: `Zen "Banded Kestrel"`,
		{17, 1}: `Zen "Raven Ridge/Great Horned Owl"`,
		{1, 0}:  `Zen "Naples/Whitehaven/Summit Ridge/Snowy Owl"`,
	},
	0x16: {
		{0, 3}: "Puma",
		{0, 0}: "Jaguar",
	},
	0x15: {
		{0, 7}: `Excavator "Stoney Ridge"`,
		{5, 6}: `Excavator "Bristol Ridge"`,
		{0, 6}: `Excavator "Carrizo"`,

		{8, 3}: `Steamroller "Godavari"`,
		{0, 3}: `Steamroller "Kaveri"`,

		{3, 1}: `Piledriver "Richland"`,
		{0, 1}: `Piledriver "Trinity"`,
		{2, 0}: "Piledriver",

		{1, 0}: "Bulldozer",
	},
	0x14: {
		{0 /*UNKNOWN*/, 0}: "Bobcat",
	},
	0x12: {
		{0, 0}: "Llano",
	},
	0x10: {
		{10, 0}: "K10",
		{9, 0}:  "K10",
		{8, 0}:  "K10",
		{6, 0}:  "K10",
		{5, 0}:  "K10",
		{4, 0}:  "K10",
		{2, 0}:  "K10",
	},
	0x0F: {
		{0x1, 0xC}: `Athlon 64 FX "Windsor"`,
		{0xF, 0x7}: `Athlon 64, Athlon, Sempron`,
		{0xF, 0x6}: `Athlon 64, Neo, Sempron`,
		{0xB, 0x6}: "Dual Core Athlon/Athlon 64/Athlon Neo/Sempron/Turion",
		{0x8, 0x6}: "Dual Core Athlon/Athlon 64/Athlon Neo/Sempron/Turion",
		{0xF, 0x5}: "Athlon 64, Athlon, Sempron",
		{0xD, 0x5}: "Opteron",
		{0xF, 0x4}: "Athlon 64, Sempron",
		{0xC, 0x4}: "Turion 64, Athlon 64, Sempron",
		{0xB, 0x4}: "Athlon 64 X2",
		{0x8, 0x4}: "Turion 64 X2, Athlon 64 X2",
		{0x3, 0x4}: `Opteron 1200 "Santa Ana", Athlon 64 FX "Windsor", Athlon 64 X2`,
		{0x1, 0x4}: `Opteron 2200/8200 "Santa Rosa"`,

		{0xF, 0x2}: "Athlon, Sempron",
		{0xC, 0x2}: "Sempron, Mobile Sempron",
		{0xB, 0x2}: `Athlon 64 X2 "Manchester"`,
		{0x7, 0x2}: `Opteron 1xx "Venus", Athlon 64 FX, Athlon`,
		{0x5, 0x2}: `Opteron 1xx/2xx/8xx "Venus/Troy/Athens"`,
		{0x4, 0x2}: "Turion, Mobile Athlon 64",
		{0x3, 0x2}: `Opteron 1xx "Denmark", Athlon 64 X2 & Athlon 64 FX "Toledo"`,
		{0x1, 0x2}: `Opteron 1xx/2xx/8xx "Denmark/Italy/Egypt"`,
		{0xF, 0x1}: "Athlon, Sempron",
		{0xC, 0x1}: "Athlon, Sempron, Mobile Athlon 64, Mobile Sempron",
		{0xB, 0x1}: "Athlon",
		{0x8, 0x1}: "Athlon, Mobile Athlon 64, Mobile Athlon XP-M, Mobile Sempron",
		{0x7, 0x1}: "Athlon, Athlon 64 FX",
		{0x5, 0x1}: "Opteron, Athlon 64 FX",
		{0x4, 0x1}: "Athlon, Mobile Athlon 64, Mobile Athlon XP-M",
		{0xF, 0x0}: "Sempron",
		{0xC, 0x0}: "Mobile Athlon 64, Mobile Athlon XP-M, Sempron, Mobile Sempron",
		{0xE, 0x0}: "Mobile Athlon 64, Mobile Athlon XP-M, Sempron, Mobile Sempron",
		{0x8, 0x0}: "Mobile Athlon 64, Mobile Athlon XP-M, Mobile Sempron",
		{0x7, 0x0}: "Athlon 64, Athlon 64 FX",
		{0x5, 0x0}: "Opteron 1xx/2xx/8xx, Athlon 64 FX",
		{0x4, 0x0}: "Athlon 64, Mobile Athlon 64, Mobile Athlon XP-M",
	},
}
