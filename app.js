var character = require('./character.js');

var Paola = new character.make("Paola", [
	"offence", 
	"strength", 
	"ballistic skill", 
	"weapon skill", 
	"defence", 
	"agility", 
	"leadership", 
	"social"
]);
Paola.enhance({id: "WS", value: 36, cost: 0, isMisc: true});
Paola.enhance({id: "BS", value: 28, cost: 0, isMisc: true});
Paola.enhance({id: "STR", value: 25, cost: 0, isMisc: true});
Paola.enhance({id: "AGI", value: 27, cost: 0, isMisc: true});
Paola.enhance({id: "TOU", value: 36, cost: 0, isMisc: true});
Paola.enhance({id: "INT", value: 31, cost: 0, isMisc: true});
Paola.enhance({id: "PER", value: 33, cost: 0, isMisc: true});
Paola.enhance({id: "WP", value: 32, cost: 0, isMisc: true});
Paola.enhance({id: "FEL", value: 38, cost: 0, isMisc: true});

Paola.enhance({id: "Weapon prof. Spear", value: true, cost: 0});
Paola.enhance({id: "Iron jaw", value: true, cost: 0});

Paola.enhance({id: "Command", value: "trained", cost: 0});
Paola.enhance({id: "Charm", value: "trained", cost: 0});
Paola.enhance({id: "Scrutiny", value: "trained", cost: 0});
Paola.enhance({id: "Common Lore: imperial cult", value: "trained", cost: 0});

Paola.enhance({id: "WS", value: 05});
Paola.enhance({id: "Parry", value: "trained", cost: 100});
Paola.enhance({id: "Bullwark of the godess", value: true, cost: 200});
Paola.enhance({id: "STR", value: 05});
Paola.enhance({id: "TOU", value: 05});
Paola.enhance({id: "AGI", value: 05});
Paola.enhance({id: "Weapon prof. Shield", value: true, cost: 300});
Paola.enhance({id: "Resistance: Disease", value: true, cost: 300});

Paola.display();