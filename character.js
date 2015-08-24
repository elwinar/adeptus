var config = require('./config');

exports.make = function(name, aptitudes) {
	var self = this;
	self.name = name;
	self.aptitudes = aptitudes;
	
	self.caracteristics = [];
	self.custom = [];
	
	self.upgrades = [];
	self.experience = 0;
	
	function defineTier(id) {
		var ups = 1;
		for(var it = 0; it < self.upgrades.length; it++) {
			if(self.upgrades[it].id != id) {
				continue;
			}
			if(self.upgrades[it].isMisc) {
				continue;
			}
			ups++;
		}
		return ups;
	}
	
	function countAptitudes(aptitudes) {
		var count = 0;
		for(var it = 0; it < aptitudes.length; it++) {
			if(self.aptitudes.indexOf(aptitudes[it]) == -1) {
				continue;
			}
			count++;
		}
		return count;
	}
	
	function findEnhancement(id) {
		for(var it = 0; it < config.enhancements.length; it++) {
			if(config.enhancements[it].id == id) {
				return config.enhancements[it];
			}
		}
	}
	
	self.enhance = function(upgrade) {
		var enhancement = findEnhancement(upgrade.id);
		
		if(typeof(enhancement) == "undefined") {		
			var cost = 0;
			if(typeof(upgrade.cost) != "undefined") {
				cost = upgrade.cost;
			}
			self.custom[upgrade.id] = upgrade.value;
			self.experience += cost;
			self.upgrades.push(upgrade);
			return;
		}
		
		// defined tier
		var tier = defineTier(upgrade.id);
		if(typeof(config.types[enhancement.type][tier]) == "undefined") {
			console.error("Undefined tier for upgrade: " + enhancement.id + "[" + tier + "]");
			return;
		}
		
		// define cost
		if(typeof(upgrade.cost) != "undefined") {
			self.experience += upgrade.cost;
		}
		else {
			var match = countAptitudes(enhancement.aptitudes);
			self.experience += config.types.caracteristic[tier][match];
		}
		
		// improves value
		var id = enhancement.id;
		var attribute;
		switch(enhancement.type) {
			case "caracteristic":
				attribute = self.caracteristics;
				break;
		}
		if(typeof(attribute[id]) == "undefined") {
			attribute[id] = 0;
		}
		attribute[id] += upgrade.value;
		
		// save upgrade
		self.upgrades.push(upgrade);
	}
	
	self.display = function() {
		console.log("---------");
		console.log("name: " + self.name);
		console.log("----");
		console.log("Caracteristics");
		for(id in self.caracteristics){
			console.log(id + ": " + self.caracteristics[id]);
		}
		console.log("----");
		console.log("Custom");
		for(id in self.custom){
			console.log(id + ": " + self.custom[id]);
		}
		console.log("----");
		console.log("xp: " + self.experience);
	}
}