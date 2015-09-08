# adeptus

Adeptus' goal is to track character sheets in a simple and intuitive manner.

## Character

A character is defined by the following informations:

- Name
- Origin
- Background
- Role
- Tarot
- Aptitudes
- Upgrades

The upgrades are a list of modifications applied to the base character sheet that can be of 4 types:

- A characteristic bonus
- A new talent
- A skill level
- A new special rule

### Name

The name of a character can be any UTF-8 encodable string.

### Origin, Background, Role & Tarot

The origin, background, role and tarot of a character define a set of bonus that are applied to the character at its creation.
They can take the form of characteristic bonuses or maluses, new aptitudes, talents or skill upgrades, or special rules.

### Aptitudes

The aptitudes of a character are a set of values determined at the character creation depending on the origin, background and role of the character. 
The owned aptitudes of a character are used in determining the cost of various upgrades depending on the presence of a pair of aptitude.

### Characteristics

The characteristics of a character are a list of values representing some of the character defining characteristics:

- Weapon Skill (`WS`)
- Ballistic Skill (`BS`)
- Strength (`STR`)
- Agility (`AGI`)
- Toughness (`TOU`)
- Intelligence (`INT`)
- Perception (`PER`)
- Willpower (`WIL`)
- Fellowship (`FEL`)

The starting value of characteristics are determined by rolling dices at the character creation. Characteristics are numbers generally in the range [1..100] (actually no extension allow going above 85).

Characteristic have 5 levels of upgrades, each upgrade providing a `+5` bonus to the characteristic starting value.

Characteristics upgrade costs depends on 2 aptitudes, the cost of each upgrade being a function of the presence of the needed aptitudes in the aptitude set of the character and upgrade level.

### Skills

The skills are a set of competences known by the caracter.

Skills have 4 levels of upgrades, each upgrade providing a bonus to actions depending on this skill.

Skills costs depends on the presence of 2 aptitudes in the character aptitudes set and of the upgrade level.

### Talents

The talents are a set of situationnal competences that provide bonuses to the use of skills.

Talents doesn't have levels or values: they are either known, or not.

Talents are splitted into tiers which determine their cost along with the presence of aptitudes in the aptitude set of the character,

Talents also have prerequisites in terms of characteristic values, skills, other talents, etc.

## Sheet file

A character sheet is stored in a file following a special format designed to be intuitive to read and write for humans. The file to use must be specified on the command line using the `sheet,s` command line option.

The sheet itself consist in a header defining the character's name, universe, origin, background, role and tarot, followed by a block of characteristic rolls, and a number of sessions. Each session consists in a headline containing a date, a label, and an experience point reward, followed by a list of upgrades to apply.

Each header line is mandatory. Origin, Background, Role and Tarot may propose choices between different skills, talents or other upgrades. Each choice made must be precised in parenthesis `()`, separated by comas `,`.

The characteristic block must containt each characteristic defined in the universe.

Any line beginning with either `#`, `;` or `//` will be ignored

Each upgrade is composed of a mark, the upgrade definition and eventually a cost (surrounded by brackets `[]`). The upgrade definition is composed of the name of the upgrade, its type being inferred from its name, and eventually any additionnal information needed like the value change for a characteristic, or an eventual specialisation for a skill or talent.

Example of sheet file:

```
# Header block
Name: Sephiam
Universe: Fantasy
Origin: Wood Elve
Background: Outcast (Jaded)
Role: Warrior (Weapon Proficiency: Sword, Iron Jaw)
Tarot: Boon & Bane (WS +3, INT -3)

# Characteristic block
WS 	38
BS 	33
STR 	29
AGI 	34
TOU 	29
INT 	33
PER 	25
WIL 	31
FEL 	26

# Session blocks
2015/07/01 Creation [1500]
	* Weapon Proficiency: Bow
	* Catfall
	* BS +5
	* AGI +5
	* Combat master

2015/07/01 First scenario [750]
	* Resistance: Disease
	+ WIL +2 [100]
	- Malignancy: Fear of the Damned
```

### Blanks

The character sheet can use blanks and tabulations without distinction.
The blanks can be repeated at will, as long as there is at least one to separate different elements.

### Marks

The marks define how the upgrade is taken into account:

- `*` is the standard way of adding an upgrade: the upgrade is taken into account for the maximum number of upgrades and cost of futures upgrades
- `+` is a special way of adding an upgrades: the upgrade isn't taken into account for the maximum number of upgrades and cost of future upgrades
- `-` is an alias of `+ <definition> [0]`

In both cases, the point cost is computed on the fly, unless the point cost is specified on the upgrade line.

The mark must be the first item of the line, not counting blanks

### Experience

The value between brackets is experience. A Session offers some, an Upgrade costs some.
- must be a positive integer
- must be between brackets
- must not contain any blank
- must be placed in second or in last position of the line

### Characteristics upgrades

Characteristics upgrades can by specified using two ways:

- `absolute value` will define a new starting point for the characteristic and reset the number of applied upgrades for the  characteristic. Their cost can generally be ommitted, in which case it will be assumed to be 0.
- `relative value` will change the current value of the characteristic based on the given modifier. Standard characteristics upgrades give a +5 bonus, but game-master awarded upgrades (or downgrades) will typically have differents bonuses/maluses, and differents costs. The value of a non-standard characteristic upgade must be specified. If it isn't specified, the program stop with an error.

In each case, the upgrade is defined by the name of the characteristic, a space, then the new value (either absolute or relative).

### Talents & skills upgrades

Talents and skills upgrades are defined by their name, eventually followed by a colon and specialization.
In case of specialisation, `-`, `:` and `()` are valid separator.

### Special rules upgrades

Special rule are defined by their name only.

### Unrecognized upgrades

If an upgrade definition isn't recognized, it is considered as a special rule with the given experience point cost.
If not point cost is given, the program must stop with an error.

## Universes

The list of skills, talents, special rules, aptitudes, etc, is stored in a JSON file named an "universe". The program must load the universe prior to doing any other action.

The universe file is named `universe.json` and is located either in the working directory or in the installation directory of the program. It is defined by the `universe` value in the header.

The file is a valid JSON file containing multiple arrays:

- `aptitudes` list the names of aptitudes
- `characteristics` & `skills` list the names and aptitudes of characteristics and skills
- `talents` list the names and the prerequisites of skills
- `rules` list the names of special rules
- `origins`, `background`, `roles` & `tarots` list the names and upgrades of origins, backgrounds, roles and tarots

## Commands

The program accept multiple commands that have different outputs. Every command is read-only, so a character sheet or universe is never modified as the result of an `adeptus` command.

### sheet

The `sheet` command parse the universe, the chosen sheet file, and display the final state of the character sheet. If an error occurs when parsing the sheet, the program stops and display the error.

The `history,h` command line switch will display the upgrades history along with the character sheet.

The `format,f` command line option will change the output format of the character sheet. Its value must be the path to a valid template for character sheet.

### aptitudes/skills/talents/rules/origins/backgrounds/roles/tarots/characteristics

Display the fill list of availables entries of the corresponding type in the selected universe, along with all available informations about it.
