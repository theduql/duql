---
coverY: 0
---

# Groups

The `group` function in DUQL is a powerful tool for aggregating and analyzing data based on specified columns. It supports a wide range of operations, allowing for complex data transformations and analyses within grouped data.

## Syntax

```yaml
group:
  by: <grouping_columns>
  <additional_steps>
```

## Parameters

| Parameter          | Type            | Required           | Description                                                                                                 |
| ------------------ | --------------- | ------------------ | ----------------------------------------------------------------------------------------------------------- |
| `by`               | string or array | Yes                | Columns to group by. Can be a single column, an array of columns, or 'table.\*' for all columns of a table. |
| `additional_steps` | object          | Yes (at least one) | One or more additional transformation steps to apply to the grouped data.                                   |

## Additional Steps

The `group` function can include any of the following steps after grouping:

* `summarize`: Perform aggregations on the grouped data.
* `filter`: Apply conditions to filter the grouped results.
* `generate`: Create new columns based on grouped data.
* `join`: Join the grouped data with another dataset.
* `select`: Choose specific columns from the grouped results.
* `sort`: Order the grouped results.
* `take`: Limit the number of grouped results returned.
* `window`: Perform window functions on the grouped data.
* `loop`: Perform iterative operations on the grouped data.

## Examples

### Basic Grouping with Aggregation and Filtering

```yaml
group:
  by: department
  steps:
  - summarize:
      avg_salary: average salary
      employee_count: count employee_id
  - filter: avg_salary > 50000
  - sort: -avg_salary
  - take: 5
```

This example groups employees by department, calculates average salary and employee count, filters for departments with high average salaries, sorts by average salary descending, and takes the top 5 results.

### Time-based Analysis with Custom Metrics

```yaml
group:
  by: [year(transaction_date), month(transaction_date)]
  generate:
    monthly_revenue: sum(transaction_amount)
  window:
    revenue_growth:
      function: (monthly_revenue - lag(monthly_revenue)) / lag(monthly_revenue) * 100
      over:
        sort: [year(transaction_date), month(transaction_date)]
  filter: revenue_growth < 0
  sort: [year(transaction_date), month(transaction_date)]
```

This example demonstrates a monthly revenue analysis, calculating month-over-month growth and identifying periods of negative growth.

### Complex Business Logic with Case Statements

```yaml
group:
  by: product_id
  generate:
    stock_status:
      case:
        - inventory_count == 0: "Out of Stock"
        - inventory_count < reorder_point: "Low Stock"
        - true: "In Stock"
  filter: stock_status != "In Stock"
  sort: [stock_status, product_id]
```

This query categorizes products based on their inventory status and filters for those needing attention.

### Advanced Analytics with Multiple Steps

```yaml
group:
  by: [store_id, date(sale_timestamp)]
  generate:
    daily_sales: sum(sale_amount)
  window:
    moving_average:
      function: avg(daily_sales)
      over:
        partition: [store_id]
        rows: -6..0
  filter: daily_sales < moving_average * 0.8
  sort: [store_id, date(sale_timestamp)]
```

This complex example calculates daily sales by store, computes a 7-day moving average, and identifies days where sales were significantly below average.

## Best Practices

1. 🎯 Choose grouping columns carefully to ensure meaningful aggregations.
2. 📊 Use `summarize` or `generate` to calculate aggregate metrics.
3. 🔍 Leverage `filter` after aggregation to focus on important results.
4. 📈 Utilize `window` functions for advanced analytics within groups.
5. 🔢 Use `sort` and `take` to prioritize and limit results for better performance.
6. 🧮 Combine multiple steps for complex analyses in a single, readable query.

## Related Functions

* [`summarize`](summarize.md): Often used within `group` to perform aggregations.
* [`filter`](../basic/filter.md): Can be used before or after grouping to refine the dataset.
* [`generate`](generate.md): Useful for creating new columns based on grouped data.
* [`window`](../advanced/window.md): Enables advanced analytical functions within grouped data.

## Limitations and Considerations

* Grouping operations can be computationally expensive on large datasets. Use indexing strategies on grouping columns when possible.
* Be mindful of the order of operations within the `group` function, as it can affect the final results.
* Some combinations of steps might not be logically valid or may produce unexpected results. Always test your queries thoroughly.

***

> 💡 **Tip:** The `group` function is the cornerstone of data analysis in DUQL. By combining it with various additional steps, you can perform complex aggregations, transformations, and analyses all within a single, readable query structure. Experiment with different combinations to unlock deeper insights from your data!
