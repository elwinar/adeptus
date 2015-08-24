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

exports.config = {
	types: include('./types'), 
	enhancements: characteristics.concat(skills).concat(talents), 
};