# DUQL Dataset

The `dataset` component in DUQL is a fundamental element that specifies the source of data for your query. It can represent a wide range of data sources, from simple table names to complex file systems or even subqueries.

## Syntax

The `dataset` can be defined in two main ways:

### Simple Form
```yaml
dataset: <table_name_or_simple_source>
```

### Advanced Form
```yaml
dataset:
  name: <data_source_name_or_path>
  format: <data_format>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | The name or path of the data source |
| `format` | string | No | The format of the data (default: `table`) |

### Supported Formats

- `table`: A database table (default)
- `csv`: Comma-separated values file
- `json`: JSON file or data
- `parquet`: Apache Parquet file

## Examples

### Basic Table Reference

```yaml
dataset: customers
```

This simple form refers to a table named `customers` in the default database.

### CSV File

```yaml
dataset:
  name: /path/to/sales_data.csv
  format: csv
```

This defines a dataset from a CSV file located at the specified path.

### Remote Parquet Files

```yaml
dataset:
  name: s3://my-bucket/transactions/*.parquet
  format: parquet
```

This example shows how to reference Parquet files stored in an Amazon S3 bucket.

### SQL Query

```yaml
dataset: sql"""
  SELECT * FROM orders
  WHERE order_date >= '2023-01-01'
"""
```

You can use raw SQL queries as datasets, which is useful for complex data sources or when transitioning from SQL to DUQL.

### Subquery

```yaml
dataset:
  name: recent_orders
  steps:
    - dataset: orders
    - filter: order_date >= @2023-01-01
```

This advanced example shows how to define a dataset based on a subquery, which can be reused in multiple parts of your DUQL script.

## Best Practices

1. 📁 Use meaningful names for your datasets to improve query readability.
2. 🔒 Ensure you have the necessary permissions to access the specified data sources.
3. 🚀 For large datasets, consider using efficient file formats like Parquet for better performance.
4. 🧩 Leverage subqueries to create reusable dataset components.
5. 📊 When working with files, use wildcards to process multiple files in a single query.
6. 🔍 Always validate the data format and structure of external sources before using them in complex queries.
7. 🔧 When using SQL strings (`sql''`), be aware that they bypass DUQL's syntax checking and optimization. Use them judiciously and consider requesting native DUQL features for frequently used SQL patterns.

## Real-World Use Case

Here's an example of a DUQL query that combines data from multiple sources:

```yaml
declare:
  recent_customers:
    dataset: customers
    steps:
      - filter: last_purchase_date >= @2023-01-01

dataset:
  name: s3://sales-data/transactions/*.parquet
  format: parquet

steps:
  - filter: transaction_date >= @2023-01-01
  - join:
      dataset: recent_customers
      where: transactions.customer_id == recent_customers.id
  - join:
      dataset: sql"""
        SELECT product_id, name, category
        FROM products
        WHERE is_active = true
      """
      where: transactions.product_id == products.product_id
  - group:
      by: [category, product_id, name]
      summarize:
        total_sales: sum(amount)
        num_transactions: count(transaction_id)
  - sort: -total_sales
  - take: 100

into: top_products_analysis
```

This query demonstrates:
1. Declaring a reusable dataset (`recent_customers`)
2. Using a Parquet file from S3 as the main dataset
3. Joining with the declared dataset and an SQL-defined product dataset
4. Performing aggregations and sorting on the combined data

The resulting `top_products_analysis` dataset provides insights into the best-selling products among recent customers, combining data from multiple sources and formats.

---

> 💡 **Tip:** The flexibility of DUQL's `dataset` component allows you to work with a wide variety of data sources. This makes it easy to integrate data from different systems and formats into a single, cohesive analysis! (Note: Only DuckDB currently supports non-table source.)