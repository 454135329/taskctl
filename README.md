# taskctl
Simple time manager

# Usage
Starting new task
```bash
$ taskctl start TSK-42
```

Stoping existing task
```bash
$ taskctl stop TSK-42
```

Complete existing task
```bash
$ taskctl done TSK-42
```

Listing logged time
```bash
$ taskctl list
+--------+-------------+-------------+
|  TASK  |   STATUS    | LOGGED TIME |
+--------+-------------+-------------+
| TSK-42 | In progress | 1 h 30 m    |
+--------+-------------+-------------+
```
