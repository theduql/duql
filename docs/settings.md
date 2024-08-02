# DUQL Settings

The `settings` component in DUQL allows you to specify metadata and configuration options for your query. This includes version information and target database specifications, which can affect how your query is processed and executed.

## Syntax

```yaml
settings:
  version: <version_string>
  target: <target_database>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `version` | string | Yes | The version of DUQL being used |
| `target` | string | Yes | The target database or SQL dialect for the query |

### Supported Targets

- `sql.clickhouse`
- `sql.duckdb`
- `sql.generic`
- `sql.glaredb`
- `sql.mysql`
- `sql.postgres`
- `sql.sqlite`

## Examples

### Basic Settings

```yaml
settings:
  version: '0.0.1'
  target: sql.postgres
```

### Settings for a Different Database

```yaml
settings:
  version: '0.0.2'
  target: sql.clickhouse
```

## Use Cases

1. **Version Control**: Specify the DUQL version to ensure compatibility with the parser and runtime environment.

2. **Database-Specific Optimizations**: The `target` setting allows the DUQL engine to generate optimized SQL for specific database systems.

3. **Feature Availability**: Certain DUQL features may only be available in specific versions or for certain database targets.

4. **Query Portability**: By explicitly stating the target, you can ensure that your query will work as expected when moved between different database environments.

## Best Practices

1. 📌 Always specify both the `version` and `target` in your settings to ensure consistent behavior across different environments.

2. 🔄 Keep your DUQL version up-to-date to leverage the latest features and improvements. Check the changelog when upgrading to be aware of any breaking changes.

3. 🎯 Choose the most specific target that matches your database system. Using `sql.generic` may result in suboptimal performance or feature limitations.

4. 📚 Familiarize yourself with the specific features and limitations of your chosen target database to make the most of DUQL's capabilities.

5. 🧪 When developing queries that need to work across multiple database systems, test with each target to ensure compatibility.

## Related Components

- All DUQL components are affected by the `settings`, as they determine the available features and the SQL generation process.

## Limitations and Considerations

- The available features and syntax may vary depending on the specified `version` and `target`. Consult the DUQL documentation for your specific version and target for detailed information.
- Some advanced or database-specific features may not be available when using the `sql.generic` target.
- Changing the `target` setting may require adjustments to your query if you've used any database-specific functions or syntax.

---

> 💡 **Tip:** The `settings` component is crucial for ensuring your DUQL queries are interpreted and executed correctly. Always start your DUQL files with a clear `settings` section to make your queries more robust and portable across different environments!