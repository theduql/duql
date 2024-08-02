# DUQL Generate

The `generate` function in DUQL is used to create new columns or modify existing ones based on expressions, calculations, and conditional logic. It's a powerful tool for data transformation and enrichment within your query pipeline.

## Syntax

```yaml
generate:
  <new_column_name>: <expression>
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `generate` | object | Yes | Key-value pairs where keys are new column names and values are expressions |

## Behavior

- Creates new columns or modifies existing ones based on the provided expressions.
- Can reference other columns, use functions, and implement conditional logic.
- Executes in the order specified within the `generate` block.

## Examples

### Basic Column Generation

```yaml
generate:
  full_name: f"{first_name} {last_name}"
  age: datediff(years, birth_date, current_date())
```

### Boolean Flags

```yaml
generate:
  is_northern: latitude > 0
  should_have_shipped_today: order_time < @08:30
```

### Conditional Logic

```yaml
generate:
  discount:
    case:
      - total_purchases > 1000: 0.15
      - total_purchases > 500: 0.10
      - total_purchases > 100: 0.05
      - true: 0
```

### Complex Calculations

```yaml
generate:
  bmi: weight / (height / 100) ^ 2
  health_status:
    case:
      - bmi < 18.5: "Underweight"
      - bmi < 25: "Normal"
      - bmi < 30: "Overweight"
      - true: "Obese"
```

### Using SQL Functions

```yaml
generate:
  clean_text: 
    sql'REGEXP_REPLACE(LOWER(text), ''[^a-z0-9 ]'', '''')'
  word_count:
    sql'ARRAY_LENGTH(SPLIT(clean_text, '' ''))'
```

### Date-based Calculations

```yaml
generate:
  is_weekend: day_of_week in [6, 7]
  season:
    case:
      - month in [12, 1, 2]: "Winter"
      - month in [3, 4, 5]: "Spring"
      - month in [6, 7, 8]: "Summer"
      - true: "Fall"
```

### Financial Calculations

```yaml
generate:
  revenue: price * quantity
  profit_margin: (revenue - cost) / revenue * 100
  performance:
    case:
      - profit_margin > 30: "Excellent"
      - profit_margin > 20: "Good"
      - profit_margin > 10: "Average"
      - true: "Poor"
```

### Geospatial Calculations

```yaml
generate:
  distance: 
    sql'ST_Distance(ST_Point(long, lat), ST_Point(-74.006, 40.7128))'
  is_nearby: distance < 10
```

## Best Practices

1. 🏷️ Use clear and descriptive names for generated columns.
2. 🧮 Break down complex calculations into multiple steps for readability.
3. 🔍 Validate your generated columns, especially when using complex expressions.
4. 🚀 Consider performance implications when generating columns with heavy calculations.
5. 📊 Use `case` statements for complex conditional logic.
6. 🔄 Remember that columns generated earlier can be used in subsequent generations within the same `generate` block.

## Real-World Use Case

Here's an example of a DUQL query that uses `generate` to enrich and analyze sales data:

```yaml
dataset: sales

steps:
  - join:
      dataset: products
      where: sales.product_id == products.id
  - join:
      dataset: customers
      where: sales.customer_id == customers.id
  - generate:
      revenue: price * quantity
      cost: wholesale_price * quantity
      profit: revenue - cost
      profit_margin: (profit / revenue) * 100
      days_since_last_purchase: 
        sql'DATEDIFF(sale_date, customers.last_purchase_date)'
      customer_segment:
        case:
          - days_since_last_purchase <= 30: "Recent"
          - days_since_last_purchase <= 90: "Active"
          - days_since_last_purchase <= 365: "Lapsed"
          - true: "Inactive"
      is_high_value:
        case:
          - revenue > 1000: true
          - profit_margin > 50 && quantity > 10: true
          - true: false
      season:
        case:
          - month(sale_date) in [12, 1, 2]: "Winter"
          - month(sale_date) in [3, 4, 5]: "Spring"
          - month(sale_date) in [6, 7, 8]: "Summer"
          - true: "Fall"
  - group:
      by: [product_id, customer_segment, season]
      summarize:
        total_revenue: sum(revenue)
        total_profit: sum(profit)
        avg_profit_margin: avg(profit_margin)
        high_value_sales: sum(is_high_value)
  - generate:
      performance_score: 
        (total_revenue * 0.4) + (total_profit * 0.4) + (avg_profit_margin * 0.2)
      rank:
        sql'ROW_NUMBER() OVER (PARTITION BY customer_segment, season ORDER BY performance_score DESC)'

into: product_performance_analysis
```

This query demonstrates:
1. Joining sales data with product and customer information
2. Generating financial metrics (revenue, cost, profit, profit margin)
3. Creating customer segments based on purchase recency
4. Identifying high-value sales
5. Determining the season of each sale
6. Aggregating data by product, customer segment, and season
7. Calculating a performance score and ranking products within each segment and season

The resulting `product_performance_analysis` provides a comprehensive view of product performance across different customer segments and seasons, showcasing the power and flexibility of the `generate` function in DUQL.

---

> 💡 **Tip:** The `generate` function is your Swiss Army knife for data transformation in DUQL. Use it to create new insights, segment your data, and prepare it for deeper analysis. Don't hesitate to combine multiple `generate` steps in your pipeline for complex transformations!