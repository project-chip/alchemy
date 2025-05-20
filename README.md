# Alchemy

Alchemy is a command line tool for modifying and transforming Matter spec documents.
<img src="./alchemy.svg" align="right"/>

It can:
- Format your spec documents, aligning table cells, removing unneeded spacing, etc.
- Disco-ball your spec documents, aligning them with Matter spec style guidelines
- Generate ZAP XML files for clusters and device types
- Generate Data Model XML files
- Generate basic test plans for clusters
- Present the Matter spec as a MySQL-compatible database to run queries against
- Print English-language explanations of Matter conformance strings

<br clear="right"/>

## Installation

### Prebuilt Binaries

Download the latest releases from the [GitHub release page](https://github.com/project-chip/alchemy/releases).

There are two builds for each architecture: plain and db. The db build includes the [DB Feature](#alchemy-db) below, but at the expense of a larger binary.

### Build from source

Alchemy is written in pure Go, so to build from source:

1. Install Go 1.24 or greater
2. Clone the alchemy repo
3. In the root of the alchemy repo, run ```go build```

## Commands

### Common flags

| Flag                                  | Default       | Description   |	
| :------------------------------------ |:-------------:| :-------------|
| `--serial`          	                | false         | Process Asciidoc files one-by-one instead of in parallel; slower
| `--dry-run` `-d`                          | false         | Run all logic for the command, but do not write results to disk
| `--patch`   	                        | false         | Write a patch file for any changes to stdout
| `--verbose` 	                        | false         | Display more verbose logging; best used with --serial
| `--attribute="<name of attribute>"` | empty string	| Sets an attribute for Asciidoc processing, e.g. "in-progress".<br/>This parameter can be specified multiple times for different attributes 


### format

Format does not alter the content of the document, but does make it easier to read the source:

- Aligns table cells so that table delimiters correctly line up with other cells in the same column
- Removes extraneous spacing in tables and lists
- Normalizes properties

#### Examples

| Flag                            | Default  | Description   |	
| :------------------------------ |:--------:| :-------------|
| --wrap                          | none     | The number of characters to wrap lines without disrupting Asciidoc syntax |

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
alchemy format connectedhomeip-spec/src/\*\*/\*.adoc
```

Word-wrap all documents in a directory to 120 characters:

```console
alchemy format connectedhomeip-spec/src/app_clusters/*.adoc --wrap=120
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

| Flag                                 | Default                | Description   |	
| :----------------------------------- |:----------------------:| :-------------|
| `--link-index-tables`                | false                  | Link table cells to child sections |
| `--add-missing-columns`              | true                   | Add standard columns missing from tables |
| `--reorder-columns`                  | true                   | Rearrange table columns into disco-ball order |
| `--rename-table-headers`             | true                   | Rename table headers to disco-ball standard names |
| `--format-access`                    | true                   | Reformat access columns in disco-ball order |
| `--promote-data-types`               | true                   | Promote inline data types to Data Types section |
| `--reorder-sections`                 | true                   | Reorder sections in disco-ball order |
| `--normalize-table-options`          | true                   | Remove existing table options and replace<br/>with standard disco-ball options |
| `--fix-command-direction`            | true                   | Normalize command directions |
| `--append-subsection-types`          | true                   | Add missing suffixes to data type sections<br/>(e.g. "Bit", "Value", "Field", etc.) |
| `--uppercase-hex`                    | true                   | Uppercase hex values |
| `--add-space-after-punctuation`      | true                   | Add missing space after punctuation |
| `--remove-extra-spaces`              | true                   | Remove extraneous spaces |
| `--normalize-feature-names`          | true                   | Normalize feature names to be compatible<br/>with downstream code generation |
| `--disambiguate-conformance-choice`  | false                  | Ensure that each document only uses each<br/>conformance choice identifier once |
| `--spec-root`                        | ./connectedhomeip-spec | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--wrap`                             | none                   | The number of characters to wrap lines<br/>without disrupting Asciidoc syntax |

#### Examples

Disco-ball a single document:

```console
alchemy disco connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```

Disco-ball the whole spec:

```console
alchemy disco --spec-root=./connectedhomeip-spec
```

Disco-ball a single document, but update any other documents if needed (e.g. rewriting a reference):

```console
alchemy disco --spec-root=./connectedhomeip-spec ./connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```

Disco-ball a single document, wrapping the text at 120 characters and linking table entries to their associated sections:

```console
alchemy disco connectedhomeip-spec/src/app_clusters/Thermostat.adoc --wrap=120 --link-index-tables
```

### zap

ZAP generates zap-template XMLs from a spec, creating new XML files for provisional clusters, and amending existing XML files with
changes.

| Flag                                       | Default                | Description   |	
| :----------------------------------------- |:----------------------:| :-------------|
| `--spec-root`                              | ./connectedhomeip-spec | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--sdk-root`                               | ./connectedhomeip      | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| `--overwrite`                              | false                  | Overwrite existing XML files instead of amending them
| `--force`                                  | false                  | Forces generation of XML files, even if there are parsing errors reading the spec
| `--ignore-errored`                         | false                  | Generates XML files for all provided spec files that had no parsing errors
| `--exclude=<file pattern>`                 |                        | Ignores a pattern of file paths for generation; this attribute may be provided multiple times
| `--provisional-policy=[none\|loose\|strict]` | none                   | Sets the provisional policy (see below)

> [!IMPORTANT]  
> ZAP generates the XML based on how the spec would render for the attributes provided. If you are attempting to generate the zap-template XML for a
> part of the spec that is hidden by an #ifdef by default, you will need to provide the necessary attributes to allow it to render.
>
> For example, if your cluster was included in the spec like so:
>
> ```
> ifdef::in-progress,my-cool-feature[]
> include::./MyAwesomeCluster.adoc[]
> endif::[]
> ```
>
> You would need to pass an `--attribute` flag for either `in-progress` or `my-cool-feature`:
>
> ```console
> alchemy zap --attribute="in-progress"  --sdk-root=./connectedhomeip/ --spec-root=./connectedhomeip-spec/
> ```

> [!WARNING]  
> By default, Alchemy will refuse to generate XML for spec files with errors. If you receive the error `Alchemy was unable to proceed due to the following fatal errors in parsing the spec`,
> you can either:
>
> 1. Fix the error. This is the best route.
> 2. Use the `--exclude` flag to exclude the file that is failing to parse.
> 3. Use the `--ignore-errored` flag to automatically exclude any files that are failing to parse
> 4. Use the `--force` flag to cause Alchemy to ignore all errors and generate XML as best as it can. This is likely to cause issues later, and should be avoided if possible.

> [!NOTE]  
> By default, existing ZAP XML files will be amended by Alchemy, leaving ordering of elements, comments and unrecognized XML attributes in place. The overwrite flag allows regenerating the XML files from scratch.

#### Generate ZAP files for a single cluster

```console
alchemy zap --attribute="in-progress"  --sdk-root=./connectedhomeip/ --spec-root=./connectedhomeip-spec/ ./connectedhomeip-spec/src/app_clusters/Thermostat.adoc
```
> [!NOTE]  
> Alchemy follows dependencies between clusters, so if the specified doc requires data types from other docs, it will also generate XML files for them as well. In the above case, Thermostat depends on an enumeration in OccupancySensor, so occupancy-sensing-cluster.xml will also be generated.
>
> If the targeted cluster references, directly or indirectly, one of the global data type files (e.g. global-structs.xml, global-enums.xml, etc.), the entire global data type file will be generated. This may pull in changes to global data types from other clusters.


#### Provisional Policy

The provisional policy changes how Alchemy handles adding new data types to the XML files.

##### None

The "none" policy causes Alchemy treat provisional data types the same way it did before these policies were introduced:

1. Clusters/attributes with provisional conformance are written with the "apiMaturity" attribute set to "provisional"; clusters/attributes without provisional conformance have this attribute removed, if it already exists
2. Structs which do not currently exist in the ZAP XML are written with "apiMaturity" set to "provisional"

##### Loose

The "loose" policy causes Alchemy to write or clear apiMaturity attributes on the following data types:

* Attributes
* Bitmaps
* Clusters
* Commands
* Enums
* Events
* Features
* Structs
* Fields on Commands, Events and Structs 

Data types are considered provisional unless one of the following is true:

1. The data type has a conformance column which does not have provisional conformance set
2. The data type is referenced by a data type which is not provisional

##### Strict

The "strict" policy has the same rules as the "loose" policy, except that it will refuse to add new data types to the generated XML if they are not provisional. 

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

| Flag                       | Default                                 | Description   |	
| :------------------------- |:---------------------------------------:| :-------------|
| `--spec-root`              | ./connectedhomeip-spec                  | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--dm-root`                | ./connectedhomeip/data_model/master     | The data model directory of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| `--force`                  | false                                   | Forces generation of data model files, even if there are parsing errors reading the spec
| `--ignore-errored`         | false                                   | Generates data model files for all provided spec files that had no parsing errors
| `--exclude=<file pattern>` |                                         | Ignores a pattern of file paths for generation; this attribute may be provided multiple times

### testplan

Testplan generates basic test plan adoc files from the spec.


| Flag                       | Default                | Description   |	
| :------------------------- |:----------------------:| :-------------|
| `--spec-root`              | ./connectedhomeip-spec | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--test-root`              | ./chip-test-plans      | The root of your clone of [the Matter test plans](https://github.com/CHIP-Specifications/chip-test-plans) |
| `--overwrite`              | false                  | Overwrite existing test plan files instead of amending them


> [!NOTE]  
> By default, any existing test plan Asciidoc files will be ignored. The overwrite flag allows regenerating the test plan Asciidoc files from scratch; this will destroy any existing tests aside from basic validation of features, attributes, etc.

### testscript

Testplan generates a basic test script from the spec.

| Flag                       | Default                | Description   |	
| :------------------------- |:----------------------:| :-------------|
| `--spec-root`              | ./connectedhomeip-spec | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--template-root`          |                        | The root of your local template files; if not specified, Alchemy will use an internal copy|
| `--overwrite`              | false                  | Overwrite existing test scripts files instead of amending them


> [!NOTE]  
> By default, any existing test script files will be ignored. The overwrite flag allows regenerating the test script files from scratch; this will destroy any existing tests aside from basic validation of features, attributes, etc.


### alchemy-db

Alchemy-db is provided as a separate binary. It loads up a set of spec docs or ZAP templates and exposes their contents as tables in a local MySQL server you can query.

| Flag                       | Default                | Description   |	
| :------------------------- |:----------------------:| :-------------|
| `--spec-root`              | ./connectedhomeip-spec | The root of your clone of [the Matter Specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/) |
| `--sdk-root`               | ./connectedhomeip      | The root of your clone of [the Matter SDK](https://github.com/project-chip/connectedhomeip/) |
| `--address`                | localhost              | The address to bind the MySQL server to |
| `--port`                   | 3306                   | The port to bind the MySQL server to |
| `--raw`                    | false                  | Populates the tables with the raw text of the associated entities,<br/> rather than parsing into an object model first |

#### Examples

##### Command line to launch, loading from specific paths
```console
alchemy-db --sdk-root=./connectedhomeip/ --spec-root=./connectedhomeip-spec/
```

##### Semantic tags used across multiple Namespaces
```sql
SELECT
    t.name AS tag,
    ns.name AS namespace,
    d.path
FROM
    tag AS t
    JOIN namespace AS ns ON t.namespace_id = ns.namespace_id
    JOIN document AS d on ns.document_id = d.document_id
WHERE
    t.name IN 
    (
        SELECT
            name
        FROM
            tag
        GROUP BY
            name
        HAVING
            COUNT(*) > 1
    )
ORDER BY
    t.name

```

##### Events with fields whose data types are enumerations with more than three values
```sql
WITH LargeEnums AS (
    SELECT
        enum_id,
        name
    FROM
        enum
    WHERE enum_id IN
    (
    SELECT 
        enum_id
    FROM   
        enum_value
    GROUP BY enum_id
    HAVING COUNT(*) > 3
    )
) 
SELECT 
    c.name AS cluster_name, 
    e.name AS event_name
FROM   
    event AS e
    JOIN cluster c ON e.cluster_id = c.id
WHERE 
    event_id IN 
    (
        SELECT 
            event_id
        FROM   
            event_field
        WHERE 
            data_type IN 
        (SELECT NAME FROM LargeEnums)
    );  
```