# Summarize

The `summarize` function in DUQL is used to perform aggregations on your data, condensing multiple rows into summary statistics. It can be used both as a top-level function and within the `group` function, offering flexibility in how you aggregate your data.

## Syntax

As a top-level function:

```yaml
summarize:
  <new_column_name>: <aggregation_function>
```

Within a `group` function:

```yaml
group:
  by: <grouping_columns>
  steps:
    summarize:
      <new_column_name>: <aggregation_function>
```

## Behavior

1. **Top-level usage**: When used at the top level, `summarize` produces a single row summarizing the entire dataset.
2. **Within `group`**: When used inside a `group` function, it produces one summary row for each unique combination of the grouping columns.

This dual behavior allows for both global summaries and group-specific summaries within the same query language.

## Common Aggregation Functions

Currently, all declared aggregation functions are `min`, `max`, `count`, `average`, `stddev`, `avg`, `sum` and `count_distinct`

## Examples

### Top-level Summarize

```yaml
dataset: sales

steps:
  - summarize:
      total_revenue: sum amount
      average_order_value: avg amount
      order_count: count order_id
      unique_customers: count_distinct customer_id
```

This query will produce a single row with overall summary statistics for the entire sales dataset.

### Summarize Within Group

```yaml
dataset: sales

steps:
  - group:
      by: [category, year(order_date)]
      summarize:
        total_revenue: sum amount
        average_order_value: avg amount
        order_count: count order_id
```

This query will produce summary statistics for each unique combination of category and year.

### Combined Usage

```yaml
dataset: sales

steps:
  - group:
      by: category
      summarize:
        category_revenue: sum amount
        category_orders: count order_id
  - summarize:
      total_revenue: sum category_revenue
      total_orders: sum category_orders
      num_categories: count category
```

This query first summarizes data by category, then provides an overall summary of those category summaries.

## Best Practices

1. 🎯 Choose appropriate aggregation functions for your data types and analysis goals.
2. 🏷️ Use clear and descriptive names for your summarized columns.
3. 🧮 Be aware of how `null` values are handled in your aggregations.
4. 🚀 Consider performance implications when summarizing large datasets.
5. 📊 Use top-level `summarize` for overall statistics and `group` with `summarize` for segmented analysis.
6. 🔍 Validate your summary results, especially when using complex aggregation expressions.

## Real-World Use Case

Here's an example of a DUQL query that uses `summarize` both within `group` and at the top level to analyze sales performance:

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
    by: [category, month(sale_date)]
    summarize:
      monthly_revenue: sum revenue
      monthly_quantity: sum quantity
      avg_margin: avg margin
      product_count: count_distinct product_id
- sort: [category, month(sale_date)]
- summarize:
    total_revenue: sum monthly_revenue
    total_quantity: sum monthly_quantity
    overall_avg_margin: avg avg_margin
    peak_monthly_revenue: max monthly_revenue
    total_products: sum product_count

into: sales_performance_summary
```

This query demonstrates:

1. Filtering recent sales data
2. Joining with product information
3. Calculating revenue and margin
4. Grouping by category and month, with monthly summaries
5. Sorting the grouped results
6. Providing an overall summary of the grouped data

The resulting `sales_performance_summary` dataset offers both detailed monthly performance by category and overall performance metrics, showcasing the flexibility of the `summarize` function in DUQL.

***

> 💡 **Tip:** The `summarize` function is a versatile tool in DUQL. Use it at the top level for quick overall insights, or within `group` for detailed breakdowns. Combine both approaches to create comprehensive analytical queries!
