# gotask

## goal

create an cli application for managing tasks in the terminal.

```
$ gotask
```

## requirements

should be able to perform crud operations via a cli on a data file of tasks. the operations should be as follows:

```
$ gotask list
$ gotask add <name> -d <description> -l <date-limit>
$ gotask complete <taskid>
$ gotask delete <taskid>
$ gotask done
```

### add

the add method should be used to create new tasks in the underlying data store.
it should take a positional argument with the task name, and description and date limit as optional arguments

```
$ gotask add <name> --description <description> --limit <date-limit>
or
$ gotask add <name> -dl <description> <date-limit>
```

for example:

```
$ gotask add "review Johh code" --description "pr: https://github.com/john/somerepo/pull/1" --limit "2024-12-15 18:00:000"
```

should add a new task to review John code until 18h00 from 15-12-24.

### list

this method should return a list of all of the **uncompleted** tasks, with the option to return all tasks regardless of whether or not they are completed.

for example:

```
$ gotask list
id  name         description      deadline    created
1   review pr    some pr link xd  in 21 days  a day ago
3   gym workout  back and biceps  in a month  27 minutes ago
```

or for showing all tasks, using a flag (such as -a or --all)

```
$ gotask list -a
id  name           description      deadline    expired  status   due_date   created
1   review pr      some pr link xd  in 21 days  false    pending  -          a day ago
2   clean my desk  -                in 4 years  false    done     a day ago  a day ago
3   gym workout    back and biceps  in a month  false    pending  -          27 minutes ago
```

### complete

to mark a task as done, add in the following method

```
$ gotask complete <taskid>
```

### delete

The following method should be implemented to delete a task from the data store

```
$ gotask delete <taskid>
```

### data file

additionally, must be possible to set tasks from csv:

```
ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false
```

**app idea based in dreamsofcode**
