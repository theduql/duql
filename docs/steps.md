# DUQL Steps

The `steps` component in DUQL represents a series of data transformation operations applied sequentially to your dataset. It's the heart of DUQL's query structure, providing a flexible and intuitive way to express complex data manipulations as a pipeline of operations.

## Syntax

```yaml
steps:
- <step_1>
- <step_2>
- <step_3>
# ... more steps as needed
```

## Behavior

- Each step in the pipeline takes the output of the previous step as its input.
- Steps are executed in the order they are listed.
- The pipeline is flexible, allowing you to add, remove, or reorder steps easily.
- You can use any DUQL function as a step in the pipeline.

## Supported Step Types

The `steps` pipeline can include any combination of the following operations:

- `filter`: Select rows based on conditions
- `generate`: Create new columns or modify existing ones
- `group`: Aggregate data
- `join`: Combine data from multiple sources
- `select`: Choose or compute columns
- `sort`: Order results
- `take`: Limit the number of rows
- `window`: Perform window functions
- `append`: Combine datasets by adding rows
- `remove`: Exclude specific rows or subsets
- `intersect`: Find common rows between datasets
- `distinct`: Remove duplicate rows
- `union`: Combine datasets, removing duplicates
- `except`: Find rows in one dataset but not in another
- `loop`: Perform iterative processing

## Flexibility and Power

The `steps` pipeline in DUQL offers several advantages:

1. **Modularity**: Each step performs a specific operation, making queries easy to understand and maintain.
2. **Reusability**: You can easily copy, paste, and modify steps between queries.
3. **Readability**: The sequential nature of steps makes the data transformation process clear and logical.
4. **Incremental Development**: You can build complex queries step by step, testing and refining as you go.
5. **Easy Debugging**: You can comment out or remove steps to isolate issues in your query.
6. **Flexibility**: Steps can be easily reordered or modified to change the query logic.

## Examples

### Basic Pipeline

```yaml
steps:
- filter: order_date >= @2023-01-01
- generate:
    total_amount: price * quantity
- group:
    by: product_id
    summarize:
      total_sales: sum(total_amount)
- sort: -total_sales
- take: 10
```

### Complex Data Transformation

```yaml
steps:
- filter: status == 'active'
- join:
    dataset: customer_details
    where: orders.customer_id == customer_details.id
- generate:
    age: datediff(years, birth_date, current_date())
    customer_segment:
      case:
      - age < 25: "Young Adult"
      - age < 40: "Adult"
      - age < 60: "Middle Age"
      - true: "Senior"
- group:
    by: [customer_segment, product_category]
    summarize:
      total_revenue: sum(price * quantity)
      average_order_value: avg(price * quantity)
- sort: [customer_segment, -total_revenue]
- generate:
    revenue_rank:
      sql'ROW_NUMBER() OVER (PARTITION BY customer_segment ORDER BY total_revenue DESC)'
- filter: revenue_rank <= 5
```

## Best Practices

1. 🔍 Start with a `filter` step to reduce data volume early in the pipeline.
2. 🧮 Use `generate` steps to create new columns needed for analysis.
3. 🔗 Place `join` steps early in the pipeline, but after initial filtering.
4. 📊 Use `group` steps for aggregations, followed by `generate` for post-aggregation calculations.
5. 🏷️ Use `select` towards the end of the pipeline to choose final output columns.
6. 🔢 End with `sort` and `take` steps to order and limit final results.
7. 📝 Comment your steps for complex queries to explain the transformation logic.

## Real-World Use Case

Here's an example of a DUQL query that uses a complex `steps` pipeline to perform a cohort analysis:

```yaml
dataset: user_actions

steps:
- filter: action_date >= @2023-01-01
- generate:
    cohort_date: date_trunc('month', first_action_date)
    months_since_first_action: datediff(months, cohort_date, action_date)
- group:
    by: [cohort_date, months_since_first_action]
    summarize:
      user_count: count_distinct(user_id)
      total_actions: count(action_id)
- join:
    dataset: 
      steps:
      - dataset: user_actions
      - group:
          by: cohort_date
          summarize:
            initial_users: count_distinct(user_id)
    where: ==cohort_date
- generate:
    retention_rate: user_count / initial_users
- sort: [cohort_date, months_since_first_action]
- select:
    Cohort: cohort_date
    "Months Since First Action": months_since_first_action
    "User Count": user_count
    "Retention Rate": retention_rate
- filter: months_since_first_action <= 12

into: user_cohort_analysis
```

This query demonstrates:
1. Filtering recent data
2. Generating cohort and time-based columns
3. Aggregating user actions
4. Joining with a subquery to get initial user counts
5. Calculating retention rates
6. Sorting results
7. Selecting and renaming final columns
8. Filtering to show only the first year of data for each cohort

The `steps` pipeline allows for this complex analysis to be expressed clearly and logically, making it easy to understand and modify the cohort analysis process.

---

> 💡 **Tip:** Think of the `steps` pipeline as a data assembly line. Each step performs a specific task, gradually shaping your data into the final form you need. Don't hesitate to experiment with the order and composition of steps to achieve your desired result!