# Alchemy

Alchemy is a command line tool for modifying and transforming Matter spec documents.

## Installation

### Prebuilt Binaries

Download the latest releases from the [GitHub release page](https://github.com/hasty/alchemy/releases).

There are two builds for each architecture: plain and db. The db build includes the [DB Feature](#DB) below, but at the expense of a larger binary.

### Build from source

Alchemy is written in pure Go, so to build from source:

1. Install Go 1.21 or greater
2. Clone the alchemy repo
3. In the root of the alchemy repo, run ```go build```

## Commands

### Common flags

| Flag                      | Default       | Description   |	
| :------------------------ |:-------------:| :-------------|
| --serial          	      |	false         | Process Asciidoc files one-by-one instead of in parallel; slower
| --dryrun -d               | false         | Run all logic for the command, but do not write results to disk
| --verbose 	              |	false         | Display more verbose logging; best used with --serial
| --attribute ```<name of attribute>``` | empty string	| Sets an attribute for Asciidoc processing, e.g. "in-progress". This parameter can be specified multiple times for different attributes 


### format

Format does not alter the content of the document, but does make it easier to read the source:

- Aligns table cells so that table delimiters correctly line up with other cells in the same column
- Removes extraneous spacing in tables and lists
- Normalizes properties

#### Examples

Format a single document:

```console
alchemy format connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```

Format all documents in a directory:

```console
alchemy format connectedhomeip-spec/src/app_clusters/*.adoc
```

Recursively format all documents in the spec:

```console
alchemy format connectedhomeip-spec/src/\*\*/*.adoc
```

### disco

Disco-ball is more aggressive than format, and attempts to rewrite the document to the disco-ball standard:

- Everything format does
- Rearranges the document into disco-ball order for clusters and device-types
- Rearranges data types into bitmap, enum, struct order
- Promotes inline data types into the Data Types section and creates references for them
- Normalizes all references to the ref_PascalCase format
- Adds missing columns to tables for each type of section
- Reorders columns in tables to match disco ball order
- Renames headers in columns to match disco ball header names
- Re-formats access columns into disco ball order and spacing
- Re-formats constraint columns to be more readable (e.g. an uint with the constraint "0 to 60" -> "max 60", )
- Fixes command directions to client (<=/=>) server format
- Appends suffixes to sections when needed (e.g. "XyzBitmap" -> "XyzBitmap Type" or "MyField" -> "MyField Field")
- Uppercases all hexadecimal numbers
- Adds spaces after punctuation, when needed
- Adds labels to anchors when missing
- Removes extra spaces at the end of lines
- Fixes common section naming mistakes

| Flag                       | Default  | Description   |	
| :------------------------- |:--------:| :-------------|
| --linkIndexTables          | false    | Link table cells to child sections |
| --addMissingColumns        | true     | Add standard columns missing from tables
| --reorderColumns           | true     | Rearrange table columns into disco-ball order
| --renameTableHeaders       | true     | Rename table headers to disco-ball standard names
| --formatAccess             | true     | Reformat access columns in disco-ball order
| --promoteDataTypes         | true     | Promote inline data types to Data Types section
| --reorderSections          | true     | Reorder sections in disco-ball order
| --normalizeTableOptions    | true     | Remove existing table options and replace with standard disco-ball options
| --fixCommandDirection      | true     | Normalize command directions
| --appendSubsectionTypes    | true     | Add missing suffixes to data type sections (e.g. "Bit", "Value", "Field", etc.)
| --uppercaseHex             | true     | Uppercase hex values
| --addSpaceAfterPunctuation | true     | Add missing space after punctuation
| --removeExtraSpaces        | true     | Remove extraneous spaces

#### Examples

Disco-ball a single document:

```console
alchemy disco connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```


### zap

ZAP generates zap-template XMLs from a spec, creating new XML files for provisional clusters, and amending existing XML files with
changes.

| Flag                       | Description   |	
| :------------------------- | :-------------|
| --specRoot                 | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| --sdkRoot                  | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| --overwrite                | Overwrite existing XML files instead of amending them

> [!NOTE]  
> By default, existing ZAP XML files will be amended by Alchemy, leaving ordering of elements, comments and unrecognized XML attributes in place. The overwrite flag allows regenerating the XML files from scratch.


#### Generate ZAP files for a single cluster

```console
alchemy zap --attribute="in-progress"  --sdkRoot=./connectedhomeip/ --specRoot=./connectedhomeip-spec/ ./connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```
> [!NOTE]  
> Alchemy follows dependencies between clusters, so if the specified doc requires data types from other docs, it will also generate XML files for them as well. In the above case, Thermostat depends on an enumeration in OccupancySensor, so occupancy-sensing-cluster.xml will also be generated.

### compare

Compare loads the spec and the ZAP template XMLs and returns their differences in JSON format.

| Flag                       | Description   |	
| :------------------------- | :-------------|
| --specRoot                 | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| --sdkRoot                  | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| --text                     | Returns differences in a text format |

#### Example

```console
alchemy compare --sdkRoot=./connectedhomeip/ --specRoot=./connectedhomeip-spec/
```

### conformance

Conformance parses a provided conformance string and explains its meaning in plain English. It can also take a series of defined
identifiers and return whether the provided conformance would consider it mandatory, optional, etc.

#### Examples

```console
$ alchemy conformance "[LT | DF & CF]"
description: optional if (LT or (DF and CF))
```

```console
$ alchemy conformance "AB, [CD]"
description: mandatory if AB, otherwise optional if CD
$ alchemy conformance "AB, [CD]" AB CD
description: mandatory if AB, otherwise optional if CD
conformance: Mandatory
$ alchemy conformance "AB, [CD]" EF CD
description: mandatory if AB, otherwise optional if CD
conformance: Optional
$ alchemy conformance "AB, [CD]" EF
description: mandatory if AB, otherwise optional if CD
conformance: Disallowed
```

### dm

Data Model generates the Data Model XML files from the spec.

| Flag                       | Description   |	
| :------------------------- | :-------------|
| --specRoot                 | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| --sdkRoot                  | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |


### testplan

Testplan generates basic test plan adoc files from the spec.

| Flag                       | Description   |	
| :------------------------- | :-------------|
| --specRoot                 | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| --testRoot                 | The root of your clone of [the Matter Test Plan Specification](https://github.com/CHIP-Specifications/chip-test-plans) |
| --overwrite                | Overwrite existing XML files instead of amending them


> [!NOTE]  
> By default, existing test plan Asciidoc files will be ignored. The overwrite flag allows regenerating the test plan Asciidoc files from scratch; this will destroy any existing tests aside from basic validation of features, attributes, etc.

### db

DB loads up a set of spec docs or ZAP templates and exposes their contents as tables in a local MySQL server you can query.

| Flag                       | Description   |	
| :------------------------- | :-------------|
| --specRoot                 | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| --sdkRoot                  | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| --address                  | The address to bind the MySQL server to; defaults to localhost |
| --port                     | The port to bind the MySQL server to; defaults to 3306 |
| --raw                      | Populates the tables with the raw text of the associated entities, rather than parsing into an object model first |

