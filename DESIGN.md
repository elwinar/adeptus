# adeptus

Adeptus' goal is to track character sheets in a simple and intuitive manner.

## Character

A character is defined by the following informations:

- Name
- Backgrounds
- Aptitudes
- Upgrades

The upgrades are a list of modifications applied to the base character sheet that can be of 5 types:

- A characteristic bonus
- A new talent
- A skill level
- A gauge
- A new special rule

### Name

The name of a character can be any UTF-8 encodable string.

### Backgrounds

The backgrounds of a character define a set of bonus that are applied to the character at its creation.
They can take the form of characteristic bonuses or maluses, new aptitudes, talents or skill upgrades, or special rules.

### Aptitudes

The aptitudes of a character are a set of values determined at the character creation depending on sbackground of the character. 
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

Characteristic have levels of upgrades called `tier`, each upgrade providing a bonus to the characteristic starting value.
The canonical bonus is `+5`, but it must be specified.
The number of available upgrades for a characteristic depends on the universe.

Characteristics upgrade costs depends on 2 aptitudes, the cost of each upgrade being a function of the presence of the needed aptitudes in the aptitude set of the character and tier.

### Skills

The skills are a set of competences known by the caracter.

Skills have levels of upgrades called `tier`, each upgrade providing a bonus to actions depending on this skill.

Skills costs depends on the presence of 2 aptitudes in the character aptitudes set and of the upgrade tier.

### Talents

The talents are a set of situationnal competences that provide bonuses to the use of skills.

Some talents can stack with themselves. In such cases, the talent's value increases by 1 for each upgrade.

Talents are splitted into tiers which determine their cost along with the presence of aptitudes in the aptitude set of the character.

Talents also have prerequisites in terms of characteristic values, skills, other talents, etc.

### Spells

The spells are similar to custom rules. However, their structure is very specific:

- Name, the name of the spell
- Cost, the XP cost of the spell
- Description, the in-game effect of the spell

Additional custom attributes can be defined.

Example:
```
{
	name: "Fireball",
	cost: 100,
	description: "Toss a fireball at an enemy, inflicting 1D10 fire damages.",
	push: "The fireball ignores toughness and armor.",
	tier: 1,
}
```

## Sheet file

A character sheet is stored in a file following a special format designed to be intuitive to read and write for humans. The file to use must be specified on the command line. It is the first argument of the command.

The sheet itself consist in a header defining the character's name and backgrounds, followed by a block of characteristic rolls, and a number of sessions. Each session consists in a headline containing a date, a label, and an experience point reward, followed by a list of upgrades to apply.

The header line is mandatory. Backgrounds may propose choices between different skills, talents or other upgrades. Each choice made must be precised in parenthesis `()`, separated by comas `,`.

The characteristic block must containt each characteristic defined in the universe.

Any character following either `#`, or `//` on the same line will be ignored.

Each upgrade is composed of a mark, the upgrade definition and eventually a cost (surrounded by brackets `[]`). The upgrade definition is composed of the name of the upgrade, its type being inferred from its name, and eventually any additionnal information needed like the value change for a characteristic, or an eventual specialisation for a skill or talent.

Example of sheet file:

```
# Header block
Name: Sephiam
Origin: Wood Elve
Background: Outcast (Jaded)
Role: Warrior (Weapon Proficiency: Sword, Iron Jaw)
Tarot: Boon & Bane (WS +3, INT -3)

# Characteristic block
WS	38
BS	33
STR	29
AGI	34
TOU	29
INT	33
PER	25
WIL	31
FEL	26

# Session blocks
2015/07/01 Creation [1500]
	+ Weapon Proficiency: Bow
	+ Catfall
	+ BS +5
	+ AGI +5
	+ Combat master

2015/07/01 First scenario [750]
	+ Resistance: Disease
	+ WIL +5 [100]
	+ Malignancy: Fear of the Damned
	+ fieldcraft
	- strength
	* STR -3
```

