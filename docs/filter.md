# DUQL Filter

The `filter` function in DUQL is used to select rows from a dataset based on specified conditions. It allows you to narrow down your data to only the records that meet certain criteria.

## Syntax

```yaml
filter: <condition>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `filter` | string | Yes | A boolean expression representing the filter condition |

## Behavior

- Evaluates the specified condition for each row in the dataset.
- Retains only the rows where the condition evaluates to true.
- Can use complex expressions involving multiple columns, functions, and operators.

## Examples

### Basic Comparison

```yaml
filter: age > 18
```

### Date Range

```yaml
filter: (created_at | in @1776-07-04..@1787-09-17)
```

### Numeric Range

```yaml
filter: (magnitude | in 50..100)
```

### String Matching

```yaml
filter: (lower(name) | like '%smith%') && age >= 21
```

### Date Calculation

```yaml
filter: order_date > current_date() - interval '30 days'
```

### Using Parameters

```yaml
filter: id == $1
```

### Array Contains

```yaml
filter: array_contains(tags, 'urgent') && status != 'completed'
```

### Complex Conditions

```yaml
filter: (revenue - costs) / revenue > 0.2 && date.year(order_date) == 2023
```

### Nested Conditions

```yaml
filter: age >= 18 && (has_parental_consent || country != "US")
```

### Using SQL Functions

```yaml
filter: sql'regexp_contains(title, ''([a-z0-9]*-){{2,}}'')'
```

## Best Practices

1. 🎯 Be as specific as possible with your filter conditions to improve query performance.
2. 🔍 Use parentheses to clarify the order of operations in complex conditions.
3. 🧪 Test your filter conditions with sample data to ensure they produce expected results.
4. 📊 Consider using multiple `filter` steps for complex logic, improving readability.
5. 🚀 Place filters early in your query pipeline to reduce data volume for subsequent operations.
6. 🔢 Use parameters (`$1`, `$2`, etc.) for dynamic filtering when appropriate.

## Real-World Use Case

Here's an example of a DUQL query that uses multiple `filter` steps to analyze high-value orders from loyal customers:

```yaml
dataset: orders

steps:
- filter: order_date >= @2023-01-01
- join:
    dataset: customers
    where: orders.customer_id == customers.id
- filter: customers.account_created_at < @2022-01-01  # Loyal customers
- join:
    dataset: products
    where: orders.product_id == products.id
- generate:
    order_total: quantity * price
    is_discounted: discount_rate > 0
- filter: order_total > 1000  # High-value orders
- group:
    by: [customer_id, customers.name]
    summarize:
    total_spent: sum(order_total)
    order_count: count(order_id)
    discounted_orders: sum(is_discounted)
- filter: order_count >= 3  # Repeat high-value customers
- generate:
    average_order_value: total_spent / order_count
    discount_rate: discounted_orders / order_count
- filter: discount_rate < 0.5  # Mostly full-price purchases
- sort: -total_spent
- take: 100

into: top_loyal_customers
```

This query demonstrates:
1. Filtering orders by date
2. Joining with customer and product data
3. Filtering for loyal customers (account created before 2022)
4. Calculating order totals and identifying discounted orders
5. Filtering for high-value orders (over $1000)
6. Aggregating customer data
7. Filtering for repeat high-value customers (3 or more orders)
8. Calculating average order value and discount rate
9. Filtering for customers who mostly pay full price
10. Sorting and limiting the results

The resulting `top_loyal_customers` dataset provides a focused view of the most valuable loyal customers who frequently make large purchases without heavy reliance on discounts.

---

> 💡 **Tip:** The `filter` function is your primary tool for data selection in DUQL. Use it strategically throughout your query pipeline to focus your analysis on the most relevant data. Combine it with `generate` and `group` functions for powerful data insights!