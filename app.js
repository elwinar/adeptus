var character = require('./character.js');

var Paola = new character.make("Paola", [
	"offence", 
	"strength", 
	"ballistic skill", 
	"weapon skill", 
	"defence", 
	"agility", 
	"leadership", 
	"social",
	"general",
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

Paola.enhance({id: "Weapon Prof. Spear", cost: 0});
Paola.enhance({id: "Iron Jaw", cost: 0});

Paola.enhance({id: "Command", cost: 0});
Paola.enhance({id: "Charm", cost: 0});
Paola.enhance({id: "Scrutiny", cost: 0});
Paola.enhance({id: "Common Lore: Imperial Cult", cost: 0});

Paola.enhance({id: "WS", value: 5});
Paola.enhance({id: "Parry", value: 1});
Paola.enhance({id: "Bullwark of the godess", cost: 200});
Paola.enhance({id: "STR", value: 5});
Paola.enhance({id: "TOU", value: 5});
Paola.enhance({id: "AGI", value: 5});

Paola.enhance({id: "Weapon Prof. Shield"});
Paola.enhance({id: "Resistance Disease"});

Paola.display();
console.log();
console.log();
console.log();
Paola.log();
