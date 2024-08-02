# Getting Started with DUQL

DUQL is a powerful, YAML-based query language designed for data transformation and analysis. It's a user-friendly version of PRQL with additional usability improvements. This guide will help you understand the basics of DUQL and get you started with writing your own queries.

## 1. Basic Query Structure

A DUQL query typically consists of the following components:

```yaml
settings:  # Optional
  version: '0.0.1'
  target: sql.postgres

dataset: your_table_name

steps:
  - step1
  - step2
  # ... more steps as needed

into: result_name  # Optional
```

- `settings`: Optional. Specifies the DUQL version and target database.
- `dataset`: Required. Specifies the main data source for your query.
- `steps`: Optional, but typically used. Defines the sequence of operations to perform on your data.
- `into`: Optional. Names the output of your query for further use.

## 2. Simple Operators and Expressions

DUQL supports various operators for constructing expressions:

- Arithmetic: `+`, `-`, `*`, `/`, `%` (modulo), `**` (exponentiation)
- Comparison: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Logical: `&&` (AND), `||` (OR), `!` (NOT)
- String: `+` (concatenation), `~=` (regex match)

Example:
```yaml
steps:
  - filter: age > 18 && (city == "New York" || city == "Los Angeles")
```

## 3. Basic Functions

DUQL includes many built-in functions for data manipulation:

- String functions: `lower()`, `upper()`, `trim()`, `length()`
- Date functions: `year()`, `month()`, `day()`, `datediff()`
- Numeric functions: `round()`, `abs()`, `sum()`, `avg()`

Example:
```yaml
steps:
- generate:
    full_name: f"{lower(first_name)} {upper(last_name)}"
    age: datediff(years, birth_date, current_date())
```

## 4. Common Transformation Steps

Here are some frequently used transformation steps:

### Filter
Selects rows based on a condition:
```yaml
- filter: price > 100 && category == "Electronics"
```

### Generate
Creates new columns or modifies existing ones:
```yaml
- generate:
    total_price: price * quantity
    discount: 
      case:
        - total_price > 1000: 0.1
        - total_price > 500: 0.05
        - true: 0
```

### Group
Aggregates data based on specified columns:
```yaml
- group:
    by: [category, year(order_date)]
    summarize:
      total_sales: sum(price * quantity)
      order_count: count(order_id)
```

### Sort
Orders the results:
```yaml
- sort: [-total_sales, category]
```

### Take
Limits the number of rows returned:
```yaml
- take: 10
```

## 5. Putting It All Together

Here's an example that combines multiple steps to analyze sales data:

```yaml
settings:
  version: '0.0.1'
  target: sql.postgres

dataset: sales

steps:
- filter: order_date >= @2023-01-01
- generate:
    total_amount: price * quantity
    order_year: year(order_date)
- group:
    by: [category, order_year]
    summarize:
      total_sales: sum(total_amount)
      order_count: count(order_id)
- filter: total_sales > 10000
- sort: [-order_year, -total_sales]
- take: 5

into: top_selling_categories
```

This query:
1. Filters for orders from 2023 onwards
2. Calculates the total amount for each order and extracts the year
3. Groups the data by category and year, summarizing total sales and order count
4. Filters for categories with over $10,000 in sales
5. Sorts the results by year (descending) and total sales (descending)
6. Takes the top 5 results
7. Stores the result in a dataset named `top_selling_categories`

## Next Steps

As you become more comfortable with DUQL, you can explore more advanced features such as:
- Joining multiple datasets
- Window functions for complex analytics
- Declaring reusable variables and functions
- Using subqueries and complex expressions

Remember, DUQL is designed to be intuitive and expressive. Don't hesitate to experiment with different combinations of steps and functions to achieve your data analysis goals!