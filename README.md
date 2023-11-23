# Alchemy

Alchemy is a command line tool for modifying and transforming Matter spec documents.

## Format

Format does not alter the content of the document, but does make it easier to read the source:

. Aligns table cells so that table delimiters correctly line up with other cells in the same column
. Removes extraneous spacing in tables and lists
. Normalizes properties

## Disco Ball

Disco ball is more aggressive than format, and attempts to rewrite the document to the disco-ball standard:

. Everything format does
. Rearranges the document into disco-ball order for clusters and device-types
. Rearranges data types into bitmap, enum, struct order
. Promotes inline data types into the Data Types section and creates references for them
. Normalizes all references to the ref_PascalCase format
. Adds missing columns to tables for each type of section
. Reorders columns in tables to match disco ball order
. Renames headers in columns to match disco ball header names
. Re-formats access columns into disco ball order and spacing
. Re-formats constraint columns to be more readable (e.g. "60" -> "max 60", )
. Fixes command directions to client (<=/=>) server format
. Appends suffixes to sections when needed (e.g. "XyzBitmap" -> "XyzBitmap Type" or "MyField" -> "MyField Field")
. Uppercases all hexadecimal numbers
. Adds spaces after punctuation, when needed
. Adds labels to anchors when missing
. Removes extra spaces at the end of lines
. Fixes common section naming mistakes


	linkAttributes           bool
	addMissingColumns        bool
	reorderColumns           bool
	renameTableHeaders       bool
	formatAccess             bool
	promoteDataTypes         bool
	reorderSections          bool
	normalizeTableOptions    bool
	fixCommandDirection      bool
	appendSubsectionTypes    bool
	uppercaseHex             bool
	addSpaceAfterPunctuation bool
	removeExtraSpaces        bool