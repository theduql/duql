# DUQL Expressions

Expressions in DUQL are powerful constructs used to perform calculations, comparisons, and transformations on data. They are fundamental to many DUQL operations, including filtering, generating new columns, and defining join conditions.

## Syntax

Expressions in DUQL can range from simple column references to complex combinations of functions, operators, and conditional logic. Here's a general syntax overview:

```
<column_name>                           # Simple column reference
<function>(<arguments>)                 # Function call
<expression> <operator> <expression>    # Binary operation
<unary_operator><expression>            # Unary operation
case: [<condition>: <result>, ...]      # Conditional expression
```

## Types of Expressions

### Column References

Simply use the column name to reference its values.

Example:
```yaml
filter: age > 18
```

### Literals

- Strings: Enclosed in double quotes `"example"`
- Numbers: `42`, `3.14`
- Booleans: `true`, `false`
- Dates: `@2023-05-01`
- Null: `null`

### Operators

#### Arithmetic Operators
- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Modulus: `%`
- Exponentiation: `**`

#### Comparison Operators
- Equal: `==`
- Not Equal: `!=`
- Greater Than: `>`
- Less Than: `<`
- Greater Than or Equal: `>=`
- Less Than or Equal: `<=`

#### Logical Operators
- AND: `&&`
- OR: `||`
- NOT: `!`

#### String Operators
- Concatenation: `+`
- Pattern Matching: `~=` (regular expression match)

### Functions

DUQL supports a wide range of functions for various operations:

#### String Functions
- `lower(text)`: Convert to lowercase
- `upper(text)`: Convert to uppercase
- `trim(text)`: Remove leading and trailing whitespace
- `length(text)`: Get string length
- `contains(substring, text)`: Check if text contains substring

#### Numeric Functions
- `abs(number)`: Absolute value
- `round(number, decimals)`: Round to specified decimal places
- `ceil(number)`: Round up to the nearest integer
- `floor(number)`: Round down to the nearest integer

#### Date Functions
- `year(date)`: Extract year from date
- `month(date)`: Extract month from date
- `day(date)`: Extract day from date
- `datediff(unit, start_date, end_date)`: Calculate difference between dates

#### Aggregate Functions (used in `group` operations)
- `sum(expression)`: Calculate sum
- `avg(expression)`: Calculate average
- `count(expression)`: Count non-null values
- `min(expression)`: Find minimum value
- `max(expression)`: Find maximum value

### Conditional Expressions

Use the `case` construct for conditional logic:

```yaml
generate:
  status:
    case:
      - age < 18: "Minor"
      - age < 65: "Adult"
      - true: "Senior"
```

### SQL Strings

Use the `sql''` syntax to include raw SQL in your DUQL expressions. This is useful for database-specific functions or complex operations not yet supported by DUQL.

Example:
```yaml
filter: sql'regexp_contains(title, ''([a-z0-9]*-){{2,}}'')'

## Examples

### Basic Arithmetic and Comparison
```yaml
filter: (price * quantity) > 1000 && discount < 0.2
```

### String Manipulation
```yaml
generate:
  full_name: f"{lower(first_name)} {upper(last_name)}"
```

### Date Calculations
```yaml
filter: datediff(days, order_date, current_date()) <= 30
```

### Complex Conditional Logic
```yaml
generate:
  shipping_cost:
    case:
      - weight < 1: 5.99
      - weight < 5: 9.99
      - weight < 10: 14.99
      - true: weight * 2
```

### Using Functions in Aggregations
```yaml
group:
  by: category
  summarize:
    total_revenue: sum(price * quantity)
    avg_order_value: avg(price * quantity)
```

## Best Practices

1. 🧠 Keep expressions readable by breaking complex logic into smaller parts.
2. 🔢 Use parentheses to clarify the order of operations in complex calculations.
3. 🧪 Test expressions with sample data to ensure they produce expected results.
4. 🏷️ Use meaningful names for generated columns to improve query understandability.
5. 🚀 Leverage built-in functions to simplify common operations.
6. 📊 Consider performance implications of complex expressions, especially in large datasets.

## Real-World Use Case

Here's an example of a DUQL query using various expressions to analyze product performance:

```yaml
dataset: sales

steps:
  - generate:
      sale_date: date.to_text "%Y-%m-%d" timestamp
      total_amount: price * quantity
      discount_amount: price * quantity * discount_rate
      net_amount: total_amount - discount_amount
  - filter: sale_date >= @2023-01-01
  - group:
      by: [product_id, category]
      summarize:
        total_sales: sum(net_amount)
        units_sold: sum(quantity)
        avg_price: avg(price)
  - generate:
      performance_score:
        case:
          - total_sales > 10000 && units_sold > 100: "High"
          - total_sales > 5000 || units_sold > 50: "Medium"
          - true: "Low"
      price_category:
        case:
          - avg_price < 50: "Budget"
          - avg_price < 200: "Mid-range"
          - true: "Premium"
  - sort: -total_sales

into: product_performance_analysis
```

This query demonstrates the use of various expressions to:
1. Generate new columns with calculated values
2. Filter data based on a date condition
3. Perform grouping and aggregation
4. Create categorical columns based on complex conditions
5. Sort the results

The resulting `product_performance_analysis` dataset provides valuable insights into product performance, considering sales amounts, units sold, and pricing strategies.

---

> 💡 **Tip:** Expressions are the building blocks of powerful DUQL queries. Master them to unlock the full potential of your data analysis capabilities!