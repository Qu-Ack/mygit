# MYGIT


I always wanted to learn about the internals of git so I tried by making my own git 

this project is a demo of how git internally work, The tree objects, the blob objects, the commit objects all of those things 

it recreates the data structure that git uses

maybe I will expand this project's scope to branches in git as well and pushing to a platform like github



## STEPS TO SETUP LOCALLY 

```sh
go build -o out
./out 
```


## Documentation

To add the files in the staging area
```sh
    ./out staging <filepaths...> 
```

To commit the files in the staging area
```sh
    ./out commit
```

To See the contents of the hashed files
```sh
    ./out decomp-zlib <filepath>
```
