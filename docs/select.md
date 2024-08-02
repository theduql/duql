# DUQL Select

The `select` function in DUQL is used to specify which columns to include in the output, rename columns, compute new columns, or select all columns from a specific table. There's also a `select!` variant used to exclude specified columns.

## Syntax

```yaml
select:
  <new_column_name>: <column_or_expression>
```

or

```yaml
select: [<column1>, <column2>, ...]
```

or

```yaml
select!: [<column1_to_exclude>, <column2_to_exclude>, ...]
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `select` | object or array | Yes (if not using `select!`) | Columns to include or rename |
| `select!` | array | Yes (if not using `select`) | Columns to exclude |

## Behavior

- When `select` is an object, it allows renaming and computing new columns.
- When `select` is an array, it simply selects the specified columns.
- `select!` excludes the specified columns and keeps all others.
- Using `table.*` selects all columns from a specific table.

## Examples

### Basic Column Selection

```yaml
select: [id, name, email]
```

### Renaming and Computing Columns

```yaml
select:
  Customer ID: customer_id
  Full Name: f"{first_name} {last_name}"
  Total Spent: sum(order_total)
```

### Selecting All Columns from a Table

```yaml
select: employees.*
```

### Excluding Columns

```yaml
select!: [password, credit_card_number]
```

### Mixed Selection with Computation

```yaml
select:
  - order_id
  - customer_name
  - order_date
  Total: price * quantity
```

### Complex Expressions

```yaml
select:
  BMI: weight / (height / 100) ^ 2
  Weight Status:
    case:
      - BMI < 18.5: "Underweight"
      - BMI < 25: "Normal"
      - BMI < 30: "Overweight"
      - true: "Obese"
```

### Selecting from Multiple Tables

```yaml
select:
  - products.*
  - inventory.quantity_on_hand
```

### Using SQL Functions

```yaml
select:
  User ID: id
  Username: username
  Last Login: 
    sql'DATE_FORMAT(last_login, ''%Y-%m-%d %H:%i:%s'')'
```

### Window Functions

```yaml
select:
  - category
  - product_name
  Sales: sum(quantity * price)
  "% of Total Sales": 
    sql'SUM(quantity * price) / SUM(SUM(quantity * price)) OVER () * 100'
```

## Best Practices

1. 🎯 Be explicit about which columns you need to improve query performance and readability.
2. 🏷️ Use meaningful names when renaming columns or creating computed columns.
3. 🧮 Leverage the power of expressions and case statements for complex transformations.
4. 🔍 Use `select!` when it's easier to specify columns to exclude rather than include.
5. 📊 Be cautious when using `table.*` as it may include unnecessary columns.
6. 🚀 Place computationally intensive selections later in your query pipeline for better performance.

## Real-World Use Case

Here's an example of a DUQL query that uses `select` to prepare a comprehensive customer report:

```yaml
dataset: customers

steps:
  - join:
      dataset: orders
      where: customers.id == orders.customer_id
  - join:
      dataset: products
      where: orders.product_id == products.id
  - group:
      by: customers.id
      summarize:
        total_spent: sum(orders.amount)
        order_count: count(orders.id)
        last_order_date: max(orders.order_date)
  - select:
      Customer ID: customers.id
      Full Name: f"{customers.first_name} {customers.last_name}"
      Email: customers.email
      Total Spent: total_spent
      Average Order Value: total_spent / order_count
      Order Count: order_count
      Last Order Date: last_order_date
      Customer Since: 
        sql'DATE_FORMAT(customers.created_at, ''%Y-%m-%d'')'
      Loyalty Status:
        case:
          - order_count > 10: "VIP"
          - order_count > 5: "Regular"
          - true: "New"
      Days Since Last Order:
        sql'DATEDIFF(CURRENT_DATE, last_order_date)'

into: customer_summary_report
```

This query demonstrates:
1. Joining customer data with orders and products
2. Aggregating order information
3. Selecting and renaming relevant columns
4. Computing new columns using various expressions and SQL functions
5. Creating a categorical column based on order count

The resulting `customer_summary_report` provides a comprehensive view of customer activity and value, showcasing the versatility of the `select` function in DUQL.

---

> 💡 **Tip:** The `select` function is your tool for shaping the final output of your DUQL query. Use it to focus on the most relevant data, create meaningful derivations, and present your results in the most useful format for your analysis or reporting needs!