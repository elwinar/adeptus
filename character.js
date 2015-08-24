function include(filename, type) {

	var data = require(filename).data;
	
	if(typeof(type) != "undefined") {
		for(key in data) {
			data[key].type = type;
		}
	}

	return data;
}

var characteristics = include('./characteristics', "characteristic");
var skills = include('./skills', "skill");
var talents = include('./talents', "talent");

var config = {};
config.types = include('./types');
config.enhancements = characteristics.concat(skills).concat(talents);

exports.make = function(name, aptitudes) {
	var self = this;
	self.name = name;
	self.aptitudes = aptitudes;
	
	self.characteristics = [];
	self.skills = [];
	self.talents = [];
	self.custom = [];
	
	self.upgrades = [];
	self.experience = 0;
	
	function defineTier(upgrade, history) {
		var tier = 1;
		for(var it = 0; it < history.length; it++) {
			if(self.upgrades[it].id != upgrade.id) {
				continue;
			}
			if(self.upgrades[it].isMisc) {
				continue;
			}
			tier++;
		}
		return tier;
	}
	
	function countAptitudes(enhancement, aptitudes) {
		var count = 0;
		for(var it = 0; it < enhancement.aptitudes.length; it++) {
			if(aptitudes.indexOf(enhancement.aptitudes[it]) == -1) {
				continue;
			}
			count++;
		}
		return count;
	}
	
	function findEnhancement(upgrade, enhancements) {
		for(var it = 0; it < enhancements.length; it++) {
			if(enhancements[it].id == upgrade.id) {
				return enhancements[it];
			}
		}
		return {type:"custom", id: upgrade.id, aptitudes: [], cost: 0};
	}
	
	function defineValue(upgrade) {
		if(typeof(upgrade.value) == "undefined") {
			return 1;
		}
		return upgrade.value;
	}
	
	function defineType(enhancement) {
		if(['talent', 'characteristic', 'skill'].indexOf(enhancement.type) != -1) {
			return enhancement.type + "s";
		}
		return "custom";
	}
	
	function defineCost(enhancement, upgrade, match, tier) {
		if(typeof(upgrade.cost) != "undefined") {
			return upgrade.cost;
		}		
		return config.types[enhancement.type][tier][match];
	}
	
	self.enhance = function(upgrade) {
		var enhancement = findEnhancement(upgrade, config.enhancements);
		var type = defineType(enhancement);
		if(typeof(self[type][enhancement.id]) == "undefined") {
			self[type][enhancement.id] = 0;
		}
		self[type][enhancement.id] += defineValue(upgrade);
		var match = countAptitudes(enhancement, self.aptitudes);
		var tier = defineTier(upgrade, self.upgrades);
		
		self.experience += defineCost(enhancement, upgrade, match, tier);
		self.upgrades.push(upgrade);
	}
	
	self.log = function() {
		for(key in self.upgrades) {
			var history = self.upgrades.slice(0, key);
			var upgrade = self.upgrades[key];
			var enhancement = findEnhancement(upgrade, config.enhancements);
			var value = defineValue(upgrade);
			var match = countAptitudes(enhancement, self.aptitudes);
			var tier = defineTier(upgrade, history);
			
			var cost = defineCost(enhancement, upgrade, match, tier);
			console.log(key + " - " + cost + " " + upgrade.id + " " + value);
		}
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
		console.log("Skills");
		for(id in self.skills){
			console.log(id + ": " + self.skills[id]);
		}
		console.log("----");
		console.log("Talents");
		for(id in self.talents){
			console.log(id + ": " + self.talents[id]);
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
