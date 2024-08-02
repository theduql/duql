# DUQL
> Dangerously Usable Query Language

<p align="center"><img src="https://tyler-mills-shared.s3.amazonaws.com/duql-logo.png" width="400"></p>

## What?
- A YAML-based query language
- Compiles to SQL & JSON
- [100% open-source](/LICENSE)
- A layer on top of [PRQL](https://github.com/PRQL/prql)

## Priorities
1. JSON Schema Definition
2. `duql` Utility
3. VSCode Extension
> [!IMPORTANT]  
> Seeking contributors! Please contact hello@assembled.studio

## Examples
### DUQL vs. SQL
#### DUQL
![basic-duql](https://tyler-mills-shared.s3.amazonaws.com/duql_0.png)
#### SQL
![basic-duql-sql-comparison](https://tyler-mills-shared.s3.amazonaws.com/duql_1_sql.png)

### Advanced Query
![duql-advanced](https://tyler-mills-shared.s3.amazonaws.com/Chalk+Screenshot.png)

## Why?
- Balanced: 
    - DUQL leverages YAML to strike a delicate balance between readability and verbosity.
- Logical: SQL isn't structured to map well to how humans think about querying and transforming data. 
    - DUQL is piped allowing you to think about transformations as sequential steps with clear "in" and "out".
- The Alternatives: 
    - SQL doesn't map well to how humans think about querying and transforming data
    - JSON is too verbose both when writing and reading
    - ORMs lock you into a specific programming language
    - PRQL queries become harder to read as queries become more complex
- Modern:
    - Queries will increasingly be generated by LLMs, not written; and DUQL is designed to maximize readability.
    - Validating a query and understanding how it works has never been easier.

## Install
> Coming Soon

## Writing / Editing
> [!NOTE]  
> DUQL is 100% valid YAML. As a result, you can leverage all your awesome YAML extensions that exist today.

> [!TIP]  
> The following line will add base validation and hints:
> `# yaml-language-server: $schema=https://json.schemastore.org/duql-schema-0.0.0.yml`

## Usage
### Basic
```console
duql -w
```
> [!TIP]  
>  Runs in `watch` mode, generating and updating `.sql` files for `.duql` file changes in current directory.
> Outputs to `duql` subdirectory.

### Advanced
```console
NAME
    duql - generate sql, json, prql from .duql files

SYNOPSIS
    duql [OPTIONS]

DESCRIPTION
    Generates sql, json, or prql files from .duql files in the input directory.
    When run without options, duql defaults to watch mode with default input and output.

OPTIONS
    -i, --input INPUT
        Specify input directory (default: current directory)

    -o, --output OUTPUT
        Specify output directory (default: 'duql' subdirectory)

    -f, --format FORMAT
        Specify output format (default: sql)

    -w, --watch
        Watch for changes and generate automatically (default when no options are provided)

    --wtf
        Generate a friendly summary using API key

EXAMPLES
    duql
        Runs in watch mode with default input and output

    duql --input ./queries --output ./generated --format json
    duql -i ./queries -o ./generated -f json -w
    duql --input ./queries --wtf

```
