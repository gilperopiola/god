# 🛠️ ***God**ly Golang Toolset* - or **God** 🛠️

Handcrafted with love over the years, ***god*** is a *marvelous selection of useful and polished **Go Tools*** that any deity or backend engineer could ever need.

###### If you're looking for the frontend then talk to my son - *JSus*.

## 🧙 Why you should care about God

**God wants to provide as much utility as it can while removing all the complexity away from your codebase**. He abstracts that complexity upon himself, making your own code *shine bright like a 💎*. God makes your code so clean that even a child can understand it! 

And if a 3-year-old can understand your core business logic, then you can hire a 3-year-old to do the job and fire all those expensive, _sweaty_ senior devs that stench of catpiss. **God is good like that**. 😌

### 🙏 A Quick Example

**So** - `context.WithTimeout(context.Background(), d)` isn't bad, not at all. 

**But** - `god.NewCtxWithTimeout(d)` **is just better*** - It's shorter, cleaner and more expressive.

**I know** - This supposedly *godly* improvement doesn't change a thing on the scope of a big project. **I know**. 

### However, ***God's Way***™ is just way better.

---

Rome wasn't built in a day, neither was it built with a single huge brick. God stacks a lot of small improvements together to make a big difference.

**And that's why it's called God**: it should be all over your project to make a real impact, he can't do his job if he's not omnipresent. So better have God around and be happy about it :)


```
 |█████   |█████  |████
|█     █--█-----█--█---██---------> Golang.
|█       |█    |█  █     █
|█       |█     █  |     | 
||  |███  | |█  |--|-----█--------> Or.
|█     █ |█     █ |█     █
|█     █ |█    |█  |    █
 |█████   |█████  |████|----------> Death.
```

And if you're against importing code from a random dude on the internet, then you're right. You should be careful, that's why I'm giving you right now all the licensing rights to just copy the code and paste it in your project. *God doesn't even need credit, **he's just happy to help***.


## 🤩 Cool Examples 
  
*god.**ToIntAndErr**(s string, err error) (int, error)*

```go
  // - Before ⬇️
  userIDStr, err := ctx.Get("user_id")
  if err != nil {
    return nil, err
  }

  userID, err := strconv.Atoi(userIDStr)
  if err != nil {
    return nil, err
  }

  // - After ⬇️
  userID, err := god.ToIntAndErr(ctx.Get("user_id"))
  if err != nil {
    return nil, err
  }
```

```go
  // Todo: Add more cool examples
```

