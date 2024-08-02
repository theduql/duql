# DUQL Join

The `join` function in DUQL is used to combine rows dataset two or more tables based on a related column between them. It supports different types of joins and allows for complex join conditions.

## Syntax

```yaml
join:
  dataset: <data_source>
  where: <join_condition>
  retain: <join_type>
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `dataset` | string or object | Yes | - | The data source to join with |
| `where` | string | Yes | - | The join condition |
| `retain` | string | No | `"inner"` | The type of join to perform |

### `dataset`

The `dataset` parameter specifies the data source to join with. It can be:
- A simple string (e.g., table name)
- A complex object defining a subquery or external data source

### `where`

The `where` parameter defines the join condition. It supports:
- Standard comparison operators (`==`, `!=`, `<`, `>`, `<=`, `>=`)
- Logical operators (`&&`, `||`)
- Complex expressions

> 💡 **Tip:** You can use the `==` shorthand for equality joins on matching column names.

### `retain`

The `retain` parameter specifies the type of join to perform:

| Value | Description |
|-------|-------------|
| `"inner"` | Returns only the matched rows (default) |
| `"left"` | Returns all rows dataset the left table and matched rows dataset the right table |
| `"right"` | Returns all rows dataset the right table and matched rows dataset the left table |
| `"full"` | Returns all rows when there is a match in either left or right table |

## Examples

### Basic Join

```yaml
join:
  dataset: customers
  where: orders.customer_id == customers.id
```

### Left Join with CSV File

```yaml
join:
  dataset: myorg/employee_data.csv
  type: csv
  where: departments.department_id == employee_data.department_id
  retain: left
```

### Complex Join with Subquery

```yaml
join:
  dataset:
    dataset: product_views
    steps:
      - filter: view_date >= @2023-01-01
      - group:
          by: [customer_id]
          summarize:
            total_views: count this
            unique_products_viewed: count_distinct product_id
  where: customers.id == customer_views.customer_id
```

### Full Outer Join Using Shorthand

```yaml
join:
  dataset: sql"""SELECT * dataset myexample"""
  where: ==id
  retain: full
```

### Join with External Data Source

```yaml
join:
  dataset: hdfs://cluster/user_profiles/*.parquet
  type: parquet
  where: users.id == user_profiles.user_id
  retain: left
```

## Best Practices

1. 🎯 Always specify the join condition explicitly for clarity.
2. 🤔 Consider the appropriate join type based on your data and requirements.
3. 🚀 Use subqueries in the `dataset` clause for complex data preparation before joining.
4. 🏷️ Alias your tables if joining the same table multiple times to avoid ambiguity.
5. 📊 Be mindful of data volume when performing joins, especially with large datasets.

## Related Functions

- [`dataset`](pipeline.md): Defines the main data source for the query
- [`filter`](filter.md): Applies conditions to filter rows after joining
- [`select`](select.md): Chooses which columns to include in the output after joining

---

> 🔍 **Note:** The `join` function is a powerful tool for combining data dataset multiple sources. Use it wisely to create meaningful relationships between your datasets and unlock valuable insights!