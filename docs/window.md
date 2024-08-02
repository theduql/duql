Certainly! I'll create comprehensive documentation for the DUQL Window Function based on the provided schema. Here's the markdown documentation:

# DUQL Window Function

The window function in DUQL is a powerful tool for performing calculations across a set of rows that are related to the current row. It allows for complex analytics and comparisons within specified ranges or groups of data.

## Syntax

```yaml
window:
  <window_type>: <value>
  <additional_steps>
```

## Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `rows` | string or integer | Specifies a range of rows relative to the current row position. |
| `range` | string | Specifies a range of values relative to the current row value. |
| `expanding` | boolean | When true, creates a cumulative window (alias for `rows: ..0`). |
| `rolling` | integer | Specifies a rolling window of n rows, including the current row. |

Additional steps can be included within the window operation, such as `sort`, `select`, or `generate`.

## Window Types

### Row Range

```yaml
window:
  rows: <start>..<end>
```

- Specifies a range of rows relative to the current row position.
- Examples:
  - `"0..2"`: Current row and next two rows
  - `"-2..0"`: Previous two rows and current row
  - `"..0"`: All preceding rows and current row
  - `".."`: All rows in the partition

**Gotcha:** Large row ranges can impact performance on large datasets.

### Value Range

```yaml
window:
  range: <start>..<end>
```

- Specifies a range of values relative to the current row value.
- Example: `"-1000..1000"`: All rows within 1000 units of the current row's value

**Gotcha:** Range is based on the values in the sorted column, not row numbers.

### Expanding Window

```yaml
window:
  expanding: true
```

- Creates a cumulative window (alias for `rows: ..0`).
- Useful for running totals or cumulative aggregations.

**Gotcha:** Can lead to performance issues on very large datasets.

### Rolling Window

```yaml
window:
  rolling: <n>
```

- Specifies a rolling window of n rows, including the current row.
- Example: `7` for a 7-day rolling window.

**Gotcha:** Ensure the window size is appropriate for your data and analysis needs.

## Examples

### Moving Average

```yaml
window:
  rows: "0..2"
  sort: [date, -amount]
  select:
    moving_average: average amount
```

This example calculates a moving average over the current row and the next two rows, sorted by date and amount descending.

### Cumulative Sum

```yaml
window:
  expanding: true
  generate:
    cumulative_sum: sum sales
```

This creates a running total of sales using an expanding window.

### Weekly Average

```yaml
window:
  rolling: 7
  generate:
    weekly_average: average daily_visitors
```

This calculates a 7-day rolling average of daily visitors.

### Centered Median

```yaml
window:
  rows: "-2..2"
  sort: date
  select:
    centered_median: median temperature
```

This computes a centered median temperature using two rows before and after the current row.

### Price Ranking

```yaml
window:
  range: "-1000..1000"
  sort: price
  select:
    price_rank: rank this
```

This ranks prices within a range of 1000 units above and below the current price.

### Percent of Total

```yaml
window:
  rows: ".."
  generate:
    percent_of_total: amount / sum(amount) * 100
```

This calculates each row's percentage of the total amount across all rows.

### Running Total and Rank

```yaml
window:
  rows: "..0"
  sort: date
  generate:
    running_total: sum amount
    rank_by_amount: rank -amount
```

This example computes both a running total and a rank based on amount in descending order.

### Volatility Calculation

```yaml
window:
  rolling: 30
  sort: date
  generate:
    volatility: stddev price / average price
```

This calculates price volatility over a 30-day rolling window.

### Rate of Change

```yaml
window:
  rows: "-1..1"
  sort: timestamp
  generate:
    rate_of_change: (lead(value) - lag(value)) / (2 * interval)
```

This computes the rate of change using the previous and next values.

### Cumulative Conversion Rate

```yaml
window:
  expanding: true
  sort: date
  generate:
    cumulative_conversion_rate: sum(conversions) / sum(visitors) * 100
```

This calculates a cumulative conversion rate over time.

## Best Practices

1. 🎯 Choose the appropriate window type for your analysis needs.
2. 🔢 Be mindful of window sizes, especially on large datasets.
3. 📊 Use sorting within window functions to ensure consistent results.
4. 🚀 Consider performance implications, especially with expanding windows or large ranges.
5. 🧪 Test window functions with sample data to verify results.
6. 📈 Combine window functions with other DUQL operations for complex analyses.

## Related Functions

- [`sort`](sort.md): Often used within window functions to order data.
- [`generate`](generate.md): Used to create new columns based on window calculations.
- [`select`](select.md): Can be used to choose specific window function results.

## Limitations and Considerations

- Window functions can be computationally expensive, especially on large datasets.
- Not all database systems support all types of window functions. Check your target database's capabilities.
- Complex window functions may require careful optimization for performance.

---

> 💡 **Tip:** Window functions are powerful tools for time-series analysis, rankings, and running calculations. They allow you to perform complex analytics without the need for self-joins or subqueries. Experiment with different window types and combinations to unlock deeper insights from your data!