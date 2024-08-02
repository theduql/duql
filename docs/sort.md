# DUQL Sort

The `sort` function in DUQL is used to specify the order of rows in the output. It supports sorting by multiple columns, in ascending or descending order, and can use complex expressions for sorting criteria.

## Syntax

```yaml
sort: <sorting_criteria>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `sort` | string, array, or object | Yes | Specifies the sorting criteria |

## Behavior

- Sorts the dataset based on the specified criteria.
- Can sort by single or multiple columns.
- Supports ascending (default) and descending order.
- Can use expressions and SQL functions for complex sorting.

## Examples

### Sort by a Single Column

```yaml
sort: last_name
```

### Sort in Descending Order

```yaml
sort: -age
```

### Sort by Multiple Columns

```yaml
sort: [department, -salary]
```

### Sort with Complex Criteria

```yaml
sort:
  - last_name
  - first_name
  - -hire_date
```

### Sort with Column Renaming

```yaml
sort:
  Order Date: order_date
  "Total Amount": -order_total
```

### Sort with Expressions

```yaml
sort:
  - category
  - -sum(sales_amount)
```

### Sort with SQL Functions

```yaml
sort:
  sql: "CASE WHEN status = 'urgent' THEN 0 ELSE 1 END, created_at DESC"
```

## Best Practices

1. 🎯 Choose sorting criteria that align with your analysis goals.
2. 🔢 Use descending order (prefix with `-`) for "top N" type queries.
3. 📊 Consider performance implications when sorting large datasets.
4. 🧮 Leverage expressions and SQL functions for complex sorting logic.
5. 🚀 Place `sort` before `take` when you want to limit results based on the sort order.
6. 🔍 Be explicit about sort order for clarity, even when using the default ascending order.

## Real-World Use Case

Here's an example of a DUQL query that uses `sort` to analyze customer purchase behavior:

```yaml
dataset: orders

steps:
  - join:
      dataset: customers
      where: orders.customer_id == customers.id
  - generate:
      total_amount: price * quantity
      days_since_last_order:
        sql'DATEDIFF(CURRENT_DATE, MAX(order_date) OVER (PARTITION BY customer_id))'
  - group:
      by: [customer_id, customers.name, customers.email]
      summarize:
        total_spent: sum(total_amount)
        avg_order_value: avg(total_amount)
        order_count: count(order_id)
        last_order_date: max(order_date)
  - generate:
      customer_value:
        case:
          - total_spent > 10000: "High"
          - total_spent > 5000: "Medium"
          - true: "Low"
  - sort:
      - -total_spent
      - -order_count
      - last_order_date
  - take: 100

into: top_customers_analysis
```

This query demonstrates:
1. Joining order data with customer information
2. Calculating total amount and days since last order
3. Grouping and summarizing by customer
4. Generating a customer value category
5. Sorting by multiple criteria:
   - Total spent (descending)
   - Order count (descending)
   - Last order date (ascending)
6. Taking the top 100 customers based on this sorting

The `sort` step ensures that we get the most valuable customers, prioritizing those who have spent more, made more orders, and purchased more recently.

---

> 💡 **Tip:** The `sort` function is crucial for organizing your data meaningfully. Combine it with `take` to efficiently retrieve the most important subset of your data based on your sorting criteria!