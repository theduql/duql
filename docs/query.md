# DUQL Query

DUQL is a powerful and intuitive query language designed for data transformation and analysis. It provides a structured way to define complex data operations using a human-readable YAML format.

## Syntax

A typical DUQL query has the following structure:

```yaml
settings:
  version: <duql_version>
  target: <target_database>

declare:
  <variable_declarations>

dataset: <main_data_source>

steps:
  - <transformation_step_1>
  - <transformation_step_2>
  # ... more steps as needed

into: <output_destination>
```

## Components

### Settings

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `version` | string | No | The version of DUQL being used |
| `target` | string | No | The target database or SQL dialect for the query |

Example:
```yaml
settings:
  version: '0.0.1'
  target: sql.postgres
```

### Declare

The `declare` section allows you to define variables, functions, or reusable query components.

Example:
```yaml
declare:
  active_users:
    dataset: users
    steps:
      - filter: last_login > @2023-01-01
  
  calculate_age: birth_date -> datediff(years, birth_date, current_date())
```

### Dataset

The `dataset` component specifies the main data source for your query. It can be a simple table name or a more complex definition.

Example:
```yaml
dataset: sales_transactions
```

or 

```yaml
dataset:
  name: s3://data-bucket/sales/*.parquet
  format: parquet
```

### Steps

The `steps` section is where you define your data transformation pipeline. Each step represents an operation on your data.

Available steps include:
- `filter`: Select rows based on conditions
- `join`: Combine data from multiple sources
- `select`: Choose or compute columns
- `group`: Aggregate data
- `sort`: Order results
- `take`: Limit the number of rows
- `generate`: Create new columns
- And more...

Example:
```yaml
steps:
  - filter: date > @2023-01-01
  - join:
      dataset: products
      where: sales_transactions.product_id == products.id
  - group:
      by: [category, product_name]
      summarize:
        total_sales: sum amount
  - sort: -total_sales
  - take: 10
```

### Into

The `into` component specifies the destination for your query results.

Example:
```yaml
into: monthly_sales_report
```

## Best Practices

1. 🏗️ Structure your query logically, starting with data sources and progressing through transformations.
2. 🧩 Use the `declare` section to create reusable components and improve query readability.
3. 📊 Leverage the power of DUQL's expressive steps to perform complex data operations.
4. 🎯 Be explicit in your joins and filters to ensure data integrity.
5. 🚀 Optimize your query by placing filters early in the pipeline and using appropriate indexes.
6. 📝 Use meaningful names for variables and result sets to enhance query understandability.
7. 🔍 Take advantage of DUQL's ability to work with various data sources and formats.

## Real-World Use Case

Here's an example of a DUQL query that analyzes customer purchasing behavior:

```yaml
settings:
  version: '0.0.1'
  target: sql.postgres

declare:
  recent_customers:
    dataset: customers
    steps:
      - filter: last_purchase > @2023-01-01

dataset: orders

steps:
  - join:
      dataset: recent_customers
      where: orders.customer_id == recent_customers.id
  - join:
      dataset: products
      where: orders.product_id == products.id
  - generate:
      total_amount: quantity * price
      purchase_month: date_trunc('month', order_date)
  - group:
      by: [customer_id, purchase_month, category]
      summarize:
        total_spent: sum total_amount
        num_orders: count order_id
  - sort: [customer_id, purchase_month, -total_spent]
  - generate:
      customer_value:
        case:
          - total_spent > 1000: "High"
          - total_spent > 500: "Medium"
          - true: "Low"

into: customer_purchase_analysis
```

This query:
1. Defines recent customers
2. Joins order data with customer and product information
3. Calculates total amount per order
4. Groups and summarizes purchases by customer, month, and product category
5. Sorts the results
6. Categorizes customers based on their spending

The resulting `customer_purchase_analysis` dataset provides valuable insights into customer behavior, allowing for targeted marketing strategies and improved customer relationship management.

---