# Join

The `join` function in DUQL is used to combine rows from two or more tables based on a related column between them. It supports different types of joins and allows for complex join conditions.

## Syntax

```yaml
join:
  dataset: <data_source>
  where: <join_condition>
  retain: <join_type>
```

## Parameters

| Parameter | Type   | Required | Default   | Description                                            |
| --------- | ------ | -------- | --------- | ------------------------------------------------------ |
| `dataset` | object | Yes      | -         | The data source to join with (ref: dataset.s.duql.yml) |
| `where`   | string | Yes      | -         | The join condition                                     |
| `retain`  | string | No       | `"inner"` | The type of join to perform                            |

## Behavior

* Combines rows from the current dataset with rows from the specified dataset based on the join condition.
* The `dataset` parameter can include its own `steps` for preprocessing before the join.
* The `where` condition supports a `==` shorthand for equality joins on matching column names.
* The `retain` parameter specifies the join type, defaulting to an inner join.

## Join Types

| Value     | Description                                                                |
| --------- | -------------------------------------------------------------------------- |
| `"inner"` | Returns only the matched rows (default)                                    |
| `"left"`  | Returns all rows from the left table and matched rows from the right table |
| `"right"` | Returns all rows from the right table and matched rows from the left table |
| `"full"`  | Returns all rows when there is a match in either left or right table       |

## Examples

### Basic Join

```yaml
join:
  dataset: customers
  steps: 
  where: orders.customer_id == customers.id
```

### Left Join

```yaml
join:
  dataset: products
  where: orders.product_id == products.id
  retain: left
```

### Join with CSV File

```yaml
join:
  dataset: myorg/employee_data.csv
  type: csv
  where: departments.department_id == employee_data.department_id
  retain: inner
```

### Join with Subquery and Full Outer Join

```yaml
join:
  dataset: recent_inventory
  steps:
  - dataset: inventory
  - filter: last_updated > @2023-01-01
  where: products.product_id == recent_inventory.product_id
  retain: full
```

### Complex Join with Aggregation

```yaml
join:
  dataset: customer_views
  where: customers.id == customer_views.customer_id
```

### Join with SQL Query

```yaml
join:
  dataset: sql'SELECT * FROM myexample'
  where: ==id
  retain: full
```

### Join with External Data Source

```yaml
join:
  dataset: hdfs://cluster/user_profiles/*.parquet
  type: parquet
  where: users.id == user_profiles.user_id
  retain: left
```

## Best Practices

1. 🎯 Always specify the join condition explicitly for clarity.
2. 🤔 Choose the appropriate join type using the `retain` parameter based on your data requirements.
3. 🚀 Utilize the `steps` within the `dataset` parameter for preprocessing before joining.
4. 🔍 Use the `==` shorthand for simple equality joins when column names match.
5. 📊 Be mindful of data volume when performing joins, especially with large datasets or complex subqueries.
6. 🏷️ Consider using aliases or clear naming conventions when joining multiple tables to avoid ambiguity.

## Real-World Use Case

Here's an example of a DUQL query that uses multiple joins to create a comprehensive sales report:

```yaml
dataset: orders

steps:
- filter: order_date >= @2023-01-01
- join:
    dataset: customers
    where: orders.customer_id == customers.id
- join:
    dataset: products
    where: orders.product_id == products.id
- join:
    dataset: inventory
    steps:
    - group:
        by: product_id
        summarize:
          current_stock: sum(quantity)
    where: products.id == inventory.product_id
- generate:
    total_amount: orders.quantity * products.price
- group:
    by: [products.category, products.id, products.name]
    summarize:
      total_sales: sum(total_amount)
      units_sold: sum(orders.quantity)
      unique_customers: count_distinct(customer_id)
- join:
    dataset: | 
      sql' SELECT category, AVG(total_sales) as avg_category_sales
      FROM (
        SELECT products.category, SUM(orders.quantity * products.price) as total_sales
        FROM orders
        JOIN products ON orders.product_id = products.id
        WHERE orders.order_date >= '2023-01-01'
        GROUP BY products.category, products.id
      ) category_sales
      GROUP BY category'
    where: ==category
- generate:
    performance_ratio: total_sales / avg_category_sales
    stock_status:
      case:
      - current_stock == 0: "Out of Stock"
      - current_stock < 10: "Low Stock"
      - true: "In Stock"
- sort: [category, -total_sales]

into: product_sales_report
```

This query demonstrates:

1. Joining order data with customer, product, and inventory information
2. Using subqueries within joins for aggregations
3. Incorporating SQL queries for complex calculations
4. Generating performance metrics and stock status
5. Sorting the final results

The `join` operations allow us to bring together data from multiple sources, creating a rich dataset for analysis that includes sales performance, customer behavior, inventory status, and category-level benchmarks.

***

> 💡 **Tip:** The `join` function in DUQL is highly flexible, allowing you to combine data from various sources and formats. Leverage its ability to include `steps` within the `dataset` parameter to preprocess your data before joining, enabling more sophisticated data integration in your analyses!
