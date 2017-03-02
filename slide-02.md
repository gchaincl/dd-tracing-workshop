# OpenTracing Concepts

Vendor neutral open standard for distributed tracing.

- Trace: Main operation, collection of spans. Defined implicitly by their Spans
- Span: Represents a logical unit of work in the system.

```
Temporal relationships between Spans in a single Trace

––|–––––––|–––––––|–––––––|–––––––|–––––––|–––––––|–––––––|–> time

 [Span A···················································]
   [Span B··············································]
      [Span D··········································]
    [Span C········································]
         [Span E·······]        [Span F··] [Span G··] [Span H··]
```