# Windows

## DUQL Window Function

The window function in DUQL is a powerful tool for performing calculations across a set of rows that are related to the current row. It allows for complex analytics and comparisons within specified ranges or groups of data.

Window applies a step pipeline to segments of rows, producing one output value for every input value.

### Syntax

```yaml
window:
  <window_type>: <value>
  <steps>
```

### Parameters

For each row, the segment over which the pipeline is applied is determined by one of:

| Parameter   | Type              | Description                                                      |
| ----------- | ----------------- | ---------------------------------------------------------------- |
| `rows`      | string or integer | Specifies a range of rows relative to the current row position.  |
| `range`     | string            | Specifies a range of values relative to the current row value.   |
| `expanding` | boolean           | When true, creates a cumulative window (alias for `rows: ..0`).  |
| `rolling`   | integer           | Specifies a rolling window of n rows, including the current row. |

{% hint style="info" %}
`0` references the current row.
{% endhint %}

{% hint style="success" %}
For ease of use, there are two flags that override `rows` or `range`:

* `expanding:true` is an alias for `rows:..0`. A sum using this window is also known as “cumulative sum”.
* `rolling:n` is an alias for `rows:(-n+1)..0`, where `n` is an integer. This will include `n` last values, including current row. An average using this window is also knows as a Simple Moving Average.
{% endhint %}

{% hint style="warning" %}
The bounds of the range are inclusive. If a bound is omitted, the segment will extend until the edge of the table or group.
{% endhint %}

### Window Types

#### Row Range

```yaml
window:
  rows: <start>..<end>
```

* Specifies a range of rows relative to the current row position.
* Examples:
  * `"0..2"`: Current row and next two rows
  * `"-2..0"`: Previous two rows and current row
  * `"..0"`: All preceding rows and current row
  * `".."`: All rows in the partition

**Gotcha:** Large row ranges can impact performance on large datasets.

#### Value Range

```yaml
window:
  range: <start>..<end>
```

* Specifies a range of values relative to the current row value.
* Example: `"-1000..1000"`: All rows within 1000 units of the current row's value

**Gotcha:** Range is based on the values in the sorted column, not row numbers.

#### Expanding Window

```yaml
window:
  expanding: true
```

* Creates a cumulative window (alias for `rows: ..0`).
* A sum using this window is also known as “cumulative sum”.
* Useful for running totals or cumulative aggregations.

**Gotcha:** Can lead to performance issues on very large datasets.

#### Rolling Window

```yaml
window:
  rolling: <n>
```

* Specifies a rolling window of n rows, including the current row.
* `rolling:n` is an alias for `rows:(-n+1)..0`, where `n` is an integer.  This will include `n` last values, including current row.&#x20;
* Example: `7` for a 7-day rolling window.
* An average using this window is also knows as a Simple Moving Average.

**Gotcha:** Ensure the window size is appropriate for your data and analysis needs.

### Examples

#### Moving Average

```yaml
window:
  rows: "0..2"
  steps:
  - sort: [date, -amount]
  - select: [moving_average: average amount]
```

This example calculates a moving average over the current row and the next two rows, sorted by date and amount descending.

#### Cumulative Sum

```yaml
window:
  expanding: true
  steps:
  - generate:
      cumulative_sum: sum sales
```

This creates a running total of sales using an expanding window.

#### Weekly Average

```yaml
window:
  rolling: 7
  steps:
  - generate:
      weekly_average: average daily_visitors
```

This calculates a 7-day rolling average of daily visitors.

#### Centered Median

```yaml
window:
  rows: "-2..2"
  steps:
  - sort: date
  - select:
      centered_median: median temperature
```

This computes a centered median temperature using two rows before and after the current row.

#### Price Ranking

```yaml
window:
  range: "-1000..1000"
  steps:
  - sort: price
  - select:
     price_rank: rank this
```

This ranks prices within a range of 1000 units above and below the current price.

#### Percent of Total

```yaml
window:
  rows: ".."
  steps:
  - generate:
      percent_of_total: amount / sum(amount) * 100
```

This calculates each row's percentage of the total amount across all rows.

#### Running Total and Rank

```yaml
window:
  rows: "..0"
  steps:
  - sort: date
  - generate:
      running_total: sum amount
      rank_by_amount: rank -amount
```

This example computes both a running total and a rank based on amount in descending order.

#### Volatility Calculation

```yaml
window:
  rolling: 30
  steps:
  - sort: date
  - generate:
      volatility: stddev price / average price
```

This calculates price volatility over a 30-day rolling window.

#### Rate of Change

```yaml
window:
  rows: "-1..1"
  steps:
  - sort: timestamp
  - generate:
      rate_of_change: (lead(value) - lag(value)) / (2 * interval)
```

This computes the rate of change using the previous and next values.

#### Cumulative Conversion Rate

```yaml
window:
  expanding: true
  steps:
  - sort: date
  - generate:
      cumulative_conversion_rate: sum(conversions) / sum(visitors) * 100
```

This calculates a cumulative conversion rate over time.

### Best Practices

1. 🎯 Choose the appropriate window type for your analysis needs.
2. 🔢 Be mindful of window sizes, especially on large datasets.
3. 📊 Use sorting within window functions to ensure consistent results.
4. 🚀 Consider performance implications, especially with expanding windows or large ranges.
5. 🧪 Test window functions with sample data to verify results.
6. 📈 Combine window functions with other DUQL operations for complex analyses.

### Related Functions

* [`sort`](../basic/sort.md): Often used within window functions to order data.
* [`generate`](../intermediate/generate.md): Used to create new columns based on window calculations.
* [`select`](../basic/select.md): Can be used to choose specific window function results.

### Limitations and Considerations

* Window functions can be computationally expensive, especially on large datasets.
* Not all database systems support all types of window functions. Check your target database's capabilities.
* Complex window functions may require careful optimization for performance.

***

> 💡 **Tip:** Window functions are powerful tools for time-series analysis, rankings, and running calculations. They allow you to perform complex analytics without the need for self-joins or subqueries. Experiment with different window types and combinations to unlock deeper insights from your data!
