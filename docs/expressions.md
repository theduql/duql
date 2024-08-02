# DUQL Expression

Expressions in DUQL are powerful constructs used to perform calculations, comparisons, and transformations on data. They are fundamental to many DUQL operations, including filtering, generating new columns, and defining conditions.

## Syntax

DUQL expressions can take one of four forms:

1. Inline Expression
2. SQL Statement
3. Case Statement
4. Pipeline

### Inline Expression

```
<boolean_expression>
```

### SQL Statement

```yaml
sql: <sql_statement>
```

### Case Statement

```yaml
case:
  - <condition_1>: <result_1>
  - <condition_2>: <result_2>
  # ...
  - true: <default_result>
```

### Pipeline

```
(<initial_value> | <function_1> | <function_2> | ... | <function_n>)
```

## Types of Expressions

### Inline Expressions

Inline expressions are boolean expressions written directly in DUQL syntax. They can include column references, literals, operators, and function calls.

Examples:
```yaml
price * quantity > 1000
date.year(order_date) == 2023
lower(email) ~= "^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,}$"
array_contains(tags, 'urgent') && status != 'completed'
```

### SQL Statements

SQL statements allow you to use raw SQL within your DUQL queries. This is useful for complex operations or database-specific functions.

Examples:
```yaml
sql: SELECT AVG(price) FROM products WHERE category = 'Electronics'

sql: |
  WITH ranked_sales AS (
    SELECT product_id, sales_amount,
           ROW_NUMBER() OVER (PARTITION BY category ORDER BY sales_amount DESC) AS rank
    FROM sales
  )
  SELECT product_id, sales_amount
  FROM ranked_sales
  WHERE rank <= 3
```

### Case Statements

Case statements provide a way to express conditional logic. They consist of a series of condition-result pairs. The first condition that evaluates to true determines the result.

Examples:
```yaml
case:
  - age < 13: "Child"
  - age < 20: "Teenager"
  - age < 65: "Adult"
  - true: "Senior"

case:
  - order_total > 1000 && is_repeat_customer: "VIP"
  - order_total > 1000: "Big Spender"
  - is_repeat_customer: "Loyal Customer"
  - true: "New Customer"
```

### Pipelines

Pipelines are a powerful feature in DUQL that allow you to chain multiple operations or functions together. They provide a clear, left-to-right reading order for complex transformations.

Syntax:
```
(<initial_value> | <function_1> | <function_2> | ... | <function_n>)
```

Examples:
```yaml
(last_name | text.lower | text.starts_with("a"))
(age | math.pow 2)
(invoice_date | date.to_text "%d/%m/%Y")
(customer_data | json.extract 'preferences' | json.extract 'theme' | text.lower | text.equals 'dark')
```

How Pipelines Work:
1. The initial value (often a column name) is passed as input to the first function.
2. The result of each function is passed as the last argument to the next function.
3. The final result is the output of the last function in the pipeline.

For example, the pipeline:
```
(a | foo 3 | bar 'hello' 'world' | baz)
```
is equivalent to:
```
baz(bar('hello', 'world', foo(3, a)))
```

This makes it easier to read and write complex nested function calls, especially when dealing with data transformations that involve multiple steps.

## Complex Expression Examples

DUQL expressions can combine multiple operations and functions for sophisticated data manipulation:

```yaml
(revenue - cost) / revenue * 100 > 20 && units_sold > 100

date.day_of_week(order_date) in [6, 7] || date.is_holiday(order_date)

(user_data | json.extract 'preferences' | json.extract 'theme' | text.lower | text.equals 'dark')

date.add_months(subscription_start, 12) > current_date() && is_active

(description | text.trim | text.split '\s+' | array.length | math.between 50 100)
```

## Best Practices

1. 🧠 Keep expressions readable by breaking complex logic into smaller parts.
2. 🔢 Use parentheses to clarify the order of operations in complex calculations.
3. 🧪 Test expressions with sample data to ensure they produce expected results.
4. 📊 Consider performance implications of complex expressions, especially in large datasets.
5. 🚀 Leverage built-in functions to simplify common operations.
6. 📝 Use case statements for clear, multi-condition logic.
7. 🔍 When using SQL statements, ensure they are compatible with your target database.
8. 🔗 Use pipelines to make complex transformations more readable and maintainable.

## Related Functions

- [`filter`](filter.md): Often uses expressions to define row selection criteria.
- [`generate`](generate.md): Uses expressions to create or modify columns.
- [`group`](group.md): Can use expressions in summarizations and having clauses.
- [`sort`](sort.md): May use expressions to define custom sorting logic.

---

> 💡 **Tip:** Pipelines in DUQL expressions are a powerful tool for creating clear and maintainable data transformations. They allow you to break down complex operations into a series of simple steps, making your queries easier to read and modify. Experiment with pipelines to streamline your data processing logic!