### Blanks

The character sheet can use blanks and tabulations without distinction.
The blanks can be repeated at will, as long as there is at least one to separate different elements.

### Marks

The marks define how the upgrade is taken into account:

- `+` is the standard way of adding an upgrade: the upgrade is taken into account for the maximum number of upgrades and cost of futures upgrades. The default cost is computed.
- `-` is the way of losing an upgrade: the previous upgrade is no longer taken into account for the maximum number of upgrades and cost. The default cost is 0.
- `*` is a special way of acquiring upgrades: the upgrade isn't taken into account for the maximum number of upgrades and cost of future upgrades. The default cost is 0.

The mark must be the first item of the line, not counting blanks

*Note: The * mark can only be used for characteristic upgrades*

### Experience

The value between brackets is experience. A Session offers some, an Upgrade costs some.
- must be a positive integer
- must be between brackets
- must not contain any blank
- must be placed in second or in last position of the line

### Characteristics upgrades

Characteristics upgrades can by specified using two ways:

- `relative value` will change the current value of the characteristic based on the given modifier. Standard characteristics upgrades give a `+5` bonus, but game-master awarded upgrades (or downgrades) will typically have differents bonuses/maluses, and differents costs. Anyhow, the value of a characteristic upgrade must always be specified. If it isn't specified, the program stop with an error.
- `absolute value` will define a new starting point for the characteristic and reset the number of applied upgrades for the  characteristic. Their cost can generally be ommitted, in which case it will be assumed to be 0.

In each case, the upgrade is defined by the name of the characteristic, a space, then the new value (either absolute or relative).

Example:

```
	  STR +5 // error, no mark is specified, unless it is used in the characteristic block
	+ STR +5 // the character gains 5 points of STR and the current STR tier is increased by 1
	+ STR -5 // error, a standard upgrade cannot provoke a characteristic loss
	- STR +5 // error, a downgrade cannot provoke a characteristic up
	- STR -5 // the character loose 5 points of STR and the current STR tier is reduced by 1
	* STR +5 // the character gains 5 points of STR, but the STR tier does not change
	* STR -5 // the character looses 5 points of STR, but the STR tier does not change
	* STR 50 // the strength of the character bumps to 50, the tier are not affected
```

*Note: In the characteristics block, the mark must be omited.*

### Talents & skills upgrades

Talents and skills upgrades are defined by their name, eventually followed by a colon and specialization. To specialise a talent or skill, use
the `:` separator.

### Special rules upgrades

Special rule are defined by their name only.

### Unrecognized upgrades

If an upgrade definition isn't recognized, it is considered as a special rule with a 0 cost value.

## Universes

The list of skills, talents, special rules, aptitudes, etc, is stored in a JSON file named an "universe". The program must load the universe prior to doing any other action.

The universe file is named `universe.json` and is located either in the working directory or in the installation directory of the program. Its location or name can be overriden by the flag `universe,u`.

The file is a valid JSON file containing multiple arrays:

- `aptitudes` list the names of aptitudes
- `characteristics` & `skills` list the names and aptitudes of characteristics and skills
- `talents` list the names and the prerequisites of skills
- `gauges` list the names of the existing gauges
- `backgrounds` list the names and upgrades of backgrounds

## Commands

The program accept multiple commands that have different outputs. Every command is read-only, so a character sheet or universe is never modified as the result of an `adeptus` command.

### History

The `history,h` command line switch will display the upgrades history along with the character sheet.

### Suggest

The `suggest,s` command line will propose purchasable upgrades for the character. Any upgrade with a cost lesser than the remaining XP
of the character is displayed. The maximum value of the proposed upgrades can be overriden with the `max` and the `all` flag.

### aptitudes/skills/talents/backgrounds/characteristics

Display the fill list of availables entries of the corresponding type in the selected universe, along with all available informations about it.
