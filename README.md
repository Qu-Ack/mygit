# MYGIT


I always wanted to learn about the internals of git so I tried by making my own git 


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
