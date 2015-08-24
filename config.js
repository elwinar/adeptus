exports.enhancements = [
	{
		id: "WS", 
		aptitudes: ["weapon skill", "offence"],
		type: "caracteristic"
	},
	{
		id: "BS", 
		aptitudes: ["ballistic skill", "finesse"],
		type: "caracteristic"
	},
	{
		id: "STR", 
		aptitudes: ["strength", "offence"],
		type: "caracteristic"
	},
	{
		id: "AGI", 
		aptitudes: ["agility", "finesse"],
		type: "caracteristic"
	},
	{
		id: "TOU", 
		aptitudes: ["toughness", "defence"],
		type: "caracteristic"
	},
	{
		id: "INT", 
		aptitudes: ["intelligence", "knowledge"],
		type: "caracteristic"
	},
	{
		id: "PER", 
		aptitudes: ["perception", "fieldcraft"],
		type: "caracteristic"
	},
	{
		id: "WP", 
		aptitudes: ["willpower", "psyker"],
		type: "caracteristic"
	},
	{
		id: "FEL", 
		aptitudes: ["fellowship", "social"],
		type: "caracteristic"
	},
];

exports.types = {
	caracteristic: {
		1: { 0: 500, 1: 250, 2: 100},
		2: { 0: 1000, 1: 500, 2: 250},
		3: { 0: 1500, 1: 1000, 2: 500},
		4: { 0: 2000, 1: 1500, 2: 750},
	},
};