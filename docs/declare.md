# DUQL Declare

The `declare` component in DUQL allows you to define reusable elements such as variables, functions, and subqueries. These declarations can be referenced throughout your DUQL query, promoting code reuse and improving query readability.

## Syntax

The `declare` section is typically placed at the beginning of a DUQL query and can contain multiple declarations:

```yaml
declare:
  <variable_name>: <value_or_expression>
  <function_name>: <function_definition>
  <subquery_name>:
    dataset: <data_source>
    steps:
      - <transformation_step_1>
      - <transformation_step_2>
      # ... more steps as needed
```

## Types of Declarations

### Simple Variables

You can declare simple variables with literal values or expressions.

Example:
```yaml
declare:
  tax_rate: 0.08
  current_year: 2023
  company_name: "Acme Corp"
```

### Functions

Functions can be declared using a simplified arrow syntax or a more detailed YAML structure.

Simple syntax:
```yaml
declare:
  calculate_total: price quantity -> price * quantity * (1 + tax_rate)
```

Detailed syntax:
```yaml
declare:
  calculate_total:
    parameters: [price, quantity]
    expression: price * quantity * (1 + tax_rate)
```

### Subqueries

Subqueries are declared as complete DUQL pipelines that can be reused in your main query.

Example:
```yaml
declare:
  active_customers:
    dataset: customers
    steps:
      - filter: last_purchase_date >= @2023-01-01
      - select: [customer_id, name, email]
```

## Examples

### Mixed Declarations

```yaml
declare:
  tax_rate: 0.08
  calculate_total: price quantity -> price * quantity * (1 + tax_rate)
  recent_orders:
    dataset: orders
    steps:
      - filter: order_date >= @2023-01-01
      - join:
          dataset: customers
          where: orders.customer_id == customers.id
```

### Using Declarations in a Query

```yaml
declare:
  high_value_threshold: 1000
  is_high_value: amount -> amount > high_value_threshold

dataset: recent_orders

steps:
  - generate:
      total_amount: calculate_total(price, quantity)
      is_high_value: is_high_value(total_amount)
  - filter: is_high_value
  - sort: -total_amount
```

## Best Practices

1. 📝 Use clear and descriptive names for your declarations to improve query readability.
2. 🔄 Leverage declarations to avoid repetition in your queries.
3. 🧩 Break down complex logic into smaller, reusable functions.
4. 🏗️ Use subquery declarations to create modular and maintainable query components.
5. 📊 Consider performance implications when using complex subqueries in declarations.
6. 🔍 Document your declarations, especially for complex functions or subqueries.

## Real-World Use Case

Here's an example of a DUQL query that makes extensive use of declarations:

```yaml
declare:
  tax_rate: 0.08
  shipping_threshold: 50
  calculate_total: price quantity -> price * quantity * (1 + tax_rate)
  apply_shipping: total -> 
    case:
    - total >= shipping_threshold: total
    - true: total + 10
  recent_customers:
    dataset: customers
    steps:
    - filter: last_purchase_date >= @2023-01-01
  product_categories:
    dataset: products
    steps:
    - select: [product_id, category]

dataset: orders

steps:
  - filter: order_date >= @2023-01-01
  - join:
      dataset: recent_customers
      where: orders.customer_id == recent_customers.customer_id
  - join:
      dataset: product_categories
      where: orders.product_id == product_categories.product_id
  - generate:
      subtotal: calculate_total(price, quantity)
      total_with_shipping: apply_shipping(subtotal)
  - group:
      by: [customer_id, category]
      summarize:
        total_spent: sum(total_with_shipping)
        num_orders: count(order_id)
  - sort: -total_spent
  - take: 100

into: top_customer_category_analysis
```

This query demonstrates:
1. Declaring constants (`tax_rate`, `shipping_threshold`)
2. Defining reusable functions (`calculate_total`, `apply_shipping`)
3. Creating subquery declarations (`recent_customers`, `product_categories`)
4. Using these declarations throughout the main query for calculations, filtering, and joins

The resulting `top_customer_category_analysis` dataset provides insights into the top-spending customers by product category, incorporating tax and shipping calculations.

---

> 💡 **Tip:** The `declare` section is a powerful tool for creating reusable and maintainable DUQL queries. Use it to define your business logic once and apply it consistently throughout your data analysis pipeline!

### Into

The `into` component specifies the destination for your query results. It's similar to declaring a variable, but it occurs at the end of a query pipeline. When you use `into`, you're essentially creating a named result set that can be referenced in subsequent queries or operations.

Key points about `into`:

1. It functions like a variable declaration that happens at the end of a query.
2. The result of all preceding steps in the query is stored in the named variable specified by `into`.
3. This named result can be used as a dataset in other DUQL queries within the same session or script.

Example:
```yaml
# ... previous query steps ...

into: monthly_sales_report
```

In this example, `monthly_sales_report` becomes a named dataset containing the results of the query. You can then use it in subsequent queries like this:

```yaml
dataset: monthly_sales_report
steps:
- filter: total_sales > 10000
  # ... more steps ...
```

The key difference between `into` and a variable declared in the `declare` section is the timing and context:

- Variables in `declare` are defined before the main query pipeline and can be used throughout the query.
- `into` creates a named result at the end of the query pipeline, making the final result available for future use.

Think of `into` as a way to save your query results for further analysis or as building blocks for more complex data operations. It's particularly useful when you want to break down a complex analysis into multiple, manageable DUQL queries.