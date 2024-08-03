# Loops

The `loop` function in DUQL is used for iterative processing. It applies a sequence of steps repeatedly to an initial dataset until a termination condition is met, typically when the step function returns an empty table.

## Syntax

```yaml
loop:
- <step_1>
- <step_2>
# ...
```

## Parameters

The `loop` function doesn't have parameters itself, but it contains a list of steps that will be executed iteratively until the input (the results from the previous step) is empty.

| Component | Type  | Required | Description                                              |
| --------- | ----- | -------- | -------------------------------------------------------- |
| `steps`   | array | Yes      | A list of transformation steps to be applied iteratively |

Each step in the loop can be any valid DUQL transformation, such as `filter`, `generate`, `group`, etc.

## Examples

### Basic Iterative Calculation

```yaml
loop:
- filter: remaining_balance > 0
- generate:
    remaining_balance: remaining_balance - payment_amount
- append: payments_made
```

This loop continues to process payments until the remaining balance is zero or negative.

### Hierarchical Data Processing

```yaml
loop:
- join:
    dataset: employees
    where: managers.id == employees.manager_id
- generate:
    level: level + 1
- filter: not is_null(manager_id)
```

This loop traverses an employee hierarchy, joining the manager table with itself until it reaches employees with no manager (top level).

## Use Cases

1. **Recursive Calculations**: Perform calculations that depend on previous results, such as compound interest or depreciation.
2. **Hierarchical Data Processing**: Traverse tree-like structures, such as organizational hierarchies or bill of materials.
3. **Iterative Data Cleaning**: Apply data cleaning steps repeatedly until certain quality criteria are met.
4. **Convergence Algorithms**: Implement algorithms that iterate until a convergence condition is satisfied.

## Best Practices

1. ⚠️ Always include a termination condition to prevent infinite loops. This is typically done using a `filter` step that will eventually return an empty result.
2. 🔢 Consider adding a maximum iteration count as a safeguard against unexpected infinite loops.
3. 📊 Use `generate` steps within the loop to create or update variables that track the iteration progress or accumulate results.
4. 🧮 When possible, try to express your logic without loops for better performance. Only use loops when iterative processing is truly necessary.
5. 📝 Document the purpose and expected behavior of your loop clearly, especially for complex iterative processes.

## Related Functions

* [`filter`](../../basic/filter.md): Often used as a termination condition in loops
* [`generate`](../../intermediate/generate.md): Used to update variables within the loop
* [`append`](../../append.md): Useful for accumulating results across iterations

## Limitations and Considerations

* Loops can be computationally expensive, especially on large datasets. Use them judiciously.
* Not all database systems support iterative processing natively.
* Complex loops can be difficult to optimize. Consider alternative non-looping approaches if performance becomes an issue.

***

> 💡 **Tip:** While loops are powerful for certain types of problems, they should be used sparingly in data processing pipelines. Often, set-based operations are more efficient and easier to optimize. Always consider if there's a non-iterative way to express your logic before resorting to a loop!
