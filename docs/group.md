# DUQL Group

The `group` function in DUQL is used to aggregate data based on specified columns. It allows you to perform calculations across sets of rows that have the same values in one or more columns.

## Syntax

```yaml
group:
  by: <grouping_columns>
  summarize:
    <new_column_name>: <aggregation_function>
  sort: <sorting_criteria>
  take: <limit>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `by` | string or array | Yes | Columns to group by |
| `summarize` | object | No | Aggregations to perform on grouped data |
| `sort` | string or array | No | Columns to sort the grouped results by |
| `take` | integer or string | No | Number of rows to return or a range |

### `by`

Specifies the columns to group by. It can be:
- A single column name
- An array of column names
- `table.*` to group by all columns of a table

### `summarize`

An object where each key is a new column name, and the value is an aggregation function or expression.

Common aggregation functions:
- `count`: Count of rows
- `sum`: Sum of values
- `avg`: Average of values
- `min`: Minimum value
- `max`: Maximum value
- `count_distinct`: Count of unique values

### `sort`

Specifies the sorting order for the grouped results. Can be:
- A single column name
- An array of column names
- Prefix with `-` for descending order

### `take`

Limits the number of rows in the result. Can be:
- An integer (e.g., `10`)
- A range (e.g., `'1..10'`)

## Examples

### Basic Grouping

```yaml
group:
  by: category
  summarize:
    total_sales: sum amount
```

### Multiple Grouping Columns

```yaml
group:
  by: [year(order_date), month(order_date)]
  summarize:
    order_count: count order_id
    total_revenue: sum total_amount
```

### Grouping with Sorting and Limit

```yaml
group:
  by: customer_id
  summarize:
    last_purchase: max order_date
    total_spent: sum amount
  sort: -total_spent
  take: 10
```

### Grouping by All Columns

```yaml
group:
  by: employees.*
  take: 1
```

### Complex Aggregations

```yaml
group:
  by: [category, subcategory]
  summarize:
    item_count: count product_id
    total_value: sum (price * stock_quantity)
    avg_price: avg price
    price_range: max price - min price
  sort: -total_value
  take: 20
```

## Best Practices

1. 🎯 Choose grouping columns carefully to ensure meaningful aggregations.
2. 🧮 Use appropriate aggregation functions for your data types and analysis goals.
3. 🏷️ Provide clear names for aggregated columns to improve query readability.
4. 🚀 Consider performance when grouping by many columns or on large datasets.
5. 📊 Use sorting and limits to focus on the most important results.
6. 🔍 Validate your grouping results, especially when using complex expressions.

## Real-World Use Case

Here's an example of a DUQL query that uses grouping to analyze sales data:

```yaml
declare:
  recent_date: @2023-01-01
  calculate_margin: revenue cost -> (revenue - cost) / revenue * 100

dataset: sales

steps:
  - filter: sale_date >= recent_date
  - join:
      dataset: products
      where: sales.product_id == products.id
  - generate:
      revenue: price * quantity
      margin: calculate_margin(revenue, cost)
  - group:
      by: [category, product_name]
      summarize:
        total_revenue: sum revenue
        total_quantity: sum quantity
        avg_margin: avg margin
        num_transactions: count sale_id
  - sort: -total_revenue
  - generate:
    category_rank:
      s'ROW_NUMBER() OVER (PARTITION BY category ORDER BY total_revenue DESC)'
  - filter: category_rank <= 5

into: top_products_by_category
```

This query demonstrates:
1. Filtering recent sales data
2. Joining with product information
3. Calculating revenue and margin
4. Grouping by product category and name
5. Aggregating various metrics (revenue, quantity, margin, transaction count)
6. Sorting by total revenue
7. Ranking products within each category
8. Filtering to show only the top 5 products per category

The resulting `top_products_by_category` dataset provides a comprehensive view of the best-performing products in each category, considering multiple performance metrics.

---

> 💡 **Tip:** The `group` function is a powerful tool for summarizing and analyzing your data. Combine it with other DUQL functions like `generate` and `filter` to create insightful analytical queries!