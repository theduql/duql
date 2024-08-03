# Developer Observability

DUQL provides powerful capabilities for analyzing and visualizing developer observability data, including metrics, logs, and application performance monitoring (APM) data. This guide will walk you through various examples of how to use DUQL to gain insights from your observability data.

## 1. Analyzing Metrics

Let's start with some common metric analysis scenarios using the assumed schema: `metric_name, data_point, time, service, host, tag3, ...`

### 1.1 Service Health Overview

This query provides an overview of service health by calculating the average response time and error rate for each service over the last 24 hours.

```yaml
dataset: metrics

steps:
- filter: time >= current_timestamp() - interval '24 hours'
- group:
    by: service
    steps:
    - summarize:
        avg_response_time: 
          avg: 
            case:
            - metric_name == 'response_time': data_point
            - true: null
        error_rate:
          sum:
            case:
            - metric_name == 'error_count': data_point
            - true: 0
          / sum:
            case:
            - metric_name == 'request_count': data_point
            - true: 0
          * 100
- filter: avg_response_time is not null
- sort: -avg_response_time
- generate:
    health_status:
      case:
      - avg_response_time > 1000 || error_rate > 5: "Critical"
      - avg_response_time > 500 || error_rate > 2: "Warning"
      - true: "Healthy"

into: service_health_overview
```

This query:

1. Filters metrics from the last 24 hours
2. Groups data by service
3. Calculates average response time and error rate
4. Filters out services without response time data
5. Sorts by average response time (descending)
6. Generates a health status based on response time and error rate thresholds

### 1.2 Host Resource Utilization

This query analyzes CPU and memory utilization across hosts, identifying potential resource constraints.

```yaml
dataset: metrics

steps:
- filter: 
    metric_name in ['cpu_utilization', 'memory_utilization'] 
    && time >= current_timestamp() - interval '1 hour'
- group:
    by: [host, metric_name]
    steps:
    - summarize:
        avg_utilization: avg data_point
        max_utilization: max data_point
- generate:
    utilization_status:
      case:
      - max_utilization > 90: "Critical"
      - max_utilization > 80: "Warning"
      - true: "Normal"
- sort: [host, -max_utilization]

into: host_resource_utilization
```

This query:

1. Filters CPU and memory utilization metrics from the last hour
2. Groups data by host and metric name
3. Calculates average and maximum utilization
4. Generates a utilization status based on maximum utilization
5. Sorts results by host and maximum utilization (descending)

## 2. Log Analysis

While the given schema is for metrics, let's assume we have a `logs` table with columns: `timestamp, log_level, service, message, trace_id`.

### 2.1 Error Log Summary

This query summarizes error logs across services.

```yaml
dataset: logs

steps:
- filter: 
    log_level == 'ERROR'
    && 
    timestamp >= current_timestamp() - interval '24 hours'
- generate:
    error_type: regexp_extract(message, '^([A-Za-z]+Error)')
- group:
    by: [service, error_type]
    steps:
    - summarize: 
      error_count: count this
- sort: [service, -error_count]
- generate:
    error_percentage:
      sql: "error_count * 100.0 / SUM(error_count) OVER (PARTITION BY service)"
- filter: error_count > 10

into: error_log_summary
```

This query:

1. Filters error logs from the last 24 hours
2. Extracts error type from the message
3. Groups errors by service and error type
4. Calculates error count and percentage within each service
5. Filters out infrequent errors (count <= 10)

### 2.2 Log Pattern Mining

This example identifies common log patterns to help spot recurring issues.

```yaml
dataset: logs

steps:
- filter: timestamp >= current_timestamp() - interval '1 hour'
- generate:
    log_pattern: regexp_replace(message, '[0-9]+', 'N')
- group:
    by: [service, log_level, log_pattern]
    steps: 
    - summarize: 
        pattern_count: count this
        example_message: first(message)
- filter: pattern_count > 10
- sort: [service, -pattern_count]
- take: 100

into: common_log_patterns
```

This query:

1. Filters logs from the last hour
2. Generates a log pattern by replacing numbers with 'N'
3. Groups logs by service, log level, and pattern
4. Counts occurrences of each pattern and keeps an example message
5. Filters patterns occurring more than 10 times
6. Sorts by service and pattern count (descending)
7. Takes the top 100 patterns

## 3. Application Performance Monitoring (APM)

For APM examples, let's assume we have a `traces` table with columns: `trace_id, parent_span_id, span_id, service, operation, start_time, duration, status`.

### 3.1 Slow Transaction Analysis

This query identifies slow transactions and their component spans.

```yaml
dataset: traces

steps:
- filter: 
    parent_span_id is null  # Root spans only
    && start_time >= current_timestamp() - interval '1 hour'
    && duration > 1000  # Transactions taking more than 1 second
- join:
    dataset: traces as spans
    where: traces.trace_id == spans.trace_id
- group:
    by: [traces.trace_id, traces.service, traces.operation]
    steps:
    - summarize:
      total_duration: max(traces.duration)
      span_breakdown:
        sql: "STRING_AGG(CONCAT(spans.service, '.', spans.operation, ': ', spans.duration, 'ms'), ', ' ORDER BY spans.duration DESC)"
- sort: -total_duration
- take: 50

into: slow_transactions
```

This query:

1. Filters for root spans (transactions) from the last hour that took more than 1 second
2. Joins with all spans in the same trace
3. Groups by trace, service, and operation
4. Calculates total duration and creates a breakdown of component spans
5. Sorts by total duration (descending)
6. Takes the top 50 slow transactions

### 3.2 Service Dependency Analysis

This query analyzes service dependencies and their impact on performance.

```yaml
dataset: traces

steps:
- filter: start_time >= current_timestamp() - interval '24 hours'
- join:
    dataset: traces as child_spans
    where: traces.span_id == child_spans.parent_span_id
- group:
    by: [traces.service, child_spans.service]
    steps:
    - summarize:
        call_count: count(distinct traces.trace_id)
        avg_duration: avg(child_spans.duration)
        error_rate: 
          sum(case when child_spans.status == 'ERROR' then 1 else 0 end) 
          / count(child_spans.span_id) 
          * 100
- filter: call_count > 100
- sort: [traces.service, -call_count]
- generate:
    dependency_health:
      case:
      - error_rate > 5 || avg_duration > 1000: "Poor"
      - error_rate > 1 || avg_duration > 500: "Fair"
      - true: "Good"

into: service_dependencies
```

This query:

1. Filters traces from the last 24 hours
2. Joins parent spans with child spans
3. Groups by parent and child services
4. Calculates call count, average duration, and error rate
5. Filters dependencies with more than 100 calls
6. Sorts by parent service and call count
7. Generates a dependency health status based on error rate and average duration

## Conclusion

These examples demonstrate the power and flexibility of DUQL in analyzing developer observability data. From basic metric analysis to complex log pattern mining and distributed trace analysis, DUQL provides a intuitive and expressive way to gain insights from your data.

By combining different data sources and leveraging DUQL's advanced features like window functions, complex aggregations, and conditional logic, you can create comprehensive observability dashboards and alerting systems.

Remember that these queries can be easily modified and combined to suit your specific needs. Experiment with different time ranges, thresholds, and aggregations to discover hidden patterns and improve your application's performance and reliability.
