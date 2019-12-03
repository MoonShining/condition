### Condition Match DSL

```
    c, err := NewCondition(`(false || a==1) && (b==2||c==3) && (d in ["hello"])`)
    env := run.NewEnvironment()
    env.Set("a", int64(1))
    env.Set("b", int64(4))
    env.Set("c", int64(3))
    env.Set("d", "hello")
    c.Match(env) // true
```

support operators
```
!=
==
>
<
>=
<=
in
```
